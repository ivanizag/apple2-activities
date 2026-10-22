package activities

import (
	"os"
	"strings"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// lodeRunnerMachine is the machine of the guide, as its command line of izapple2
const lodeRunnerMachine = `izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s6 'diskii,disk1=disks/Lode Runner.woz'`

// lodeRunnerLevel is the level of the guide, a row of the board a line, the
// key of each shape in its place and a dot where there is nothing
const lodeRunnerLevel = "../guides/listings/lode-runner.txt"

/*
lodeRunnerScreenshots is Lode Runner of 1983 on an Apple ][+, from its
original disk: the title, the demonstration that plays itself, and a game
started and played from the keyboard, two pieces of gold taken before a
guard catches the runner. Then the board editor: a data disk initialized, a
level of our own drawn on it shape by shape, saved, and played to the end.

	cp disks/blank.dsk my-disk.dsk
*/
func lodeRunnerScreenshots(t *testing.T) {
	pictures := newAlbum("lode-runner", album.Color)
	myDisk := disk(t, "blank.dsk")
	o := start(t, lodeRunnerMachine, nil)

	// The title
	o.Run(16 * 60)
	must(t, pictures.Screenshot(o, "title"))

	// The demonstration, recorded from the level appearing
	o.Run(4 * 60)
	demo := pictures.Record(o)
	demo.Capture(6)
	demo.Run(20*60, 6)
	must(t, pictures.SaveRecording(demo, "demo", 100))

	// A game started with Space, the first level drawn
	must(t, o.Key("Space"))
	o.Run(10 * 60)
	must(t, pictures.Screenshot(o, "game"))

	// Played from the keyboard: Control-K, then a key sets the runner going
	// until the next one. Right to the gold, up the ladder, and left to the
	// next gold, until a guard catches him.
	must(t, o.Key("Ctrl+K"))
	o.Run(30)
	play := pictures.Record(o)
	play.Capture(10)
	must(t, o.Type("L"))
	play.Run(210, 6)
	must(t, o.Type("I"))
	play.Run(72, 6)
	must(t, o.Type("J"))
	play.Run(10*60, 6)
	must(t, pictures.SaveRecording(play, "play", 200))

	// Started again, not to write the high scores on the disk of the game,
	// and the board editor, Control-E, at the title
	o = start(t, lodeRunnerMachine, nil)
	o.Run(16 * 60)
	must(t, o.Key("Ctrl+E"))
	must(t, o.WaitForKeyboard(10))
	must(t, pictures.Screenshot(o, "editor"))

	// A data disk: I, my disk put in for the one of the game, and Y
	must(t, o.Type("I"))
	must(t, o.WaitForKeyboard(10))
	must(t, pictures.Screenshot(o, "initialize"))
	must(t, o.InsertDisk(0, myDisk))
	must(t, o.Type("Y"))
	must(t, o.WaitForKeyboard(60))

	// Level 001 of the data disk, empty, and drawn: a key moves the cursor
	// and a digit puts a shape where it is, a row at a time
	must(t, o.Type("E"))
	must(t, o.WaitForKeyboard(10))
	must(t, o.Type("001\n"))
	must(t, o.WaitForKeyboard(10))
	waitForCursor(t, o, pictures)
	must(t, pictures.Screenshot(o, "empty"))
	level, err := os.ReadFile(lodeRunnerLevel)
	must(t, err)
	drawing := pictures.Record(o)
	drawing.Capture(50)
	for _, row := range strings.Split(strings.TrimSpace(string(level)), "\n") {
		must(t, drawRow(o, row))
		drawing.Capture(25)
	}
	o.Run(30)
	drawing.Capture(25)
	must(t, pictures.SaveRecording(drawing, "drawing", 300))

	// Saved with Control-S, and back to the editor with Control-Q
	must(t, o.Key("Ctrl+S"))
	must(t, o.WaitForKeyboard(30))
	must(t, o.Key("Ctrl+Q"))
	must(t, o.WaitForKeyboard(30))

	// Played, P and its number: the five pieces of gold taken, the ladder
	// that appears climbed to the top, and the level won
	must(t, o.Type("P"))
	must(t, o.WaitForKeyboard(10))
	must(t, o.Type("001\n"))
	o.Run(10 * 60)
	must(t, o.Key("Ctrl+K"))
	o.Run(30)
	mine := pictures.Record(o)
	mine.Capture(10)
	run := func(key string, done func(column, row int) bool) {
		t.Helper()
		must(t, o.Type(key))
		limit := o.Frames() + 20*60
		for !done(runnerAt(o)) {
			if o.Frames() > limit {
				column, row := runnerAt(o)
				t.Fatalf("the runner stopped at %v,%v after %v", column, row, key)
			}
			mine.Run(2, 6)
		}
	}
	across := func(to int) func(column, row int) bool {
		return func(column, _ int) bool { return column == to && runnerCentred(o) }
	}
	upOrDown := func(to int) func(column, row int) bool {
		return func(_, row int) bool { return row == to }
	}
	run("L", across(21))
	run("J", across(15))
	run("I", upOrDown(9))
	run("J", across(6))
	run("I", upOrDown(3))
	run("J", across(2))
	run("L", across(15))
	run("K", upOrDown(9))
	run("L", across(25))
	run("J", across(22))
	run("I", upOrDown(0))
	mine.Run(3*60, 6)
	must(t, pictures.SaveRecording(mine, "mine", 300))
}

// waitForCursor waits for the cursor of the board editor, which blinks, to be
// drawn at the top left of the board
func waitForCursor(t *testing.T, o *operator.Operator, pictures *album.Album) {
	t.Helper()
	limit := o.Frames() + 5*60
	for !promptDrawn(pictures.Screen(o)) {
		if o.Frames() > limit {
			t.Fatal("the cursor of the editor did not show")
		}
		o.Run(2)
	}
}

/*
drawRow draws a row of a level in the board editor, from the cursor at its
left: the digit of each shape, and K to the right after each but the last,
then M down and J back to the left
*/
func drawRow(o *operator.Operator, row string) error {
	for i, shape := range row {
		if shape != '.' {
			if err := o.Type(string(shape)); err != nil {
				return err
			}
		}
		if i < len(row)-1 {
			if err := o.Type("K"); err != nil {
				return err
			}
		}
	}
	return o.Type("M" + strings.Repeat("J", len(row)-1))
}

// runnerAt is the column and row of the runner, which the game keeps at $00
// and $01
func runnerAt(o *operator.Operator) (column int, row int) {
	return int(o.Apple2().Peek(0x00)), int(o.Apple2().Peek(0x01))
}

/*
runnerCentred tells if the runner is near the middle of his column, where he
takes the gold and can climb: the game keeps where he is in it at $02, 0 to
4, 2 in the middle, and moves him more than one a step
*/
func runnerCentred(o *operator.Operator) bool {
	at := o.Apple2().Peek(0x02)
	return at >= 1 && at <= 3
}
