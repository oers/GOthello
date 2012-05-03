package gothello

import (
	"testing"
)

//func Test_Replay(t *testing.T) {
//   //black Win
//   Replay("F5F6E6F4E3D6C5F3G4E2G5G6C7C3D3C2D2C6F7B5F1H4H3H5E7D7B3E1B4F8C1G1A5D8B6A6F2H2G3C8E8A4C4G2H1D1A3A2B2A1B1B7H6H7H8G7G8B8A7A8")
//}

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

func Test_Bitmasks(t *testing.T) {
	b := new(Board)

	for i := 1; i < 8; i++ {
		for j := 1; j < 8; j++ {
			assertTrue(t, b.IsEmpty(i, j), "Empty", i, j)
		}
	}

	//black
	assertTrue(t, b.IsEmpty(5, 5), "Empty", 5, 5)
	b.setStone(5, 5, 1)
	assertFalse(t, b.IsEmpty(5, 5), "Not Empty", 5, 5)
	assertTrue(t, b.IsStone(5, 5, 1), "Black", 5, 5)

	assertTrue(t, b.IsEmpty(0, 0), "Empty", 0, 0)
	assertTrue(t, b.IsEmpty(1, 1), "Empty", 1, 1)
	assertTrue(t, b.IsEmpty(2, 2), "Empty", 2, 2)
	assertTrue(t, b.IsEmpty(3, 3), "Empty", 3, 3)
	assertTrue(t, b.IsEmpty(4, 4), "Empty", 4, 4)
	assertTrue(t, b.IsEmpty(6, 6), "Empty", 6, 6)
	assertTrue(t, b.IsEmpty(7, 7), "Empty", 7, 7)

	//white
	assertTrue(t, b.IsEmpty(6, 6), "Empty", 6, 6)
	b.setStone(6, 6, 0)
	assertFalse(t, b.IsEmpty(6, 6), "Not Empty", 6, 6)
	assertTrue(t, b.IsStone(6, 6, 0), "White", 6, 6)

	b.setStone(7, 7, 0) //h8 is special
	assertFalse(t, b.IsEmpty(7, 7), "Not Empty", 7, 7)
	assertTrue(t, b.IsStone(7, 7, 0), "White", 7, 7)
}

func Test_Gameplay(t *testing.T) {     
        b := makeBoard()
        //wb
        //bw     
        assertTrue(t, b.IsStone(3, 4, 0), "Black", 4, 3)
        assertTrue(t, b.IsStone(4, 3, 0), "Black", 4, 3)
        assertTrue(t, b.IsStone(3, 3, 1), "White", 3, 3)
        assertTrue(t, b.IsStone(4, 4, 1), "White", 4, 4)
     
        assertTrue(t, b.markNextMoves(), b.ToString()); //mark available moves
        assertTrue(t, b.markNextMoves(), b.ToString()); //mark available moves, must work twice
        
        b.PrintBoard()
        
        
        assertTrue(t, b.IsStone(3, 4, 0), "Black", 4, 3)
        assertTrue(t, b.IsStone(4, 3, 0), "Black", 4, 3)
        assertTrue(t, b.IsStone(3, 3, 1), "White", 3, 3)
        assertTrue(t, b.IsStone(4, 4, 1), "White", 4, 4)
        
        //new possible moves all Black of Course
        assertTrue(t, b.IsPossibleMove(2, 3), "Selectable", 2, 3)
        assertTrue(t, b.IsPossibleMove(4, 5), "Selectable", 4, 5)
        assertTrue(t, b.IsPossibleMove(3, 2), "Selectable", 3, 2)
        assertTrue(t, b.IsPossibleMove(5, 4), "Selectable", 5, 4)
//        assertEquals("Selectable", b.getState(2, 3), STATE.SELECTABLE);
//        assertEquals("Selectable", b.getState(4, 5), STATE.SELECTABLE);
//        assertEquals("Selectable", b.getState(3, 2), STATE.SELECTABLE);
//        assertEquals("Selectable", b.getState(5, 4), STATE.SELECTABLE);
        
        //assertEquals(true, b.isNextPlayerBlack());
//        
//        //make Illegal Move and expect an Exception
//
//        assertFalse(b.makeMove(1, 1));
//   
//        //nothing has changed
//        assertEquals("Black", b.getState(3, 4), STATE.BLACK);
//        assertEquals("Black", b.getState(4, 3), STATE.BLACK);
//        assertEquals("White", b.getState(3, 3), STATE.WHITE);
//        assertEquals("White", b.getState(4, 4), STATE.WHITE);
//        
//        //new possible moves all Black of Course
//        assertEquals("Selectable", b.getState(2, 3), STATE.SELECTABLE);
//        assertEquals("Selectable", b.getState(4, 5), STATE.SELECTABLE);
//        assertEquals("Selectable", b.getState(3, 2), STATE.SELECTABLE);
//        assertEquals("Selectable", b.getState(5, 4), STATE.SELECTABLE);
//        
//        //make Legal Move
//        assertTrue(b.makeMove(2, 3));
//           //nothing has changed
//        assertEquals("Black", b.getState(2, 3), STATE.BLACK);
//        assertEquals("Black", b.getState(3, 4), STATE.BLACK);
//        assertEquals("Black", b.getState(4, 3), STATE.BLACK);
//        assertEquals("White", b.getState(3, 3), STATE.BLACK);
//        assertEquals("White", b.getState(4, 4), STATE.WHITE);
//        
//         b.markNextMoves(); //mark available moves
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
