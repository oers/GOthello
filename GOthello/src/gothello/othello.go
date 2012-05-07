package gothello

import (
	"fmt"
	"strconv"
	"strings"
	"container/list"
	"math/rand"
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
	possibleMoves    uint64
}

type direction struct {
	hor int
	ver int
}

//all directions
var dirs [8]direction = [8]direction{
	direction{ 0, 1},
	direction{ 0,-1},		
	direction{ 1, 1},	
	direction{-1,-1},
	direction{ 1, 0},
	direction{-1, 0},
	direction{-1, 1},
	direction{ 1,-1}}

func Hello() {
	fmt.Println("hello")
	//list := new(list.List)
}

func Replay(moves string) (board *Board){
	//fmt.Println("Replay: " + moves)
	board = MakeBoard()
	for len(moves) > 0  { //len is a system function which returns the length of a string
		move := moves[0:2] //this is called slicing and creates a substring
		moves = moves[2:len(moves)]
		
		result := board.Move(move)
		if !result {
		  fmt.Println("Move not Made:", move)
		  return
		}
	    //board.PrintBoard()
	}
	return
	


}

func MakeBoard() (result *Board) {
	result = new(Board)
	result.initBoard()
	return
}

func (board *Board) initBoard() { //define a method on a struct
	board.setStoneInt(4, 4, 1) //e5 --> White (Field Occupied)
	board.setStoneInt(3, 3, 1) //d4 --> White (Field Occupied)
	board.setStoneInt(3, 4, 0) //d5 --> Black (Field Occupied)
	board.setStoneInt(4, 3, 0) //e4 --> Black (Field Occupied)
	board.moves = list.New()
	board.markNextMoves()
}

/**
*
* PUBLIC METHODS
*
**/

func (board *Board) PrintBoard() {
	fmt.Println(board.ToString())
}

func (board *Board) ToString() (result string) {
	result = "\n_|a|b|c|d|e|f|g|h|\n"
	for i := 0; i < 8; i++ {
		result += strconv.Itoa(i+1) + "|"
		for j := 0; j < 8; j++ {
			if board.isPossibleMoveInt(i, j){
			   result += "o|"
			} else
			if board.isEmptyInt(i, j) {
				result += "_|"
			} else if board.isStoneInt(i, j, 0) {
				result += "b|"
			} else  {
				result += "w|"
			}
		}
		result += "\n"
	}
	return
}


func (board *Board) Move(move string) (result bool) {
	row, column := stringMoveToInt(move)
	return board.makeMoveInt(row, column)
}

func (board *Board) SetStoneUnsafe(field string, stone int) {
	row, column := stringMoveToInt(field)
	board.setStoneInt(row, column, stone)
}

func (board *Board) IsPossibleMove(field string) (bool) {
	row, column := stringMoveToInt(field)
	return board.isPossibleMoveInt(row, column)
}

func (board *Board) IsFinished() (bool) {
	return board.finished
}

func (board *Board) GetState(field string, stone int) (state int){
	row, column := stringMoveToInt(field)
	if board.isPossibleMoveInt(row, column){
		return 2; //possible move
	} else if board.isEmptyInt(row, column){
		return -1 //empty field
	} else if board.isStoneInt(row, column, 0){
		return 0 // black
	}
	return 1 //white
	
}

func (board *Board) MakeRandomMove(){

	for {
		move := rand.Int() % 64
		if board.isPossibleMoveInt(move/8, move %8) {
			board.makeMoveInt(move/8, move%8)
			return
		}
	}
}

func (board *Board) GetPossibleMoves() (moves *list.List){
    moves = list.New()  
	for i := 1; i <64; i++ {
		if board.isPossibleMoveInt(i/8, i%8) {
			moves.PushBack(intMoveToString(i/8, i%8))
		}
	}
	return
	
}

func (board *Board) GetResult() (black, white int){
	return board.blackStonesCount, board.whiteStonesCount
	
}

