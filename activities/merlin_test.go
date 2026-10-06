package activities

import (
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// merlinMachine is the machine of the guide, as its command line of izapple2
const merlinMachine = `izapple2 -model _base -board 2plus -cpu 6502 \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s0 language \
    -s6 'diskii,disk1="disks/Merlin Macroassembler Side 1 (SDS, 1983).dsk"'`

// barsListing is the program of the page, in the assembly language of Merlin
const barsListing = "../guides/listings/bars.s"

/*
merlinScreenshots is a program in the assembly language of the 6502, written,
assembled and run with the Merlin assembler of 1983 on an Apple ][+: the
source typed into its editor, assembled into 120 bytes, run from the
Monitor, moving bars of colour, and its machine code disassembled.
*/
func merlinScreenshots(t *testing.T) {
	pictures := newAlbum("merlin", album.Green)
	o := start(t, merlinMachine, nil)

	// The title, and the menu after a key
	must(t, o.WaitForText("GLEN BREDON", 60))
	must(t, o.WaitForKeyboard(30))
	o.Run(30)
	must(t, pictures.Screenshot(o, "title"))
	must(t, o.Key("Return"))
	must(t, o.WaitForText("ENTER ED/ASM", 60))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "menu"))

	// The editor, and the source added a line at a time
	must(t, o.Type("E"))
	must(t, o.WaitForText(":", 10))
	must(t, o.WaitForKeyboard(10))
	must(t, o.TypeLines("A"))
	must(t, typeMerlin(o, barsListing))
	must(t, o.TypeLines(""))
	must(t, o.WaitForKeyboard(10))
	must(t, o.TypeLines("L1,25"))
	must(t, o.WaitForPrompt(30))
	must(t, pictures.Screenshot(o, "source"))

	// Assembled
	must(t, o.TypeLines("ASM"))
	must(t, o.WaitForText("UPDATE SOURCE", 10))
	must(t, o.Type("N"))
	must(t, o.WaitForText("NUMERICAL ORDER", 60))
	must(t, o.WaitForPrompt(30))
	must(t, pictures.Screenshot(o, "assembled"))

	// Run from the Monitor, at the address of the object code, until a key
	must(t, o.TypeLines("MON"))
	must(t, o.WaitForKeyboard(10))
	color := pictures.On(album.Color)
	bars := color.Record(o)
	must(t, o.TypeLines("8000G"))
	bars.Capture(10)
	bars.Run(2*60, 6)
	must(t, color.SaveRecording(bars, "bars", 100))
	must(t, o.Key("Space"))
	must(t, o.WaitForKeyboard(10))

	// The machine code, disassembled by the Monitor
	must(t, o.TypeLines("8000L"))
	must(t, o.WaitForPrompt(10))
	must(t, pictures.Screenshot(o, "disassembled"))
}

// merlinSpaces are the spaces between the fields of a line of Merlin's
// source, which are one space when typed: Merlin puts each field in its column
var merlinSpaces = regexp.MustCompile(` +`)

/*
typeMerlin types a source of the guides into the editor of Merlin, in add
mode: comments as they are, and the other lines with a single space between
their fields, starting with one when they have no label
*/
func typeMerlin(o interface{ TypeLines(...string) error }, path string) error {
	text, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	for _, line := range strings.Split(strings.TrimRight(string(text), "\n"), "\n") {
		if !strings.HasPrefix(line, "*") {
			typed := merlinSpaces.ReplaceAllString(strings.TrimSpace(line), " ")
			if strings.HasPrefix(line, " ") {
				typed = " " + typed
			}
			line = typed
		}
		if err := o.TypeLines(line); err != nil {
			return err
		}
	}
	return nil
}
