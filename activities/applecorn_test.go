package activities

import (
	"image"
	"strings"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// applecornMachine is the machine of the guide, as its command line of izapple2
const applecornMachine = `izapple2 -model none -board 2e -cpu 65c02 -screen green \
    -rom "<internal>/Apple2e_Enhanced.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s7 smartport,image1=disks/applecorn-0.54-beta.po`

// The programs of the guide, one in each language of the BBC Micro
const (
	roseListing      = "../guides/listings/rose.bas"
	stepsListing     = "../guides/listings/steps.comal"
	hanoiListing     = "../guides/listings/hanoi.lisp"
	fibonacciListing = "../guides/listings/fibonacci.forth"
	familyListing    = "../guides/listings/family.prolog"
	sieveListing     = "../guides/listings/sieve.bcpl"
	triangleListing  = "../guides/listings/triangle.pas"
)

/*
applecornScreenshots is an enhanced Apple //e running Applecorn, the
operating system of the BBC Micro on ProDOS: its menu of the language ROMs of
the BBC, and a program in each, the roses of BBC BASIC drawn in colour, COMAL,
Lisp, Forth, micro-PROLOG, BCPL compiled, ISO-Pascal edited and compiled, and
a text in the View word processor; between them, *QUIT to the program
selector of ProDOS and back to the menu. Then the seven languages loaded at
once, as sideways ROMs.
*/
func applecornScreenshots(t *testing.T) {
	pictures := newAlbum("applecorn", album.Green)
	o := start(t, applecornMachine, nil)

	// The menu of the ROMs, and BBC BASIC
	waitForApplecornMenu(t, o)
	must(t, pictures.Screenshot(o, "menu"))
	chooseROM(t, o, "1")
	must(t, pictures.Screenshot(o, "basic"))

	// The roses, typed and listed, then drawn in colour
	must(t, typeListing(o, roseListing))
	must(t, o.TypeLines("LIST"))
	must(t, o.WaitForPrompt(30))
	must(t, pictures.Screenshot(o, "rose-listing"))
	color := pictures.On(album.Color)
	drawing := color.Record(o)
	drawing.Capture(50)
	must(t, o.TypeLines("RUN"))
	from := o.Frames()
	drawUntilPrompt(t, o, drawing, color)
	t.Logf("the roses took %.1f seconds", float64(o.Frames()-from)/operator.FramesPerSecond)
	must(t, color.SaveRecording(drawing, "roses", 300))
	must(t, color.Screenshot(o, "roses"))

	// Back to the menu, by the program selector of ProDOS
	must(t, o.TypeLines("*QUIT"))
	must(t, o.WaitForText("BITSY  BYE", 30))
	must(t, o.WaitForKeyboard(10))
	must(t, pictures.Screenshot(o, "bitsy-bye"))
	backToMenu(t, o)

	// COMAL
	chooseROM(t, o, "2")
	must(t, typeListing(o, stepsListing))
	must(t, o.TypeLines("LIST"))
	must(t, o.WaitForPrompt(30))
	must(t, o.TypeLines("RUN"))
	must(t, o.WaitForPrompt(30))
	must(t, pictures.Screenshot(o, "comal"))
	quitTo(t, o, "*QUIT")

	// Lisp
	chooseROM(t, o, "3")
	must(t, typeListing(o, hanoiListing))
	must(t, o.WaitForPrompt(30))
	must(t, pictures.Screenshot(o, "lisp"))
	quitTo(t, o, "(* 'QUIT)")

	// Forth
	chooseROM(t, o, "4")
	must(t, typeListing(o, fibonacciListing))
	must(t, o.WaitForPrompt(30))
	must(t, pictures.Screenshot(o, "forth"))
	quitTo(t, o, "OS' QUIT'")

	// micro-PROLOG
	chooseROM(t, o, "5")
	must(t, typeListing(o, familyListing))
	must(t, o.WaitForPrompt(30))
	must(t, o.TypeLines("LIST ALL"))
	must(t, o.WaitForPrompt(30))
	must(t, pictures.Screenshot(o, "prolog"))
	quitTo(t, o, "** QUIT")

	// BCPL: the source copied from the keyboard to a file, until Escape,
	// then compiled and run
	chooseROM(t, o, "6")
	must(t, o.TypeLines("*DIR BCPL", "COPY /C TO SIEVE"))
	must(t, typeListing(o, sieveListing))
	must(t, o.Key("Escape"))
	must(t, o.WaitForPrompt(30))
	must(t, pictures.Screenshot(o, "bcpl-source"))
	from = o.Frames()
	must(t, o.TypeLines("BCPL SIEVE SIEVEC"))
	must(t, o.WaitForText("CINTCODE size", 300))
	must(t, o.WaitForPrompt(30))
	t.Logf("BCPL took %.1f seconds", float64(o.Frames()-from)/operator.FramesPerSecond)
	must(t, o.TypeLines("SIEVEC"))
	must(t, o.WaitForPrompt(30))
	must(t, pictures.Screenshot(o, "bcpl"))
	quitTo(t, o, "*QUIT")

	// ISO-Pascal: the program in the editor, left with Shift-f4, then
	// compiled and run
	chooseROM(t, o, "7")
	must(t, o.TypeLines("EDIT"))
	must(t, o.WaitForText("mark(s)", 30))
	must(t, o.WaitForKeyboard(10))
	must(t, typeText(o, triangleListing))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "pascal-editor"))
	functionKey(t, o, 1, "4")
	must(t, o.WaitForPrompt(30))
	from = o.Frames()
	must(t, o.TypeLines("COMPILE"))
	must(t, o.WaitForText("Code size", 300))
	must(t, o.WaitForPrompt(30))
	t.Logf("ISO-Pascal took %.1f seconds", float64(o.Frames()-from)/operator.FramesPerSecond)
	must(t, o.TypeLines("GO"))
	must(t, o.WaitForPrompt(30))
	must(t, pictures.Screenshot(o, "pascal"))
	quitTo(t, o, "*QUIT")

	// View: a title centred with a stored command, Shift-f8 and CE, and a
	// paragraph, then shown as it will be printed
	chooseROM(t, o, "9")
	must(t, pictures.Screenshot(o, "view-start"))
	must(t, o.TypeLines("NEW"))
	must(t, o.WaitForPrompt(30))
	must(t, o.Key("Escape"))
	must(t, o.WaitForKeyboard(10))
	functionKey(t, o, 1, "8")
	must(t, o.TypeLines("CE", "View on the Apple //e", ""))
	must(t, o.TypeLines("View was the word processor of the BBC Micro, in a ROM "+
		"of 16 KB. Here it runs on an Apple //e, the words that do not fit "+
		"carried to the next line and the spaces spread to make the lines even."))
	must(t, o.WaitForPrompt(30))
	must(t, pictures.Screenshot(o, "view-text"))
	must(t, o.Key("Escape"))
	must(t, o.WaitForPrompt(30))
	must(t, o.TypeLines("SCREEN"))
	must(t, o.WaitForText("lines even.", 30))
	must(t, o.WaitForPrompt(30))
	must(t, pictures.Screenshot(o, "view-screen"))

	// SCREEN shows the whole page, 66 lines, a screen at a time, and waits
	// for Shift, which on the Apple is Open Apple, to go on
	for i := 0; !o.HasText("=>") || strings.Contains(o.Text(), "SCREEN"); i++ {
		if i == 10 {
			t.Fatal("SCREEN did not end")
		}
		o.HoldButton(0)
		o.RunSeconds(1)
		o.ReleaseButton(0)
		o.RunSeconds(1)
	}
	must(t, o.WaitForKeyboard(10))
	quitTo(t, o, "*QUIT")

	// The seven languages at once, as sideways ROMs, and from one to another
	chooseROM(t, o, "8")
	must(t, o.TypeLines("*HELP"))
	must(t, o.WaitForPrompt(30))
	must(t, o.TypeLines("*COMAL"))
	must(t, o.WaitForPrompt(30))
	must(t, o.TypeLines("*BASIC"))
	must(t, o.WaitForPrompt(30))
	must(t, pictures.Screenshot(o, "sideways"))
}

