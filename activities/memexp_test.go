package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// memexpMachine is an Apple ][+ with the Apple II Memory Expansion Card of
// 1 MB in slot 4, and ProDOS
const memexpMachine = `izapple2 -model none -board 2plus -cpu 6502 -screen green \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s0 language \
    -s4 memexp \
    -s6 diskii,disk1=disks/ProDOS_2_4_3.po`

// memexpDOS is the same machine with DOS 3.3
const memexpDOS = `izapple2 -model none -board 2plus -cpu 6502 -screen green \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s0 language \
    -s4 memexp \
    -s6 'diskii,disk1=disks/DOS 3.3 System Master - 680-0210-A (1982).dsk'`

/*
memexpScreenshots is the Apple II Memory Expansion Card on an Apple ][+: the
RAM disk ProDOS finds on it, a file copied to it and loaded from it, the
startup it can't do yet, the test in its ROM, and the card used from DOS
3.3.
*/
func memexpScreenshots(t *testing.T) {
	pictures := newAlbum("memory-expansion", album.Green)
	o := start(t, memexpMachine, nil)

	// The card, a disk for ProDOS, seen from Bitsy Bye with Tab
	must(t, o.WaitForText("BITSY", 60))
	must(t, o.WaitForPrompt(10))
	must(t, o.Key("Tab"))
	must(t, o.WaitForShownText("S4,D1:/RAM4", 30))
	must(t, o.WaitForPrompt(10))
	must(t, pictures.Screenshot(o, "bitsy-bye"))

	// BASIC.SYSTEM, from the diskette
	must(t, o.Key("Tab"))
	must(t, o.WaitForShownText("S6,D1:/PRODOS.2.4.3", 30))
	must(t, o.WaitForPrompt(10))
	for range 3 {
		must(t, o.Key("Down"))
	}
	must(t, o.Key("Return"))
	must(t, o.WaitForText("PRODOS BASIC", 120))
	must(t, o.WaitForKeyboard(30))

	// ProDOS copied to the card, and loaded from the diskette and from the
	// card, timed
	must(t, o.TypeLines(
		"CAT /RAM4",
		"BLOAD PRODOS,TSYS,A$2000",
		"CREATE /RAM4/PRODOS,TSYS",
		"BSAVE /RAM4/PRODOS,TSYS,A$2000,L17128",
		"CAT /RAM4",
	))
	must(t, o.WaitForText("BLOCKS USED:   42", 30))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "copied"))
	fromDisk := memexpTimed(t, o, "BLOAD /PRODOS.2.4.3/PRODOS,TSYS,A$2000")
	fromCard := memexpTimed(t, o, "BLOAD /RAM4/PRODOS,TSYS,A$2000")
	if fromCard*5 > fromDisk {
		t.Fatalf("the card is not faster than the diskette: %v frames against %v", fromCard, fromDisk)
	}

	// The test in the ROM of the card, at $C40A, a pass of three minutes
	must(t, o.TypeLines("CALL -151", "C40AG"))
	must(t, o.WaitForText("PASSES = 0000", 10))
	if !o.WaitUntil(240, func() bool { return o.HasText("PASSES = 0001") }) {
		t.Fatalf("the test of the card did not pass:\n%v", o.Text())
	}
	must(t, pictures.Screenshot(o, "test"))

	// Started from the card, which is not a startup disk
	must(t, o.Key("Escape"))
	o.Reset()
	must(t, o.WaitForKeyboard(10))
	must(t, o.TypeLines("HOME", "PR#4"))
	must(t, o.WaitForText("UNABLE TO START UP FROM MEMORY CARD.", 10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "startup"))

	// DOS 3.3: IN#4 makes the card a disk for it, slot 4 drive 1
	o = start(t, memexpDOS, nil)
	must(t, o.WaitForText("SYSTEM MASTER", 30))
	must(t, o.WaitForKeyboard(30))
	must(t, o.TypeLines(
		"IN#4",
		`10 PRINT "SAVED ON THE CARD"`,
		"SAVE NOTE,S4,D1",
		"CATALOG,S4,D1",
		"RUN NOTE,S4,D1",
	))
	must(t, o.WaitForText("\nSAVED ON THE CARD", 30))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "dos"))
}

// memexpTimed types a command and counts the frames until the keyboard is
// read again
func memexpTimed(t *testing.T, o *operator.Operator, command string) uint64 {
	t.Helper()
	must(t, o.TypeLines(command))
	start := o.Frames()
	must(t, o.WaitForKeyboard(60))
	frames := o.Frames() - start
	t.Logf("%v: %.2f s", command, float64(frames)/operator.FramesPerSecond)
	return frames
}
