package graphics

import (
	"chess/board"
	"chess/pieces"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	BOARD_WIDTH = 800
	BOARD_HEIGHT = 600
)
var BACKGROUND = sdl.Color{R: 0, G: 0, B: 255, A: 255}

func DrawBoard(b *board.Board, surface *sdl.Surface) {
	var rectangle sdl.Rect
	rectangle.W = BOARD_WIDTH/board.BOARD_SIZE
	rectangle.H = BOARD_HEIGHT/board.BOARD_SIZE

	for x, row := range b.Nodes {
		for y, piece := range row {
			rectangle.X = int32(float32(rectangle.W) * float32(x))
			rectangle.Y = int32(float32(rectangle.H) * float32(y))

			var color sdl.Color
			if piece.PieceType != pieces.NONE {
				if (piece.Color == pieces.WHITE) {
					color = sdl.Color{R: 255,G: 255,B: 255,A: 255}
				} else {
					color = sdl.Color{R: 0, G: 0, B: 0, A: 255}
				}
			} else {
				color = BACKGROUND
			}

			err := surface.FillRect(&rectangle,color.Uint32())
			if (err != nil) {
				panic(err)
			}
			
		}
	}
}