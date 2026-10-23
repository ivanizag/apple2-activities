package activities

import (
	"os"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// cpm65Machine is the machine of the guide, as its command line of izapple2
const cpm65Machine = `izapple2 -model none -board 2e -cpu 65c02 -screen green \
    -rom "<internal>/Apple2e_Enhanced.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s6 diskii,disk1=disks/cpm65-apple2e.po`

// The programs of the guide, one in the assembly language of CP/M-65 and one
// in its Atari BASIC
const (
	alphabetListing = "../guides/listings/alphabet.asm"
	powersListing   = "../guides/listings/powers.bas"
)

/*
cpm65Screenshots is CP/M-65 on an enhanced Apple //e: started, its disk
listed, the program of its demonstration assembled and run, room made on the
disk, a program of the page written in the editor QE, assembled and run, and
Atari BASIC with its example and a program of the page, saved.
*/
func cpm65Screenshots(t *testing.T) {
	pictures := newAlbum("cpm65", album.Green)
	o := start(t, cpm65Machine, nil)

	// Started, at the prompt of drive A
	must(t, o.WaitForText("A>", 60))
	must(t, o.WaitForPrompt(10))
	must(t, pictures.Screenshot(o, "started"))

	// The disk, its files and the space on it, and the processor
	cpm65Command(t, o, "DIR")
	cpm65Command(t, o, "STAT")
	cpm65Command(t, o, "CPUINFO")
	must(t, pictures.Screenshot(o, "dir"))

	// The demonstration: a source shown, assembled and run, by SUBMIT
	cpm65Command(t, o, "CLS")
	cpm65Command(t, o, "SUBMIT DEMO")
	must(t, o.WaitForText("Hello, world!", 60))
	must(t, o.WaitForPrompt(30))
	must(t, pictures.Screenshot(o, "demo"))

	// Room on the disk, the source of BEDIT erased
	cpm65Command(t, o, "CLS")
	cpm65Command(t, o, "ERA BEDIT.ASM")
	cpm65Command(t, o, "STAT")
	must(t, pictures.Screenshot(o, "room"))

	// The program of the page written in QE, typed into its insert mode,
	// and saved with :wq
	must(t, o.TypeLines("QE ALPHABET.ASM"))
	must(t, o.WaitForText("~", 30))
	must(t, o.WaitForKeyboard(10))
	editor := pictures.Record(o)
	editor.Capture(50)
	must(t, editor.Type("i"))
	source, err := os.ReadFile(alphabetListing)
	must(t, err)
	must(t, editor.Type(string(source)))
	must(t, o.Key("Escape"))
	editor.Run(60, 6)
	must(t, pictures.SaveRecording(editor, "qe", 300))
	must(t, o.TypeLines(":wq"))
	must(t, o.WaitForText("Goodbye!", 30))
	must(t, o.WaitForPrompt(30))

	// Assembled and run
	cpm65Command(t, o, "CLS")
	cpm65Command(t, o, "ASM ALPHABET.ASM ALPHABET.COM")
	cpm65Command(t, o, "ALPHABET")
	must(t, pictures.Screenshot(o, "alphabet"))
	cpm65Command(t, o, "DUMP ALPHABET.COM")
	must(t, pictures.Screenshot(o, "dump"))

	// Atari BASIC, its example entered from the text of the disk and run
	cpm65Command(t, o, "CLS")
	must(t, o.TypeLines("ATBASIC"))
	must(t, o.WaitForText("Ready", 30))
	must(t, o.WaitForPrompt(10))
	cpm65Command(t, o, `ENTER "ATBASIC.TXT"`)
	cpm65Command(t, o, "LIST")
	cpm65Command(t, o, "RUN")
	must(t, pictures.Screenshot(o, "atbasic"))

	// The program of the page, run, saved and loaded back
	cpm65Command(t, o, "NEW")
	must(t, typeListing(o, powersListing))
	cpm65Command(t, o, "RUN")
	must(t, pictures.Screenshot(o, "powers"))
	cpm65Command(t, o, `SAVE "POWERS.BAS"`)
	cpm65Command(t, o, "NEW")
	cpm65Command(t, o, `LOAD "POWERS.BAS"`)
	cpm65Command(t, o, "LIST")
	must(t, pictures.Screenshot(o, "loaded"))
	must(t, o.TypeLines("BYE"))
	must(t, o.WaitForText("A>", 30))
	must(t, o.WaitForPrompt(10))
	cpm65Command(t, o, "DIR")
	must(t, pictures.Screenshot(o, "saved"))
}

// cpm65Command types a command, of CP/M-65 or of a program on it, and waits
// for the prompt after it
func cpm65Command(t *testing.T, o *operator.Operator, line string) {
	t.Helper()
	must(t, o.TypeLines(line))
	must(t, o.WaitForPrompt(60))
}
