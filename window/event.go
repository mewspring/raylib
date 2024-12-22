package window

// #include <raylib.h>
import "C"

import (
	"fmt"
	"image"

	"github.com/gammazero/deque"
	"github.com/mewpkg/clog"
	"github.com/mewspring/we"
)

func init() {
	// initialize keyFromRaylibKey.
	for key, raylibKey := range raylibKeyFromKey {
		keyFromRaylibKey[raylibKey] = key
	}
	// initialize modFromRaylibKey.
	for mod, raylibKeys := range raylibKeysFromMod {
		for _, raylibKey := range raylibKeys {
			modFromRaylibKey[raylibKey] = mod
		}
	}
	// initialize mouseButtonFromRaylibMouseButton.
	for button, raylibButton := range raylibMouseButtonFromMouseButton {
		mouseButtonFromRaylibMouseButton[raylibButton] = button
	}
}

// eventQueue holds pending input events.
var eventQueue = deque.New[we.Event]()

// PollEvent returns a pending event from the event queue or nil if the queue
// was empty. Note that more than one event may be present in the event queue.
//
// Note: the event queue is populated once per frame upon call to
// Window.Display.
func (*Window) PollEvent() we.Event {
	if eventQueue.Len() > 0 {
		e := eventQueue.PopFront()
		return e
	}
	// no pending input events.
	return nil
}

// --- [ keyboard modifier ] ---------------------------------------------------

// getModState returns a bitfield with the current state of the keyboard
// modifiers.
func getModState() we.Mod {
	var mod we.Mod
	for m := we.ModFirst; m <= we.ModLast; m <<= 1 {
		if isModPressed(m) {
			mod |= m
		}
	}
	return mod
}

// isModPressed reports whether all of the given keyboard modifiers are pressed.
func isModPressed(mods ...we.Mod) bool {
	for _, mod := range mods {
		raylibKeys, ok := raylibKeysFromMod[mod]
		if !ok {
			clog.Warnf("support for keyboard modifier %v using raylib not yet implemented", mod)
			return false
		}
		if !isModifierPressed(raylibKeys) {
			return false
		}
	}
	return true
}

// isModifierPressed reports whether any of the given keyboard keys are pressed,
// thus indicating that a given modifier is pressed.
//
// Example:
//
//	mod  = "left shift"
//	keys = ["left shift", "right shift"]
func isModifierPressed(raylibKeys []raylibKeyType) bool {
	for _, raylibKey := range raylibKeys {
		if C.IsKeyDown(raylibKey) {
			return true
		}
	}
	return false
}

var (
	// modFromRaylibKey maps from raylib keyboard key to keyboard modifier.
	// Initialized by init.
	modFromRaylibKey = make(map[raylibKeyType]we.Mod)
	// raylibKeyFromMod maps from keyboard modifier to raylib keyboard modifiers.
	raylibKeysFromMod = map[we.Mod][]raylibKeyType{
		we.ModAlt:     []raylibKeyType{C.KEY_LEFT_ALT, C.KEY_RIGHT_ALT},
		we.ModControl: []raylibKeyType{C.KEY_LEFT_CONTROL, C.KEY_RIGHT_CONTROL},
		// "windows key"
		we.ModSuper: []raylibKeyType{C.KEY_LEFT_SUPER, C.KEY_RIGHT_SUPER},
		we.ModShift: []raylibKeyType{C.KEY_LEFT_SHIFT, C.KEY_RIGHT_SHIFT},
	}
)

// --- [ keyboard key ] --------------------------------------------------------

// Alias for raylib keyboard key and modifier types.
type raylibKeyType = C.int

