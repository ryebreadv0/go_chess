package graphics

import (
	"chess/board"
	"chess/pieces"
	"github.com/veandco/go-sdl2/sdl"
)

var BACKGROUND = sdl.Color{R: 0, G: 0, B: 255, A: 255}

func DrawBoard(b *board.Board, window *sdl.Window, surface *sdl.Surface) {
	winWidth, winHeight := window.GetSize()
	
	var rectangle sdl.Rect
	rectangle.W = winWidth/board.BOARD_SIZE
	rectangle.H = winHeight/board.BOARD_SIZE

	for x, row := range b.Nodes {
		for y, piece := range row {
			rectangle.X = int32(float32(rectangle.W) * float32(y))
			rectangle.Y = int32(float32(rectangle.H) * float32(x))

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