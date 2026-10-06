package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// cardCatMachine is the machine of the guide, as its command line of izapple2
const cardCatMachine = `izapple2 -model none -board 2e -cpu 65c02 -screen green \
    -rom "<internal>/Apple2e_Enhanced.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -ramworks 8192 -nsc main \
    -s0 language \
    -s1 parallel,file=printer.out \
    -s2 vidhd \
    -s3 fastchip \
    -s4 mockingboard \
    -s5 thunderclock \
    -s6 'diskii,disk1=<internal>/Card Cat 1.7.dsk' \
    -s7 mouse`

/*
cardCatScreenshots is Card Cat, a program that finds what card is in each slot,
on an enhanced Apple //e with a card in each: the scan, its result, and the
firmware of the mouse card.
*/
func cardCatScreenshots(t *testing.T) {
	pictures := newAlbum("card-cat", album.Green)
	o := start(t, cardCatMachine, map[string]string{
		"printer.out": t.TempDir() + "/printer.out",
	})

	// The slots checked one by one
	must(t, o.WaitForShownText("Checking Card Slot", 60))
	o.Run(30)
	must(t, pictures.Screenshot(o, "scanning"))

	// What is in each
	must(t, o.WaitForShownText("[Q]uit", 120))
	o.Run(30)
	must(t, pictures.Screenshot(o, "slots"))

	// The firmware of the mouse card
	must(t, o.Type("V"))
	must(t, o.WaitForShownText("View Card ROM", 10))
	must(t, o.Type("7"))
	must(t, o.WaitForShownText("C750:", 10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "mouse-rom"))
}
