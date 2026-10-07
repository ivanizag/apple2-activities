/*
Package album makes the pictures of the activities: screenshots of the screen
of the machine, as PNG files, and recordings of it as it changes, as GIF files,
all in the same frame, the glass and the case of the monitor around them.

The Apple II was watched on two kinds of screen, and an album says which: a
monochrome monitor, green, sharp, and with the 80 columns readable, or a
colour television, that shows the colours the machine makes out of the NTSC
signal.
*/
package album

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"

	"github.com/ivanizag/izapple2/screen"

	"github.com/ivanizag/apple2-activities/operator"
)

// Monitor is the screen the pictures of an album are taken on
type Monitor int

const (
	// Green is a monochrome monitor with a green phosphor
	Green Monitor = iota
	// Color is a colour television
	Color
)

// screenMode is the izapple2 rendering of the monitor
func (m Monitor) screenMode() int {
	if m == Color {
		return screen.ScreenModeColor
	}
	return screen.ScreenModeGreen
}

/*
An Album is the folder of the pictures of one activity, taken on a monitor.
The folder is made with the first picture.
*/
type Album struct {
	folder  string
	monitor Monitor
}

// New is an album in a folder, its pictures taken on a monitor
func New(folder string, monitor Monitor) *Album {
	return &Album{folder: folder, monitor: monitor}
}

// On is the same album with the pictures taken from now on a different
// monitor, for a page that shows both
func (a *Album) On(monitor Monitor) *Album {
	return &Album{folder: a.folder, monitor: monitor}
}

// Path is where a picture of the album is, by its name and extension
func (a *Album) Path(file string) string {
	return filepath.Join(a.folder, file)
}

// create makes a file of the album, and the folder when it is not there yet
func (a *Album) create(file string) (*os.File, error) {
	if err := os.MkdirAll(a.folder, 0o755); err != nil {
		return nil, err
	}
	return os.Create(a.Path(file))
}

// Screen is what the machine shows now, on the monitor of the album, with
// the lines doubled so that it is in the proportions of the screen
func (a *Album) Screen(o *operator.Operator) *image.RGBA {
	return doubleLines(screen.Snapshot(o.Apple2().GetVideoSource(), a.monitor.screenMode()))
}

// Screenshot writes the screen of a machine as it is now, in its frame
func (a *Album) Screenshot(o *operator.Operator, name string) error {
	return a.Write(a.Screen(o), name)
}

// ScreenshotOf writes the screen of a machine as it is now, in its frame with
// a label under it, which says which machine it is when a page shows two
func (a *Album) ScreenshotOf(o *operator.Operator, name string, label string) error {
	return a.WriteLabelled(a.Screen(o), name, label)
}

// Write writes a screen taken before, or put together, in its frame
func (a *Album) Write(screen image.Image, name string) error {
	return a.WriteLabelled(screen, name, "")
}

// WriteLabelled writes a screen taken before in its frame, with a label under
// it
func (a *Album) WriteLabelled(screen image.Image, name string, label string) error {
	f, err := a.create(name + ".png")
	if err != nil {
		return err
	}
	if err := png.Encode(f, FrameLabelled(screen, label)); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}

// SaveRecording writes a recording as a GIF, its last picture held for
// longer, by hold hundredths of a second
func (a *Album) SaveRecording(r *Recording, name string, hold int) error {
	f, err := a.create(name + ".gif")
	if err != nil {
		return err
	}
	if err := r.Encode(f, hold); err != nil {
		f.Close()
		return fmt.Errorf("the recording %v: %w", name, err)
	}
	return f.Close()
}

/*
doubleLines draws each line of the screen twice. izapple2 draws the 192 lines
of the screen 560 dots across, the two dots of each of its 280 pixels, which
is twice as wide as it is high; doubled, it is in the proportions of the
monitor.
*/
func doubleLines(in *image.RGBA) *image.RGBA {
	b := in.Bounds()
	if b.Dy() > b.Dx()/2 {
		// Already in proportion
		return in
	}
	out := image.NewRGBA(image.Rect(0, 0, b.Dx(), 2*b.Dy()))
	stride := b.Dx() * 4
	for y := 0; y < b.Dy(); y++ {
		line := in.Pix[y*in.Stride : y*in.Stride+stride]
		copy(out.Pix[2*y*out.Stride:], line)
		copy(out.Pix[(2*y+1)*out.Stride:], line)
	}
	return out
}
