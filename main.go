package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

var winTitle string = "Experiments"
var winWidth, winHeight int = 800, 600
var imageName string = "test.png"

func main() {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var texture *sdl.Texture
	var src, dst sdl.Rect

	window = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	if window == nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", sdl.GetError())
		os.Exit(1)
	}

	renderer = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if renderer == nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", sdl.GetError())
		os.Exit(2)
	}

	image := img.Load(imageName)
	if image == nil {
		fmt.Fprintf(os.Stderr, "Failed to load PNG: %s\n", sdl.GetError())
		os.Exit(3)
	}

	texture = renderer.CreateTextureFromSurface(image)
	if texture == nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", sdl.GetError())
		os.Exit(4)
	}

	src = sdl.Rect{0, 0, 512, 512}
	dst = sdl.Rect{100, 50, 512, 512}
	var angle float64 = 0

	var quit bool = false
	var event sdl.Event

	for !quit {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				quit = true
			case *sdl.KeyUpEvent:
				switch t.Keysym.Sym {
				case sdl.K_q:
					quit = true
				}
			case *sdl.KeyDownEvent:
				switch t.Keysym.Sym {
				case sdl.K_LEFT:
					angle = angle - 1
				case sdl.K_RIGHT:
					angle = angle + 1
				}
			}
		}
		renderer.Clear()
		renderer.SetDrawColor(255, 0, 0, 255)
		renderer.FillRect(&sdl.Rect{0, 0, int32(winWidth), int32(winHeight)})
		renderer.CopyEx(texture, &src, &dst, angle, nil, 0)
		renderer.Present()
	}

	image.Free()
	texture.Destroy()
	renderer.Destroy()
	window.Destroy()
}
