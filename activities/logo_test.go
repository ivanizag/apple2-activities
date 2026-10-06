package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// logoMachine is the machine of the guide, as its command line of izapple2
const logoMachine = `izapple2 -model _base -board 2plus -cpu 6502 \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s0 language \
    -s6 'diskii,disk1=disks/Apple LOGO.dsk'`

/*
logoScreenshots is Apple Logo on an Apple ][+ with 64 KB: words and lists
printed, the turtle moved by hand, procedures of its own taught to it, a
flower of squares and a spiral that calls itself drawn, and the procedures
saved on the diskette.
*/
func logoScreenshots(t *testing.T) {
	pictures := newAlbum("logo", album.Green)
	o := start(t, logoMachine, nil)

	// Started, waiting at its prompt, ?
	must(t, o.WaitForText("?", 60))
	must(t, o.WaitForPrompt(30))

	// Words, numbers and lists
	must(t, o.TypeLines(
		"PRINT [HELLO FROM LOGO]",
		"PRINT 3 * 4 + 1",
		"PRINT FIRST [TURTLE GRAPHICS]",
		`PRINT BUTFIRST "TURTLE`,
	))
	must(t, o.WaitForText("URTLE\n?", 10))
	must(t, o.WaitForPrompt(10))
	must(t, pictures.Screenshot(o, "words"))

	// The turtle, moved by hand
	turtle := pictures.Record(o)
	must(t, turtle.TypeLines("FORWARD 60"))
	turtle.Run(30, 6)
	must(t, turtle.TypeLines("RIGHT 90", "FORWARD 60", "RIGHT 135", "FORWARD 85"))
	turtle.Run(60, 6)
	must(t, pictures.SaveRecording(turtle, "turtle", 300))

	// A procedure of its own, with an input
	must(t, o.TypeLines(
		"CLEARSCREEN",
		"TO SQUARE :SIZE",
		"REPEAT 4 [FORWARD :SIZE RIGHT 90]",
		"END",
		"SQUARE 40",
		"SQUARE 70",
	))
	must(t, o.WaitForText("SQUARE 70\n?", 30))
	must(t, o.WaitForPrompt(10))
	must(t, pictures.Screenshot(o, "squares"))
	must(t, o.TypeLines("TEXTSCREEN"))
	must(t, o.WaitForPrompt(10))
	must(t, pictures.Screenshot(o, "procedure"))

	// A procedure made of another: the flower
	must(t, o.TypeLines(
		"CLEARSCREEN",
		"TO FLOWER",
		"REPEAT 36 [SQUARE 60 RIGHT 10]",
		"END",
	))
	must(t, o.WaitForPrompt(10))
	flower := pictures.Record(o)
	must(t, flower.TypeLines("FLOWER"))
	drawUntil(t, o, flower, "FLOWER\n?")
	must(t, pictures.SaveRecording(flower, "flower", 300))

	// A procedure that calls itself
	must(t, o.TypeLines(
		"CLEARSCREEN",
		"TO SPIRAL :SIDE",
		"IF :SIDE > 120 [STOP]",
		"FORWARD :SIDE",
		"RIGHT 121",
		"SPIRAL :SIDE + 3",
		"END",
	))
	must(t, o.WaitForPrompt(10))
	spiral := pictures.Record(o)
	must(t, spiral.TypeLines("SPIRAL 1"))
	drawUntil(t, o, spiral, "SPIRAL 1\n?")
	must(t, pictures.SaveRecording(spiral, "spiral", 300))

	// The procedures listed, saved, and the diskette catalogued
	must(t, o.TypeLines("TEXTSCREEN", "POTS", `SAVE "SHAPES`))
	must(t, o.WaitForText("PROCEDURES SAVED", 30))
	must(t, o.TypeLines("CATALOG"))
	must(t, o.WaitForText("SHAPES.LOGO", 30))
	must(t, o.WaitForPrompt(10))
	must(t, pictures.Screenshot(o, "saved"))
}

// drawUntil records the turtle drawing until a text is on the screen, the
// prompt after the command
func drawUntil(t *testing.T, o *operator.Operator, r *album.Recording, text string) {
	t.Helper()
	limit := o.Frames() + 3*60*60
	for !o.HasText(text) {
		if o.Frames() > limit {
			t.Fatalf("the drawing did not end:\n%v", o.Text())
		}
		r.Run(6, 6)
	}
	r.Run(60, 6)
}
