package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// deskTopMachine is the machine of the guide, as its command line of izapple2
const deskTopMachine = `izapple2 -model none -board 2e -cpu 65c02 -screen color \
    -rom "<internal>/Apple2e_Enhanced.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" -rgb \
    -s0 language \
    -s4 mouse \
    -s7 smartport,image1=disks/A2DeskTop-1.4-en_800k.2mg`

/*
deskTopScreenshots is Apple II DeskTop on an enhanced Apple //e with a mouse:
the desktop, a disk opened, a file read, the machine described, and the
Calculator.
*/
func deskTopScreenshots(t *testing.T) {
	pictures := newAlbum("desktop", album.Color)
	o := start(t, deskTopMachine, nil)

	// The desktop, with the pointer out of the way
	dot(o, 280, 120)
	o.Run(6 * 60)
	must(t, pictures.Screenshot(o, "desktop"))

	// The disk opened with a double click
	opening := pictures.Record(o)
	opening.Capture(50)
	opening.Glide(atDot(517, 26))
	opening.Run(10, 2)
	opening.DoubleClick()
	opening.Run(180, 3)
	must(t, pictures.SaveRecording(opening, "open-disk", 300))

	// The Read.Me file, read, and closed with Escape
	dot(o, 194, 55)
	o.Run(20)
	o.DoubleClick()
	o.Run(5 * 60)
	must(t, pictures.Screenshot(o, "read-me"))
	must(t, o.Key("Escape"))
	o.Run(2 * 60)

	// About This Apple II, from the Apple menu, with the menu recorded
	dot(o, 14, 5)
	o.Run(30)
	menu := pictures.Record(o)
	menu.Capture(50)
	menu.Press(true)
	menu.Run(20, 2)
	menu.Glide(atDot(20, 29))
	menu.Run(30, 2)
	must(t, pictures.SaveRecording(menu, "apple-menu", 200))
	o.Release()
	o.Run(5 * 60)
	must(t, pictures.Screenshot(o, "about"))
	must(t, o.Key("Escape"))
	o.Run(2 * 60)

	// The folder of samples, and a picture in it opened, each chosen by
	// typing the start of its name and opened with Open Apple and O
	openByName(t, o, "SAMPLE")
	must(t, pictures.Screenshot(o, "sample-media"))
	openByName(t, o, "MONARCH")
	o.Run(3 * 60)
	must(t, pictures.Screenshot(o, "monarch"))
	must(t, o.Key("Escape"))
	o.Run(5 * 60)
	closeWindow(t, o)

	// The screen savers, from the Apple menu, and the flying toasters until
	// a key is pressed
	chooseFromAppleMenu(o, 64)
	must(t, pictures.Screenshot(o, "screen-savers"))
	must(t, o.Type("FLY"))
	o.Run(30)
	toasters := pictures.Record(o)
	toasters.Capture(50)
	appleKey(t, o, "O")
	toasters.Run(12*60, 6)
	must(t, pictures.SaveRecording(toasters, "toasters", 100))
	must(t, o.Key("Escape"))
	o.Run(3 * 60)
	closeWindow(t, o)

	// The Calculator, and twelve times three worked out on its keys
	chooseFromAppleMenu(o, 84)
	calculating := pictures.Record(o)
	calculating.Capture(50)
	for _, key := range [][2]int{{232, 132}, {261, 132}, {318, 87}, {290, 132}, {290, 87}} {
		calculating.Glide(atDot(key[0], key[1]))
		calculating.Click()
	}
	calculating.Run(30, 6)
	must(t, pictures.SaveRecording(calculating, "calculator", 300))
}

// closeWindow closes the front window with Open Apple and W
func closeWindow(t *testing.T, o *operator.Operator) {
	t.Helper()
	appleKey(t, o, "W")
	o.Run(3 * 60)
}

// openByName selects an icon of the front window by typing the start of its
// name, and opens it with Open Apple and O
func openByName(t *testing.T, o *operator.Operator, name string) {
	t.Helper()
	must(t, o.Type(name))
	o.Run(30)
	appleKey(t, o, "O")
	o.Run(5 * 60)
}

// appleKey types a key with the Open Apple key, button 0, held down
func appleKey(t *testing.T, o *operator.Operator, key string) {
	t.Helper()
	o.HoldButton(0)
	o.Run(5)
	must(t, o.Type(key))
	o.ReleaseButton(0)
}
// The screen of Apple II DeskTop, in the dots of the 80 column graphics
const (
	deskTopWidth  = 560
	deskTopHeight = 192
)

// atDot is a place of the screen of Apple II DeskTop, by its dots, as a part
// of the screen
func atDot(x int, y int) (float64, float64) {
	return (float64(x) + 0.5) / deskTopWidth, (float64(y) + 0.5) / deskTopHeight
}

// dot moves the mouse to a place of the screen of Apple II DeskTop
func dot(o *operator.Operator, x int, y int) {
	o.MouseTo(atDot(x, y))
}

// chooseFromAppleMenu pulls down the Apple menu and chooses the item at a
// height, going down the menu a little at a time as a hand does
func chooseFromAppleMenu(o *operator.Operator, y int) {
	dot(o, 14, 5)
	o.Hold()
	o.Run(20)
	for down := 5; down < y; down += 4 {
		dot(o, 20, down)
	}
	dot(o, 20, y)
	o.Run(20)
	o.Release()
	o.Run(5 * 60)
}
