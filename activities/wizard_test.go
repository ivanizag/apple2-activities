package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// wizardMachine is the machine of the guide, as its command line of
// izapple2: an Apple ][+ with Wizard and the Princess in drive 1
const wizardMachine = `izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s6 'diskii,disk1=disks/The Wizard and the Princess (1980-On-Line Systems).nib'`

// wizardShort is the model 2plus of izapple2, an Apple ][+ with more cards,
// with the disk of the game
const wizardShort = `izapple2 -model 2plus 'disks/The Wizard and the Princess (1980-On-Line Systems).nib'`

// wizard is the game
var wizard = adventure{
	short:       wizardShort,
	name:        "wizard",
	walkthrough: "../guides/listings/wizard.txt",
	disk:        func([]string) (string, string, bool) { return "", "", false },
	ending:      "JUNIOR-MASTER ADVENTURER",
}

/*
wizardScreenshots plays Wizard and the Princess from the start to the end, on an Apple ][+,
as the walkthrough says, a picture of the screen for each command, and
writes the walkthrough of the page with the pictures.
*/
func wizardScreenshots(t *testing.T) {
	wizard.screenshots(t, wizardMachine, wizardBegin)
}

// wizardBegin takes the game to its first prompt
func wizardBegin(t *testing.T, o *operator.Operator, pictures *album.Album) {
	// The game starts at its first place
}

// TestWizardWalkthrough checks that the page has the walkthrough
func TestWizardWalkthrough(t *testing.T) {
	wizard.checkPage(t)
}
