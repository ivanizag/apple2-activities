/*
Package activities makes the pictures of the guides, by taking a machine
through what each guide tells its reader to do. It is not a test: it does
nothing unless asked for with A2_ACTIVITIES, and then writes the pictures of
each guide in guides/images, over the ones there.

	A2_ACTIVITIES=1 go test -count=1 -run 'TestActivities/switch-on' ./activities
*/
package activities

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// activities are the guides, the machines their pages give, and the
// generators of their pictures
var activities = []struct {
	guide    string
	machines []string
	pictures func(t *testing.T)
}{
	{"switch-on", []string{switchOnMachine}, switchOnScreenshots},
	{"dos33", []string{dos33Machine, dos33OwnDisk}, dos33Screenshots},
	{"paddle-game", []string{paddleMachine}, paddleScreenshots},
	{"desktop", []string{deskTopMachine}, deskTopScreenshots},
	{"apple-ii", []string{appleIIMachine}, appleIIScreenshots},
	{"pascal", []string{pascalMachine}, pascalScreenshots},
	{"cpm", []string{cpmMachine}, cpmScreenshots},
	{"lode-runner", []string{lodeRunnerMachine}, lodeRunnerScreenshots},
	{"mockingboard", []string{mockingboardMachine}, mockingboardScreenshots},
	{"apple-iie", []string{appleIIeMachine}, appleIIeScreenshots},
	{"card-cat", []string{cardCatMachine}, cardCatScreenshots},
	{"ultraterm", []string{ultratermMachine}, ultratermScreenshots},
	{"total-replay", []string{totalReplayMachine}, totalReplayScreenshots},
	{"forth", []string{forthMachine}, forthScreenshots},
	{"prodos", []string{prodosMachine}, prodosScreenshots},
	{"visicalc", []string{visiCalcMachine}, visiCalcScreenshots},
	{"logo", []string{logoMachine}, logoScreenshots},
}

func TestActivities(t *testing.T) {
	if os.Getenv("A2_ACTIVITIES") == "" {
		t.Skip("this makes the pictures of the guides, with A2_ACTIVITIES=1")
	}
	for _, a := range activities {
		t.Run(a.guide, a.pictures)
	}
}

// TestMachines checks that each page gives the command lines its pictures
// were made with, word for word
func TestMachines(t *testing.T) {
	for _, a := range activities {
		page, err := os.ReadFile(filepath.Join("../guides", a.guide+".md"))
		if err != nil {
			t.Fatal(err)
		}
		shown := map[string]bool{}
		for _, block := range strings.Split(string(page), "```bash\n")[1:] {
			command, _, _ := strings.Cut(block, "```")
			shown[strings.Join(strings.Fields(command), " ")] = true
		}
		for _, machine := range a.machines {
			if !shown[strings.Join(strings.Fields(machine), " ")] {
				t.Errorf("%v.md does not give the command line\n%v", a.guide, machine)
			}
		}
	}
}
