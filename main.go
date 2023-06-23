package main

import (
	"chess/board"
	"chess/graphics"
	"chess/utils"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func loop() {
	b := board.NewBoard()
	var err error
	graphics.Init()
	defer sdl.Quit()
	
	window := graphics.CreateWindow("chess", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 800)
	renderer := graphics.CreateRenderer(window)

	graphics.InitBitmaps(renderer)
	defer graphics.DestroyBitmaps()

	renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)

	var prevCoord utils.Vec2
	var selected bool = false
	var validMoves []utils.Vec2
	var shouldUpdate bool = true


	window_open := true
	for window_open {
		for curEvent := sdl.PollEvent(); curEvent != nil; curEvent = sdl.PollEvent() {
			switch event := curEvent.(type) {
			case *sdl.QuitEvent:
			{
				window_open = false
				break
			}
			case *sdl.MouseButtonEvent:
			{
				if event.Type == sdl.MOUSEBUTTONDOWN {
					if event.Button == sdl.BUTTON_LEFT {
						var mousePos board.Vec2
						mousePos.X = int(curEvent.(*sdl.MouseButtonEvent).X)
						mousePos.Y = int(curEvent.(*sdl.MouseButtonEvent).Y)

						width, height := window.GetSize()
						width = int32(float32(mousePos.X)/(float32(width)/float32(board.BOARD_SIZE)))
						height = int32(float32(mousePos.Y)/(float32(height)/float32(board.BOARD_SIZE)))

						boardPos := board.Vec2{X: int(width), Y: int(height)}

						if (selected) {
							err = b.MovePiece(prevCoord, boardPos)
							if (err != nil) {
								fmt.Println(err)
							}
							selected = false
						} else {
							if b.ValidSelection(boardPos) {
								prevCoord = boardPos
								selected = true
								validMoves = b.ListValidMoves(prevCoord)
							}
						}
					} else if event.Button == sdl.BUTTON_RIGHT {
						selected = false
					}
					shouldUpdate = true
				} 
			}
			case *sdl.KeyboardEvent:
			{
				if event.Type == sdl.KEYDOWN {
					switch event.Keysym.Sym {
					case sdl.K_ESCAPE:
						selected = false
						shouldUpdate = true
					}
				}
			}
			case *sdl.WindowEvent:
			{
				if event.Event == sdl.WINDOWEVENT_RESIZED {
					shouldUpdate = true
				}
			}

			} // end switch
		}
		
		if shouldUpdate {
			graphics.DrawBoard(&b, window, renderer)
	
			if selected {
				graphics.DrawHightlight(window, renderer, prevCoord)
				graphics.DrawValidMoves(window, renderer, &validMoves)
			}
	
			renderer.Present()
			shouldUpdate = false
		}
	}
}

func main() {
	loop()

}

