package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// appleWorksMachine is the machine of the guide, as its command line of izapple2
const appleWorksMachine = `izapple2 -model _base -board 2e -cpu 65c02 \
    -rom "<internal>/Apple2e_Enhanced.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s7 smartport,image1=disks/a2_AppleWorks_3.0_8-bit.2mg`

/*
appleWorksScreenshots is AppleWorks 3.0 on an enhanced Apple //e, its three
programs in one: a spreadsheet of the costs of a trip, printed to the
clipboard, a letter in the word processor with the table brought in, a data
base of friends, the Desktop with the three, and all saved on quitting and
found on the disk again.
*/
func appleWorksScreenshots(t *testing.T) {
	pictures := newAlbum("appleworks", album.Green)
	o := start(t, appleWorksMachine, nil)

	// Started: the date, and the main menu
	must(t, o.WaitForText("Type today's date", 120))
	must(t, o.WaitForKeyboard(10))
	must(t, o.Key("Return"))
	waitForMainMenu(t, o)
	must(t, pictures.Screenshot(o, "main-menu"))

	// A new spreadsheet, from scratch, named Trip
	chooseMenu(t, o, "1", "Make a new file for the")
	must(t, pictures.Screenshot(o, "add-files"))
	chooseMenu(t, o, "5", "From scratch")
	must(t, o.Key("Return"))
	must(t, o.WaitForText("Type a name for this new file", 10))
	must(t, o.TypeLines("Trip"))
	must(t, o.WaitForText("REVIEW/ADD/CHANGE", 30))
	must(t, o.WaitForKeyboard(10))

	// Its cells: Right Arrow enters one and goes to the next
	sheet := pictures.Record(o)
	sheet.Capture(50)
	rows := [][]string{
		{"Item", "Days", "Per day", "Cost"},
		{"Hotel", "4", "85", "+B2*C2"},
		{"Car", "4", "40", "+B3*C3"},
		{"Food", "4", "60", "+B4*C4"},
		{"Museum", "1", "25", "+B5*C5"},
		{"Total", "", "", "@SUM(D2.D5)"},
	}
	for _, row := range rows {
		for i, cell := range row {
			must(t, sheet.Type(cell))
			if i < len(row)-1 {
				must(t, o.Key("Right"))
				sheet.Capture(6)
			}
		}
		must(t, o.Key("Return"))
		for range len(row) - 1 {
			must(t, o.Key("Left"))
		}
		must(t, o.Key("Down"))
		sheet.Capture(6)
	}
	must(t, o.Key("Up"))
	for range 3 {
		must(t, o.Key("Right"))
	}
	sheet.Run(60, 6)
	must(t, pictures.SaveRecording(sheet, "spreadsheet", 300))
	must(t, o.WaitForText("Total                            765", 10))

	// Printed to the clipboard, for the word processor
	withOpenApple(t, o, "P")
	must(t, o.WaitForText("Print?", 10))
	must(t, o.Key("Return"))
	must(t, o.WaitForText("The clipboard (for the Word Processor)", 10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "print-to"))
	must(t, o.TypeLines("2"))
	must(t, o.WaitForText("Type report date", 10))
	must(t, o.Key("Return"))
	must(t, o.WaitForText("The report is now on the clipboard", 30))
	must(t, o.WaitForKeyboard(10))
	must(t, o.Key("Space"))
	must(t, o.WaitForKeyboard(10))

	// A letter, the table brought in from the clipboard, and its end
	toMainMenu(t, o)
	chooseMenu(t, o, "1", "Make a new file for the")
	chooseMenu(t, o, "3", "From scratch")
	must(t, o.Key("Return"))
	must(t, o.WaitForText("Type a name for this new file", 10))
	must(t, o.TypeLines("To.Ada"))
	must(t, o.WaitForText("File: To.Ada", 30))
	must(t, o.WaitForKeyboard(10))
	must(t, o.Type("Dear Ada,\n\nHere is what the trip to the Computer Faire will cost, worked out in the spreadsheet:\n\n"))
	withOpenApple(t, o, "C")
	must(t, o.WaitForText("From clipboard", 10))
	must(t, o.Type("F"))
	must(t, o.WaitForKeyboard(10))
	withOpenApple(t, o, "9")
	must(t, o.WaitForKeyboard(10))
	must(t, o.Type("\nSee you there.\n\nGrace\n"))
	must(t, o.WaitForKeyboard(10))
	withOpenApple(t, o, "1")
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "letter"))

	// A data base of friends: its categories, and four records
	toMainMenu(t, o)
	chooseMenu(t, o, "1", "Make a new file for the")
	chooseMenu(t, o, "4", "From scratch")
	must(t, o.Key("Return"))
	must(t, o.WaitForText("Type a name for this new file", 10))
	must(t, o.TypeLines("Friends"))
	must(t, o.WaitForText("Category names", 30))
	must(t, o.WaitForKeyboard(10))
	must(t, o.Key("Ctrl+Y"))
	must(t, o.TypeLines("Name", "City", "Phone"))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "categories"))
	must(t, o.Key("Escape"))
	must(t, o.WaitForText("Insert New Records feature", 10))
	must(t, o.WaitForKeyboard(10))
	must(t, o.Key("Space"))
	must(t, o.WaitForText("INSERT NEW RECORDS", 10))
	must(t, o.TypeLines(
		"Ada Lovelace", "London", "555-1815",
		"Grace Hopper", "Arlington", "555-1906",
		"Steve Wozniak", "Cupertino", "555-1950",
		"Margaret Hamilton", "Boston", "555-1936",
	))
	must(t, o.WaitForKeyboard(10))
	must(t, o.Key("Escape"))
	must(t, o.WaitForText("Selection: All records", 10))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "friends"))

	// The Desktop, the three files in memory at once
	withOpenApple(t, o, "Q")
	must(t, o.WaitForText("Desktop Index", 10))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "desktop"))
	must(t, o.Key("Escape"))
	o.Run(60)

	// Quit, each file saved on the disk
	toMainMenu(t, o)
	chooseMenu(t, o, "6", "Do you really want to do this?")
	must(t, o.Type("Y"))
	for _, file := range []string{`"TRIP"`, `"TO.ADA"`, `"FRIENDS"`} {
		if err := o.WaitForText(file, 30); err != nil {
			t.Fatalf("%v\n%v", err, o.Text())
		}
		must(t, o.WaitForText("Save the file on the current disk", 10))
		must(t, o.WaitForKeyboard(10))
		if file == `"TRIP"` {
			o.Run(30)
			must(t, pictures.Screenshot(o, "save"))
		}
		must(t, o.Key("Return"))
		o.Run(60)
	}

	// Back in the program selector, AppleWorks started again, and the files
	// on the disk listed
	must(t, o.WaitForText("RETURN: SELECT FILE", 60))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "selector"))
	must(t, o.Key("Down"))
	must(t, o.Key("Return"))
	must(t, o.WaitForText("Type today's date", 120))
	must(t, o.WaitForKeyboard(10))
	must(t, o.Key("Return"))
	waitForMainMenu(t, o)
	chooseMenu(t, o, "1", "Make a new file for the")
	chooseMenu(t, o, "1", "Friends")
	o.Run(30)
	must(t, pictures.Screenshot(o, "on-disk"))
}

