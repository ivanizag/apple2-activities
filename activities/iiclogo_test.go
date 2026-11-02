package activities

import (
	"fmt"
	"image"
	"strings"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// iicLogoMachine is the machine of the guide, as its command line of
// izapple2, with Exploring Apple Logo
const iicLogoMachine = `izapple2 -model none -board 2c -cpu 65c02 -screen color \
    -rom "<internal>/Apple2c_ROM255.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s1 iicserial \
    -s2 iicserial \
    -s6 "iicdisk,disk1=disks/Apple Presents The IIc - Exploring Apple Logo 691-0015-A (Apple-1984).DSK"`

/*
iicLogoLessons is the walk through the six topics of the diskette, from the
introduction, a step a line:

	r:N      Return at N pages
	t:TEXT   a line typed, with Return, when the guide asks for it
	c:TEXT   keys typed, with no Return
	w:N      N seconds
	s:NAME   a picture of the screen
	g:NAME   a GIF of the pages from here, until the next g or e
	e        the end of the GIF

The guide draws its pages in high resolution, so the walk counts them: the
counts are those of a walk that types what is asked as soon as it is asked.
*/
const iicLogoLessons = `
g:introduction
r:13
g:story
r:11
g:square
r:8
s:forward
t:FORWARD 30
r:1
t:RIGHT 90
r:3
t:FORWARD 30
t:RIGHT 90
t:FORWARD 30
t:RIGHT 90
t:FORWARD 30
e
s:congratulations
r:2
s:leo
r:11
s:repeat
r:6
t:REPEAT 4 [FD 30 RT 90]
s:repeated
g:procedure
t:CS
r:3
t:SQUARE
r:6
t:TO SQUARE
r:4
t:SQUARE
e
s:square-defined
g:flower
t:CS
r:4
t:ADA
r:8
t:ADA
w:20
e
s:flower
r:11
g:house
t:HOUSE
r:2
s:house-program
r:22
g:gallery
r:12
e
r:1
g:numbers
r:16
t:12
s:interest
r:4
t:PRINT [HI KIDS]
e
s:hi-kids
r:5
g:postcards
t:POSTCARD
r:3
c:.
r:4
t:THE TURTLE SAYS HELLO
r:2
t:POSTCARD
r:2
e
s:postcard
c:.
r:4
s:done
`

/*
iicLogoScreenshots is Exploring Apple Logo, a diskette of Apple Presents the
IIc: the turtle drawing its title, its menu, and its six topics walked from
the first to the last as the guide asks, the story of Logo, a square typed a
command at a time, then with REPEAT, then taught to the turtle as a new word,
a house built of words, a gallery, numbers, and postcards of random phrases
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

	// The topics, from the introduction
	iic.typed("1\n")
	w := &iicLogoWalk{t: t, o: o, pictures: pictures}
	for _, step := range strings.Split(iicLogoLessons, "\n") {
		if step = strings.TrimSpace(step); step != "" {
			w.step(step)
		}
	}
}

// iicLogoWalk is the person at the topics, and the GIF being recorded
type iicLogoWalk struct {
	t        *testing.T
	o        *operator.Operator
	pictures *album.Album
	rec      *album.Recording
	name     string
}

// step does a step of iicLogoLessons
func (w *iicLogoWalk) step(step string) {
	w.t.Helper()
	command, arg, _ := strings.Cut(step, ":")
	switch command {
	case "r":
		var pages int
		fmt.Sscan(arg, &pages)
		for range pages {
			w.page()
		}
	case "t", "c":
		if command == "t" {
			arg += "\n"
		}
		must(w.t, w.o.WaitForKeyboard(90))
		w.still()
		if w.rec != nil {
			must(w.t, w.rec.Type(arg))
		} else {
			must(w.t, w.o.Type(arg))
		}
		w.run(60)
		w.still()
		w.hold()
	case "w":
		var seconds int
		fmt.Sscan(arg, &seconds)
		w.run(seconds * operator.FramesPerSecond)
	case "s":
		must(w.t, w.pictures.Screenshot(w.o, arg))
	case "g", "e":
		if w.rec != nil {
			w.rec.Capture(300)
			must(w.t, w.pictures.SaveRecording(w.rec, w.name, 300))
			w.rec = nil
		}
		if command == "g" {
			w.rec = w.pictures.Record(w.o)
			w.rec.Capture(100)
			w.name = arg
		}
	default:
		w.t.Fatalf("no step %q", step)
	}
}

/*
page presses Return when the guide reads the keyboard, and waits for the
page to change, as the guide may load the next from the disk first, and then
to be drawn
*/
func (w *iicLogoWalk) page() {
	w.t.Helper()
	must(w.t, w.o.WaitForKeyboard(90))
	before := w.pictures.Screen(w.o)
	must(w.t, w.o.Key("Return"))
	for frames := 0; frames < 20*operator.FramesPerSecond; frames += 6 {
		w.run(6)
		if changedDots(w.pictures.Screen(w.o), before) >= 100 {
			break
		}
	}
	w.still()
	w.hold()
}

/*
still runs until the screen stays the same for a second and a half, but for
the cursor that blinks, or for eight seconds at most, as the reminders to
press Return blink too
*/
func (w *iicLogoWalk) still() {
	var last *image.RGBA = w.pictures.Screen(w.o)
	for still, frames := 0, 0; still < 6 && frames < 8*operator.FramesPerSecond; frames += 15 {
		w.run(15)
		screen := w.pictures.Screen(w.o)
		if changedDots(screen, last) < 60 {
			still++
		} else {
			still = 0
		}
		last = screen
	}
}

// run runs the machine, and the recording if any
func (w *iicLogoWalk) run(frames int) {
	if w.rec != nil {
		w.rec.Run(frames, 6)
	} else {
		w.o.Run(frames)
	}
}

// hold holds the page in the GIF, if any, for two seconds
func (w *iicLogoWalk) hold() {
	if w.rec != nil {
		w.rec.Capture(200)
	}
}
