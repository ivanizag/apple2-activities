/*
Sheet puts pictures side by side in one PNG, to look at many at once: the
screenshots of a page, or the frames of a GIF recording as a browser shows
them, each drawn over the ones before. It is how the pictures of the
activities are checked before they are committed.

	go run ./tools/sheet -o sheet.png guides/images/dos33/*.png
	go run ./tools/sheet -o frames.png -step 10 guides/images/dos33/boot.gif

The pictures are scaled to the same width, -width, in rows of -columns. For a
GIF it prints how long each frame kept is shown.
*/
package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	out := flag.String("o", "sheet.png", "the PNG file to write")
	step := flag.Int("step", 1, "for a GIF, take one frame in so many, and always the last")
	width := flag.Int("width", 296, "the width of each picture on the sheet")
	columns := flag.Int("columns", 4, "pictures in a row")
	flag.Parse()

	var pictures []image.Image
	for _, path := range flag.Args() {
		loaded, err := load(path, *step)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v: %v\n", path, err)
			os.Exit(1)
		}
		pictures = append(pictures, loaded...)
	}
	if len(pictures) == 0 {
		fmt.Fprintln(os.Stderr, "no pictures given")
		os.Exit(1)
	}

	if err := save(*out, sheet(pictures, *width, *columns)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// load is the picture of a PNG, or the frames of a GIF, each one in so many
// and the last, as they are seen
func load(path string, step int) ([]image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if strings.ToLower(filepath.Ext(path)) != ".gif" {
		picture, _, err := image.Decode(f)
		if err != nil {
			return nil, err
		}
		return []image.Image{picture}, nil
	}

	animation, err := gif.DecodeAll(f)
	if err != nil {
		return nil, err
	}
	bounds := animation.Image[0].Bounds()
	canvas := image.NewRGBA(bounds)
	var frames []image.Image
	for i, frame := range animation.Image {
		// Each frame is drawn over the ones before, its transparent dots
		// leaving them as they are
		draw.Draw(canvas, frame.Bounds(), frame, frame.Bounds().Min, draw.Over)
		last := i == len(animation.Image)-1
		if i%step == 0 || last {
			kept := image.NewRGBA(bounds)
			draw.Draw(kept, bounds, canvas, bounds.Min, draw.Src)
			frames = append(frames, kept)
			fmt.Printf("%v frame %v: %v hundredths\n", filepath.Base(path), i, animation.Delay[i])
		}
	}
	fmt.Printf("%v: %v frames\n", filepath.Base(path), len(animation.Image))
	return frames, nil
}

// sheet lays pictures out in rows, each scaled to a width, on white
func sheet(pictures []image.Image, width int, columns int) *image.RGBA {
	heights := make([]int, len(pictures))
	rowHeight := 0
	for i, picture := range pictures {
		b := picture.Bounds()
		heights[i] = b.Dy() * width / b.Dx()
		rowHeight = max(rowHeight, heights[i])
	}
	columns = min(columns, len(pictures))
	rows := (len(pictures) + columns - 1) / columns
	const gap = 4
	out := image.NewRGBA(image.Rect(0, 0, columns*(width+gap), rows*(rowHeight+gap)))
	draw.Draw(out, out.Bounds(), image.NewUniform(color.White), image.Point{}, draw.Src)

	for i, picture := range pictures {
		x, y := (i%columns)*(width+gap), (i/columns)*(rowHeight+gap)
		scale(out, image.Rect(x, y, x+width, y+heights[i]), picture)
	}
	return out
}

// scale draws a picture into a rectangle, taking the nearest dot
func scale(dst *image.RGBA, r image.Rectangle, src image.Image) {
	b := src.Bounds()
	for y := r.Min.Y; y < r.Max.Y; y++ {
		sy := b.Min.Y + (y-r.Min.Y)*b.Dy()/r.Dy()
		for x := r.Min.X; x < r.Max.X; x++ {
			sx := b.Min.X + (x-r.Min.X)*b.Dx()/r.Dx()
			dst.Set(x, y, src.At(sx, sy))
		}
	}
}

func save(path string, picture image.Image) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	if err := png.Encode(f, picture); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}
