package operator

/*
The game port: four paddles, or two joysticks of two paddles each, and three
buttons. On the Apple //e the first two buttons are also the Open Apple and
the Closed Apple keys. A paddle not plugged in reads as turned to the end.
*/

// paddles is the izapple2.JoysticksProvider of the operator
type paddles struct {
	value   [4]uint8
	plugged [4]bool
	button  [3]bool
}

// ReadButton says whether a button is held down
func (p *paddles) ReadButton(i int) bool {
	return i < len(p.button) && p.button[i]
}

// ReadPaddle is where a paddle is turned to, from 0 to 255, if it is plugged
func (p *paddles) ReadPaddle(i int) (uint8, bool) {
	if i >= len(p.value) {
		return 0, false
	}
	return p.value[i], p.plugged[i]
}

// TurnPaddle turns a paddle, 0 to 3, to a value from 0 to 255. A joystick is
// two paddles, 0 and 1 for the first, its left and right and its up and
// down.
func (o *Operator) TurnPaddle(paddle int, value uint8) {
	o.paddles.value[paddle] = value
	o.paddles.plugged[paddle] = true
}

// Joystick moves the first joystick, each axis from 0 to 255 and 127 in the
// middle
func (o *Operator) Joystick(x uint8, y uint8) {
	o.TurnPaddle(0, x)
	o.TurnPaddle(1, y)
}

// HoldButton presses a button of the game port, 0 to 2, and leaves it down,
// and ReleaseButton lets go of it. Buttons 0 and 1 are the Open Apple and the
// Closed Apple keys of the Apple //e.
func (o *Operator) HoldButton(button int) {
	o.paddles.button[button] = true
}

// ReleaseButton lets go of a button of the game port
func (o *Operator) ReleaseButton(button int) {
	o.paddles.button[button] = false
}

// PressButton presses and releases a button of the game port
func (o *Operator) PressButton(button int) {
	o.HoldButton(button)
	o.Run(buttonFrames)
	o.ReleaseButton(button)
	o.Run(buttonFrames)
}

// buttonFrames is how long a button is held when pressed
const buttonFrames = 8
