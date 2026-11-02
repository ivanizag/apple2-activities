package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// dazzleDrawMachine is the machine of the guide, as its command line of
// izapple2: an Apple IIc with Dazzle Draw in its drive, and the other side of
// its disk, with its pictures, in a second drive
const dazzleDrawMachine = `izapple2 -model none -board 2c -cpu 65c02 -screen color \
    -rom "<internal>/Apple2c_ROM255.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s1 iicserial \
    -s2 iicserial \
    -s6 "diskii,disk1=disks/Dazzle Draw 1.2.dsk,disk2=disks/Dazzle Draw Slide Show.DSK"`

/*
dazzleDrawScreenshots is Dazzle Draw on the Apple IIc, painting in the
sixteen colours of double high resolution with the mouse: the pictures on
the back of its disk, and a picture of its own
*/
func dazzleDrawScreenshots(t *testing.T) {
	pictures := newAlbum("dazzle-draw", album.Color)
	o := start(t, dazzleDrawMachine, nil)
	dd := &dazzleDraw{t: t, o: o, pictures: pictures}

	// The title, and the start, with the mouse
	o.RunSeconds(14)
	must(t, pictures.Screenshot(o, "title"))
	o.RunSeconds(11)
	must(t, pictures.Screenshot(o, "start"))
	must(t, o.Key("Return"))
	o.RunSeconds(20)
	o.MouseTo(0, 0)
	o.RunSeconds(1)
	must(t, pictures.Screenshot(o, "canvas"))

	// File, Load, a picture, of the disk in drive 2: On Line, Prefix, the
	// other side, Okay, and its name
	dd.menu(ddFile, 17)
	o.RunSeconds(3)
	dd.clickAt(237, 63)
	o.RunSeconds(3)
	must(t, pictures.Screenshot(o, "load"))
	dd.clickAt(237, 75)
	o.RunSeconds(3)
	dd.clickAt(237, 64)
	o.RunSeconds(3)
	dd.clickAt(57, 75)
	o.RunSeconds(2)
	dd.clickAt(237, 86)
	o.RunSeconds(3)
	must(t, o.Type("MONARCH"))
	o.RunSeconds(1)
	must(t, pictures.Screenshot(o, "monarch-named"))
	dd.clickAt(237, 86)
	o.RunSeconds(8)
	must(t, pictures.Screenshot(o, "monarch"))

	// Goodies, View Picture: the whole screen, without the menus
	dd.menu(ddGoodies, 61)
	o.RunSeconds(3)
	must(t, pictures.Screenshot(o, "monarch-view"))
	dd.clickAt(140, 100)
	o.RunSeconds(3)

	// Another picture, the prefix kept
	dd.menu(ddFile, 17)
	o.RunSeconds(3)
	dd.clickAt(237, 63)
	o.RunSeconds(3)
	must(t, o.Type("ROOM"))
	dd.clickAt(237, 86)
	o.RunSeconds(8)
	dd.menu(ddGoodies, 61)
	o.RunSeconds(3)
	must(t, pictures.Screenshot(o, "room-view"))
	dd.clickAt(140, 100)
	o.RunSeconds(3)

	// Goodies, Clear Picture, and Yes
	dd.menu(ddGoodies, 72)
	o.RunSeconds(3)
	must(t, pictures.Screenshot(o, "clear"))
	dd.clickAt(233, 85)
	o.RunSeconds(3)

	// A picture of its own, recorded
	dd.rec = pictures.Record(o)
	dd.rec.Capture(100)

	// The sky, the whole of the black filled with light blue
	dd.tool(ddFloodFill)
	dd.colour(ddLightBlue)
	dd.clickAt(140, 60)

	// The ground, a rectangle filled with green, and the sun, an oval
	dd.tool(ddShapes)
	dd.colour(ddGreen)
	dd.clickAt(ddFilledRectangle[0], ddFilledRectangle[1])
	dd.drag(2, 120, 278, 163)
	dd.colour(ddYellow)
	dd.clickAt(ddFilledOval[0], ddFilledOval[1])
	dd.drag(200, 22, 240, 52)

	// The walls, orange, the door, olive, and two windows, white
	dd.colour(ddOrange)
	dd.clickAt(ddFilledRectangle[0], ddFilledRectangle[1])
	dd.drag(60, 80, 130, 125)
	dd.colour(ddOlive)
	dd.drag(88, 100, 102, 125)
	dd.colour(ddWhite)
	dd.drag(68, 90, 80, 102)
	dd.drag(110, 90, 122, 102)

	// The roof, three lines in red, and filled
	dd.tool(ddLines)
	dd.colour(ddRed)
	dd.drag(55, 80, 95, 55)
	dd.drag(95, 55, 135, 80)
	dd.drag(55, 80, 135, 80)
	dd.tool(ddFloodFill)
	dd.colour(ddRed)
	dd.clickAt(95, 72)

	// Flowers, sprayed in magenta
	dd.tool(ddSprayPaint)
	dd.colour(ddMagenta)
	dd.drag(10, 140, 270, 150)

	// A line of text
	dd.tool(ddText)
	dd.clickAt(20, 35)
	must(t, dd.rec.Type("Dazzle Draw, 1985"))

	dd.run(60)
	must(t, pictures.SaveRecording(dd.rec, "painting", 300))
	dd.rec = nil
	must(t, pictures.Screenshot(o, "painted"))

	// Zoom, on the door
	dd.tool(ddZoom)
	dd.clickAt(95, 105)
	o.RunSeconds(3)
	must(t, pictures.Screenshot(o, "zoom"))

	// View Picture
	dd.menu(ddGoodies, 61)
	o.RunSeconds(3)
	must(t, pictures.Screenshot(o, "painted-view"))
	dd.clickAt(140, 100)
	o.RunSeconds(3)
}

