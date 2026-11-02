package activities

import (
	"fmt"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// ultratermModes is the program of the page that shows the eight modes
const ultratermModes = "../guides/listings/ultraterm-modes.bas"

// ultratermMachine is the machine of the guide, as its command line of izapple2
const ultratermMachine = `izapple2 -model none -board 2plus -cpu 6502 -screen green \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s3 videxultraterm \
    -s6 'diskii,disk1=disks/Videx Ultraterm Utilities disk.dsk'`

/*
ultratermScreenshots is the demonstration of the Videx Ultraterm, the card of
up to 160 columns, on an Apple ][+, from the disk of utilities of the card,
and then a program of the page that shows each of its eight modes in turn,
filled with a ruler.

The card blinks its cursor by the clock of the host, so the pages of the
demonstration are waited for by their text, and then by the dots of the
screen staying the same.
*/
func ultratermScreenshots(t *testing.T) {
	pictures := newAlbum("ultraterm", album.Green)
	o := start(t, ultratermMachine, nil)

	// "Videx presents", the first page
	must(t, o.WaitForText("INC. TM", 180))
	waitForStillDots(t, o, pictures)
	must(t, pictures.Screenshot(o, "presents"))

	// The introduction and the modes, and then the firmware, each page
	// written under the introduction
	must(t, o.WaitForText("160 columns by 24 lines", 180))
	waitForStillDots(t, o, pictures)
	must(t, pictures.Screenshot(o, "modes"))
	must(t, o.WaitForText("unmatched by any other card", 180))
	waitForStillDots(t, o, pictures)
	must(t, pictures.Screenshot(o, "firmware"))

	// The demonstration stopped with Reset, the card taken back with PR#3,
	// and the program of the page typed and run
	o.Reset()
	must(t, o.WaitForKeyboard(10))
	must(t, o.TypeLines("PR#3", "NEW"))
	must(t, typeListing(o, ultratermModes))
	must(t, o.TypeLines("RUN"))

	// Each mode, once its ruler is drawn, and Space for the next. The
	// seventh, 132 columns, is not taken: izapple2 shows it 160 wide (see
	// IZAPPLE2.md)
	for mode := 1; mode <= 8; mode++ {
		waitForStillDots(t, o, pictures)
		if mode != 7 {
			must(t, pictures.Screenshot(o, fmt.Sprintf("mode-%d", mode)))
		}
		must(t, o.Key("Space"))
	}
}

/*
waitForStillDots runs the machine until a screen of the Ultraterm is drawn:
there is text on it, and the lit dots have stayed nearly the same for a
second, but for the cursor, which blinks
*/
func waitForStillDots(t *testing.T, o *operator.Operator, pictures *album.Album) {
	t.Helper()
	limit := o.Frames() + 120*60
	last, still := -1, 0
	for still < 2 {
		if o.Frames() > limit {
			t.Fatal("the screen did not stop changing")
		}
		o.Run(30)
		lit := litDots(pictures.Screen(o).Pix)
		if lit > 5000 && last >= 0 && lit-last < 500 && last-lit < 500 {
			still++
		} else {
			still = 0
		}
		last = lit
	}
}

// litDots counts the dots of a screen that are lit, by their green
func litDots(pix []uint8) int {
	lit := 0
	for i := 1; i < len(pix); i += 4 {
		if pix[i] > 100 {
			lit++
		}
	}
	return lit
}
