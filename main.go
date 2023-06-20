package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"chess/board"
	"chess/graphics"
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
		b, err = b.MovePiece(src, dest)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func loop() {
	b := board.NewBoard()
	graphics.Init()
	
	defer sdl.Quit()
	window := graphics.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600)
	surface := graphics.GetSurface(window)
	surface.FillRect(nil, 0)

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
					}
				}
			}

			} // end switch
		}

		rect := sdl.Rect{X: 50, Y: 50, W: 100, H: 100}

		graphics.DrawBoard(&b, surface)
		
		surface.FillRect(&rect,255)

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

