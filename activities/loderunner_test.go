package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// lodeRunnerMachine is the machine of the guide, as its command line of izapple2
const lodeRunnerMachine = `izapple2 -model _base -board 2plus -cpu 6502 \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s6 'diskii,disk1=disks/Lode Runner.woz'`

/*
lodeRunnerScreenshots is Lode Runner of 1983 on an Apple ][+, from its
original disk: the title, the demonstration that plays itself, and a game
started and played from the keyboard, two pieces of gold taken before a
guard catches the runner.
*/
func lodeRunnerScreenshots(t *testing.T) {
	pictures := newAlbum("lode-runner", album.Color)
	o := start(t, lodeRunnerMachine, nil)

	// The title
	o.Run(16 * 60)
	must(t, pictures.Screenshot(o, "title"))

	// The demonstration, recorded from the level appearing
	o.Run(4 * 60)
	demo := pictures.Record(o)
	demo.Capture(6)
	demo.Run(20*60, 6)
	must(t, pictures.SaveRecording(demo, "demo", 100))

	// A game started with Space, the first level drawn
	must(t, o.Key("Space"))
	o.Run(10 * 60)
	must(t, pictures.Screenshot(o, "game"))

	// Played from the keyboard: Control-K, then a key sets the runner going
	// until the next one. Right to the gold, up the ladder, and left to the
	// next gold, until a guard catches him.
	must(t, o.Key("Ctrl+K"))
	o.Run(30)
	play := pictures.Record(o)
	play.Capture(10)
	must(t, o.Type("L"))
	play.Run(210, 6)
	must(t, o.Type("I"))
	play.Run(72, 6)
	must(t, o.Type("J"))
	play.Run(10*60, 6)
	must(t, pictures.SaveRecording(play, "play", 200))
}
