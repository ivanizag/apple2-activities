package activities

import (
	"regexp"
	"strconv"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// game2048Listing is the program of the page, the game 2048 in Pascal
const game2048Listing = "../guides/listings/2048.pas"

/*
pascal2048Screenshots is a whole game written in Apple Pascal 1.3 on an
enhanced Apple //e, 2048: the editor set not to indent by itself, the
program of the page typed, compiled and run, and played, a key at a time,
as a player who keeps the big tiles in a corner.
*/
func pascal2048Screenshots(t *testing.T) {
	pictures := newAlbum("pascal-2048", album.Green)
	o := start(t, pascalMachine, nil)
	must(t, o.WaitForText("[1.3]", 120))
	must(t, o.WaitForKeyboard(10))

	// A new file in the editor, and its environment: no automatic indent
	must(t, o.Type("E"))
	must(t, o.WaitForText("Edit what file?", 30))
	must(t, o.Key("Return"))
	must(t, o.WaitForText(">Edit:", 30))
	must(t, o.Type("SE"))
	must(t, o.WaitForText("A(uto indent", 10))
	must(t, o.Type("AF"))
	must(t, o.WaitForText("A(uto indent   False", 10))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "environment"))
	must(t, o.Key("Ctrl+C"))
	must(t, o.WaitForText(">Edit:", 10))

	// The program typed in, and kept in the work file
	must(t, o.Type("I"))
	o.Run(30)
	must(t, typeText(o, game2048Listing))
	must(t, o.Key("Ctrl+C"))
	o.Run(60)
	must(t, pictures.Screenshot(o, "typed"))
	must(t, o.Type("Q"))
	must(t, o.WaitForText("C(hange", 30))
	must(t, o.Type("U"))
	must(t, o.WaitForText("bytes long", 60))
	must(t, o.Type("E"))
	must(t, o.WaitForText("Command:", 30))

	// Compiled, the compiler showing the procedures as it goes
	must(t, o.Type("R"))
	must(t, o.WaitForText("Listing file", 60))
	must(t, o.Key("Return"))
	must(t, o.WaitForText("GAME2048 [", 300))
	must(t, pictures.Screenshot(o, "compiling"))

	// Running: the board, and the first moves recorded
	must(t, o.WaitForText("MOVES    0", 120))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "board"))
	first := pictures.Record(o)
	first.Capture(50)
	for range 30 {
		playMove(t, o, first)
	}
	first.Run(60, 6)
	must(t, pictures.SaveRecording(first, "playing", 300))

	// More moves, until the game ends, or after 400 moves Q ends it
	for range 370 {
		if gameEnded(o) {
			break
		}
		playMove(t, o, nil)
	}
	if !gameEnded(o) {
		must(t, o.Type("Q"))
	}
	must(t, o.WaitForText("PRESS A KEY", 60))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "end"))
	must(t, o.Key("Space"))
	must(t, o.WaitForText("Command:", 30))
}

// movesText is the count of moves of the game, under the board
var movesText = regexp.MustCompile(`MOVES +([0-9]+)`)

// gameMoves is the count of moves the game shows
func gameMoves(o *operator.Operator) int {
	m := movesText.FindStringSubmatch(o.Text())
	if m == nil {
		return -1
	}
	n, _ := strconv.Atoi(m[1])
	return n
}

/*
playMove makes one move as a player who keeps the big tiles at the bottom
left: down if the tiles slide, else left, else right, else up. Each key is
waited out until the screen is still, the board drawn again or not, and
recorded when there is a recording.
*/
func playMove(t *testing.T, o *operator.Operator, r *album.Recording) {
	t.Helper()
	before := gameMoves(o)
	for _, key := range "KJLI" {
		if gameEnded(o) {
			return
		}
		must(t, o.Type(string(key)))
		still, last := 0, o.Text()
		for range 60 {
			if r != nil {
				r.Run(6, 6)
			} else {
				o.Run(6)
			}
			if text := o.Text(); text != last {
				still, last = 0, text
			} else if still++; still == 5 {
				break
			}
		}
		if gameMoves(o) != before {
			return
		}
	}
}

// gameEnded tells whether the game says it has ended
func gameEnded(o *operator.Operator) bool {
	return o.HasText("PRESS A KEY")
}
