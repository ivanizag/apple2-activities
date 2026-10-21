package activities

import (
	"fmt"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// romxMachine is the machine of the guide, as its command line of izapple2:
// an enhanced Apple //e with a RomXce
const romxMachine = `izapple2 -model none -board 2e -cpu 65c02 -screen green \
    -rom "<internal>/Apple2e_Enhanced.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -romx \
    -s0 language`

// romxListing is the program of the page, the sixteen fonts one after the
// other
const romxListing = "../guides/listings/romx-fonts.bas"

// romxFonts are the names of the sixteen fonts of the first bank, in the
// guide of the RomXce, the labels of their pictures
var romxFonts = []string{
	"Apple US Enhanced",
	"Apple US Un-Enhanced",
	"Clinton Turner V1",
	"ReActiveMicro",
	"Dan Paymar",
	"Blippo Black",
	"Byte",
	"Colossal",
	"Count",
	"Flow",
	"Gothic",
	"Outline",
	"Pigfont",
	"Pinocchio",
	"Slant",
	"Stop",
}

/*
romxScreenshots is the RomXce on an enhanced Apple //e changing the font of
the screen: a program that calls it, through the addresses it listens to, for
each of the sixteen fonts of its first bank, recorded, and a picture of each.
*/
func romxScreenshots(t *testing.T) {
	pictures := newAlbum("romx", album.Green)
	o := start(t, romxMachine, nil)
	must(t, o.WaitForKeyboard(10))

	// The program typed and run
	must(t, typeListing(o, romxListing))
	must(t, o.TypeLines("RUN"))

	// Each font, a key for the next
	fonts := pictures.Record(o)
	for font := range 16 {
		must(t, o.WaitForText(fmt.Sprintf("FONT %d", font), 10))
		must(t, o.WaitForKeyboard(10))
		o.Run(10)
		fonts.Capture(150)
		label := fmt.Sprintf("Font %d, %v", font, romxFonts[font])
		must(t, pictures.ScreenshotOf(o, fmt.Sprintf("font-%d", font), label))
		must(t, o.Type(" "))
	}
	must(t, pictures.SaveRecording(fonts, "fonts", 150))
}
