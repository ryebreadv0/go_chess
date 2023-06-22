package graphics

import (
	"chess/board"
	"chess/pieces"
	"chess/utils"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var BACKGROUND = sdl.Color{R: 0, G: 150, B: 150, A: 255}
var HIGHLIGHT = sdl.Color{R: 255, G: 255, B: 255, A: 100}
var VALID_MOVE = sdl.Color{R: 255, G: 255, B: 255, A: 120}
var BLACK_PAWN_BITMAP *sdl.Texture = nil
var BLACK_ROOK_BITMAP *sdl.Texture = nil
var BLACK_KNIGHT_BITMAP *sdl.Texture = nil
var BLACK_BISHOP_BITMAP *sdl.Texture = nil
var BLACK_QUEEN_BITMAP *sdl.Texture = nil
var BLACK_KING_BITMAP *sdl.Texture = nil
var WHITE_PAWN_BITMAP *sdl.Texture = nil
var WHITE_ROOK_BITMAP *sdl.Texture = nil
var WHITE_KNIGHT_BITMAP *sdl.Texture = nil
var WHITE_BISHOP_BITMAP *sdl.Texture = nil
var WHITE_QUEEN_BITMAP *sdl.Texture = nil
var WHITE_KING_BITMAP *sdl.Texture = nil

func bitmapToTexture(renderer *sdl.Renderer, asset string) *sdl.Texture {
	surface, err := sdl.LoadBMP(asset)
	if err != nil {
		err := fmt.Sprintf("Error loading bitmap: %s", asset)
		panic(err)
	}

	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		panic(err)
	}
	surface.Free()

	return texture
}

func InitBitmaps(renderer *sdl.Renderer) {
	BLACK_PAWN_BITMAP = bitmapToTexture(renderer, "assets/black_pawn.bmp")
	BLACK_ROOK_BITMAP = bitmapToTexture(renderer, "assets/black_rook.bmp")
	BLACK_KNIGHT_BITMAP = bitmapToTexture(renderer, "assets/black_knight.bmp")
	BLACK_BISHOP_BITMAP = bitmapToTexture(renderer, "assets/black_bishop.bmp")
	BLACK_QUEEN_BITMAP = bitmapToTexture(renderer, "assets/black_queen.bmp")
	BLACK_KING_BITMAP = bitmapToTexture(renderer, "assets/black_king.bmp")
	WHITE_PAWN_BITMAP = bitmapToTexture(renderer, "assets/white_pawn.bmp")
	WHITE_ROOK_BITMAP = bitmapToTexture(renderer, "assets/white_rook.bmp")
	WHITE_KNIGHT_BITMAP = bitmapToTexture(renderer, "assets/white_knight.bmp")
	WHITE_BISHOP_BITMAP = bitmapToTexture(renderer, "assets/white_bishop.bmp")
	WHITE_QUEEN_BITMAP = bitmapToTexture(renderer, "assets/white_queen.bmp")
	WHITE_KING_BITMAP = bitmapToTexture(renderer, "assets/white_king.bmp")
}

func DestroyBitmaps() {
	BLACK_PAWN_BITMAP.Destroy()
	BLACK_ROOK_BITMAP.Destroy()
	BLACK_KNIGHT_BITMAP.Destroy()
	BLACK_BISHOP_BITMAP.Destroy()
	BLACK_QUEEN_BITMAP.Destroy()
	BLACK_KING_BITMAP.Destroy()
	WHITE_PAWN_BITMAP.Destroy()
	WHITE_ROOK_BITMAP.Destroy()
	WHITE_KNIGHT_BITMAP.Destroy()
	WHITE_BISHOP_BITMAP.Destroy()
	WHITE_QUEEN_BITMAP.Destroy()
	WHITE_KING_BITMAP.Destroy()
}

