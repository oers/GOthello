package gothello

import (
	"testing"
	"math/rand"
	"time"
)

func BenchmarkReplay(b *testing.B) { //benchmark function starts with "Benchmark" and takes a pointer to type testing.B
    for i := 0; i < b.N; i++ { //use b.N for looping 
         Replay("F5F6E6F4E3D6C5F3G4E2G5G6C7C3D3C2D2C6F7B5F1H4H3H5E7D7B3E1B4F8C1G1A5D8B6A6F2H2G3C8E8A4C4G2H1D1A3A2B2A1B1B7H6H7H8G7G8B8A7A8")
    }
}

func BenchmarkRandomPlay(b *testing.B) { //benchmark function starts with "Benchmark" and takes a pointer to type testing.B
    for i := 0; i < b.N; i++ { //use b.N for looping 
    	 rand.Seed(time.Now().UnixNano()) //set the seed
         playRandom()
    }
}

func BenchmarkSolve10(b *testing.B) { //benchmark function starts with "Benchmark" and takes a pointer to type testing.B
    for i := 0; i < b.N; i++ { //use b.N for looping 
    	 rand.Seed(time.Now().UnixNano()) //set the seed
         b := playNMoves(10)
         b.SolveBoard()
    }
}

func TestReplayBlackWin(t *testing.T) {
   //black Win
   board := Replay("F5F6E6F4E3D6C5F3G4E2G5G6C7C3D3C2D2C6F7B5F1H4H3H5E7D7B3E1B4F8C1G1A5D8B6A6F2H2G3C8E8A4C4G2H1D1A3A2B2A1B1B7H6H7H8G7G8B8A7A8")
   black, white := board.GetResult()
   assertTrue(t, black == 40, "40 Black Discs", black)
   assertTrue(t, white == 24, "24 White Discs", white)
   
   //Assert.assertEquals("A8".toLowerCase(), b.getLastMove());

}

func TestSolveBlackWin(t *testing.T) {
   //black Win
   Solve("F5F6E6F4E3D6C5F3G4E2G5G6C7C3D3C2D2C6F7B5F1H4H3H5E7D7B3E1B4F8C1G1A5D8B6A6F2H2G3C8E8A4C4G2H1D1A3A2B2A1") //B1B7H6H7H8G7G8B8A7A8")
//   black, white := board.GetResult()
//   assertTrue(t, black == 40, "40 Black Discs", black)
//   assertTrue(t, white == 24, "24 White Discs", white)
   
   //Assert.assertEquals("A8".toLowerCase(), b.getLastMove());

}

func TestSolveBlackWin2(t *testing.T) {
   //black Win
   Solve("F5F6E6F4E3D6C5F3G4E2G5G6C7C3D3C2D2C6F7B5F1H4H3H5E7D7B3E1B4F8C1G1A5D8B6A6F2H2G3C8E8A4C4G2H1D1A3A2B2A1B1B7H6H7H8G7G8B8A7") //A8")
//   black, white := board.GetResult()
//   assertTrue(t, black == 40, "40 Black Discs", black)
//   assertTrue(t, white == 24, "24 White Discs", white)
   
   //Assert.assertEquals("A8".toLowerCase(), b.getLastMove());

}

func TestSolveBlackWin3(t *testing.T) {
   //black Win
   Solve("F5F6E6F4E3D6C5F3G4E2G5G6C7C3D3C2D2C6F7B5F1H4H3H5E7D7B3E1B4F8C1G1A5D8B6A6F2H2G3C8E8A4C4G2H1D1A3A2B2A1B1B7H6H7H8G7G8B8") //A7A8")
//   black, white := board.GetResult()
//   assertTrue(t, black == 40, "40 Black Discs", black)
//   assertTrue(t, white == 24, "24 White Discs", white)
   
   //Assert.assertEquals("A8".toLowerCase(), b.getLastMove());

}

