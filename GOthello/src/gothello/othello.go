package gothello

import (
	"fmt"
	"strconv"
	"strings"
	"container/list"
)

type BoardMove struct{
	row int
	column int
}

type Board struct {
	whiteStones      uint64
	blackStones      uint64
	whiteStonesCount int
	blackStonesCount int
	nextplayer       int //black = 0, white = 1
	finished         bool
	moves            *list.List
	possibleMoves    *list.List
}

type direction struct {
	hor int
	ver int
}

func Hello() {
	fmt.Println("hello")
	//list := new(list.List)
}

func Replay(moves string) {
	fmt.Println("Replay: " + moves)
	board := makeBoard()
	for len(moves) > 0 { //len is a system function which returns the length of a string
		move := moves[0:2] //this is called slicing and creates a substring
		board.MakeMove(move)
		moves = moves[2:len(moves)]
	}

}

func makeBoard() (result *Board) {
	result = new(Board)
	result.initBoard()
	result.PrintBoard()
	return
}

func (board *Board) initBoard() { //define a method on a struct
	board.setStone(4, 4, 0) //e5 --> White (Field Occupied)
	board.setStone(3, 3, 0) //d4 --> White (Field Occupied)
	board.setStone(3, 4, 1) //d5 --> Black (Field Occupied)
	board.setStone(4, 3, 1) //e4 --> Black (Field Occupied)
	board.moves = list.New()
	board.possibleMoves = list.New()
}

func (board *Board) PrintBoard() {
	result := "_|a|b|c|d|e|f|g|h|\n"
	for i := 0; i < 8; i++ {
		result += strconv.Itoa(i+1) + "|"
		for j := 0; j < 8; j++ {
			if board.isEmpty(i, j) {
				result += "_|"
			} else if board.isStone(i, j, 0) {
				result += "b|"
			} else {
				result += "w|"
			}
		}
		result += "\n"
	}
	fmt.Println(result)
}

func (board *Board) MakeMove(move string) (result bool) {
	move = strings.ToLower(move)
	//fmt.Println(move)

	var row int
	switch move[0:1] {
	case "a":
		row = 0
	case "b":
		row = 1
	case "c":
		row = 2
	case "d":
		row = 3
	case "e":
		row = 4
	case "f":
		row = 5
	case "g":
		row = 6
	case "h":
		row = 7
	}

	if row < 0 || row > 7 {
		panic(" " + string(row) + " is not in range 0 .. 7")
	}

	column, error := strconv.Atoi(move[1:2])
	if error != nil {
		panic(error)
	}

	column = column - 1 //make it zero based

	if column < 0 || column > 7 {
		panic(string(column) + " is not in range 0 .. 7")
	}

	return board.makeMoveInt(row, column)
}

//no Method overloading: http://golang.org/doc/go_faq.html#overloading
func (board *Board) makeMoveInt(row, column int) (result bool) {
	if board.IsLegalMove(row, column) { //next player means last player now
		flipped := board.Flip(row, column)
		if !flipped {
		fmt.Println("Not Legal: ", row, column)
			return false
		}
		board.moves.PushBack(BoardMove{row, column})
	}
	//first Move is Black
	if board.nextplayer == 0 {
		board.nextplayer = 1
		canMove := board.markNextMoves()
		if !canMove {
			board.nextplayer = 0
			canMove = board.markNextMoves()
			if canMove {
				//if(log.isDebugEnabled()) {log.debug("White has to skip");}

			} else {
				board.finished = true
				//if(log.isDebugEnabled()) {log.debug("End of Game");}
			}
		}

	} else {
		board.nextplayer = 0
		canMove := board.markNextMoves()
		if !canMove {
			board.nextplayer = 1
			canMove = board.markNextMoves()
			if canMove {
				//if(log.isDebugEnabled()) {log.debug("Black has to skip");}                  
			} else {
				board.finished = true
				//if(log.isDebugEnabled()) {log.debug("End of Game");}
			}
		}

	}

	//if(log.isDebugEnabled()) {log.debug("NextPlayer: " + (nextPlayerBlack?"black":"white"));};

	return true
}

func (board *Board) setStone(row, column, stone int) {
	if stone == 0 {
		board.blackStones = board.blackStones | (1 << uint64(row*8+column))  //set black Stone
		board.whiteStones = board.whiteStones & ^(1 << uint64(row*8+column)) //unset white Stone (or leave unset)
	} else {
		board.blackStones = board.blackStones & ^(1 << uint64(row*8+column)) //unset blackStone (or leave unset)
		board.whiteStones = board.whiteStones | (1 << uint64(row*8+column))  //set white Stone
	}
}

func (board *Board) isEmpty(row, column int) bool {
	return ((board.whiteStones | board.blackStones) & (1 << uint64(row*8+column))) == 0
}

