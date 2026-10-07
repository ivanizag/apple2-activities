package activities

import (
	"bufio"
	"bytes"
	"fmt"
	"html"
	"image"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
	izscreen "github.com/ivanizag/izapple2/screen"
)

/*
An adventure is one of the Hi-Res Adventures of On-Line Systems, played from
its start to its end as a walkthrough says, a command a line. They share the
way they talk: a picture, four lines of text under it, and the prompt
"--------------- ENTER COMMAND?", where the game reads a line with the
keyboard routine of the ROM; a longer text stops every four lines, in a loop
at $6474, for Return. What each game has of its own is here.
*/
type adventure struct {
	// name is the page, guides/<name>.md, and its pictures, images/<name>/
	name string
	// walkthrough is the file of the commands, in sections
	walkthrough string
	// waits are the other places where the game waits for keys, by the
	// processor
	waits []adventureWait
	// disk reads a request for a disk at the bottom of the text: what to say
	// on the page, as "Side 1B", and the file to put in drive 1
	disk func(lines []string) (side string, file string, asked bool)
	// ending is a text of the end of the game
	ending string
	// prompt is where the game reads a command, if not adventurePrompt
	prompt string
}

// commandPrompt is where the game reads a command
func (a adventure) commandPrompt() string {
	if a.prompt != "" {
		return a.prompt
	}
	return adventurePrompt
}

// adventureWait is a place where the game waits for keys
type adventureWait struct {
	from, to uint16
	keys     adventureKeys
}

// adventureKeys is what the game waits for
type adventureKeys int

const (
	adventureLine adventureKeys = iota // a line, a command or an answer
	adventureMore                      // Return, after a page of text
	adventureOver                      // nothing, the game is won
)

// adventurePrompt is the prompt where the games read a command
const adventurePrompt = "--------------- ENTER COMMAND"

// The marks of the page between which the generator writes the walkthrough
const (
	adventureBegin = "<!-- The walkthrough, written by the generator -->\n"
	adventureEnd   = "<!-- End of the walkthrough -->\n"
)

// adventureTheEnd is the title of the last section, the ending of the game
const adventureTheEnd = "The end"

// adventureTextTop is the first line of the four of text, of the 192 of the
// screen
const adventureTextTop = 160

// adventureStep is a line of the walkthrough: a section or a command
type adventureStep struct {
	section string
	command string
}

// steps reads the walkthrough
func (a adventure) steps() ([]adventureStep, error) {
	f, err := os.Open(a.walkthrough)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var steps []adventureStep
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		switch {
		case line == "", strings.HasPrefix(line, "# "):
		case strings.HasPrefix(line, "## "):
			steps = append(steps, adventureStep{section: line[3:]})
		default:
			steps = append(steps, adventureStep{command: line})
		}
	}
	return steps, s.Err()
}

/*
waitForKeys runs the machine until the game waits for keys, by where the
processor is: in the keyboard routine of the ROM at $FD1B, or in the loop at
$6474 for Return after a page of text, or in one of the places of the game. The screen is still by then, the picture drawn and the
text printed. At the end of the game it waits for nothing.
*/
func (a adventure) waitForKeys(o *operator.Operator) (adventureKeys, error) {
	keys, frames := adventureKeys(-1), 0
	for range 60 * 60 {
		o.Run(1)
		pc := o.Apple2().GetPC()
		now := adventureKeys(-1)
		switch {
		case pc >= 0xfd1b && pc <= 0xfd2e:
			now = adventureLine
		case pc >= 0x6474 && pc <= 0x647b:
			now = adventureMore
		}
		for _, w := range a.waits {
			if pc >= w.from && pc <= w.to {
				now = w.keys
			}
		}
		if now != keys {
			keys, frames = now, 0
		}
		frames++
		if keys >= 0 && frames >= 20 {
			return keys, nil
		}
	}
	if o.HasText(a.ending) {
		return adventureOver, nil
	}
	return 0, fmt.Errorf("the game did not wait for keys:\n%v\nthe processor runs at %v", o.Text(), adventurePlaces(o))
}

