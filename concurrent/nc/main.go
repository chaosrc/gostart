package main

import (
	"log"
	"net"
	"io"
	"os"
)

func main() {
	 conn, err := net.Dial("tcp", ":8090")
	 if err != nil {
		 log.Fatal(err)
	 }
	 defer conn.Close()
	 if _, err := io.Copy(os.Stdout, conn); err != nil {
		 log.Fatal(err)
	 }
}