package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// holidayMusicMachine is the machine of the guide, as its command line of
// izapple2
const holidayMusicMachine = `izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s4 mockingboard \
    -s6 "diskii,disk1=disks/Sweet Micro Systems Mockingboard presents Holiday Music.dsk"`

/*
holidayMusicScreenshots is Mockingboard presents Holiday Music, the Christmas
songs of Sweet Micro Systems for the Mockingboard, on an Apple ][+: its menu
of seven songs, and the first, Song of Joy, its words on the screen and the
whole of its music recorded, before it starts again
*/
func holidayMusicScreenshots(t *testing.T) {
	pictures := newAlbum("holiday-music", album.Color)
	o := start(t, holidayMusicMachine, nil)
	sound := album.Listen(o)

	// The menu, once drawn: the screen still, and many dots lit
	(&iicTour{t: t, o: o, pictures: pictures}).settle(nil, 2, 20000)
	must(t, o.WaitForKeyboard(10))
	must(t, pictures.Screenshot(o, "menu"))

	// Song of Joy, a key alone, no Return: the song loads for five seconds,
	// plays for fifty, and starts again
	must(t, o.Type("1"))
	o.RunSeconds(5)
	from := sound.Now()
	o.RunSeconds(5)
	must(t, pictures.Screenshot(o, "song-of-joy"))
	o.RunSeconds(48)
	must(t, pictures.SaveSound(sound.Clip(from, sound.Now()), "song-of-joy"))
	sound.Stop()
}
