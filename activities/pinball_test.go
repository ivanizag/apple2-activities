package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// pinballMachine is the machine of the guide, as its command line of izapple2
const pinballMachine = `izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s6 'diskii,disk1=disks/Pinball Construction Set (4am and san inc crack).dsk'`

// The places of the page, as the hand has them: the parts box and the icons,
// where the finger touches 12 lines above the hand, and the board
var (
	pinballLeftFlipper  = [2]int{208, 12}
	pinballRightFlipper = [2]int{228, 12}
	pinballBumper       = [2]int{173, 62}
	pinballBall         = [2]int{249, 19}
	pinballPlay         = [2]int{265, 134}
	pinballDisk         = [2]int{264, 185}
	pinballLoad         = [2]int{180, 91}
	pinballPlayGame     = [2]int{180, 163}
)

/*
pinballScreenshots is Bill Budge's Pinball Construction Set of 1983: a table
built on the empty board, two flippers, two bumpers and a ball dragged from
the parts box with the hand, and played; then DEMO1, a table of the disk,
loaded and played as a game, with its spring.

The hand is moved by the joystick, faster the further the lever is from the
middle, and the operator steers it by its place in memory. The flippers are played by a player that watches the ball, the dots
that change from one picture to the next, and holds both buttons while the
ball is just above them.
*/
func pinballScreenshots(t *testing.T) {
	pictures := newAlbum("pinball-construction-set", album.Color)
	o := start(t, pinballMachine, nil)
	o.Joystick(128, 128)

	// The title, and the construction set: the board and the parts box
	o.RunSeconds(10)
	must(t, pictures.Screenshot(o, "title"))
	o.RunSeconds(20)
	must(t, pictures.Screenshot(o, "construction"))

	// A table of its own: the parts dragged onto the board
	build := pictures.Record(o)
	build.Capture(50)
	pinballDrag(t, o, build, pinballLeftFlipper, [2]int{45, 160})
	pinballDrag(t, o, build, pinballRightFlipper, [2]int{95, 160})
	pinballDrag(t, o, build, pinballBumper, [2]int{50, 90})
	pinballDrag(t, o, build, pinballBumper, [2]int{100, 70})
	pinballDrag(t, o, build, pinballBall, [2]int{70, 40})
	build.Run(60, 3)
	must(t, pictures.SaveRecording(build, "build", 300))
	must(t, pictures.Screenshot(o, "built"))

	// Played, with PLAY
	pinballClick(t, o, pinballPlay)
	play := pictures.Record(o)
	play.Capture(10)
	pinballFlippers(t, o, pictures, play, 8)
	must(t, pictures.SaveRecording(play, "play", 300))
	must(t, o.Key("Escape"))
	o.RunSeconds(3)

	// DEMO1, from the disk
	pinballClick(t, o, pinballDisk)
	o.RunSeconds(3)
	must(t, pictures.Screenshot(o, "disk"))
	pinballClick(t, o, pinballLoad)
	o.RunSeconds(2)
	must(t, o.Type("DEMO1\n"))
	o.RunSeconds(3)
	must(t, pictures.Screenshot(o, "insert"))
	must(t, o.Key("Space"))
	o.RunSeconds(15)
	must(t, pictures.Screenshot(o, "demo1"))

	// A game of it: the space bar, the button, the spring pulled down and
	// pushed up, and the button to let the ball go
	pinballClick(t, o, pinballPlayGame)
	o.RunSeconds(3)
	game := pictures.Record(o)
	game.Capture(10)
	must(t, o.Key("Space"))
	game.Run(120, 3)
	o.HoldButton(0)
	game.Run(20, 3)
	o.ReleaseButton(0)
	game.Run(120, 3)
	o.Joystick(128, 255)
	game.Run(60, 3)
	o.Joystick(128, 0)
	game.Run(60, 3)
	o.Joystick(128, 128)
	game.Run(10, 3)
	o.HoldButton(0)
	game.Run(10, 3)
	o.ReleaseButton(0)
	pinballFlippers(t, o, pictures, game, 16)
	must(t, pictures.SaveRecording(game, "game", 300))
	must(t, pictures.Screenshot(o, "score"))
}

