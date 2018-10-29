// Getting an Image on the Screen
// Adapted from http://lazyfoo.net/tutorials/SDL/02_getting_an_image_on_the_screen/index.php
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
	helloWorld, err = sdl.LoadBMP("hello_world.bmp")
	if err != nil {
		return err
	}

	return nil
}

func Close() {
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
	sdl.Delay(2000)

	Close()
}
