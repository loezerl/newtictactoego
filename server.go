package main

import (
	"os"
	"fmt"
	"net"
	"encoding/json"
)

type Message struct {
    Name string
    Body string
    Time int64
}

func main () {
	
	listener, err := net.Listen("tcp", "localhost:13000")
	if err != nil {
		fmt.Println("Erro no Listen: ", err)
		os.Exit(1)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erro no Accept: ", err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn (conn net.Conn) {
	defer conn.Close()
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, _ := json.Marshal(m)
	fmt.Println(b)
	conn.Write(b)
}