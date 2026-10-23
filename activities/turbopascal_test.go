package activities

import (
	"os"
	"strings"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// turboPascalMachine is the machine of the guide, as its command line of
// izapple2
const turboPascalMachine = `izapple2 -model none -board 2e -cpu 65c02 -screen green \
    -rom "<internal>/Apple2e_Enhanced.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s3 videx \
    -s4 z80softcard \
    -s6 'diskii,disk1=disks/Borland Turbo Pascal 2.00A - Apple II CPM.dsk'`

// mazeListing is the program of the page, a maze dug and solved
const mazeListing = "../guides/listings/maze.pas"

/*
turboPascalScreenshots is Turbo Pascal 2.00A on the CP/M of the SoftCard, in
the 80 columns of a Videx Videoterm: the keys of its editor shown by TINST,
Turbo started without its error messages, the program of the page typed with
a mistake and compiled, Error 41 and nothing more; Turbo started again with
its messages, the same error explained, the mistake mended, the program
compiled and run; run again without {$A-}, which recursion needs, and
compiled with it into a command of CP/M, run from CP/M.

The Videoterm shows the menu of Turbo in inverse, characters with their top
bit set, which the text of the screen does not read as letters: the generator
waits for the keyboard, and for the text that is not inverse. The editor
reads the keyboard while it draws the screen, so a picture of it waits for
the screen to stay still.
*/
func turboPascalScreenshots(t *testing.T) {
	pictures := newAlbum("turbo-pascal", album.Green)
	o := start(t, turboPascalMachine, nil)
	source, err := os.ReadFile(mazeListing)
	must(t, err)
	// The same program with a mistake: GotoX for the first GotoXY
	mistaken := strings.Replace(string(source), "GotoXY(X, Y)", "GotoX(X, Y)", 1)

	// Started, and the disk listed
	must(t, o.WaitForText("A>", 60))
	must(t, o.WaitForPrompt(10))
	turboCommand(t, o, "DIR")
	turboCommand(t, o, "STAT")
	must(t, pictures.Screenshot(o, "started"))

	// The keys of the editor, as TINST has them for this CP/M, down to End
	// edit, the 38th, then the rest let be and TINST left
	turboCommand(t, o, "TINST")
	must(t, o.Type("C"))
	must(t, o.WaitForText("CURSOR MOVEMENTS", 10))
	for !o.HasText("38:  End edit") {
		must(t, o.WaitForKeyboard(10))
		must(t, o.Type("\n"))
	}
	must(t, o.WaitForKeyboard(10))
	must(t, pictures.Screenshot(o, "keys"))
	for !o.HasText("Enter S, C, or Q") {
		must(t, o.Type("\n"))
		must(t, o.WaitForKeyboard(10))
	}
	must(t, o.Type("Q"))
	must(t, o.WaitForPrompt(10))

	// Turbo Pascal, without its error messages
	must(t, o.TypeLines("TURBO"))
	must(t, o.WaitForText("(Y/N)?", 30))
	must(t, o.WaitForKeyboard(10))
	must(t, pictures.Screenshot(o, "question"))
	must(t, o.Type("N"))
	must(t, o.WaitForPrompt(30))
	must(t, pictures.Screenshot(o, "menu"))

	// The program typed in the editor, the indent of the editor off, as the
	// program brings its own, with the mistake
	must(t, o.Type("E"))
	must(t, o.WaitForKeyboard(10))
	must(t, o.TypeLines("MAZE"))
	must(t, o.WaitForKeyboard(30))
	must(t, o.Key("Ctrl+Q"))
	must(t, o.Type("I"))
	must(t, o.WaitForKeyboard(10))
	must(t, o.Type(mistaken))
	must(t, o.WaitForKeyboard(10))
	must(t, pictures.Screenshot(o, "typed"))
	turboEndEdit(t, o)

	// Compiled: the error by its number only
	must(t, o.Type("C"))
	must(t, o.WaitForText("Press <ESC>", 60))
	must(t, o.WaitForKeyboard(10))
	must(t, pictures.Screenshot(o, "error-number"))
	must(t, o.Key("Escape"))
	must(t, o.WaitForPrompt(30))
	turboEndEdit(t, o)

	// Turbo left, the program saved, and started again with its messages
	must(t, o.Type("Q"))
	must(t, o.WaitForText("Save (Y/N)?", 10))
	must(t, o.Type("Y"))
	must(t, o.WaitForPrompt(30))
	must(t, o.TypeLines("TURBO"))
	must(t, o.WaitForText("(Y/N)?", 30))
	must(t, o.Type("Y"))
	must(t, o.WaitForPrompt(30))
	must(t, o.Type("W"))
	must(t, o.WaitForKeyboard(10))
	must(t, o.TypeLines("MAZE"))
	must(t, o.WaitForPrompt(30))

	// Compiled again: the error explained, and the editor where it is
	must(t, o.Type("C"))
	must(t, o.WaitForText("Press <ESC>", 60))
	must(t, o.WaitForKeyboard(10))
	must(t, pictures.Screenshot(o, "error-text"))
	must(t, o.Key("Escape"))
	must(t, o.WaitForPrompt(30))
	must(t, pictures.Screenshot(o, "error-place"))

	// The mistake mended: the cursor past GotoX, and the Y typed
	for range 5 {
		must(t, o.Key("Ctrl+D"))
	}
	must(t, o.Type("Y"))
	must(t, o.WaitForPrompt(10))
	turboEndEdit(t, o)

	// Compiled into memory, and run
	must(t, o.Type("C"))
	must(t, o.WaitForText("Data:", 60))
	must(t, o.WaitForPrompt(10))
	must(t, pictures.Screenshot(o, "compiled"))
	must(t, o.Type("R"))
	maze := pictures.Record(o)
	maze.Capture(50)
	for !o.HasText("Press a key") {
		maze.Run(6, 6)
	}
	maze.Run(60, 6)
	must(t, pictures.SaveRecording(maze, "maze", 300))
	must(t, pictures.Screenshot(o, "solved"))
	must(t, o.Type(" "))
	must(t, o.WaitForPrompt(10))

	// Without {$A-}, the second line, deleted with Ctrl-Y: the variables of
	// the procedures are in one place for all their calls, and the digging
	// stops short
	must(t, o.Type("E"))
	must(t, o.WaitForPrompt(30))
	turboSecondLine(t, o)
	must(t, o.Key("Ctrl+Y"))
	must(t, o.WaitForPrompt(10))
	turboEndEdit(t, o)
	must(t, o.Type("C"))
	must(t, o.WaitForText("Data:", 60))
	must(t, o.WaitForPrompt(10))
	must(t, o.Type("R"))
	must(t, o.WaitForText("Press a key", 120))
	must(t, o.WaitForKeyboard(10))
	must(t, pictures.Screenshot(o, "static"))
	must(t, o.Type(" "))
	must(t, o.WaitForPrompt(10))

	// The line typed back
	must(t, o.Type("E"))
	must(t, o.WaitForPrompt(30))
	turboSecondLine(t, o)
	must(t, o.TypeLines("{$A-}"))
	must(t, o.WaitForPrompt(10))
	turboEndEdit(t, o)

	// Compiled into a command of CP/M, MAZE.COM
	must(t, o.Type("O"))
	must(t, o.WaitForKeyboard(10))
	must(t, o.Type("C"))
	must(t, o.WaitForKeyboard(10))
	must(t, pictures.Screenshot(o, "options"))
	must(t, o.Type("Q"))
	must(t, o.WaitForKeyboard(10))
	must(t, o.Type("C"))
	must(t, o.WaitForText("Data:", 60))
	must(t, o.WaitForPrompt(10))
	must(t, pictures.Screenshot(o, "com"))

	// Turbo left, the program saved, and the command run from CP/M
	must(t, o.Type("Q"))
	must(t, o.WaitForText("Save (Y/N)?", 10))
	must(t, o.Type("Y"))
	must(t, o.WaitForPrompt(30))
	turboCommand(t, o, "DIR MAZE.*")
	must(t, o.TypeLines("MAZE"))
	must(t, o.WaitForText("Press a key", 120))
	must(t, o.WaitForKeyboard(10))
	must(t, pictures.Screenshot(o, "command"))
}

// turboCommand types a command of CP/M and waits for the prompt after it
func turboCommand(t *testing.T, o *operator.Operator, line string) {
	t.Helper()
	must(t, o.TypeLines(line))
	must(t, o.WaitForPrompt(60))
}

// turboSecondLine takes the cursor of the editor to the second line of the
// program, with Ctrl-Q Ctrl-R, to the top of the file, and Ctrl-X, a line
// down
func turboSecondLine(t *testing.T, o *operator.Operator) {
	t.Helper()
	must(t, o.Key("Ctrl+Q"))
	must(t, o.Key("Ctrl+R"))
	must(t, o.Key("Ctrl+X"))
	must(t, o.WaitForPrompt(10))
}

// turboEndEdit leaves the editor of Turbo Pascal for its menu, with Ctrl-O
// Ctrl-D, End edit as this version has it
func turboEndEdit(t *testing.T, o *operator.Operator) {
	t.Helper()
	must(t, o.Key("Ctrl+O"))
	must(t, o.Key("Ctrl+D"))
	must(t, o.WaitForPrompt(10))
}
