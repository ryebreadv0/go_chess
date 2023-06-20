package graphics

import (
	"github.com/veandco/go-sdl2/sdl"
)

func Init() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
}

func CreateWindow(title string, x, y, w, h int32) *sdl.Window {
	window, err := sdl.CreateWindow(title, x, y, w, h, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	return window
}

func CreateRenderer(window *sdl.Window) *sdl.Renderer {
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	return renderer
}

func GetSurface(window *sdl.Window) *sdl.Surface {
	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	return surface
}

