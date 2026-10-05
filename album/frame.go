package album

import (
	"image"
	"image/color"
	"image/draw"
)

/*
The frame around a screenshot: the black glass around what the tube shows,
with rounded corners. The corners outside it are transparent.
*/
const (
	frameWidth  = 16
	frameRadius = 14
)

// Frame puts a screen in the frame every picture of the activities has
func Frame(screen image.Image) *image.RGBA {
	b := screen.Bounds()
	width, height := b.Dx()+2*frameWidth, b.Dy()+2*frameWidth
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
	return out
}
