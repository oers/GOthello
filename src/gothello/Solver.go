package gothello

import (
	"fmt"
)

func Solve(moves string) {
	board := Replay(moves)
	board.SolveBoard()
}

func (board *Board) SolveBoard() {
	l := board.GetPossibleMoves()
	localResult := -9999
	var bestMove string
	//fmt.Println(board.ToString())
	for e := l.Front(); e != nil; e = e.Next() {
		 m := e.Value.(string)
		 //fmt.Println("Try", m)
		 result,_ := solveForMove(board.CopyOf(), m)
		 //fmt.Println(board.ToString(), result, m)
		 if isBetterResult(result, localResult, board.nextplayer) {
		 	localResult = result
		 	bestMove = m
		 }
	}	
	fmt.Println(board.ToString(), localResult, bestMove)
}
/*
* Recursive
*/
func solveForMove(board Board, move string) (localResult int, bestmove string) {
	board.Move(move)
	//fmt.Println(board.ToString())
	if board.IsFinished() {
		black, white := board.GetResult()
		return  black - white, move
	}
	l := board.GetPossibleMoves()
	localResult = -9999
	moveForPlayer := board.nextplayer
	for e := l.Front(); e != nil; e = e.Next() {
		 m := e.Value.(string)
		 result,_ := solveForMove(board, m)
		 //fmt.Println(board.ToString(), result, m)
		 if isBetterResult(result, localResult, moveForPlayer) {
		 	localResult = result
		 	bestmove = m
		 }
	}
	
	//fmt.Println(board.ToString(), localResult, bestmove)
	return localResult, bestmove
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