var (
	// keyFromRaylibKey maps from raylib keyboard key to keyboard key.
	// Initialized by init.
	keyFromRaylibKey = make(map[raylibKeyType]we.Key)
	// raylibKeyFromKey maps from keyboard key to raylib keyboard key.
	raylibKeyFromKey = map[we.Key]raylibKeyType{
		we.Key0:            C.KEY_ZERO,          // '0'
		we.Key1:            C.KEY_ONE,           // '1'
		we.Key2:            C.KEY_TWO,           // '2'
		we.Key3:            C.KEY_THREE,         // '3'
		we.Key4:            C.KEY_FOUR,          // '4'
		we.Key5:            C.KEY_FIVE,          // '5'
		we.Key6:            C.KEY_SIX,           // '6'
		we.Key7:            C.KEY_SEVEN,         // '7'
		we.Key8:            C.KEY_EIGHT,         // '8'
		we.Key9:            C.KEY_NINE,          // '9'
		we.KeyA:            C.KEY_A,             // 'a'
		we.KeyB:            C.KEY_B,             // 'b'
		we.KeyC:            C.KEY_C,             // 'c'
		we.KeyD:            C.KEY_D,             // 'd'
		we.KeyE:            C.KEY_E,             // 'e'
		we.KeyF:            C.KEY_F,             // 'f'
		we.KeyG:            C.KEY_G,             // 'g'
		we.KeyH:            C.KEY_H,             // 'h'
		we.KeyI:            C.KEY_I,             // 'i'
		we.KeyJ:            C.KEY_J,             // 'j'
		we.KeyK:            C.KEY_K,             // 'k'
		we.KeyL:            C.KEY_L,             // 'l'
		we.KeyM:            C.KEY_M,             // 'm'
		we.KeyN:            C.KEY_N,             // 'n'
		we.KeyO:            C.KEY_O,             // 'o'
		we.KeyP:            C.KEY_P,             // 'p'
		we.KeyQ:            C.KEY_Q,             // 'q'
		we.KeyR:            C.KEY_R,             // 'r'
		we.KeyS:            C.KEY_S,             // 's'
		we.KeyT:            C.KEY_T,             // 't'
		we.KeyU:            C.KEY_U,             // 'u'
		we.KeyV:            C.KEY_V,             // 'v'
		we.KeyW:            C.KEY_W,             // 'w'
		we.KeyX:            C.KEY_X,             // 'x'
		we.KeyY:            C.KEY_Y,             // 'y'
		we.KeyZ:            C.KEY_Z,             // 'z'
		we.KeyGraveAccent:  C.KEY_GRAVE,         // '`'
		we.KeyMinus:        C.KEY_MINUS,         // '-'
		we.KeyEqual:        C.KEY_EQUAL,         // '='
		we.KeyLeftBracket:  C.KEY_LEFT_BRACKET,  // '['
		we.KeyRightBracket: C.KEY_RIGHT_BRACKET, // ']'
		we.KeyBackslash:    C.KEY_BACKSLASH,     // '\\'
		we.KeySemicolon:    C.KEY_SEMICOLON,     // ';'
		we.KeyApostrophe:   C.KEY_APOSTROPHE,    // '\''
		we.KeyPeriod:       C.KEY_PERIOD,        // '.'
		we.KeyComma:        C.KEY_COMMA,         // ','
		we.KeySlash:        C.KEY_SLASH,         // '/'
		we.KeySpace:        C.KEY_SPACE,         // ' '
		we.KeyTab:          C.KEY_TAB,           // '\t'
		we.KeyEscape:       C.KEY_ESCAPE,
		we.KeyF1:           C.KEY_F1,
		we.KeyF2:           C.KEY_F2,
		we.KeyF3:           C.KEY_F3,
		we.KeyF4:           C.KEY_F4,
		we.KeyF5:           C.KEY_F5,
		we.KeyF6:           C.KEY_F6,
		we.KeyF7:           C.KEY_F7,
		we.KeyF8:           C.KEY_F8,
		we.KeyF9:           C.KEY_F9,
		we.KeyF10:          C.KEY_F10,
		we.KeyF11:          C.KEY_F11,
		we.KeyF12:          C.KEY_F12,
		we.KeyPrintScreen:  C.KEY_PRINT_SCREEN,
		we.KeyScrollLock:   C.KEY_SCROLL_LOCK,
		we.KeyPause:        C.KEY_PAUSE,
		we.KeyInsert:       C.KEY_INSERT,
		we.KeyDelete:       C.KEY_DELETE,
		we.KeyHome:         C.KEY_HOME,
		we.KeyEnd:          C.KEY_END,
		we.KeyPageUp:       C.KEY_PAGE_UP,
		we.KeyPageDown:     C.KEY_PAGE_DOWN,
		we.KeyBackspace:    C.KEY_BACKSPACE,
		we.KeyCapsLock:     C.KEY_CAPS_LOCK,
		we.KeyEnter:        C.KEY_ENTER,
		we.KeyMenu:         C.KEY_KB_MENU,
		we.KeyUp:           C.KEY_UP,
		we.KeyDown:         C.KEY_DOWN,
		we.KeyLeft:         C.KEY_LEFT,
		we.KeyRight:        C.KEY_RIGHT,
		we.KeyKp0:          C.KEY_KP_0,
		we.KeyKp1:          C.KEY_KP_1,
		we.KeyKp2:          C.KEY_KP_2,
		we.KeyKp3:          C.KEY_KP_3,
		we.KeyKp4:          C.KEY_KP_4,
		we.KeyKp5:          C.KEY_KP_5,
		we.KeyKp6:          C.KEY_KP_6,
		we.KeyKp7:          C.KEY_KP_7,
		we.KeyKp8:          C.KEY_KP_8,
		we.KeyKp9:          C.KEY_KP_9,
		we.KeyKpAdd:        C.KEY_KP_ADD,
		we.KeyKpDecimal:    C.KEY_KP_DECIMAL,
		we.KeyKpDivide:     C.KEY_KP_DIVIDE,
		we.KeyKpEnter:      C.KEY_KP_ENTER,
		we.KeyKpEqual:      C.KEY_KP_EQUAL,
		we.KeyKpMultiply:   C.KEY_KP_MULTIPLY,
		we.KeyKpSubtract:   C.KEY_KP_SUBTRACT,
		we.KeyNumLock:      C.KEY_NUM_LOCK,
		we.KeyLeftShift:    C.KEY_LEFT_SHIFT,
		we.KeyRightShift:   C.KEY_RIGHT_SHIFT,
	}
)

