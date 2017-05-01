package main

import (
	"os"
	"fmt"
	"net"
	"encoding/json"
	"net/textproto"
	"bufio"
)

type Message struct {
    Name string
    Body string
    Time int64
}


type MessageTTT struct{
	ID string
	Array []int //default [array of zeros]
	Play int //default -1
}
// IDs
// Checkplay -> will check if the client's play is correct
// CheckTrue -> Client's play True
// CheckFalse -> Client's play False
// MakeAPlay -> Server's play
// GameOver -> TicTacToe | Draw
// Exit -> ConClose on servers side
// ERROR -> ID problem
// DEFAULT -> do nothing



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
	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)
	exit := false
	var ID string
	var Play int
	Resp:= MessageTTT{"DEFAULT", make([]int, 9), -1}

	for{
		if(exit){break}
		fmt.Println("Waiting for client to send the json..")
		line, _ := tp.ReadLine()
		fmt.Println("Client's json recieved")
	 	var m MessageTTT
		err := json.Unmarshal([]byte(line), &m)
		if err != nil{
			fmt.Println("Erro em descompactar o json")
		}else{
			fmt.Println("ID: ", m.ID)
			if(m.ID == "Checkplay"){
				if(Checkplay(m.Array, m.Play)){
					ID = "CheckTrue"
				}else{
					ID = "CheckFalse"
				}
				Play = -1
			}else if(m.ID == "MakeAPlay"){ //Recieve a array of avaliable plays
				ID = m.ID
				Play = m.Array[0]
			}else if(m.ID == "Exit"){
				exit = true
				fmt.Println("Closing connection...")
				ID = m.ID
				Play = m.Play
			}else if(m.ID == "GameOver"){
				ID = m.ID
				Play = m.Play
				fmt.Println("Someone has won or draw the game!")
			}else{
				fmt.Println("This ID isn't avaliable!")
				ID = "ERROR"
				Play = m.Play
			}
			Resp.ID = ID
			Resp.Array = m.Array
			Resp.Play = Play
			b, _ := json.Marshal(Resp)
			//Sending json message
			conn.Write(b)
			conn.Write([]byte("\n"))
		}

	}
	fmt.Println("Connection closed.")
}


//fazer a funcao de checar a jogada
func Checkplay(board []int, play int) bool{
	if((play >= 0) && (play <=8)){
		if(board[play] == 0){
			return true
		}
	}
	return false
}