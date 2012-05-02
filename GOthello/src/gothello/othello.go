package gothello

import (
	"fmt"
)

type Board struct{
	whiteStones uint64
	blackStones uint64
	blackTurn bool
}

func Hello() {
	fmt.Println("hello")
}

func Replay(moves string) {
	fmt.Println("Replay: " + moves)
	board := makeBoard()
	for len(moves) > 0 { //len is a system function which returns the length of a string
	    move := moves[0:2] //this is called slicing and creates a substring
	    MakeMove(board, move)
		moves = moves[2:len(moves)]
	}

	
}

func makeBoard() (result *Board){
    result = new(Board)  //TODO: find out why = new Board does not work
    initBoard(result)
	return 
}

func initBoard(board *Board){
        setStone(board, 4, 4, 0);//e5 --> White (Field Occupied)
        setStone(board, 3, 3, 0);//d4 --> White (Field Occupied)
        setStone(board, 3, 4, 1);//d5 --> Black (Field Occupied)
        setStone(board, 4, 3, 1);//e4 --> Black (Field Occupied)
}

func PrintBoard(board *Board){
}

func MakeMove(board *Board, move string) (result bool){
    fmt.Println(move)
	return false
}

//no Method overloading: http://golang.org/doc/go_faq.html#overloading
func makeMoveInt(board *Board, row, column int8) (result bool){

	return false
}

func setStone(board *Board, row,column,stone uint64){
	if(stone == 0){
		board.blackStones = board.blackStones | (1 << (row *8 + column)) //set black Stone
		board.whiteStones = board.whiteStones & ^(1 << (row *8 + column)) //unset white Stone (or leave unset)
	} else {
	    board.blackStones = board.blackStones & ^(1 << (row *8 + column)) //unset blackStone (or leave unset)
		board.whiteStones = board.whiteStones | (1 << (row *8 + column)) //set white Stone
	}
}

func isEmpty(board *Board, row,column uint64) (bool){
	return ((board.whiteStones & board.blackStones) & (1 << (row *8 + column))) == 0
}

func isStone(board *Board, row,column,stone uint64) (bool){
	if(stone == 0){
		return (board.blackStones & (1 << (row *8 + column))) > 0
	} else {
		return (board.whiteStones & (1 << (row *8 + column))) > 0
	}
	panic("undefined state")
}





