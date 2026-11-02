package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// iicLogoMachine is the machine of the guide, as its command line of
// izapple2, with Exploring Apple Logo
const iicLogoMachine = `izapple2 -model none -board 2c -cpu 65c02 -screen color \
    -rom "<internal>/Apple2c_ROM255.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s1 iicserial \
    -s2 iicserial \
    -s6 "diskii,disk1=disks/Apple Presents The IIc - Exploring Apple Logo 691-0015-A (Apple-1984).DSK"`

/*
iicLogoScreenshots is Exploring Apple Logo, a diskette of Apple Presents the
IIc: the turtle drawing its title, its menu, and its lesson on pictures, a
square typed a command at a time, then with REPEAT, and then taught to the
turtle as a new word, SQUARE, with Leo Logolover
*/
func iicLogoScreenshots(t *testing.T) {
	pictures := newAlbum("iic-logo", album.Color)
	o := start(t, iicLogoMachine, nil)
	iic := &iicTour{t: t, o: o, pictures: pictures}

	// The title, drawn by the turtle
	opening := pictures.Record(o)
	opening.Capture(50)
	iic.settle(opening, 4, 10000)
	must(t, pictures.SaveRecording(opening, "opening", 300))
	iic.key("Return")
	must(t, pictures.Screenshot(o, "menu"))

	// Using Logo for Pictures: a square, a command at a time
	iic.typed("3\n")
	iic.lessonReturns(nil, 7)
	must(t, pictures.Screenshot(o, "forward"))
	square := pictures.Record(o)
	square.Capture(100)
	iic.lessonType(square, "FORWARD 30\n")
	iic.lessonReturns(square, 1)
	iic.lessonType(square, "RIGHT 90\n")
	iic.lessonReturns(square, 4)
	for _, command := range []string{"FORWARD 30", "RIGHT 90", "FORWARD 30", "RIGHT 90", "FORWARD 30"} {
		iic.lessonType(square, command+"\n")
	}
	must(t, pictures.SaveRecording(square, "square", 300))
	must(t, pictures.Screenshot(o, "congratulations"))

	// Leo Logolover, and REPEAT
	iic.lessonReturns(nil, 2)
	must(t, pictures.Screenshot(o, "leo"))
	iic.lessonReturns(nil, 10)
	must(t, pictures.Screenshot(o, "repeat"))
	iic.lessonReturns(nil, 6)
	iic.lessonType(nil, "REPEAT 4 [FD 30 RT 90]\n")
	must(t, pictures.Screenshot(o, "repeated"))

	// A new word: SQUARE, unknown, then taught with TO and END
	procedure := pictures.Record(o)
	procedure.Capture(100)
	iic.lessonReturns(procedure, 1)
	iic.lessonType(procedure, "CS\n")
	iic.lessonReturns(procedure, 3)
	iic.lessonType(procedure, "SQUARE\n")
	iic.lessonReturns(procedure, 6)
	iic.lessonType(procedure, "TO SQUARE\n")
	iic.lessonReturns(procedure, 5)
	iic.lessonType(procedure, "SQUARE\n")
	must(t, pictures.SaveRecording(procedure, "procedure", 300))
	must(t, pictures.Screenshot(o, "square-defined"))
}

/*
lessonReturns presses Return at a number of pages of the lesson: each when
the guide reads the keyboard, after drawing the page, and until the screen
changes. A recording, if any, takes the pages.
*/
func (i *iicTour) lessonReturns(r *album.Recording, pages int) {
	i.t.Helper()
	for range pages {
		must(i.t, i.o.WaitForKeyboard(90))
		before := i.pictures.Screen(i.o)
		must(i.t, i.o.Key("Return"))
		i.goneOnRecording(r, before)
		if r != nil {
			r.Capture(200)
		}
	}
}

// lessonType types a command when the guide reads the keyboard, and lets the
// turtle draw it
func (i *iicTour) lessonType(r *album.Recording, text string) {
	i.t.Helper()
	must(i.t, i.o.WaitForKeyboard(90))
	i.o.Run(10)
	if r != nil {
		must(i.t, r.Type(text))
		r.Run(180, 6)
		r.Capture(200)
	} else {
		must(i.t, i.o.Type(text))
		i.o.RunSeconds(3)
	}
}
