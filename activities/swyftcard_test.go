package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// swyftCardMachine is the machine of the guide, as its command line of
// izapple2, with the diskette of the tutorial
const swyftCardMachine = `izapple2 -model none -board 2e -cpu 65c02 -screen green \
    -rom "<internal>/Apple2e_Enhanced.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s3 swyftcard \
    -s6 diskii,disk1=disks/SwyftWare_-_SwyftCard_Tutorial.woz`

// swyftCardEmpty is the same machine with no diskette in the drive, for a
// Text of one's own
const swyftCardEmpty = `izapple2 -model none -board 2e -cpu 65c02 -screen green \
    -rom "<internal>/Apple2e_Enhanced.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s3 swyftcard \
    -s6 diskii`

// The Leap keys of the SwyftCard are the Apple keys of the //e, the buttons
// of the game port
const (
	leapBackward = 0 // Open Apple
	leapForward  = 1 // Solid Apple
)

/*
swyftCardScreenshots is the SwyftCard, Jef Raskin's editor in a card for the
Apple //e: its tutorial, leaping to the next page and to words, a chunk of
text highlighted, deleted and brought back; and then a letter written from
nothing, with a sentence moved by leaping.
*/
func swyftCardScreenshots(t *testing.T) {
	pictures := newAlbum("swyftcard", album.Green)
	o := start(t, swyftCardMachine, nil)

	// The first page of the tutorial
	must(t, o.WaitForText("HOW TO USE SWYFTCARD", 30))
	o.RunSeconds(1)
	must(t, pictures.Screenshot(o, "tutorial"))

	// Solid Apple held and = tapped three times: to the next page
	first := pictures.Record(o)
	first.Capture(100)
	hand := swyftHand{o, first}
	hand.leap(t, leapForward, "===")
	hand.waitFor(t, "THE CURSOR")
	hand.waitFor(t, "_==")
	first.Capture(10)
	must(t, pictures.SaveRecording(first, "first-leap", 300))

	// The leaping cursor: forward to j, backward to x, forward to fr
	quick := swyftHand{o, nil}
	quick.leap(t, leapForward, "===")
	quick.waitFor(t, "THE LEAPING CURSOR")
	leaping := pictures.Record(o)
	leaping.Capture(100)
	hand = swyftHand{o, leaping}
	hand.leap(t, leapForward, "j")
	hand.leap(t, leapBackward, "x")
	hand.leap(t, leapForward, "fr")
	hand.waitFor(t, "_rog")
	leaping.Capture(10)
	must(t, pictures.SaveRecording(leaping, "leaping", 300))

	// On to the page of the chunk: Martin's first sentence highlighted,
	// deleted, and brought back with Use Front A
	for !o.HasText("HOW TO DELETE A CHUNK OF TEXT AND BRING IT BACK") {
		quick.leap(t, leapForward, "===")
	}
	chunk := pictures.Record(o)
	chunk.Capture(100)
	hand = swyftHand{o, chunk}
	hand.leap(t, leapForward, "Mar")
	hand.leap(t, leapForward, ".")
	hand.highlight()
	chunk.Run(60, 6)
	must(t, o.Key("Delete"))
	hand.waitGone(t, "peered out the window")
	chunk.Run(60, 6)
	must(t, o.Key("Ctrl+A"))
	hand.waitFor(t, "peered out the window")
	chunk.Run(30, 6)
	must(t, pictures.SaveRecording(chunk, "chunk", 300))

	// Started again with no diskette: a Text of one's own, empty
	o = start(t, swyftCardEmpty, nil)
	must(t, o.WaitForText("= 1 =", 30))
	o.RunSeconds(1)
	swyftHand{o, nil}.waitFor(t, "\n_=")
	must(t, pictures.Screenshot(o, "empty"))

	// A letter typed
	typed := pictures.Record(o)
	typed.Capture(100)
	for _, line := range letter {
		must(t, typed.Type(line+"\n"))
	}
	hand = swyftHand{o, typed}
	hand.waitFor(t, "Jef_")
	typed.Capture(10)
	must(t, pictures.SaveRecording(typed, "letter", 300))

	// The thanks moved before the date: leap back to it, leap forward to
	// its full stop, highlight, Delete, leap back to See, Use Front A
	moved := pictures.Record(o)
	moved.Capture(100)
	hand = swyftHand{o, moved}
	hand.leap(t, leapBackward, "Tha")
	hand.leap(t, leapForward, ".")
	hand.highlight()
	moved.Run(60, 6)
	must(t, o.Key("Delete"))
	hand.waitGone(t, "Thank you")
	moved.Run(60, 6)
	hand.leap(t, leapBackward, "See")
	must(t, o.Key("Ctrl+A"))
	hand.waitFor(t, "twice.See")
	must(t, moved.Type("  "))
	hand.waitFor(t, "twice.  _ee")
	moved.Capture(10)
	must(t, pictures.SaveRecording(moved, "moved", 300))
	must(t, pictures.Screenshot(o, "moved"))
}

// letter is the letter of the page, a line at a time
var letter = []string{
	"Dear Ada,",
	"",
	"See you on Tuesday at the club.  Thank you for the books, I have read the one about the Apple II twice.",
	"",
	"Jef",
}

// swyftHand is a hand on the keys of the SwyftCard, with the machine
// recorded when it has a recording
type swyftHand struct {
	o *operator.Operator
	r *album.Recording
}

// run runs the machine some frames, recording them if recorded
func (h swyftHand) run(frames int) {
	if h.r != nil {
		h.r.Run(frames, 6)
	} else {
		h.o.Run(frames)
	}
}

// leap holds a Leap key, taps the keys of a pattern and lets go: the cursor
// goes to the first letter of the pattern found, forward or backward. It
// takes the SwyftCard less than half a second.
func (h swyftHand) leap(t *testing.T, key int, pattern string) {
	t.Helper()
	h.o.HoldButton(key)
	h.run(10)
	for _, c := range pattern {
		must(t, h.o.Type(string(c)))
		h.run(20)
	}
	h.run(30)
	h.o.ReleaseButton(key)
	h.run(60)
}

// highlight presses the two Leap keys at once: the text the cursor went over
// in its last leap is highlighted
func (h swyftHand) highlight() {
	h.o.HoldButton(leapBackward)
	h.o.HoldButton(leapForward)
	h.run(20)
	h.o.ReleaseButton(leapBackward)
	h.o.ReleaseButton(leapForward)
	h.run(60)
}

// waitFor runs the machine until a text is on the screen
func (h swyftHand) waitFor(t *testing.T, text string) {
	t.Helper()
	for limit := h.o.Frames() + 10*operator.FramesPerSecond; !h.o.HasText(text); {
		if h.o.Frames() > limit {
			t.Fatalf("%q did not come:\n%v", text, h.o.Text())
		}
		h.run(6)
	}
}

// waitGone runs the machine until a text is no longer on the screen
func (h swyftHand) waitGone(t *testing.T, text string) {
	t.Helper()
	for limit := h.o.Frames() + 10*operator.FramesPerSecond; h.o.HasText(text); {
		if h.o.Frames() > limit {
			t.Fatalf("%q did not go:\n%v", text, h.o.Text())
		}
		h.run(6)
	}
}
