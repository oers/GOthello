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
}

func makeBoard(moves string) (result Board){
    local := new(Board)  //TODO: find out why = new Board does not work
    initBoard(local)
    result = *local
	return 
}

func initBoard(board *Board){
}

func PrintBoard(board *Board){
}

func setStone(board *Board, row,column,stone uint8){
	if(stone == 0){
		board.blackStones = board.blackStones | (1 << (row *8 + column))
		board.whiteStones = board.whiteStones | ^(1 << (row *8 + column))
	} else {
	    board.blackStones = board.blackStones | ^(1 << (row *8 + column))
		board.whiteStones = board.whiteStones | (1 << (row *8 + column))
	}
}

func isEmpty(board *Board, row,column uint8) (bool){
	return ((board.whiteStones & board.blackStones) & (1 << (row *8 + column))) == 0
}

func isStone(board *Board, row,column,stone uint8) (bool){
	if(stone == 0){
		return (board.blackStones & (1 << (row *8 + column))) > 0
	} else {
		return (board.whiteStones & (1 << (row *8 + column))) > 0
	}
	panic("undefined state")
}