// pinballHand is the place of the hand, alone or dragging a part: $83 and
// $84 the byte and the dot across, $82 the line, 11 above the hand
func pinballHand(o *operator.Operator) (int, int) {
	return int(o.Apple2().Peek(0x83))*7 + int(o.Apple2().Peek(0x84)), int(o.Apple2().Peek(0x82)) + 11
}

/*
pinballMove steers the hand to a place with the joystick, and stops it there,
recording if there is a recording. A part dragged stops where the walls of
the board let it, as near as it gets in ten seconds.
*/
func pinballMove(t *testing.T, o *operator.Operator, r *album.Recording, to [2]int, dragging bool) {
	t.Helper()
	for frame := 0; frame < 600; frame++ {
		x, y := pinballHand(o)
		dx, dy := to[0]-x, to[1]-y
		if dx >= -1 && dx <= 1 && dy >= -1 && dy <= 1 {
			break
		}
		o.Joystick(uint8(128+pinballLever(dx)), uint8(128+pinballLever(dy)))
		if r != nil && frame%6 == 0 {
			r.Capture(10)
		}
		o.Run(1)
		if frame == 599 && !dragging {
			t.Fatalf("the hand is at %v,%v, not at %v", x, y, to)
		}
	}
	o.Joystick(128, 128)
	o.Run(2)
	x, y := pinballHand(o)
	t.Logf("to %v: at %v,%v", to, x, y)
}

// pinballLever is how far to push the lever for a distance: past the dead
// zone of about 30 around the middle, and further for more
func pinballLever(d int) int {
	switch {
	case d > 0:
		return min(36+d/2, 120)
	case d < 0:
		return -min(36-d/2, 120)
	}
	return 0
}

// pinballDrag takes a part with the hand and drops it on the board
func pinballDrag(t *testing.T, o *operator.Operator, r *album.Recording, from, to [2]int) {
	t.Helper()
	pinballMove(t, o, r, from, false)
	o.HoldButton(0)
	r.Run(10, 5)
	pinballMove(t, o, r, to, true)
	o.ReleaseButton(0)
	r.Run(20, 5)
}

// pinballClick selects an icon or a word of a menu with the hand
func pinballClick(t *testing.T, o *operator.Operator, at [2]int) {
	t.Helper()
	pinballMove(t, o, nil, at, false)
	o.PressButton(0)
	o.Run(20)
}

/*
pinballFlippers plays the flippers for some seconds: the ball is the dots of
the board that changed since the picture before, and both buttons are held
while it is above the flippers, a few lines over them
*/
func pinballFlippers(t *testing.T, o *operator.Operator, pictures *album.Album, r *album.Recording, seconds int) {
	t.Helper()
	before := append([]uint8{}, pictures.Screen(o).Pix...)
	held := false
	for frame := 0; frame < seconds*60; frame += 2 {
		screen := pictures.Screen(o)
		width := screen.Bounds().Dx()
		changed, sumX, sumY := 0, 0, 0
		for y := 0; y < 384; y += 2 {
			for x := 0; x < 290; x += 2 {
				i := (y*width + x) * 4
				if screen.Pix[i] != before[i] || screen.Pix[i+1] != before[i+1] {
					changed++
					sumX += x
					sumY += y
				}
			}
		}
		copy(before, screen.Pix)
		ball := false
		if changed > 0 && changed <= 80 {
			x, y := sumX/changed/2, sumY/changed/2
			ball = y > 140 && y < 175 && x < 140
		}
		if ball != held {
			if ball {
				o.HoldButton(0)
				o.HoldButton(1)
			} else {
				o.ReleaseButton(0)
				o.ReleaseButton(1)
			}
			held = ball
		}
		o.Run(2)
		if frame%6 == 0 {
			r.Capture(10)
		}
	}
	o.ReleaseButton(0)
	o.ReleaseButton(1)
}
