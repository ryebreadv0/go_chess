package main

import (
	"chess/board"
	"chess/graphics"
	"chess/utils"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func board_commandline(b board.Board) {
	var err error
	for {
		fmt.Println(b.String())
		// user input for move, src, dest
		var src, dest board.Vec2

		fmt.Println("Enter src: ")
		fmt.Scanf("%d %d", &src.X, &src.Y)
		fmt.Println("Enter dest: ")
		fmt.Scanf("%d %d", &dest.X, &dest.Y)

		fmt.Printf("Moving piece from %v to %v\n", src, dest)
		err = b.MovePiece(src, dest)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func loop() {
	b := board.NewBoard()
	var err error
	graphics.Init()
	
	defer sdl.Quit()
	
	window := graphics.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600)
	surface := graphics.GetSurface(window)

	var prevCoord utils.Vec2
	var selected bool = false


	window_open := true
	for window_open {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
			{

				window_open = false
				break
			}
			case *sdl.MouseButtonEvent:
			{
				if event.(*sdl.MouseButtonEvent).Type == sdl.MOUSEBUTTONDOWN {
					if event.(*sdl.MouseButtonEvent).Button == sdl.BUTTON_LEFT {
						fmt.Println("Left mouse button pressed")
						// get mouse position
						var mousePos board.Vec2
						mousePos.X = int(event.(*sdl.MouseButtonEvent).X)
						mousePos.Y = int(event.(*sdl.MouseButtonEvent).Y)
						fmt.Println("Mouse position: ", mousePos)

						width, height := window.GetSize()
						width = int32(float32(mousePos.X)/(float32(width)/float32(board.BOARD_SIZE)))
						height = int32(float32(mousePos.Y)/(float32(height)/float32(board.BOARD_SIZE)))

						boardPos := board.Vec2{X: int(width), Y: int(height)}
						fmt.Println("Board position: ", boardPos)

						if (selected) {
							err = b.MovePiece(prevCoord, boardPos)
							if (err != nil) {
								fmt.Println(err)
							}
							selected = false
						} else {
							prevCoord = boardPos
							selected = true
						}
					}
				}
			}

			} // end switch
		}


		graphics.DrawBoard(&b, window, surface)

		window.UpdateSurface()
	}
}

func main() {
	// b := board.NewBoard()
	// board_commandline(b)

	
	loop()
	



	// window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
	// 	800, 600, sdl.WINDOW_SHOWN)
	// if err != nil {
	// 	panic(err)
	// }
	// defer window.Destroy()

	// surface, err := window.GetSurface()
	// if err != nil {
	// 	panic(err)
	// }
	// surface.FillRect(nil, 0)

	// rect := sdl.Rect{X: 0, Y: 0, W: 200, H: 200}
	// colour := sdl.Color{R: 255, G: 0, B: 255, A: 255} // purple
	// pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
	// surface.FillRect(&rect, pixel)
	// window.UpdateSurface()

	// running := true
	// for running {
	// 	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
	// 		switch event.(type) {
	// 		case *sdl.QuitEvent:
	// 			running = false
	// 			break
	// 		}
	// 	}
	// }
}

