// Loading PNGs with SDL_image
// Adapted from http://lazyfoo.net/tutorials/SDL/06_extension_libraries_and_loading_other_image_formats/index2.php

package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	err              error
	window           *sdl.Window
	renderer         *sdl.Renderer
	texture          *sdl.Texture
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

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return err
	}

	renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)

	imgFlags := img.INIT_PNG
	imgInitResult := img.Init(imgFlags)
	if (imgInitResult & imgFlags) != imgFlags {
		return img.GetError()
	}

	screenSurface, err = window.GetSurface()
	if err != nil {
		return err
	}

	return nil
}

func loadTexture(path string) (*sdl.Texture, error) {
	var newTexture *sdl.Texture

	loadedSurface, err := img.Load(path)
	if err != nil {
		return nil, err
	}

	newTexture, err = renderer.CreateTextureFromSurface(loadedSurface)
	if err != nil {
		return nil, err
	}

	loadedSurface.Free()

	return newTexture, nil
}

func loadSurface(path string) (*sdl.Surface, error) {

	loadedSurface, err := img.Load(path)
	if err != nil {
		return nil, err
	}

	optimizedSurface, err := loadedSurface.Convert(screenSurface.Format, 0)
	if err != nil {
		return nil, err
	}

	// Get rid of old loaded surface
	loadedSurface.Free()

	return optimizedSurface, nil
}

func loadMedia() error {
	texture, err = loadTexture("texture.png")
	if err != nil {
		return err
	}

	return nil
}

func close() {
	texture.Destroy()
	renderer.Destroy()
	window.Destroy()
	img.Quit()
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

	quit = false
	for !quit {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				quit = true
			}
		}

		renderer.Clear()
		renderer.Copy(texture, nil, nil)
		renderer.Present()
	}

	close()
}
