package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// snakeMachine is the machine of the guide, as its command line of izapple2
const snakeMachine = `izapple2 -model _base -board 2plus -cpu 6502 \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps`

// snakeListing is the program of the page, the game of the snake in Applesoft
const snakeListing = "../guides/listings/snake.bas"

// The colours of the low resolution graphics the game uses
const (
	snakeWall = 15
	snakeBody = 12
	snakeHead = 13
	snakeFood = 9
)

/*
snakeScreenshots is a whole game written in Applesoft on an Apple ][+ with no
disk, the snake: the program of the page typed and run, and played by a
player that reads the screen, until it has eaten twenty times, and then
steered into a wall to see the end.
*/
func snakeScreenshots(t *testing.T) {
	pictures := newAlbum("applesoft-snake", album.Color)
	o := start(t, snakeMachine, nil)
	must(t, o.WaitForKeyboard(5))

	// Typed, and listed in part
	must(t, typeListing(o, snakeListing))
	must(t, o.TypeLines("HOME", "LIST 100,200"))
	must(t, o.WaitForPrompt(10))
	must(t, pictures.On(album.Green).Screenshot(o, "listed"))

	// Run, and played: the first meals recorded
	must(t, o.TypeLines("RUN"))
	must(t, o.WaitForText("TURN THE SNAKE", 10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "start"))
	game := pictures.Record(o)
	game.Capture(10)
	player := &snakePlayer{}
	for player.meals < 6 {
		player.step(t, o, game)
	}
	game.Run(30, 6)
	must(t, pictures.SaveRecording(game, "playing", 200))

	// On to twenty meals, unrecorded
	for player.meals < 20 {
		player.step(t, o, nil)
	}
	must(t, pictures.Screenshot(o, "long"))

	// Into the wall on purpose, and the end
	player.crash = true
	for !o.HasText("GAME OVER") {
		player.step(t, o, nil)
	}
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "game-over"))
	must(t, o.Type("N"))
	must(t, o.WaitForText("SCORE", 10))
	must(t, o.WaitForPrompt(10))
}

// loResColor is the colour of a block of the low resolution page 1, x and y
// from 0 to 39
func loResColor(o *operator.Operator, x int, y int) uint8 {
	line := y / 2
	b := o.Apple2().Peek(uint16(0x400 + 0x80*(line%8) + 0x28*(line/8) + x))
	if y%2 == 1 {
		b >>= 4
	}
	return b & 0x0f
}

// snakePlayer plays the game of the snake from what is on the screen
type snakePlayer struct {
	headX, headY int
	dx, dy       int
	meals        int
	crash        bool
	frames       int
}

// The keys of the game, by direction
var snakeKeys = map[[2]int]string{{0, -1}: "I", {0, 1}: "K", {-1, 0}: "J", {1, 0}: "L"}

/*
step runs the machine a frame, recorded every few when there is a recording,
and when the head has moved, chooses where to go next and presses its key:
towards the food, never into the wall or the snake, and never into a space
too small for the snake to get out of
*/
func (p *snakePlayer) step(t *testing.T, o *operator.Operator, r *album.Recording) {
	t.Helper()
	if r != nil {
		r.Run(2, 2)
	} else {
		o.Run(1)
	}
	p.frames += 2
	if p.frames > 20*60*60 {
		t.Fatal("the game went on for too long")
	}
	if o.HasText("GAME OVER") {
		if !p.crash {
			t.Fatalf("the snake died after %v meals", p.meals)
		}
		return
	}

	// The screen: the head, the food, the length
	var grid [40][40]uint8
	hx, hy, fx, fy, length := -1, -1, -1, -1, 0
	for y := range 40 {
		for x := range 40 {
			c := loResColor(o, x, y)
			grid[x][y] = c
			switch c {
			case snakeHead:
				hx, hy = x, y
			case snakeFood:
				fx, fy = x, y
			case snakeBody:
				length++
			}
		}
	}
	if hx < 0 || (hx == p.headX && hy == p.headY) {
		return
	}
	if p.headX != 0 || p.headY != 0 {
		p.dx, p.dy = hx-p.headX, hy-p.headY
		if fx < 0 {
			p.meals++
		}
	}
	p.headX, p.headY = hx, hy
	if p.crash {
		return
	}

	// The best way: open, room enough after it, and nearest to the food
	best, bestDistance := [2]int{}, 1000
	for way := range snakeKeys {
		if way[0] == -p.dx && way[1] == -p.dy {
			continue
		}
		x, y := hx+way[0], hy+way[1]
		if c := grid[x][y]; c == snakeWall || c == snakeBody || c == snakeHead {
			continue
		}
		distance := abs(x-fx) + abs(y-fy)
		if room(grid, x, y) < length+10 {
			distance += 500
		}
		if distance < bestDistance {
			best, bestDistance = way, distance
		}
	}
	if bestDistance < 1000 && (best[0] != p.dx || best[1] != p.dy) {
		must(t, o.Type(snakeKeys[best]))
	}
}

// room counts the free blocks that can be reached from one, up to 400
func room(grid [40][40]uint8, x int, y int) int {
	seen := map[[2]int]bool{}
	pending := [][2]int{{x, y}}
	for len(pending) > 0 && len(seen) < 400 {
		at := pending[len(pending)-1]
		pending = pending[:len(pending)-1]
		if seen[at] {
			continue
		}
		if at[0] < 0 || at[0] > 39 || at[1] < 0 || at[1] > 39 {
			continue
		}
		c := grid[at[0]][at[1]]
		if c == snakeWall || c == snakeBody || c == snakeHead {
			continue
		}
		seen[at] = true
		for way := range snakeKeys {
			pending = append(pending, [2]int{at[0] + way[0], at[1] + way[1]})
		}
	}
	return len(seen)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
