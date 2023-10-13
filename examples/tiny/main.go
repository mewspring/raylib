// The tiny example demonstrates how to render images onto the window using the
// Draw and DrawRect methods.
package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"path/filepath"

	"github.com/mewspring/raylib/window"
	"github.com/mewspring/we"
	"github.com/pkg/errors"
)

func main() {
	if err := example(); err != nil {
		log.Fatalln(err)
	}
}

// dataDir specifies the path of the data directory (raylib/examples/data).
const dataDir = "../data"

func example() error {
	// Open a window with the specified dimensions.
	const (
		width  = 640
		height = 480
	)
	win, err := window.Open(width, height)
	if err != nil {
		return errors.WithStack(err)
	}
	defer win.Close()

	// Load background and foreground textures.
	bg, err := window.LoadTexture(filepath.Join(dataDir, "bg.png"))
	if err != nil {
		return errors.WithStack(err)
	}
	fg, err := window.LoadTexture(filepath.Join(dataDir, "fg.png"))
	if err != nil {
		return errors.WithStack(err)
	}

	// Drawing and event loop.
	for {
		// Poll events until the event queue is empty.
		for e := win.PollEvent(); e != nil; e = win.PollEvent() {
			fmt.Printf("%T: %v\n", e, e)
			switch e.(type) {
			case we.Close:
				// Close the window.
				return nil
			}
		}

		// Clear the window with white color.
		win.Clear(color.White)
		// Draw the entire background texture onto the window.
		err = win.Draw(image.ZP, bg)
		if err != nil {
			return errors.WithStack(err)
		}
		// Draw a subset of the foreground texture, as defined by the source
		// rectangle (90, 90, 225, 225), onto the window starting at the
		// destination point (10, 10).
		dp := image.Pt(10, 10)
		sr := image.Rect(90, 90, 225, 225)
		err = win.DrawRect(dp, fg, sr)
		if err != nil {
			return errors.WithStack(err)
		}

		// Display what has been rendered so far to the window.
		win.Display()
	}
}