// adventurePlaces are the addresses the processor runs at in a moment, to
// find where a game waits that the adventure does not know
func adventurePlaces(o *operator.Operator) string {
	seen := map[uint16]bool{}
	for range 120 {
		o.Run(1)
		seen[o.Apple2().GetPC()] = true
	}
	places := make([]string, 0, len(seen))
	for pc := range seen {
		places = append(places, fmt.Sprintf("$%04X", pc))
	}
	sort.Strings(places)
	return strings.Join(places, " ")
}

// adventureLines are the last lines of the text, the four the game shows
// under its picture
func adventureLines(o *operator.Operator, n int) []string {
	lines := strings.Split(strings.TrimRight(o.Text(), "\n"), "\n")
	return lines[max(len(lines)-n, 0):]
}

// adventurePicture is a picture of the walkthrough, of the screen or of its
// last lines of text, and the four lines of text
type adventurePicture struct {
	file string
	text []string
}

// adventureShown is what the game shows after a command, its pictures, named
// after the step of the page that shows them, and the disks asked for on the
// way
type adventureShown struct {
	name     string
	command  string
	pictures []adventurePicture
	disks    []string
	// The picture above the text, and the text, of the last one taken
	graphics []byte
	text     []string
}

/*
take keeps the screen as a picture. When the picture above the text is the
same as the last one, it keeps only the lines of text that are new since
then, cut from the bottom of the screen, which the page shows under that
picture; when there are none, nothing.
*/
func (s *adventureShown) take(o *operator.Operator, pictures *album.Album, file string) error {
	screen := pictures.Screen(o)
	graphics := screen.Pix
	mixed := o.Apple2().GetVideoSource().GetCurrentVideoMode()&izscreen.VideoMixTextMask != 0
	if mixed {
		// The lines above the four of text, drawn twice
		graphics = graphics[:2*adventureTextTop*screen.Stride]
	}
	text := adventureLines(o, 4)
	var picture image.Image = screen
	if mixed && s.graphics != nil && bytes.Equal(graphics, s.graphics) {
		lines := adventureNewLines(s.text, text)
		if strings.TrimSpace(strings.Join(text[len(text)-lines:], "")) == "" {
			// Nothing new, or only empty lines and the cursor
			return nil
		}
		// Each line of text is 8 lines of the screen, drawn twice
		b := screen.Bounds()
		picture = screen.SubImage(image.Rect(b.Min.X, b.Max.Y-2*8*lines, b.Max.X, b.Max.Y))
	}
	if err := pictures.Write(picture, file); err != nil {
		return err
	}
	s.pictures = append(s.pictures, adventurePicture{file: file, text: text})
	s.graphics, s.text = graphics, text
	return nil
}

// adventureNewLines is how many lines at the bottom of the text after are not
// in the text before, which scrolled up to make room for them
func adventureNewLines(before, after []string) int {
	if strings.Join(before, "\n") == strings.Join(after, "\n") {
		return 0
	}
	for k := min(len(before), len(after)) - 1; k > 0; k-- {
		if strings.Join(before[len(before)-k:], "\n") == strings.Join(after[:k], "\n") {
			return len(after) - k
		}
	}
	return len(after)
}

