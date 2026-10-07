package activities

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

/*
TestExplore is for trying a program on the machine before writing the
generator of its page. It starts the machine of a command line and runs a
script of steps, both from the environment, printing the text of the screen
and writing a picture at each shot:

	EXPLORE  the folder for the pictures, outside the repository; without
	         it the test is skipped
	MACHINE  the command line of izapple2, as the pages give it
	MONITOR  color for the colour monitor; green if not given
	STEPS    a step a line, its name and its argument after a colon

The steps:

	type:TEXT     types a text, \n for Return
	key:NAME      presses a key by its name: Return, Escape, Ctrl+C...
	apple:TEXT    types a text with Open Apple held
	wait:TEXT     runs until a text is on the screen, two minutes at most
	kbd           runs until the program reads the keyboard
	run:SECONDS   runs the machine some seconds
	shot:NAME     prints the screen and writes NN-NAME.png
	reset         presses Reset
	listing:FILE  types a listing, a line at a time
	text:FILE     types a file whole, into an editor
	insert:N,DISK puts a disk of disks/ in drive N, 0 the first
	joy:X,Y       the joystick, 0 to 255 each
	button:N      presses and releases a button of the game port
	dot:X,Y       the mouse to a dot of the screen, 560 by 192
	drag:X,Y      holds the mouse button and drags to a dot
	click, dclick, hold, release   the button of the mouse

For example, DOS 3.3 started and catalogued, the steps one a line:

	EXPLORE=/tmp/x MACHINE="$(cat machine.txt)" STEPS='
	    wait:SYSTEM MASTER
	    kbd
	    type:CATALOG\n
	    wait:BOOT13
	    shot:catalog' go test -count=1 -run TestExplore -v ./activities

kbd is no good after CATALOG: DOS reads the keyboard while it lists, to
pause, so wait for the last file instead.
*/
func TestExplore(t *testing.T) {
	if os.Getenv("EXPLORE") == "" {
		t.Skip("explores a machine, with EXPLORE, MACHINE and STEPS")
	}
	monitor := album.Green
	if os.Getenv("MONITOR") == "color" {
		monitor = album.Color
	}
	a := album.New(os.Getenv("EXPLORE"), monitor)
	o := start(t, os.Getenv("MACHINE"), nil)
	exploreSteps(t, o, a, os.Getenv("STEPS"))
}

// exploreSteps runs the steps of a script of TestExplore
func exploreSteps(t *testing.T, o *operator.Operator, a *album.Album, steps string) {
	t.Helper()
	shots := 0
	for _, step := range strings.Split(steps, "\n") {
		step = strings.TrimSpace(step)
		if step == "" {
			continue
		}
		name, arg, _ := strings.Cut(step, ":")
		switch name {
		case "type":
			must(t, o.Type(strings.ReplaceAll(arg, `\n`, "\n")))
		case "key":
			must(t, o.Key(arg))
		case "apple":
			o.HoldButton(0)
			o.Run(5)
			must(t, o.Type(arg))
			o.ReleaseButton(0)
		case "wait":
			if err := o.WaitForText(arg, 120); err != nil {
				t.Fatal(err)
			}
		case "kbd":
			must(t, o.WaitForKeyboard(60))
		case "run":
			var seconds float64
			fmt.Sscan(arg, &seconds)
			o.RunSeconds(seconds)
		case "shot":
			shots++
			fmt.Printf("== %s mode=%x\n%v\n", arg, o.Apple2().GetVideoSource().GetCurrentVideoMode(), o.Text())
			must(t, a.Screenshot(o, fmt.Sprintf("%02d-%s", shots, arg)))
		case "reset":
			o.Reset()
		case "listing":
			must(t, typeListing(o, arg))
		case "text":
			must(t, typeText(o, arg))
		case "insert":
			var drive int
			var file string
			fmt.Sscanf(arg, "%d,%s", &drive, &file)
			must(t, o.InsertDisk(drive, disk(t, file)))
		case "joy":
			var x, y int
			fmt.Sscanf(arg, "%d,%d", &x, &y)
			o.Joystick(uint8(x), uint8(y))
		case "button":
			var button int
			fmt.Sscan(arg, &button)
			o.PressButton(button)
		case "dot":
			var x, y int
			fmt.Sscanf(arg, "%d,%d", &x, &y)
			dot(o, x, y)
			o.Run(10)
		case "drag":
			var x, y int
			fmt.Sscanf(arg, "%d,%d", &x, &y)
			o.Hold()
			o.Run(20)
			fromX, fromY := o.MousePosition()
			toX, toY := atDot(x, y)
			for i := 1; i <= 20; i++ {
				o.MouseTo(fromX+(toX-fromX)*float64(i)/20, fromY+(toY-fromY)*float64(i)/20)
				o.Run(2)
			}
			o.Run(20)
			o.Release()
		case "click":
			o.Click()
		case "dclick":
			o.DoubleClick()
		case "hold":
			o.Hold()
		case "release":
			o.Release()
		default:
			t.Fatalf("no step %q", name)
		}
	}
}
