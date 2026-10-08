package activities

import (
	"fmt"
	"os"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// roGame reads and drives the game
type roGame struct {
	t *testing.T
	o *operator.Operator
	// target is the robot the player is going into, which is not in its way
	target int
	// light plans paths for the player alone, without what it carries
	light bool
	// ignore are the robots the player is next to as a path starts, and how
	// far from each it is then
	ignore map[int]int
	// rec records the keys pressed, when not nil, recSpeed times faster
	// than the machine
	rec      *album.Recording
	recSpeed int
}

func (g *roGame) peek(addr int) int { return int(g.o.Apple2().Peek(uint16(0x6000 + addr))) }

// place is the room of the player once it holds for a few frames: next to a
// robot it flickers
func (g *roGame) place() int {
	last, same := -1, 0
	for range 60 {
		r := g.room(0)
		if r == last {
			same++
			if same >= 4 {
				return r
			}
		} else {
			last, same = r, 0
		}
		g.o.Run(1)
	}
	return last
}

// The objects: room, x and y
func (g *roGame) room(obj int) int { return g.peek(0x300 + obj) }
func (g *roGame) x(obj int) int    { return g.peek(0x400 + obj) }
func (g *roGame) y(obj int) int    { return g.peek(0x500 + obj) }

// exit is the room through a side of a room: 0 up, 1 down, 2 right, 3 left
func (g *roGame) exit(room, side int) int { return g.peek(0xEC0 + side*0x40 + room) }

// wall says whether a tile of a room is a wall, column 0 to 19 from the
// left, row 0 to 11 from the top
func (g *roGame) wall(room, col, row int) bool {
	i := (row/4)*10 + col/2
	bit := (col%2)*4 + row%4
	return g.peek(0x1000+room*30+i)>>bit&1 == 1
}

// press presses a key, and runs until the player stops moving
func (g *roGame) press(key string) {
	frames := 2
	defer func() {
		if g.rec != nil {
			g.rec.Capture(max(frames*100/operator.FramesPerSecond/g.recSpeed, 2))
		}
	}()
	if len(key) > 1 {
		must(g.t, g.o.Key(key))
	} else {
		must(g.t, g.o.Type(key))
	}
	g.o.Run(2)
	last, still := -1, 0
	for range 120 {
		now := g.room(0)<<16 | g.x(0)<<8 | g.y(0)
		if now == last {
			still++
			if still >= 6 {
				return
			}
		} else {
			last, still = now, 0
		}
		g.o.Run(1)
		frames++
	}
}

// run runs the machine some frames, recorded when recording
func (g *roGame) run(frames int) {
	if g.rec == nil {
		g.o.Run(frames)
		return
	}
	for done := 0; done < frames; done += 6 {
		n := min(6, frames-done)
		g.o.Run(n)
		g.rec.Capture(max(n*100/operator.FramesPerSecond/g.recSpeed, 2))
	}
}

// record starts recording the keys pressed, a picture after each, played
// some times faster than the machine
func (g *roGame) record(pictures *album.Album, faster int) {
	g.rec = pictures.Record(g.o)
	g.recSpeed = faster
	g.rec.Capture(50)
}

// saveRecording writes what was recorded since record, and stops
func (g *roGame) saveRecording(pictures *album.Album, name string) {
	g.o.Run(30)
	g.rec.Capture(10)
	must(g.t, pictures.SaveRecording(g.rec, name, 200))
	g.rec = nil
}

// align moves the player a unit at a time onto the grid of the tiles
func (g *roGame) align() {
	for range 20 {
		dx := g.x(0) % 7
		dy := g.y(0) % 16
		switch {
		case dx != 0 && dx <= 3:
			g.press("Ctrl+J")
		case dx != 0:
			g.press("Ctrl+K")
		case dy != 0 && dy <= 8:
			g.press("Ctrl+M")
		case dy != 0:
			g.press("Ctrl+I")
		default:
			return
		}
	}
	g.t.Fatalf("could not align the player at %v,%v", g.x(0), g.y(0))
}

type roTile struct{ room, col, row int }

// robots are the robots in a room, by their left halves
func (g *roGame) robots(room int) []int {
	var list []int
	for obj := 1; obj < 0xFE; obj++ {
		// The Dalek, that kicks the player back to the start, is one too
		if g.room(obj) == room && (g.peek(0x100+obj) == 0x2A && g.peek(0x600+obj) == 0xFF || obj == 0x22) && obj != g.carried() {
			list = append(list, obj)
		}
	}
	return list
}

// touches says whether boxes come within the reach of a robot, from which
// a step may take the player into it
func (g *roGame) touches(robot int, boxes [][4]int) bool {
	rx, ry := g.x(robot), g.y(robot)
	for _, b := range boxes {
		if b[0] <= rx+19 && b[2] >= rx-7 && b[1] <= ry+15 && b[3] >= ry {
			return true
		}
	}
	return false
}

// distance is how far the player at x,y is from a robot, squared
func (g *roGame) distance(robot, x, y int) int {
	dx, dy := g.x(robot)+6-x-3, g.y(robot)+8-y-8
	return dx*dx + dy*dy
}

// free says whether the player fits at a tile of a room: no wall under it,
// and no robot it would step into
func (g *roGame) free(t roTile) bool {
	px, py := t.col*7, (11-t.row)*16
	boxes := [][4]int{{px, py, px + 6, py + 15}}
	// What the player carries goes through walls: only the player counts
	// The robots of the room, but the one chased and those the player starts
	// from, are in the way
	for _, obj := range g.robots(t.room) {
		if obj == g.target {
			continue
		}
		if d, ok := g.ignore[obj]; ok {
			// Beside it already: past it, but not into its body nearer than
			// at the start
			rx, ry := g.x(obj), g.y(obj)
			body := false
			for _, b := range boxes {
				if b[0] <= rx+12 && b[2] >= rx && b[1] <= ry+15 && b[3] >= ry {
					body = true
				}
			}
			if body && g.distance(obj, px, py) < d {
				return false
			}
			continue
		}
		if g.touches(obj, boxes) {
			return false
		}
	}
	for _, b := range boxes {
		if !g.clear(t.room, b) {
			return false
		}
	}
	return true
}

// clear is whether a box of a room, left, bottom, right and top, has no wall
func (g *roGame) clear(room int, b [4]int) bool {
	for x := b[0]; x <= b[2]; x += 3 {
		for y := b[1]; y <= b[3]; y += 5 {
			for _, p := range [][2]int{{x, y}, {b[2], y}, {x, b[3]}, {b[2], b[3]}} {
				c, r := p[0]/7, 11-p[1]/16
				if p[0] < 0 || p[1] < 0 || c > 19 || r < 0 {
					continue
				}
				if g.wall(room, c, r) {
					return false
				}
			}
		}
	}
	return true
}

// clearAt is whether the player would have no wall at x,y
func (g *roGame) clearAt(x, y int) bool {
	return g.clear(g.room(0), [4]int{x, y, x + 6, y + 15})
}

// tile is where the player is, once aligned
func (g *roGame) tile() roTile {
	return roTile{g.room(0), g.x(0) / 7, 11 - g.y(0)/16}
}

// path finds the keys from the player to a tile, through the rooms
func (g *roGame) path(to roTile) []string {
	keys, ok := g.tryPath(to)
	if !ok {
		g.t.Fatalf("no way from %+v to %+v", g.tile(), to)
	}
	return keys
}

// tryPath finds the keys from the player to a tile, if there is a way
func (g *roGame) tryPath(to roTile) ([]string, bool) {
	return g.search(func(t roTile) bool { return t == to })
}

// nearRobots notes the robots the player touches now, which it can only
// move away from
func (g *roGame) nearRobots() {
	g.ignore = map[int]int{}
	here := [][4]int{{g.x(0), g.y(0), g.x(0) + 6, g.y(0) + 15}}
	for _, r := range g.robots(g.room(0)) {
		if g.touches(r, here) {
			g.ignore[r] = g.distance(r, g.tile().col*7, (11-g.tile().row)*16)
		}
	}
}

// search finds the keys from the player to the nearest tile that is a goal
func (g *roGame) search(goal func(roTile) bool) ([]string, bool) {
	from := g.tile()
	g.nearRobots()
	type step struct {
		prev roTile
		key  string
	}
	seen := map[roTile]step{from: {}}
	queue := []roTile{from}
	moves := []struct {
		key        string
		dc, dr, sd int
	}{{"I", 0, -1, 0}, {"M", 0, 1, 1}, {"K", 1, 0, 2}, {"J", -1, 0, 3}}
	for len(queue) > 0 {
		at := queue[0]
		queue = queue[1:]
		if goal(at) {
			var keys []string
			for at != from {
				s := seen[at]
				keys = append([]string{s.key}, keys...)
				at = s.prev
			}
			return keys, true
		}
		for _, m := range moves {
			next := roTile{at.room, at.col + m.dc, at.row + m.dr}
			switch {
			case next.col < 0:
				next.room, next.col = g.exit(at.room, m.sd), 19
			case next.col > 19:
				next.room, next.col = g.exit(at.room, m.sd), 0
			case next.row < 0:
				next.room, next.row = g.exit(at.room, m.sd), 11
			case next.row > 11:
				next.room, next.row = g.exit(at.room, m.sd), 0
			}
			// An exit to the room itself is none, as the doors of a robot
			if next.room >= 0x40 || next.room == at.room && (next.col != at.col+m.dc || next.row != at.row+m.dr) ||
				!g.free(next) {
				continue
			}
			if _, ok := seen[next]; ok {
				continue
			}
			seen[next] = step{at, m.key}
			queue = append(queue, next)
		}
	}
	return nil, false
}

// goTo walks the player to a tile. Carrying something that the paths of
// boxes don't let through, it plans for the player alone and nudges a few
// units across when a step is blocked, as a player would.
func (g *roGame) goTo(to roTile) {
	for range 400 {
		g.align()
		keys, ok := g.tryPath(to)
		if !ok && g.carried() >= 0 {
			g.light = true
			keys, ok = g.tryPath(to)
			g.light = false
		}
		if os.Getenv("RO_DEBUG") != "" {
			fmt.Printf("    goTo %+v from %+v (%d,%d): %v %v free %v ignore %v\n", to, g.tile(), g.x(0), g.y(0), ok, keys, g.free(g.tile()), g.ignore)
		}
		if !ok {
			// A robot in the way: it moves
			g.run(10)
			continue
		}
		if len(keys) == 0 {
			break
		}
		from := g.place()
		g.step(keys[0])
		if now := g.place(); now != from && now != to.room && roRobotOf[now] != 0 && roRobotOf[now] != roRobotOf[from] {
			// Into a robot on the way: out again
			g.leave()
		}
	}
	if g.tile() != to {
		g.t.Fatalf("the player is at %+v, not %+v", g.tile(), to)
	}
}

// step presses a key for a step, and when the player doesn't move, nudges
// it across the way a unit at a time, up to a few, and tries again
func (g *roGame) step(key string) {
	before := [3]int{g.room(0), g.x(0), g.y(0)}
	g.press(key)
	if [3]int{g.room(0), g.x(0), g.y(0)} != before {
		return
	}
	across := map[string][2]string{"I": {"Ctrl+J", "Ctrl+K"}, "M": {"Ctrl+J", "Ctrl+K"},
		"J": {"Ctrl+I", "Ctrl+M"}, "K": {"Ctrl+I", "Ctrl+M"}}[key]
	back := map[string]string{"Ctrl+J": "Ctrl+K", "Ctrl+K": "Ctrl+J", "Ctrl+I": "Ctrl+M", "Ctrl+M": "Ctrl+I"}
	for _, side := range across {
		for n := 1; n <= 6; n++ {
			g.press(side)
			moved := [3]int{g.room(0), g.x(0), g.y(0)}
			g.press(key)
			if [3]int{g.room(0), g.x(0), g.y(0)} != moved {
				return
			}
		}
		for range 6 {
			g.press(back[side])
		}
	}
}

// roInside is the room inside each robot
var roInside = map[int]int{0xF0: 0x09, 0xF2: 0x0A, 0xF4: 0x0B}

// objectTile is the tile of an object
func (g *roGame) objectTile(obj int) roTile {
	return roTile{g.room(obj), (g.x(obj) + 3) / 7, 11 - (g.y(obj)+8)/16}
}

// chase walks the player into a robot that may move: it closes in on the
// robot from its left or its right, and steps in, a whole tile, when the
// robot is level with it and within a step
func (g *roGame) chase(obj int) {
	g.goToRoom(g.room(obj))
	g.target = obj
	defer func() { g.target = -1 }()
	start := g.place()
	inside := roInside[obj]
	last := [2]int{-1, -1}
	for range 600 {
		r := g.place()
		if r == inside {
			return
		}
		if r != start && roRobotOf[r] != 0 {
			// Into another robot: out again
			g.leave()
			continue
		}
		away := r != g.room(obj)
		at := [2]int{g.x(obj), g.y(obj)}
		still := at == last
		last = at
		dx := g.x(obj) - g.x(0)
		dy := g.y(obj) - g.y(0)
		level := dy >= -8 && dy <= 8
		if os.Getenv("RO_DEBUG") != "" {
			fmt.Printf("  robot %d,%d player %d,%d dx %d dy %d\n", g.x(obj), g.y(obj), g.x(0), g.y(0), dx, dy)
		}
		near := dy >= -24 && dy <= 24
		// From the right, a whole tile toward it from 17 to 19 units away
		fromRight := g.x(obj)+19 <= 133 && !g.wall(g.room(0), (g.x(obj)+19)/7, g.tile().row)
		switch {
		case away:
			// The way to it goes round through other rooms
			g.towards(obj)
		case level && dx >= 5 && dx <= 7 && dy > 4 && g.clearAt(g.x(0), g.y(0)+1):
			// At the edge of level: a unit nearer first
			g.press("Ctrl+I")
		case level && dx >= 5 && dx <= 7 && dy < -4 && g.clearAt(g.x(0), g.y(0)-1):
			g.press("Ctrl+M")
		case level && dx >= 5 && dx <= 7:
			g.press("K")
		case fromRight && level && dx <= -17 && dx >= -19:
			g.press("J")
		case fromRight && near && dx < -19 && dx >= -26:
			g.press("Ctrl+J")
		case near && dx > 7 && dx <= 14:
			g.press("Ctrl+K")
		case near && dx > -7 && dx < 5:
			// Too close to step in: back off to the left
			g.press("Ctrl+J")
		case near && dx >= 5 && dx <= 7 && still && dy > 0 && g.clearAt(g.x(0), g.y(0)+1):
			// Beside it, and it stays: up or down to be level
			g.press("Ctrl+I")
		case near && dx >= 5 && dx <= 7 && still && dy < 0 && g.clearAt(g.x(0), g.y(0)-1):
			g.press("Ctrl+M")
		case near && dx >= 5 && dx <= 7:
			// Beside it: wait for it to be level
			g.run(1)
		default:
			g.towards(obj)
		}
	}
	g.t.Fatalf("could not get into robot %02X", obj)
}

// leave walks the player out of the room it is in by the nearest opening
// in its edge, as out of a robot
func (g *roGame) leave() {
	from := g.place()
	for _, side := range []string{"", "M", "I", "J", "K"} {
		if g.tryLeave(side) {
			return
		}
	}
	g.t.Fatalf("the player did not leave room %02X", from)
}

// leaveBy leaves the robot the player is in through the opening of a side
func (g *roGame) leaveBy(side string) {
	from := g.place()
	if !g.tryLeave(side) {
		g.t.Fatalf("the player did not leave room %02X", from)
	}
}

// tryLeave leaves the robot the player is in through the opening of a
// side, I, M, K or J, or the nearest one when empty, and says if it did
func (g *roGame) tryLeave(side string) bool {
	g.align()
	from := g.place()
	start := g.tile()
	g.nearRobots()
	type step struct {
		prev roTile
		key  string
	}
	seen := map[roTile]step{start: {}}
	queue := []roTile{start}
	moves := []struct {
		key    string
		dc, dr int
	}{{"I", 0, -1}, {"M", 0, 1}, {"K", 1, 0}, {"J", -1, 0}}
	for len(queue) > 0 {
		at := queue[0]
		queue = queue[1:]
		for _, m := range moves {
			c, r := at.col+m.dc, at.row+m.dr
			if c < 0 || c > 19 || r < 0 || r > 11 {
				if side != "" && m.key != side {
					continue
				}
				var keys []string
				for p := at; p != start; p = seen[p].prev {
					keys = append([]string{seen[p].key}, keys...)
				}
				for _, k := range append(keys, m.key) {
					if g.room(0) != from {
						break
					}
					g.press(k)
					if os.Getenv("RO_DEBUG") != "" {
						fmt.Printf("    leave %v: player %d,%d room %02X carried %02X at %d,%d\n", k, g.x(0), g.y(0), g.room(0), g.carried()&0xFF, g.x(0xFA), g.y(0xFA))
					}
				}
				if g.room(0) == from {
					return false
				}
				// Out beside the robot: on the same way until clear of it
				out := g.room(0)
				for range 4 {
					clear := true
					for _, r := range g.robots(out) {
						if g.touches(r, [][4]int{{g.x(0), g.y(0), g.x(0) + 6, g.y(0) + 15}}) {
							clear = false
						}
					}
					if clear {
						break
					}
					// On the same way, or across when that leaves the room,
					// but away from the robots
					key := m.key
					steps := map[string][2]int{"I": {0, 16}, "M": {0, -16}, "K": {7, 0}, "J": {-7, 0}}
					away := func(x, y int) bool {
						for _, r := range g.robots(out) {
							if g.distance(r, x, y) <= g.distance(r, g.x(0), g.y(0)) {
								return false
							}
						}
						return true
					}
					inside := func(k string) (int, int, bool) {
						x, y := g.x(0)+steps[k][0], g.y(0)+steps[k][1]
						return x, y, x >= 0 && x <= 133 && y >= 0 && y <= 176
					}
					if _, _, ok := inside(m.key); !ok {
						for _, k := range []string{"I", "M", "K", "J"} {
							if x, y, ok := inside(k); ok && g.clearAt(x, y) && away(x, y) {
								key = k
								break
							}
						}
					}
					g.press(key)
					if g.place() != out {
						g.t.Fatalf("leaving room %02X the player went into room %02X", from, g.room(0))
					}
				}
				return true
			}
			next := roTile{from, c, r}
			if _, ok := seen[next]; ok || !g.free(next) {
				continue
			}
			seen[next] = step{at, m.key}
			queue = append(queue, next)
		}
	}
	return false
}

// carried is the object the player carries, from $091F, or -1
func (g *roGame) carried() int {
	c := int(g.o.Apple2().Peek(0x091F))
	if c == 0xFF {
		return -1
	}
	return c
}

// state is a line on where the player is and what it carries
func (g *roGame) state() string {
	return fmt.Sprintf("room %02X x=%d y=%d tile %d,%d carries %02X key %02X magnet %02X %d,%d by %02X orange %02X %d,%d",
		g.room(0), g.x(0), g.y(0), g.x(0)/7, 11-g.y(0)/16, g.carried()&0xFF, g.room(0xFA),
		g.room(0xF6), g.x(0xF6), g.y(0xF6), g.peek(0x6F6), g.room(0xF0), g.x(0xF0), g.y(0xF0)) +
		func() string {
			if c := g.carried(); c >= 0 {
				return fmt.Sprintf(" offset %d,%d", g.x(c)-g.x(0), g.y(c)-g.y(0))
			}
			return ""
		}()
}

// pick picks up an object in the room: the player goes to its tile, then
// tries Space at positions around it, a unit at a time, until it carries it
func (g *roGame) pick(obj int) {
	g.goTo(g.objectTile(obj))
	moves := []string{"Ctrl+J", "Ctrl+M", "Ctrl+K", "Ctrl+I"}
	for ring := 1; ring <= 10; ring++ {
		for i, m := range moves {
			n := ring
			if i%2 == 1 {
				n = ring * 2
			}
			for range n {
				g.press(" ")
				if os.Getenv("RO_DEBUG") != "" {
					fmt.Printf("    space at %d,%d: object %d,%d by %02X carried %02X\n", g.x(0), g.y(0), g.x(obj), g.y(obj), g.peek(0x600+obj), g.carried()&0xFF)
				}
				if g.carried() == obj {
					return
				}
				if other := g.carried(); other >= 0 {
					// Another one on top: aside, down, and back
					x, y := g.x(0), g.y(0)
					for _, k := range []string{"Ctrl+K", "Ctrl+J", "Ctrl+I", "Ctrl+M"} {
						for range 10 {
							g.press(k)
						}
						if g.x(other) != g.x(obj) || g.y(other) != g.y(obj) {
							break
						}
					}
					g.press(" ")
					for range 30 {
						switch {
						case g.x(0) < x:
							g.press("Ctrl+K")
						case g.x(0) > x:
							g.press("Ctrl+J")
						case g.y(0) < y:
							g.press("Ctrl+I")
						case g.y(0) > y:
							g.press("Ctrl+M")
						}
					}
					continue
				}
				g.press(m)
			}
		}
	}
	g.t.Fatalf("could not pick up object %02X at %d,%d, the player at %d,%d", obj, g.x(obj), g.y(obj), g.x(0), g.y(0))
}

// roRobotOf is the robot whose inside is a room
var roRobotOf = map[int]int{0x09: 0xF0, 0x0A: 0xF2, 0x0B: 0xF4}

// pickRobot picks up a robot: the player goes to stand two tiles to its
// right, on its row, then comes closer a unit at a time with Space after
// each, until it carries the robot
func (g *roGame) pickRobot(obj int) {
	g.pickRobotFrom(obj, true)
	if g.x(obj) >= g.x(0) {
		return
	}
	// Held from its right, it would not go into another robot with the
	// player: down again where it is open to its left, and up from there
	keys, ok := g.search(func(t roTile) bool {
		if t.room != g.room(0) || t.col < 5 || t.col > 15 {
			return false
		}
		for c := t.col - 5; c <= t.col+3; c++ {
			if !g.free(roTile{t.room, c, t.row}) {
				return false
			}
		}
		return true
	})
	if !ok {
		return
	}
	for _, k := range keys {
		g.step(k)
	}
	g.press(" ")
	g.pickRobotFrom(obj, true)
}

// pickRobotFrom picks up a robot from its left, or else its right, or the
// other way round
func (g *roGame) pickRobotFrom(obj int, leftFirst bool) {
	robot := g.objectTile(obj)
	type side struct {
		dc, dr int
		move   string
	}
	var left, right []side
	// On its row, or one off where a wall is in the way
	for _, dr := range []int{0, 1, -1} {
		left = append(left, side{-3, dr, "Ctrl+K"}, side{-2, dr, "Ctrl+K"}, side{-1, dr, "Ctrl+K"})
		right = append(right, side{4, dr, "Ctrl+J"}, side{3, dr, "Ctrl+J"})
	}
	sides := append(left, right...)
	if !leftFirst {
		sides = append(right, left...)
	}
	for _, side := range sides {
		target := roTile{robot.room, robot.col + side.dc, robot.row + side.dr}
		if target.row < 0 || target.row > 11 {
			continue
		}
		// Beside it, but not into it on the way
		g.target = -1
		_, reach := g.tryPath(target)
		if !reach {
			g.target = obj
			if g.free(target) {
				_, reach = g.tryPath(target)
			}
			g.target = -1
			if !reach || g.touches(obj, [][4]int{{target.col * 7, (11 - target.row) * 16, target.col*7 + 6, (11-target.row)*16 + 15}}) && g.distance(obj, target.col*7, (11-target.row)*16) < 60 {
				if !reach {
					continue
				}
			}
			g.target = obj
		}
		if os.Getenv("RO_DEBUG") != "" {
			fmt.Printf("    robot %02X at %d,%d tile %+v side %+v free %v reach %v\n", obj, g.x(obj), g.y(obj), robot, target, target.col >= 0 && target.col <= 19 && g.free(target), reach)
		}
		if target.col < 0 || target.col > 19 || !reach {
			g.target = -1
			continue
		}
		g.goTo(target)
		g.target = -1
		for range 21 {
			g.press(" ")
			if os.Getenv("RO_DEBUG") != "" {
				fmt.Printf("    pick: player %d,%d robot %d,%d carried %02X\n", g.x(0), g.y(0), g.x(obj), g.y(obj), g.carried()&0xFF)
			}
			if g.carried() == obj {
				return
			}
			if g.carried() >= 0 {
				// Something else came up: down again
				g.press(" ")
			}
			px := g.x(0)
			g.press(side.move)
			if g.x(0) == px && g.place() == robot.room {
				// A wall in the way
				break
			}
			if g.place() != robot.room {
				// Into it instead: out again, and from another side
				g.leave()
				if g.place() != robot.room {
					g.t.Fatalf("the player went into the robot %02X instead of picking it up", obj)
				}
				break
			}
		}
	}
	g.t.Fatalf("could not pick up robot %02X", obj)
}

// wiggle moves the player a unit at a time around where it is, in a growing
// square, until done says it is done
func (g *roGame) wiggle(done func() bool) bool {
	moves := []string{"Ctrl+J", "Ctrl+M", "Ctrl+K", "Ctrl+I"}
	for ring := 1; ring <= 12; ring++ {
		for i, m := range moves {
			n := ring
			if i%2 == 1 {
				n++
			}
			for range n {
				if done() {
					return true
				}
				g.press(m)
			}
		}
	}
	return done()
}

// goToRoom walks the player into a room, as far as its first tile
func (g *roGame) goToRoom(room int) {
	for range 400 {
		if g.place() == room {
			return
		}
		g.align()
		keys, ok := g.search(func(t roTile) bool { return t.room == room })
		if !ok {
			g.run(10)
			continue
		}
		if len(keys) > 0 {
			g.press(keys[0])
		}
	}
	g.t.Fatalf("could not get to room %02X", room)
}

// meet walks the player toward an object that wanders from room to room,
// a step at a time, until done says it is done
func (g *roGame) meet(obj int, done func() bool) {
	for range 600 {
		if done() {
			return
		}
		g.align()
		there := g.objectTile(obj)
		if os.Getenv("RO_DEBUG") != "" {
			fmt.Printf("  meet: %02X in %02X at %d,%d, player in %02X at %d,%d carries %02X\n", obj, g.room(obj), g.x(obj), g.y(obj), g.room(0), g.x(0), g.y(0), g.carried()&0xFF)
		}
		keys, ok := g.search(func(t roTile) bool {
			return t.room == there.room && abs(t.col-there.col) <= 1 && t.row == there.row
		})
		if !ok || len(keys) == 0 {
			g.run(4)
			continue
		}
		g.step(keys[0])
	}
	g.t.Fatalf("could not meet object %02X", obj)
}

// putRobot puts the robot the player carries down with its left half at x,y:
// the player stands where the robot it holds comes there. When that is a
// wall, it puts the robot down, takes it from its other side, and tries
// again.
func (g *roGame) putRobot(robot, x, y int) {
	// Just in, the robot may be still in the room before: further in
	for range 4 {
		if g.room(robot) == g.room(0) {
			break
		}
		t := g.tile()
		switch {
		case t.row <= 1:
			g.step("M")
		case t.row >= 10:
			g.step("I")
		case t.col <= 1:
			g.step("K")
		default:
			g.step("J")
		}
	}
	for try := range 3 {
		ox, oy := g.x(robot)-g.x(0), g.y(robot)-g.y(0)
		px, py := x-ox, y-oy
		stand := roTile{g.room(0), (px + 3) / 7, 11 - (py+8)/16}
		if os.Getenv("RO_DEBUG") != "" {
			fmt.Printf("    put: try %d offset %d,%d stand %d,%d tile %+v free %v player %d,%d carries %02X\n", try, ox, oy, px, py, stand, g.free(stand), g.x(0), g.y(0), g.carried()&0xFF)
		}
		if stand.col >= 0 && stand.col <= 19 && stand.row >= 0 && stand.row <= 11 && g.free(stand) {
			if _, ok := g.tryPath(stand); ok {
				g.goTo(stand)
				// The last units, one at a time
				for range 20 {
					switch {
					case g.x(0) < px:
						g.press("Ctrl+K")
					case g.x(0) > px:
						g.press("Ctrl+J")
					case g.y(0) < py:
						g.press("Ctrl+I")
					case g.y(0) > py:
						g.press("Ctrl+M")
					}
				}
				g.press(" ")
				return
			}
		}
		if try == 2 {
			break
		}
		g.press(" ")
		g.pickRobotFrom(robot, ox < 0)
	}
	g.t.Fatalf("could not put robot %02X at %d,%d", robot, x, y)
}

// moveCarried moves the player until what it carries is at x,y, across
// first, or up or down when a wall is in the way
func (g *roGame) moveCarried(x, y int) {
	obj := g.carried()
	for range 60 {
		var moves []string
		ox, oy := g.x(obj), g.y(obj)
		if ox < x {
			moves = append(moves, "Ctrl+K")
		} else if ox > x {
			moves = append(moves, "Ctrl+J")
		}
		if oy < y {
			moves = append(moves, "Ctrl+I")
		} else if oy > y {
			moves = append(moves, "Ctrl+M")
		}
		if len(moves) == 0 {
			return
		}
		moved := false
		for _, m := range moves {
			px, py := g.x(0), g.y(0)
			g.press(m)
			if g.x(0) != px || g.y(0) != py {
				moved = true
				break
			}
		}
		if !moved {
			return
		}
	}
}

// unlock opens a door with the key the player carries: the key goes into
// its lock, at x,y, from the left, a unit at a time until the wall at
// col,row of the room goes, and out again. A lock opens or closes its door
// as the key goes in, so the door is checked once the key is out.
func (g *roGame) unlock(x, y, room, col, row int) {
	key := g.carried()
	if g.x(key) > g.x(0) {
		// Held on the right, it would go into the lock again on the way
		// back: down, and up again from its right
		g.press(" ")
		for range g.x(key) - g.x(0) + 2 {
			g.press("Ctrl+K")
		}
		g.press(" ")
		if g.carried() != key {
			g.t.Fatalf("could not take the key %02X again", key)
		}
	}
	// The key may have gone in on the way: the door slides a while
	g.moveCarried(x-4, y)
	g.run(5 * operator.FramesPerSecond)
	if !g.wall(room, col, row) {
		return
	}
	for range 4 {
		g.moveCarried(x-4, y)
		for range 6 {
			g.press("Ctrl+K")
			if !g.wall(room, col, row) {
				break
			}
		}
		g.moveCarried(x-4, y)
		g.run(5 * operator.FramesPerSecond)
		if os.Getenv("RO_DEBUG") != "" {
			fmt.Printf("    unlock: the key at %d,%d, the wall %v\n", g.x(key), g.y(key), g.wall(room, col, row))
		}
		if !g.wall(room, col, row) {
			return
		}
	}
	g.t.Fatalf("the door of room %02X did not open", room)
}

// towards takes a step toward a tile beside a robot, to its left or right
func (g *roGame) towards(obj int) {
	g.align()
	robot := g.objectTile(obj)
	moved := false
	for _, dc := range []int{-1, 3} {
		target := roTile{robot.room, robot.col + dc, robot.row}
		if target.col < 0 || target.col > 19 || g.wall(target.room, target.col, target.row) {
			continue
		}
		keys, ok := g.tryPath(target)
		if os.Getenv("RO_DEBUG") != "" {
			fmt.Printf("    path to %+v: %v %v ignore %v free %v\n", target, ok, keys, g.ignore, g.free(g.tile()))
		}
		if ok && len(keys) > 0 {
			g.press(keys[0])
			moved = true
			break
		}
	}
	if !moved {
		g.run(2)
	}
}

// switchRobot turns the switch of the robot the player is in, in its lower
// right corner
func (g *roGame) switchRobot() {
	g.goTo(roTile{g.room(0), 17, 9})
	g.press(" ")
}

// remote turns the remote control on or off, R and Space, which stops all
// the robots or lets them go on, and turns the player back with C
func (g *roGame) remote() {
	must(g.t, g.o.Type("R"))
	g.o.Run(10)
	g.press(" ")
	must(g.t, g.o.Type("C"))
	g.o.Run(10)
}

// waitFor runs the machine until done says so, or fails after some seconds;
// with a recording, it is recorded
func (g *roGame) waitFor(what string, seconds int, rec *album.Recording, done func() bool) {
	for i := 0; !done(); i++ {
		if i > 2*seconds {
			g.t.Fatalf("%v did not happen in %v seconds", what, seconds)
		}
		if rec != nil {
			rec.Run(operator.FramesPerSecond/2, 6)
		} else {
			g.o.RunSeconds(0.5)
		}
	}
}

// fetched is what a guard keeps in a room, the magnet or the energy
// crystal, whichever the game put there
func (g *roGame) fetched(room int) int {
	for _, obj := range []int{roMagnet, roCrystal} {
		if g.room(obj) == room {
			return obj
		}
	}
	g.t.Fatalf("neither the magnet nor the crystal is in room %02X", room)
	return -1
}

// regrip puts down what the player carries and takes it again, dy higher
// than its feet
func (g *roGame) regrip(dy int) {
	obj := g.carried()
	g.press(" ")
	px, py := g.x(0), g.y(obj)-dy
	for range 40 {
		switch {
		case g.y(0) < py:
			g.press("Ctrl+I")
		case g.y(0) > py:
			g.press("Ctrl+M")
		case g.x(0) < px:
			g.press("Ctrl+K")
		case g.x(0) > px:
			g.press("Ctrl+J")
		}
	}
	g.press(" ")
	if g.carried() != obj {
		g.t.Fatalf("could not take %02X again, the player at %d,%d", obj, g.x(0), g.y(0))
	}
}

// dropAt puts down what the player carries at x,y
func (g *roGame) dropAt(x, y int) {
	g.moveCarried(x, y)
	g.press(" ")
}

// stand takes the player to x,y a unit at a time, and stops if the room
// changes on the way
func (g *roGame) stand(x, y int) {
	room := g.room(0)
	for range 40 {
		switch {
		case g.room(0) != room:
			return
		case g.x(0) < x:
			g.press("Ctrl+K")
		case g.x(0) > x:
			g.press("Ctrl+J")
		case g.y(0) < y:
			g.press("Ctrl+I")
		case g.y(0) > y:
			g.press("Ctrl+M")
		default:
			return
		}
	}
}

// stash takes an object into the white robot, wherever it is, and puts it
// down at x,y in there
func (g *roGame) stash(obj, x, y int) {
	g.pick(obj)
	g.chase(roWhite)
	g.dropAt(x, y)
	g.leave()
}

// collect puts the magnet into the white robot, and the energy crystal into
// the blue robot, inside the white one, where its crystal sensor does not see
// it, wherever the robots that fetched them let go of them
func (g *roGame) collect() {
	if roRobotOf[g.room(roMagnet)] == 0 {
		g.stash(roMagnet, 50, 45)
	}
	if g.room(roCrystal) != roInside[roBlue] {
		if g.room(roCrystal) == roInside[roWhite] {
			g.chase(roWhite)
		}
		g.pick(roCrystal)
		if g.place() != roInside[roWhite] {
			g.chase(roWhite)
		}
		g.chase(roBlue)
		g.dropAt(60, 100)
		g.leave()
		g.leave()
	}
}

// The objects of the Sewer: the player, the robots, and what they find
const (
	roPlayer      = 0x00
	roMagnet      = 0x03
	roOrange      = 0xF0
	roWhite       = 0xF2
	roBlue        = 0xF4
	roCrystal     = 0xF6
	roKey         = 0xFA
	roChip        = 0xE7
	roTokenSensor = 0x13
	roDirection   = 0x0F
)
