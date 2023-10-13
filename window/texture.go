package window

// #include <stdlib.h>
// #include <raylib.h>
import "C"

import (
	"image"
	"image/draw"
	"runtime"
	"time"
	"unsafe"

	"github.com/mewspring/wandi"
)

// Texture represent a read-only texture. It implements the wandi.Image
// interface.
type Texture struct {
	// A read-only GPU texture.
	_tex C.Texture2D
}

// LoadTexture loads the provided file and converts it into a read-only texture.
//
// Note: a finalizer is registered to unload the texture.
func LoadTexture(path string) (*Texture, error) {
	// Load the texture from file.
	_path := C.CString(path)
	defer C.free(unsafe.Pointer(_path))
	_tex := C.LoadTexture(_path)
	// TODO: figure out how to check error.
	tex := newTexture(_tex)
	return tex, nil
}

// LoadTextureFromImage reads the provided image and converts it into a
// read-only texture.
//
// Note: a finalizer is registered to unload the texture.
func LoadTextureFromImage(src image.Image) (*Texture, error) {
	// Use fallback conversion for unknown image formats.
	rgba, ok := src.(*image.RGBA)
	if !ok {
		return LoadTextureFromImage(fallbackRGBAImage(src))
	}
	// Use fallback conversion for subimages.
	bounds := rgba.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	const npixelBytes = 4 // RGBA
	if rgba.Stride != npixelBytes*width {
		return LoadTextureFromImage(fallbackRGBAImage(src))
	}
	// Create a read-only texture based on the pixels of the src image.
	pix := unsafe.Pointer(&rgba.Pix[0])
	_img := C.Image{
		data:    pix,
		width:   C.int(width),
		height:  C.int(height),
		mipmaps: 1,
		format:  C.PIXELFORMAT_UNCOMPRESSED_R8G8B8A8,
	}
	_tex := C.LoadTextureFromImage(_img)
	// TODO: figure out how to check error.
	tex := newTexture(_tex)
	return tex, nil
}

// Width returns the width of the texture.
func (tex *Texture) Width() int {
	return int(tex._tex.width)
}

// Height returns the height of the texture.
func (tex *Texture) Height() int {
	return int(tex._tex.height)
}

// ### [ Helper functions ] ####################################################

// newTexture returns a new read-only texture.
//
// Note: a finalizer is registered to unload the texture.
func newTexture(_tex C.Texture2D) *Texture {
	tex := &Texture{
		_tex: _tex,
	}
	// Register finalizer to unload texture.
	free := func(obj any) {
		C.UnloadTexture(_tex)
	}
	runtime.SetFinalizer(tex, free)
	return tex
}

// fallbackRGBAImage converts the provided image or subimage into an RGBA image.
func fallbackRGBAImage(src image.Image) *image.RGBA {
	start := time.Now()
	// Create a new RGBA image and draw the src image onto it.
	bounds := src.Bounds()
	dr := image.Rect(0, 0, bounds.Dx(), bounds.Dy())
	dst := image.NewRGBA(dr)
	draw.Draw(dst, dr, src, bounds.Min, draw.Src)
	warn.Printf("fallback conversion for non-RGBA image (%T) finished in: %v", src, time.Since(start))
	return dst
}

// Ensure that Texture implements wandi.Image.
var _ wandi.Image = (*Texture)(nil)
