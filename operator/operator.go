/*
Package operator is someone sitting at an Apple II that izapple2 emulates: hands
on the keyboard, the paddles and the mouse, and eyes on the screen. It runs the
machine itself, a frame at a time, so that what the machine does depends only
on what it is given and when, and not on how fast the host is.

The timings are frames of the machine, sixtieths of a second of its time,
which runs as fast as the host can go.
*/
package operator

import (
	"fmt"
	"strings"

	"github.com/ivanizag/izapple2"
	"github.com/ivanizag/izapple2/screen"
)

// CyclesPerFrame is a frame of NTSC video, 65 cycles for each of 262 lines
const CyclesPerFrame = 17030

// FramesPerSecond is how many frames make a second of the machine
const FramesPerSecond = 60

/*
Operator drives one machine. It is the keyboard, the paddles and the mouse of
the machine, and the only one that runs it.
*/
type Operator struct {
	a        *izapple2.Apple2
	keyboard *keyboard
	paddles  *paddles
	mouse    *mouse
	frames   uint64
}

/*
New sits an operator at a machine just built, with izapple2.CreateAppleFromModel.
It connects the keyboard, the paddles and the mouse, and switches the machine
on. Nothing runs until Run is called.
*/
func New(a *izapple2.Apple2) *Operator {
	o := &Operator{
		a:        a,
		keyboard: &keyboard{},
		paddles:  &paddles{},
		mouse:    &mouse{},
	}
	a.SetKeyboardProvider(o.keyboard)
	a.SetJoysticksProvider(o.paddles)
	a.SetMouseProvider(o.mouse)
	a.Init()
	return o
}

// Start builds a machine of a model, with the configuration changed as given
// and the files on its drives, and sits an operator at it
func Start(model string, overrides map[string]string, files ...string) (*Operator, error) {
	a, err := izapple2.CreateAppleFromModel(model, overrides, files)
	if err != nil {
		return nil, fmt.Errorf("could not build a %v: %w", model, err)
	}
	return New(a), nil
}

// Apple2 is the machine the operator is at
func (o *Operator) Apple2() *izapple2.Apple2 {
	return o.a
}

// Frames is how many frames the machine has run
func (o *Operator) Frames() uint64 {
	return o.frames
}

// Run lets the machine run for a number of frames
func (o *Operator) Run(frames int) {
	for range frames {
		o.a.RunCycles(CyclesPerFrame)
		o.frames++
	}
}

// RunSeconds lets the machine run for a number of seconds of its time
func (o *Operator) RunSeconds(seconds float64) {
	o.Run(int(seconds * FramesPerSecond))
}

/*
WaitUntil runs the machine a frame at a time until something has happened, for
as many seconds of the machine as given, and tells whether it did
*/
func (o *Operator) WaitUntil(seconds float64, done func() bool) bool {
	limit := int(seconds * FramesPerSecond)
	for range limit {
		if done() {
			return true
		}
		o.Run(1)
	}
	return done()
}

// Text is the text on the screen, a line per row
func (o *Operator) Text() string {
	return o.a.ScreenText()
}

// WaitForText runs the machine until a text is on the screen, for as many
// seconds of the machine as given
func (o *Operator) WaitForText(text string, seconds float64) error {
	if !o.WaitUntil(seconds, func() bool { return strings.Contains(o.Text(), text) }) {
		return fmt.Errorf("%q was not on the screen after %v seconds:\n%v", text, seconds, o.Text())
	}
	return nil
}

// Reset presses Control and Reset, which stops the program running and,
// depending on the machine and what it has loaded, goes back to BASIC or
// starts again
func (o *Operator) Reset() {
	o.a.SendCommand(izapple2.CommandReset)
	o.Run(resetFrames)
}

// InsertDisk puts a disk in a drive, numbered as izapple2 lists them: the
// first drive of slot 6 is 0
func (o *Operator) InsertDisk(drive int, path string) error {
	return o.a.LoadDisk(drive, path)
}

// resetFrames is how long the reset key is held, and then let go
const resetFrames = 10

// HasText says whether a text is on the screen now
func (o *Operator) HasText(text string) bool {
	return strings.Contains(o.Text(), text)
}

// InTextMode says whether the screen shows text, in 40 or 80 columns, and not
// graphics: the text of the text page is there even when it is not shown
func (o *Operator) InTextMode() bool {
	mode := o.a.GetVideoSource().GetCurrentVideoMode() & screen.VideoBaseMask
	return mode == screen.VideoText40 || mode == screen.VideoText80 || mode == screen.VideoVidex
}

// WaitForShownText runs the machine until a text is on the screen and the
// screen shows text, for as many seconds of the machine as given
func (o *Operator) WaitForShownText(text string, seconds float64) error {
	if !o.WaitUntil(seconds, func() bool { return o.InTextMode() && o.HasText(text) }) {
		return fmt.Errorf("%q was not shown after %v seconds", text, seconds)
	}
	return nil
}
