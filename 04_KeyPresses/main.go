// Key Presses
// Adapted from http://lazyfoo.net/tutorials/SDL/04_key_presses/index.php
package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

const (
	surfaceDefault = iota
	surfaceUp      = iota
	surfaceDown    = iota
	surfaceLeft    = iota
	surfaceRight   = iota
	surfaceTotal   = iota
)

var (
	err              error
	window           *sdl.Window
	screenSurface    *sdl.Surface
	currentSurface   *sdl.Surface
	quit             bool
	event            sdl.Event
	keyPressSurfaces map[int]*sdl.Surface
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

	return loadedSurface, nil
}

func loadMedia() error {
	keyPressSurfaces = make(map[int]*sdl.Surface)

	keyPressSurfaces[surfaceDefault], err = loadSurface("press.bmp")
	if err != nil {
		return err
	}

	keyPressSurfaces[surfaceUp], err = loadSurface("up.bmp")
	if err != nil {
		return err
	}

	keyPressSurfaces[surfaceDown], err = loadSurface("down.bmp")
	if err != nil {
		return err
	}

	keyPressSurfaces[surfaceLeft], err = loadSurface("left.bmp")
	if err != nil {
		return err
	}

	keyPressSurfaces[surfaceRight], err = loadSurface("right.bmp")
	if err != nil {
		return err
	}

	return nil
}

func Close() {

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

	currentSurface = keyPressSurfaces[surfaceDefault]

	window.UpdateSurface()

	quit = false
	for !quit {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				quit = true
			case *sdl.KeyboardEvent:
				if t.Type == sdl.KEYDOWN {
					switch t.Keysym.Sym {
					case sdl.K_UP:
						currentSurface = keyPressSurfaces[surfaceUp]
					case sdl.K_DOWN:
						currentSurface = keyPressSurfaces[surfaceDown]
					case sdl.K_LEFT:
						currentSurface = keyPressSurfaces[surfaceLeft]
					case sdl.K_RIGHT:
						currentSurface = keyPressSurfaces[surfaceRight]
					}
				}
			}
		}
		currentSurface.Blit(nil, screenSurface, nil)
		window.UpdateSurface()
	}

	Close()
}
