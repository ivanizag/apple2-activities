package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

/*
appleIIScreenshots is the Apple ][ of 1977, with the ROM of Integer BASIC and
no cards: the Monitor, its Mini-Assembler, Integer BASIC and its colours.

	izapple2 -model 2
*/
func appleIIScreenshots(t *testing.T) {
	pictures := newAlbum("apple-ii", album.Green)
	o := start(t, "2", nil)

	// Switched on, in the Monitor
	must(t, o.WaitForKeyboard(5))
	o.Run(30)
	must(t, pictures.Screenshot(o, "switched-on"))

	// Memory shown in hexadecimal, and the ROM disassembled
	clearMonitorScreen(t, o)
	must(t, o.TypeLines("F800L"))
	o.Run(30)
	must(t, pictures.Screenshot(o, "monitor"))

	// A program in the Mini-Assembler, listed and run
	clearMonitorScreen(t, o)
	assembling := pictures.Record(o)
	assembling.Capture(50)
	must(t, assembling.TypeLines(
		"F666G",
		"300:LDA #C1",
		" JSR FDED",
		" CLC",
		" ADC #1",
		" CMP #DB",
		" BNE 302",
		" RTS",
		"$FF69G",
		"300G",
	))
	assembling.Run(60, 6)
	must(t, pictures.SaveRecording(assembling, "mini-assembler", 300))

	// Integer BASIC
	must(t, o.Key("Ctrl+B"))
	must(t, o.Key("Return"))
	must(t, o.TypeLines(
		"CALL -936",
		`10 PRINT "HELLO"`,
		"20 GOTO 10",
		"LIST",
		"PRINT 7/2",
		"PRINT 32767+1",
	))
	o.Run(30)
	must(t, pictures.Screenshot(o, "integer-basic"))

	// The sixteen colours of the low resolution graphics
	must(t, o.TypeLines(
		"NEW",
		"CALL -936",
		"10 GR",
		"20 FOR I=0 TO 15",
		"30 COLOR=I",
		"40 VLIN 0,39 AT I*2+4",
		"50 VLIN 0,39 AT I*2+5",
		"60 NEXT I",
		"70 END",
		"RUN",
	))
	o.Run(60)
	must(t, pictures.On(album.Color).Screenshot(o, "colours"))
}

// clearMonitorScreen clears the screen from the Monitor, with Escape and @
func clearMonitorScreen(t *testing.T, o *operator.Operator) {
	t.Helper()
	must(t, o.Key("Escape"))
	must(t, o.Type("@"))
}
