package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// akalabethMachine is the machine of the guide, as its command line of
// izapple2
const akalabethMachine = `izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s6 diskii,disk1=disks/Akalabeth_original.dsk`

// The ways of the page, a letter a move: N, S, W and E outside, the keys
// Return, /, ← and →; in the dungeons the same keys are forward, turn
// around, turn left and turn right, and A R attacks with the rapier, X
// climbs a ladder
const (
	akalabethToCastle  = "NNWNNNNNNNWWWWWWNNNW"
	akalabethToDungeon = "ESSS"
	akalabethBack      = "NNNW"
	// The first level, to the ladder down, a skeleton and a thief met
	akalabethLevel1 = "NNEEARARSARENNNNNN"
	// The second level, the giant rat, and the ladder up
	akalabethLevel2 = "SNNENARARSNWNN"
	// The first level again, to the ladder up at the entrance
	akalabethOut = "SNNNNNNWARARSARSNN"
)

/*
akalabethScreenshots is Akalabeth, Richard Garriott's game of 1980, played
with the lucky number 5: the title, the qualities and the shop, the walk to
the castle of Lord British and his quest, a giant rat, into a dungeon, down
a ladder to the second level, where the rat is, killed, and back to the
castle, the quest done.

The game is all in Applesoft, and the lucky number starts its random
numbers: the same keys make the same world, the same dungeons and the same
fights. Applesoft looks at the keyboard after every statement, for
Control-C, so the generator does not wait for the keyboard but for the line
of the program that waits for a key, from the line Applesoft is running, at
$75.
*/
func akalabethScreenshots(t *testing.T) {
	pictures := newAlbum("akalabeth", album.Color)
	o := start(t, akalabethMachine, nil)

	// The title, then the picture of the title, drawn as the game loads
	akalabethWait(t, o, 4)
	must(t, o.WaitForText("CONTINUE)", 10))
	must(t, pictures.Screenshot(o, "welcome"))
	must(t, o.Key("Space"))
	title := pictures.Record(o)
	title.Capture(50)
	for o.InTextMode() {
		o.Run(4)
	}
	// Until line 13 of the title, AKA0, which wipes the picture off
	for !o.InTextMode() && akalabethLine(o) != 13 {
		title.Capture(7)
		o.Run(4)
	}
	must(t, pictures.SaveRecording(title, "title", 300))

	// A character: the lucky number, the level, and the qualities it gives
	akalabethWait(t, o, 60000)
	must(t, o.TypeLines("5"))
	akalabethWait(t, o, 60005)
	must(t, o.TypeLines("1"))
	akalabethWait(t, o, 60060)
	must(t, o.WaitForText("QUALITIES?", 10))
	must(t, pictures.Screenshot(o, "qualities"))
	must(t, o.Type("Y"))
	akalabethWait(t, o, 60061)
	must(t, o.Type("F"))

	// The shop: a rapier, and food with the rest of the gold
	akalabethWait(t, o, 60210)
	for _, item := range "RFFFFFFFF" {
		must(t, o.Type(string(item)))
		akalabethWait(t, o, 60210)
	}
	must(t, pictures.Screenshot(o, "shop"))
	must(t, o.Type("Q"))

	// The world, from the town, and the walk to the castle
	akalabethWait(t, o, 1001)
	must(t, pictures.Screenshot(o, "town"))
	walk := pictures.Record(o)
	walk.Capture(50)
	akalabethMoves(t, o, walk, akalabethToCastle)
	walk.Capture(200)
	must(t, pictures.SaveRecording(walk, "walk", 300))

	// Lord British, and his quest
	must(t, o.Type("X"))
	akalabethWait(t, o, 7030)
	must(t, o.TypeLines("ADA"))
	akalabethWait(t, o, 7040)
	must(t, o.Type("Y"))
	akalabethWait(t, o, 7070)
	must(t, pictures.Screenshot(o, "quest"))
	must(t, o.Key("Space"))

	// The dungeon, its first level
	akalabethWait(t, o, 1001)
	akalabethMoves(t, o, nil, akalabethToDungeon)
	must(t, o.Type("X"))
	akalabethWait(t, o, 1001)
	must(t, pictures.Screenshot(o, "dungeon"))
	akalabethMoves(t, o, nil, akalabethLevel1)
	must(t, pictures.Screenshot(o, "ladder"))

	// Down the ladder, the giant rat, and back to the ladder up
	must(t, o.Type("X"))
	akalabethWait(t, o, 1001)
	rat := pictures.Record(o)
	rat.Capture(100)
	akalabethMoves(t, o, rat, akalabethLevel2[:9])
	rat.Capture(300)
	must(t, pictures.SaveRecording(rat, "rat", 300))
	must(t, pictures.Screenshot(o, "killed"))
	akalabethMoves(t, o, nil, akalabethLevel2[9:])

	// Up, and out of the dungeon
	must(t, o.Type("X"))
	akalabethWait(t, o, 1001)
	akalabethMoves(t, o, nil, akalabethOut)
	must(t, o.Type("X"))
	akalabethWait(t, o, 1001)
	must(t, pictures.Screenshot(o, "out"))

	// Back to Lord British: the quest done, and the next one
	akalabethMoves(t, o, nil, akalabethBack)
	must(t, o.Type("X"))
	akalabethWait(t, o, 7070)
	must(t, pictures.Screenshot(o, "done"))

}

// akalabethLine is the line of the program Applesoft is running, CURLIN
func akalabethLine(o *operator.Operator) int {
	return int(o.Apple2().Peek(0x75)) | int(o.Apple2().Peek(0x76))<<8
}

// akalabethWait runs the machine until the program is at a line, one that
// waits for a key, two minutes at most
func akalabethWait(t *testing.T, o *operator.Operator, line int) {
	t.Helper()
	if !o.WaitUntil(120, func() bool { return akalabethLine(o) == line }) {
		t.Fatalf("the program did not get to line %v, it is at %v:\n%v", line, akalabethLine(o), o.Text())
	}
}

// akalabethMoves presses the keys of a way, each when the game asks for a
// command, recording the screen if there is a recording
func akalabethMoves(t *testing.T, o *operator.Operator, r *album.Recording, way string) {
	t.Helper()
	keys := map[rune]string{'N': "Return", 'S': "/", 'W': "Left", 'E': "Right", 'A': "A", 'R': "R", 'X': "X"}
	for _, move := range way {
		key := keys[move]
		if key == "R" {
			// The weapon, asked for after A
			akalabethWait(t, o, 1650)
		} else {
			akalabethWait(t, o, 1001)
		}
		if key == "Return" || key == "Left" || key == "Right" {
			must(t, o.Key(key))
		} else {
			must(t, o.Type(key))
		}
		for r != nil && akalabethLine(o) != 1001 && akalabethLine(o) != 1650 {
			r.Run(4, 4)
		}
	}
	akalabethWait(t, o, 1001)
}
