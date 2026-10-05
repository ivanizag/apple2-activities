package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// cpmMachine is the machine of the guide, as its command line of izapple2
const cpmMachine = `izapple2 -model _base -board 2plus -cpu 6502 \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s0 language \
    -s4 z80softcard \
    -s6 diskii,disk1=disks/cpm-2.20b.po`

/*
cpmScreenshots is CP/M on an Apple ][+ with the Microsoft Z80 SoftCard: the
system started, its disk listed, an assembler source typed on the screen, and
a program in Microsoft BASIC-80.
*/
func cpmScreenshots(t *testing.T) {
	pictures := newAlbum("cpm", album.Green)
	o := start(t, cpmMachine, nil)

	// Started, at the prompt of drive A
	must(t, o.WaitForText("A>", 120))
	must(t, o.WaitForKeyboard(10))
	must(t, pictures.Screenshot(o, "started"))

	// The disk, its files and the space on it
	must(t, o.TypeLines("DIR"))
	must(t, o.WaitForPrompt(30))
	must(t, o.TypeLines("STAT"))
	must(t, o.WaitForPrompt(30))
	o.Run(30)
	must(t, pictures.Screenshot(o, "dir"))

	// A source file on the screen, stopped with Control-C
	must(t, o.TypeLines("TYPE DUMP.ASM"))
	o.Run(6 * 60)
	must(t, pictures.Screenshot(o, "type"))
	must(t, o.Key("Ctrl+C"))
	must(t, o.WaitForPrompt(10))

	// Microsoft BASIC-80, a program and its run
	must(t, o.TypeLines("MBASIC"))
	must(t, o.WaitForText("OK", 120))
	must(t, o.WaitForKeyboard(10))
	basic := pictures.Record(o)
	basic.Capture(50)
	must(t, basic.TypeLines(
		"10 FOR I = 1 TO 10",
		`20 PRINT USING "###  #####.##"; I, SQR(I)*1000`,
		"30 NEXT",
		"RUN",
	))
	basic.Run(120, 6)
	must(t, pictures.SaveRecording(basic, "mbasic", 300))

	// Back to CP/M
	must(t, o.TypeLines("SYSTEM"))
	must(t, o.WaitForText("\nA>", 60))
}
