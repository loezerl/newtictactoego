package tictactoe

//player1 (client) = 1
//player2 (server) = 2
//empty cell = 0
		  //0, 1, 2, 3, 4, 5, 6, 7, 8	-> positions
// Board = [0, 0, 0, 0, 0, 0, 0, 0, 0]

// Board = [0, 1, 2]
		// [3, 4, 5]
		// [6, 7, 8]

import "fmt"


type Board struct{
	Board []int
	Player int
	Size int
}
func(b Board) GetBoard(){
	fmt.Println(b.Board)
}//ok
//Reset the board | Start a new game
func(b *Board) SetBoard(){
	for i := 0; i < len(b.Board); i++{
		b.Board[i] = 0
	}
	b.Player = 1
}//ok
func(b Board) SetPlay(play int){
	b.Board[play] = b.Player
}//ok
func(b *Board) ChangePlayer(){
	if(b.Player == 1){
		b.Player = 2
	}else{
		b.Player = 1
	}
}//ok 


func(b Board) GetPlayer() int{
	return b.Player
}//ok

//Verify if the column is complete with the actual player
func(b Board) VerifyColumn(column int) bool{
	for i:=0; i< len(b.Board); i+=3{
		if(!(b.Board[i + column] == b.Player)){
			return false
		}
	}
	return true
}//OK

//Verify if the row is complete with the actual player
func(b Board) VerifyRow(row int) bool{
	//Convert 0, 1, 2 -> 0, 3, 6
	initial := row * 3
	final := initial + 2
	for i:= initial; i<= final; i++{
		if(!(b.Board[i] == b.Player)){
			return false
		}
	}
	return true
} //OK
// Diagonal sentido horario -> 0 4 8
func(b Board) IsDH(cord int) bool{
	if(cord == 0){
		return true
	}else if(cord == 4){
		return true
	}else if(cord == 8){
		return true
	}else{
		return false
	}
}//ok

// Diagonal anti horario -> 2 4 6
func(b Board) IsDAH(cord int) bool{
	if(cord == 2){
		return true
	}else if(cord == 4){
		return true
	}else if(cord == 6){
		return true
	}else{
		return false
	}	
}//ok

func(b Board) IsWinner(x, y int) bool{
	play := b.ConvertPlay(x, y)
	if(b.VerifyRow(x)){
		return true
	}else if(b.VerifyColumn(y)){
		return true
	}
	if(b.IsDH(play) || b.IsDAH(play)){
		contdh := 0
		contdah := 0
		for i:= 0; i< len(b.Board); i++{
			if(b.IsDH(i) && b.Board[i] == b.Player){
				contdh++
			}
			if(b.IsDAH(i) && b.Board[i] == b.Player){
				contdah++
			}
		}
		if(contdh == b.Size){
			return true
		}else if(contdah == b.Size){
			return true
		}
	}
	return false
}//ok

func(b Board) ConvertPlay(row, column int) int{
	x := [3]int{0, 1, 2}
	y := [3]int{3, 4, 5}
	z := [3]int{6, 7, 8}

	if(row == 0){
		return x[column]
	}else if(row == 1){
		return y[column]
	}else{
		return z[column]
	}
}//ok

func(b Board) AvaliablePlays() []int{
	v := make([]int, 0)
	for i:=0; i < len(b.Board); i++{
		if(b.Board[i] == 0){
			v = append(v, i)
		}
	}
	return v
}//ok

func(b Board) IsThisPlayAvaliable(row, column int) bool{
	if((0<=row && row<=2)&&(0<=column && column<=2)){
		x := b.ConvertPlay(row, column)
		if(b.Board[x] == 0){
			return true
		}
	}
	return false
}

func(b Board) IsBoardFull() bool{
	for i:= 0; i< len(b.Board); i++{
		if(b.Board[i] == 0){
			return false
		}
	}
	return true
}

func(b Board) DesConvertPlay(play int) (int, int){
	switch play{
		case 0: return 0,0
		case 1: return 0,1
		case 2: return 0,2
		case 3: return 1,0
		case 4: return 1,1
		case 5: return 1,2
		case 6: return 2,0
		case 7: return 2,1
		case 8: return 2,2
	}
	return 0,0
}