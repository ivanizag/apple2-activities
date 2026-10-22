package activities

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/png"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// printShopMachine is the machine of the guide, as its command line of
// izapple2
const printShopMachine = `izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s1 parallel,file=printer.out \
    -s6 'diskii,disk1=disks/The Print Shop (4am crack) side A.dsk'`

/*
printShopScreenshots is The Print Shop of 1984 on an Apple ][+ with a
printer: set up for an ImageWriter on the parallel card, then a greeting card
of its own design, a sign, a banner, and a kaleidoscope of Screen Magic
printed. What the card prints, in the codes of the ImageWriter, is drawn on
paper by imagewriter2.

The Print Shop is driven by the arrow keys and Return, and remembers the
choice of each menu: each key is pressed once the screen has settled.
*/
func printShopScreenshots(t *testing.T) {
	pictures := newAlbum("print-shop", album.Color)
	printer := filepath.Join(t.TempDir(), "printer.out")
	o := start(t, printShopMachine, map[string]string{"printer.out": printer})
	ps := &printShop{t: t, o: o, pictures: pictures, printer: printer}

	// The title, and the main menu, its picture changing with the choice
	waitForScreen(t, o, pictures, printShopTitle)
	o.RunSeconds(1)
	must(t, pictures.Screenshot(o, "title"))
	ps.settle(nil)
	menu := pictures.Record(o)
	menu.Capture(100)
	for range 6 {
		ps.key(menu, "Right")
	}
	must(t, pictures.SaveRecording(menu, "menu", 100))

	// Setup: the ImageWriter, which it has chosen already, the Apple II
	// Parallel Interface in slot 1, one drive, and the test of the printer,
	// saved on the disk
	ps.key(nil, "Return")
	must(t, pictures.Screenshot(o, "printer"))
	ps.key(nil, "Return")
	ps.key(nil, "Right")
	must(t, pictures.Screenshot(o, "interface"))
	ps.key(nil, "Return")
	ps.key(nil, "Return")
	ps.key(nil, "Left")
	ps.key(nil, "Return")
	ps.key(nil, "Return")
	ps.waitForPrinter("")
	ps.printout("welcome")
	must(t, pictures.Screenshot(o, "tested"))
	ps.key(nil, "Return")

	// A greeting card of its own design. The front: the border of stars, the
	// gift box, large, and a message in Party, in three dimensions, at the top
	ps.key(nil, "Right")
	ps.key(nil, "Return")
	ps.key(nil, "Return")
	borders := pictures.Record(o)
	borders.Capture(100)
	for range 6 {
		ps.key(borders, "Right")
	}
	must(t, pictures.SaveRecording(borders, "borders", 200))
	ps.key(nil, "Return")
	ps.key(nil, "Return")
	graphics := pictures.Record(o)
	graphics.Capture(100)
	for range 3 {
		ps.key(graphics, "Right")
	}
	must(t, pictures.SaveRecording(graphics, "graphics", 200))
	ps.key(nil, "Return")
	ps.key(nil, "Right")
	ps.key(nil, "Right")
	ps.key(nil, "Return")
	fonts := pictures.Record(o)
	fonts.Capture(100)
	for range 4 {
		ps.key(fonts, "Right")
	}
	must(t, pictures.SaveRecording(fonts, "fonts", 200))
	ps.key(nil, "Return")
	message := pictures.Record(o)
	message.Capture(100)
	ps.typeText(message, "HAPPY")
	ps.key(message, "Ctrl+F")
	ps.key(message, "Ctrl+F")
	ps.key(message, "Return")
	ps.typeText(message, "BIRTHDAY")
	ps.key(message, "Ctrl+F")
	ps.key(message, "Ctrl+F")
	must(t, pictures.SaveRecording(message, "front", 200))
	ps.endMessage()
	ps.key(nil, "Right")
	must(t, pictures.Screenshot(o, "centre"))
	ps.key(nil, "Return")

	// The inside: a thin line, the cake, small, in a layout of its own, three
	// at the top and three at the bottom, and a message in Alexia
	ps.key(nil, "Return")
	ps.key(nil, "Return")
	ps.key(nil, "Return")
	ps.key(nil, "Return")
	ps.key(nil, "Right")
	ps.key(nil, "Right")
	ps.key(nil, "Return")
	layout := pictures.Record(o)
	layout.Capture(100)
	for place := 1; place <= 13; place++ {
		if place <= 3 || place >= 11 {
			ps.key(layout, "Return")
		}
		ps.key(layout, "Right")
	}
	must(t, pictures.SaveRecording(layout, "layout", 200))
	ps.key(nil, "Return")
	ps.key(nil, "Right")
	ps.key(nil, "Return")
	inside := pictures.Record(o)
	inside.Capture(100)
	ps.typeText(inside, "MANY HAPPY")
	ps.key(inside, "Return")
	ps.typeText(inside, "RETURNS")
	ps.key(inside, "Return")
	ps.typeText(inside, "ADA")
	must(t, pictures.SaveRecording(inside, "inside", 200))
	ps.endMessage()
	ps.key(nil, "Return")

	// Printed, with a line of credit for the back
	ps.key(nil, "Return")
	ps.typeText(nil, "MADE ON AN APPLE II")
	must(t, pictures.Screenshot(o, "credit"))
	ps.key(nil, "Return")
	ps.key(nil, "Right")
	ps.key(nil, "Right")
	must(t, pictures.Screenshot(o, "print"))
	must(t, o.Key("Return"))
	ps.waitForPrinter("thinking")
	ps.printout("card")
	ps.goodbye()

	// A sign: a thick line, the apple of the second page of graphics, large,
	// and a text above it and one below in Stencil, not in the middle
	ps.key(nil, "Right")
	ps.key(nil, "Return")
	ps.key(nil, "Right")
	ps.key(nil, "Right")
	ps.key(nil, "Return")
	ps.key(nil, "Return")
	for range 10 {
		ps.key(nil, "Right")
	}
	ps.key(nil, "Return")
	ps.key(nil, "Right")
	ps.key(nil, "Right")
	must(t, pictures.Screenshot(o, "apple"))
	ps.key(nil, "Return")
	ps.key(nil, "Right")
	ps.key(nil, "Right")
	ps.key(nil, "Return")
	for range 7 {
		ps.key(nil, "Right")
	}
	ps.key(nil, "Return")
	sign := pictures.Record(o)
	sign.Capture(100)
	ps.typeText(sign, "APPLE II CLUB")
	for range 12 {
		ps.key(sign, "Return")
	}
	ps.typeText(sign, "MEETS TONIGHT")
	must(t, pictures.SaveRecording(sign, "sign-message", 200))
	ps.endMessage()
	ps.key(nil, "Right")
	ps.key(nil, "Return")
	ps.key(nil, "Right")
	ps.key(nil, "Right")
	must(t, o.Key("Return"))
	ps.waitForPrinter("")
	ps.printout("sign")
	ps.goodbye()

	// A banner: in Party, solid, and the cake at both ends
	ps.key(nil, "Right")
	ps.key(nil, "Right")
	ps.key(nil, "Return")
	for range 4 {
		ps.key(nil, "Right")
	}
	ps.key(nil, "Return")
	ps.key(nil, "Return")
	banner := pictures.Record(o)
	banner.Capture(100)
	ps.typeText(banner, "HAPPY BIRTHDAY ADA")
	must(t, pictures.SaveRecording(banner, "banner-message", 200))
	ps.key(nil, "Return")
	ps.key(nil, "Return")
	ps.key(nil, "Return")
	ps.key(nil, "Right")
	ps.key(nil, "Right")
	must(t, pictures.Screenshot(o, "banner-ends"))
	ps.key(nil, "Return")
	ps.key(nil, "Right")
	must(t, o.Key("Return"))
	ps.waitForPrinter("")
	ps.printout("banner")
	ps.goodbye()

	// Screen Magic: the first kaleidoscope, a pattern and the next, frozen
	// with Escape, and printed on the top half of a page, its lit dots in
	// ink, in a frame
	ps.key(nil, "Right")
	ps.key(nil, "Return")
	ps.key(nil, "Return")
	ps.key(nil, "Return")
	must(t, o.Key("Space"))
	kaleidoscope := pictures.Record(o)
	kaleidoscope.Run(6*operator.FramesPerSecond, 4)
	must(t, o.Key("Return"))
	kaleidoscope.Run(6*operator.FramesPerSecond, 4)
	must(t, pictures.SaveRecording(kaleidoscope, "kaleidoscope", 0))
	must(t, o.Key("Escape"))
	ps.settle(nil)
	for range 4 {
		ps.key(nil, "Right")
	}
	ps.key(nil, "Return")
	ps.key(nil, "Return")
	ps.key(nil, "Return")
	must(t, pictures.Screenshot(o, "frame"))
	ps.key(nil, "Return")
	ps.key(nil, "Right")
	ps.key(nil, "Right")
	must(t, o.Key("Return"))
	ps.waitForPrinter("")
	ps.printout("kaleidoscope-printed")
}

