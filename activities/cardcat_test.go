package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

/*
cardCatScreenshots is Card Cat, a program that finds what card is in each slot,
on an enhanced Apple //e with a card in each: the scan, its result, and the
firmware of the mouse card.

	izapple2 -model 2enh -s1 parallel -s5 thunderclock -s7 mouse cardcat
*/
func cardCatScreenshots(t *testing.T) {
	pictures := newAlbum("card-cat", album.Green)
	o := start(t, "2enh", map[string]string{
		"s1": "parallel,file=" + t.TempDir() + "/printer.out",
		"s5": "thunderclock",
		"s7": "mouse",
	}, "cardcat")

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
