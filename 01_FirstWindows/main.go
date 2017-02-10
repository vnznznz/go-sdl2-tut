// Hello SDL: Your First Graphics Window
// Adapted from http://lazyfoo.net/tutorials/SDL/01_hello_SDL/index2.php
package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

func main() {
	var window *sdl.Window
	var screenSurface *sdl.Surface
	var err error

	err = sdl.Init(sdl.INIT_VIDEO)

	if err != nil {
		log.Fatal("SDL could not initialize! Error:", err)
	}

	window, err = sdl.CreateWindow("SDL Tutorial", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_SHOWN)

	if err != nil {
		log.Fatal("Window could not be created! Error:", err)
	}

	screenSurface, err = window.GetSurface()

	if err != nil {
		log.Fatal("Unable to get Surface from Window! Error:", err)
	}

	screenSurface.FillRect(nil, sdl.MapRGB(screenSurface.Format, 0xFF, 0xFF, 0xFF))

	err = window.UpdateSurface()

	if err != nil {
		log.Fatal("Unable to update Surface! Error:", err)
	}

	sdl.Delay(2000)

	window.Destroy()
	sdl.Quit()
}
