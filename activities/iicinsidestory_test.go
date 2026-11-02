package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// iicInsideStoryMachine is the machine of the guide, as its command line of
// izapple2, with The Inside Story
const iicInsideStoryMachine = `izapple2 -model none -board 2c -cpu 65c02 -screen color \
    -rom "<internal>/Apple2c_ROM255.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s1 iicserial \
    -s2 iicserial \
    -s6 "diskii,disk1=disks/Apple Presents The IIc - The Inside Story 691-0020-A (Apple-1984).dsk"`

/*
iicInsideStoryScreenshots is The Inside Story, a diskette of Apple Presents
the IIc: the case of Sherlock Capslock and Nancy Novice, the tour of the
inside of the IIc that he gives her, and a turn of its game
*/
func iicInsideStoryScreenshots(t *testing.T) {
	pictures := newAlbum("iic-inside-story", album.Color)
	o := start(t, iicInsideStoryMachine, nil)
	iic := &iicTour{t: t, o: o, pictures: pictures}

	// The title, and the menu
	iic.settle(nil, 4, 10000)
	must(t, pictures.Screenshot(o, "title"))
	iic.key("Return")
	must(t, pictures.Screenshot(o, "menu"))

	// Part 1, The Problem at Hand: the case, a page at each Return
	iic.typed("2\n")
	iic.pages("case", 11)
	must(t, pictures.Screenshot(o, "tour-offered"))

	// Sherlock's inside tour of the IIc
	iic.pages("tour", 26)
	must(t, pictures.Screenshot(o, "rest-offered"))

	// Escape, back to the menu, and the game
	iic.key("Escape")
	iic.typed("5\n")
	must(t, pictures.Screenshot(o, "game"))
	iic.key("Return")
	must(t, pictures.Screenshot(o, "board"))

	// A key stops the spinner, and an arrow moves as many squares that way
	must(t, o.Key("Space"))
	o.RunSeconds(2)
	must(t, o.Key("Right"))
	o.RunSeconds(5)
	must(t, pictures.Screenshot(o, "question"))
	must(t, o.Type("2"))
	o.RunSeconds(3)
	must(t, pictures.Screenshot(o, "answered"))
}

// pages records a GIF of a number of pages of the guide, a Return each,
// each page held three seconds
func (i *iicTour) pages(name string, pages int) {
	i.t.Helper()
	r := i.pictures.Record(i.o)
	r.Capture(300)
	for range pages {
		before := i.pictures.Screen(i.o)
		must(i.t, i.o.Key("Return"))
		i.goneOnRecording(r, before)
		r.Capture(300)
	}
	must(i.t, i.pictures.SaveRecording(r, name, 300))
}
