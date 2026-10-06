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
	"regexp"
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
	{"karateka", []string{karatekaMachine}, karatekaScreenshots},
	{"pascal-2048", []string{pascalMachine}, pascal2048Screenshots},
	{"printing", []string{printingMachine}, printingScreenshots},
	{"applesoft-snake", []string{snakeMachine}, snakeScreenshots},
	{"merlin", []string{merlinMachine}, merlinScreenshots},
	{"rgb-card", []string{rgbCardMachine}, rgbCardScreenshots},
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

// listings are the programs of the guides, and the pages that show them
var listings = []struct {
	guide, listing string
}{
	{"ultraterm", ultratermModes},
	{"pascal-2048", game2048Listing},
	{"printing", calendarListing},
	{"applesoft-snake", snakeListing},
	{"merlin", barsListing},
}

/*
TestListings checks that each page shows its programs whole, as the files
the generators type and the reader downloads: the code blocks of the
language of a listing, one after the other, are the listing
*/
func TestListings(t *testing.T) {
	for _, l := range listings {
		listing, err := os.ReadFile(l.listing)
		if err != nil {
			t.Fatal(err)
		}
		page, err := os.ReadFile(filepath.Join("../guides", l.guide+".md"))
		if err != nil {
			t.Fatal(err)
		}
		language := listingLanguage(l.listing)
		if blocks := strings.Join(codeBlocks(string(page), language), ""); blocks != string(listing) {
			t.Errorf("the %v blocks of %v.md are not %v", language, l.guide, l.listing)
		}
	}
}

// printouts are what the generators printed, and the pages that show them
var printouts = []struct {
	guide, printout string
}{
	{"printing", "../guides/images/printing/listing.txt"},
	{"printing", "../guides/images/printing/calendar.txt"},
}

// TestPrintouts checks that each page shows what its generator printed, whole,
// in a block of text
func TestPrintouts(t *testing.T) {
	for _, p := range printouts {
		printout, err := os.ReadFile(p.printout)
		if err != nil {
			t.Fatal(err)
		}
		page, err := os.ReadFile(filepath.Join("../guides", p.guide+".md"))
		if err != nil {
			t.Fatal(err)
		}
		shown := false
		for _, block := range codeBlocks(string(page), "text") {
			shown = shown || block == string(printout)
		}
		if !shown {
			t.Errorf("%v.md does not show %v", p.guide, p.printout)
		}
	}
}

// listingLanguage is the language of the code blocks of a listing, by its
// extension
func listingLanguage(path string) string {
	switch filepath.Ext(path) {
	case ".bas":
		return "basic"
	case ".pas":
		return "pascal"
	case ".s":
		return "asm"
	}
	return strings.TrimPrefix(filepath.Ext(path), ".")
}

/*
codeBlocks are the code blocks of a language of a page, each without the
indentation of its fence, as Markdown takes it off inside a list item
*/
func codeBlocks(page string, language string) []string {
	fence := regexp.MustCompile("(?m)^( *)```" + language + "\n")
	var blocks []string
	for _, at := range fence.FindAllStringSubmatchIndex(page, -1) {
		indent := page[at[2]:at[3]]
		code, _, _ := strings.Cut(page[at[1]:], indent+"```")
		var lines []string
		for _, line := range strings.SplitAfter(code, "\n") {
			lines = append(lines, strings.TrimPrefix(line, indent))
		}
		blocks = append(blocks, strings.Join(lines, ""))
	}
	return blocks
}