/*
answer runs the game after a command until it asks for the next one, taking a
picture each time it waits: for Return after a page of text, or a key after a
message, which it is given, for a disk, which is put in drive 1, and for the
next command or the answer to a question, where it stops.
*/
func (a adventure) answer(t *testing.T, o *operator.Operator, pictures *album.Album, s *adventureShown) {
	page := func() {
		must(t, s.take(o, pictures, fmt.Sprintf("%v-%d", s.name, len(s.pictures)+1)))
	}
	asked := ""
	for range 50 {
		keys, err := a.waitForKeys(o)
		must(t, err)
		last := strings.TrimSpace(adventureLines(o, 1)[0])
		side, file, asking := a.disk(adventureLines(o, 3))
		switch {
		case keys == adventureOver:
			o.RunSeconds(5)
		case last == "*" || last == "]":
			// The end of a game that leaves to the Monitor of the ROM, or to
			// BASIC, after its last page: its prompt, and the last of the text
			must(t, s.take(o, pictures, s.name))
			return
		case o.HasText(a.ending) && keys == adventureLine:
			// The end, which may ask to play again
		case keys == adventureMore:
			page()
			must(t, o.Type("\n"))
			continue
		case strings.Contains(o.Text(), "WRONG DISK"):
			t.Fatalf("%v %v: the game did not take the disk:\n%v", s.name, s.command, o.Text())
		case asking:
			if side == asked {
				t.Fatalf("%v %v: the game asked for the disk again:\n%v", s.name, s.command, o.Text())
			}
			asked = side
			page()
			must(t, o.InsertDisk(0, disk(t, file)))
			s.disks = append(s.disks, side+": "+file)
			// The drive stops before the game reads the new disk
			o.RunSeconds(3)
			must(t, o.Type("\n"))
			continue
		case strings.Contains(last, "PLAY AGAIN"):
			t.Fatalf("%v %v: the game was lost:\n%v", s.name, s.command, o.Text())
		case !strings.HasPrefix(last, a.commandPrompt()) && !strings.HasSuffix(last, "?"):
			// A message that waits for a key
			page()
			must(t, o.Type("\n"))
			continue
		}
		// The prompt, or a question the next command answers
		must(t, s.take(o, pictures, s.name))
		if os.Getenv("ADVENTURE_LOG") != "" {
			fmt.Printf("%v %-18v| %v\n", s.name, s.command, adventureAnswer(o, s.command))
		}
		return
	}
	t.Fatalf("%v %v: the game never asked for the next command:\n%v", s.name, s.command, o.Text())
}

/*
play plays the walkthrough, from the first prompt of the game, the one it is
coming to now, to the end, and writes it on the page: each step is what the
game shows, and the command typed then.
*/
func (a adventure) play(t *testing.T, o *operator.Operator, pictures *album.Album) {
	steps, err := a.steps()
	must(t, err)
	shown := adventureShown{name: "0001"}
	a.answer(t, o, pictures, &shown)
	var walkthrough strings.Builder
	walkthrough.WriteString(a.index(steps))
	number, section := 0, 0
	for _, step := range steps {
		if step.section != "" {
			section++
			fmt.Fprintf(&walkthrough, "### %v\n\n", adventureSection(section, step.section))
			continue
		}
		number++
		a.write(&walkthrough, shown, step.command)
		must(t, o.Type(step.command+"\n"))
		shown = adventureShown{name: fmt.Sprintf("%04d", number+1), command: step.command}
		a.answer(t, o, pictures, &shown)
	}
	fmt.Fprintf(&walkthrough, "### %v\n\n", adventureTheEnd)
	a.write(&walkthrough, shown, "")
	if !o.HasText(a.ending) {
		t.Fatalf("the walkthrough did not end the game:\n%v", o.Text())
	}
	must(t, a.writePage(walkthrough.String()))
	must(t, a.prune())
}

// prune removes the pictures of the folder of the game that its page does
// not show, left by a run that stopped on the way
func (a adventure) prune() error {
	page, err := os.ReadFile(a.page())
	if err != nil {
		return err
	}
	folder := "../guides/images/" + a.name
	files, err := os.ReadDir(folder)
	if err != nil {
		return err
	}
	for _, f := range files {
		if !strings.Contains(string(page), "images/"+a.name+"/"+f.Name()) {
			if err := os.Remove(folder + "/" + f.Name()); err != nil {
				return err
			}
		}
	}
	return nil
}

// page is the file of the page
func (a adventure) page() string {
	return "../guides/" + a.name + ".md"
}

// adventureSection is the title of a section, with its number
func adventureSection(number int, title string) string {
	return fmt.Sprintf("%d. %v", number, title)
}

// adventureAnchor is the name GitHub gives a title, to link to it: in small
// letters, without punctuation, the spaces made hyphens
func adventureAnchor(title string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(title) {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9', r == '-', r == '_':
			b.WriteRune(r)
		case r == ' ':
			b.WriteRune('-')
		}
	}
	return b.String()
}

