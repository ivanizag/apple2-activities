package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// cpmMachine is the machine of the guide, as its command line of izapple2
const cpmMachine = `izapple2 -model none -board 2plus -cpu 6502 -screen green \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s0 language \
    -s4 z80softcard \
    -s6 diskii,disk1=disks/CPM1.PO`

/*
cpmScreenshots is CP/M on an Apple ][+ with the Microsoft Z80 SoftCard: the
system started, its disk listed, an assembler source typed on the screen, a
program in Microsoft BASIC-80 run and saved as text, and the graphics of the
Apple II drawn from GBASIC.
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

	// Saved as text, and back to CP/M to see it among the files
	must(t, o.TypeLines(`SAVE "ROOTS",A`))
	must(t, o.WaitForKeyboard(30))
	o.Run(60)
	must(t, o.TypeLines("SYSTEM"))
	must(t, o.WaitForText("\nA>", 60))
	must(t, o.WaitForKeyboard(10))
	must(t, o.TypeLines("DIR"))
	must(t, o.WaitForPrompt(30))
	must(t, o.TypeLines("TYPE ROOTS.BAS"))
	must(t, o.WaitForText("30 NEXT", 30))
	must(t, o.WaitForPrompt(30))
	must(t, pictures.Screenshot(o, "saved"))

	// GBASIC, the BASIC with the graphics of the Apple II, drawing
	must(t, o.TypeLines("GBASIC"))
	must(t, o.WaitForText("BYTES FREE", 120))
	must(t, o.WaitForKeyboard(10))
	must(t, o.TypeLines(
		"10 HGR : HCOLOR = 3",
		"20 FOR I = 0 TO 159 STEP 6",
		"30 HPLOT 0,I TO I*1.75,159",
		"40 HPLOT 279,159-I TO 279-I*1.75,0",
		"50 NEXT",
	))
	must(t, o.WaitForKeyboard(10))
	color := pictures.On(album.Color)
	drawing := color.Record(o)
	drawing.Capture(50)
	must(t, drawing.TypeLines("RUN"))
	for !o.HasText("RUN\nOK") {
		if o.Frames() > 60*60*10 {
			t.Fatal("the drawing did not end")
		}
		drawing.Run(6, 6)
	}
	drawing.Run(60, 6)
	must(t, color.SaveRecording(drawing, "gbasic", 300))

	// Back to CP/M
	must(t, o.TypeLines("SYSTEM"))
	must(t, o.WaitForText("\nA>", 60))
}
