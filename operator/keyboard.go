package operator

import (
	"fmt"
	"strings"
)

/*
The keyboard of the Apple II has no buffer: a key pressed sets a latch, and
the program clears it when it has read the key. A key pressed before that
replaces the one waiting, and is lost. The operator waits instead, as a person
does who looks at the screen: each key is pressed when the one before has been
taken.
*/

// The pace of a hand on the keyboard, in frames
const (
	// keyFrames is the time between two keys typed, about six a second
	keyFrames = 10
	// keyTakenFrames is the longest a key waits to be read by the program
	keyTakenFrames = 10 * FramesPerSecond
)

// keyboard is the izapple2.KeyboardProvider of the operator, a key waiting to
// be pressed
type keyboard struct {
	pending []uint8
	// polls counts the times the program looked for a new key
	polls uint64
}

// GetKey gives the machine the next key once the program has read the last
func (k *keyboard) GetKey(strobed bool) (uint8, bool) {
	if strobed {
		k.polls++
	}
	if !strobed || len(k.pending) == 0 {
		return 0, false
	}
	key := k.pending[0]
	k.pending = k.pending[1:]
	return key, true
}

/*
Type types a text a key at a time, at the pace of a hand. A new line is the
Return key. The Apple ][ and ][+ have no lower case, and their letters are typed
as capitals.
*/
func (o *Operator) Type(text string) error {
	for _, r := range text {
		code, err := o.charCode(r)
		if err != nil {
			return err
		}
		if err := o.press(code); err != nil {
			return err
		}
		o.Run(keyFrames)
	}
	return nil
}

// TypeLines types lines of text, each followed by Return, which is how a
// program is typed in BASIC
func (o *Operator) TypeLines(lines ...string) error {
	for _, line := range lines {
		if err := o.Type(line + "\n"); err != nil {
			return err
		}
	}
	return nil
}

/*
Key presses a key by its name: Return, Escape, Left, Right, Up, Down, Tab,
Delete or Space, or a letter with the Control key held as "Ctrl+C"
*/
func (o *Operator) Key(name string) error {
	code, err := keyCode(name)
	if err != nil {
		return err
	}
	if err := o.press(code); err != nil {
		return err
	}
	o.Run(keyFrames)
	return nil
}

/*
WaitForKeyboard runs the machine until the program looks at the keyboard for
a key, for as many seconds of the machine as given. A prompt that is on the
screen is not yet a prompt that takes keys: a command like INIT shows none
while it works, and the one before it is still there.
*/
func (o *Operator) WaitForKeyboard(seconds float64) error {
	polls := o.keyboard.polls
	if !o.WaitUntil(seconds, func() bool { return o.keyboard.polls > polls }) {
		return fmt.Errorf("the keyboard was not read in %v seconds", seconds)
	}
	return nil
}

// press hands a key to the machine and runs it until the program has read it
func (o *Operator) press(code uint8) error {
	o.keyboard.pending = append(o.keyboard.pending, code)
	if !o.WaitUntil(float64(keyTakenFrames)/FramesPerSecond, func() bool {
		return len(o.keyboard.pending) == 0
	}) {
		o.keyboard.pending = nil
		return fmt.Errorf("the key %v was not read in %v seconds", code, keyTakenFrames/FramesPerSecond)
	}
	return nil
}

// charCode is the code a character types, in capitals on a machine without
// lower case
func (o *Operator) charCode(r rune) (uint8, error) {
	switch {
	case r == '\n':
		return 13, nil
	case r < ' ' || r > '~':
		return 0, fmt.Errorf("there is no key for %q on the keyboard", r)
	case o.a.IsForceCaps() && r >= 'a' && r <= 'z':
		return uint8(r - 'a' + 'A'), nil
	}
	return uint8(r), nil
}

// keyCode is the code of a key by its name
func keyCode(name string) (uint8, error) {
	if letter, ok := strings.CutPrefix(name, "Ctrl+"); ok && len(letter) == 1 {
		c := strings.ToUpper(letter)[0]
		if c >= '@' && c <= '_' {
			return c - '@', nil
		}
	}
	codes := map[string]uint8{
		"Return": 13, "Escape": 27, "Left": 8, "Right": 21, "Up": 11,
		"Down": 10, "Tab": 9, "Delete": 127, "Space": 32,
	}
	if code, ok := codes[name]; ok {
		return code, nil
	}
	return 0, fmt.Errorf("there is no key called %q", name)
}
