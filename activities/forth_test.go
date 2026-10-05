package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// forthMachine is the machine of the guide, as its command line of izapple2
const forthMachine = `izapple2 -model _base -board 2plus -cpu 6502 \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s0 forthrom`

/*
forthScreenshots is an Apple ][+ with Forth in ROM, on a card of Offete
Industries: switched on, words used and new ones defined, loops and
decisions, memory read in hexadecimal, the speaker clicked into a tone, the
low resolution screen painted byte by byte, and the dictionary listed.
*/
func forthScreenshots(t *testing.T) {
	pictures := newAlbum("forth", album.Green)
	o := start(t, forthMachine, nil)
	sound := album.Listen(o)

	// Switched on, in Forth
	must(t, o.WaitForText("FORTH", 10))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "switched-on"))

	// Words used, and new ones defined
	words := pictures.Record(o)
	words.Capture(50)
	must(t, words.TypeLines(
		"2 3 + .",
		": SQUARE DUP * ;",
		"7 SQUARE .",
		": STARS 0 DO 42 EMIT LOOP ;",
		"10 STARS",
	))
	words.Run(60, 6)
	must(t, pictures.SaveRecording(words, "words", 300))

	// Words made of words, loops and a decision
	must(t, o.TypeLines(
		": PYRAMID 1+ 1 DO CR I STARS LOOP ;",
		"8 PYRAMID",
		": COUNTDOWN BEGIN DUP . 1- DUP 0= UNTIL DROP ;",
		"10 COUNTDOWN",
		`: SIGN? 0< IF ." NEGATIVE" ELSE ." NOT NEGATIVE" THEN ;`,
		"-5 SIGN?",
	))
	must(t, o.WaitForText("-5 SIGN? NEGATIVEOK", 10))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "control"))

	// Memory, in hexadecimal: the first bytes of the Monitor
	must(t, o.TypeLines("HEX F800 20 DUMP", "FF DECIMAL . HEX"))
	must(t, o.WaitForText("255 OK", 10))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "dump"))

	// The speaker, clicked by reading its address, into a tone
	must(t, o.TypeLines(
		": CLICK C030 C@ DROP ;",
		": TONE 0 DO CLICK 20 0 DO LOOP LOOP ;",
	))
	must(t, o.WaitForKeyboard(10))
	from := sound.Now()
	must(t, o.TypeLines("200 TONE"))
	must(t, o.WaitForKeyboard(30))
	o.Run(30)
	sound.Stop()
	must(t, pictures.SaveSound(sound.Clip(from, sound.Now()), "tone"))

	// The low resolution graphics, switched on and painted a byte at a time
	must(t, o.TypeLines(
		": GRAPHICS C050 C@ DROP C056 C@ DROP C052 C@ DROP ;",
		": NORMAL C051 C@ DROP 400 400 A0 FILL ;",
		": BARS 400 0 DO 78 0 DO I 28 MOD 2 * 5 / 11 * J I + 400 + C! LOOP 80 +LOOP ;",
		": SHOW GRAPHICS BARS KEY DROP NORMAL ;",
	))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "graphics-words"))
	color := pictures.On(album.Color)
	bars := color.Record(o)
	must(t, o.TypeLines("SHOW"))
	bars.Capture(20)
	if !o.WaitUntil(1, func() bool { return !o.InTextMode() }) {
		t.Fatal("SHOW did not switch to graphics")
	}
	// The bars are done when the last byte of the screen, at the bottom
	// right, has its colour, the last of the sixteen
	for o.Apple2().Peek(0x7f7) != 0xff {
		if o.Frames() > 60*60*5 {
			t.Fatal("the bars were not drawn")
		}
		bars.Run(30, 30)
	}
	bars.Run(60, 6)
	must(t, color.SaveRecording(bars, "bars", 300))
	must(t, o.Key("Space"))
	must(t, o.WaitForKeyboard(10))

	// The dictionary, the newest words first
	must(t, o.TypeLines("VLIST"))
	must(t, o.WaitForPrompt(30))
	must(t, pictures.Screenshot(o, "vlist"))
}
