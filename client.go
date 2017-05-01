package main

import(
	 "io"
	 "os"
	 "fmt"
	 "net"
	 "encoding/json"
	 "net/textproto"
	 "bufio"
	 "time"
	 "./extfiles"
	 "math"
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
// ERROR -> Id Problem
// DEFAULT -> do nothing


func main () {
	
	fmt.Println("\033[H\033[2J")
	fmt.Println("Welcome to TicTacToe!")
	fmt.Println("Trying to connect to the server...")
	board_example := make([]int, 9)
	for i:=0; i<len(board_example); i++{
		board_example[i] = i
	}
	conn, err := net.Dial("tcp", "localhost:13000")
	if err != nil {
		fmt.Println("Dial's Error: ", err)
		os.Exit(1)
	}
	defer conn.Close()
	
	quit := false
	board := tictactoe.Board{make([]int, 3*3), 1, 3}
	for{
		if quit{ break }
		quit2:= false
		choice := PrintInitialInterface()
		board.SetBoard()
		switch choice{
			case 1: // start a new game
				for{
					if quit2{break}
					choice2 := PrintSecondInterface(board.Board)
					switch choice2{
						case 1: //Make a play
							play := PrintThirdInterface(board.Board, board_example)
							if play == 9{
								fmt.Println("Invalid choice!")
							}else{
								if((board.IsThisPlayAvaliable(board.DesConvertPlay(play))) && (ServerCheckPlay(conn, play, board.Board))){
									board.SetPlay(play)
									if(board.IsWinner(board.DesConvertPlay(play))){
										Someonehaswon(board.GetPlayer(), board.Board)
										quit2 = true
									}else if(board.IsBoardFull()){
										DrawGame(board.Board)
										quit2 = true
									}else{ //Computer plays
										board.ChangePlayer() //Change to computer
										play2 := ComputerIsPlaying(board, conn)
										if(play2 == 9){
											fmt.Println("Connect problems.. exiting")
											time.Sleep(time.Second * 2)
											quit2 = true
											quit = true
											ExitServer(conn)
											break
										}
										board.SetPlay(play2)
										if(board.IsWinner(board.DesConvertPlay(play2))){
											Someonehaswon(board.GetPlayer(), board.Board)
											quit2 = true
										}else if(board.IsBoardFull()){
											DrawGame(board.Board)
											quit2 = true
										}else{
											board.ChangePlayer() //Change to client
											break;
										}
									}
								}else{
									fmt.Println("Your play isn't avaliable! Try another!")
									time.Sleep(time.Second * 2)
									break;
								}
							}
						case 2: 
							quit2 = true
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
			case 2: //exit for
				ExitServer(conn)
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
	fmt.Println("===== TIC TAC TOE =====\n")
	fmt.Println("(1) - New Game")
	fmt.Println("(2) - Exit")
	fmt.Println("\n=======================")
	var i int
	fmt.Print("Your choice -> ")
	_, err := fmt.Scanf("%d", &i)

	if err != nil{
		fmt.Println("Scanf error: ", err)
		return 3
	}
	return i
}

func PrintSecondInterface(board []int) int{
	fmt.Println("\033[H\033[2J")
	fmt.Println("===== GAME =====\n")
	fmt.Println("-- Actual Board --")
	NormalizeBoard(board)
	fmt.Println("(1) - Make a play")
	fmt.Println("(2) - Back (will reset the game)")
	fmt.Println("\n================")
	var i int
	fmt.Print("Your choice -> ")
	_, err := fmt.Scanf("%d", &i)

	if err != nil{
		fmt.Println("Scanf error: ", err)
		return 3
	}
	return i
}

func PrintThirdInterface(board []int, example []int) int{
	fmt.Println("\033[H\033[2J")
	fmt.Println("===== MAKE A PLAY =====\n")
	fmt.Println("Make a play using the numbers below: ")
	fmt.Println("-- Example Board --")
	NormalizeBoard(example)
	fmt.Println("-- Actual Board --")
	NormalizeBoard(board)
	var i int
	fmt.Print("Your choice -> ")
	fmt.Println("\n=========================")
	_, err := fmt.Scanf("%d", &i)

	if err != nil{
		fmt.Println("Scanf error: ", err)
		return 9
	}
	if ((i < 0) || (i > 8)){
		return 9
	}
	return i	
}

func NormalizeBoard(board []int){
	tam := int(math.Sqrt(float64(len(board))))
	x:= make([]int, tam)
	y:=make([]int, tam)
	z:=make([]int, tam)
	for i:= 0; i< tam; i++{
		x[i] = board[i]
		y[i] = board[i + 3]
		z[i] = board[i + 6]
	}
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func Someonehaswon(player int, board []int){
	fmt.Println("\033[H\033[2J")
	fmt.Println("===== TIC TAC TOE =====\n")
	fmt.Printf("Player %d has won!!\n", player)
	NormalizeBoard(board)
	fmt.Println("=========================")
	time.Sleep(time.Second * 2)
}

func DrawGame(board []int){
	fmt.Println("\033[H\033[2J")
	fmt.Println("===== DRAW GAME =====\n")
	NormalizeBoard(board)
	fmt.Println("=======================")
	time.Sleep(time.Second * 2)	
}

func ComputerIsPlaying(board tictactoe.Board, conn net.Conn) int{
	fmt.Println("\033[H\033[2J")
	fmt.Println("===== Computer is Playing =====\n")
	fmt.Println("=================================")
	
	//Creating json
	avaliableplays := board.AvaliablePlays()
	if len(avaliableplays) <= 0{
		return 9
	}
	m2 := MessageTTT{"MakeAPlay", avaliableplays, -1}
	b, _ := json.Marshal(m2)
	//Sending the json to the server with an array of avaliable plays
	_, erro := conn.Write(b)
	conn.Write([]byte("\n"))
	if erro != nil{
		fmt.Println("Problemas ao enviar o json")
		return 9
	}

	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)

	//Recieve the servers play
	line, _ := tp.ReadLine()
	//Unmarshal the json
 	var m MessageTTT
	err2 := json.Unmarshal([]byte(line), &m)
	if err2 != nil{
		fmt.Println("Erro em descompactar o json")
		return 9
	}
	time.Sleep(time.Second * 1)
	return m.Play
}

func ExitServer(conn net.Conn){
	m := MessageTTT{"Exit", make([]int, 0), -1}
	b, _ := json.Marshal(m)
	_, erro := conn.Write(b)
	conn.Write([]byte("\n"))
	if erro != nil{
		fmt.Println("Problemas ao enviar o json")
	}

	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)

	//Recieve the servers play
	line, _ := tp.ReadLine()
	//Unmarshal the json
 	var m2 MessageTTT
	err2 := json.Unmarshal([]byte(line), &m2)
	if err2 != nil{
		fmt.Println("Erro em descompactar o json")
	}
	if(m2.ID == "Exit"){
		fmt.Println("Closing connection...")
	}
}

func ServerCheckPlay(conn net.Conn, play int, board []int) bool{
	m := MessageTTT{"Checkplay", board, play}
	b, _ := json.Marshal(m)
	_, erro := conn.Write(b)
	conn.Write([]byte("\n"))
	if erro != nil{
		fmt.Println("Problemas ao enviar o json")
	}

	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)

	//Recieve the servers play
	line, _ := tp.ReadLine()
	//Unmarshal the json
 	var m2 MessageTTT
	err2 := json.Unmarshal([]byte(line), &m2)
	if err2 != nil{
		fmt.Println("Erro em descompactar o json")
	}
	if(m2.ID == "CheckTrue"){
		return true
	}
	return false
}