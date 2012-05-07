package main 

import (
   "fmt"
   "gothello"
   "math/rand"
   "time"
)

func main() {
   fmt.Println("GOthello")
   fmt.Println("How would you like to play?")
   fmt.Println("Start as (b)lack or (w)hite, (a)utomatic or (t)wo players?")
   var input string
   fmt.Scanln(&input)
   
   var automatic bool
   var twoPlayer bool
   var blackStart bool
   
   if input[0:1] == "b" {
      blackStart = true
      automatic = false
   }
   
   if input[0:1] == "a" {
      blackStart = true
      automatic = true
   }
   
   if input[0:1] == "w" {
      blackStart = false
      automatic = false
   }
   
   if input[0:1] == "t" {
      blackStart = true
      automatic = false
   }
   
   rand.Seed(time.Now().UnixNano()) //set the seed
   
   board := gothello.MakeBoard()
   
   if(!blackStart || automatic) {
   		board.MakeRandomMove()
   }
    
   board.PrintBoard() 
    
   for !board.IsFinished() {
		if automatic  {
			board.MakeRandomMove()
			board.PrintBoard()
		} else if twoPlayer {
			fmt.Println("Make a Move:")
		   	var move string
		   	fmt.Scanln(&move)
		   	fmt.Println(move)
		   	if !board.Move(move) { 
		   		fmt.Println("Move was illegal. Try again.")
		   	}
		} else {			
		   	fmt.Println("Your Move:")
		   	var move string
		   	fmt.Scanln(&move)
		   	fmt.Println(move)
		   	if board.Move(move) { //returns false if move was not right/possible
			   	board.PrintBoard()
			   	board.MakeRandomMove()
			   	board.PrintBoard()
		   	}
	   	}
   }
   
   black, white := board.GetResult()
   winner := "Draw"
   if black > white {
   		winner = "Black Wins!"
   } else if white > black {
   		winner = "White Wins!"
   }
   
   fmt.Println(winner, " - Black: ", black, "White", white)
   
}

