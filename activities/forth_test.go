package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

/*
forthScreenshots is an Apple ][+ with Forth in ROM, on a card of Offete
Industries: switched on, words used and new ones defined, and the
dictionary listed.

	izapple2 -model forth
*/
func forthScreenshots(t *testing.T) {
	pictures := newAlbum("forth", album.Green)
	o := start(t, "forth", nil)

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

	// The dictionary, the newest words first
	must(t, o.TypeLines("VLIST"))
	must(t, o.WaitForPrompt(30))
	must(t, pictures.Screenshot(o, "vlist"))
}