// printShop is The Print Shop at work, and what its printer printed
type printShop struct {
	t        *testing.T
	o        *operator.Operator
	pictures *album.Album
	printer  string
	printed  int
}

/*
settle runs the machine until the screen stays the same for a second and the
program reads the keyboard: The Print Shop draws its menus and its previews a
part at a time, and loads them from the disk, and a key pressed while it does
is lost. The cursors that blink, where a text is typed or a place of the page
is chosen, are too small to count. A recording, if any, takes a picture as it
goes.
*/
func (ps *printShop) settle(r *album.Recording) {
	ps.t.Helper()
	var last *image.RGBA
	still := 0
	limit := ps.o.Frames() + 120*operator.FramesPerSecond
	for still < 4 {
		if ps.o.Frames() > limit {
			ps.t.Fatal("the screen of The Print Shop did not settle")
		}
		if r != nil {
			r.Run(15, 5)
		} else {
			ps.o.Run(15)
		}
		screen := ps.pictures.Screen(ps.o)
		if last != nil && changedDots(screen, last) < 1500 {
			still++
		} else {
			still = 0
		}
		last = screen
	}
	must(ps.t, ps.o.WaitForKeyboard(10))
}

// changedDots counts the dots of the screen that are not the same in two
// pictures of it
func changedDots(a *image.RGBA, b *image.RGBA) int {
	changed := 0
	for i := 0; i < len(a.Pix); i += 4 {
		if a.Pix[i] != b.Pix[i] || a.Pix[i+1] != b.Pix[i+1] || a.Pix[i+2] != b.Pix[i+2] {
			changed++
		}
	}
	return changed
}

