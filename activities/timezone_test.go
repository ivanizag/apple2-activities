package activities

import (
	"bufio"
	"bytes"
	"fmt"
	"html"
	"image"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
	izscreen "github.com/ivanizag/izapple2/screen"
)

// timeZoneMachine is the machine of the guide, as its command line of
// izapple2: an Apple ][+ with side A of Time Zone in drive 1
const timeZoneMachine = `izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s6 'diskii,disk1=disks/Time Zone (4am and san inc crack) disk A.dsk'`

// timeZoneWalkthrough is the game played from the start to the end, a
// command a line, in sections
const timeZoneWalkthrough = "../guides/listings/timezone.txt"

// timeZonePage is the page, whose walkthrough the generator writes between
// two marks
const (
	timeZonePage      = "../guides/timezone.md"
	timeZoneBegin     = "<!-- The walkthrough, written by the generator -->\n"
	timeZoneEnd       = "<!-- End of the walkthrough -->\n"
	timeZonePrompt    = "--------------- ENTER COMMAND"
	timeZoneEnding    = "ULTIMATE ADVENTURER"
	timeZoneDiskNamed = "Time Zone (4am and san inc crack) disk %v.dsk"
)

// timeZoneDisk is the game asking for a side of its six disks, 1A to 6L; the
// letter names the file
var timeZoneDisk = regexp.MustCompile(`INSERT DISK NUMBER (\d([A-L]))`)

// timeZoneStep is a line of the walkthrough: a section or a command
type timeZoneStep struct {
	section string
	command string
}

// timeZoneSteps reads the walkthrough
func timeZoneSteps(path string) ([]timeZoneStep, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var steps []timeZoneStep
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		switch {
		case line == "", strings.HasPrefix(line, "# "):
		case strings.HasPrefix(line, "## "):
			steps = append(steps, timeZoneStep{section: line[3:]})
		default:
			steps = append(steps, timeZoneStep{command: line})
		}
	}
	return steps, s.Err()
}

// timeZoneWait is what the game waits for
type timeZoneWait int

const (
	timeZoneLine timeZoneWait = iota // a line, in the keyboard routine of the ROM or its pause
	timeZoneMore                     // Return, after a page of text
	timeZoneOver                     // nothing, the game is won
)

/*
timeZoneWaitForKeys runs the machine until the game waits for keys, by where
the processor is: in the keyboard routine of the ROM at $FD1B, or in the
pause at $4780 that the game makes in the time machine before its prompt,
reading the keyboard, for a line; in the loop at $6477, for Return to go on.
The screen is still by then, the picture drawn and the text printed. At the
end of the game it waits for nothing.
*/
func timeZoneWaitForKeys(o *operator.Operator) (timeZoneWait, error) {
	wait, frames := timeZoneWait(-1), 0
	for range 60 * 60 {
		o.Run(1)
		pc := o.Apple2().GetPC()
		now := timeZoneWait(-1)
		switch {
		case pc >= 0xfd1b && pc <= 0xfd2e, pc >= 0x4780 && pc <= 0x47e1:
			now = timeZoneLine
		case pc >= 0x6477 && pc <= 0x647b:
			now = timeZoneMore
		}
		if now != wait {
			wait, frames = now, 0
		}
		frames++
		if wait >= 0 && frames >= 20 {
			return wait, nil
		}
	}
	if o.HasText(timeZoneEnding) {
		return timeZoneOver, nil
	}
	return 0, fmt.Errorf("the game did not wait for keys:\n%v", o.Text())
}

// timeZoneLines are the last lines of the text, the four the game shows
// under its picture
func timeZoneLines(o *operator.Operator, n int) []string {
	lines := strings.Split(strings.TrimRight(o.Text(), "\n"), "\n")
	return lines[max(len(lines)-n, 0):]
}

// timeZonePicture is a picture of the walkthrough, of the screen or of its
// four lines of text, and that text
type timeZonePicture struct {
	file string
	text []string
}

// timeZoneCommand is the answer of the game to a command, its pictures, named
// after the step of the page that shows them, and the disks asked for on the
// way
type timeZoneCommand struct {
	name     string
	command  string
	pictures []timeZonePicture
	disks    []string
	// The picture above the text, and the text, of the last one taken
	graphics []byte
	text     []string
}

