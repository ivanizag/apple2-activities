package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// prodosMachine is the machine of the guide, as its command line of izapple2
const prodosMachine = `izapple2 -model none -board 2e -cpu 65c02 -screen green \
    -rom "<internal>/Apple2e_Enhanced.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s6 diskii,disk1=disks/ProDOS_2_4_3.po`

/*
prodosScreenshots is ProDOS 2.4.3 on an enhanced Apple //e: the program
selector it starts with, BASIC.SYSTEM chosen in it, the disk catalogued from
BASIC in 80 columns, a folder made on the RAM disk with a program that writes
a text file and reads it back, and Copy II Plus, its menu and the map of the
disk, file by file.
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

	// The RAM disk, a folder on it, and a program that writes a text file
	// there and reads it back
	must(t, o.TypeLines("HOME", "CATALOG /RAM"))
	must(t, o.WaitForText("TOTAL BLOCKS", 30))
	must(t, o.WaitForKeyboard(30))
	o.Run(30)
	must(t, pictures.Screenshot(o, "ram-disk"))
	must(t, o.TypeLines("HOME", "CREATE /RAM/NOTES", "PREFIX /RAM/NOTES"))
	must(t, o.TypeLines(noteProgram...))
	must(t, o.TypeLines("SAVE WRITE.NOTE", "RUN"))
	must(t, o.WaitForText("THE NOTE SAYS", 30))
	must(t, o.TypeLines("CATALOG"))
	must(t, o.WaitForText("TOTAL BLOCKS:  127", 30))
	must(t, o.WaitForKeyboard(30))
	o.Run(30)
	must(t, pictures.Screenshot(o, "note"))

	// Back to Bitsy Bye, which shows the RAM disk now, and to the diskette
	must(t, o.TypeLines("BYE"))
	must(t, o.WaitForShownText("S3,D2:/RAM", 30))
	must(t, o.WaitForPrompt(10))
	must(t, pictures.Screenshot(o, "bye"))
	must(t, o.Key("Tab"))
	must(t, o.WaitForShownText("S6,D1:/PRODOS.2.4.3", 30))
	must(t, o.WaitForPrompt(10))

	// Copy II Plus, with no date
	for range 4 {
		must(t, o.Key("Down"))
	}
	must(t, o.Key("Return"))
	must(t, o.WaitForShownText("ENTER DATE", 60))
	must(t, o.WaitForKeyboard(10))
	must(t, o.Key("Escape"))
	must(t, o.WaitForShownText("SELECT FUNCTION", 30))
	must(t, o.WaitForPrompt(10))
	must(t, pictures.Screenshot(o, "copy-ii-plus"))

	// The map of the diskette in drive 1, and then of each file
	for range 9 {
		must(t, o.Key("Down"))
	}
	must(t, o.Key("Return"))
	must(t, o.WaitForShownText("SELECT DEVICE", 10))
	must(t, o.WaitForPrompt(10))
	must(t, o.Key("Return"))
	must(t, o.WaitForShownText("[RETURN]-CONTINUE", 30))
	must(t, o.WaitForPrompt(30))
	must(t, pictures.Screenshot(o, "disk-map"))
	files := pictures.Record(o)
	files.Capture(60)
	must(t, o.Key("Return"))
	must(t, o.WaitForShownText("MAP OTHER FILES", 30))
	for range 6 {
		files.Run(60, 6)
		must(t, o.Key("Right"))
	}
	files.Run(90, 6)
	must(t, pictures.SaveRecording(files, "file-map", 300))
}

// noteProgram writes a line to a text file, NOTE, and reads it back, with the
// commands of ProDOS printed after Control-D, CHR$(4)
var noteProgram = []string{
	`10 D$ = CHR$ (4)`,
	`20 PRINT D$;"OPEN NOTE"`,
	`30 PRINT D$;"WRITE NOTE"`,
	`40 PRINT "MADE ON AN APPLE //E"`,
	`50 PRINT D$;"CLOSE NOTE"`,
	`60 PRINT D$;"OPEN NOTE"`,
	`70 PRINT D$;"READ NOTE"`,
	`80 INPUT A$`,
	`90 PRINT D$;"CLOSE NOTE"`,
	`100 PRINT "THE NOTE SAYS: ";A$`,
}
