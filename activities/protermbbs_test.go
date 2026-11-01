package activities

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

/*
proTermMachine is the machine of the guide, as its command line of
izapple2, with a modem to the internet on the Super Serial Card and a copy
of the first side of the disk of ProTERM in the drive, which ProTERM writes
its setup and its dialing directory to:

	cp disks/ProTERM_3_Greg_Schaefer_Intrec_Software_1993_PT3120925_Side_1.dsk proterm.dsk
*/
const proTermMachine = `izapple2 -model none -board 2e -cpu 65c02 -screen green \
    -rom "<internal>/Apple2e_Enhanced.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s2 ssc,device=netmodem \
    -s6 diskii,disk1=proterm.dsk`

// proTermSide1 is the first side of the disk of ProTERM, to set it up
const proTermSide1 = "ProTERM_3_Greg_Schaefer_Intrec_Software_1993_PT3120925_Side_1.dsk"

// proTermSide2 is the other side of the disk of ProTERM, its program
const proTermSide2 = "ProTERM_3_Greg_Schaefer_Intrec_Software_1993_PT3120925_Side_2.dsk"

// duraEuropos is the address of the BBS, a host and a port of the internet
const duraEuropos = "dura-bbs.net:6359"

/*
proTermScreenshots is ProTERM 3.1 on an enhanced Apple //e calling the BBS
Dura-Europos, an Apple //e running GBBS Pro, through the modem of izapple2:
ProTERM installed, the BBS entered in its dialing directory with the
ProTERM Special emulation, the call, a look around as a guest and a game of
Concentration+ drawn in MouseText.

The BBS may be gone the next time: this runs only with A2_BBS, on the time of
the host and at the speed of a real machine.
*/
func proTermScreenshots(t *testing.T) {
	if os.Getenv("A2_BBS") == "" {
		t.Skip("calls a BBS on the internet, with A2_BBS=1")
	}
	pictures := newAlbum("proterm-bbs", album.Green)
	o := startOnHostTime(t, proTermMachine, map[string]string{"proterm.dsk": disk(t, proTermSide1)})

	// The modem: the Null Modem Driver, ProTERM leaves the modem to the
	// one at the keyboard, 24 down the list
	must(t, o.WaitForText("Install Modem", 60))
	o.RunSeconds(1)
	must(t, pictures.Screenshot(o, "install-modem"))
	for range 24 {
		must(t, o.Key("Down"))
	}
	choose(t, o)

	// The port: the Super Serial Card, 5 down, in slot 2
	must(t, o.WaitForText("Install Modem Port", 10))
	o.RunSeconds(1)
	for range 5 {
		must(t, o.Key("Down"))
	}
	choose(t, o)
	must(t, o.WaitForText("Modem Port Slot", 10))
	o.RunSeconds(1)
	must(t, o.Type("2\n"))
	o.RunSeconds(1)
	must(t, o.Key("Return"))

	// No printer, 16 down
	must(t, o.WaitForText("Install Printer Port", 30))
	o.RunSeconds(1)
	for range 16 {
		must(t, o.Key("Down"))
	}
	choose(t, o)

	// The whole installation, accepted
	must(t, o.WaitForText("Init:", 30))
	o.RunSeconds(1)
	must(t, pictures.Screenshot(o, "installed"))
	must(t, o.Key("Return"))

	// The other side of the disk, the program
	must(t, o.WaitForText("flip disk", 30))
	o.RunSeconds(1)
	must(t, o.InsertDisk(0, disk(t, proTermSide2)))
	must(t, o.Key("Return"))
	must(t, o.WaitForText("Main Menu", 60))
	o.RunSeconds(1)
	must(t, pictures.Screenshot(o, "menu"))

	// The BBS in the dialing directory: Dial, Create System, its name, no
	// number, saved as DURA
	must(t, o.Type("D"))
	o.RunSeconds(1)
	must(t, o.Key("Return"))
	must(t, o.WaitForText("Enter a System", 30))
	o.RunSeconds(1)
	must(t, o.Type("Dura-Europos"))
	for range 4 {
		must(t, o.Key("Tab"))
	}
	must(t, o.Type("DURA"))
	o.RunSeconds(1)
	must(t, o.Key("Return"))
	o.RunSeconds(1)
	must(t, o.Key("Return"))

	// Its parameters: Emulate, nine fields on, ProTERM Special, nine on
	must(t, o.WaitForText("Edit System Parms", 30))
	o.RunSeconds(1)
	for range 9 {
		must(t, o.Key("Tab"))
	}
	for range 9 {
		must(t, o.Key("Right"))
	}
	o.RunSeconds(1)
	if !o.HasText("Emulate: ProTERM Special") {
		t.Fatalf("ProTERM Special not chosen:\n%v", o.Text())
	}
	must(t, pictures.Screenshot(o, "system"))

	// Dial, Open Apple and Return, the default button: online with the
	// modem, at the speed of a real machine from now on
	o.RealTime(true)
	o.HoldButton(0)
	o.Run(5)
	must(t, o.Key("Return"))
	o.Run(5)
	o.ReleaseButton(0)
	must(t, o.WaitForText("Online", 30))
	o.RunSeconds(1)

	// The call, typed to the modem: the BBS finds ProTERM Special, draws its
	// title, and asks who calls. Its check is quick, and fails now and then
	// over the internet: then it asks for ANSI, Q hangs up, and the call is
	// made again.
	for attempt := 0; ; attempt++ {
		if attempt == 3 {
			t.Fatalf("the BBS did not find ProTERM Special:\n%v", o.Text())
		}
		call := pictures.Record(o)
		call.Capture(50)
		must(t, bbsType(o, call, "ATDT"+duraEuropos+"\n"))
		recordUntil(t, o, call, 60, func() bool {
			return o.HasText("If you just want a quick look around") || o.HasText("Could not detect")
		})
		if o.HasText("If you just want a quick look around") {
			call.Run(60, 6)
			must(t, pictures.SaveRecording(call, "call", 300))
			break
		}
		must(t, bbsType(o, nil, "Q"))
		must(t, o.WaitForText("NO CARRIER", 60))
		o.RunSeconds(2)
	}
	must(t, pictures.Screenshot(o, "title"))

	// A guest, the news of the system and the welcome
	must(t, bbsType(o, nil, "GUEST\n"))
	must(t, o.WaitForText("Hit [Enter] to continue", 60))
	o.RunSeconds(2)
	must(t, pictures.Screenshot(o, "news"))
	must(t, o.Key("Return"))
	must(t, o.WaitForText("Press [Enter] to continue", 60))
	o.RunSeconds(2)
	must(t, pictures.Screenshot(o, "guest"))
	must(t, o.Key("Return"))
	must(t, o.WaitForText("[MAIN] (?=Help):", 60))
	o.RunSeconds(1)

	// The programs, and the games
	must(t, bbsType(o, nil, "?\n"))
	must(t, o.WaitForText("Which:", 30))
	o.RunSeconds(1)
	must(t, bbsType(o, nil, "P\n"))
	must(t, o.WaitForText("LEADER BOARD", 30))
	o.RunSeconds(2)
	must(t, pictures.Screenshot(o, "programs"))
	must(t, bbsType(o, nil, "GAMES\n"))
	must(t, o.WaitForText("CONCENTRATION+", 30))
	o.RunSeconds(2)
	must(t, pictures.Screenshot(o, "games"))

	// Concentration+, played to the end
	must(t, bbsType(o, nil, "26\n"))
	must(t, o.WaitForText("First block ID#", 60))
	o.RunSeconds(2)
	must(t, pictures.Screenshot(o, "concentration"))
	game := pictures.Record(o).Faster(4)
	game.Capture(50)
	playConcentration(t, o, game)
	waitRecording(t, o, game, "matched all the blocks", 30)
	game.Run(120, 6)
	must(t, pictures.SaveRecording(game, "game", 300))
	must(t, pictures.Screenshot(o, "matched"))

	// Back to the games, to the main prompt, and goodbye: the BBS hangs up
	must(t, bbsType(o, nil, "\n"))
	must(t, o.WaitForText("Enter ID number, * or Q", 30))
	o.RunSeconds(1)
	must(t, bbsType(o, nil, "Q\n"))
	must(t, o.WaitForText("[MAIN] (?=Help):", 30))
	o.RunSeconds(1)
	bye := pictures.Record(o)
	bye.Capture(50)
	must(t, bbsType(o, bye, "BYE\n"))
	waitRecording(t, o, bye, "Are you sure? Y/N", 30)
	bye.Run(30, 6)
	must(t, bbsType(o, bye, "Y\n"))
	waitRecording(t, o, bye, "NO CARRIER", 60)
	bye.Run(60, 6)
	must(t, pictures.SaveRecording(bye, "bye", 300))
}

