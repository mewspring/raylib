// The input_events example creates a new window and handles input events.
package main

import (
	"fmt"
	"log"

	"github.com/mewspring/raylib/window"
	"github.com/mewspring/we"
	"github.com/pkg/errors"
)

func main() {
	if err := example(); err != nil {
		log.Fatalf("%+v", err)
	}
}

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

		// NOTE: logic updates goes here.

		// NOTE: drawing goes here.

		// Display what has been rendered so far to the window (and populate input
		// event queue).
		win.Display()
	}
	return nil
}
