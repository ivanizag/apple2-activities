package activities

import (
	"image"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

/*
ultratermScreenshots is the demonstration of the Videx Ultraterm, the card of
up to 160 columns, on an Apple ][+, from the disk of utilities of the card.

	izapple2 -model ultraterm

The card blinks its cursor by the clock of the host, and the pages of the
demonstration do not come at the same frame on every run: the first is waited
for by what is on the screen, and the others taken at the times they stay
after it.
*/
func ultratermScreenshots(t *testing.T) {
	pictures := newAlbum("ultraterm", album.Green)
	o := start(t, "ultraterm", nil)

	// "Videx presents", the first page, is drawn 1440 dots across
	for {
		screen := pictures.Screen(o)
		if screen.Bounds().Dx() == 1440 && litDots(screen.Pix) > 50000 {
			break
		}
		if o.Frames() > 180*60 {
			t.Fatal("the demonstration did not start")
		}
		o.Run(30)
	}
	o.RunSeconds(2)
	must(t, pictures.Screenshot(o, "presents"))

	// The pages that follow, each kept once it stops changing: the
	// introduction, the modes, the characters and the firmware
	var pages []*image.RGBA
	last, stable := -1, 0
	for len(pages) < 4 {
		if o.Frames() > 600*60 {
			t.Fatalf("only %v pages of the demonstration came", len(pages))
		}
		o.Run(60)
		screen := pictures.Screen(o)
		lit := litDots(screen.Pix)
		if screen.Bounds().Dx() == 1440 || lit < 1000 {
			continue
		}
		if lit != last {
			last, stable = lit, 0
			continue
		}
		stable++
		if stable == 2 {
			pages = append(pages, screen)
		}
	}
	must(t, pictures.Write(pages[1], "modes"))
	must(t, pictures.Write(pages[3], "firmware"))
}

// litDots counts the dots of a screen that are lit, by their green
func litDots(pix []uint8) int {
	lit := 0
	for i := 1; i < len(pix); i += 4 {
		if pix[i] > 100 {
			lit++
		}
	}
	return lit
}
