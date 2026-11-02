package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// iicAtWorkMachine is the machine of the guide, as its command line of
// izapple2, with The Apple at Work, its side for 40 columns
const iicAtWorkMachine = `izapple2 -model none -board 2c -cpu 65c02 -screen color \
    -rom "<internal>/Apple2c_ROM255.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s1 iicserial \
    -s2 iicserial \
    -s6 "diskii,disk1=disks/Apple Presents The IIc - The Apple At Work Side 2 40-Col 691-0014-A (Apple-1984).dsk"`

/*
iicAtWorkScreenshots is The Apple at Work, a diskette of Apple Presents the
IIc: its menu, its introduction with a survey and a quiz on the programs of
work, and the lesson of its word processor, a letter changed for other
friends, in the AppleWorks it imitates
*/
func iicAtWorkScreenshots(t *testing.T) {
	pictures := newAlbum("iic-at-work", album.Color)
	o := start(t, iicAtWorkMachine, nil)
	iic := &iicTour{t: t, o: o, pictures: pictures}

	// The title, and the menu after it
	iic.settle(nil, 4, 10000)
	must(t, pictures.Screenshot(o, "title"))
	iic.goneOn(pictures.Screen(o))
	must(t, pictures.Screenshot(o, "menu"))

	// The introduction: a desk in a mess, and one in order
	iic.key("Return")
	must(t, pictures.Screenshot(o, "introduction"))
	iic.key("Return")
	must(t, pictures.Screenshot(o, "mess"))
	iic.key("Return")
	iic.key("Return")
	must(t, pictures.Screenshot(o, "order"))
	for range 3 {
		iic.key("Return")
	}
	must(t, pictures.Screenshot(o, "ziegfeld"))

	// The survey, five questions, all answered B
	iic.key("Return")
	iic.key("Return")
	must(t, pictures.Screenshot(o, "survey"))
	for range 5 {
		iic.typed("B\n")
	}
	must(t, pictures.Screenshot(o, "congratulations"))

	// The quiz, the three kinds of programs
	for range 4 {
		iic.key("Return")
	}
	iic.typed("B\n")
	iic.key("Return")
	iic.typed("A\n")
	iic.key("Return")
	must(t, pictures.Screenshot(o, "quiz"))
	iic.typed("A\n")

	// The Walkers, AppleWorks, and the word processor
	for range 3 {
		iic.key("Return")
	}
	must(t, pictures.Screenshot(o, "walkers"))
	iic.key("Return")
	iic.key("Return")
	must(t, pictures.Screenshot(o, "appleworks"))
	for range 8 {
		iic.key("Return")
	}
	must(t, pictures.Screenshot(o, "letter"))
	iic.key("Return")

	// The lesson: down to the end of the letter, Open Apple-1 back to the
	// beginning, right to the names, and new ones typed over them
	lesson := pictures.Record(o)
	lesson.Capture(100)
	for range 40 {
		must(t, o.Key("Down"))
		lesson.Run(30, 6)
	}
	for range 2 {
		before := pictures.Screen(o)
		must(t, lesson.Type("\n"))
		iic.goneOnRecording(lesson, before)
	}
	o.HoldButton(0)
	lesson.Run(6, 6)
	must(t, lesson.Type("1"))
	o.ReleaseButton(0)
	lesson.Run(120, 6)
	for range 5 {
		must(t, o.Key("Right"))
		lesson.Run(30, 6)
	}
	lesson.Run(60, 6)
	must(t, lesson.Type("Frank and Leslie,"))
	lesson.Run(180, 6)
	must(t, pictures.SaveRecording(lesson, "lesson", 300))
	must(t, pictures.Screenshot(o, "personalized"))

	// Open Apple-Escape, back to the menu, the topic marked as done
	before := pictures.Screen(o)
	o.HoldButton(0)
	o.Run(6)
	must(t, o.Key("Escape"))
	o.ReleaseButton(0)
	iic.goneOn(before)
	must(t, pictures.Screenshot(o, "done"))
}
