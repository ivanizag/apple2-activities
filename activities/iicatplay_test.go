package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// iicAtPlayMachine is the machine of the guide, as its command line of
// izapple2, with The Apple at Play
const iicAtPlayMachine = `izapple2 -model none -board 2c -cpu 65c02 -screen color \
    -rom "<internal>/Apple2c_ROM255.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s1 iicserial \
    -s2 iicserial \
    -s6 "iicdisk,disk1=disks/Apple Presents The IIc - The Apple At Play 691-0016-A (Apple-1984).dsk"`

/*
iicAtPlayScreenshots is The Apple at Play, a diskette of Apple Presents the
IIc with games and programs of others: a hand of Apple 21, two days of
Lemonade Stand, the sonata of the Music Recital listened to, a question of
the Quick Quiz, and the demonstration of Space Quarks
*/
func iicAtPlayScreenshots(t *testing.T) {
	pictures := newAlbum("iic-at-play", album.Color)
	o := start(t, iicAtPlayMachine, nil)
	sound := album.Listen(o)

	// The menu
	must(t, o.WaitForText("The Apple at Play Menu", 30))
	must(t, o.WaitForKeyboard(10))
	o.RunSeconds(1)
	must(t, pictures.Screenshot(o, "menu"))

	// Lemonade Stand: a new game, one stand, and two days. Its text is drawn
	// in high resolution, and its questions are waited for by time.
	must(t, o.Type("4\n"))
	o.RunSeconds(20)
	must(t, pictures.Screenshot(o, "lemonsville"))
	for _, line := range []string{"Y", "1", "", ""} {
		must(t, o.TypeLines(line))
		o.RunSeconds(3)
	}
	must(t, pictures.Screenshot(o, "weather"))
	for day, decisions := range [][]string{{"30", "3", "10"}, {"40", "5", "15"}} {
		for _, decision := range decisions {
			must(t, o.TypeLines(decision))
			o.RunSeconds(2)
		}
		if day == 0 {
			must(t, pictures.Screenshot(o, "decisions"))
		}
		must(t, o.TypeLines("N"))
		o.RunSeconds(10)
		must(t, pictures.Screenshot(o, []string{"report-1", "report-2"}[day]))
		must(t, o.TypeLines(""))
		o.RunSeconds(3)
	}
	iicMenuAgain(t, o)

	// The Music Recital: its sonata, fifteen seconds of it
	must(t, o.Type("5\n"))
	must(t, o.WaitForText("PLAYING", 30))
	from := sound.Now()
	o.RunSeconds(15)
	must(t, pictures.SaveSound(sound.Clip(from, sound.Now()), "sonata"))
	must(t, pictures.Screenshot(o, "recital"))
	must(t, o.Key("Escape"))
	must(t, o.WaitForText("The Apple at Play Menu", 30))
	must(t, o.WaitForKeyboard(10))

	// The Quick Quiz: its first question, on the Introduction
	must(t, o.Type("7\n"))
	must(t, o.WaitForText("Test Your Knowledge", 30))
	o.RunSeconds(5)
	must(t, pictures.Screenshot(o, "quiz-menu"))
	must(t, o.TypeLines("1"))
	must(t, o.WaitForText("about disks is false", 30))
	o.RunSeconds(5)
	must(t, pictures.Screenshot(o, "question"))
	must(t, o.TypeLines("A"))
	o.RunSeconds(5)
	must(t, pictures.Screenshot(o, "answer"))
	iicMenuAgain(t, o)

	// Apple 21: one player, the $5 table, $10 of chips and $5 on the hand.
	// It keeps the program where the text of the screen would be, and its
	// questions are waited for by time.
	must(t, o.Type("2\n"))
	o.RunSeconds(15)
	must(t, pictures.Screenshot(o, "apple-21"))
	for _, keys := range []string{"1", "5", "\n", "5", "\n", "10\n"} {
		must(t, o.Type(keys))
		o.RunSeconds(5)
	}
	o.RunSeconds(3)
	hand := pictures.Record(o)
	hand.Capture(100)
	must(t, hand.TypeLines("5"))
	hand.Run(8*60, 6)
	must(t, hand.TypeLines("3"))
	hand.Run(8*60, 6)
	must(t, pictures.SaveRecording(hand, "hand", 300))
	iicMenuAgain(t, o)

	// Space Quarks, of Broderbund: its title, and the game playing itself
	must(t, o.Type("6\n"))
	must(t, o.WaitForText("CAPS LOCK", 30))
	must(t, o.WaitForKeyboard(10))
	must(t, o.TypeLines(""))
	o.RunSeconds(25)
	must(t, pictures.Screenshot(o, "space-quarks"))
	demo := pictures.Record(o)
	demo.Capture(50)
	demo.Run(20*60, 4)
	must(t, pictures.SaveRecording(demo, "demo", 300))
}

/*
iicMenuAgain leaves a program of The Apple at Play with Open
Apple-Control-Reset, which starts the diskette again, as its introduction
says to do with the programs that have no way out
*/
func iicMenuAgain(t *testing.T, o *operator.Operator) {
	t.Helper()
	o.HoldButton(0)
	o.Reset()
	o.Run(30)
	o.ReleaseButton(0)
	must(t, o.WaitForText("The Apple at Play Menu", 30))
	must(t, o.WaitForKeyboard(10))
}
