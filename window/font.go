package window

// #include <stdlib.h>
// #include <raylib.h>
import "C"

import (
	"runtime"
	"unsafe"
)

// A Font provides glyphs (visual characters) and metrics used for text
// rendering.
type Font struct {
	// underlying raylib font.
	_font C.Font
}

// LoadFont loads the provided TTF font.
//
// Note: a finalizer is registered to unload the font.
func LoadFont(ttfPath string) (*Font, error) {
	_ttfPath := C.CString(ttfPath)
	defer C.free(unsafe.Pointer(_ttfPath))
	_font := C.LoadFont(_ttfPath)
	// TODO: figure out how to check for error.
	font := &Font{
		_font: _font,
	}
	free := func(obj any) {
		C.UnloadFont(_font)
	}
	runtime.SetFinalizer(font, free)
	return font, nil
}