/*
take keeps the screen of the answer of a command as a picture. When the
picture above the text is the same as the last one, it keeps only the lines
of text that are new since then, cut from the bottom of the screen, which the
page shows under that picture; when there are none, nothing.
*/
func (c *timeZoneCommand) take(o *operator.Operator, pictures *album.Album, file string) error {
	screen := pictures.Screen(o)
	graphics := screen.Pix
	mixed := o.Apple2().GetVideoSource().GetCurrentVideoMode()&izscreen.VideoMixTextMask != 0
	if mixed {
		// The lines above the four of text, drawn twice
		graphics = graphics[:2*timeZoneTextTop*screen.Stride]
	}
	text := timeZoneLines(o, 4)
	var picture image.Image = screen
	if mixed && c.graphics != nil && bytes.Equal(graphics, c.graphics) {
		lines := timeZoneNewLines(c.text, text)
		if lines == 0 {
			return nil
		}
		// Each line of text is 8 lines of the screen, drawn twice
		b := screen.Bounds()
		picture = screen.SubImage(image.Rect(b.Min.X, b.Max.Y-2*8*lines, b.Max.X, b.Max.Y))
	}
	if err := pictures.Write(picture, file); err != nil {
		return err
	}
	c.pictures = append(c.pictures, timeZonePicture{file: file, text: text})
	c.graphics, c.text = graphics, text
	return nil
}

