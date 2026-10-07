package activities

import (
	"bytes"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// basis108Machine is the machine of the guide, as its command line of
// izapple2
const basis108Machine = `izapple2 -model none -board basis108 -cpu 6502 -screen green \
    -rom "<custom>" \
    -charrom "<internal>/D29_basis_cg_2532.rom.BIN" \
    -s0 language`

// basis108DOS is the same machine with a disk drive and the DOS 3.3 System
// Master
const basis108DOS = `izapple2 -model none -board basis108 -cpu 6502 -screen green \
    -rom "<custom>" \
    -charrom "<internal>/D29_basis_cg_2532.rom.BIN" \
    -s0 language \
    -s6 'diskii,disk1=disks/DOS 3.3 System Master - 680-0210-A (1982).dsk'`

/*
basis108Screenshots is the Basis 108, an Apple ][+ made in West Germany: in
80 columns from the start, with small letters, its four sets of characters,
ASCII, German, APL and the Apple's, chosen from BASIC, and DOS 3.3 started
in 80 columns from Apple's System Master.

izapple2 reads its text in 80 columns a column out of two, so the generator
waits for the keyboard to be read, not for a text.
*/
func basis108Screenshots(t *testing.T) {
	pictures := newAlbum("basis108", album.Green)
	o := start(t, basis108Machine, nil)

	// Switched on: its name, in 80 columns, and Applesoft
	basisReady(t, o, pictures)
	must(t, pictures.Screenshot(o, "switched-on"))

	// Small letters, and the signs of ASCII
	must(t, o.TypeLines(`PRINT "@ABC [\]^_ abc {|}~ 0123"`))
	basisReady(t, o, pictures)
	must(t, pictures.Screenshot(o, "ascii"))

	// The other sets of characters: German, APL and the Apple's
	for _, set := range []struct {
		name  string
		pokes string
	}{
		{"german", "POKE 49155,0: POKE 49156,0"},
		{"apl", "POKE 49157,0"},
		{"apple", "POKE 49154,0: POKE 49156,0"},
	} {
		must(t, o.TypeLines(set.pokes))
		basisReady(t, o, pictures)
		must(t, pictures.Screenshot(o, set.name))
	}

	// Apple's DOS 3.3, in 80 columns
	o = start(t, basis108DOS, nil)
	basisReady(t, o, pictures)
	must(t, o.TypeLines("CATALOG"))
	basisReady(t, o, pictures)
	must(t, pictures.Screenshot(o, "dos"))
}

// basisReady runs the Basis 108 until it waits for a key and its screen
// has not changed for half a second
func basisReady(t *testing.T, o *operator.Operator, pictures *album.Album) {
	t.Helper()
	must(t, o.WaitForKeyboard(60))
	before := pictures.Screen(o).Pix
	for limit := o.Frames() + 60*operator.FramesPerSecond; ; {
		o.Run(30)
		now := pictures.Screen(o).Pix
		if bytes.Equal(before, now) {
			return
		}
		if o.Frames() > limit {
			t.Fatal("the screen did not settle")
		}
		before = now
	}
}
