package activities

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// printingMachine is the machine of the guide, as its command line of izapple2
const printingMachine = `izapple2 -model _base -board 2plus -cpu 6502 \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s1 parallel,file=printer.out`

// calendarListing is the program of the page, a calendar of a year
const calendarListing = "../guides/listings/calendar.bas"

/*
printingScreenshots is an Apple ][+ with a parallel printer card: the
calendar of the page typed, its listing printed with PR#1, and the program
run, printing a year three months across, wider than the screen. What the
card prints is kept as text, as the pages show it.
*/
func printingScreenshots(t *testing.T) {
	pictures := newAlbum("printing", album.Green)
	printer := filepath.Join(t.TempDir(), "printer.out")
	o := start(t, printingMachine, map[string]string{"printer.out": printer})
	printed = 0
	must(t, o.WaitForKeyboard(5))

	// The program typed
	must(t, typeListing(o, calendarListing))
	must(t, o.WaitForPrompt(10))
	must(t, pictures.Screenshot(o, "typed"))

	// Its listing, printed: the output sent to slot 1 and back
	must(t, o.TypeLines("PR#1", "LIST 10,90", "PR#0"))
	must(t, o.WaitForPrompt(30))
	savePrintout(t, printer, "listing")

	// The calendar of 1977, printed, and the screen while it does
	must(t, o.TypeLines("RUN"))
	must(t, o.WaitForText("YEAR?", 10))
	must(t, o.TypeLines("1977"))
	if !o.WaitUntil(300, func() bool { return bytes.Contains(printable(printer), []byte("DECEMBER")) }) {
		t.Fatal("the calendar did not get to December")
	}
	must(t, o.WaitForPrompt(60))
	must(t, pictures.Screenshot(o, "printing"))
	savePrintout(t, printer, "calendar")
}

// printed is how much of the printer's file the printouts before have taken
var printed int

/*
savePrintout keeps what was printed since the printout before as a text file
of the guide: the Apple sends its characters with the top bit set and ends
its lines with a carriage return and a line feed, and the file has plain
ASCII, lines ended with a line feed, no spaces at their ends and no empty
lines around
*/
func savePrintout(t *testing.T, printer string, name string) {
	t.Helper()
	text := printable(printer)
	part := text[printed:]
	printed = len(text)
	lines := bytes.Split(part, []byte("\n"))
	for i := range lines {
		lines[i] = bytes.TrimRight(lines[i], " ")
	}
	part = append(bytes.Trim(bytes.Join(lines, []byte("\n")), "\n"), '\n')
	must(t, os.WriteFile(filepath.Join(guideImages, "printing", name+".txt"), part, 0o644))
}

// printable is what the printer got, as plain text
func printable(printer string) []byte {
	raw, err := os.ReadFile(printer)
	if err != nil {
		return nil
	}
	text := make([]byte, 0, len(raw))
	for _, b := range raw {
		if b&0x7f != '\r' {
			text = append(text, b&0x7f)
		}
	}
	return text
}
