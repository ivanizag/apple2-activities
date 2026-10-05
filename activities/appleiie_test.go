package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

/*
appleIIeScreenshots is an enhanced Apple //e with no disk drive: lower case,
80 columns, MouseText, and the self test of its ROM.

	izapple2 -model 2enh -s6 empty
*/
func appleIIeScreenshots(t *testing.T) {
	pictures := newAlbum("apple-iie", album.Green)
	o := start(t, "2enh", map[string]string{"s6": "empty"})

	// Switched on, in Applesoft
	must(t, o.WaitForKeyboard(5))
	o.Run(30)
	must(t, pictures.Screenshot(o, "switched-on"))

	// Lower case
	must(t, o.TypeLines(
		`print "Hello, //e"`,
		`10 for i = 1 to 3: print i; " lower case": next`,
		"list",
		"run",
	))
	o.Run(30)
	must(t, pictures.Screenshot(o, "lower-case"))

	// 80 columns, and MouseText
	must(t, o.TypeLines(
		"PR#3",
		"HOME",
		"LIST",
		`PRINT CHR$(27);: INVERSE: PRINT "@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_": NORMAL: PRINT CHR$(24);`,
	))
	o.Run(30)
	must(t, pictures.Screenshot(o, "eighty-columns"))

	// The self test, with both Apple keys held through a reset
	o.HoldButton(0)
	o.HoldButton(1)
	o.Reset()
	o.RunSeconds(2)
	o.ReleaseButton(0)
	o.ReleaseButton(1)
	o.RunSeconds(10)
	must(t, pictures.On(album.Color).Screenshot(o, "self-test"))
	must(t, o.WaitForShownText("System OK", 120))
	o.Run(30)
	must(t, pictures.Screenshot(o, "system-ok"))
}
