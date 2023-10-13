// The fonts example demonstrates how to render text using TTF fonts.
package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"path/filepath"
	"time"

	"github.com/mewspring/raylib/window"
	"github.com/mewspring/we"
	"github.com/pkg/errors"
)

func main() {
	err := fonts()
	if err != nil {
		log.Fatalln(err)
	}
}

// dataDir specifies the path of the data directory (raylib/examples/data).
const dataDir = "../data"

// fonts demonstrates how to render text using TTF fonts.
func fonts() (err error) {
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

	// Load background texture.
	bg, err := window.LoadTexture(filepath.Join(dataDir, "bg2.png"))
	if err != nil {
		return errors.WithStack(err)
	}
	// Load the text TTF font.
	textFont, err := window.LoadFont(filepath.Join(dataDir, "Exocet.ttf"))
	if err != nil {
		return errors.WithStack(err)
	}
	// Create a new graphical text entry based on the Exocet TTF font and
	// initialize its text to "TTF fonts", its font size to 32 (the default is
	// 12) and its color to white (the default is black).
	text := window.NewText(textFont, "TTF fonts", 32, color.White)
	// Load the fps TTF font.
	fpsFont, err := window.LoadFont(filepath.Join(dataDir, "DejaVuSansMono.ttf"))
	if err != nil {
		return errors.WithStack(err)
	}
	// Create a graphical FPS text entry. The text of this graphical text entry
	// will be updated repeatedly using SetText.
	fps := window.NewText(fpsFont, 14, color.White)

	// start and frames will be used to calculate the average FPS of the
	// application.
	start := time.Now()
	frames := 0.0

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
		if err := win.Draw(image.Pt(0, 0), bg); err != nil {
			return errors.WithStack(err)
		}
		// Draw the entire text onto the window starting the destination point
		// (420, 12).
		dp := image.Pt(420, 12)
		if err := win.Draw(dp, text); err != nil {
			return errors.WithStack(err)
		}
		// Update the text of the FPS text entry.
		fps.SetText(getFPS(start, frames))
		// Draw the entire FPS text entry onto the screen starting at the
		// destination point (8, 4).
		dp = image.Pt(8, 4)
		if err := win.Draw(dp, fps); err != nil {
			return errors.WithStack(err)
		}

		// Display what has been rendered so far to the window.
		win.Display()
		frames++
	}
}

// getFPS returns the average FPS as a string, based on the provided start time
// and frame count.
func getFPS(start time.Time, frames float64) (text string) {
	// Average FPS.
	fps := frames / time.Since(start).Seconds()
	return fmt.Sprintf("FPS: %.2f", fps)
}