/*
bbsType types a text slower than the operator does, a key every quarter of
a second: the BBS loses keys that come faster. The recording, if any, takes
a picture after each.
*/
func bbsType(o *operator.Operator, r *album.Recording, text string) error {
	for _, c := range text {
		if r != nil {
			if err := r.Type(string(c)); err != nil {
				return err
			}
			r.Run(6, 6)
			continue
		}
		if err := o.Type(string(c)); err != nil {
			return err
		}
		o.Run(6)
	}
	return nil
}

// choose ticks the line chosen in a list of ProTERM and saves it
func choose(t *testing.T, o *operator.Operator) {
	t.Helper()
	o.RunSeconds(0.5)
	must(t, o.Key("Return"))
	o.RunSeconds(1)
	must(t, o.Key("Return"))
}

// waitRecording runs a recording until a text is on the screen
func waitRecording(t *testing.T, o *operator.Operator, r *album.Recording, text string, seconds float64) {
	t.Helper()
	recordUntil(t, o, r, seconds, func() bool { return o.HasText(text) })
}

// recordUntil runs a recording until something has happened
func recordUntil(t *testing.T, o *operator.Operator, r *album.Recording, seconds float64, done func() bool) {
	t.Helper()
	for limit := o.Frames() + uint64(seconds*operator.FramesPerSecond); !done(); {
		if o.Frames() > limit {
			t.Fatalf("it did not come:\n%v", o.Text())
		}
		r.Run(6, 6)
	}
}

