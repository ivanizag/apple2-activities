package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// missionAsteroidMachine is the machine of the guide, as its command line of
// izapple2: an Apple ][+ with Mission: Asteroid in drive 1
const missionAsteroidMachine = `izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s6 'diskii,disk1=disks/Mission Asteroid (4am and san inc crack).dsk'`

// missionAsteroidShort is the model 2plus of izapple2, an Apple ][+ with more cards,
// with the disk of the game
const missionAsteroidShort = `izapple2 -model 2plus 'disks/Mission Asteroid (4am and san inc crack).dsk'`

// missionAsteroid is the game
var missionAsteroid = adventure{
	short:       missionAsteroidShort,
	name:        "missionasteroid",
	walkthrough: "../guides/listings/missionasteroid.txt",
	disk:        func([]string) (string, string, bool) { return "", "", false },
	ending:      "THE ASTEROID HAS EXPLODED",
}

/*
missionAsteroidScreenshots plays Mission: Asteroid from the start to the end, on an Apple ][+,
as the walkthrough says, a picture of the screen for each command, and
writes the walkthrough of the page with the pictures.
*/
func missionAsteroidScreenshots(t *testing.T) {
	missionAsteroid.screenshots(t, missionAsteroidMachine, missionAsteroidBegin)
}

// missionAsteroidBegin takes the game to its first prompt
func missionAsteroidBegin(t *testing.T, o *operator.Operator, pictures *album.Album) {
	// The game starts at its first place
}

// TestMissionAsteroidWalkthrough checks that the page has the walkthrough
func TestMissionAsteroidWalkthrough(t *testing.T) {
	missionAsteroid.checkPage(t)
}
