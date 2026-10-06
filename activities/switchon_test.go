package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// switchOnMachine is the machine of the guide, as its command line of izapple2
const switchOnMachine = `izapple2 -model none -board 2plus -cpu 6502 -screen green \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps`

/*
switchOnScreenshots is an Apple ][+ switched on with no disk drive: Applesoft
BASIC waiting, a few commands, a program typed, listed and run, and a loop
that never ends stopped with Control-C.
*/
func switchOnScreenshots(t *testing.T) {
	pictures := newAlbum("switch-on", album.Green)
	o := start(t, switchOnMachine, nil)

	// Switched on: it beeps and is ready at once
	must(t, o.WaitForText("]", 5))
	o.Run(30)
	must(t, pictures.Screenshot(o, "switched-on"))

	// The first commands, typed
	first := pictures.Record(o)
	first.Capture(50)
	must(t, first.TypeLines(`PRINT "HELLO"`, "PRINT 2+2", "PRINT 355/113"))
	first.Run(60, 6)
	must(t, pictures.SaveRecording(first, "first-commands", 300))

	// A program typed and listed
	must(t, o.TypeLines("HOME"))
	must(t, o.TypeLines(
		`10 INPUT "WHAT IS YOUR NAME? ";N$`,
		"20 FOR I = 1 TO 10",
		`30 PRINT I;" HELLO, ";N$`,
		"40 NEXT I",
		"LIST",
	))
	o.Run(30)
	must(t, pictures.Screenshot(o, "program"))

	// Run, with a name typed at its question
	must(t, o.TypeLines("HOME"))
	running := pictures.Record(o)
	running.Capture(50)
	must(t, running.TypeLines("RUN"))
	must(t, running.TypeLines("ADA"))
	running.Run(60, 6)
	must(t, pictures.SaveRecording(running, "run", 300))

	// A loop that never ends, stopped with Control-C
	must(t, o.TypeLines("NEW", "HOME"))
	must(t, o.TypeLines(`10 PRINT "APPLE ][+ ";`, "20 GOTO 10"))
	loop := pictures.Record(o)
	loop.Capture(50)
	must(t, loop.TypeLines("RUN"))
	loop.Run(90, 4)
	must(t, o.Key("Ctrl+C"))
	loop.Run(30, 6)
	must(t, pictures.SaveRecording(loop, "break", 300))
	must(t, o.WaitForText("BREAK IN", 2))
}