func DrawBoard(b *board.Board, window *sdl.Window, renderer *sdl.Renderer) {
	winWidth, winHeight := window.GetSize()
	
	var err error
	
	var rectangle sdl.FRect
	rectangle.W = float32(winWidth)/float32(board.BOARD_SIZE)
	rectangle.H = float32(winHeight)/float32(board.BOARD_SIZE)

	for y, row := range b.Nodes {
		for x, piece := range row {
			rectangle.X = float32(rectangle.W) * float32(x)
			rectangle.Y = float32(rectangle.H) * float32(y)

			// determine if even or odd
			even := (y + x) % 2
			if even == 0 {
				renderer.SetDrawColor(0, 150, 190, 255)
			} else {
				renderer.SetDrawColor(0, 190, 150, 255)
			}
			err = renderer.FillRectF(&rectangle)
			if err != nil {
				panic(err)
			}
			

			switch piece.PieceType {
			case pieces.PAWN:
			{
				if piece.Color == pieces.BLACK {
					err = renderer.CopyF(BLACK_PAWN_BITMAP, nil, &rectangle)
					if err != nil {
						panic(err)
					}
				} else {
					err = renderer.CopyF(WHITE_PAWN_BITMAP, nil, &rectangle)
					if err != nil {
						panic(err)
					}
				}
			}
			case pieces.ROOK:
			{
				if piece.Color == pieces.BLACK {
					err = renderer.CopyF(BLACK_ROOK_BITMAP, nil, &rectangle)
					if err != nil {
						panic(err)
					}
				} else {
					err = renderer.CopyF(WHITE_ROOK_BITMAP, nil, &rectangle)
					if err != nil {
						panic(err)
					}
				}
			}
			case pieces.KNIGHT:
			{
				if piece.Color == pieces.BLACK {
					err = renderer.CopyF(BLACK_KNIGHT_BITMAP, nil, &rectangle)
					if err != nil {
						panic(err)
					}
				} else {
					err = renderer.CopyF(WHITE_KNIGHT_BITMAP, nil, &rectangle)
					if err != nil {
						panic(err)
					}
				}
			}
			case pieces.BISHOP:
			{
				if piece.Color == pieces.BLACK {
					err = renderer.CopyF(BLACK_BISHOP_BITMAP, nil, &rectangle)
					if err != nil {
						panic(err)
					}
				} else {
					err = renderer.CopyF(WHITE_BISHOP_BITMAP, nil, &rectangle)
					if err != nil {
						panic(err)
					}
				}
			}
			case pieces.QUEEN:
			{
				if piece.Color == pieces.BLACK {
					err = renderer.CopyF(BLACK_QUEEN_BITMAP, nil, &rectangle)
					if err != nil {
						panic(err)
					}
				} else {
					err = renderer.CopyF(WHITE_QUEEN_BITMAP, nil, &rectangle)
					if err != nil {
						panic(err)
					}
				}
			}
			case pieces.KING:
			{
				if piece.Color == pieces.BLACK {
					err = renderer.CopyF(BLACK_KING_BITMAP, nil, &rectangle)
					if err != nil {
						panic(err)
					}
				} else {
					err = renderer.CopyF(WHITE_KING_BITMAP, nil, &rectangle)
					if err != nil {
						panic(err)
					}
				}
			}
			} // end switch


			// renderer.SetDrawColor(0, 0, 0, 255)
			// err = renderer.DrawRect(&rectangle)
			// if err != nil {
			// 	panic(err)
			// }
			
		}
	}
}

func DrawHightlight(window *sdl.Window, renderer *sdl.Renderer, location utils.Vec2) {
	winWidth, winHeight := window.GetSize()
	
	var rectangle sdl.FRect
	rectangle.W = float32(winWidth)/float32(board.BOARD_SIZE)
	rectangle.H = float32(winHeight)/float32(board.BOARD_SIZE)

	rectangle.X = float32(rectangle.W) * float32(location.X)
	rectangle.Y = float32(rectangle.H) * float32(location.Y)

	renderer.SetDrawColor(HIGHLIGHT.R, HIGHLIGHT.G, HIGHLIGHT.B, HIGHLIGHT.A)
	err := renderer.FillRectF(&rectangle)
	if (err != nil) {
		panic(err)
	}
}

func DrawValidMoves(window *sdl.Window, renderer *sdl.Renderer, moves *[]utils.Vec2) {
	winWidth, winHeight := window.GetSize()
	
	var rectangle sdl.FRect
	rectangle.W = float32(winWidth)/float32(board.BOARD_SIZE)
	rectangle.H = float32(winHeight)/float32(board.BOARD_SIZE)

	// moves := b.ListValidMoves(location)
	// fmt.Println(moves)

	for _, move := range *moves {
		rectangle.X = float32(rectangle.W) * float32(move.X)
		rectangle.Y = float32(rectangle.H) * float32(move.Y)

		renderer.SetDrawColor(VALID_MOVE.R, VALID_MOVE.G, VALID_MOVE.B, VALID_MOVE.A)
		err := renderer.FillRectF(&rectangle)
		if (err != nil) {
			panic(err)
		}
	}

}