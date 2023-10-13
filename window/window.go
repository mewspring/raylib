// Package window handles window creation, drawing and events. It uses a small
// subset of the features provided by the raylib library version 4.5.
package window

// #include <raylib.h>
//
//#cgo LDFLAGS: -lraylib
import "C"

import (
	"image"
	"io"
	"log"
	"os"

	"github.com/mewkiz/pkg/term"
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
// and window events.
type Window struct {
}

// Open opens a new window of the specified dimensions.
//
// Note: the caller is responsible for invoking Close when finished using the
// window.
func Open(width, height int) (*Window, error) {
	const title = "raylib"
	C.InitWindow(C.int(width), C.int(height), C.CString(title))
	// TODO: figure out how to detect errors.
	win := &Window{}
	return win, nil
}

// Close closes the window.
func (*Window) Close() error {
	C.CloseWindow()
	// TODO: figure out how to detect errors.
	return nil
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

// ref: https://www.raylib.com/cheatsheet/cheatsheet.html
