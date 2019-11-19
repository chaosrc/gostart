package main

import (
	"log"
	"net"
	"time"
	"io"
)

func clockServer() {
	server, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalln("create server error")
	}

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Panicln(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("03:04:05\n"))
		if err != nil {
			return 
		}
		time.Sleep(time.Second * 1)
	}
}
