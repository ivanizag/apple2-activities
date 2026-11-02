package activities

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
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
iicBasicLessons is the walk through the lessons of the diskette, a step a
line:

	u:TEXT   Return at each page until TEXT is on the screen
	f:TEXT   until TEXT is on the screen, with no key
	t:TEXT   a line typed, with Return
	r        Return
	w:N      N seconds
	s:NAME   a picture of the screen
	g:NAME   a GIF of the pages from here, until the next g or e
	e        the end of the GIF
*/
const iicBasicLessons = `
g:introduction
u:first name
t:ADA
s:name
t:Y
g:numbers
u:Type PRINT
t:PRINT 1984
s:very-well
t:N
u:addition problem
t:PRINT 1977 + 7
f:the symbols
s:sum
u:other calculations
t:Y
t:PRINT 4 * 12
t:N
g:words
u:to print a literal
t:PRINT "HELLO, ADA"
t:N
u:PRUNT 50
t:PRUNT 50
s:prunt
g:programs
u:execute the new program
t:RUN
u:notice that the
t:RUN
u:statement on line 40
t:40 PRINT "AND SO DO YOU"
u:Why don't you run it
t:RUN
s:poem
u:Why not try it now
t:40 PRINT "AND IT LOVES ME TOO"
u:Now type LIST
t:LIST
u:Now run the new program
t:RUN
u:Try erasing line 30
t:30
u:Now list the
t:LIST
g:saving
u:To review what your program
t:LIST
u:Let's name your program POEM
t:SAVE POEM
u:display the catalog
t:CAT
s:catalog
u:NEW         (and press RETURN)
t:NEW
u:Try running your
t:RUN
u:You saved your program as POEM
t:LOAD POEM
u:Try listing your program now
t:LIST
u:why don't you change line
t:30 PRINT "SUGAR IS SWEET"
u:save your modified poem now
t:SAVE SWEET
u:Now catalog the
t:CAT
s:two-poems
g:variables
u:Do you know what to type to tell
t:20 PRINT A
u:three-line program
t:30 END
u:becoming a first-rate
t:RUN
s:variable
u:program using variables
t:Y
u:Try it yourself
t:10 LET A = 1984
u:Now run the program and
t:RUN
u:First of all, please clear
t:NEW
u:We saved the program on the
t:LOAD NUMBER
w:5
r
t:RUN
u:enter you favorite
t:7
s:number
u:Then the INPUT statement waits
t:42
u:and line 50 causes the
t:N
u:Would you like to see
t:Y
u:type in your name
t:GRACE
u:Please erase this program
s:grace
t:NEW
u:There's no INPUT
t:20 INPUT I$
u:Your program is complete.  Why
t:LIST
u:Now try running it
t:RUN
u:Type in your favorite flavor
t:PISTACHIO
s:ice-cream
u:Very good!  It did
t:N
u:You are about to leave
e
s:leave
`

/*
iicBasicScreenshots is Getting Down to BASIC, a diskette of Apple Presents
the IIc that teaches Applesoft: its title and menu, and all its lessons,
walked from the first to the last as the guide asks, a GIF of each, from
PRINT to the variables of INPUT
*/
func iicBasicScreenshots(t *testing.T) {
	pictures := newAlbum("iic-basic", album.Color)
	o := start(t, iicBasicMachine, nil)
	w := &iicBasicWalk{t: t, o: o, pictures: pictures}

	// The title, its border always moving, and the menu
	must(t, o.WaitForText("Please press RETURN to start", 60))
	o.RunSeconds(2)
	must(t, pictures.Screenshot(o, "title"))
	must(t, o.WaitForKeyboard(60))
	must(t, o.Key("Return"))
	w.step("f:Getting Down to BASIC Menu")
	must(t, pictures.Screenshot(o, "menu"))

	// The lessons, the introduction chosen with Return
	for _, step := range strings.Split(iicBasicLessons, "\n") {
		if step = strings.TrimSpace(step); step != "" {
			w.step(step)
		}
	}
}

// iicBasicWalk is the person at the lessons, and the GIF being recorded
type iicBasicWalk struct {
	t        *testing.T
	o        *operator.Operator
	pictures *album.Album
	rec      *album.Recording
	name     string
}

// step does a step of iicBasicLessons
func (w *iicBasicWalk) step(step string) {
	w.t.Helper()
	command, arg, _ := strings.Cut(step, ":")
	switch command {
	case "u":
		for tries := 0; !w.o.HasText(arg); tries++ {
			if tries > 60 {
				w.t.Fatalf("%q never came:\n%v", arg, w.o.Text())
			}
			w.key("Return")
		}
	case "f":
		for frames := 0; !w.o.HasText(arg); frames += 6 {
			if frames > 60*60 {
				w.t.Fatalf("%q never came:\n%v", arg, w.o.Text())
			}
			w.run(6)
		}
		w.settle()
	case "t":
		must(w.t, w.o.WaitForKeyboard(60))
		if w.rec != nil {
			must(w.t, w.rec.Type(arg+"\n"))
		} else {
			must(w.t, w.o.Type(arg+"\n"))
		}
		w.settle()
	case "r":
		w.key("Return")
	case "w":
		var seconds int
		fmt.Sscan(arg, &seconds)
		w.run(seconds * 60)
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

// key presses a key when the lesson reads the keyboard, and waits for the
// page to settle
func (w *iicBasicWalk) key(name string) {
	w.t.Helper()
	must(w.t, w.o.WaitForKeyboard(60))
	must(w.t, w.o.Key(name))
	w.settle()
}

// run runs the machine, and the recording if any
func (w *iicBasicWalk) run(frames int) {
	if w.rec != nil {
		w.rec.Run(frames, 6)
	} else {
		w.o.Run(frames)
	}
}

/*
settle runs until the text of the screen stays the same for a second and a
half, but for the cursor that blinks in it, and then holds the page in the
GIF: the lessons type their pages a letter at a time. The pages of graphics,
whose text doesn't settle, are given twenty seconds.
*/
func (w *iicBasicWalk) settle() {
	w.t.Helper()
	last, still := "", 0
	for frames := 0; still < 90 && frames < 20*60; frames += 6 {
		w.run(6)
		text := strings.ReplaceAll(w.o.Text(), "_", " ")
		if text == last {
			still += 6
		} else {
			last, still = text, 0
		}
	}
	if w.rec != nil {
		w.rec.Capture(200)
	}
}