// key presses a key, and waits for the screen to settle
func (ps *printShop) key(r *album.Recording, name string) {
	ps.t.Helper()
	if len(name) == 1 {
		must(ps.t, ps.o.Type(name))
	} else {
		must(ps.t, ps.o.Key(name))
	}
	ps.settle(r)
}

// typeText types a text at the pace of a hand, the editor of a message
// taking each key as it comes, and waits for the screen to settle
func (ps *printShop) typeText(r *album.Recording, text string) {
	ps.t.Helper()
	if r != nil {
		must(ps.t, r.Type(text))
	} else {
		must(ps.t, ps.o.Type(text))
	}
	ps.settle(r)
}

// endMessage presses Return down the lines of a message to the last, after
// which the program asks if the text goes in the middle, top to bottom
func (ps *printShop) endMessage() {
	ps.t.Helper()
	for range 20 {
		ps.key(nil, "Return")
		if centreQuestion(ps.pictures.Screen(ps.o)) {
			return
		}
	}
	ps.t.Fatal("the message did not end")
}

// centreQuestion tells if the editor of a message asks if the text goes in
// the middle, its YES in white at the bottom right
func centreQuestion(s *image.RGBA) bool {
	size := s.Bounds().Size()
	white, all := 0, 0
	for y := size.Y * 866 / 1000; y < size.Y*900/1000; y++ {
		for x := size.X * 825 / 1000; x < size.X*945/1000; x++ {
			c := s.RGBAAt(x, y)
			if c.R > 200 && c.G > 200 && c.B > 200 {
				white++
			}
			all++
		}
	}
	return white > all/2
}

// goodbye goes from the menu of printing, after a print, back to the main
// menu: down to its last choice, and Return to say so
func (ps *printShop) goodbye() {
	ps.t.Helper()
	ps.key(nil, "Right")
	ps.key(nil, "Right")
	ps.key(nil, "Return")
	ps.key(nil, "Return")
}

/*
waitForPrinter runs the machine until the printer has had something new, and
then nothing more for 45 seconds, and the screen settles. The Print Shop
works out each part of a card before it prints it, and is quiet for up to
half a minute while it does, the colours of THINKING on the screen, of which
a picture is taken when a name is given.
*/
func (ps *printShop) waitForPrinter(thinking string) {
	ps.t.Helper()
	size := int64(-1)
	still := 0
	limit := ps.o.Frames() + 600*operator.FramesPerSecond
	for still < 45 {
		if ps.o.Frames() > limit {
			ps.t.Fatal("the printer did not stop")
		}
		ps.o.RunSeconds(1)
		if thinking != "" && colourful(ps.pictures.Screen(ps.o)) {
			ps.o.RunSeconds(1)
			must(ps.t, ps.pictures.Screenshot(ps.o, thinking))
			thinking = ""
		}
		now := int64(0)
		if f, err := os.Stat(ps.printer); err == nil {
			now = f.Size()
		}
		if now > int64(ps.printed) && now == size {
			still++
		} else {
			still = 0
		}
		size = now
	}
	ps.settle(nil)
}

// colourful tells if most of the screen is in colour, neither black nor
// white
func colourful(s *image.RGBA) bool {
	coloured := 0
	for i := 0; i < len(s.Pix); i += 4 {
		r, g, b := s.Pix[i], s.Pix[i+1], s.Pix[i+2]
		if max(r, g, b)-min(r, g, b) > 80 {
			coloured++
		}
	}
	return coloured > len(s.Pix)/4/2
}

