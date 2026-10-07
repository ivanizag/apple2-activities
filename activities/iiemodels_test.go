package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// iieOriginal is the Apple //e of 1983, as its command line of izapple2
const iieOriginal = `izapple2 -model none -board 2e -cpu 6502 -screen green \
    -rom "<internal>/Apple2e.rom" \
    -charrom "<internal>/Apple IIe Video Unenhanced.bin" \
    -s0 language`

// iieEnhanced is the enhanced Apple //e of 1985
const iieEnhanced = `izapple2 -model none -board 2e -cpu 65c02 -screen green \
    -rom "<internal>/Apple2e_Enhanced.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language`

// mouseTextLine prints the 32 characters from @ to _ in inverse, with
// MouseText turned on in 80 columns, Control-[, and off again, Control-X
const mouseTextLine = `PRINT CHR$(27): INVERSE: PRINT "@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_": NORMAL: PRINT CHR$(24)`

/*
iieModelsScreenshots is the Apple //e of 1983 and the enhanced one of 1985,
the same things done on both: switched on, a command of Applesoft in lower
case, the inverse capitals in 80 columns that are MouseText on the enhanced
one, and the Mini-Assembler of the Monitor, which only the enhanced one has.
*/
func iieModelsScreenshots(t *testing.T) {
	pictures := newAlbum("apple-iie-models", album.Green)
	for _, model := range []struct {
		name    string
		command string
	}{
		{"original", iieOriginal},
		{"enhanced", iieEnhanced},
	} {
		o := start(t, model.command, nil)
		must(t, o.WaitForText("Apple ", 10))
		must(t, o.WaitForKeyboard(10))
		o.Run(30)

		// A command of Applesoft in small letters
		must(t, o.TypeLines(`print "hello"`))
		must(t, o.WaitForKeyboard(10))
		o.Run(30)
		must(t, pictures.Screenshot(o, model.name+"-lower"))

		// The inverse capitals, in 80 columns, with MouseText turned on
		must(t, o.TypeLines("PR#3", mouseTextLine))
		must(t, o.WaitForKeyboard(10))
		o.Run(30)
		must(t, pictures.Screenshot(o, model.name+"-mousetext"))

		// The Monitor, and ! for the Mini-Assembler
		must(t, o.TypeLines("CALL -151", "!"))
		must(t, o.WaitForKeyboard(10))
		if model.name == "enhanced" {
			iieAssemble(t, o)
		}
		o.Run(30)
		must(t, pictures.Screenshot(o, model.name+"-monitor"))
	}
}

// iieAssemble writes a program of three instructions with the Mini-Assembler
// of the enhanced //e, leaves it with an empty line, and runs it: it prints
// an A
func iieAssemble(t *testing.T, o *operator.Operator) {
	t.Helper()
	must(t, o.TypeLines("300:LDA #$C1", " JSR $FDED", " RTS", "", "300G"))
	must(t, o.WaitForKeyboard(10))
}
