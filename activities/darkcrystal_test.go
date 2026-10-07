package activities

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// darkCrystalMachine is the machine of the guide, as its command line of
// izapple2: an Apple ][+ with The Dark Crystal in drive 1
const darkCrystalMachine = `izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s6 'diskii,disk1=disks/The Dark Crystal (4am and san inc crack) disk 1A.dsk'`

// darkCrystalDisk is the game asking for a side of its two disks
var darkCrystalDisk = regexp.MustCompile(`INSERT DISK #(\d), SIDE "([AB])"`)

// darkCrystal is the game, on the four sides of two disks
var darkCrystal = adventure{
	name:        "darkcrystal",
	walkthrough: "../guides/listings/darkcrystal.txt",
	disk: func(lines []string) (string, string, bool) {
		m := darkCrystalDisk.FindStringSubmatch(strings.Join(lines, " "))
		if m == nil || strings.TrimSpace(lines[len(lines)-1]) != "AND PRESS RETURN." {
			return "", "", false
		}
		return fmt.Sprintf("Disk %v, side %v", m[1], m[2]),
			fmt.Sprintf("The Dark Crystal (4am and san inc crack) disk %v%v.dsk", m[1], m[2]), true
	},
	ending: "THANKS FOR PLAYING",
	prompt: "----> ENTER COMMAND",
	waits: []adventureWait{
		{0x64f9, 0x6500, adventureMore},
	},
}

/*
darkCrystalScreenshots plays The Dark Crystal from the start to the end, on an Apple ][+,
as the walkthrough says, a picture of the screen for each command, and
writes the walkthrough of the page with the pictures.
*/
func darkCrystalScreenshots(t *testing.T) {
	pictures := newAlbum("darkcrystal", album.ColorWhiteText)
	o := start(t, darkCrystalMachine, nil)

	// The game asks for the side of its disks it starts from

	darkCrystal.play(t, o, pictures)
}

// TestDarkCrystalWalkthrough checks that the page has the walkthrough
func TestDarkCrystalWalkthrough(t *testing.T) {
	darkCrystal.checkPage(t)
}
