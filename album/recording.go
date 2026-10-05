package album

import (
	"errors"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"io"

	"github.com/ivanizag/apple2-activities/operator"
)

/*
A Recording is an animated screenshot, for what is worth watching move: a GIF
of the screen in its frame, a picture taken every few frames of the machine.
Each picture after the first is only the part of the screen that changed, with
the colours it has, and a picture that changed nothing makes the one before it
last longer, up to LongestStill.

Delays are in hundredths of a second, as a GIF counts them. Browsers slow down
anything shorter than two, so a picture is not worth taking more often than
every two frames of the machine.
*/
type Recording struct {
	album  *Album
	speed  int
	op     *operator.Operator
	frames []*image.Paletted
	delays []int
	last   *image.RGBA
}

// LongestStill is the longest a picture stays, in hundredths of a second: a
// wait on the machine with nothing changing is shortened to that
const LongestStill = 300

// Record starts a recording of the machine an operator is at, on the monitor
// of the album. Nothing is taken until Capture or Run.
func (a *Album) Record(op *operator.Operator) *Recording {
	return &Recording{album: a, op: op, speed: 1}
}

// Faster makes the recording play the time Run runs the machine a number of
// times faster than the machine did, for a wait that would be tedious at its
// real speed
func (r *Recording) Faster(times int) *Recording {
	r.speed = times
	return r
}

// Capture takes a picture of the screen, shown for a time in hundredths of a
// second
func (r *Recording) Capture(delay int) {
	whole := Frame(r.album.Screen(r.op))
	if r.last != nil && whole.Bounds() != r.last.Bounds() {
		// The screen changed its size, as Super Hi-Res does: the picture is
		// put on one of the size of the first
		fitted := image.NewRGBA(r.last.Bounds())
		draw.Draw(fitted, fitted.Bounds(), whole, whole.Bounds().Min, draw.Src)
		whole = fitted
	}

	changed := whole.Bounds()
	if r.last != nil {
		changed = changedArea(r.last, whole)
		if changed.Empty() {
			last := len(r.delays) - 1
			r.delays[last] = min(r.delays[last]+delay, LongestStill)
			return
		}
	}
	before := r.last
	r.last = whole
	r.frames = append(r.frames, paletted(whole, before, changed))
	r.delays = append(r.delays, delay)
}

// Run runs the machine for some frames, taking a picture every few of them,
// shown for as long as they took, or less in a faster recording
func (r *Recording) Run(frames int, every int) {
	for done := 0; done < frames; done += every {
		r.op.Run(every)
		r.Capture(max(every*100/operator.FramesPerSecond/r.speed, 2))
	}
}

// Type types a text at the pace of a hand, taking a picture after each key
func (r *Recording) Type(text string) error {
	for _, c := range text {
		if err := r.op.Type(string(c)); err != nil {
			return err
		}
		r.Capture(keyDelay)
	}
	return nil
}

// TypeLines types lines of text, each followed by Return, with pictures
func (r *Recording) TypeLines(lines ...string) error {
	for _, line := range lines {
		if err := r.Type(line + "\n"); err != nil {
			return err
		}
	}
	return nil
}

// keyDelay is how long a picture stays after a key, the time between two keys
// of the operator
const keyDelay = 100 / 6

/*
Encode writes the recording as a GIF that plays in a loop, the last picture
held for longer, by hold hundredths of a second, so that what it ends on is
seen before it starts again
*/
func (r *Recording) Encode(w io.Writer, hold int) error {
	if len(r.delays) == 0 {
		return errors.New("the recording has nothing in it")
	}
	delays := append([]int(nil), r.delays...)
	delays[len(delays)-1] += hold
	return gif.EncodeAll(w, &gif.GIF{Image: r.frames, Delay: delays})
}