// timeZoneNewLines is how many lines at the bottom of the text after are not
// in the text before, which scrolled up to make room for them
func timeZoneNewLines(before, after []string) int {
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

// timeZoneTextTop is the first line of the four of text, of the 192 of the
// screen
const timeZoneTextTop = 160

/*
timeZoneScreenshots plays Time Zone from the start to the end, on an Apple
][+, as the walkthrough says, a picture of the screen for each command, and
one for each page of a longer answer. It changes the disks when the game asks
for them, and writes the walkthrough of the page with the pictures.
*/
func timeZoneScreenshots(t *testing.T) {
	pictures := newAlbum("timezone", album.ColorWhiteText)
	steps, err := timeZoneSteps(timeZoneWalkthrough)
	must(t, err)
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

	// Each command and its answer, a picture each time the game waits
	answer := func(c *timeZoneCommand) {
		name := c.name
		page := func() {
			must(t, c.take(o, pictures, fmt.Sprintf("%v-%d", name, len(c.pictures)+1)))
		}
		asked := ""
		for range 50 {
			wait, err := timeZoneWaitForKeys(o)
			must(t, err)
			last := strings.TrimSpace(timeZoneLines(o, 1)[0])
			disk := timeZoneDisk.FindStringSubmatch(strings.Join(timeZoneLines(o, 3), " "))
			switch {
			case wait == timeZoneOver:
				o.RunSeconds(5)
			case wait == timeZoneMore:
				page()
				must(t, o.Type("\n"))
				continue
			case strings.Contains(o.Text(), "WRONG DISK"):
				t.Fatalf("%v %v: the game did not take the disk:\n%v", c.name, c.command, o.Text())
			case disk != nil && last == "AND PRESS RETURN.":
				if disk[0] == asked {
					t.Fatalf("%v %v: the game asked for the disk again:\n%v", c.name, c.command, o.Text())
				}
				asked = disk[0]
				page()
				must(t, o.InsertDisk(0, timeZoneDiskFile(t, disk[2])))
				c.disks = append(c.disks, disk[1])
				// The drive stops before the game reads the new disk
				o.RunSeconds(3)
				must(t, o.Type("\n"))
				continue
			case strings.Contains(last, "PLAY AGAIN"):
				t.Fatalf("%v %v: the game was lost:\n%v", c.name, c.command, o.Text())
			case !strings.HasPrefix(last, timeZonePrompt) && !strings.HasSuffix(last, "?"):
				// A message that waits for a key
				page()
				must(t, o.Type("\n"))
				continue
			}
			// The prompt, or a question the next command answers
			must(t, c.take(o, pictures, name))
			return
		}
		t.Fatalf("%v %v: the game never asked for the next command", c.name, c.command)
	}

	// Each step of the page is what the game shows, and the command typed
	// then
	shown := timeZoneCommand{name: "0001"}
	answer(&shown)
	var walkthrough strings.Builder
	number := 0
	for _, step := range steps {
		if step.section != "" {
			fmt.Fprintf(&walkthrough, "### %v\n\n", step.section)
			continue
		}
		number++
		timeZoneWrite(&walkthrough, number, shown, step.command)
		must(t, o.Type(step.command+"\n"))
		shown = timeZoneCommand{name: fmt.Sprintf("%04d", number+1), command: step.command}
		answer(&shown)
	}
	walkthrough.WriteString("### The end\n\n")
	timeZoneWrite(&walkthrough, 0, shown, "")
	if !o.HasText(timeZoneEnding) {
		t.Fatalf("the walkthrough did not end the game:\n%v", o.Text())
	}
	must(t, timeZoneWritePage(walkthrough.String()))
}

// timeZoneDiskFile is a copy of a side of the game, by its letter
func timeZoneDiskFile(t *testing.T, letter string) string {
	return disk(t, fmt.Sprintf(timeZoneDiskNamed, letter))
}

// timeZoneWrite writes a step on the page: its number, what the game shows,
// the disks it asks for, and the command to type then
func timeZoneWrite(w *strings.Builder, number int, c timeZoneCommand, command string) {
	if number > 0 {
		fmt.Fprintf(w, "**%d.**\n\n", number)
	}
	for i, p := range c.pictures {
		if i > 0 {
			// The pictures one under the other, the text under its picture
			w.WriteString("<br>\n")
		}
		alt := strings.Join(strings.Fields(strings.Join(p.text, " ")), " ")
		alt = strings.TrimSpace(strings.TrimSuffix(alt, timeZonePrompt+"?"))
		fmt.Fprintf(w, "<img src=\"images/timezone/%v.png\" width=\"400\" alt=\"%v\">",
			p.file, html.EscapeString(alt))
	}
	w.WriteString("\n\n")
	for _, d := range c.disks {
		fmt.Fprintf(w, "*Side %v: drop `%v` on drive 1, and press Return.*\n\n",
			d, fmt.Sprintf(timeZoneDiskNamed, d[1:]))
	}
	if command != "" {
		fmt.Fprintf(w, "**Type `%v`**\n\n", command)
	}
}

// timeZoneWritePage puts the walkthrough on the page, between its marks
func timeZoneWritePage(walkthrough string) error {
	page, err := os.ReadFile(timeZonePage)
	if err != nil {
		return err
	}
	text := string(page)
	begin := strings.Index(text, timeZoneBegin)
	end := strings.Index(text, timeZoneEnd)
	if begin < 0 || end < begin {
		return fmt.Errorf("the page has no marks for the walkthrough")
	}
	text = text[:begin+len(timeZoneBegin)] + "\n" + walkthrough + text[end:]
	return os.WriteFile(timeZonePage, []byte(text), 0o644)
}

/*
TestTimeZoneWalkthrough checks that the page has the walkthrough the
generator played, its sections and commands in their order, so that a change
to the walkthrough is not left out of the page
*/
func TestTimeZoneWalkthrough(t *testing.T) {
	steps, err := timeZoneSteps(timeZoneWalkthrough)
	must(t, err)
	page, err := os.ReadFile(timeZonePage)
	must(t, err)
	var want, got []string
	number := 0
	for _, s := range steps {
		if s.section != "" {
			want = append(want, "### "+s.section)
		} else {
			number++
			want = append(want, fmt.Sprintf("**%d.**", number), fmt.Sprintf("**Type `%v`**", s.command))
		}
	}
	want = append(want, "### The end")
	step := regexp.MustCompile("^\\*\\*([0-9]+\\.|Type `.*`)\\*\\*$")
	for _, line := range strings.Split(string(page), "\n") {
		if strings.HasPrefix(line, "### ") || step.MatchString(line) {
			got = append(got, line)
		}
	}
	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Errorf("the walkthrough of %v is not the one of %v: run its generator", timeZonePage, timeZoneWalkthrough)
	}
}
