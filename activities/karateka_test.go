package activities

import (
	"image"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// karatekaMachine is the machine of the guide, as its command line of izapple2
const karatekaMachine = `izapple2 -model _base -board 2plus -cpu 6502 \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s6 diskii,disk1=disks/Karateka.woz`

/*
karatekaScreenshots is Karateka, Jordan Mechner's game of 1984, from its
original disk on an Apple ][+: the titles, the prologue, Akuma and the
princess until she faints in her room, and a game started, the karateka
climbing up from the cliff and beating the first guard, played with the
joystick.

The parts of the game are waited for by their colours, counted on the
screen: the orange of the title, the white of the prologue, the blue of the
floors and the sky.
*/
func karatekaScreenshots(t *testing.T) {
	pictures := newAlbum("karateka", album.Color)
	o := start(t, karatekaMachine, nil)
	o.Joystick(127, 127)
	colours := func() dotColours { return countColours(pictures.Screen(o)) }

	// The titles, recorded from the name of Brøderbund to the one of the game
	waitForColours(t, o, colours, 30, func(c dotColours) bool { return c.white > 1000 })
	titles := pictures.Record(o)
	titles.Capture(50)
	for colours().orange < 15000 {
		if o.Frames() > 60*60 {
			t.Fatal("the title did not come")
		}
		titles.Run(30, 30)
	}
	titles.Run(150, 30)
	must(t, pictures.SaveRecording(titles, "titles", 300))

	// The prologue, when it is all on the screen
	waitForColours(t, o, colours, 60, func(c dotColours) bool { return c.white > 14000 })
	must(t, pictures.Screenshot(o, "prologue"))

	// Akuma and the princess, and the princess in her room until she
	// faints, when there is less orange of her standing on the screen
	waitForColours(t, o, colours, 60, func(c dotColours) bool { return c.blue > 17000 })
	dungeon := pictures.Record(o)
	dungeon.Capture(10)
	for c := colours(); c.blue > 15500 || c.orange > 420; c = colours() {
		if o.Frames() > 4*60*60 {
			t.Fatal("the princess did not faint")
		}
		dungeon.Run(6, 6)
	}
	dungeon.Run(2*60, 6)
	must(t, pictures.SaveRecording(dungeon, "akuma", 100))

	// A game started with a key, in the demonstration that follows: the
	// castle, and the cliff
	waitForColours(t, o, colours, 30, func(c dotColours) bool { return c.blue > 100000 })
	must(t, o.Key("Space"))
	waitForColours(t, o, colours, 30, func(c dotColours) bool { return c.white > 30000 })
	o.RunSeconds(1)
	must(t, pictures.Screenshot(o, "castle"))
	waitForColours(t, o, colours, 30, func(c dotColours) bool { return c.blue > 100000 })
	climb := pictures.Record(o)
	climb.Capture(10)
	climb.Run(12*60, 6)

	// Run forward, the joystick up and to the right, and stop in the
	// fighting stance, the joystick let go
	o.Joystick(255, 0)
	climb.Run(3*60, 6)
	o.Joystick(127, 127)
	climb.Run(2*60, 6)
	must(t, pictures.SaveRecording(climb, "climb", 100))

	// The fight, a kick, button 0, every 36 frames, until the guard is down
	fight := pictures.Record(o)
	fight.Capture(10)
	for range 90 {
		o.PressButton(0)
		fight.Run(36, 6)
	}
	fight.Run(60, 6)
	must(t, pictures.SaveRecording(fight, "fight", 300))
	must(t, pictures.Screenshot(o, "won"))
}

// dotColours are how many dots of a screen are of the colours Karateka is
// told apart by
type dotColours struct {
	orange, white, blue int
}

// countColours counts the dots of a screen that are orange, white and blue
func countColours(screen *image.RGBA) dotColours {
	var c dotColours
	pix := screen.Pix
	for i := 0; i < len(pix); i += 4 {
		r, g, b := pix[i], pix[i+1], pix[i+2]
		switch {
		case r > 200 && g > 200 && b > 200:
			c.white++
		case r > 180 && g < 140 && b < 100:
			c.orange++
		case b > 180 && r < 100:
			c.blue++
		}
	}
	return c
}

// waitForColours runs the machine until the colours of its screen pass a
// test, looking every 30 frames
func waitForColours(t *testing.T, o *operator.Operator, colours func() dotColours, seconds float64, done func(dotColours) bool) {
	t.Helper()
	limit := o.Frames() + uint64(seconds*operator.FramesPerSecond)
	for !done(colours()) {
		if o.Frames() > limit {
			t.Fatalf("the screen did not change in %v seconds: %+v", seconds, colours())
		}
		o.Run(30)
	}
}