// waitForMainMenu waits for the main menu of AppleWorks, and its keyboard
func waitForMainMenu(t *testing.T, o *operator.Operator) {
	t.Helper()
	must(t, o.WaitForText("Add files to the Desktop", 60))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
}

/*
toMainMenu goes to the main menu with Escape: Escape goes from a file to
the menu and back, so it is pressed once and then again only if the menu has
not come
*/
func toMainMenu(t *testing.T, o *operator.Operator) {
	t.Helper()
	for range 3 {
		must(t, o.Key("Escape"))
		if o.WaitUntil(3, func() bool { return o.HasText("Add files to the Desktop") }) {
			must(t, o.WaitForKeyboard(10))
			o.Run(30)
			return
		}
	}
	t.Fatalf("the main menu did not come:\n%v", o.Text())
}

// chooseMenu types the number of an item of a menu of AppleWorks, and Return,
// and waits for a text of what it opens
func chooseMenu(t *testing.T, o *operator.Operator, number string, then string) {
	t.Helper()
	must(t, o.TypeLines(number))
	must(t, o.WaitForText(then, 30))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
}

// withOpenApple types a key with the Open Apple key, button 0, held down
func withOpenApple(t *testing.T, o *operator.Operator, key string) {
	t.Helper()
	o.HoldButton(0)
	o.Run(5)
	must(t, o.Type(key))
	o.ReleaseButton(0)
}
