package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"

	"github.com/NutcrackerCom/go-backend-journey/mini-projects/mini-chat/internal/protocol"
)

type Client struct {
	conn net.Conn
	out  chan string
	mu   sync.RWMutex
	name string
}

type whoReq struct {
	client *Client
}

type message struct {
	msg string
	who *Client
}

type renameReq struct {
	client *Client
	name   string
}

type Hub struct {
	join      chan *Client
	leave     chan *Client
	broadcast chan message
	clients   map[*Client]string
	who       chan whoReq
	rename    chan renameReq
}

func NewHub() *Hub {
	return &Hub{
		join:      make(chan *Client, 5),
		leave:     make(chan *Client, 5),
		broadcast: make(chan message, 5),
		clients:   make(map[*Client]string),
		who:       make(chan whoReq, 5),
		rename:    make(chan renameReq, 5),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case c := <-h.join:
			c.mu.RLock()
			h.clients[c] = c.name
			name := c.name
			c.mu.RUnlock()
			for cl := range h.clients {
				if cl == c {
					continue
				}
				select {
				case cl.out <- fmt.Sprintf("%s joined", name):
				default:
				}
			}
		case c := <-h.leave:
			c.mu.RLock()
			name := c.name
			c.mu.RUnlock()
			for cl := range h.clients {
				if cl == c {
					continue
				}
				select {
				case cl.out <- fmt.Sprintf("%s left", name):
				default:
				}
			}
			delete(h.clients, c)
			close(c.out)
			c.conn.Close()
		case msg := <-h.broadcast:
			for c := range h.clients {
				if c == msg.who {
					continue
				}
				select {
				case c.out <- msg.msg:
				default:
				}
			}
		case name := <-h.rename:
			name.client.mu.Lock()
			h.clients[name.client] = name.name
			name.client.mu.Unlock()
		case req := <-h.who:
			var response []string
			for c := range h.clients {
				c.mu.RLock()
				response = append(response, c.name)
				c.mu.RUnlock()
			}
			select {
			case req.client.out <- "online: " + strings.Join(response, ", "):
			default:
			}
		}
	}
}

func clientWriter(c *Client) {
	for msg := range c.out {
		_, err := fmt.Fprintln(c.conn, msg)
		if err != nil {
			return
		}
	}
}

func clientReader(c *Client, h *Hub) {
	scanner := bufio.NewScanner(c.conn)
	defer func() { h.leave <- c }()
	for scanner.Scan() {
		line := scanner.Text()
		p, err := protocol.ParseLine(line)
		if err != nil {
			select {
			case c.out <- fmt.Sprintf("error: %v", err):
			default:
			}
			continue
		}
		switch p.Kind {
		case protocol.KindMessage:
			c.mu.RLock()
			name := c.name
			c.mu.RUnlock()
			h.broadcast <- message{msg: fmt.Sprintf("%s: %s", name, p.Text), who: c}
		case protocol.KindCommand:
			if p.Cmd == "quit" {
				select {
				case c.out <- "bye":
				default:
				}
				return
			}
			if p.Cmd == "name" {
				c.mu.Lock()
				c.name = p.Arg
				c.mu.Unlock()
				select {
				case h.rename <- renameReq{client: c, name: c.name}:
				default:
				}
			}
			if p.Cmd == "who" {
				h.who <- whoReq{client: c}
			}
		}
	}

}

func main() {
	addr := flag.String("addr", ":9091", "listen address")
	flag.Parse()
	ln, err := net.Listen("tcp", *addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "can not connect. Error: %v", err)
		return
	}
	defer ln.Close()
	hub := NewHub()
	go hub.Run()
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		client := &Client{conn: conn, out: make(chan string, 20), name: "anonymous"}
		hub.join <- client
		go clientWriter(client)
		go clientReader(client, hub)
	}
}