func TestReplayWhiteWin(t *testing.T) {
   //black Win
   board := Replay("F5F6E6F4G5G6G4E7F3D6F7H3D8D3H4H5D7E3E2D2G3F8C5E8G8B5C4B4C3D1F2C6F1H2B3C2B1C7A5A3A4A6C8B6H7C1E1G2B7A7H6H8B2B8G1H1G7A1A2A8")
   black, white := board.GetResult()
   
   assertTrue(t, black == 20, "20 Black Discs", black)
   assertTrue(t, white == 44, "44 White Discs", white)
   
   //Assert.assertEquals("A8".toLowerCase(), b.getLastMove());
}

func TestReplayWipeOut(t *testing.T) {
   //black Win
   board := Replay("F5F6E6F4G5G6G4C6F3F7E7D6D7F8E8F2G3C7G8H5H6H3D8G7H4H7H8E3H2C5B6C8C4B3C3B8A3A6A7A8B7B5B4D3C2D1G2F1G1H1D2E2E1B1A4A5A2C1")
   black, white := board.GetResult()
   
   assertTrue(t, black == 64, "64 Black Discs", black)
   assertTrue(t, white == 0, "0 White Discs", white)

   
	//Assert.assertEquals("C1".toLowerCase(), b.getLastMove());
}

func TestReplayDraw(t *testing.T) {
   //black Win
   board := Replay("F5D6C3D3C4F4F6F3E6E7D7G6F8F7H6C5C6D8E3B6G4B4B5H3H4E2D2G3F1F2G5E8C8C7A3E1D1C1B1C2A6G2B8A5G8H5H1A7B2G1A4B7H2A2A8A1B3H7G7H8")
   black, white := board.GetResult()
   
   assertTrue(t, black == 32, "32 Black Discs", black)
   assertTrue(t, white == 32, "32 White Discs", white)
   //Assert.assertEquals("H8".toLowerCase(), b.getLastMove());

}

func playNMoves(n int) (b *Board){
	b = MakeBoard()
	for i := 0; i <  60 - n; i++{
		b.MakeRandomMove()
	}
	return
	//b.PrintBoard()
}


func playRandom() {
	b := MakeBoard()
	for !b.IsFinished(){
		b.MakeRandomMove()
	}
	//b.PrintBoard()
}

func TestBitmasks(t *testing.T) {
	b := new(Board)

	for i := 1; i < 8; i++ {
		for j := 1; j < 8; j++ {
			assertTrue(t, b.isEmptyInt(i, j), "Empty", i, j)
		}
	}

	//black
	assertTrue(t, b.isEmptyInt(5, 5), "Empty", 5, 5)
	b.setStoneInt(5, 5, 1)
	assertFalse(t, b.isEmptyInt(5, 5), "Not Empty", 5, 5)
	assertTrue(t, b.isStoneInt(5, 5, 1), "Black", 5, 5)

	assertTrue(t, b.isEmptyInt(0, 0), "Empty", 0, 0)
	assertTrue(t, b.isEmptyInt(1, 1), "Empty", 1, 1)
	assertTrue(t, b.isEmptyInt(2, 2), "Empty", 2, 2)
	assertTrue(t, b.isEmptyInt(3, 3), "Empty", 3, 3)
	assertTrue(t, b.isEmptyInt(4, 4), "Empty", 4, 4)
	assertTrue(t, b.isEmptyInt(6, 6), "Empty", 6, 6)
	assertTrue(t, b.isEmptyInt(7, 7), "Empty", 7, 7)

	//white
	assertTrue(t, b.isEmptyInt(6, 6), "Empty", 6, 6)
	b.setStoneInt(6, 6, 0)
	assertFalse(t, b.isEmptyInt(6, 6), "Not Empty", 6, 6)
	assertTrue(t, b.isStoneInt(6, 6, 0), "White", 6, 6)

	b.setStoneInt(7, 7, 0) //h8 is special
	assertFalse(t, b.isEmptyInt(7, 7), "Not Empty", 7, 7)
	assertTrue(t, b.isStoneInt(7, 7, 0), "White", 7, 7)
}

