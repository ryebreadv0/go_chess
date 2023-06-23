package board

import (
	"chess/pieces"
)

const BLACK = pieces.BLACK
const WHITE = pieces.WHITE

type piece = pieces.Piece

const BOARD_SIZE = 8

type Board struct {
	Nodes [BOARD_SIZE][BOARD_SIZE]piece
	Turn int
}

var emptyRow = [8]piece{pieces.NewNone(), pieces.NewNone(), pieces.NewNone(), pieces.NewNone(), pieces.NewNone(), pieces.NewNone(), pieces.NewNone(), pieces.NewNone()}

var defaultBoard = Board{
	Nodes: [8][8]piece{
		{pieces.NewRook(BLACK), pieces.NewKnight(BLACK), pieces.NewBishop(BLACK), pieces.NewKing(BLACK), pieces.NewQueen(BLACK), pieces.NewBishop(BLACK), pieces.NewKnight(BLACK), pieces.NewRook(BLACK)},
		{pieces.NewPawn(BLACK), pieces.NewPawn(BLACK), pieces.NewPawn(BLACK), pieces.NewPawn(BLACK), pieces.NewPawn(BLACK), pieces.NewPawn(BLACK), pieces.NewPawn(BLACK), pieces.NewPawn(BLACK)},
		emptyRow,
		emptyRow,
		emptyRow,
		emptyRow,
		{pieces.NewPawn(WHITE), pieces.NewPawn(WHITE), pieces.NewPawn(WHITE), pieces.NewPawn(WHITE), pieces.NewPawn(WHITE), pieces.NewPawn(WHITE), pieces.NewPawn(WHITE), pieces.NewPawn(WHITE)},
		{pieces.NewRook(WHITE), pieces.NewKnight(WHITE), pieces.NewBishop(WHITE), pieces.NewQueen(WHITE), pieces.NewKing(WHITE), pieces.NewBishop(WHITE), pieces.NewKnight(WHITE), pieces.NewRook(WHITE)},
	},
	Turn: BLACK,
}

func NewBoard() Board {
	return defaultBoard
}

func (b *Board) String() string {
	var str string
	for _, row := range b.Nodes {
		for _, piece := range row {
			str += piece.String()
		}
		str += "\n"
	}
	return str
}

