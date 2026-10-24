package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// princeOfPersiaMachine is the machine of the guide, as its command line of
// izapple2
const princeOfPersiaMachine = `izapple2 -model none -board 2e -cpu 65c02 -screen color \
    -rom "<internal>/Apple2e_Enhanced.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s6 'diskii,disk1=disks/prince of persia side a (san inc crack).dsk'`

/*
princeOfPersiaScreenshots is Prince of Persia, Jordan Mechner's game of 1989,
on an enhanced Apple //e: the titles and the story, the princess and Jaffar
with his hourglass, and the game playing itself, the run, the jumps and a
fight with the sword; then a game, the prince running from his cell into the
room of a guard, without a sword, and killed.

The attract mode runs by itself, the same on every run, so its parts are
taken at the times they come, in seconds from the start. The game reads the
keyboard for the key down: izapple2 has no key up, and a key pressed again
and again, a third of a second apart, keeps the prince running as a key held.
*/
func princeOfPersiaScreenshots(t *testing.T) {
	pictures := newAlbum("prince-of-persia", album.Color)
	o := start(t, princeOfPersiaMachine, nil)

	// The titles and the story
	o.RunSeconds(47)
	must(t, pictures.Screenshot(o, "title"))
	o.RunSeconds(13)
	must(t, pictures.Screenshot(o, "story"))

	// The princess, and Jaffar
	o.RunSeconds(18)
	princess := pictures.Record(o)
	princess.Capture(50)
	princess.Run(44*60, 6)
	must(t, pictures.SaveRecording(princess, "princess", 300))
	o.RunSeconds(28)
	must(t, pictures.Screenshot(o, "jaffar"))

	// The game playing itself
	o.RunSeconds(32)
	demo := pictures.Record(o)
	demo.Capture(50)
	demo.Run(50*60, 6)
	must(t, pictures.SaveRecording(demo, "demo", 300))

	// A game: Return, and the cell of the prince
	o.RunSeconds(13)
	must(t, o.Key("Return"))
	o.RunSeconds(30)
	must(t, pictures.Screenshot(o, "cell"))

	// To the right, down to the floor below, and on into the room of a
	// guard, without a sword
	run := pictures.Record(o)
	run.Capture(50)
	princeOfPersiaKeys(t, o, run, "L", 8)
	run.Run(120, 6)
	princeOfPersiaKeys(t, o, run, "L", 16)
	run.Run(300, 6)
	must(t, pictures.SaveRecording(run, "run", 300))
	must(t, pictures.Screenshot(o, "killed"))
}

// princeOfPersiaKeys presses a key again and again, a third of a second
// apart, as a key held down
func princeOfPersiaKeys(t *testing.T, o *operator.Operator, r *album.Recording, key string, times int) {
	t.Helper()
	for range times {
		must(t, o.Type(key))
		r.Run(18, 6)
	}
}
