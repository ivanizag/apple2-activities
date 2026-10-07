package album

import (
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

/*
The frame around a screenshot: the black glass around what the tube shows,
with rounded corners. The corners outside it are transparent. A label, when
there is one, goes under the screen, in a frame taller at the bottom: it says
which machine a picture is of when a page shows two.
*/
const (
	frameWidth  = 16
	frameRadius = 14
	frameLabel  = 14
)

// labelGrey is the colour of the label in the frame
var labelGrey = color.Gray{Y: 0xb0}

// Frame puts a screen in the frame every picture of the activities has
func Frame(screen image.Image) *image.RGBA {
	return FrameLabelled(screen, "")
}

// FrameLabelled puts a screen in its frame, with a label under it
func FrameLabelled(screen image.Image, label string) *image.RGBA {
	b := screen.Bounds()
	bottom := frameWidth
	if label != "" {
		bottom += frameLabel
	}
	width, height := b.Dx()+2*frameWidth, b.Dy()+frameWidth+bottom
	out := image.NewRGBA(image.Rect(0, 0, width, height))

	// Black, but for what the rounding leaves out of each corner
	corner := func(x int, last int) int {
		return max(frameRadius-x, x-(last-frameRadius), 0)
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			dx, dy := corner(x, width-1), corner(y, height-1)
			if dx*dx+dy*dy <= frameRadius*frameRadius {
				out.Set(x, y, color.Black)
			}
		}
	}
	draw.Draw(out, image.Rect(frameWidth, frameWidth, frameWidth+b.Dx(), frameWidth+b.Dy()),
		screen, b.Min, draw.Src)

	if label != "" {
		writer := font.Drawer{
			Dst:  out,
			Src:  image.NewUniform(labelGrey),
			Face: basicfont.Face7x13,
		}
		textWidth := writer.MeasureString(label).Round()
		writer.Dot = fixed.P((width-textWidth)/2, frameWidth+b.Dy()+frameLabel-2)
		writer.DrawString(label)
	}
	return out
}
