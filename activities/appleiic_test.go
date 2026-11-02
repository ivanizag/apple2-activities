package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// appleIIcMachine is the machine of the guide, as its command line of izapple2
const appleIIcMachine = `izapple2 -model none -board 2c -cpu 65c02 -screen green \
    -rom "<internal>/Apple2c_ROM255.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s1 iicserial \
    -s2 iicserial \
    -s6 iicdisk`

/*
appleIIcScreenshots is the Apple IIc of 1984 switched on: with no disk, it
waits for one; Applesoft after a reset, in small letters, and its 80 columns
*/
func appleIIcScreenshots(t *testing.T) {
	pictures := newAlbum("apple-iic", album.Green)
	o := start(t, appleIIcMachine, nil)

	// Switched on with no disk: it waits for one in the drive
	must(t, o.WaitForText("Apple //c", 5))
	o.RunSeconds(3)
	must(t, pictures.Screenshot(o, "switched-on"))

	// Control-Reset, and Applesoft
	o.Reset()
	must(t, o.WaitForKeyboard(5))
	must(t, o.TypeLines(`print "Hello, //c"`, "print 6*7"))
	must(t, o.WaitForKeyboard(5))
	o.Run(30)
	must(t, pictures.Screenshot(o, "applesoft"))

	// 80 columns, from the firmware in slot 3
	must(t, o.TypeLines(
		"PR#3",
		`10 FOR I = 1 TO 8: PRINT I, I * I, I * I * I, SQR(I): NEXT`,
		"LIST",
		"RUN",
	))
	must(t, o.WaitForText("2.82842713", 10))
	must(t, o.WaitForKeyboard(5))
	o.Run(30)
	must(t, pictures.Screenshot(o, "eighty-columns"))
}