/*
printout draws on paper what was printed since the printout before, with
imagewriter2, which reads the codes of the ImageWriter and places each dot
where the print head puts it, at 144 dots an inch, and keeps it as a picture
of the guide, with a grey line at the edges of the paper and its strips of
holes. A banner, longer than it is wide, is turned to be read as it hangs.
*/
func (ps *printShop) printout(name string) {
	ps.t.Helper()
	raw, err := os.ReadFile(ps.printer)
	must(ps.t, err)
	part := raw[ps.printed:]
	ps.printed = len(raw)
	if len(part) == 0 {
		ps.t.Fatalf("nothing was printed for %v", name)
	}

	paper := filepath.Join(ps.t.TempDir(), name+".gif")
	imagewriter := exec.Command("go", "tool", "imagewriter2", "-gif", paper)
	imagewriter.Stdin = bytes.NewReader(part)
	var complaint bytes.Buffer
	imagewriter.Stderr = &complaint
	if err := imagewriter.Run(); err != nil {
		ps.t.Fatalf("imagewriter2: %v\n%v", err, complaint.String())
	}
	f, err := os.Open(paper)
	must(ps.t, err)
	defer f.Close()
	drawn, err := gif.Decode(f)
	must(ps.t, err)

	page := edged(drawn)
	if b := page.Bounds(); b.Dy() > 2*b.Dx() {
		page = turned(page)
	}
	out, err := os.Create(ps.pictures.Path(name + ".png"))
	must(ps.t, err)
	defer out.Close()
	must(ps.t, png.Encode(out, page))
}

// edged is a picture of the paper with a grey line around it, a margin of
// paper around what was printed, and the strips of the continuous paper at
// its sides, half an inch with a hole every half inch, torn off along a
// perforation; a line or two, too short for them, has none
func edged(drawn image.Image) *image.RGBA {
	const (
		margin = 24
		strip  = 72 // half an inch at 144 dots an inch
		hole   = 22 // 5/32 of an inch
	)
	b := drawn.Bounds()
	side := strip
	if b.Dy()+2*margin < 4*strip {
		side = 0
	}
	page := image.NewRGBA(image.Rect(0, 0, b.Dx()+2*side, b.Dy()+2*margin))
	draw.Draw(page, page.Bounds(), image.White, image.Point{}, draw.Src)
	draw.Draw(page, b.Add(image.Pt(side, margin)), drawn, b.Min, draw.Src)
	edge := color.RGBA{0xb0, 0xb0, 0xb0, 0xff}
	size := page.Bounds().Size()
	for x := 0; x < size.X; x++ {
		page.Set(x, 0, edge)
		page.Set(x, size.Y-1, edge)
	}
	for y := 0; y < size.Y; y++ {
		page.Set(0, y, edge)
		page.Set(size.X-1, y, edge)
	}
	if side == 0 {
		return page
	}
	shadow := color.RGBA{0xe8, 0xe8, 0xe8, 0xff}
	for _, x := range []int{side, size.X - 1 - side} {
		for y := 0; y < size.Y; y += 6 {
			for d := 0; d < 3 && y+d < size.Y; d++ {
				page.Set(x, y+d, edge)
			}
		}
	}
	const r = hole / 2
	for _, cx := range []int{side / 2, size.X - side/2} {
		for cy := strip / 2; cy+r < size.Y; cy += strip {
			for dy := -r; dy <= r; dy++ {
				for dx := -r; dx <= r; dx++ {
					d := dx*dx + dy*dy
					switch {
					case d <= (r-1)*(r-1):
						page.Set(cx+dx, cy+dy, shadow)
					case d <= r*r:
						page.Set(cx+dx, cy+dy, edge)
					}
				}
			}
		}
	}
	return page
}

// turned is a picture turned a quarter to the left, its right side up
func turned(in *image.RGBA) *image.RGBA {
	b := in.Bounds()
	out := image.NewRGBA(image.Rect(0, 0, b.Dy(), b.Dx()))
	for y := 0; y < b.Dy(); y++ {
		for x := 0; x < b.Dx(); x++ {
			out.Set(y, b.Dx()-1-x, in.At(x, y))
		}
	}
	return out
}

// printShopTitle tells if the title is shown, white letters on black under
// the crown of Broderbund
func printShopTitle(s *image.RGBA) bool {
	size := s.Bounds().Size()
	white := 0
	for y := size.Y * 25 / 100; y < size.Y*65/100; y++ {
		for x := 0; x < size.X; x += 2 {
			c := s.RGBAAt(x, y)
			if c.R > 200 && c.G > 200 && c.B > 200 {
				white++
			}
		}
	}
	return white > 2000 && blackAt(s, 0.02, 0.5)
}
