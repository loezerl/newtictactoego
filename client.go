package main

import "io"
import "os"
import "fmt"
import "net"
import "encoding/json"
import "net/textproto"
import "bufio"

type Message struct {
    Name string
    Body string
    Time int64
}



func main () {
	conn, err := net.Dial("tcp", "localhost:13000")
	if err != nil {
		fmt.Println("Erro no Dial: ", err)
		os.Exit(1)
	}
	defer conn.Close()
	
	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)
	line, _ := tp.ReadLine()
 	var m Message
	err = json.Unmarshal([]byte(line), &m)
	if err != nil{
		fmt.Println("Erro em descompactar o json")
	}
 	fmt.Println(m)
}

func mustCopy (dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)

	if err != nil {
		fmt.Println("Erro no mustCopy: ", err)
		os.Exit(1)
	}
}