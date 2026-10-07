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
	"slices"
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
	{"apple-ii", []string{appleIIMachine}, appleIIScreenshots},
	{"apple-ii-cassettes", []string{cassettesBreakout, cassettesColor, cassettesHires, cassettesRevision1}, cassettesScreenshots},
	{"apple-iie", []string{appleIIeMachine}, appleIIeScreenshots},
	{"apple-iie-models", []string{iieOriginal, iieEnhanced}, iieModelsScreenshots},
	{"applesoft-snake", []string{snakeMachine}, snakeScreenshots},
	{"appleworks", []string{appleWorksMachine}, appleWorksScreenshots},
	{"base64a", []string{base64aMachine, base64aDOS}, base64aScreenshots},
	{"basis108", []string{basis108Machine, basis108DOS}, basis108Screenshots},
	{"card-cat", []string{cardCatMachine}, cardCatScreenshots},
	{"cpm", []string{cpmMachine}, cpmScreenshots},
	{"desktop", []string{deskTopMachine}, deskTopScreenshots},
	{"dos32", []string{dos32Machine, dos32Upgraded}, dos32Screenshots},
	{"dos33", []string{dos33Machine, dos33OwnDisk}, dos33Screenshots},
	{"forth", []string{forthMachine}, forthScreenshots},
	{"integer-and-applesoft", []string{basicsTape, basicsPlusNoCard, basicsPlus, basicsInteger}, basicsScreenshots},
	{"karateka", []string{karatekaMachine}, karatekaScreenshots},
	{"lode-runner", []string{lodeRunnerMachine}, lodeRunnerScreenshots},
	{"logo", []string{logoMachine}, logoScreenshots},
	{"memory-expansion", []string{memexpMachine, memexpDOS}, memexpScreenshots},
	{"merlin", []string{merlinMachine}, merlinScreenshots},
	{"mockingboard", []string{mockingboardMachine}, mockingboardScreenshots},
	{"paddle-game", []string{paddleMachine}, paddleScreenshots},
	{"pascal", []string{pascalMachine}, pascalScreenshots},
	{"pascal-2048", []string{pascalMachine}, pascal2048Screenshots},
	{"printing", []string{printingMachine}, printingScreenshots},
	{"prodos", []string{prodosMachine}, prodosScreenshots},
	{"rgb-card", []string{rgbCardMachine}, rgbCardScreenshots},
	{"switch-on", []string{switchOnMachine}, switchOnScreenshots},
	{"swyftcard", []string{swyftCardMachine, swyftCardEmpty}, swyftCardScreenshots},
	{"thunderclock", []string{thunderclockMachine}, thunderclockScreenshots},
	{"timezone", []string{timeZoneMachine}, timeZoneScreenshots},
	{"total-replay", []string{totalReplayMachine}, totalReplayScreenshots},
	{"ultraterm", []string{ultratermMachine}, ultratermScreenshots},
	{"ulysses", []string{ulyssesMachine}, ulyssesScreenshots},
	{"visicalc", []string{visiCalcMachine}, visiCalcScreenshots},
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
// were made with, word for word, each a whole machine with its monitor
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
			if !strings.HasPrefix(machine, "izapple2 -model none ") || !strings.Contains(machine, " -screen ") {
				t.Errorf("the machine of %v.md is not whole, from -model none and with -screen\n%v", a.guide, machine)
			}
		}
	}
}

// TestOrder checks that the activities, their listings and the disks are in
// the order of their names, so that two pull requests adding to them add in
// different places and don't conflict
func TestOrder(t *testing.T) {
	var guides, listed []string
	for _, a := range activities {
		guides = append(guides, a.guide)
	}
	for _, l := range listings {
		listed = append(listed, l.guide)
	}
	list, err := os.ReadFile("../disks.tsv")
	if err != nil {
		t.Fatal(err)
	}
	var disks []string
	for _, line := range strings.Split(string(list), "\n") {
		if line != "" && !strings.HasPrefix(line, "#") {
			name, _, _ := strings.Cut(line, "\t")
			disks = append(disks, name)
		}
	}
	for what, names := range map[string][]string{
		"activities": guides, "listings": listed, "disks.tsv": disks,
	} {
		if !slices.IsSorted(names) {
			t.Errorf("%v are not in the order of their names: %v", what, names)
		}
	}
}

// listings are the programs of the guides, and the pages that show them
var listings = []struct {
	guide, listing string
}{
	{"apple-ii-cassettes", coloursListing},
	{"applesoft-snake", snakeListing},
	{"merlin", barsListing},
	{"pascal-2048", game2048Listing},
	{"printing", calendarListing},
	{"thunderclock", clockListing},
	{"ultraterm", ultratermModes},
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
	{"base64a", "../guides/images/base64a/letter.txt"},
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
