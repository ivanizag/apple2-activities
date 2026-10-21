package activities

import (
	"strings"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// flightSimMachine is the machine of the guide, as its command line of izapple2
const flightSimMachine = `izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s0 language \
    -s6 "diskii,disk1=disks/Flight Simulator II v2.0.woz"`

/*
flightSimScreenshots is subLOGIC's Flight Simulator II on an Apple ][+: its
demonstration, and then a flight from Meigs Field in Chicago, looked at from
the cockpit and on the radar, the sound turned on in the editor, the take-off
and a turn past the towers of the city, a dive into Lake Michigan, and the
same airport at night.

The simulator draws a frame a second or so and does not tell when it is done,
so the steps of a flight are timed, as the pilot keeps time; the flight is the
same each run, as the machine is.
*/
func flightSimScreenshots(t *testing.T) {
	pictures := newAlbum("flight-simulator", album.Color)

	// The demonstration, until it has flown a while
	o := start(t, flightSimMachine, nil)
	waitForFlightSimMenu(t, o)
	must(t, pictures.Screenshot(o, "menu"))
	must(t, o.Type("A"))
	must(t, o.WaitForKeyboard(30))
	o.RunSeconds(1)
	must(t, pictures.Screenshot(o, "modes"))
	must(t, o.Type("A"))
	must(t, o.WaitForKeyboard(30))
	o.RunSeconds(1)
	must(t, pictures.Screenshot(o, "demo-start"))
	demo := pictures.Record(o).Faster(4)
	demo.Capture(100)
	must(t, o.Key("Space"))
	demo.Run(60*150, 60)
	must(t, pictures.SaveRecording(demo, "demo", 300))

	// Started again, for a flight of one's own
	o = start(t, flightSimMachine, nil)
	waitForFlightSimMenu(t, o)
	must(t, o.Type("A"))
	must(t, o.WaitForKeyboard(30))
	must(t, o.Type("B"))
	o.RunSeconds(10)
	must(t, pictures.Screenshot(o, "runway"))

	// The views out of the cockpit, and the radar
	view(t, o, "5F")
	must(t, pictures.Screenshot(o, "left"))
	view(t, o, "5B")
	must(t, pictures.Screenshot(o, "back"))
	view(t, o, "5T")
	view(t, o, "4")
	must(t, pictures.Screenshot(o, "radar"))
	view(t, o, "..")
	must(t, pictures.Screenshot(o, "radar-zoom"))
	view(t, o, "5T")

	// The editor: the sound on, and the second page
	openEditor(t, o)
	must(t, o.Type("\n1\n"))
	o.RunSeconds(1)
	must(t, pictures.Screenshot(o, "editor"))
	for i := 0; i < 19; i++ {
		must(t, o.Key("Return"))
	}
	must(t, o.WaitForText("ENVIRONMENTAL CONTROL", 10))
	o.RunSeconds(1)
	must(t, pictures.Screenshot(o, "editor-2"))
	closeEditor(t, o)

	// The take-off: elevator up three notches, the throttle full
	takeoff := pictures.Record(o)
	takeoff.Capture(100)
	must(t, o.Type("BBB"))
	for i := 0; i < 16; i++ {
		must(t, o.Key("Right"))
	}
	takeoff.Run(60*30, 30)

	// The elevator down two notches, a gentle bank left, and the towers
	must(t, o.Type("TT"))
	takeoff.Run(60*10, 30)
	must(t, pictures.SaveRecording(takeoff, "takeoff", 300))
	past := pictures.Record(o)
	past.Capture(100)
	must(t, o.Type("F"))
	past.Run(60, 30)
	must(t, o.Type("GB"))
	past.Run(60*30, 30)
	must(t, pictures.SaveRecording(past, "tower", 300))

	// Into the ground: a steep bank, the nose down, and the crash
	dive := pictures.Record(o)
	dive.Capture(100)
	must(t, o.Type("FFFTTTT"))
	dive.Run(60*36, 30)
	must(t, pictures.SaveRecording(dive, "crash", 300))
	must(t, pictures.Screenshot(o, "after-crash"))

	// Nine in the evening, on the start of the runway, as the editor had
	// them at first: the position on its first page, the hour on the second
	openEditor(t, o)
	editorTo(t, o, "NORTH POSITION")
	must(t, o.Type("17188\n16671\n591\n0\n0\n0\n"))
	editorTo(t, o, "TIME: HOURS")
	must(t, o.Type("21\n"))
	closeEditor(t, o)
	waitForNightSky(t, o, pictures)
	o.RunSeconds(5)
	must(t, pictures.Screenshot(o, "night"))
	view(t, o, "4")
	must(t, pictures.Screenshot(o, "night-radar"))
	view(t, o, "5T")
}

// waitForFlightSimMenu waits for the first question of the simulator, which
// reads the keyboard once it is loaded
func waitForFlightSimMenu(t *testing.T, o *operator.Operator) {
	t.Helper()
	must(t, o.WaitForKeyboard(120))
	o.RunSeconds(1)
}

// view types the keys of a view, or of the radar, and gives the simulator the
// seconds it takes to draw it
func view(t *testing.T, o *operator.Operator, keys string) {
	t.Helper()
	must(t, o.Type(keys))
	o.RunSeconds(5)
}

// openEditor stops the flight for the editor, Escape, which comes from the disk
func openEditor(t *testing.T, o *operator.Operator) {
	t.Helper()
	must(t, o.Key("Escape"))
	must(t, o.WaitForText("HIT ESC TO EXIT", 30))
	must(t, o.WaitForKeyboard(10))
}

// closeEditor goes back to the flight, and waits for its view to be drawn
func closeEditor(t *testing.T, o *operator.Operator) {
	t.Helper()
	must(t, o.Key("Escape"))
	o.RunSeconds(10)
}

// editorTo presses Return until the arrow of the editor is at a line
func editorTo(t *testing.T, o *operator.Operator, line string) {
	t.Helper()
	for i := 0; i < 50; i++ {
		for _, l := range strings.Split(o.Text(), "\n") {
			if strings.Contains(l, line) && strings.Contains(l, "-->") {
				return
			}
		}
		must(t, o.Key("Return"))
		o.Run(10)
	}
	t.Fatalf("the editor has no line %q", line)
}

// waitForNightSky runs the machine until the sky, at the top of the view, is
// drawn black
func waitForNightSky(t *testing.T, o *operator.Operator, pictures *album.Album) {
	t.Helper()
	for i := 0; i < 60; i++ {
		screen := pictures.Screen(o)
		c := screen.RGBAAt(screen.Bounds().Dx()/2, screen.Bounds().Dy()/5)
		if c.R < 0x40 && c.G < 0x40 && c.B < 0x40 {
			return
		}
		o.RunSeconds(1)
	}
	t.Fatal("the sky did not go black")
}