func (board *Board) isStone(row, column, stone int) bool {
	if stone == 0 {
		return (board.blackStones & (1 << uint64(row*8+column))) > 0
	}
	return (board.whiteStones & (1 << uint64(row*8+column))) > 0
}

func (board *Board) IsLegalMove(row, column int) bool {
	return board.executeFlip(row, column, false) //don't flip just check  
}

func (board *Board) Flip(row, column int) bool {
	return board.executeFlip(row, column, true)
}

/**
 * Flips all Stones on the board for the given move.
 * @param board
 * @param row the row of the move
 * @param column the column of the move
 * @param endflip the desired stone that makes the flip
 * @param executeFlip if false then only a check is performed, no actual flipping is done
 * @return 
 */
func (board *Board) executeFlip(row, column int, executeFlip bool) bool {
	toFlip := (board.nextplayer - 1) * (board.nextplayer - 1) //1*1=1; -1*-1 = 1, 0*0 = 0
	endflip := board.nextplayer

	dirs := [10]direction{
	direction{0,1},
	direction{0,-1},		
	direction{1,1},	
	direction{-1,-1},
	direction{1,0},
	direction{-1,0},
	direction{-1,1},
	direction{1,-1}}

	flipped := 0
	//look for flip in every direction
	for _, dir := range dirs { //first value is index in array, 2nd value is value in value
		nextRow := row + dir.hor

		//we must jump over at least one stone, so certain flips don't need to be evaluated
		//if the next stone in direction is on directly on the edge rows/columns skip evaluation

		//read as: if this or the next move in the same vertical direction will be out of bounds --> ignore
		if nextRow+dir.hor < 0 || nextRow+dir.hor > 7 {
			continue //at the end of the board
		}
		nextColumn := column + dir.ver

		//read as: if this or the next move in the same horizontal direction will be out of bounds --> ignore
		if nextColumn+dir.ver < 0 || nextColumn+dir.ver > 7 {
			continue //at the end of the board
		}

		//            log.info(nextRow + "/" + nextRow + "/" + isEmpty(nextRow, nextColumn) + "/" + isStone(nextRow, nextColumn, toFlip) + "/" + isStone(nextRow, nextColumn, endflip));
		if board.isStone(nextRow, nextColumn, toFlip) { //the direction is right, stone of opposite colour in that direction
			//                if (log.isDebugEnabled()) {
			//                    log.debug("Flip candidate found for " + dir);
			//                }

			//can be flipped, if we can find a beginning i.e. stone of other colour in same direction
			for { //can't think of an appropiate recursion end right now
				nextRow = nextRow + dir.hor
				if nextRow == -1 || nextRow == 8 {
					break //at the end of the board
				}
				nextColumn = nextColumn + dir.ver
				if nextColumn == -1 || nextColumn == 8 {
					break //at the end of the board
				}
				//if we find an empty field break;
				if board.isEmpty(nextRow, nextColumn) {
					break
				}

				if board.isStone(nextRow, nextColumn, endflip) { //found a stone of same colour, lines between can be flipped
					//                       if (log.isDebugEnabled()) {
					//                            log.debug("Possible Move found for Flip " + dir);
					//                        }
					if !executeFlip { //don't change the board, just check
						return true
					}

					for !(row == nextRow && column == nextColumn) { //backwards flipping, flip till we reach the start
						nextRow = nextRow - dir.hor
						nextColumn = nextColumn - dir.ver

						//                            if (log.isDebugEnabled()) {
						//                                log.debug("Flipped: " + nextColumn + "/" + nextRow + " to " + endflip);
						//                            }

						//flip and count stones that are not already flipped
						if !board.isStone(nextRow, nextColumn, endflip) {
							board.setStone(nextRow, nextColumn, endflip)
							flipped++
						}
					}
					break
				}
			}
		}

		if flipped == 0 {
			//                if (log.isDebugEnabled()) {
			//                    log.debug(dir + " did not flip");
			//                }
		}
	}

	if executeFlip {
		if toFlip == 0 { //White
			board.blackStonesCount += flipped
			board.whiteStonesCount -= (flipped - 1) //flipped contains the new stone
		} else {
			board.blackStonesCount -= (flipped - 1) //flipped contains the new stone
			board.whiteStonesCount += flipped
		}
	}

	return flipped > 0
}

func (board *Board) markNextMoves() bool {
	marked := false
	board.possibleMoves = list.New()
	for i := 0; i < 64; i++ {
	   if board.isEmpty(i/8, i%8) && board.IsLegalMove(i/8, i%8) {
	       board.possibleMoves.PushBack(BoardMove{i/8, i%8});
	       marked = true;
	   }
	}

	return marked
}
