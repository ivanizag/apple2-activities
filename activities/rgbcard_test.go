package activities

import (
	"fmt"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
	"github.com/ivanizag/izapple2/screen"
)

// rgbCardMachine is the machine of the guide, as its command line of izapple2
const rgbCardMachine = `izapple2 -model _base -board 2e -cpu 65c02 \
    -rom "<internal>/Apple2e_Enhanced.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" -rgb \
    -s0 language \
    -s6 diskii,disk1=disks/video7-rgb-demo.dsk`

/*
rgbCardScreenshots is the demonstration disk of the Video-7 RGB card on an
enhanced Apple //e: its menu, and each of the fourteen video modes of its
Video Modes part, the six of the //e and the eight the card adds.
*/
func rgbCardScreenshots(t *testing.T) {
	pictures := newAlbum("rgb-card", album.Color)
	o := start(t, rgbCardMachine, nil)

	// The menu of the demonstration, drawn in graphics
	waitForMode(t, o, screen.VideoDHGR, 60)
	waitForStillScreen(t, o, pictures)
	must(t, pictures.Screenshot(o, "menu"))

	// The menu of the video modes
	must(t, o.Type("5"))
	waitForModesMenu(t, o, pictures)
	must(t, pictures.Screenshot(o, "modes"))

	// Modes 1 and 2, the text of the //e in 40 and 80 columns, its two
	// character sets each
	for mode := 1; mode <= 2; mode++ {
		chooseMode(t, o, pictures, mode)
		must(t, o.Type("2"))
		waitForStillScreen(t, o, pictures)
		must(t, pictures.Screenshot(o, fmt.Sprintf("mode-%d", mode)))
		leaveMode(t, o, pictures)
	}

	// Mode 3, the low resolution graphics: the sixteen colours, and a
	// pattern that changes, after Space
	chooseMode(t, o, pictures, 3)
	must(t, pictures.Screenshot(o, "mode-3"))
	pattern := pictures.Record(o)
	pattern.Capture(10)
	must(t, o.Key("Space"))
	pattern.Run(10*60, 6)
	must(t, pictures.SaveRecording(pattern, "mode-3-pattern", 100))
	leaveMode(t, o, pictures)

	// Modes 4, 5 and 8 move: a game of bricks that plays itself, lines
	// drawn in high resolution, and the bricks again with coloured text
	for _, mode := range []int{4, 5, 8} {
		chooseMode(t, o, pictures, mode)
		o.RunSeconds(8)
		must(t, pictures.Screenshot(o, fmt.Sprintf("mode-%d", mode)))
		leaveMode(t, o, pictures)
	}

	// Mode 6, the six colours of high resolution, named in 80 columns
	chooseMode(t, o, pictures, 6)
	must(t, pictures.Screenshot(o, "mode-6"))
	leaveMode(t, o, pictures)

	// Mode 7, text in a colour on another: black on yellow, and the same
	// colour twice to go back to the menu
	chooseMode(t, o, pictures, 7)
	must(t, o.WaitForText("FOREGROUND COLOR", 10))
	must(t, o.TypeLines("0"))
	must(t, o.WaitForText("BACKGROUND COLOR", 10))
	must(t, o.TypeLines("13"))
	must(t, o.WaitForText("FOREGROUND COLOR", 10))
	waitForStillScreen(t, o, pictures)
	must(t, pictures.Screenshot(o, "mode-7"))
	must(t, o.TypeLines("5"))
	must(t, o.WaitForText("BACKGROUND COLOR", 10))
	must(t, o.TypeLines("5"))
	waitForModesMenu(t, o, pictures)

	// Modes 9 to 14, still pictures of each
	for mode := 9; mode <= 14; mode++ {
		chooseMode(t, o, pictures, mode)
		must(t, pictures.Screenshot(o, fmt.Sprintf("mode-%d", mode)))
		leaveMode(t, o, pictures)
	}
}

// chooseMode types the number of a mode in the menu of the modes, and waits
// for its screen to be drawn
func chooseMode(t *testing.T, o *operator.Operator, pictures *album.Album, mode int) {
	t.Helper()
	must(t, o.TypeLines(fmt.Sprint(mode)))
	o.RunSeconds(2)
	waitForStillScreen(t, o, pictures)
}

// leaveMode goes back to the menu of the modes with Escape
func leaveMode(t *testing.T, o *operator.Operator, pictures *album.Album) {
	t.Helper()
	must(t, o.Key("Escape"))
	waitForModesMenu(t, o, pictures)
}

// waitForModesMenu waits for the menu of the modes, all of it written
func waitForModesMenu(t *testing.T, o *operator.Operator, pictures *album.Album) {
	t.Helper()
	must(t, o.WaitForText("RETURNS YOU TO MAIN MENU", 60))
	waitForStillScreen(t, o, pictures)
}

/*
waitForStillScreen runs the machine until its screen stops changing: fewer
than a few hundred dots change in half a second, twice, which leaves out a
blinking cursor
*/
func waitForStillScreen(t *testing.T, o *operator.Operator, pictures *album.Album) {
	t.Helper()
	limit := o.Frames() + 60*60
	last := pictures.Screen(o)
	for still := 0; still < 2; {
		if o.Frames() > limit {
			t.Fatal("the screen did not stop changing")
		}
		o.Run(30)
		now := pictures.Screen(o)
		changed := 0
		if now.Bounds() != last.Bounds() {
			changed = len(now.Pix)
		} else {
			for i := 0; i < len(now.Pix); i += 4 {
				if now.Pix[i] != last.Pix[i] || now.Pix[i+1] != last.Pix[i+1] || now.Pix[i+2] != last.Pix[i+2] {
					changed++
				}
			}
		}
		if changed < 400 {
			still++
		} else {
			still = 0
		}
		last = now
	}
}
