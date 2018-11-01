// color key example inspired by http://lazyfoo.net/tutorials/SDL/10_color_keying/index.php
// Code license:		MIT
// Image license:		CC0-1.0
// 2018 Alexander Kahl <kahl@magneticcoffee.eu>
package main

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/sdl"

)

const (
	screenWidth 	= 640
	screenHeight 	= 480
	colorKey			= 0xff00ff // C-C-C-E-C ♫
)


// TextureMap holds the textures and wraps loading/destroying
type TextureMap map[string](*sdl.Texture)

var (
	window		*sdl.Window
	renderer	*sdl.Renderer
	textures	TextureMap = TextureMap{}
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

	for n := range(tm) {
		delete(tm, n)
	}

}


func initSDL() error {

	err := sdl.Init(sdl.INIT_VIDEO)

	if err != nil {
		return err
	}

	window, err = sdl.CreateWindow(
		"Moo",
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



	// load all required images into the texture map
	for _, e := range([]string{"landscape", "cow", "bubble"}) {
		_, err = textures.LoadBMP(e, fmt.Sprintf("./%s.bmp", e))
		if err != nil {
			log.Fatalf("Error loading BMP: %s", err)
		}
	}


	// prepare rendering
	renderer.Clear()

	// back to front: render the background
	renderer.Copy(textures["landscape"], nil, nil)

	// and show it
	renderer.Present()
	sdl.Delay(1000)

	// get dimensions of the cow
	_, _, w, h, err = textures["cow"].Query()
	if (err != nil) {
		log.Fatalf("Can't query cow texture: %s", err)
	}

	// specify the location and size
	pos := &sdl.Rect{100, 300, w, h}


	// draw it to the position
	renderer.Copy(textures["cow"],nil, pos)
	renderer.Present()
	sdl.Delay(2000)

	// get dimensions of the bubble
	_, _, w, h, err = textures["bubble"].Query()
	if (err != nil) {
		log.Fatalf("Can't query bubble texture: %s", err)
	}
	pos = &sdl.Rect{250, 250, w, h}
	renderer.Copy(textures["bubble"],nil, pos)
	renderer.Present()
	sdl.Delay(2000)
	
	// done!

}



