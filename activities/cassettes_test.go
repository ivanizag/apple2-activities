package activities

import (
	"strings"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// cassettesBreakout is the first Apple ][ of the guide, with the tape of
// Breakout in its cassette recorder
const cassettesBreakout = `izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/341-000x_integer.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -mods four-colors \
    -tape disks/k7_apple_002000101_breakout.wav`

// cassettesColor is the same machine with the tape of Color Graphics
const cassettesColor = `izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/341-000x_integer.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -mods four-colors \
    -tape disks/k7_apple_002000101_colorgraphics.wav`

// cassettesHires is the same machine with the tape of the High-Resolution
// Graphics demonstrations
const cassettesHires = `izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/341-000x_integer.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -mods four-colors \
    -tape disks/k7_apple_600201600_highresolutiongraphics.wav`

// cassettesRevision1 is an Apple ][ of a later board, with six colours in
// the high resolution graphics
const cassettesRevision1 = `izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/341-000x_integer.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps`

// coloursListing is the program of the page, bands of the colours of the
// high resolution graphics
const coloursListing = "../guides/listings/colours.bas"

/*
cassettesScreenshots is the first Apple ][ and the programs that came with it
on cassette: Breakout loaded and played with the paddle, the demonstrations of
the colour graphics, and the ones of the high resolution graphics, their
machine code read with the Monitor and their BASIC with LOAD. Then the
colours of the high resolution graphics, four on the first boards and six on
the later ones, as Wozniak told the readers of Byte to modify the first.
*/
func cassettesScreenshots(t *testing.T) {
	pictures := newAlbum("apple-ii-cassettes", album.Color)

	// Breakout, loaded from its tape and played with paddle 0
	o := start(t, cassettesBreakout, nil)
	integerBasic(t, o)
	must(t, o.TypeLines("LOAD"))
	must(t, o.WaitForKeyboard(300))
	must(t, o.TypeLines("RUN"))
	must(t, o.WaitForText("WHAT'S YOUR NAME?", 10))
	must(t, o.WaitForKeyboard(10))
	must(t, pictures.Screenshot(o, "breakout-name"))
	must(t, o.TypeLines("ADA"))
	must(t, o.WaitForText("STANDARD COLORS ADA?", 10))
	must(t, o.TypeLines("Y"))
	must(t, o.WaitForText("BALLS LEFT", 10))
	game := pictures.Record(o)
	game.Capture(10)
	for range 300 {
		breakoutPaddle(o)
		game.Run(6, 6)
	}
	must(t, pictures.SaveRecording(game, "breakout", 100))

	// Color Graphics, and its kaleidoscope
	o = start(t, cassettesColor, nil)
	integerBasic(t, o)
	must(t, o.TypeLines("LOAD"))
	must(t, o.WaitForKeyboard(300))
	must(t, o.TypeLines("RUN"))
	must(t, o.WaitForText("WHICH WOULD YOU LIKE?", 10))
	must(t, o.WaitForKeyboard(10))
	must(t, pictures.Screenshot(o, "color-menu"))
	must(t, o.TypeLines("3"))
	kaleidoscope := pictures.Record(o)
	kaleidoscope.Capture(10)
	kaleidoscope.Run(15*60, 6)
	must(t, pictures.SaveRecording(kaleidoscope, "kaleidoscope", 100))

	// The High-Resolution Graphics: its machine code first, read by the
	// Monitor into $C00 to $FFF, then its program in Integer BASIC
	o = start(t, cassettesHires, nil)
	o.RunSeconds(1)
	must(t, o.TypeLines("C00.FFFR"))
	must(t, o.WaitForKeyboard(300))
	integerBasic(t, o)
	must(t, o.TypeLines("LOAD"))
	must(t, o.WaitForKeyboard(300))
	must(t, o.TypeLines("RUN"))
	must(t, o.WaitForText("WHICH DEMO # DO YOU WANT ?", 10))
	must(t, o.WaitForKeyboard(10))
	must(t, pictures.Screenshot(o, "hires-menu"))
	must(t, o.TypeLines("6"))
	donut := pictures.Record(o)
	donut.Capture(10)
	donut.Run(60*60, 6)
	must(t, pictures.SaveRecording(donut, "donut", 300))

	// The colours of the first board: Reset, a new Integer BASIC, and the
	// program of the bands
	o.Reset()
	o.RunSeconds(1)
	integerBasic(t, o)
	must(t, typeListing(o, coloursListing))
	coloursRun(t, o)
	must(t, pictures.Screenshot(o, "revision-0"))

	// And the ones of a later board
	o = start(t, cassettesRevision1, nil)
	integerBasic(t, o)
	must(t, typeListing(o, coloursListing))
	coloursRun(t, o)
	must(t, pictures.Screenshot(o, "revision-1"))
}

// integerBasic goes from the Monitor of the Apple ][ to Integer BASIC, with
// Control-B and Return
func integerBasic(t *testing.T, o *operator.Operator) {
	t.Helper()
	must(t, o.WaitForKeyboard(10))
	must(t, o.Key("Ctrl+B"))
	must(t, o.Key("Return"))
	must(t, o.WaitForKeyboard(10))
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

// breakoutPaddle turns paddle 0 so that the bat of Breakout, six blocks
// down the left edge from (PDL(0)-20)/6, is in front of the ball, white
func breakoutPaddle(o *operator.Operator) {
	_, y, ok := loResFind(o, 15)
	if !ok {
		return
	}
	pdl := (y-2)*6 + 23
	o.TurnPaddle(0, uint8(min(max(pdl, 0), 255)))
}
