package activities

import (
	"bytes"
	"image/gif"
	"image/png"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// iicPrintingMachine is the machine of the guide, as its command line of
// izapple2, with a printer on the printer port that is a file
const iicPrintingMachine = `izapple2 -model none -board 2c -cpu 65c02 -screen green \
    -rom "<internal>/Apple2c_ROM255.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s1 iicserial,device=file,file=printer.out \
    -s2 iicserial \
    -s6 diskii`

const (
	// iicPrintingDisk is MousePaint, whose picture is printed
	iicPrintingDisk = "077_MousePaint_Drawing_Program_68C-3005-A_1984.dsk"

	// squaresListing is the program of the page, printed
	squaresListing = "../guides/listings/squares.bas"
)

/*
iicPrintingScreenshots is the Apple IIc with an ImageWriter on its printer
port: a program typed in Applesoft and its listing printed with PR#1, which
the screen doesn't show, then the picture of the mouse of MousePaint
printed from MousePaint, and drawn by imagewriter2 as the ImageWriter would
*/
func iicPrintingScreenshots(t *testing.T) {
	pictures := newAlbum("iic-printing", album.Green)
	printer := filepath.Join(t.TempDir(), "printer.out")
	o := start(t, iicPrintingMachine, map[string]string{"printer.out": printer})

	// No disk, and Control-Reset for Applesoft
	must(t, o.WaitForText("Apple //c", 5))
	o.RunSeconds(1)
	o.Reset()
	must(t, o.WaitForKeyboard(5))

	// A program, its listing and what it prints, to the printer
	must(t, typeListing(o, squaresListing))
	must(t, o.TypeLines(
		"PR#1",
		"LIST",
		"RUN",
		"PR#0",
	))
	must(t, o.WaitForKeyboard(5))
	o.Run(30)
	must(t, pictures.Screenshot(o, "listed"))
	iicSavePrintout(t, printer, "listing")

	// MousePaint, started with Open Apple-Control-Reset
	must(t, o.InsertDisk(0, disk(t, iicPrintingDisk)))
	o.HoldButton(0)
	o.Reset()
	o.Run(30)
	o.ReleaseButton(0)
	must(t, o.WaitForText("directly to MousePaint", 30))
	o.RunSeconds(1)
	mp := &mousePaint{t: t, o: o, pictures: pictures}
	mp.click()
	o.RunSeconds(12)
	o.MouseTo(0, 0)
	o.RunSeconds(1)

	// MOUSE.PIC, and File, Print Picture
	mp.menu(mpFile, 20)
	o.RunSeconds(2)
	must(t, o.Type("MOUSE.PIC\n"))
	o.RunSeconds(5)
	before := len(raw(t, printer))
	mp.menu(mpFile, 79)
	o.RunSeconds(1)
	must(t, pictures.Screenshot(o, "printer-slot"))
	must(t, o.Type("1\n"))
	// Printed when the printer has got something, and nothing more for
	// five seconds
	for last, waited := before, 0; ; waited += 5 {
		if waited > 600 {
			t.Fatal("MousePaint did not finish printing")
		}
		o.RunSeconds(5)
		size := len(raw(t, printer))
		if size > before && size == last {
			break
		}
		last = size
	}
	iicPaper(t, raw(t, printer)[before:], pictures.Path("paper.png"))
}

// raw is all that the printer got
func raw(t *testing.T, printer string) []byte {
	t.Helper()
	data, err := os.ReadFile(printer)
	if err != nil {
		return nil
	}
	return data
}

// iicSavePrintout keeps what was printed as a text file of the guide, as
// savePrintout does
func iicSavePrintout(t *testing.T, printer string, name string) {
	t.Helper()
	lines := bytes.Split(printable(printer), []byte("\n"))
	for i := range lines {
		lines[i] = bytes.TrimRight(lines[i], " ")
	}
	text := append(bytes.Trim(bytes.Join(lines, []byte("\n")), "\n"), '\n')
	must(t, os.WriteFile(filepath.Join(guideImages, "iic-printing", name+".txt"), text, 0o644))
}

/*
lowCommands takes the top bit off what was printed, but the dots of the
graphics: MousePaint sends the codes of the ImageWriter as the Apple keeps
its characters, ESC as $9B, and imagewriter2 knows them as $1B. The bytes
of the dots after ESC G or ESC S and a count of four digits, or ESC g and a
count of three digits of eight bytes, use their eight bits.
*/
func lowCommands(printed []byte) []byte {
	out := make([]byte, 0, len(printed))
	for i := 0; i < len(printed); i++ {
		b := printed[i] & 0x7f
		out = append(out, b)
		if b != 0x1b || i+1 >= len(printed) {
			continue
		}
		i++
		command := printed[i] & 0x7f
		out = append(out, command)
		digits, times := 0, 1
		switch command {
		case 'G', 'S':
			digits = 4
		case 'g':
			digits, times = 3, 8
		}
		count := 0
		for d := 0; d < digits && i+1 < len(printed); d++ {
			i++
			digit := printed[i] & 0x7f
			out = append(out, digit)
			count = count*10 + int(digit-'0')
		}
		for n := 0; n < count*times && i+1 < len(printed); n++ {
			i++
			out = append(out, printed[i])
		}
	}
	return out
}

// iicPaper draws what was printed with imagewriter2, on its paper
func iicPaper(t *testing.T, printed []byte, path string) {
	t.Helper()
	paper := filepath.Join(t.TempDir(), "paper.gif")
	imagewriter := exec.Command("go", "tool", "imagewriter2", "-gif", paper)
	imagewriter.Stdin = bytes.NewReader(lowCommands(printed))
	var complaint bytes.Buffer
	imagewriter.Stderr = &complaint
	if err := imagewriter.Run(); err != nil {
		t.Fatalf("imagewriter2: %v\n%v", err, complaint.String())
	}
	f, err := os.Open(paper)
	must(t, err)
	defer f.Close()
	drawn, err := gif.Decode(f)
	must(t, err)
	out, err := os.Create(path)
	must(t, err)
	defer out.Close()
	must(t, png.Encode(out, edged(drawn)))
}
