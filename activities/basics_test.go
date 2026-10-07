package activities

import (
	"strings"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// basicsTape is the first Apple ][ with the tape of Applesoft in its
// cassette recorder
const basicsTape = `izapple2 -model none -board 2plus -cpu 6502 -screen green \
    -rom "<internal>/341-000x_integer.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -tape disks/k7_apple_600200600_applesoftiia.wav`

// basicsPlus is an Apple ][+ with a Language Card and DOS 3.3
const basicsPlus = `izapple2 -model none -board 2plus -cpu 6502 -screen green \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s0 language \
    -s6 'diskii,disk1=disks/DOS 3.3 System Master - 680-0210-A (1982).dsk'`

// basicsInteger is the first Apple ][ with a Language Card and DOS 3.3
const basicsInteger = `izapple2 -model none -board 2plus -cpu 6502 -screen green \
    -rom "<internal>/341-000x_integer.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s0 language \
    -s6 'diskii,disk1=disks/DOS 3.3 System Master - 680-0210-A (1982).dsk'`

/*
basicsScreenshots is how the Apple ][ went from Integer BASIC to Applesoft:
Integer BASIC counting in whole numbers, Applesoft loaded from its tape into
the first Apple ][, and then, with the Language Card, the two BASICs on one
machine, the one of its ROM and the other loaded by DOS 3.3 into the card,
on an Apple ][+ and on the first Apple ][.
*/
func basicsScreenshots(t *testing.T) {
	pictures := newAlbum("integer-and-applesoft", album.Green)

	// Integer BASIC, from the Monitor with Control-B and Return
	o := start(t, basicsTape, nil)
	must(t, o.WaitForKeyboard(10))
	must(t, o.Key("Ctrl+B"))
	must(t, o.Key("Return"))
	must(t, o.WaitForKeyboard(10))
	must(t, o.TypeLines("PRINT 7/2", "PRINT 1/3", "PRINT 32767+1"))
	must(t, o.WaitForText("*** >32767 ERR", 10))
	must(t, o.WaitForKeyboard(10))
	must(t, pictures.Screenshot(o, "integer"))

	// Applesoft, loaded from its tape by Integer BASIC, and run
	must(t, o.TypeLines("LOAD"))
	must(t, o.WaitForKeyboard(300))
	must(t, o.TypeLines("RUN"))
	must(t, o.WaitForText("APPLESOFT ][ FLOATING POINT BASIC", 10))
	must(t, o.WaitForKeyboard(10))
	must(t, o.TypeLines("PRINT 7/2", "PRINT 1/3", "PRINT 32767+1", "PRINT SQR(2)"))
	must(t, o.WaitForText("1.41421356", 10))
	must(t, o.WaitForKeyboard(10))
	must(t, pictures.Screenshot(o, "applesoft-tape"))

	// An Apple ][+: Applesoft in its ROM, and Integer BASIC loaded by DOS
	// 3.3 into the Language Card, switched with INT and FP
	o = start(t, basicsPlus, nil)
	must(t, o.WaitForText("SYSTEM MASTER", 30))
	must(t, o.WaitForKeyboard(30))
	must(t, o.TypeLines("PRINT 1/3", "INT", "PRINT 1/3", "FP", "PRINT 1/3"))
	waitForLines(t, o, ".333333333", 2)
	must(t, o.WaitForKeyboard(10))
	must(t, pictures.Screenshot(o, "card-plus"))

	// The first Apple ][: Integer BASIC in its ROM, and Applesoft loaded
	// into the card. It starts in the Monitor: 6, Control-P and Return
	// start the disk in slot 6
	o = start(t, basicsInteger, nil)
	must(t, o.WaitForKeyboard(10))
	must(t, o.Type("6"))
	must(t, o.Key("Ctrl+P"))
	must(t, o.Key("Return"))
	must(t, o.WaitForText("SYSTEM MASTER", 30))
	must(t, o.WaitForKeyboard(30))
	must(t, o.TypeLines("PRINT 1/3", "FP", "PRINT 1/3", "INT", "PRINT 1/3"))
	waitForLines(t, o, "0", 2)
	must(t, o.WaitForKeyboard(10))
	must(t, pictures.Screenshot(o, "card-integer"))
}

// waitForLines runs the machine until a line of the screen is a text, the
// answer of BASIC to a PRINT, as many times as given
func waitForLines(t *testing.T, o *operator.Operator, text string, times int) {
	t.Helper()
	count := func() int {
		n := 0
		for _, line := range strings.Split(o.Text(), "\n") {
			if strings.TrimSpace(line) == text {
				n++
			}
		}
		return n
	}
	if !o.WaitUntil(10, func() bool { return count() >= times }) {
		t.Fatalf("%q was not on the screen %v times:\n%v", text, times, o.Text())
	}
}