// --- [ mouse input ] ---------------------------------------------------------

// Alias for raylib mouse button type.
type raylibMouseButtonType = C.int

var (
	// mouseButtonFromRaylibMouseButton maps from raylib mouse button to mouse
	// button. Initialized by init.
	mouseButtonFromRaylibMouseButton = make(map[raylibMouseButtonType]we.Button)
	// raylibMouseButtonFromMouseButton maps from mouse button to raylib mouse
	// button.
	raylibMouseButtonFromMouseButton = map[we.Button]raylibMouseButtonType{
		we.ButtonLeft:   C.MOUSE_BUTTON_LEFT,
		we.ButtonRight:  C.MOUSE_BUTTON_RIGHT,
		we.ButtonMiddle: C.MOUSE_BUTTON_MIDDLE,
	}
)

// --- [ fill event queue ] ----------------------------------------------------

// fillEventQueue fills the input event queue with input events received since
// last call to Window.Display.
//
// Note: fillEventQueue should be invoked at most once each frame (after call to
// EndDrawing in Window.Display), and must be invoked before polling events
// through Window.PollEvent.
func fillEventQueue() {
	mod := getModState()
	// fill keyboard key press events since last frame.
	const (
		raylibKeyMin raylibKeyType = 1   // `key > 0` in IsKeyPressed (ref: raylib/src/rcore.c)
		raylibKeyMax raylibKeyType = 512 // MAX_KEYBOARD_KEYS=512 (ref: raylib/src/rcore.c)
	)
	for raylibKey := raylibKeyMin; raylibKey < raylibKeyMax; raylibKey++ {
		if C.IsKeyPressed(raylibKey) {
			key, ok := keyFromRaylibKey[raylibKey]
			if !ok {
				if _, ok := modFromRaylibKey[raylibKey]; ok {
					// skip keyboard modifiers
					continue
				}
				clog.Warnf("support for raylib key %v not yet implemented", raylibKey)
			}
			event := we.KeyPress{
				Key: key,
				Mod: mod,
			}
			eventQueue.PushBack(event)
		}
	}
	// fill keyboard key release events since last frame.
	for raylibKey := raylibKeyMin; raylibKey < raylibKeyMax; raylibKey++ {
		if C.IsKeyReleased(raylibKey) {
			key, ok := keyFromRaylibKey[raylibKey]
			if !ok {
				if _, ok := modFromRaylibKey[raylibKey]; ok {
					// skip keyboard modifiers
					continue
				}
				clog.Warnf("support for raylib key %v not yet implemented", raylibKey)
			}
			event := we.KeyRelease{
				Key: key,
				Mod: mod,
			}
			eventQueue.PushBack(event)
		}
	}

	// fill mouse button press events since last frame.
	const (
		raylibMouseButtonMin raylibMouseButtonType = C.MOUSE_BUTTON_LEFT
		raylibMouseButtonMax raylibMouseButtonType = C.MOUSE_BUTTON_MIDDLE
	)
	_mousePos := C.GetMousePosition()
	mousePos := image.Pt(int(_mousePos.x), int(_mousePos.y))
	for raylibButton := raylibMouseButtonMin; raylibButton <= raylibMouseButtonMax; raylibButton++ {
		if C.IsMouseButtonPressed(raylibButton) {
			button, ok := mouseButtonFromRaylibMouseButton[raylibButton]
			if !ok {
				panic(fmt.Errorf("support for raylib mouse button %v not yet implemented", raylibButton))
			}
			event := we.MousePress{
				Point:  mousePos,
				Button: button,
				Mod:    mod,
			}
			eventQueue.PushBack(event)
		}
	}
	// fill mouse button release events since last frame.
	for raylibButton := raylibMouseButtonMin; raylibButton <= raylibMouseButtonMax; raylibButton++ {
		if C.IsMouseButtonReleased(raylibButton) {
			button, ok := mouseButtonFromRaylibMouseButton[raylibButton]
			if !ok {
				panic(fmt.Errorf("support for raylib mouse button %v not yet implemented", raylibButton))
			}
			event := we.MouseRelease{
				Point:  mousePos,
				Button: button,
				Mod:    mod,
			}
			eventQueue.PushBack(event)
		}
	}
	// fill mouse movement events since last frame.

	// TODO: implement we.MouseDrag (record mouse position on mouse press and
	// check if the recorded position differs on mouse release).

	if mousePos != prevMousePos {
		// Mouse movement detected.
		event := we.MouseMove{
			Point: mousePos,
			From:  prevMousePos,
			//Mod:   mod, // TODO: add Mod to we.MouseMove?
		}
		eventQueue.PushBack(event)

		prevMousePos = mousePos
	}
	// fill typed rune events since last frame.
	for {
		char := C.GetCharPressed()
		if char == 0 {
			break
		}
		event := we.KeyRune(char)
		eventQueue.PushBack(event)
	}
}

// prevMousePos tracks the position of the mouse cursor at the previous frame.
var prevMousePos image.Point
