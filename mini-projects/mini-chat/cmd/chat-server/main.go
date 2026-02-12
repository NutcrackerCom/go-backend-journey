package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/NutcrackerCom/go-backend-journey/mini-projects/mini-chat/internal/protocol"
)

type Client struct {
	conn net.Conn
	out  chan string
	mu   sync.RWMutex
	name string
}

type Hub struct {
	join      chan *Client
	leave     chan *Client
	broadcast chan string
	clients   map[*Client]bool
}

func NewHub() *Hub {
	return &Hub{
		join:      make(chan *Client, 5),
		leave:     make(chan *Client, 5),
		broadcast: make(chan string, 5),
		clients:   make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case c := <-h.join:
			h.clients[c] = true
		case c := <-h.leave:
			delete(h.clients, c)
			close(c.out)
			c.conn.Close()
		case msg := <-h.broadcast:
			for c := range h.clients {
				select {
				case c.out <- msg:
				default:
				}
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
			h.broadcast <- fmt.Sprintf("%s: %s", name, p.Text)
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
