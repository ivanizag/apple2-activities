package activities

import (
	"math"
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
1981, from Total Replay, from the start to the end: found in the launcher, its
title, the controls, the story of the cellmate, the first room, its two guards
shot and one searched, the stairs to the second floor, and the castle crossed
floor by floor, a uniform found in a chest on the third, out of the castle at
the end of the fifth.

The game runs in real time, so each key is pressed at a set time, the same on
every run, as the machine is. The times of the first room were found by
playing; those of the floors after it by a program that played them, trying
many ways through each room and keeping one that got out alive. The sound is
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

	// Up the stairs, to the second floor and its guards
	stairs := pictures.Record(o)
	stairs.Capture(50)
	castleWolfensteinPlay(t, o, stairs, "A:1.2 S:0.3 W:3 -:1")
	must(t, pictures.SaveRecording(stairs, "stairs", 300))
	sound.Stop()

	// The second floor, to the stairs up
	castleWolfensteinPlay(t, o, nil, castleWolfensteinFloor2)
	must(t, pictures.Screenshot(o, "floor3"))

	// The third floor, to a chest with uniforms in it, unlocked
	castleWolfensteinPlay(t, o, nil, castleWolfensteinFloor3)
	must(t, pictures.Screenshot(o, "chest"))
	castleWolfensteinPlay(t, o, nil, "-:27 -:2")
	must(t, pictures.Screenshot(o, "uniforms"))
	castleWolfensteinPlay(t, o, nil, "U:2 -:10")
	must(t, pictures.Screenshot(o, "uniform"))

	// In uniform, the rest of the third floor and the fourth
	castleWolfensteinPlay(t, o, nil, castleWolfensteinFloor3Rest)
	must(t, pictures.Screenshot(o, "floor4"))
	castleWolfensteinPlay(t, o, nil, castleWolfensteinFloor4)
	must(t, pictures.Screenshot(o, "floor5"))

	// The fifth floor, and out of the castle by its last room
	castleWolfensteinPlay(t, o, nil, castleWolfensteinFloor5)
	escape := pictures.Record(o)
	escape.Capture(50)
	castleWolfensteinPlay(t, o, escape, castleWolfensteinEscape+" -:2")
	must(t, pictures.SaveRecording(escape, "escape", 300))
	must(t, pictures.Screenshot(o, "escaped"))
	must(t, o.WaitForShownText("CORPORAL", 10))
	must(t, pictures.Screenshot(o, "corporal"))
}

