package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:9091", "server address")
	flag.Parse()
	conn, err := net.Dial("tcp", *addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "can not connect. Error: %v", err)
	}
	defer conn.Close()
	fmt.Printf("connected to %s\n", *addr)

	go func() {
		in := bufio.NewScanner(os.Stdin)
		for in.Scan() {
			line := in.Text()

			_, err := fmt.Fprintln(conn, line)
			if err != nil {
				return
			}
		}

		_ = conn.Close()
	}()

	out := bufio.NewScanner(conn)
	for out.Scan() {
		fmt.Println(out.Text())
	}

	if err := out.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "read error: %v\n", err)
	}
	fmt.Println("disconnected")
}