func TestGameplay(t *testing.T) {     
        b := MakeBoard()
        //wb
        //bw     
        assertTrue(t, b.isStoneInt(3, 4, 0), "Black", 4, 3)
        assertTrue(t, b.isStoneInt(4, 3, 0), "Black", 4, 3)
        assertTrue(t, b.isStoneInt(3, 3, 1), "White", 3, 3)
        assertTrue(t, b.isStoneInt(4, 4, 1), "White", 4, 4)
     
        assertTrue(t, b.markNextMoves(), b.ToString()); //mark available moves
        assertTrue(t, b.markNextMoves(), b.ToString()); //mark available moves, must work twice
        
        assertTrue(t, b.isStoneInt(3, 4, 0), "Black", 4, 3)
        assertTrue(t, b.isStoneInt(4, 3, 0), "Black", 4, 3)
        assertTrue(t, b.isStoneInt(3, 3, 1), "White", 3, 3)
        assertTrue(t, b.isStoneInt(4, 4, 1), "White", 4, 4)
        
        //new possible moves all Black of Course
        assertTrue(t, b.isPossibleMoveInt(2, 3), "Selectable", 2, 3)
        assertTrue(t, b.isPossibleMoveInt(4, 5), "Selectable", 4, 5)
        assertTrue(t, b.isPossibleMoveInt(3, 2), "Selectable", 3, 2)
        assertTrue(t, b.isPossibleMoveInt(5, 4), "Selectable", 5, 4)
        
        assertTrue(t, b.IsNextPlayerBlack(), "Blacks Turn", b.ToString())
      
        //make Illegal Move and expect an Exception
        assertFalse(t, b.makeMoveInt(1, 1), "Illegal Move expected", 1, 1,b.ToString());
 
        //nothing has changed
        assertTrue(t, b.isStoneInt(3, 4, 0), "Black", 4, 3)
        assertTrue(t, b.isStoneInt(4, 3, 0), "Black", 4, 3)
        assertTrue(t, b.isStoneInt(3, 3, 1), "White", 3, 3)
        assertTrue(t, b.isStoneInt(4, 4, 1), "White", 4, 4)
        
        //new possible moves all Black of Course
        assertTrue(t, b.isPossibleMoveInt(2, 3), "Selectable", 2, 3)
        assertTrue(t, b.isPossibleMoveInt(4, 5), "Selectable", 4, 5)
        assertTrue(t, b.isPossibleMoveInt(3, 2), "Selectable", 3, 2)
        assertTrue(t, b.isPossibleMoveInt(5, 4), "Selectable", 5, 4)
        
        assertTrue(t, b.IsNextPlayerBlack(), "Blacks Turn", b.ToString())        
         
        //make Legal Move
		assertTrue(t, b.makeMoveInt(2, 3), "Legal Move expected", b.ToString());

        assertTrue(t, b.isStoneInt(2, 3, 0), "Black", 2, 3, b.ToString())
        assertTrue(t, b.isStoneInt(3, 4, 0), "Black", 3, 4, b.ToString())
        assertTrue(t, b.isStoneInt(4, 3, 0), "Black", 4, 3, b.ToString())
        assertTrue(t, b.isStoneInt(3, 3, 0), "Black", 3, 3, b.ToString())
        assertTrue(t, b.isStoneInt(4, 4, 1), "White", 4, 4, b.ToString())
        
        t.Log("Mark Next moves")
              
		assertTrue(t, b.markNextMoves(), b.ToString()); //mark available moves
	    assertFalse(t, b.IsNextPlayerBlack(), "Whites Turn", b.ToString())   
	    
	    t.Log("Make Move", 2, 3)
	    //try this move again, should fail
		assertFalse(t, b.makeMoveInt(2, 3), "Illegal Move expected", 1, 1, b.ToString());
//        
//        //all marked fields have to be playable
//        for(BoardMove move : b.getPossibleMoves())
//        {
//           assertTrue(move+ " is not legal but was marked as legal",b.isLegalMove(move.getRow(), move.getColumn()));
//
//        }
//        //try this move again, should fail
//       assertFalse(b.makeMove(2, 3));  
}

func assertTrue(t *testing.T, condition bool, messages ...interface{}) {
	if condition != true {
		t.Log(messages)
		t.FailNow()
	}
}

func assertFalse(t *testing.T, condition bool, messages ...interface{}) {
	if condition != false {
		t.Log(messages)
		t.FailNow()
	}
}
