// ref: https://www.raylib.com/cheatsheet/cheatsheet.html

// Package window handles window creation, drawing and events. It uses a small
// subset of the features provided by the raylib library version 4.5.
package window

// #include <stdlib.h>
// #include <raylib.h>
//
//#cgo LDFLAGS: -lraylib
import "C"

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"log"
	"os"

	"github.com/mewkiz/pkg/term"
	"github.com/mewspring/wandi"
)

var (
	// dbg is a logger with the "raylib:" prefix which logs debug messages to
	// standard error.
	dbg = log.New(os.Stderr, term.MagentaBold("raylib:")+" ", 0)
	// warn is a logger with the "raylib:" prefix which logs warning messages to
	// standard error.
	warn = log.New(os.Stderr, term.RedBold("raylib:")+" ", log.Lshortfile)
)

func init() {
	if !debug {
		dbg.SetOutput(io.Discard)
	}
}

// Enable debug output.
const debug = true

// A Window represents a graphical window capable of handling draw operations
// and window events. It implements the wandi.Window interface.
type Window struct {
}

// Open opens a new window of the specified dimensions.
//
// Note: the caller is responsible for invoking Close when finished using the
// window.
func Open(width, height int) (*Window, error) {
	const title = "raylib"
	C.SetTraceLogLevel(C.LOG_WARNING)
	C.InitWindow(C.int(width), C.int(height), C.CString(title))
	// TODO: figure out how to detect errors.
	win := &Window{}
	return win, nil
}

// Close closes the window.
func (*Window) Close() {
	C.CloseWindow()
}

// SetTitle sets the title of the window.
func (*Window) SetTitle(title string) {
	C.SetWindowTitle(C.CString(title))
}

// ShowCursor displays or hides the mouse cursor depending on the value of
// visible. It is visible by default.
func (*Window) ShowCursor(visible bool) {
	if visible {
		C.ShowCursor()
	} else {
		C.HideCursor()
	}
}

// Width returns the width of the window.
func (*Window) Width() int {
	// TODO: double-check that renderer width corresponds to window width.
	width := int(C.GetRenderWidth())
	return width
}

// Height returns the height of the window.
func (*Window) Height() int {
	// TODO: double-check that renderer width corresponds to window width.
	height := int(C.GetRenderHeight())
	return height
}

// Draw draws the entire src image onto the window starting at the destination
// point dp.
func (win *Window) Draw(dp image.Point, src wandi.Image) error {
	sr := image.Rect(0, 0, src.Width(), src.Height()) // TODO: add support for subimages; bounds.Min instead of (0,0).
	return win.DrawRect(dp, src, sr)
}

// DrawRect draws a subset of the src image, as defined by the source rectangle
// sr, onto the window starting at the destination point dp.
func (win *Window) DrawRect(dp image.Point, src wandi.Image, sr image.Rectangle) error {
	switch src := src.(type) {
	case *Texture:
		_sr := raylibRectangle(sr)
		_dp := vector2FromPoint(dp)
		_tint := raylibColor(color.White)
		C.DrawTextureRec(src._tex, _sr, _dp, _tint)
	case *Text:
		_dp := vector2FromPoint(dp)
		C.DrawTextEx(src.font._font, src._str, _dp, C.float(src.fontSize), defaultSpacing, raylibColor(src.c))
	default:
		panic(fmt.Errorf("support for image format %T not yet implemented", src))
	}
	return nil
}

// Clear clears the entire window with the given color.
func (win *Window) Clear(c color.Color) {
	C.ClearBackground(raylibColor(c))
}

// Display displays what has been rendered so far to the window.
func (*Window) Display() {
	// draw everything + SwapScreenBuffer + PollInputEvents.
	C.EndDrawing()
	// populate the input event queue.
	fillEventQueue()
}

// CursorPos returns the current cursor position within the given window.
func (*Window) CursorPos() image.Point {
	_pt := C.GetMousePosition()
	pt := image.Pt(int(_pt.x), int(_pt.y))
	return pt
}

// SetCursorPos sets the position of the cursor in the given window.
func (*Window) SetCursorPos(pt image.Point) {
	// TODO: double-check that mouse position is within window (i.e. (0,0) is
	// top-left corner of window).
	C.SetMousePosition(C.int(pt.X), C.int(pt.Y))
}

// ### [ Helper functions ] ####################################################

// raylibRectangle converts the given Go rectangle to the corresponding raylib
// rectangle.
func raylibRectangle(rect image.Rectangle) C.Rectangle {
	return C.Rectangle{
		x:      C.float(rect.Min.X),
		y:      C.float(rect.Min.Y),
		width:  C.float(rect.Dx()),
		height: C.float(rect.Dy()),
	}
}

// vector2FromPoint converts the given Go point to the corresponding raylib
// vector2.
func vector2FromPoint(pt image.Point) C.Vector2 {
	return C.Vector2{
		x: C.float(pt.X),
		y: C.float(pt.Y),
	}
}

// raylibColor converts the given Go color to the corresponding raylib color.
func raylibColor(c color.Color) C.Color {
	r, g, b, a := c.RGBA()
	return C.Color{
		r: C.uchar(r & 0xFF),
		g: C.uchar(g & 0xFF),
		b: C.uchar(b & 0xFF),
		a: C.uchar(a & 0xFF),
	}
}

// Ensure that Window implements wandi.Window.
var _ wandi.Window = (*Window)(nil)
