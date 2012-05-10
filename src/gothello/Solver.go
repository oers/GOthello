package gothello

import (
	"fmt"
	"strconv"
	"runtime"
)

var c = make(chan int, 3)

type MoveAnalysis struct{
	rating int
	move string
	nextMove *MoveAnalysis
}

func Solve(moves string) {
	fmt.Println("Moves to solve:" , 60 - len(moves)/2)
	board := Replay(moves)
	board.SolveBoard()
}

func (board *Board) SolveBoard() { //this method is allowed to work on Strings, for a nice start
	runtime.GOMAXPROCS(runtime.NumCPU())
	l := board.GetPossibleMoves()
	ana := new(MoveAnalysis)
	ana.rating = -9999
	fmt.Println("Top",l.Len())
	//fmt.Println(board.ToString())
	
	i := l.Len()

	for e := l.Front(); e != nil; e = e.Next() {
		go execute(e.Value.(string), board)
	}	
	
	for j:=0; j < i; j++ {
		<- c
	}
	l = nil
	//fmt.Println(board.ToString(), ana.ToString())
}

func execute(move string, board *Board) {
			 m := move
			 //fmt.Println("Try", m)
			 row, column := stringMoveToInt(m)
			 fmt.Println("Start", m)
			 result := solveForMove(board.CopyOf(), row, column)
			 fmt.Println("Top",m, result.ToString())
//			 //fmt.Println(board.ToString(), result, m)
//			 if isBetterResult(result.rating, ana.rating, board.nextplayer) {
//			 	ana.rating = result.rating
//			 	ana.move = m
//			 	ana.nextMove = result
//			 }
			 c <- 1
		 }

/*
* Recursive, works only on int interpretation
*/
func solveForMove(board Board, row, column int) (analysis *MoveAnalysis) {
	
	ana := new(MoveAnalysis)
	ana.rating = -9999
	board.makeMoveInt(row, column)
	//fmt.Println(board.ToString())
	if board.IsFinished() {
		black, white := board.GetResult()
		analysis = new(MoveAnalysis)
		analysis.move = intMoveToString(row, column)
		analysis.rating = black -  white
		return 
	}
	analysis = new(MoveAnalysis)
	analysis.rating = -9999
	
	moveForPlayer := board.nextplayer
	for i:= 0; i< 64; i++ {
		 if !board.isPossibleMoveInt(i/8, i%8) {
		 	continue
		 }
		 result:= solveForMove(board, i/8, i%8)
		 //fmt.Println("Deeper", result.ToString())
		 //fmt.Println(board.ToString(), result, i/8, i%8)
		 if isBetterResult(result.rating, analysis.rating, moveForPlayer) {
		 	analysis.rating = result.rating
		 	analysis.move = intMoveToString(i/8, i%8)
		 	analysis.nextMove = result
		 } 
	}
	//fmt.Println(board.ToString(), localResult, bestmove)
	return
}

func isBetterResult(newresult, oldresult, stone int) (bool) {

    if oldresult == -9999 {
    	return true
    }
    
	if stone == 0 {
		return newresult > oldresult
	} 
	
	return newresult < oldresult
}

func (analysis *MoveAnalysis) ToString() (result string) {
	result += analysis.move + "(" + strconv.Itoa(analysis.rating) + ") "
	if analysis.nextMove != nil {
		result += "-" + analysis.nextMove.ToString()
	}
	return
}
