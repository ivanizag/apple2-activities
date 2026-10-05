package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// prodosMachine is the machine of the guide, as its command line of izapple2
const prodosMachine = `izapple2 -model _base -board 2e -cpu 65c02 \
    -rom "<internal>/Apple2e_Enhanced.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s6 diskii,disk1=disks/prodos-2.4.3.po`

/*
prodosScreenshots is ProDOS 2.4.3 on an enhanced Apple //e: the program
selector it starts with, BASIC.SYSTEM chosen in it, and the disk catalogued
from BASIC in 80 columns.
*/
func prodosScreenshots(t *testing.T) {
	pictures := newAlbum("prodos", album.Green)
	o := start(t, prodosMachine, nil)

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
