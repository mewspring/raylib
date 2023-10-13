package window

// #include <stdlib.h>
// #include <raylib.h>
import "C"

import (
	"image/color"
	"unsafe"
)

// defaultSpacing specifies the default font spacing.
const defaultSpacing = 1

// Text represent a graphical text entry with a specific font, font size, and
// colour. It implements the wandi.Image interface.
type Text struct {
	// Font to use for rendering.
	font *Font
	// Text string.
	_str *C.char
	// Font size in pixels.
	fontSize int
	// Text colour.
	c color.Color
}

// NewText returns a new graphical text entry. The initial text, font, font
// size, and colour of the graphical text entry can be customized through
// string, *Font, int, and color.Color arguments respectively, depending on the
// type of the argument.
//
// The default font, font size and colour of the text is default font, 12 and
// black, respectively.
func NewText(args ...interface{}) *Text {
	// Create a text entry.
	text := &Text{}
	// Set the default font, font size, and colour of the text.
	text.SetFont(nil) // TODO: figure out how to handle default font.
	text.SetFontSize(12)
	text.SetColor(color.Black)
	// Customize the text, font, font size, and colour based on the provided
	// arguments.
	for _, arg := range args {
		switch v := arg.(type) {
		case string:
			text.SetText(v)
		case *Font:
			text.SetFont(v)
		case int:
			text.SetFontSize(v)
		case color.Color:
			text.SetColor(v)
		}
	}
	return text
}

// SetText sets the text of the text entry.
func (text *Text) SetText(s string) {
	if text._str != nil {
		// free old string if present.
		C.free(unsafe.Pointer(text._str))
	}
	text._str = C.CString(s)
}

// SetFont sets the font of the text.
func (text *Text) SetFont(font *Font) {
	text.font = font
}

// SetFontSize sets the font size, in pixels, of the text.
func (text *Text) SetFontSize(fontSize int) {
	text.fontSize = fontSize
}

// SetColor sets the colour of the text.
func (text *Text) SetColor(c color.Color) {
	text.c = c
}

// Width returns the width of the text entry.
func (text *Text) Width() int {
	_size := C.MeasureTextEx(text.font._font, text._str, C.float(text.fontSize), defaultSpacing)
	return int(_size.x)
}

// Height returns the height of the text entry.
func (text *Text) Height() int {
	_size := C.MeasureTextEx(text.font._font, text._str, C.float(text.fontSize), defaultSpacing)
	return int(_size.y)
}
