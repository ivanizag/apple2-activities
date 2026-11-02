package activities

import (
	"strings"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// iicAdaMachine is Ada's Apple IIc, as its command line of izapple2: its
// modem port waits for a call on port 1978 of the computer
const iicAdaMachine = `izapple2 -model none -board 2c -cpu 65c02 -screen green \
    -rom "<internal>/Apple2c_ROM255.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s1 iicserial \
    -s2 iicserial,device=tcp \
    -s6 iicdisk`

// iicGraceMachine is Grace's Apple IIc: a modem on its modem port, which
// dials the internet
const iicGraceMachine = `izapple2 -model none -board 2c -cpu 65c02 -screen green \
    -rom "<internal>/Apple2c_ROM255.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s1 iicserial \
    -s2 iicserial,device=netmodem \
    -s6 iicdisk`

/*
iicTerminalScreenshots is two Apple IIcs, Ada's and Grace's, talking through
their modem ports with nothing but the firmware of the IIc: each in its
terminal mode, Grace's modem dials Ada's IIc, they write to each other, and
Grace hangs up. Ada's pictures are labelled with her name and go on the
left of the page, Grace's on the right.
*/
func iicTerminalScreenshots(t *testing.T) {
	pictures := newAlbum("iic-terminal", album.Green)
	ada := start(t, iicAdaMachine, nil)
	grace := start(t, iicGraceMachine, nil)
	pair := iicPair{ada, grace}

	// What one IIc sends reaches the other on the clock of the host. At the
	// speed of real machines, it does well within the frame it was sent in,
	// and the pictures are the same on every run
	ada.RealTime(true)
	grace.RealTime(true)

	// Switched on with no disk, each IIc looks for one: Reset, to BASIC
	pair.run(3*60, nil)
	ada.Reset()
	grace.Reset()
	pair.run(60, nil)

	// Ada's IIc in its terminal mode, echoing what she types
	pair.typeOn(t, ada, nil, "HOME\n")
	pair.typeOn(t, ada, nil, "IN#2\n")
	pair.command(t, ada, nil, "T")
	pair.command(t, ada, nil, "I")
	must(t, pictures.ScreenshotOf(ada, "ada-terminal", "Ada"))

	// Grace's in its terminal mode too, and her modem dialing Ada's IIc
	dial := pictures.Record(grace).Labelled("Grace")
	dial.Capture(50)
	pair.typeOn(t, grace, dial, "HOME\n")
	pair.typeOn(t, grace, dial, "IN#2\n")
	pair.command(t, grace, dial, "T")
	pair.typeOn(t, grace, dial, "AT\n")
	pair.waitFor(t, grace, dial, "OK")
	// The connection is made on the clock of the host, after a different
	// number of frames each time: the recording waits for it without
	// pictures, for the same ones on every run
	pair.typeOn(t, grace, dial, "ATD localhost:1978\n")
	pair.waitFor(t, grace, nil, "CONNECT")
	dial.Capture(50)
	// Connected, the modem doesn't echo any more: the IIc does
	pair.command(t, grace, dial, "I")
	pair.run(60, dial)
	must(t, pictures.SaveRecording(dial, "grace-dial", 300))

	// Grace writes, and Ada reads it, and answers
	hello := pictures.Record(grace).Labelled("Grace")
	hello.Capture(50)
	pair.typeOn(t, grace, hello, "Hello Ada, can you read me?\n")
	pair.waitFor(t, ada, nil, "can you read me?")
	pair.run(60, hello)
	must(t, pictures.SaveRecording(hello, "grace-hello", 300))

	answer := pictures.Record(ada).Labelled("Ada")
	answer.Capture(50)
	pair.typeOn(t, ada, answer, "Loud and clear, Grace!\n")
	pair.waitFor(t, grace, nil, "Loud and clear")
	pair.run(60, answer)
	must(t, pictures.SaveRecording(answer, "ada-answer", 300))
	must(t, pictures.ScreenshotOf(grace, "grace-read", "Grace"))

	// Grace hangs up: a second of nothing, +++, a second more, and ATH
	bye := pictures.Record(grace).Labelled("Grace")
	bye.Capture(50)
	pair.command(t, grace, bye, "40N")
	pair.run(90, bye)
	answers := strings.Count(grace.Text(), "OK")
	pair.typeOn(t, grace, bye, "+++")
	// The modem answers a second after, timed on the clock of the host: two
	// seconds of the machines, without pictures, for the same ones on every
	// run
	pair.run(2*operator.FramesPerSecond, nil)
	if strings.Count(grace.Text(), "OK") == answers {
		t.Fatalf("the modem did not answer +++:\n%v", grace.Text())
	}
	bye.Capture(50)
	pair.typeOn(t, grace, bye, "ATH\n")
	for limit := grace.Frames() + 10*operator.FramesPerSecond; strings.Count(grace.Text(), "OK") == answers+1; {
		if grace.Frames() > limit {
			t.Fatalf("the modem did not hang up:\n%v", grace.Text())
		}
		pair.run(6, bye)
	}
	pair.run(60, bye)
	must(t, pictures.SaveRecording(bye, "grace-hangup", 300))
}

// iicPair is the two IIcs, which run together, a frame each in turn, as
// the two machines on a desk did
type iicPair struct {
	ada, grace *operator.Operator
}

// run runs both for some frames, with a picture every six for a recording
func (p iicPair) run(frames int, r *album.Recording) {
	for i := range frames {
		p.ada.Run(1)
		p.grace.Run(1)
		if r != nil && i%6 == 5 {
			r.Capture(10)
		}
	}
}

// typeOn types a text on one of the IIcs, a key at a time, the other
// running too
func (p iicPair) typeOn(t *testing.T, o *operator.Operator, r *album.Recording, text string) {
	t.Helper()
	for _, c := range text {
		must(t, o.Type(string(c)))
		if r != nil {
			r.Capture(10)
		}
		p.run(4, r)
	}
}

// command gives a command to the firmware of the modem port: Control-A
// and its letters, as T for the terminal mode or I to echo on the screen
// what is typed
func (p iicPair) command(t *testing.T, o *operator.Operator, r *album.Recording, command string) {
	t.Helper()
	must(t, o.Key("Ctrl+A"))
	p.run(10, r)
	p.typeOn(t, o, r, command)
	p.run(30, r)
}

// waitFor runs both until a text is on the screen of one, ten seconds at
// most
func (p iicPair) waitFor(t *testing.T, o *operator.Operator, r *album.Recording, text string) {
	t.Helper()
	for limit := o.Frames() + 10*operator.FramesPerSecond; !o.HasText(text); {
		if o.Frames() > limit {
			t.Fatalf("%q did not come:\n%v", text, o.Text())
		}
		p.run(6, r)
	}
}