/*
castleWolfensteinPlay presses keys, each followed by a time in seconds, as
"W:0.3": a letter, "_" for the space bar, a key by its name, or "-" for none;
recording the screen every six frames if there is a recording, or every three
for a time that is not a multiple of six frames, as are many of the floors
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
		frames := int(math.Round(s * 60))
		switch {
		case r == nil:
			o.Run(frames)
		case frames%6 == 0:
			r.Run(frames, 6)
		default:
			// The recording runs whole steps of its own: these of three
			// frames, as the times are
			r.Run(frames, 3)
		}
	}
}

// castleWolfensteinFloor2 is the second floor, from the stairs of the first to those of the third
const castleWolfensteinFloor2 = `
P:0.1 L:0.3 -:0.05 E:0.15 -:0.20 E:0.45 -:0.10 E:0.10 -:0.10 E:0.10
-:0.10 D:0.25 -:0.10 O:0.1 L:0.4 A:0.30 -:0.05 E:0.65 Q:0.25 W:0.45
-:0.35 A:0.70 -:0.05 W:0.55 ;:0.1 L:0.4 E:0.60 -:0.15 E:0.75 -:0.20
E:0.50 W:0.60 -:0.15 D:0.75 -:0.10 E:0.40 -:0.10 X:0.55 -:0.20 E:0.70
-:0.25 E:0.15 -:0.30 W:0.25 -:0.20 W:0.65 ;:0.1 L:0.4 -:0.25 S:0.55
E:0.20 -:0.05 E:0.25 -:0.05 D:0.55 -:0.30 W:0.20 -:0.25 W:0.35 -:0.20
;:0.1 L:0.4 X:0.15 -:0.15 E:0.20 -:0.20 -:1.5 C:0.15 S:0.1 ;:0.1 _:3 -:1
-:0.05 X:0.65 D:0.35 W:0.50 S:0.10 W:0.05 D:0.25 .:0.10 L:1.70 X:0.5
S:0.1 .:0.1 _:3 Return:1.5 D:0.10 S:0.05 D:0.35 W:0.20 S:0.05`

// castleWolfensteinFloor3 is the third floor, to the chest with the uniforms, which is unlocked
const castleWolfensteinFloor3 = `
S:2.00 -:0.10 D:0.05 S:0.05 D:0.40 X:0.75 D:0.20 S:0.15 D:2.05 A:0.05
S:2.00 -:0.25 S:0.05 A:0.05 X:0.05 S:0.20 X:0.70 A:0.20 S:0.20 A:0.50
S:0.05 P:0.10 L:0.40 A:0.45 S:0.15 A:0.30 S:0.20 ;:0.10 L:0.40 X:0.10
S:0.05 S:0.25 ;:0.10 L:1.65 -:0.10 X:0.70 S:0.10 X:0.05 A:1.00 S:0.15
S:0.05 K:0.10 L:1.85 -:0.20 A:0.35 W:0.15 S:0.15 W:0.35 A:0.05 K:0.10
L:0.40 A:0.05 S:0.25 X:0.10 S:0.55 X:0.10 S:0.05 X:1.15 -:0.20 A:0.05
S:0.05 A:0.25 S:0.05 -:0.20 A:0.15 S:0.10 A:0.40 W:0.40 A:0.60 S:0.05
W:0.45 I:0.10 L:0.40 W:0.05 D:0.55 S:0.45 X:0.05 D:0.30 S:0.05 ;:0.10
_:1.00 -:1 ;:0.1 L:0.5 D:0.3 S:0.05 X:0.2 S:0.05 D:0.25 S:0.05 W:0.05
S:0.05 K:0.1 _:1`

// castleWolfensteinFloor3Rest is the rest of the third floor, in uniform
const castleWolfensteinFloor3Rest = `
-:0.20 X:0.05 A:0.35 W:0.35 A:0.15 S:0.05 A:0.35 W:0.05 A:0.05 S:0.05
A:0.05 X:0.40 S:0.10 X:0.05 A:0.05 S:2.00 -:0.25 A:0.90 S:0.05 A:0.05
X:0.10 S:0.05 X:0.40 D:0.05 X:0.05 S:2.00 D:0.30 S:0.05 D:1.00 S:0.05
S:2.00 -:0.30 X:0.30 S:0.20 X:0.05 D:0.05 S:0.10 D:0.35 S:0.10 D:1.10
W:0.75 D:0.05 S:2.00 -:0.20 W:0.70 A:0.05 S:0.05 W:0.05 D:0.25 S:0.25
D:0.10 S:0.55 D:0.30 X:0.05 S:0.05 D:0.10 S:0.55 D:0.90 S:0.05 S:2.00
-:0.15 S:0.20 D:0.30 W:0.05 S:0.05 D:0.05 X:0.15 S:0.05 X:0.30 A:0.05
S:0.05 X:0.15 S:0.15 X:0.05 D:0.05 S:0.05 D:0.35 S:0.45 D:0.40 X:0.10
S:0.20 X:0.35 S:2.00 D:0.6 S:0.05 ;:0.1 _:1 -:10 -:0.45 D:0.40 X:0.10
S:0.30 X:0.60 S:0.10 X:0.05 A:1.05 W:0.15 S:0.10 W:0.05 S:0.20 W:0.05
A:0.05 X:0.50 S:0.05 X:0.75 D:0.20 S:0.50 D:0.20 S:0.55 D:0.05 X:0.05
S:2.00 -:0.30 A:0.40 X:0.10 S:0.25 X:0.75 S:0.10 X:0.70 D:0.05 S:0.05
X:0.05 S:0.25 X:0.30 D:0.05 S:0.05 X:0.30 D:0.05 S:0.05 X:0.30 D:0.05
S:0.05 X:0.30 S:0.90 S:0.55 X:0.30 D:0.05 S:0.05 X:0.05 A:0.30 X:0.05
S:0.05 A:0.40 S:2.00 -:0.25 W:0.15 S:0.15 W:0.20 S:0.05 W:0.05 S:0.05
W:0.25 S:0.05`

// castleWolfensteinFloor4 is the fourth floor
const castleWolfensteinFloor4 = `
S:2.00 A:0.05 W:0.25 S:0.15 W:0.05 S:2.00 -:0.15 A:0.05 W:0.20 O:0.10
L:0.40 W:0.05 A:0.20 I:0.10 L:0.40 A:0.20 I:0.10 L:0.40 A:0.25 I:0.10
L:0.40 A:0.20 I:0.10 L:0.40 A:0.30 W:0.05 S:0.05 A:0.30 I:0.10 L:0.40
I:0.10 L:0.40 A:0.05 W:0.35 A:0.05 S:0.05 W:0.60 D:0.50 S:0.55 D:0.55
S:0.05 S:2.00 X:0.15 S:0.05 X:0.05 D:0.80 S:0.15 D:0.40 S:0.10 D:0.05
W:0.15 S:0.15 W:0.35 D:0.05 S:2.00 -:0.05 W:0.40 D:0.05 W:0.05 D:0.30
S:0.25 D:0.05 X:1.90 S:0.05 X:0.05 S:0.20 X:0.05 S:2.00 D:0.6 S:0.05
;:0.1 _:1 -:10 D:0.30 S:0.15 D:0.30 X:0.05 S:0.05 D:0.05 X:0.05 S:0.15
X:0.05 D:0.05 S:0.35 D:0.25 S:0.05 S:2.00 -:0.30 D:0.30 W:0.05 S:0.05
D:0.10 S:0.15 D:0.10 ;:0.10 L:0.40 D:0.05 W:0.30 D:0.05 S:0.05 W:0.05
O:0.10 L:0.40 O:0.10 L:0.40 W:0.05 S:0.10 W:0.05 A:0.05 W:0.35 D:0.15
S:0.35 D:0.30 W:0.05 S:0.05 D:0.30 W:0.05 S:0.05 D:0.30 S:0.45 D:0.30
D:0.4 S:0.05 ;:0.1 _:1 -:10 -:0.25 D:0.35 X:0.15 S:0.05 X:0.05 D:0.05
X:0.20 S:0.15 X:0.05 D:0.05 W:0.55 S:0.25 W:1.05 A:0.05 W:0.10 S:0.05
S:2.00 -:0.40 S:0.50 P:0.10 L:0.40 D:0.60 S:0.30 D:0.10 S:0.10 D:0.05
W:0.75 S:0.05 W:0.05 K:0.10 L:0.40 A:1.2 S:0.05 K:0.1 _:1 -:10 -:0.20
A:0.05 S:0.05 A:0.05 W:1.00 S:0.05 W:0.30 S:0.05 S:2.00 -:0.10 W:0.35
S:0.10 W:0.45 S:0.30 W:0.20 S:0.15 W:0.20 S:0.05 S:2.00 S:0.35 D:0.40
W:0.40 A:0.30 S:0.50 W:0.10 O:0.10 L:0.40 W:0.05 K:0.10 L:0.40 X:0.05
A:0.60 W:0.05 S:0.05 A:0.35 W:0.05 S:0.05 A:0.25 S:0.30 X:0.25 S:0.05
X:0.05 A:0.35 X:0.05 S:0.05 A:0.05 S:2.00 X:0.10 S:0.40 X:1.45 A:1.05
S:0.05 K:0.1 _:1 -:10 A:1.10 S:0.05 A:0.05 W:0.10 S:0.05 W:0.40 A:0.05
S:2.00 -:0.05 W:0.25 S:0.05 W:0.05 S:2.00`

// castleWolfensteinFloor5 is the fifth floor, to the room of the way out
const castleWolfensteinFloor5 = `
-:0.05 A:0.75 Z:0.55 A:0.40 .:0.10 L:0.40 .:0.10 L:0.40 X:0.75 -:0.20
W:0.40 C:0.80 C:0.50 Q:0.75 Z:0.15 /:0.10 L:0.40 -:0.15 C:0.90 D:0.65
D:0.95 S:0.40 -:3 X:0.25 S:0.05 X:0.05 S:2.00 -:0.40 A:0.05 X:0.05
S:0.20 X:0.30 S:0.60 X:0.40 A:0.20 S:0.05 K:0.1 _:1 -:10 A:0.40 X:0.05
S:0.05 X:0.05 A:0.05 X:0.10 S:2.00 D:0.6 S:0.05 ;:0.1 _:1 -:10 -:0.30
;:0.10 L:0.40 D:0.55 S:0.40 D:0.45 S:0.20 ;:0.10 L:0.40 D:0.05 X:0.10
S:0.30 X:0.10 S:0.05 X:0.05 D:0.05 S:0.05 S:2.00 -:0.45 D:0.20 S:0.05
K:0.10 L:0.40 D:0.45 W:0.40 S:0.15 W:1.10 S:0.15 W:0.35 D:0.05 S:0.05
W:0.05 S:2.00 A:0.7 S:0.05 K:0.1 _:1 -:10 -:2 -:0.10 A:0.40 W:0.30
S:0.05 W:0.40 D:0.05 W:0.25 S:0.55 W:0.20 S:0.05 S:2.00 W:0.2 S:0.05
-:0.05 D:0.05 W:0.25 S:0.50 W:0.35 S:0.40 W:0.05 S:0.3 D:0.3 S:0.05
;:0.1 _:1 -:10 -:0.05 D:0.65 S:0.15 D:0.30 X:0.05 S:0.05 D:0.40 X:0.05
S:0.05 X:0.45 S:0.05 X:0.25 S:0.25 X:0.05 D:0.05 S:2.00 -:0.40 D:0.20
S:0.05 D:0.20 S:0.55 D:0.35 S:0.55 D:0.30 X:0.05 S:0.05 D:0.05 W:0.35
S:0.40 D:0.15 S:0.75 D:0.05 X:0.40 S:0.70 X:0.05 S:0.50 A:0.3 S:0.05
K:0.1 _:1 -:10 X:0.05 A:0.20 S:0.05 A:0.05 X:0.40 A:0.05 X:0.05 S:2.00
X:0.05 S:0.05 X:0.05 D:0.05 X:0.30 S:0.10 X:1.10 S:0.05`

// castleWolfensteinEscape is the last room, and out
const castleWolfensteinEscape = `
S:2.00 -:0.45 D:0.60 S:0.55 D:0.30 X:0.05 S:0.05 D:0.10 S:0.05 D:0.40
X:0.05 S:0.05 D:0.40 W:0.05 S:0.05 D:0.45 X:0.30 S:0.45 D:0.3`
