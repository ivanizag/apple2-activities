package activities

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// thunderclockMachine is the machine of the guide, as its command line of
// izapple2: an Apple ][+ with a ThunderClock Plus and ProDOS. What it writes
// is kept in the folder changes, and the image left as it is.
const thunderclockMachine = `izapple2 -model none -board 2plus -cpu 6502 -screen green \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -saveDir changes \
    -s0 language \
    -s4 thunderclock \
    -s6 diskii,disk1=disks/ProDOS_2_4_3.po`

// clockListing is the program of the page, the time read from the card
const clockListing = "../guides/listings/clock.bas"

/*
thunderclockScreenshots is the ThunderClock Plus on an Apple ][+: the time
read from BASIC through the firmware of the card, and a program saved with
the date and the time ProDOS takes from it. The pictures show the time they
were made at.
*/
func thunderclockScreenshots(t *testing.T) {
	pictures := newAlbum("thunderclock", album.Green)
	changes := filepath.Join(t.TempDir(), "changes")
	must(t, os.Mkdir(changes, 0o755))
	o := start(t, thunderclockMachine, map[string]string{"changes": changes})

	// ProDOS, and BASIC.SYSTEM chosen in Bitsy Bye
	must(t, o.WaitForText("BITSY", 60))
	must(t, o.WaitForPrompt(10))
	for range 3 {
		must(t, o.Key("Down"))
	}
	must(t, o.Key("Return"))
	must(t, o.WaitForText("PRODOS BASIC", 120))
	must(t, o.WaitForKeyboard(30))

	// The time, read from the card
	must(t, typeListing(o, clockListing))
	must(t, o.TypeLines("RUN"))
	if !o.WaitUntil(10, func() bool { return o.HasText(" AM\n") || o.HasText(" PM\n") }) {
		t.Fatalf("the time was not read:\n%v", o.Text())
	}
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "time"))

	// The program saved, and its date and time in the catalog
	must(t, o.TypeLines("SAVE CLOCK", "CAT"))
	must(t, o.WaitForText("BLOCKS FREE", 30))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "catalog"))
}
