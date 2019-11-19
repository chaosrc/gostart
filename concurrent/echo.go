package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo() {
	server, err := net.Listen("tcp", ":8091")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Panicln(err)
			continue
		}
		go handleEchoConn(conn)
	}
}

func handleEchoConn(c net.Conn) {
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		go handleEcho(c, scanner.Text())
	}
	c.Close()
}

func handleEcho(c net.Conn, out string) {
	fmt.Fprintln(c, "\t", strings.ToUpper(out))
	time.Sleep(time.Second * 3)

	fmt.Fprintln(c, "\t", out)
	time.Sleep(time.Second * 3)

	fmt.Fprintln(c, "\t", strings.ToLower(out))
	time.Sleep(time.Second * 3)
}

func writeEcho(c net.Conn, s string) {
	defer c.Close()

	for len(s) > 1 {
		c.Write([]byte(s + "\n"))
		time.Sleep(time.Second * 1)

		s = s[len(s)/2:]
	}

	// c.Write([]byte(s + "\n"))
	// time.Sleep(time.Second * 1)

	// c.Write([]byte(s[len(s)/2:] + "\n"))
	// time.Sleep(time.Second * 1)

	// c.Write([]byte(s[len(s)/4:] + "\n"))
	// time.Sleep(time.Second * 1)

}
