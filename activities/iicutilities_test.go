package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// iicUtilitiesMachine is the machine of the guide, as its command line of
// izapple2: an Apple IIc with its System Utilities in the drive, and a
// diskette never formatted in an external drive
const iicUtilitiesMachine = `izapple2 -model none -board 2c -cpu 65c02 -screen green \
    -rom "<internal>/Apple2c_ROM255.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s1 iicserial \
    -s2 iicserial \
    -s6 "diskii,disk1=disks/System Utilities for the IIc - 680-3066-A (Apple-1984).dsk,disk2=my-disk.woz"`

// iicUtilitiesBlank is the diskette of the external drive, a copy of the
// blank WOZ of disks/
const iicUtilitiesBlank = "my-disk.woz"

/*
iicUtilitiesScreenshots is the System Utilities of the Apple IIc, its disk of
ProDOS: the catalog of its own disk, a diskette never formatted formatted
for ProDOS in the external drive, two files copied to it, and its catalog
*/
func iicUtilitiesScreenshots(t *testing.T) {
	pictures := newAlbum("iic-utilities", album.Green)
	files := map[string]string{iicUtilitiesBlank: disk(t, "blank.woz")}
	o := start(t, iicUtilitiesMachine, files)

	// The title, and the menu
	must(t, o.WaitForText("Copyright Apple Computer, Inc., 1984", 30))
	o.RunSeconds(2)
	must(t, pictures.Screenshot(o, "title"))
	iicUtilitiesWait(t, o, "Exit System Utilities")
	must(t, pictures.Screenshot(o, "menu"))

	// 7, Identify and Catalog a Disk: the built-in drive, on the display
	must(t, o.TypeLines("7"))
	iicUtilitiesWait(t, o, "Where is your disk?")
	must(t, o.TypeLines("1"))
	iicUtilitiesWait(t, o, "Where do you want the list?")
	must(t, o.TypeLines("D"))
	iicUtilitiesWait(t, o, "Listing complete")
	must(t, pictures.Screenshot(o, "catalog"))
	must(t, o.Key("Escape"))
	iicUtilitiesWait(t, o, "Exit System Utilities")

	// 6, Format a Disk: the external drive, ProDOS, a name, and the diskette
	// in the drive
	must(t, o.TypeLines("6"))
	iicUtilitiesWait(t, o, "Where is your disk?")
	must(t, o.TypeLines("2"))
	iicUtilitiesWait(t, o, "Select the Operating System")
	must(t, pictures.Screenshot(o, "systems"))
	must(t, o.TypeLines("1"))
	iicUtilitiesWait(t, o, "Enter Name of New Volume")
	must(t, o.TypeLines("MYDISK"))
	iicUtilitiesWait(t, o, "Place the disk you wish to format")
	must(t, o.TypeLines(""))
	iicUtilitiesWait(t, o, "Format complete")
	must(t, pictures.Screenshot(o, "formatted"))
	must(t, o.Key("Escape"))
	iicUtilitiesWait(t, o, "Exit System Utilities")

	// 1, Copy Files: from the built-in drive to the external one, some of
	// them, STARTUP and SU marked with the right arrow
	must(t, o.TypeLines("1"))
	iicUtilitiesWait(t, o, "Where is your source disk?")
	must(t, o.TypeLines("1"))
	iicUtilitiesWait(t, o, "Where is your destination disk?")
	must(t, o.TypeLines("2"))
	iicUtilitiesWait(t, o, "Select Some or All")
	must(t, o.TypeLines("S"))
	iicUtilitiesWait(t, o, "To Accept Marked Files")
	o.RunSeconds(2)
	for range 2 {
		must(t, o.Key("Down"))
		o.RunSeconds(1)
		must(t, o.Key("Right"))
		o.RunSeconds(1)
	}
	must(t, pictures.Screenshot(o, "marked"))
	must(t, o.TypeLines(""))
	iicUtilitiesWait(t, o, "Copying Complete")
	must(t, pictures.Screenshot(o, "copied"))
	must(t, o.Key("Escape"))
	iicUtilitiesWait(t, o, "Exit System Utilities")

	// 7 again, the external drive: the new disk, and the two files on it
	must(t, o.TypeLines("7"))
	iicUtilitiesWait(t, o, "Where is your disk?")
	must(t, o.TypeLines("2"))
	iicUtilitiesWait(t, o, "Where do you want the list?")
	must(t, o.TypeLines("D"))
	iicUtilitiesWait(t, o, "Listing complete")
	must(t, pictures.Screenshot(o, "mydisk"))
	must(t, o.Key("Escape"))
	iicUtilitiesWait(t, o, "Exit System Utilities")

	// 9, Exit System Utilities
	must(t, o.TypeLines("9"))
	iicUtilitiesWait(t, o, "Are you sure you want to leave")
	must(t, o.TypeLines("Y"))
	must(t, o.WaitForText("]", 30))
	must(t, o.WaitForKeyboard(30))

	// BASIC, with ProDOS: the catalog of the new disk
	must(t, o.TypeLines("CAT /MYDISK"))
	must(t, o.WaitForText("BLOCKS FREE", 30))
	must(t, o.WaitForKeyboard(30))
	o.Run(30)
	must(t, pictures.Screenshot(o, "basic"))
}

// iicUtilitiesWait runs until a text is on the screen and the keyboard read
func iicUtilitiesWait(t *testing.T, o *operator.Operator, text string) {
	t.Helper()
	must(t, o.WaitForText(text, 300))
	must(t, o.WaitForKeyboard(30))
	o.Run(30)
}
