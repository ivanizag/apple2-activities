package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// mousePaintMachine is the machine of the guide, as its command line of
// izapple2
const mousePaintMachine = `izapple2 -model none -board 2c -cpu 65c02 -screen green \
    -rom "<internal>/Apple2c_ROM255.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s1 iicserial \
    -s2 iicserial \
    -s6 diskii,disk1=disks/077_MousePaint_Drawing_Program_68C-3005-A_1984.dsk`

/*
mousePaintScreenshots is MousePaint on the Apple IIc, with its mouse: the
picture of the mouse on its disk, and a drawing of a house made with its
tools and patterns, a line of text in one of its fonts, and saved on the disk
*/
func mousePaintScreenshots(t *testing.T) {
	pictures := newAlbum("mousepaint", album.Green)
	o := start(t, mousePaintMachine, nil)
	mp := &mousePaint{t: t, o: o, pictures: pictures}

	// Its first screen, which offers a lesson on the mouse
	must(t, o.WaitForText("directly to MousePaint", 30))
	o.RunSeconds(1)
	must(t, pictures.Screenshot(o, "title"))

	// A click, and MousePaint, with the pointer up in a corner, where the
	// one of izapple2 and the one of MousePaint meet
	mp.click()
	o.RunSeconds(12)
	o.MouseTo(0, 0)
	o.RunSeconds(1)
	must(t, pictures.Screenshot(o, "desktop"))

	// File, Get picture: MOUSE.PIC, on the disk, and Aids, Show Page, the
	// whole of it
	mp.menu(mpFile, 20)
	o.RunSeconds(2)
	must(t, o.Type("MOUSE.PIC\n"))
	o.RunSeconds(5)
	must(t, pictures.Screenshot(o, "mouse-pic"))
	mp.menu(mpAids, 31)
	o.RunSeconds(5)
	must(t, pictures.Screenshot(o, "show-page"))
	mp.clickAt(60, 60)
	o.RunSeconds(3)

	// File, Blank HiRes Screen: a page to draw on, and MOUSE.PIC not saved
	mp.menu(mpFile, 32)
	o.RunSeconds(1)
	must(t, o.Type("N\n"))
	o.RunSeconds(3)
	must(t, pictures.Screenshot(o, "blank"))

	// The drawing, recorded
	mp.rec = pictures.Record(o)
	mp.rec.Capture(100)

	// The ground, a line across
	mp.tool(mpLine)
	mp.drag(52, 140, 276, 140)

	// The walls, filled with bricks, two windows, white, and a door, black
	mp.pattern(237, 173)
	mp.tool(mpFilledRectangle)
	mp.drag(100, 95, 180, 140)
	mp.pattern(68, 183)
	mp.drag(108, 104, 124, 118)
	mp.drag(156, 104, 172, 118)
	mp.pattern(68, 173)
	mp.drag(132, 114, 148, 140)

	// The roof, a polygon filled with tiles: a click at each corner, and a
	// double click at the last
	mp.pattern(167, 173)
	mp.tool(mpFilledPolygon)
	mp.clickAt(94, 95)
	mp.clickAt(140, 58)
	mp.clickAt(186, 95)
	mp.doubleClickAt(94, 95)

	// The sun, an oval filled with dots
	mp.pattern(83, 183)
	mp.tool(mpFilledOval)
	mp.drag(228, 36, 256, 60)

	// Grass, with the spray can
	mp.tool(mpSpray)
	for y := 144; y <= 156; y += 6 {
		mp.drag(54, y, 274, y)
	}

	// A line of text, in Venice, which MousePaint loads from the disk
	mp.tool(mpText)
	mp.menu(mpFonts, 67)
	mp.run(300)
	mp.clickAt(70, 60)
	must(t, mp.rec.Type("Home, sweet home"))
	mp.run(300)
	mp.at(278, 168)
	mp.rec.Run(60, 6)
	must(t, pictures.SaveRecording(mp.rec, "drawing", 300))
	mp.rec = nil
	must(t, pictures.Screenshot(o, "home"))

	// File, Put a copy in: saved on the disk, as HOME
	mp.menu(mpFile, 55)
	o.RunSeconds(2)
	must(t, o.Type("HOME\n"))
	o.RunSeconds(5)
	must(t, pictures.Screenshot(o, "saved"))
}

// Places of the menus and the tools of MousePaint, on the high resolution
// screen, 280 by 192
const (
	mpFile  = 40
	mpAids  = 103
	mpFonts = 140

	mpText            = 1
	mpSpray           = 2
	mpLine            = 3
	mpFilledRectangle = 4
	mpFilledOval      = 5
	mpFilledPolygon   = 6
)

// mpTools are where the tools are, in the two columns at the left
var mpTools = map[int][2]int{
	mpText:            {35, 42},
	mpSpray:           {14, 57},
	mpLine:            {14, 70},
	mpFilledRectangle: {35, 85},
	mpFilledOval:      {35, 112},
	mpFilledPolygon:   {35, 138},
}

// mousePaint is the person at MousePaint
type mousePaint struct {
	t        *testing.T
	o        *operator.Operator
	pictures *album.Album
	rec      *album.Recording
}

// run runs the machine, and the recording if any
func (m *mousePaint) run(frames int) {
	if m.rec != nil {
		m.rec.Run(frames, 4)
	} else {
		m.o.Run(frames)
	}
}

// at moves the mouse to a dot of the high resolution screen, 280 by 192:
// MousePaint counts a dot for each step of the mouse of the IIc, which
// izapple2 makes 560 steps across the window
func (m *mousePaint) at(x int, y int) {
	fx, fy := (float64(x)+0.5)/560, (float64(y)+0.5)/192
	if m.rec != nil {
		m.rec.Glide(fx, fy)
	} else {
		m.o.MouseTo(fx, fy)
	}
	m.run(20)
}

// click presses the button long enough for MousePaint to see it
func (m *mousePaint) click() {
	m.o.Hold()
	m.run(20)
	m.o.Release()
	m.run(20)
}

// clickAt clicks at a dot
func (m *mousePaint) clickAt(x int, y int) {
	m.at(x, y)
	m.click()
}

// doubleClickAt clicks twice at a dot, quickly
func (m *mousePaint) doubleClickAt(x int, y int) {
	m.at(x, y)
	m.o.Hold()
	m.run(6)
	m.o.Release()
	m.run(6)
	m.o.Hold()
	m.run(6)
	m.o.Release()
	m.run(20)
}

// tool chooses a tool
func (m *mousePaint) tool(tool int) {
	m.clickAt(mpTools[tool][0], mpTools[tool][1])
}

// pattern chooses a pattern, at a dot of the patterns at the bottom
func (m *mousePaint) pattern(x int, y int) {
	m.clickAt(x, y)
}

// drag holds the button from a place to another, a little at a time
func (m *mousePaint) drag(x0 int, y0 int, x1 int, y1 int) {
	m.at(x0, y0)
	m.o.Hold()
	m.run(10)
	m.at(x1, y1)
	m.o.Release()
	m.run(20)
}

/*
menu chooses an item of a menu of the bar, at a height. The pointer goes up
against the top of the screen first: the pointer of MousePaint and the one of
izapple2 drift apart when MousePaint keeps the pointer in its window, as it
draws, and they meet again on an edge.
*/
func (m *mousePaint) menu(x int, y int) {
	m.at(x, -10)
	m.run(30)
	m.at(x, 5)
	m.o.Hold()
	m.run(20)
	m.at(x, y)
	m.o.Release()
	m.run(20)
}