//no Method overloading: http://golang.org/doc/go_faq.html#overloading
func (board *Board) makeMoveInt(row, column int) (result bool) {
	if board.isPossibleMoveInt(row, column) { //next player means last player now
		flipped := board.flip(row, column)
		if !flipped {
		fmt.Println("Not Legal: ", row, column)
			return false
		}
		board.moves.PushBack(BoardMove{row, column})
	} else {
	    return false
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

func (board *Board) setStoneInt(row, column, stone int) {
	if stone == 0 {
		board.blackStones = board.blackStones | (1 << uint64(row*8+column))  //set black Stone
		board.whiteStones = board.whiteStones & ^(1 << uint64(row*8+column)) //unset white Stone (or leave unset)
	} else {
		board.blackStones = board.blackStones & ^(1 << uint64(row*8+column)) //unset blackStone (or leave unset)
		board.whiteStones = board.whiteStones | (1 << uint64(row*8+column))  //set white Stone
	}
}

func (board *Board) isEmptyInt(row, column int) bool {
	return ((board.whiteStones | board.blackStones) & (1 << uint64(row*8+column))) == 0
}

func (board *Board) isStoneInt(row, column, stone int) bool {
	if stone == 0 {
		return (board.blackStones & (1 << uint64(row*8+column))) > 0
	}
	return (board.whiteStones & (1 << uint64(row*8+column))) > 0
}

func (board *Board) isLegalMoveInt(row, column int) bool {
	return board.executeFlip(row, column, false) //don't flip just check  
}

func (board *Board) flip(row, column int) bool {
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

		//            log.info(nextRow + "/" + nextRow + "/" + isEmptyInt(nextRow, nextColumn) + "/" + isStoneInt(nextRow, nextColumn, toFlip) + "/" + isStoneInt(nextRow, nextColumn, endflip));
		if board.isStoneInt(nextRow, nextColumn, toFlip) { //the direction is right, stone of opposite colour in that direction
			//                if (log.isDebugEnabled()) {
			//                    log.debug("Flip candidate found for " + dir);
			//                }

			//can be flipped, if we can find a beginning i.e. stone of other colour in same direction
			for { //can't think of an appropiate recursion end right now
				//fmt.Println("Endless for", nextRow, nextColumn, dir.hor, dir.ver)
				nextRow = nextRow + dir.hor
				if nextRow <= -1 || nextRow >= 8 {
					break //at the end of the board
				}
				nextColumn = nextColumn + dir.ver
				if nextColumn <= -1 || nextColumn >= 8 {
					break //at the end of the board
				}
				//if we find an empty field break;
				if board.isEmptyInt(nextRow, nextColumn) {
					break
				}

				if board.isStoneInt(nextRow, nextColumn, endflip) { //found a stone of same colour, lines between can be flipped
					//                       if (log.isDebugEnabled()) {
					//                            log.debug("Possible Move found for Flip " + dir);
					//                        }
					if !executeFlip { //don't change the board, just check
						return true
					}

					for !(row == nextRow && column == nextColumn) { //backwards flipping, flip till we reach the start
						nextRow = nextRow - dir.hor
						nextColumn = nextColumn - dir.ver
						//fmt.Println("Endless recursive for", nextRow, nextColumn)
						//                            if (log.isDebugEnabled()) {
						//                                log.debug("Flipped: " + nextColumn + "/" + nextRow + " to " + endflip);
						//                            }

						//flip and count stones that are not already flipped
						if !board.isStoneInt(nextRow, nextColumn, endflip) {
							board.setStoneInt(nextRow, nextColumn, endflip)
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
	board.possibleMoves = 0 //unset everything
	for i := 0; i < 64; i++ {
	   if board.isEmptyInt(i/8, i%8) && board.isLegalMoveInt(i/8, i%8) {
	       board.possibleMoves = board.possibleMoves | (1 << uint64(i))
	       marked = true;
	   }
	}

	return marked
}

func (board *Board) isPossibleMoveInt(row, column int) (bool){

    //check that possibleMoves are marked
    if !board.finished && board.possibleMoves == 0 && !board.markNextMoves() {
    	panic("Unfinished Board has no possible moves")
    }
	return (board.possibleMoves & (1 << uint64(row*8 + column))) > 0;
}

func (board *Board) IsNextPlayerBlack() (bool){
	return board.nextplayer == 0;
}

func intMoveToString(row, column int) (move string){

   strColumn := strconv.Itoa((column + 1))
   switch row {
   	case 0: return "a" + strColumn
   	case 1: return "b" + strColumn
   	case 2: return "c" + strColumn
   	case 3: return "d" + strColumn
   	case 4: return "e" + strColumn
   	case 5: return "f" + strColumn
   	case 6: return "g" + strColumn
   	case 7: return "h" + strColumn
   }
   
   panic("row > 7 or row < 0")
}

func stringMoveToInt(move string) (row, column int){
	move = strings.ToLower(move)
	switch move[0:1] {
	case "a":
		column = 0
	case "b":
		column = 1
	case "c":
		column = 2
	case "d":
		column = 3
	case "e":
		column = 4
	case "f":
		column = 5
	case "g":
		column = 6
	case "h":
		column = 7
	}

	if column < 0 || column > 7 {
		panic(" " + string(row) + " is not in range 0 .. 7")
	}

	row, error := strconv.Atoi(move[1:2])
	if error != nil {
		panic(error)
	}

	row = row - 1 //make it zero based

	if row < 0 || row > 7 {
		panic(string(column) + " is not in range 0 .. 7")
	}
	//fmt.Println(move, row, column)
	
	return
}
