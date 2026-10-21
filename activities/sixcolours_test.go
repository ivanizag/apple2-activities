package activities

import (
	"strings"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// sixColoursRevision0 is the first Apple ][, of the board of Revision 0, with
// four colours in the high resolution graphics
const sixColoursRevision0 = `izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/341-000x_integer.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -mods four-colors`

// sixColoursRevision1 is an Apple ][ of a later board, with six colours in
// the high resolution graphics
const sixColoursRevision1 = `izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/341-000x_integer.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps`

// coloursListing is the program of the page, bands of the colours of the
// high resolution graphics
const coloursListing = "../guides/listings/colours.bas"

/*
sixColoursScreenshots is the colours of the high resolution graphics, four on
the first boards of the Apple ][ and six on the later ones, as Wozniak told
the readers of Byte to modify the first: the same program on both, and the
later one on the green monitor too, where the eighth bit moves the dots.
*/
func sixColoursScreenshots(t *testing.T) {
	pictures := newAlbum("six-colours", album.Color)

	// The colours of the first board
	o := start(t, sixColoursRevision0, nil)
	integerBasic(t, o)
	must(t, typeListing(o, coloursListing))
	coloursRun(t, o)
	must(t, pictures.Screenshot(o, "revision-0"))

	// And the ones of a later board, on both monitors
	o = start(t, sixColoursRevision1, nil)
	integerBasic(t, o)
	must(t, typeListing(o, coloursListing))
	coloursRun(t, o)
	must(t, pictures.Screenshot(o, "revision-1"))
	must(t, pictures.On(album.Green).Screenshot(o, "revision-1-green"))
}

// coloursRun runs the program of the bands until Integer BASIC prompts again
// under its RUN, in the text page behind the graphics: while a program runs,
// Integer BASIC looks at the keyboard for Control-C, so the keyboard can't
// tell. The text before the RUN may end in a prompt too, so the RUN and the
// prompt have to be the last lines.
func coloursRun(t *testing.T, o *operator.Operator) {
	t.Helper()
	must(t, o.TypeLines("RUN"))
	ended := func() bool {
		var lines []string
		for _, line := range strings.Split(o.Text(), "\n") {
			if line = strings.TrimSpace(line); line != "" {
				lines = append(lines, line)
			}
		}
		n := len(lines)
		return n >= 2 && lines[n-2] == ">RUN" && lines[n-1] == ">"
	}
	if !o.WaitUntil(120, ended) {
		t.Fatalf("the program did not end:\n%v", o.Text())
	}
	o.Run(30)
}