// Places of Dazzle Draw, on the screen as it counts it, 280 by 192
const (
	ddFile    = 37
	ddTools   = 82
	ddEdit    = 120
	ddGoodies = 164
)

// The tools, by their height in the menu Tools
const (
	ddPaintBrush = 17
	ddSprayPaint = 28
	ddFloodFill  = 39
	ddZoom       = 50
	ddText       = 61
	ddShapes     = 72
	ddLines      = 83
)

// The sixteen colours, where they are at the bottom left, two rows of eight
var (
	ddBlack     = [2]int{19, 181}
	ddMagenta   = [2]int{62, 181}
	ddRed       = [2]int{77, 181}
	ddOrange    = [2]int{91, 181}
	ddYellow    = [2]int{119, 181}
	ddOlive     = [2]int{48, 187}
	ddGreen     = [2]int{62, 187}
	ddLightBlue = [2]int{91, 187}
	ddWhite     = [2]int{119, 187}
)

// The shapes, at the bottom, filled with the colour
var (
	ddFilledRectangle = [2]int{151, 173}
	ddFilledOval      = [2]int{151, 186}
)

// dazzleDraw is the person at Dazzle Draw
type dazzleDraw struct {
	t        *testing.T
	o        *operator.Operator
	pictures *album.Album
	rec      *album.Recording
}

// run runs the machine, and the recording if any
func (d *dazzleDraw) run(frames int) {
	if d.rec != nil {
		d.rec.Run(frames, 4)
	} else {
		d.o.Run(frames)
	}
}

/*
at moves the mouse to a place, 280 by 192: Dazzle Draw counts a place for
each step of the mouse of the IIc, which izapple2 makes 560 steps across the
window
*/
func (d *dazzleDraw) at(x int, y int) {
	fx, fy := (float64(x)+0.5)/560, (float64(y)+0.5)/192
	if d.rec != nil {
		d.rec.Glide(fx, fy)
	} else {
		d.o.MouseTo(fx, fy)
	}
	d.run(20)
}

// clickAt clicks at a place
func (d *dazzleDraw) clickAt(x int, y int) {
	d.at(x, y)
	d.o.Hold()
	d.run(20)
	d.o.Release()
	d.run(20)
}

// drag holds the button from a place to another
func (d *dazzleDraw) drag(x0 int, y0 int, x1 int, y1 int) {
	d.at(x0, y0)
	d.o.Hold()
	d.run(10)
	d.at(x1, y1)
	d.o.Release()
	d.run(20)
}

// still runs until the screen stays the same for a second, but for a cursor
// that blinks, as when Flood Fill has filled what it fills, a line at a time
func (d *dazzleDraw) still() {
	d.t.Helper()
	last := d.pictures.Screen(d.o)
	for quiet, waited := 0, 0; quiet < 4; waited++ {
		if waited > 4*60 {
			d.t.Fatal("the screen of Dazzle Draw did not stay still")
		}
		d.run(15)
		screen := d.pictures.Screen(d.o)
		if changedDots(screen, last) < 1000 {
			quiet++
		} else {
			quiet = 0
		}
		last = screen
	}
}

// tool chooses a tool in the menu Tools
func (d *dazzleDraw) tool(y int) {
	d.menu(ddTools, y)
	d.run(60)
}

// colour chooses a colour
func (d *dazzleDraw) colour(c [2]int) {
	d.clickAt(c[0], c[1])
}

/*
menu chooses an item of a menu of the bar, at a height. The pointer goes up
against the top of the screen first, where the pointer of Dazzle Draw and
the one of izapple2 meet if they drifted apart
*/
func (d *dazzleDraw) menu(x int, y int) {
	d.still()
	d.at(x, -10)
	d.run(30)
	d.at(x, 5)
	d.o.Hold()
	d.run(20)
	d.at(x, y)
	d.o.Release()
	d.run(20)
}
