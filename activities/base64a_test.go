package activities

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// base64aMachine is the machine of the guide, as its command line of
// izapple2, with a printer
const base64aMachine = `izapple2 -model none -board base64a -cpu 6502 -screen green \
    -rom "<custom>" \
    -charrom "<internal>/BASE64A_ROM7_CharGen.BIN" \
    -s0 language \
    -s1 parallel,file=printer.out`

// base64aDOS is the same machine with a disk drive and the DOS 3.3 System
// Master, and no printer
const base64aDOS = `izapple2 -model none -board base64a -cpu 6502 -screen green \
    -rom "<custom>" \
    -charrom "<internal>/BASE64A_ROM7_CharGen.BIN" \
    -s0 language \
    -s6 'diskii,disk1=disks/DOS 3.3 System Master - 680-0210-A (1982).dsk'`

/*
base64aScreenshots is the Base 64A, an Apple ][+ made in Taiwan: its name
when it starts, the lower case it shows, Mini-Writer, the word processor in
its ROM, with a letter written, a name replaced in it and the letter printed,
and DOS 3.3 started from Apple's System Master.
*/
func base64aScreenshots(t *testing.T) {
	pictures := newAlbum("base64a", album.Green)
	printer := filepath.Join(t.TempDir(), "printer.out")
	o := start(t, base64aMachine, map[string]string{"printer.out": printer})

	// Switched on: its name, and Applesoft
	must(t, o.WaitForText("BASE 64A", 10))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "switched-on"))

	// Small letters, on the screen too
	must(t, o.TypeLines(`PRINT "Hello from a Base 64A"`))
	must(t, o.WaitForText("\nHello from a Base 64A\n", 10))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "lower-case"))

	// Mini-Writer, from the ROM
	must(t, o.TypeLines("WRITER"))
	must(t, o.WaitForText("SELECT:", 10))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "writer"))

	// A new file, edited
	must(t, o.TypeLines("N"))
	must(t, o.WaitForText("ERASE FILE IN MEMORY (Y/N)?", 10))
	must(t, o.Type("Y"))
	must(t, o.WaitForKeyboard(10))
	must(t, o.TypeLines("E"))
	if !o.WaitUntil(10, func() bool { return !o.HasText("SELECT") }) {
		t.Fatalf("the editor did not open:\n%v", o.Text())
	}
	must(t, o.WaitForKeyboard(10))
	typed := pictures.Record(o)
	typed.Capture(100)
	for _, line := range base64aLetter {
		must(t, typed.Type(line+"\n"))
	}
	typed.Run(90, 6)
	must(t, pictures.SaveRecording(typed, "letter", 300))

	// Ada replaced by Grace: to the top with Control-B, and Control-S
	must(t, o.Key("Ctrl+B"))
	must(t, o.Key("Ctrl+S"))
	must(t, o.WaitForText("SEARCH & REPLACE, ENTER", 10))
	must(t, o.TypeLines("/Ada/Grace/"))
	must(t, o.WaitForText("REPLACE (A) AUTOMATIC, OR (M) MANUAL", 10))
	must(t, o.Type("A"))
	must(t, o.WaitForKeyboard(10))
	must(t, o.Key("Return"))
	must(t, o.WaitForText("Dear Grace,", 10))
	must(t, o.WaitForKeyboard(10))
	must(t, pictures.Screenshot(o, "replaced"))

	// Back to the menu, Escape and Control-Q, and the letter printed, on
	// paper that goes on, not a sheet at a time
	must(t, o.Key("Escape"))
	must(t, o.Key("Ctrl+Q"))
	must(t, o.WaitForText("EDITOR MENU", 10))
	must(t, o.WaitForKeyboard(10))
	must(t, o.TypeLines("P"))
	must(t, o.WaitForText("PRINT MENU", 10))
	must(t, o.WaitForKeyboard(10))
	must(t, o.TypeLines("P"))
	must(t, o.WaitForText("PRINTER CONSTANTS ARE SET", 10))
	must(t, o.WaitForKeyboard(10))
	must(t, pictures.Screenshot(o, "printer"))
	must(t, o.TypeLines("G"))
	must(t, o.WaitForText("ENTER NEW VALUE", 10))
	must(t, o.TypeLines("0", ""))
	must(t, o.WaitForText("ENTER PAGE HEADING", 10))
	must(t, o.TypeLines(""))
	must(t, o.WaitForText("PRESS RETURN TO START PRINTING", 10))
	must(t, o.TypeLines(""))
	if !o.WaitUntil(30, func() bool { return bytes.Contains(printable(printer), []byte("Alan")) }) {
		t.Fatalf("the letter was not printed: %q", printable(printer))
	}
	must(t, o.WaitForText("PRINT MENU", 30))
	saveBase64aPrintout(t, printer)

	// Apple's DOS 3.3, started from its System Master
	o = start(t, base64aDOS, nil)
	must(t, o.WaitForText("SYSTEM MASTER", 30))
	must(t, o.WaitForKeyboard(30))
	must(t, o.TypeLines("CATALOG"))
	must(t, o.WaitForText("BOOT13", 30))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "dos"))
}

// base64aLetter is the letter of the page, a line at a time
var base64aLetter = []string{
	"Dear Ada,",
	"",
	"I write to you with Mini-Writer, the",
	"word processor in the ROM of my Base",
	"64A. I typed WRITER, and there it was.",
	"",
	"Yours,",
	"Alan",
}

// saveBase64aPrintout keeps what was printed as a text file of the guide,
// plain ASCII with no spaces at the ends of the lines and no empty lines
// around
func saveBase64aPrintout(t *testing.T, printer string) {
	t.Helper()
	lines := bytes.Split(printable(printer), []byte("\n"))
	for i := range lines {
		lines[i] = bytes.TrimRight(lines[i], " ")
	}
	text := append(bytes.Trim(bytes.Join(lines, []byte("\n")), "\n"), '\n')
	must(t, os.WriteFile(filepath.Join(guideImages, "base64a", "letter.txt"), text, 0o644))
}
