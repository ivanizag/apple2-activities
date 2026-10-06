package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// visiCalcMachine is the machine of the guide, as its command line of izapple2
const visiCalcMachine = `izapple2 -model _base -board 2plus -cpu 6502 \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s6 'diskii,sectors13=true,disk1=disks/VisiCalc v1.37.woz'`

/*
visiCalcScreenshots is VisiCalc 1.37, the first spreadsheet, on an Apple ][+
with a Disk II of 13 sectors: a household budget typed in, totalled with
@SUM, a column of the year made by replicating one formula, the share of each
item, and the rent changed to see everything worked out again.
*/
func visiCalcScreenshots(t *testing.T) {
	pictures := newAlbum("visicalc", album.Green)
	o := start(t, visiCalcMachine, nil)

	// Started, the empty sheet
	must(t, o.WaitForText("SOFTWARE ARTS", 60))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "started"))

	// The items and their cost in a month, each cell gone to with >
	entering := pictures.Record(o)
	entering.Capture(50)
	must(t, entering.Type(">A1\nITEM\n>B1\n\"MONTH\n"))
	for _, item := range []struct{ row, name, cost string }{
		{"2", "RENT", "450"},
		{"3", "FOOD", "210"},
		{"4", "FUEL", "65"},
		{"5", "PHONE", "30"},
	} {
		must(t, entering.Type(">A"+item.row+"\n"+item.name+"\n>B"+item.row+"\n"+item.cost+"\n"))
	}
	entering.Run(60, 6)
	must(t, pictures.SaveRecording(entering, "entering", 300))

	// The total
	must(t, o.Type(">A6\nTOTAL\n>B6\n@SUM(B2.B5)\n"))
	must(t, o.WaitForText("TOTAL          755", 10))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "total"))

	// The year: a formula, and the same down the column, B2 relative
	must(t, o.Type(">C1\n\"YEAR\n>C2\n+B2*12\n"))
	must(t, o.Type("/R\nC3.C6\n"))
	must(t, o.WaitForText("R=RELATIVE", 10))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "replicate"))
	must(t, o.Type("R"))
	must(t, o.WaitForText("9060", 10))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "year"))

	// The share of each item: B2 relative, the total B6 not
	must(t, o.Type(">D1\n\"SHARE %\n>D2\n+B2/B6*100\n"))
	must(t, o.Type("/R\nD3.D6\nRN"))
	must(t, o.WaitForText("100", 10))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "share"))

	// Every number shown as a whole number
	must(t, o.Type("/GFI"))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "integers"))

	// What if the rent goes up: everything that depends on it changes
	whatIf := pictures.Record(o)
	whatIf.Capture(50)
	must(t, whatIf.Type(">B2\n500\n"))
	whatIf.Run(120, 6)
	must(t, pictures.SaveRecording(whatIf, "what-if", 300))
	if !o.HasText("TOTAL          805     9660      100") {
		t.Fatalf("the sheet was not worked out again:\n%v", o.Text())
	}
}
