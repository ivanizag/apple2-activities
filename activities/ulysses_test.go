package activities

import (
	"strings"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// ulyssesMachine is the machine of the guide, as its command line of
// izapple2: an Apple ][+ with Ulysses and the Golden Fleece in drive 1
const ulyssesMachine = `izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s6 'diskii,disk1=disks/Ulysses and the Golden Fleece v1.1 (4am crack) side A.dsk'`

// ulyssesShort is the model 2plus of izapple2, an Apple ][+ with more cards,
// with the disk of the game
const ulyssesShort = `izapple2 -model 2plus 'disks/Ulysses and the Golden Fleece v1.1 (4am crack) side A.dsk'`

// ulyssesSide is the side of the disk in the drive, A or B
var ulyssesSide = "A"

// ulysses is the game, on the two sides of a disk: it asks to flip it over,
// and waits for the key at $26DB or $26FB
var ulysses = adventure{
	short:       ulyssesShort,
	name:        "ulysses",
	walkthrough: "../guides/listings/ulysses.txt",
	waits: []adventureWait{
		{0x26db, 0x26df, adventureLine},
		{0x26fb, 0x26ff, adventureLine},
	},
	disk: func(lines []string) (string, string, bool) {
		if strings.TrimSpace(lines[len(lines)-1]) != "PLEASE FLIP DISK OVER AND PRESS ANY KEY" {
			return "", "", false
		}
		if ulyssesSide == "A" {
			ulyssesSide = "B"
		} else {
			ulyssesSide = "A"
		}
		return "Side " + ulyssesSide, "Ulysses and the Golden Fleece v1.1 (4am crack) side " + ulyssesSide + ".dsk", true
	},
	ending: "THE KING DELIGHTEDLY TAKES THE FLEECE",
}

/*
ulyssesScreenshots plays Ulysses and the Golden Fleece from the start to the end, on an Apple ][+,
as the walkthrough says, a picture of the screen for each command, and
writes the walkthrough of the page with the pictures.
*/
func ulyssesScreenshots(t *testing.T) {
	ulysses.screenshots(t, ulyssesMachine, ulyssesBegin)
}

// ulyssesBegin takes the game to its first prompt
func ulyssesBegin(t *testing.T, o *operator.Operator, pictures *album.Album) {
	ulyssesSide = "A"

	// The game asks for the other side as it starts
}

// TestUlyssesWalkthrough checks that the page has the walkthrough
func TestUlyssesWalkthrough(t *testing.T) {
	ulysses.checkPage(t)
}
