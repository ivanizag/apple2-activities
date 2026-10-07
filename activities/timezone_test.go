package activities

import (
	"bufio"
	"fmt"
	"html"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
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
	timeZoneLine timeZoneWait = iota // a line, in the keyboard routine of the ROM or its own
	timeZoneMore                     // Return, after a page of text
	timeZoneOver                     // nothing, the game is won
)

/*
timeZoneWaitForKeys runs the machine until the game waits for keys, by where
the processor is: in the keyboard routine of the ROM at $FD1B, or in the loop
at $4780 that hums the time machine while it reads the keyboard, for a line;
in the loop at $6477, for Return to go on. The screen is still by then, the
picture drawn and the text printed. At the end of the game it plays a tune
and waits for nothing.
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

// timeZonePicture is a picture of the walkthrough, and the text under it
type timeZonePicture struct {
	file string
	text []string
}

// timeZoneCommand is a command played, the pictures of its answer, and the
// disks asked for on the way
type timeZoneCommand struct {
	number   int
	command  string
	pictures []timeZonePicture
	disks    []string
}

/*
timeZoneScreenshots plays Time Zone from the start to the end, on an Apple
][+, as the walkthrough says, a picture of the screen for each command, and
one for each page of a longer answer. It changes the disks when the game asks
for them, and writes the walkthrough of the page with the pictures.
*/
func timeZoneScreenshots(t *testing.T) {
	pictures := newAlbum("timezone", album.Color)
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
		name := fmt.Sprintf("%04d", c.number)
		page := func() {
			file := fmt.Sprintf("%v-%d", name, len(c.pictures)+1)
			must(t, pictures.Screenshot(o, file))
			c.pictures = append(c.pictures, timeZonePicture{file, timeZoneLines(o, 4)})
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
				t.Fatalf("%v %v: the game did not take the disk:\n%v", c.number, c.command, o.Text())
			case disk != nil && last == "AND PRESS RETURN.":
				if disk[0] == asked {
					t.Fatalf("%v %v: the game asked for the disk again:\n%v", c.number, c.command, o.Text())
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
				t.Fatalf("%v %v: the game was lost:\n%v", c.number, c.command, o.Text())
			case !strings.HasPrefix(last, timeZonePrompt) && !strings.HasSuffix(last, "?"):
				// A message that waits for a key
				page()
				must(t, o.Type("\n"))
				continue
			}
			// The prompt, or a question the next command answers
			file := name
			if c.number == 0 {
				file = "start"
			}
			must(t, pictures.Screenshot(o, file))
			c.pictures = append(c.pictures, timeZonePicture{file, timeZoneLines(o, 4)})
			return
		}
		t.Fatalf("%v %v: the game never asked for the next command", c.number, c.command)
	}

	start := timeZoneCommand{}
	answer(&start)
	var walkthrough strings.Builder
	timeZoneWrite(&walkthrough, start)
	number := 0
	for _, step := range steps {
		if step.section != "" {
			fmt.Fprintf(&walkthrough, "### %v\n\n", step.section)
			continue
		}
		number++
		c := timeZoneCommand{number: number, command: step.command}
		must(t, o.Type(step.command+"\n"))
		answer(&c)
		timeZoneWrite(&walkthrough, c)
	}
	if !o.HasText(timeZoneEnding) {
		t.Fatalf("the walkthrough did not end the game:\n%v", o.Text())
	}
	must(t, timeZoneWritePage(walkthrough.String()))
}

// timeZoneDiskFile is a copy of a side of the game, by its letter
func timeZoneDiskFile(t *testing.T, letter string) string {
	return disk(t, fmt.Sprintf(timeZoneDiskNamed, letter))
}

// timeZoneWrite writes a command, its pictures, and the disks it asked for,
// on the page
func timeZoneWrite(w *strings.Builder, c timeZoneCommand) {
	if c.number > 0 {
		fmt.Fprintf(w, "**%d.** `%v`\n\n", c.number, c.command)
	}
	for i, p := range c.pictures {
		if i > 0 {
			w.WriteString("\n")
		}
		alt := strings.Join(strings.Fields(strings.Join(p.text, " ")), " ")
		alt = strings.TrimSpace(strings.TrimSuffix(alt, timeZonePrompt+"?"))
		fmt.Fprintf(w, `<img src="images/timezone/%v.png" width="400" alt="%v">`,
			p.file, html.EscapeString(alt))
	}
	w.WriteString("\n\n")
	for _, d := range c.disks {
		fmt.Fprintf(w, "*Side %v: drop `%v` on drive 1, and press Return.*\n\n",
			d, fmt.Sprintf(timeZoneDiskNamed, d[1:]))
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
			want = append(want, fmt.Sprintf("**%d.** `%v`", number, s.command))
		}
	}
	for _, line := range strings.Split(string(page), "\n") {
		if strings.HasPrefix(line, "### ") || strings.HasPrefix(line, "**") && strings.Contains(line, ".** `") {
			got = append(got, line)
		}
	}
	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Errorf("the walkthrough of %v is not the one of %v: run its generator", timeZonePage, timeZoneWalkthrough)
	}
}
