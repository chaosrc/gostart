package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

var (
	entering = make(chan client)
	leaveing = make(chan client)
	message  = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaveing:
			delete(clients, cli)
			close(cli)
		case msg := <-message:
			for cli := range clients {
				cli <- msg
			}
		}
	}

}

func main() {
	listen, err := net.Listen("tcp", ":9900")
	if err != nil {
		log.Fatal(err)
	}
	// 运行广播
	go broadcaster()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	name := conn.RemoteAddr().String()

	ch := make(chan string)
	entering <- ch
	message <- fmt.Sprintf("%s join\n", name)

	scanner := bufio.NewScanner(conn)

	// 监听 ch 事件将广播的 message 输出给客户端
	go func() {
		for msg := range ch {
			fmt.Fprintln(conn, msg)
		}
	}()

	// 读取客户端输入信息
	for scanner.Scan() {
		str := scanner.Text()
		message <- fmt.Sprintf("%s: %s\n", name, str)
	}

	leaveing <- ch
	conn.Close()
}
