package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// airheartMachine is the machine of the guide, as its command line of
// izapple2: an Apple IIc with Airheart in its drive and a joystick
const airheartMachine = `izapple2 -model none -board 2c -cpu 65c02 -screen color \
    -rom "<internal>/Apple2c_ROM255.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s1 iicserial \
    -s2 iicserial \
    -s6 "iicdisk,disk1=disks/airheart (san inc crack).dsk"`

/*
airheartScreenshots is Airheart, Dan Gorlin's game of 1986 in double high
resolution, for 128 KB: its title, its legend, the game playing itself, and
a game started and played a while with the joystick
*/
func airheartScreenshots(t *testing.T) {
	pictures := newAlbum("airheart", album.Color)
	o := start(t, airheartMachine, nil)
	o.Joystick(128, 128)

	// The title, Broderbund, the name, and the legend
	o.RunSeconds(8)
	must(t, pictures.Screenshot(o, "title"))
	o.RunSeconds(16)
	must(t, pictures.Screenshot(o, "broderbund"))
	o.RunSeconds(16)
	must(t, pictures.Screenshot(o, "airheart"))
	o.RunSeconds(6)
	must(t, pictures.Screenshot(o, "legend"))

	// The game playing itself
	demo := pictures.Record(o)
	demo.Capture(50)
	demo.Run(40*60, 4)
	must(t, pictures.SaveRecording(demo, "demo", 300))

	// The button of the joystick starts a game: the two guardians, and what
	// they ask
	game := pictures.Record(o)
	game.Capture(50)
	o.PressButton(0)
	game.Run(12*60, 4)
	must(t, pictures.Screenshot(o, "bring-the-sword"))
	game.Run(4*60, 4)

	// The raft, steered with the joystick, the button firing
	moves := [][2]uint8{{128, 0}, {255, 0}, {255, 128}, {128, 255}, {0, 128}, {0, 0}, {128, 0}, {128, 128}}
	for _, move := range moves {
		o.Joystick(move[0], move[1])
		for range 4 {
			o.HoldButton(0)
			game.Run(12, 4)
			o.ReleaseButton(0)
			game.Run(18, 4)
		}
	}
	o.Joystick(128, 128)
	game.Run(60, 4)
	must(t, pictures.SaveRecording(game, "game", 300))
}