// concentrationHidden is an object not turned over, as the text of the
// screen reads its MouseText
const concentrationHidden = "VWVW"

/*
concentrationBoard reads the objects of the board of Concentration+ from the
text of the screen: four rows of seven, each object under its number, four
characters of MouseText wide and two lines high
*/
func concentrationBoard(o *operator.Operator) map[int]string {
	lines := strings.Split(o.Text(), "\n")
	board := map[int]string{}
	for i, line := range lines {
		if !strings.HasPrefix(line, "|  1 ") && !strings.HasPrefix(line, "|  8 ") &&
			!strings.HasPrefix(line, "| 15 ") && !strings.HasPrefix(line, "| 22 ") || i+1 >= len(lines) {
			continue
		}
		var first int
		fmt.Sscan(line[1:], &first)
		row := lines[i+1]
		for c := range 7 {
			at := 1 + 5*c
			if at+4 <= len(row) {
				board[first+c] = row[at : at+4]
			}
		}
	}
	return board
}

// concentrationPrompt is a prompt of Concentration+ waiting for an ID, with
// no digit typed after it yet, only the cursor
var concentrationPrompt = regexp.MustCompile(`(?m)^(First|Second) block ID# \(1-28\)[^:]*:[^0-9\n]*$`)

// prompting tells if the game waits for the first or the second object of a
// move
func prompting(o *operator.Operator, which string) bool {
	m := concentrationPrompt.FindStringSubmatch(o.Text())
	return m != nil && m[1] == which
}

