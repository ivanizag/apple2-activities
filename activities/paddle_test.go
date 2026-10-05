package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// paddleGame is the program of the paddle game, as typed
var paddleGame = []string{
	"10 GR : HOME",
	"20 X = 20:Y = 5:DX = 1:DY = 1:S = 0:Q = -1",
	"30 COLOR= 15: HLIN 0,39 AT 0: VLIN 0,39 AT 0: VLIN 0,39 AT 39",
	"40 P = INT ( PDL (0) * 32 / 255) + 1",
	"50 IF P <> Q THEN COLOR= 0: HLIN 1,38 AT 38: COLOR= 13: HLIN P,P + 5 AT 38:Q = P",
	"60 COLOR= 0: PLOT X,Y",
	"70 IF X + DX < 1 OR X + DX > 38 THEN DX = - DX",
	"80 IF Y + DY < 1 THEN DY = - DY",
	`90 IF Y + DY = 38 AND X + DX >= P AND X + DX <= P + 5 THEN DY = - DY:S = S + 1: VTAB 22: PRINT "SCORE ";S`,
	"100 X = X + DX:Y = Y + DY",
	`110 IF Y > 38 THEN VTAB 22: PRINT "GAME OVER, SCORE ";S: END`,
	"120 COLOR= 9: PLOT X,Y",
	"130 GOTO 40",
}

/*
loResFind looks for the first block of a colour on the low resolution screen,
page 1, in the 40 rows above the text of mixed mode, and tells where it is.
Each byte of the text page is two blocks, one above the other: the low four
bits are the top one.
*/
func loResFind(o *operator.Operator, color uint8) (x int, y int, found bool) {
	for y := range 40 {
		line := y / 2
		base := uint16(0x400 + 0x80*(line%8) + 0x28*(line/8))
		for x := range 40 {
			b := o.Apple2().Peek(base + uint16(x))
			if y%2 == 1 {
				b >>= 4
			}
			if b&0x0f == color {
				return x, y, true
			}
		}
	}
	return 0, 0, false
}

// batPaddle is where to turn the paddle for the bat of the game to be under
// a place across the screen, from 1 to 38
func batPaddle(x int) uint8 {
	// The bat starts at INT(PDL(0) * 32 / 255) + 1 and is six blocks wide
	p := min(max(x-2, 1), 33)
	return uint8(min((p-1)*255/32+4, 255))
}

// paddleMachine is the machine of the guide, as its command line of izapple2
const paddleMachine = `izapple2 -model _base -board 2plus -cpu 6502 \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps`

/*
paddleScreenshots is a game written in Applesoft on an Apple ][+ with no disk:
low resolution graphics, a bat on a paddle and a ball, typed, run and played,
on a colour television.
*/
func paddleScreenshots(t *testing.T) {
	pictures := newAlbum("paddle-game", album.Color)
	o := start(t, paddleMachine, nil)
	must(t, o.WaitForKeyboard(5))

	// The program, typed, and its end listed
	must(t, o.TypeLines(paddleGame...))
	must(t, o.TypeLines("HOME", "LIST 60,130"))
	o.Run(30)
	must(t, pictures.Screenshot(o, "listing"))

	// Run, and played by following the ball with the bat
	o.TurnPaddle(0, batPaddle(20))
	must(t, o.TypeLines("RUN"))
	o.Run(30)
	follow := func() {
		if x, _, ok := loResFind(o, loResOrange); ok {
			o.TurnPaddle(0, batPaddle(x))
		}
	}
	playing := pictures.Record(o)
	playing.Capture(4)
	for range 20 * 60 / 4 {
		for range 4 {
			follow()
			o.Run(1)
		}
		playing.Capture(4 * 100 / 60)
	}
	must(t, pictures.SaveRecording(playing, "playing", 100))

	// The bat taken away from the ball, which goes past it
	for !o.HasText("GAME OVER") {
		if x, y, ok := loResFind(o, loResOrange); ok && y > 30 {
			o.TurnPaddle(0, batPaddle((x+20)%38+1))
		}
		o.Run(1)
		if o.Frames() > 10*60*60 {
			t.Fatal("the game did not end")
		}
	}
	o.Run(30)
	must(t, pictures.Screenshot(o, "game-over"))
}

// loResOrange is the colour of the ball, COLOR= 9
const loResOrange = 9
