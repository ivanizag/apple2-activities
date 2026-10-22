package activities

import (
	"image"
	"math"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// fantavisionMatinee is the machine of the guide with the demonstrations of
// Fantavision, as its command line of izapple2
const fantavisionMatinee = `izapple2 -model none -board 2e -cpu 65c02 -screen color \
    -rom "<internal>/Apple2e_Enhanced.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language -s4 mouse \
    -s6 'diskii,disk1=disks/Fantavision side B - demos.woz,disk2=my-movies.po'`

// fantavisionMachine is the machine of the guide with the program of
// Fantavision, as its command line of izapple2
const fantavisionMachine = `izapple2 -model none -board 2e -cpu 65c02 -screen color \
    -rom "<internal>/Apple2e_Enhanced.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language -s4 mouse \
    -s6 'diskii,disk1=disks/Fantavision side A - program.woz,disk2=my-movies.po'`

/*
fantavisionScreenshots is Fantavision of 1985 on an enhanced Apple //e with a
mouse: the Movie Matinee of the back of its disk, a show of movies made with
it, and then the program, with three movies made and played, a star that
turns into a circle and a square, a title that spins, and a triangle that
leaves its trace as it turns, saved on a diskette of ProDOS and listed.

	cp disks/ProDOS_2_4_3.po my-movies.po
*/
func fantavisionScreenshots(t *testing.T) {
	pictures := newAlbum("fantavision", album.Color)
	files := map[string]string{"my-movies.po": disk(t, "ProDOS_2_4_3.po")}

	// The Movie Matinee: its title while the movies load, and the movies one
	// after the other, the same on every run, each recorded by its seconds
	// from the end of the title
	o := start(t, fantavisionMatinee, files)
	waitForScreen(t, o, pictures, matineeTitle)
	must(t, pictures.Screenshot(o, "matinee"))
	waitForScreen(t, o, pictures, func(s *image.RGBA) bool { return !matineeTitle(s) })
	shown := o.Frames()
	for _, movie := range []struct {
		name     string
		from, to float64
	}{
		{"welcome", 16.75, 31},
		{"faces", 33, 51},
		{"fish", 68, 85},
		{"ship", 177.75, 197.75},
		{"eye", 198.5, 218.5},
	} {
		o.Run(int(shown + uint64(movie.from*operator.FramesPerSecond) - o.Frames()))
		show := pictures.Record(o)
		show.Capture(6)
		show.Run(int((movie.to-movie.from)*operator.FramesPerSecond), 6)
		must(t, pictures.SaveRecording(show, movie.name, 200))
	}

	// The program: its title, which says it found the mouse, and the screen
	// to draw on
	o = start(t, fantavisionMachine, files)
	waitForScreen(t, o, pictures, fantavisionTitle)
	must(t, pictures.Screenshot(o, "title"))
	waitForScreen(t, o, pictures, fantavisionEditor)
	o.MouseTo(fvMiddle.mouse())
	o.RunSeconds(1)
	must(t, pictures.Screenshot(o, "editor"))

	// The first movie, in palette 3 on black: an orange star, drawn a point
	// at a time and closed on the first
	fvClick(o, fvPaletteDown)
	fvClick(o, fvPaletteDown)
	fvClick(o, fvBackground)
	fvClick(o, fvColour(0, 0))
	fvClick(o, fvObject(1))
	fvClick(o, fvColour(5, 0))
	fvClick(o, fvDrawTool)
	star := pictures.Record(o)
	star.Capture(10)
	for i := 0; i <= 10; i++ {
		radius := 40.0
		if i%2 == 1 {
			radius = 16
		}
		angle := -math.Pi/2 + float64(i%10)*math.Pi/5
		star.Glide(fvPlace{80 + 1.2*radius*math.Cos(angle), 60 + radius*math.Sin(angle)}.mouse())
		star.Click()
	}
	star.Run(60, 6)
	must(t, pictures.SaveRecording(star, "star", 200))

	// Frame 2, a blue circle in the middle, and frame 3, a white square
	fvClick(o, fvNextFrame)
	fvClick(o, fvCircleTool)
	fvDrag(o, fvPlace{142, 85}, fvPlace{142, 30})
	fvClick(o, fvColour(2, 0))
	o.MouseTo(fvMiddle.mouse())
	o.RunSeconds(1)
	must(t, pictures.Screenshot(o, "circle"))
	fvClick(o, fvNextFrame)
	fvClick(o, fvRectangleTool)
	fvDrag(o, fvPlace{180, 90}, fvPlace{240, 150})
	fvClick(o, fvColour(6, 1))
	o.MouseTo(fvMiddle.mouse())
	o.RunSeconds(1)
	must(t, pictures.Screenshot(o, "square"))

	// GO plays it from frame 1, the first of the film strip, to 3 and back,
	// the frames between drawn by the program
	fvClick(o, fvFirstFrame)
	fvClick(o, fvGo)
	o.Run(30)
	shapes := pictures.Record(o)
	shapes.Capture(6)
	shapes.Run(10*operator.FramesPerSecond, 6)
	must(t, pictures.SaveRecording(shapes, "shapes", 100))
	fvClick(o, fvMiddle)
	o.RunSeconds(1)

	// Saved on the diskette in drive 2, which ",2" after the name asks
	// for, and cleared
	fvSave(o, "SHAPES,2")
	fvClear(o)

	// The second movie, in palette 4 on black: a title, a word in magenta
	// and a number in green, four letters to an object
	fvClick(o, fvPaletteDown)
	fvClick(o, fvBackground)
	fvClick(o, fvColour(0, 0))
	fvClick(o, fvObject(1))
	fvClick(o, fvTextTool)
	fvClick(o, fvPlace{82, 80})
	o.RunSeconds(1)
	for _, key := range []string{"A", "P", "P", "L", "E", "Right", "I", "I", "Return"} {
		if len(key) == 1 {
			must(t, o.Type(key))
		} else {
			must(t, o.Key(key))
		}
		o.Run(20)
	}
	for _, object := range []struct {
		number int
		colour fvPlace
	}{{1, fvColour(2, 0)}, {2, fvColour(2, 0)}, {3, fvColour(5, 0)}} {
		fvClick(o, fvObject(object.number))
		fvClick(o, object.colour)
	}
	o.MouseTo(fvMiddle.mouse())
	o.RunSeconds(1)
	must(t, pictures.Screenshot(o, "title-frame"))

	// Three more frames, each a clone of the last, turned a quarter, four
	// turns of the Goodies, and made bigger or smaller again by the zoom
	for frame := 2; frame <= 4; frame++ {
		fvMenu(o, fvEdit, 4) // Clone
		fvMenu(o, fvGoodies, 1)
		for range 4 {
			fvClick(o, fvGoodiesMore)
		}
		fvMenu(o, fvGoodies, 0)
		if frame == 3 {
			fvClick(o, fvGoodiesLess)
		} else {
			fvClick(o, fvGoodiesMore)
		}
	}
	o.MouseTo(fvMiddle.mouse())
	o.RunSeconds(1)
	must(t, pictures.Screenshot(o, "title-turned"))
	fvClick(o, fvFirstFrame)
	fvClick(o, fvGo)
	o.Run(30)
	spin := pictures.Record(o)
	spin.Capture(6)
	spin.Run(10*operator.FramesPerSecond, 6)
	must(t, pictures.SaveRecording(spin, "spin", 100))
	fvClick(o, fvMiddle)
	o.RunSeconds(1)
	fvSave(o, "SPIN,2")
	fvClear(o)

	// The third movie: a small triangle of lines in trace mode, which leaves
	// behind each frame it draws, turned a sixteenth and made bigger in each
	// of eleven clones
	fvClick(o, fvBackground)
	fvClick(o, fvColour(0, 0))
	fvClick(o, fvObject(1))
	fvClick(o, fvLines)
	fvClick(o, fvTrace)
	fvClick(o, fvColour(5, 0))
	fvClick(o, fvDrawTool)
	for _, p := range []fvPlace{{142, 70}, {160, 100}, {124, 100}, {142, 70}} {
		fvClick(o, p)
	}
	for frame := 2; frame <= 12; frame++ {
		fvMenu(o, fvEdit, 4) // Clone
		fvMenu(o, fvGoodies, 1)
		fvClick(o, fvGoodiesMore)
		fvMenu(o, fvGoodies, 0)
		fvClick(o, fvGoodiesMore)
		if frame%2 == 0 {
			fvClick(o, fvColour(2, 0))
		} else {
			fvClick(o, fvColour(5, 0))
		}
	}

	// GO from frame 1 at speed 2, from 1 to 7, the fastest with the fewest
	// frames in between
	fvClick(o, fvFirstFrame)
	fvClick(o, fvGo)
	o.Run(10)
	must(t, o.Type("2"))
	trace := pictures.Record(o)
	trace.Capture(6)
	trace.Run(6*operator.FramesPerSecond, 6)
	must(t, pictures.SaveRecording(trace, "trace", 300))
	fvClick(o, fvMiddle)
	o.RunSeconds(1)
	fvSave(o, "TRACE,2")

	// The movies of drive 2, in Load Movie
	fvMenu(o, fvFile, 0)
	o.RunSeconds(2)
	fvClick(o, fvDrive2)
	o.RunSeconds(4)
	must(t, pictures.Screenshot(o, "load"))
}

// fvSave saves the movie with a name, and waits for the drive
func fvSave(o *operator.Operator, name string) {
	fvMenu(o, fvFile, 1)
	o.RunSeconds(1)
	_ = o.Type(name)
	_ = o.Key("Return")
	o.RunSeconds(3)
}

// fvClear clears the movie, saying yes
func fvClear(o *operator.Operator) {
	fvMenu(o, fvFile, 2)
	o.RunSeconds(1)
	fvClick(o, fvYes)
	o.RunSeconds(1)
}

/*
The places of Fantavision's screen, in dots of the high resolution, 280
across and 192 down: the tools down the left, the palette, GO, the object
counter and the animation and dimension boxes along the bottom, the film strip
down the right and the menus along the top
*/
type fvPlace struct{ x, y float64 }

var (
	fvMiddle        = fvPlace{140, 85}
	fvDrawTool      = fvPlace{8, 18}
	fvCircleTool    = fvPlace{8, 96}
	fvRectangleTool = fvPlace{8, 112}
	fvTextTool      = fvPlace{8, 132}
	fvPaletteDown   = fvPlace{8, 183}
	fvGo            = fvPlace{163, 167}
	fvBackground    = fvPlace{200, 167} // the French window of the object counter
	fvTrace         = fvPlace{232, 167}
	fvLines         = fvPlace{255, 176}
	fvNextFrame     = fvPlace{273, 152}
	fvFirstFrame    = fvPlace{273, 22}
	fvGoodiesLess   = fvPlace{197, 4}
	fvGoodiesMore   = fvPlace{225, 4}
	fvYes           = fvPlace{121, 97}
	fvDrive2        = fvPlace{99, 30} // in Load Movie
)

// The menus, across the top, and their items, 12 dots apart from the first
const (
	fvFile    = 40
	fvEdit    = 80
	fvGoodies = 135
)

// mouse is the place as the mouse gives it, a part of the screen
func (p fvPlace) mouse() (float64, float64) {
	return p.x / 280, p.y / 192
}

// fvColour is a colour of the palette, seven across in two rows
func fvColour(column int, row int) fvPlace {
	return fvPlace{26 + 14*float64(column), 168 + 14*float64(row)}
}

// fvObject is an object number of the object counter, 1 to 8, three in a
// row from the bottom as on a calculator
func fvObject(number int) fvPlace {
	return fvPlace{180 + 10*float64((number-1)%3), 185 - 9*float64((number-1)/3)}
}

// fvClick clicks at a place of the screen
func fvClick(o *operator.Operator, p fvPlace) {
	o.MouseTo(p.mouse())
	o.Click()
}

// fvDrag holds the button down from a place to another, as a circle or a
// rectangle is made
func fvDrag(o *operator.Operator, from fvPlace, to fvPlace) {
	o.MouseTo(from.mouse())
	o.Hold()
	o.Run(10)
	for i := 1; i <= 8; i++ {
		f := float64(i) / 8
		o.MouseTo(fvPlace{from.x + (to.x-from.x)*f, from.y + (to.y-from.y)*f}.mouse())
	}
	o.Run(10)
	o.Release()
	o.Run(30)
}

// fvMenu pulls down a menu and lets go of the button on an item, numbered
// from 0
func fvMenu(o *operator.Operator, menu float64, item int) {
	o.MouseTo(fvPlace{menu, 4}.mouse())
	o.Hold()
	o.Run(10)
	o.MouseTo(fvPlace{menu, 16 + 12*float64(item)}.mouse())
	o.Run(10)
	o.Release()
	o.Run(30)
}

// waitForScreen runs the machine until its screen is as asked, looking at it
// every half second
func waitForScreen(t *testing.T, o *operator.Operator, pictures *album.Album, is func(*image.RGBA) bool) {
	t.Helper()
	limit := o.Frames() + 120*operator.FramesPerSecond
	for !is(pictures.Screen(o)) {
		if o.Frames() > limit {
			t.Fatal("the screen did not come")
		}
		o.Run(30)
	}
}

// dotAt is the colour of the screen at a place, as a part of it across and
// down
func dotAt(s *image.RGBA, x float64, y float64) (r, g, b uint8) {
	size := s.Bounds().Size()
	c := s.RGBAAt(int(x*float64(size.X)), int(y*float64(size.Y)))
	return c.R, c.G, c.B
}

func whiteAt(s *image.RGBA, x float64, y float64) bool {
	r, g, b := dotAt(s, x, y)
	return r > 200 && g > 200 && b > 200
}

func blackAt(s *image.RGBA, x float64, y float64) bool {
	r, g, b := dotAt(s, x, y)
	return r < 40 && g < 40 && b < 40
}

// matineeTitle tells if the title of the Movie Matinee is shown, in its white
// frame
func matineeTitle(s *image.RGBA) bool {
	return whiteAt(s, 0.06, 0.34) && whiteAt(s, 0.94, 0.34) &&
		whiteAt(s, 0.5, 0.14) && whiteAt(s, 0.5, 0.56)
}

// fantavisionTitle tells if the title of the program says, in green at the
// bottom, which drawing tool it found
func fantavisionTitle(s *image.RGBA) bool {
	size := s.Bounds().Size()
	green := 0
	for y := size.Y * 88 / 100; y < size.Y; y++ {
		for x := 0; x < size.X; x += 2 {
			c := s.RGBAAt(x, y)
			if c.G > 180 && c.R < 120 && c.B < 160 {
				green++
			}
		}
	}
	return green > 100 && blackAt(s, 0.5, 0.7)
}

// fantavisionEditor tells if the screen to draw on is shown, white under the
// line of the menus
func fantavisionEditor(s *image.RGBA) bool {
	return whiteAt(s, 0.5, 18.0/384) && blackAt(s, 0.5, 22.0/384) &&
		whiteAt(s, 0.5, 0.3) && whiteAt(s, 0.9, 0.7)
}