/*
playConcentration plays Concentration+ as one with a good memory: a pair
already seen is taken at once; if not, an object not seen yet is turned
over, then its pair if it was seen, or another not seen yet, and both are
remembered
*/
func playConcentration(t *testing.T, o *operator.Operator, r *album.Recording) {
	t.Helper()
	seen := map[int]string{}
	matched := map[int]bool{}
	// Some pairs look the same in MouseText, as the game tells them apart
	// by their colour: two that look the same and didn't match
	failed := map[[2]int]bool{}
	var last [2]int
	alike := func(a, b int) bool {
		return a != b && !matched[a] && !matched[b] && seen[a] != "" && seen[a] == seen[b] &&
			!failed[[2]int{min(a, b), max(a, b)}]
	}
	// look runs until an object is turned over, and reads it, or gives up
	look := func(id int) string {
		for limit := o.Frames() + 4*operator.FramesPerSecond; o.Frames() < limit; {
			// Whole, drawn over the hidden one: a pattern of two characters
			// twice
			shown := concentrationBoard(o)[id]
			if len(shown) == 4 && shown != concentrationHidden && shown[:2] == shown[2:] {
				return shown
			}
			r.Run(2, 2)
		}
		return ""
	}
	// pick types the ID of an object at its prompt and reads the object
	// turned over, or "" if the game didn't take it. The BBS loses a key
	// now and then: the ID is typed until it shows after the prompt.
	pick := func(id int, which string) string {
		recordUntil(t, o, r, 30, func() bool { return prompting(o, which) })
		r.Run(30, 6)
		typed := regexp.MustCompile(fmt.Sprintf(`(?m)^%v block ID# .*:%d\D*$`, which, id))
		must(t, bbsType(o, r, fmt.Sprint(id)))
		recordUntil(t, o, r, 10, func() bool { return typed.MatchString(o.Text()) })
		must(t, bbsType(o, r, "\n"))
		return look(id)
	}
	for moves := 0; len(matched) < 28; moves++ {
		if moves > 50 {
			t.Fatalf("the game went on for too long:\n%v", o.Text())
		}
		// At the prompt of the first object, the objects shown are the
		// pairs found: the two of a move that didn't match are turned back
		// before it. At the end, the board all turned over and no prompt.
		recordUntil(t, o, r, 30, func() bool {
			if prompting(o, "First") {
				return true
			}
			for _, object := range concentrationBoard(o) {
				if object == concentrationHidden {
					return false
				}
			}
			return true
		})
		for id, object := range concentrationBoard(o) {
			if object != concentrationHidden {
				matched[id] = true
			}
		}
		if len(matched) == 28 {
			break
		}
		if last[0] != 0 && !matched[last[0]] {
			failed[[2]int{min(last[0], last[1]), max(last[0], last[1])}] = true
		}
		last = [2]int{}

		// A pair already seen, or one not seen yet and its pair if it was
		// seen, or another not seen yet
		first, second := 0, 0
		for a := 1; a <= 28 && first == 0; a++ {
			for b := a + 1; b <= 28; b++ {
				if alike(a, b) {
					first, second = a, b
					break
				}
			}
		}
		if first == 0 {
			for a := 1; a <= 28 && first == 0; a++ {
				if !matched[a] && seen[a] == "" {
					first = a
				}
			}
		}
		shown := pick(first, "First")
		recordUntil(t, o, r, 10, func() bool { return prompting(o, "Second") || prompting(o, "First") })
		if shown == "" || !prompting(o, "Second") {
			// The game didn't take it: the move again
			continue
		}
		seen[first] = shown
		for b := 1; b <= 28 && second == 0; b++ {
			if alike(first, b) {
				second = b
			}
		}
		for b := 1; b <= 28 && second == 0; b++ {
			if b != first && !matched[b] && seen[b] == "" {
				second = b
			}
		}
		if shown := pick(second, "Second"); shown != "" {
			seen[second] = shown
			last = [2]int{first, second}
		}
	}
}
