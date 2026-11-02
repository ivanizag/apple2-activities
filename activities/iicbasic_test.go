package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// iicBasicMachine is the machine of the guide, as its command line of
// izapple2, with Getting Down to BASIC
const iicBasicMachine = `izapple2 -model none -board 2c -cpu 65c02 -screen color \
    -rom "<internal>/Apple2c_ROM255.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s1 iicserial \
    -s2 iicserial \
    -s6 "diskii,disk1=disks/Apple Presents The IIc - Getting Down To BASIC 680-3065-A (Apple-1984).dsk"`

/*
iicBasicScreenshots is Getting Down to BASIC, a diskette of Apple Presents
the IIc: its title and menu, its introduction, which asks for a name, and
its lesson on PRINT, with numbers and an addition typed
*/
func iicBasicScreenshots(t *testing.T) {
	pictures := newAlbum("iic-basic", album.Color)
	o := start(t, iicBasicMachine, nil)
	iic := &iicTour{t: t, o: o, pictures: pictures}

	// The title, its border always moving, and the menu
	must(t, o.WaitForText("Please press RETURN to start", 60))
	o.RunSeconds(2)
	must(t, pictures.Screenshot(o, "title"))
	iic.key("Return")
	must(t, pictures.Screenshot(o, "menu"))

	// The introduction, and a name
	for range 5 {
		iic.key("Return")
	}
	must(t, pictures.Screenshot(o, "name"))
	iic.typed("ADA\n")
	iic.typed("Y\n")
	for range 3 {
		iic.key("Return")
	}
	must(t, pictures.Screenshot(o, "ready"))

	// The lesson on PRINT: a number, watched and then typed, and a sum
	lesson := pictures.Record(o)
	lesson.Capture(100)
	iic.recordedKeys(lesson, "\n", 6)
	iic.recordedKeys(lesson, "PRINT 1984\n", 1)
	must(t, pictures.Screenshot(o, "very-well"))
	iic.recordedKeys(lesson, "N\n", 1)
	iic.recordedKeys(lesson, "\n", 5)
	iic.recordedKeys(lesson, "PRINT 1977 + 7\n", 1)
	lesson.Run(120, 6)
	must(t, pictures.SaveRecording(lesson, "lesson", 300))
	must(t, pictures.Screenshot(o, "sum"))
}

// recordedKeys types a text a number of times, the screen of the guide
// recorded as it goes on after each
func (i *iicTour) recordedKeys(r *album.Recording, text string, times int) {
	i.t.Helper()
	for range times {
		before := i.pictures.Screen(i.o)
		must(i.t, r.Type(text))
		i.goneOnRecording(r, before)
		r.Capture(200)
	}
}
