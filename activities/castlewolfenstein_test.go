package activities

import (
	"strconv"
	"strings"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// castleWolfensteinMachine is the machine of the guide, as its command line
// of izapple2: the machine of Total Replay, which has the game
const castleWolfensteinMachine = totalReplayMachine

/*
castleWolfensteinScreenshots is Castle Wolfenstein, Silas Warner's game of
1981, from Total Replay: found in the launcher, its title, the controls, the
story of the cellmate, the first room, its two guards shot and one searched,
the stairs to the next room, and the game saved with Escape.

The game runs in real time, so each key is pressed at a set time, the same on
every run, as the machine is: the times were found by playing. The sound is
recorded from the start, and the part of the first fight saved: the guards
call out, in German, and the shots.
*/
func castleWolfensteinScreenshots(t *testing.T) {
	pictures := newAlbum("castle-wolfenstein", album.Color)
	o := start(t, castleWolfensteinMachine, nil)
	sound := album.Listen(o)

	// Found in the launcher, by typing its name
	must(t, o.WaitForKeyboard(30))
	o.RunSeconds(5)
	must(t, o.Type("castle wolf"))
	o.RunSeconds(2)
	must(t, pictures.Screenshot(o, "search"))

	// The title, and the choice of controls
	must(t, o.Key("Return"))
	o.RunSeconds(30)
	must(t, pictures.Screenshot(o, "title"))
	must(t, o.Key("Return"))
	o.RunSeconds(5)
	must(t, pictures.Screenshot(o, "controls"))

	// The keyboard, and the story of the cellmate
	story := pictures.Record(o)
	story.Capture(10)
	must(t, o.Type("K"))
	castleWolfensteinPlay(t, o, story, "-:10")
	must(t, pictures.SaveRecording(story, "story", 300))

	// The first room: the guard beside, shot to the right, and the other,
	// coming, to the left
	fight := pictures.Record(o)
	fight.Capture(50)
	from := sound.Now()
	castleWolfensteinPlay(t, o, fight, ";:0.3 L:1.5 K:0.3 L:3")
	must(t, pictures.SaveRecording(fight, "fight", 300))
	must(t, pictures.SaveSound(sound.Clip(from, sound.Now()), "fight"))

	// The guard on the right searched, and the supplies
	castleWolfensteinPlay(t, o, nil, ";:0.5 _:3")
	must(t, pictures.Screenshot(o, "keys"))
	castleWolfensteinPlay(t, o, nil, "Return:2")
	must(t, pictures.Screenshot(o, "supplies"))

	// Up the stairs, to the next room and its guards
	stairs := pictures.Record(o)
	stairs.Capture(50)
	castleWolfensteinPlay(t, o, stairs, "A:1.2 S:0.3 W:3 -:1.5")
	must(t, pictures.SaveRecording(stairs, "stairs", 300))

	// Saved with Escape, and back to Total Replay
	castleWolfensteinPlay(t, o, nil, "Escape:0.5")
	must(t, pictures.Screenshot(o, "saving"))
	o.RunSeconds(5.5)
	sound.Stop()
	must(t, pictures.Screenshot(o, "saved"))
}

/*
castleWolfensteinPlay presses keys, each followed by a time in seconds, as
"W:0.3": a letter, "_" for the space bar, a key by its name, or "-" for none;
recording the screen every six frames if there is a recording
*/
func castleWolfensteinPlay(t *testing.T, o *operator.Operator, r *album.Recording, steps string) {
	t.Helper()
	for _, step := range strings.Fields(steps) {
		key, seconds, _ := strings.Cut(step, ":")
		s, err := strconv.ParseFloat(seconds, 64)
		must(t, err)
		switch {
		case key == "-":
		case key == "_":
			must(t, o.Type(" "))
		case len(key) == 1:
			must(t, o.Type(key))
		default:
			must(t, o.Key(key))
		}
		if r != nil {
			r.Run(int(s*60), 6)
		} else {
			o.Run(int(s * 60))
		}
	}
}