// index is the list of the sections, each a link to it
func (a adventure) index(steps []adventureStep) string {
	var b strings.Builder
	section := 0
	for _, step := range steps {
		if step.section != "" {
			section++
			title := adventureSection(section, step.section)
			fmt.Fprintf(&b, "- [%v](#%v)\n", title, adventureAnchor(title))
		}
	}
	fmt.Fprintf(&b, "- [%v](#%v)\n\n", adventureTheEnd, adventureAnchor(adventureTheEnd))
	return b.String()
}

// write writes a step on the page: what the game shows, the disks it asks
// for, and the command to type then
func (a adventure) write(w *strings.Builder, s adventureShown, command string) {
	for i, p := range s.pictures {
		if i > 0 {
			// The pictures one under the other, the text under its picture
			w.WriteString("<br>\n")
		}
		alt := strings.Join(strings.Fields(strings.Join(p.text, " ")), " ")
		alt = strings.TrimSpace(strings.TrimSuffix(alt, a.commandPrompt()+"?"))
		fmt.Fprintf(w, "<img src=\"images/%v/%v.png\" width=\"400\" alt=\"%v\">",
			a.name, p.file, html.EscapeString(alt))
	}
	w.WriteString("\n\n")
	for _, d := range s.disks {
		side, file, _ := strings.Cut(d, ": ")
		fmt.Fprintf(w, "*%v: drop `%v` on drive 1, and press Return.*\n\n", side, file)
	}
	if command != "" {
		fmt.Fprintf(w, "**Type `%v`**\n\n", command)
	}
}

// writePage puts the walkthrough on the page, between its marks
func (a adventure) writePage(walkthrough string) error {
	page, err := os.ReadFile(a.page())
	if err != nil {
		return err
	}
	text := string(page)
	begin := strings.Index(text, adventureBegin)
	end := strings.Index(text, adventureEnd)
	if begin < 0 || end < begin {
		return fmt.Errorf("%v has no marks for the walkthrough", a.page())
	}
	text = text[:begin+len(adventureBegin)] + "\n" + walkthrough + text[end:]
	return os.WriteFile(a.page(), []byte(text), 0o644)
}

/*
checkPage checks that the page has the walkthrough the generator played, its
index, its sections and its commands in their order, so that a change to the
walkthrough is not left out of the page
*/
func (a adventure) checkPage(t *testing.T) {
	steps, err := a.steps()
	must(t, err)
	page, err := os.ReadFile(a.page())
	must(t, err)
	var want, got []string
	section := 0
	for _, s := range steps {
		if s.section != "" {
			section++
			want = append(want, "### "+adventureSection(section, s.section))
		} else {
			want = append(want, fmt.Sprintf("**Type `%v`**", s.command))
		}
	}
	want = append(want, "### "+adventureTheEnd)
	for _, line := range strings.Split(string(page), "\n") {
		if strings.HasPrefix(line, "### ") || strings.HasPrefix(line, "**Type `") {
			got = append(got, line)
		}
	}
	if !strings.Contains(string(page), a.index(steps)) {
		t.Errorf("the index of %v is not the one of %v: run its generator", a.page(), a.walkthrough)
	}
	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Errorf("the walkthrough of %v is not the one of %v: run its generator", a.page(), a.walkthrough)
	}
}

/*
adventureAnswer is the text the game wrote after a command, as far back as the
screen keeps it, in one line: ADVENTURE_LOG prints it for each command, to
follow the game when the walkthrough goes wrong
*/
func adventureAnswer(o *operator.Operator, command string) string {
	// The lines of 40 characters go on in the next one
	var text strings.Builder
	for _, line := range strings.Split(o.Text(), "\n") {
		text.WriteString(line)
		if len(line) < 40 {
			text.WriteString(" ")
		}
	}
	all := text.String()
	if i := strings.LastIndex(all, "COMMAND?"+command); i >= 0 && command != "" {
		all = all[i+len("COMMAND?"+command):]
	}
	return strings.Join(strings.Fields(all), " ")
}
