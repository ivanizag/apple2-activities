package activities

import (
	"regexp"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// spiralProgram is the program of the Apple Pascal page, as typed in the
// editor
const spiralProgram = `(*$S+*)
program spiral;
uses turtlegraphics;
var i: integer;
begin
initturtle;
pencolor(white);
for i := 1 to 90 do
begin
move(i * 2);
turn(89)
end;
readln
end.
`

// pascalMachine is the machine of the guide, as its command line of izapple2
const pascalMachine = `izapple2 -model none -board 2e -cpu 65c02 -screen green \
    -rom "<internal>/Apple2e_Enhanced.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s6 'diskii,disk1=<internal>/Apple II Pascal 1.3 APPLE1_ 680-0283-A.dsk,disk2=<internal>/Apple II Pascal 1.3 APPLE2_ 680-0284-A.dsk'`

/*
pascalScreenshots is Apple Pascal 1.3 on an enhanced Apple //e with two
drives: the system started, a disk listed in the Filer, a program written in
the editor, compiled and run, drawing with the turtle.
*/
func pascalScreenshots(t *testing.T) {
	pictures := newAlbum("pascal", album.Green)
	o := start(t, pascalMachine, nil)

	// Started, at the command line
	must(t, o.WaitForText("[1.3]", 120))
	must(t, o.WaitForKeyboard(10))
	must(t, pictures.Screenshot(o, "started"))

	// The boot disk listed in the Filer
	must(t, o.Type("F"))
	must(t, o.WaitForText("Filer:", 30))
	must(t, o.Type("L"))
	must(t, o.WaitForKeyboard(10))
	must(t, o.Type("*\n"))
	must(t, o.WaitForText("files <listed/in dir>", 30))
	o.Run(30)
	must(t, pictures.Screenshot(o, "filer"))
	must(t, o.Type("Q"))
	must(t, o.WaitForText("Command:", 30))

	// The program written in the editor
	must(t, o.Type("E"))
	must(t, o.WaitForText("Edit what file?", 30))
	must(t, o.Key("Return"))
	must(t, o.WaitForText(">Edit:", 30))
	must(t, o.Type("I"))
	o.Run(30)
	must(t, o.Type(spiralProgram))
	must(t, o.Key("Ctrl+C"))
	o.Run(60)
	must(t, pictures.Screenshot(o, "editor"))

	// Kept in the work file, and back to the command line
	must(t, o.Type("Q"))
	must(t, o.WaitForText("C(hange", 30))
	must(t, o.WaitForKeyboard(10))
	must(t, pictures.Screenshot(o, "quit"))
	must(t, o.Type("U"))
	must(t, o.WaitForText("bytes long", 60))
	must(t, o.Type("E"))
	must(t, o.WaitForText("Command:", 30))

	// Compiled and run, the drawing recorded
	must(t, o.Type("R"))
	must(t, o.WaitForText("Listing file", 60))
	must(t, o.Key("Return"))
	if !o.WaitUntil(120, func() bool { return compiledSpace.MatchString(o.Text()) }) {
		t.Fatalf("the program was not compiled:\n%v", o.Text())
	}
	must(t, pictures.Screenshot(o, "compiled"))
	must(t, o.WaitForText("Running...", 30))
	drawing := pictures.On(album.Color).Record(o)
	drawing.Capture(10)
	drawing.Run(25*60, 6)
	must(t, pictures.On(album.Color).SaveRecording(drawing, "drawing", 300))
}

// compiledSpace is the last line the compiler writes, with its number
var compiledSpace = regexp.MustCompile(`Smallest available space = [0-9]+ words`)
