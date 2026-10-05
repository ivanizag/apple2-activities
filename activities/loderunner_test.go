package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

/*
lodeRunnerScreenshots is Lode Runner of 1983 on an Apple ][+, from its
original disk: the title, the demonstration that plays itself, and a game
started.

	izapple2 -model 2plus disks/lode-runner.woz
*/
func lodeRunnerScreenshots(t *testing.T) {
	pictures := newAlbum("lode-runner", album.Color)
	o := start(t, "2plus", nil, disk(t, "lode-runner.woz"))

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
}
