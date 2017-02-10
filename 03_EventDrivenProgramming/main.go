// Event Driven Programming
// Adapted from http://lazyfoo.net/tutorials/SDL/03_event_driven_programming/index.php
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
	err           error
	window        *sdl.Window
	screenSurface *sdl.Surface
	helloWorld    *sdl.Surface
	quit          bool
	event         sdl.Event
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

func loadMedia() error {
	helloWorld, err = sdl.LoadBMP("x_to_close.bmp")
	if err != nil {
		return err
	}

	return nil
}

func close() {
	helloWorld.Free()
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
		log.Fatal("Error initializing SDL:", err)
	}

	helloWorld.Blit(nil, screenSurface, nil)
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