// waitForApplecornMenu waits for the menu of the ROMs, all of it written
func waitForApplecornMenu(t *testing.T, o *operator.Operator) {
	t.Helper()
	must(t, o.WaitForText("9. Acornsoft View", 60))
	must(t, o.WaitForKeyboard(10))
}

// chooseROM presses the number of a ROM in the menu, and waits for its prompt
func chooseROM(t *testing.T, o *operator.Operator, number string) {
	t.Helper()
	must(t, o.Type(number))
	must(t, o.WaitForPrompt(60))
}

// quitTo leaves a language with its command for *QUIT, and goes back to the
// menu from the program selector of ProDOS
func quitTo(t *testing.T, o *operator.Operator, command string) {
	t.Helper()
	must(t, o.TypeLines(command))
	must(t, o.WaitForText("BITSY  BYE", 30))
	must(t, o.WaitForKeyboard(10))
	backToMenu(t, o)
}

// backToMenu starts APLCORN.SYSTEM, the second in the list of Bitsy Bye
func backToMenu(t *testing.T, o *operator.Operator) {
	t.Helper()
	must(t, o.Key("Down"))
	must(t, o.Key("Return"))
	waitForApplecornMenu(t, o)
}

/*
functionKey presses a function key of the BBC Micro: the number with an
Apple key held, Open Apple, button 0, for the key and Closed Apple, button 1,
for the key with Shift
*/
func functionKey(t *testing.T, o *operator.Operator, button int, number string) {
	t.Helper()
	o.HoldButton(button)
	o.Run(5)
	must(t, o.Type(number))
	o.Run(5)
	o.ReleaseButton(button)
}

/*
drawUntilPrompt records the machine as it draws in high resolution, until
BASIC writes its prompt, >, at the top left of the screen that MODE 2 left
black. Applecorn draws the text in the graphics, so there is no text to wait
for, and BASIC stops drawing for a while between the roses, to work out
their points, so a still screen is not the end either.
*/
func drawUntilPrompt(t *testing.T, o *operator.Operator, drawing *album.Recording, pictures *album.Album) {
	t.Helper()
	limit := o.Frames() + 60*60*5
	for o.InTextMode() || !promptDrawn(pictures.Screen(o)) {
		if o.Frames() > limit {
			t.Fatal("the drawing did not end")
		}
		drawing.Run(30, 6)
	}
	drawing.Run(60, 6)
}

// promptDrawn tells if there are dots lit in the place of the first character
// of the screen, a fortieth of it across and a twenty-fourth down
func promptDrawn(screen *image.RGBA) bool {
	size := screen.Bounds().Size()
	lit := 0
	for y := 0; y < size.Y/24; y++ {
		for x := 0; x < size.X/40; x++ {
			c := screen.RGBAAt(x, y)
			if c.R > 0x40 || c.G > 0x40 || c.B > 0x40 {
				lit++
			}
		}
	}
	return lit > 0
}
