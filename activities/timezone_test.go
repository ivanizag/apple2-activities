package activities

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// timeZoneMachine is the machine of the guide, as its command line of
// izapple2: an Apple ][+ with side A of Time Zone in drive 1
const timeZoneMachine = `izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s6 'diskii,disk1=disks/Time Zone (4am and san inc crack) disk A.dsk'`

// timeZoneDisk is the game asking for a side of its six disks, 1A to 6L; the
// letter names the file
var timeZoneDisk = regexp.MustCompile(`INSERT DISK NUMBER (\d([A-L]))`)

// timeZone is the game: it pauses at $4780 in the time machine before its
// prompt, reading the keyboard
var timeZone = adventure{
	name:        "timezone",
	walkthrough: "../guides/listings/timezone.txt",
	waits: []adventureWait{
		{0x4780, 0x47e1, adventureLine},
	},
	disk: func(lines []string) (string, string, bool) {
		m := timeZoneDisk.FindStringSubmatch(strings.Join(lines, " "))
		if m == nil || strings.TrimSpace(lines[len(lines)-1]) != "AND PRESS RETURN." {
			return "", "", false
		}
		return "Side " + m[1], fmt.Sprintf("Time Zone (4am and san inc crack) disk %v.dsk", m[2]), true
	},
	ending: "ULTIMATE ADVENTURER",
}

/*
timeZoneScreenshots plays Time Zone from the start to the end, on an Apple
][+, as the walkthrough says, a picture of the screen for each command, and
one for each page of a longer answer. It changes the disks when the game asks
for them, and writes the walkthrough of the page with the pictures.
*/
func timeZoneScreenshots(t *testing.T) {
	pictures := newAlbum("timezone", album.ColorWhiteText)
	o := start(t, timeZoneMachine, nil)

	// The title, until a key, and the menu
	must(t, o.WaitForKeyboard(60))
	must(t, pictures.Screenshot(o, "title"))
	must(t, o.Type("\n"))
	must(t, o.WaitForText("WHICH WOULD YOU LIKE?", 60))
	must(t, o.WaitForKeyboard(60))
	must(t, o.Type("1"))
	must(t, pictures.Screenshot(o, "menu"))
	must(t, o.Type("\n"))
	must(t, o.WaitForText("INSERT DISK", 60))

	timeZone.play(t, o, pictures)
}

// TestTimeZoneWalkthrough checks that the page has the walkthrough
func TestTimeZoneWalkthrough(t *testing.T) {
	timeZone.checkPage(t)
}
