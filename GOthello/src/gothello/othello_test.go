package gothello

import (
	"testing"
)

func BenchmarkReplay(b *testing.B) { //benchmark function starts with "Benchmark" and takes a pointer to type testing.B
    for i := 0; i < b.N; i++ { //use b.N for looping 
         Replay("F5F6E6F4E3D6C5F3G4E2G5G6C7C3D3C2D2C6F7B5F1H4H3H5E7D7B3E1B4F8C1G1A5D8B6A6F2H2G3C8E8A4C4G2H1D1A3A2B2A1B1B7H6H7H8G7G8B8A7A8")
    }
}

func BenchmarkRandomPlay(b *testing.B) { //benchmark function starts with "Benchmark" and takes a pointer to type testing.B
    for i := 0; i < b.N; i++ { //use b.N for looping 
         playRandom()
    }
}

func TestReplay(t *testing.T) {
   //black Win
   Replay("F5F6E6F4E3D6C5F3G4E2G5G6C7C3D3C2D2C6F7B5F1H4H3H5E7D7B3E1B4F8C1G1A5D8B6A6F2H2G3C8E8A4C4G2H1D1A3A2B2A1B1B7H6H7H8G7G8B8A7A8")
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
