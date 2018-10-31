// clip and sprite sheet example inspired by http://lazyfoo.net/tutorials/SDL/11_clip_rendering_and_sprite_sheets/index.php
// Code license:		MIT
// Image license:		CC0-1.0
// 2018 Alexander Kahl <kahl@magneticcoffee.eu>
package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

const (
	screenWidth  = 640
	screenHeight = 480
	colorKey     = 0xff00ff // C-C-C-E-C ♫
)

// TextureMap holds the textures and wraps loading/destroying
type TextureMap map[string](*sdl.Texture)

var (
	window   *sdl.Window
	renderer *sdl.Renderer
	textures TextureMap = TextureMap{}
)

// Load BMP as `name`, removing the previous texture for `name`
func (tm TextureMap) LoadBMP(name, path string) (*sdl.Texture, error) {
	tmp, err := sdl.LoadBMP(path)

	if err != nil {
		return nil, err
	}

	// set color to transparent
	err = tmp.SetColorKey(true, colorKey)
	if err != nil {
		log.Fatalf("Can't set color key: %s", err)
	}

	// create texture
	tex, err := renderer.CreateTextureFromSurface(tmp)
	if err != nil {
		log.Fatalf("Can't convert surface to texture: %s", err)
	}

	// free the surface
	tmp.Free()

	// successfully loaded, replace previous texture of the same name

	// get rid of it
	tm.Destroy(name)

	// (re)assign
	tm[name] = tex

	return tex, err
}

// Remove (and free) surface from map
func (tm TextureMap) Destroy(name string) {

	t, present := tm[name]
	if present {
		t.Destroy()
	}
	delete(tm, name)

}

// Remove all surfaces (free surface + remove from map)
func (tm TextureMap) DestroyAll() {

	for n := range tm {
		delete(tm, n)
	}

}

// Render a texture to a renderer at position `x`, `y`.
// Optional `clip` rectangle for selecting a specific part of the texture
func (tm TextureMap) Render(name string, r *sdl.Renderer, x, y int32, clip *sdl.Rect) error {
	tex, present := tm[name]
	if !present {
		return fmt.Errorf("Render: Missing texture '%s'", name)
	}

	if r == nil {
		return fmt.Errorf("Render: Missing renderer target")
	}

	renderRect := &sdl.Rect{
		X: x,
		Y: y,
		W: screenWidth,
		H: screenHeight,
	}

	if clip != nil {
		renderRect.W = clip.W
		renderRect.H = clip.H
	}

	return renderer.Copy(tex, clip, renderRect)

}

func initSDL() error {

	err := sdl.Init(sdl.INIT_VIDEO)

	if err != nil {
		return err
	}

	window, err = sdl.CreateWindow(
		"Showing clips",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_SHOWN)

	if err != nil {
		return err
	}

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	if err != nil {
		return err
	}

	return nil

}

// Shut down everything when we're done
func teardown() {

	if textures != nil {
		textures.DestroyAll()
	}
	if renderer != nil {
		renderer.Destroy()
	}

	if window != nil {
		window.Destroy()
	}
	sdl.Quit()

}

func main() {

	err := initSDL()
	if err != nil {
		log.Fatalf("Error initializing SDL: %s", err)
	}

	var w, h int32

	// make sure to clean up
	defer teardown()

	// load the clip sheet image
	_, err = textures.LoadBMP("sheet", "sheet.bmp")
	if err != nil {
		log.Fatalf("Error loading BMP: %s", err)
	}

	// get sheet dimensions
	_, _, w, h, err = textures["sheet"].Query()
	if err != nil {
		log.Fatalf("Can't query sheet texture: %s", err)
	}

	// define the clips, one for every corner
	var clips []*sdl.Rect = []*sdl.Rect{
		&sdl.Rect{
			X: 0,
			Y: 0,
			W: w / 2,
			H: h / 2,
		},
		&sdl.Rect{
			X: w / 2,
			Y: 0,
			W: w / 2,
			H: h / 2,
		},
		&sdl.Rect{
			X: 0,
			Y: h / 2,
			W: w / 2,
			H: h / 2,
		},
		&sdl.Rect{
			X: w / 2,
			Y: h / 2,
			W: w / 2,
			H: h / 2,
		},
	}

	// prepare rendering
	renderer.SetDrawColor(0xff, 0xff, 0xff, 0xff)
	renderer.Clear()

	// render the clips
	textures.Render("sheet", renderer, 0, 0, clips[0]) // the easy one: top left
	textures.Render("sheet", renderer, screenWidth-clips[1].W, 0, clips[1])
	textures.Render("sheet", renderer, 0, screenHeight-clips[2].H, clips[2])
	textures.Render("sheet", renderer, screenWidth-clips[3].W, screenHeight-clips[3].H, clips[3])

	// and show them
	renderer.Present()

	// quit on quit event and when Escape or q are pressed
	quit := false
	var event sdl.Event
	for !quit {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				quit = true
			case *sdl.KeyboardEvent:
				if t.Type == sdl.KEYDOWN {
					switch t.Keysym.Sym {
					case sdl.K_ESCAPE:
						fallthrough
					case sdl.K_q:
						quit = true
					}
				}
			}
		}
	}

	// done!

}