// changedArea is the smallest rectangle that has all that changed between two
// pictures of the same size
func changedArea(before *image.RGBA, after *image.RGBA) image.Rectangle {
	changed := image.Rectangle{}
	b := after.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		row := y * after.Stride
		for x := b.Min.X; x < b.Max.X; x++ {
			i := row + 4*x
			if before.Pix[i] != after.Pix[i] || before.Pix[i+1] != after.Pix[i+1] ||
				before.Pix[i+2] != after.Pix[i+2] || before.Pix[i+3] != after.Pix[i+3] {
				changed = changed.Union(image.Rect(x, y, x+1, y+1))
			}
		}
	}
	return changed
}

/*
paletted is a part of a picture with a palette of its own, made of its
colours. The screen of an Apple II has few, and only when there are more than
a GIF can hold, as the colour television can show, are they taken to the
nearest of a palette of all. When there is a picture before, what has not
changed since is left transparent, to show the one before through it: two
sprites far apart change a large rectangle, but few of its dots.
*/
func paletted(picture *image.RGBA, before *image.RGBA, area image.Rectangle) *image.Paletted {
	index := map[color.RGBA]uint8{}
	colors := color.Palette{}
	if before != nil {
		colors = append(colors, color.Transparent)
	}
	unchanged := func(x, y int) bool {
		return before != nil && before.RGBAAt(x, y) == picture.RGBAAt(x, y)
	}
	for y := area.Min.Y; y < area.Max.Y && colors != nil; y++ {
		for x := area.Min.X; x < area.Max.X; x++ {
			c := picture.RGBAAt(x, y)
			if _, ok := index[c]; ok || unchanged(x, y) {
				continue
			}
			if len(colors) == 256 {
				colors = nil
				break
			}
			index[c] = uint8(len(colors))
			colors = append(colors, c)
		}
	}

	if colors == nil {
		all := append(color.Palette{color.Transparent}, palette.Plan9[:255]...)
		out := image.NewPaletted(area, all)
		draw.Draw(out, area, picture, area.Min, draw.Src)
		return out
	}
	out := image.NewPaletted(area, colors)
	for y := area.Min.Y; y < area.Max.Y; y++ {
		for x := area.Min.X; x < area.Max.X; x++ {
			if unchanged(x, y) {
				continue // Index 0, transparent
			}
			out.SetColorIndex(x, y, index[picture.RGBAAt(x, y)])
		}
	}
	return out
}

// glideStep is how far the pointer goes between two pictures of a glide, as
// a part of the screen
const glideStep = 0.03

/*
Glide takes the mouse to a place, from 0 to 1 across and down, a little at a
time, as a hand would, with a picture at each step: in a recording the pointer
is seen going there, where the operator's MouseTo would take it there between
two pictures.
*/
func (r *Recording) Glide(toX float64, toY float64) {
	fromX, fromY := r.op.MousePosition()
	distance := max(abs(toX-fromX), abs(toY-fromY))
	steps := int(distance/glideStep) + 1
	for step := 1; step <= steps; step++ {
		f := float64(step) / float64(steps)
		r.op.MouseTo(fromX+(toX-fromX)*f, fromY+(toY-fromY)*f)
		r.Capture(7)
	}
}

func abs(f float64) float64 {
	if f < 0 {
		return -f
	}
	return f
}

// Press presses the button of the mouse, or lets go of it, and takes a
// picture after
func (r *Recording) Press(down bool) {
	if down {
		r.op.Hold()
	} else {
		r.op.Release()
	}
	r.Run(4, 2)
}

// Click presses and releases the button of the mouse, with pictures
func (r *Recording) Click() {
	r.Press(true)
	r.Run(6, 2)
	r.Press(false)
	r.Run(10, 2)
}

// DoubleClick clicks twice in a row, close enough for a double click, with
// pictures
func (r *Recording) DoubleClick() {
	r.Press(true)
	r.Run(4, 2)
	r.Press(false)
	r.Run(2, 2)
	r.Press(true)
	r.Run(4, 2)
	r.Press(false)
}
