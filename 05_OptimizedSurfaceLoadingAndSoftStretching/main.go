// Optimized Surface Loading and Soft Stretching
// Adapted from http://lazyfoo.net/tutorials/SDL/05_optimized_surface_loading_and_soft_stretching/index.php
package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	err              error
	window           *sdl.Window
	screenSurface    *sdl.Surface
	stretchedSurface *sdl.Surface
	quit             bool
	event            sdl.Event
)

func initSDL() error {
	err = sdl.Init(sdl.INIT_VIDEO)

	if err != nil {
		return err
	}

	window, err = sdl.CreateWindow("SDL Tutorial", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		return err
	}

	screenSurface, err = window.GetSurface()
	if err != nil {
		return err
	}

	return nil
}

func loadSurface(path string) (*sdl.Surface, error) {

	loadedSurface, err := sdl.LoadBMP(path)
	if err != nil {
		return nil, err
	}

	optimizedSurface, err := loadedSurface.Convert(screenSurface.Format, 0)
	if err != nil {
		return nil, err
	}

	return optimizedSurface, nil
}

func loadMedia() error {
	stretchedSurface, err = loadSurface("stretch.bmp")
	if err != nil {
		return err
	}

	return nil
}

func close() {

	window.Destroy()
	sdl.Quit()
}

func main() {
	err = initSDL()
	if err != nil {
		log.Fatal("Error initializing SDL:", err)
	}

	err = loadMedia()
	if err != nil {
		log.Fatal("Error loading Media:", err)
	}

	var stretchRect sdl.Rect
	stretchRect.X = 0
	stretchRect.Y = 0
	stretchRect.W = screenWidth
	stretchRect.H = screenHeight

	stretchedSurface.BlitScaled(nil, screenSurface, &stretchRect)
	window.UpdateSurface()

	quit = false
	for !quit {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				quit = true
			}
		}
	}

	close()
}
