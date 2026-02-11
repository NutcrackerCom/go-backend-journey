package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/NutcrackerCom/go-backend-journey/mini-projects/mini-chat/internal/protocol"
)

func handleConn(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	defer conn.Close()
	for scanner.Scan() {
		line := scanner.Text()
		pars, err := protocol.ParseLine(line)
		if err != nil {
			fmt.Fprintln(conn, err)
		} else {
			if pars.Kind == protocol.KindCommand {
				fmt.Fprintf(conn, "cmd: %s   arg: %s\n", pars.Cmd, pars.Arg)
			} else if pars.Kind == protocol.KindMessage {
				fmt.Fprintf(conn, "msg: %s\n", pars.Text)
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

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}
}
