package operator

/*
The mouse of the Apple II mouse card. izapple2 gives the program where the
pointer is rather than how far the mouse moved, so a place is reached at once.
Positions are fractions of the screen, from 0 to 1 across and down, as the
program decides how many places across it wants.
*/

// mouse is the izapple2.MouseProvider of the operator
type mouse struct {
	x, y    uint16
	pressed bool
}

// ReadMouse is where the mouse is, in the full range of 16 bits, and whether
// its button is down
func (m *mouse) ReadMouse() (uint16, uint16, bool) {
	return m.x, m.y, m.pressed
}

// The pace of a hand on the mouse, in frames
const (
	mouseMoveFrames = 4
	clickDownFrames = 8
	clickUpFrames   = 20
	doubleGapFrames = 6
)

// MouseTo moves the mouse to a place on the screen, from 0 to 1 across and
// down
func (o *Operator) MouseTo(x float64, y float64) {
	o.mouse.x = toMouseRange(x)
	o.mouse.y = toMouseRange(y)
	o.Run(mouseMoveFrames)
}

func toMouseRange(f float64) uint16 {
	return uint16(max(min(f*65536, 65535), 0))
}

// Hold presses the button of the mouse and leaves it down, and Release lets
// go of it
func (o *Operator) Hold() {
	o.mouse.pressed = true
}

// Release lets go of the button of the mouse
func (o *Operator) Release() {
	o.mouse.pressed = false
}

// Click presses and releases the button of the mouse where it is
func (o *Operator) Click() {
	o.Hold()
	o.Run(clickDownFrames)
	o.Release()
	o.Run(clickUpFrames)
}

// DoubleClick clicks twice, close enough to be one double click
func (o *Operator) DoubleClick() {
	o.Hold()
	o.Run(clickDownFrames)
	o.Release()
	o.Run(doubleGapFrames)
	o.Click()
}
