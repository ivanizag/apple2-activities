package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// cranstonManorMachine is the machine of the guide, as its command line of
// izapple2: an Apple ][+ with Cranston Manor in drive 1
const cranstonManorMachine = `izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s6 'diskii,disk1=disks/Cranston Manor (4am and san inc crack).dsk'`

// cranstonManorShort is the model 2plus of izapple2, an Apple ][+ with more cards,
// with the disk of the game
const cranstonManorShort = `izapple2 -model 2plus 'disks/Cranston Manor (4am and san inc crack).dsk'`

// cranstonManor is the game
var cranstonManor = adventure{
	short:       cranstonManorShort,
	name:        "cranston",
	walkthrough: "../guides/listings/cranston.txt",
	disk:        func([]string) (string, string, bool) { return "", "", false },
	ending:      "LEVEL 3",
}

/*
cranstonManorScreenshots plays Cranston Manor from the start to the end, on an Apple ][+,
as the walkthrough says, a picture of the screen for each command, and
writes the walkthrough of the page with the pictures.
*/
func cranstonManorScreenshots(t *testing.T) {
	cranstonManor.screenshots(t, cranstonManorMachine, cranstonManorBegin)
}

// cranstonManorBegin takes the game to its first prompt
func cranstonManorBegin(t *testing.T, o *operator.Operator, pictures *album.Album) {
	// The game starts at its first place
}

// TestCranstonManorWalkthrough checks that the page has the walkthrough
func TestCranstonManorWalkthrough(t *testing.T) {
	cranstonManor.checkPage(t)
}
