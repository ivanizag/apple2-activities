package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// mysteryHouseMachine is the machine of the guide, as its command line of
// izapple2: an Apple ][+ with Mystery House in drive 1
const mysteryHouseMachine = `izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s6 'diskii,disk1=disks/Mystery House (4am crack).dsk'`

// mysteryHouseShort is the model 2plus of izapple2, an Apple ][+ with more cards,
// with the disk of the game
const mysteryHouseShort = `izapple2 -model 2plus 'disks/Mystery House (4am crack).dsk'`

// mysteryHouse is the game, on one side of a disk
var mysteryHouse = adventure{
	short:       mysteryHouseShort,
	name:        "mysteryhouse",
	walkthrough: "../guides/listings/mysteryhouse.txt",
	disk:        func([]string) (string, string, bool) { return "", "", false },
	ending:      "YOU HAVE BEATEN",
}

/*
mysteryHouseScreenshots plays Mystery House from the start to the end, on an
Apple ][+, as the walkthrough says, a picture of the screen for each command,
and writes the walkthrough of the page with the pictures.
*/
func mysteryHouseScreenshots(t *testing.T) {
	mysteryHouse.screenshots(t, mysteryHouseMachine, mysteryHouseBegin)
}

// mysteryHouseBegin takes the game to its first prompt
func mysteryHouseBegin(t *testing.T, o *operator.Operator, pictures *album.Album) {
	// The title, and G for the game
	must(t, o.WaitForText("ENTER G FOR GAME", 60))
	must(t, o.WaitForKeyboard(60))
	must(t, pictures.Screenshot(o, "title"))
	must(t, o.Type("G\n"))
}

// TestMysteryHouseWalkthrough checks that the page has the walkthrough
func TestMysteryHouseWalkthrough(t *testing.T) {
	mysteryHouse.checkPage(t)
}
