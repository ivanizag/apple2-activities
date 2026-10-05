package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

/*
prodosScreenshots is ProDOS 2.4.3 on an enhanced Apple //e: the program
selector it starts with, BASIC.SYSTEM chosen in it, and the disk catalogued
from BASIC in 80 columns.

	izapple2 -model prodos
*/
func prodosScreenshots(t *testing.T) {
	pictures := newAlbum("prodos", album.Green)
	o := start(t, "prodos", nil)

	// The program selector, Bitsy Bye
	must(t, o.WaitForText("BITSY", 60))
	must(t, o.WaitForPrompt(10))
	must(t, pictures.Screenshot(o, "bitsy-bye"))

	// BASIC.SYSTEM chosen with the arrows
	for range 3 {
		must(t, o.Key("Down"))
	}
	o.Run(30)
	must(t, pictures.Screenshot(o, "selected"))
	must(t, o.Key("Return"))
	must(t, o.WaitForText("PRODOS BASIC", 120))
	must(t, o.WaitForKeyboard(30))
	o.Run(30)
	must(t, pictures.Screenshot(o, "basic"))

	// The disk, catalogued in 80 columns
	must(t, o.TypeLines("PR#3", "CATALOG"))
	must(t, o.WaitForText("BLOCKS FREE", 60))
	must(t, o.WaitForKeyboard(30))
	o.Run(30)
	must(t, pictures.Screenshot(o, "catalog"))
}
