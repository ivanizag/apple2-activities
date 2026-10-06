package activities

import (
	"testing"

	"github.com/ivanizag/izapple2/screen"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// totalReplayMachine is the machine of the guide, as its command line of izapple2
const totalReplayMachine = `izapple2 -model none -board 2e -cpu 65c02 -screen color \
    -rom "<internal>/Apple2e_Enhanced.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s2 vidhd \
    -s7 'smartport,image1=disks/Total Replay v6.1.hdv'`

/*
totalReplayScreenshots is Total Replay, hundreds of games on one hard disk, on
an enhanced Apple //e with a VidHD: the launcher, the box art of its attract
mode in Super Hi-Res, a game found by typing its name, and started.
*/
func totalReplayScreenshots(t *testing.T) {
	pictures := newAlbum("total-replay", album.Color)
	o := start(t, totalReplayMachine, nil)

	// The launcher
	must(t, o.WaitForKeyboard(30))
	o.RunSeconds(5)
	must(t, pictures.Screenshot(o, "launcher"))

	// Left alone, the attract mode: box art in Super Hi-Res
	waitForMode(t, o, screen.VideoSHR, 120)
	for coloured(pictures.Screen(o).Pix) < 100000 {
		if o.Frames() > 300*60 {
			t.Fatal("no box art was drawn")
		}
		o.Run(30)
	}
	o.RunSeconds(1)
	must(t, pictures.Screenshot(o, "box-art"))

	// A game found by typing its name
	must(t, o.Key("Escape"))
	o.RunSeconds(5)
	must(t, o.Type("karateka"))
	o.RunSeconds(2)
	must(t, pictures.Screenshot(o, "search"))

	// And started, its box art first
	must(t, o.Key("Return"))
	waitForMode(t, o, screen.VideoSHR, 30)
	o.RunSeconds(2)
	must(t, pictures.Screenshot(o, "karateka-box"))
	o.RunSeconds(20)
	must(t, pictures.Screenshot(o, "karateka"))
}

// waitForMode runs the machine until the screen is in a video mode, the base
// mode of izapple2's screen package
func waitForMode(t *testing.T, o *operator.Operator, mode uint32, seconds float64) {
	t.Helper()
	if !o.WaitUntil(seconds, func() bool {
		return o.Apple2().GetVideoSource().GetCurrentVideoMode()&screen.VideoBaseMask == mode
	}) {
		t.Fatalf("the screen did not change to the mode %x in %v seconds", mode, seconds)
	}
}

// coloured counts the dots of a screen that are not black
func coloured(pix []uint8) int {
	count := 0
	for i := 0; i < len(pix); i += 4 {
		if pix[i] > 40 || pix[i+1] > 40 || pix[i+2] > 40 {
			count++
		}
	}
	return count
}
