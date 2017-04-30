package main

import "io"
import "os"
import "fmt"
import "net"
import "encoding/json"
import "net/textproto"
import "bufio"
import "time"
//import "strconv"

type Message struct {
    Name string
    Body string
    Time int64
}



func main () {
	
	fmt.Println("\033[H\033[2J")
	fmt.Println("Welcome to TicTacToe!")
	fmt.Println("Trying to connect to the server...")

	conn, err := net.Dial("tcp", "localhost:13000")
	if err != nil {
		fmt.Println("Dial's Error: ", err)
		os.Exit(1)
	}
	defer conn.Close()
	
	quit := false
	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)
	for{
		if quit{ break }

		choice := PrintInitialInterface()
		switch choice{
			case 1: // start a new game
				line, _ := tp.ReadLine()
			 	var m Message
				err = json.Unmarshal([]byte(line), &m)
				if err != nil{
					fmt.Println("Erro em descompactar o json")
				}
			 	fmt.Println(m)
			 	time.Sleep(time.Second * 2)
			case 2: //exit for
				quit = true
				break
			case 3:
				fmt.Println("Your choice isn't avaliable")
				time.Sleep(time.Second * 2)
				break
			default:
				fmt.Println("Your choice isn't avaliable")
				time.Sleep(time.Second * 2)
				break
		}
 	}
}

func mustCopy (dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)

	if err != nil {
		fmt.Println("Erro no mustCopy: ", err)
		os.Exit(1)
	}
}


func PrintInitialInterface() int{
	fmt.Println("\033[H\033[2J")
	fmt.Println("===== TIC TAC TOE =====")
	fmt.Println("(1) - New Game")
	fmt.Println("(2) - Exit")
	fmt.Println("=======================")
	var i int
	fmt.Print("Your choice -> ")
	_, err := fmt.Scanf("%d", &i)

	if err != nil{
		fmt.Println("Scanf error: ", err)
		return 3
	}
	return i
}