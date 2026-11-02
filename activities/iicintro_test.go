package activities

import (
	"image"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// iicIntroductionMachine is the machine of the guide, as its command line of
// izapple2, with the first diskette of Apple Presents the IIc
const iicIntroductionMachine = `izapple2 -model none -board 2c -cpu 65c02 -screen color \
    -rom "<internal>/Apple2c_ROM255.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s1 iicserial \
    -s2 iicserial \
    -s6 "iicdisk,disk1=disks/Apple Presents the IIc - An Introduction.dsk"`

/*
iicIntroductionScreenshots is the first diskette of Apple Presents the IIc,
An Introduction: its opening, the keyboard, its menu, a name typed, the
Diner that has the dish of the name, the colours and the 80 columns, and the
end, that asks for the next diskette
*/
func iicIntroductionScreenshots(t *testing.T) {
	pictures := newAlbum("iic-introduction", album.Color)
	o := start(t, iicIntroductionMachine, nil)
	iic := &iicTour{t: t, o: o, pictures: pictures}

	// The opening, from the IIc switched on, to the keyboard
	opening := pictures.Record(o).Faster(3)
	opening.Capture(50)
	iic.settle(opening, 8, 10000)
	must(t, pictures.SaveRecording(opening, "opening", 300))
	must(t, pictures.Screenshot(o, "keyboard"))

	// Escape, for those who have seen it before, to the menu
	iic.key("Escape")
	must(t, pictures.Screenshot(o, "menu"))

	// Topic 2: the name, and the menu has it
	iic.typed("2\n")
	iic.typed("Ada\n")
	must(t, pictures.Screenshot(o, "name"))
	iic.typed("YES\n")
	iic.key("Return")
	must(t, pictures.Screenshot(o, "ada-menu"))

	// Topic 3: a menu of meals, and one chosen
	iic.typed("3\n")
	must(t, pictures.Screenshot(o, "diner"))
	iic.typed("11\n")
	must(t, pictures.Screenshot(o, "quiche"))
	iic.key("Return")

	// Topic 5: colour or black and white, and 80 columns
	iic.typed("5\n")
	must(t, pictures.Screenshot(o, "butterfly"))
	iic.typed("YES\n")
	iic.key("Return")
	must(t, pictures.Screenshot(o, "fringes"))
	iic.key("Return")
	iic.key("Return")
	must(t, pictures.Screenshot(o, "forty-columns"))
	iic.key("Return")
	iic.key("Return")
	must(t, pictures.Screenshot(o, "eighty-columns"))
	for range 3 {
		iic.key("Return")
	}
	must(t, pictures.Screenshot(o, "switch"))
	for range 3 {
		iic.key("Return")
	}

	// Topic 12, the end, and the next diskette
	iic.typed("12\n")
	must(t, pictures.Screenshot(o, "topics"))
	iic.key("Return")
	must(t, pictures.Screenshot(o, "goodbye"))
}

// iicTour is the person at the IIc of the guide, through its diskettes
type iicTour struct {
	t        *testing.T
	o        *operator.Operator
	pictures *album.Album
}

// openAppleReset presses Control-Reset with Open Apple held, which starts the
// IIc again from the disk in its drive
func (i *iicTour) openAppleReset() {
	i.o.HoldButton(0)
	i.o.Reset()
	i.o.Run(30)
	i.o.ReleaseButton(0)
}

/*
settle runs the machine until the screen stays the same for some seconds,
with at least some dots lit: the tour
draws its pages in high resolution, a line at a time, with pauses to read,
and loads the next from the disk with the screen blank, and a key pressed
while it does is lost. A recording, if any, takes a picture as it goes.
*/
func (i *iicTour) settle(r *album.Recording, seconds int, lit int) {
	i.t.Helper()
	var last *image.RGBA
	still := 0
	limit := i.o.Frames() + 300*operator.FramesPerSecond
	for still < seconds*4 {
		if i.o.Frames() > limit {
			i.t.Fatal("the screen of the tour did not settle")
		}
		if r != nil {
			r.Run(15, 5)
		} else {
			i.o.Run(15)
		}
		screen := i.pictures.Screen(i.o)
		if last != nil && len(screen.Pix) == len(last.Pix) && changedDots(screen, last) < 1500 &&
			litDots(screen.Pix) >= lit {
			still++
		} else {
			still = 0
		}
		last = screen
	}
}

// key presses a key, and waits for the tour to go on
func (i *iicTour) key(name string) {
	i.t.Helper()
	before := i.pictures.Screen(i.o)
	must(i.t, i.o.Key(name))
	i.goneOn(before)
}

// typed types a text, and waits for the tour to go on
func (i *iicTour) typed(text string) {
	i.t.Helper()
	before := i.pictures.Screen(i.o)
	must(i.t, i.o.Type(text))
	i.goneOn(before)
}

// goneOn waits for the tour to change the screen from what it was before a
// key, as it may load what comes next from the disk first, and then to settle
func (i *iicTour) goneOn(before *image.RGBA) {
	i.t.Helper()
	i.goneOnRecording(nil, before)
}

// goneOnRecording is goneOn with a recording, if any, taking pictures
func (i *iicTour) goneOnRecording(r *album.Recording, before *image.RGBA) {
	i.t.Helper()
	limit := i.o.Frames() + 60*operator.FramesPerSecond
	for {
		screen := i.pictures.Screen(i.o)
		if len(screen.Pix) != len(before.Pix) || changedDots(screen, before) >= 1500 {
			break
		}
		if i.o.Frames() > limit {
			i.t.Fatal("the tour did not go on")
		}
		if r != nil {
			r.Run(6, 6)
		} else {
			i.o.Run(6)
		}
	}
	i.settle(r, 2, 2000)
}
