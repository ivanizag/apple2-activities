package activities

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
	xdraw "golang.org/x/image/draw"
)

// robotOdysseyMachine is the machine of the guide, as its command line of
// izapple2: an Apple ][+ with a Language Card, for its 64 KB, and side A of
// Robot Odyssey in drive 1
const robotOdysseyMachine = `izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s0 language \
    -s6 "diskii,disk1=disks/Robot Odyssey v1.1 (4am crack) side A.dsk"`

/*
robotOdysseySewerScreenshots plays the first level of Robot Odyssey, the
Sewer, on an Apple ][+, from the start to the transporter to the Subway, with
all that the next levels need: the three robots, the blue key, the magnet, the
energy crystal, the two sensors of the subway token and the first chip.

The game is driven by its memory, where the world is, from $6000: the rooms,
their walls and exits, and where each object is. The player is walked a tile
at a time on paths found around the walls and the robots, as roGame does.
Where some things are is chosen at random by the game: the chip, and which of
two rooms has the magnet and which the crystal, so each is looked for where
it is.
*/
func robotOdysseySewerScreenshots(t *testing.T) {
	pictures := newAlbum("robotodyssey-sewer", album.Color)
	o := start(t, robotOdysseyMachine, nil)
	g := &roGame{t: t, o: o, target: -1}
	// A picture once the screen is drawn: a room takes a moment
	shot := func(name string) {
		o.Run(30)
		must(t, pictures.Screenshot(o, name))
	}

	// The drive, the title, which comes out of a growing circle, and the
	// menu
	intro := pictures.Record(o)
	intro.Capture(10)
	intro.Run(20*operator.FramesPerSecond, 4)
	must(t, pictures.SaveRecording(intro, "intro", 300))
	shot("menu")
	must(t, o.Type("\n"))
	o.RunSeconds(40)
	shot("ready")

	// The game starts with a dream, and the fall into Robotropolis
	must(t, o.Type("\n"))
	// Recorded from when the dream is drawn, the question gone: the screen
	// has more than its few letters lit then
	lit := func() int {
		n := 0
		pix := pictures.Screen(o).Pix
		for i := 0; i < len(pix); i += 4 {
			if pix[i]|pix[i+1]|pix[i+2] > 0x40 {
				n++
			}
		}
		return n
	}
	asked := lit()
	waited := 0
	for ; waited < 30*operator.FramesPerSecond && lit() < 4*asked; waited += 6 {
		o.Run(6)
	}
	dream := pictures.Record(o)
	dream.Capture(10)
	dream.Run(90*operator.FramesPerSecond-waited, 6)
	must(t, pictures.SaveRecording(dream, "dream", 300))
	shot("start")
	rooms := map[int]image.Image{0x28: pictures.Screen(o)}

	// The three robots, switched off, so that they stay where they are; the
	// first one recorded
	g.goTo(roTile{0x02, 0, 9})
	shot("robots")
	rooms[0x02] = pictures.Screen(o)
	for i, robot := range []struct {
		obj  int
		name string
	}{{roOrange, "orange"}, {roBlue, "blue"}, {roWhite, "white"}} {
		if i == 0 {
			g.record(pictures, 2)
		}
		g.chase(robot.obj)
		shot("inside-" + robot.name)
		g.switchRobot()
		if i == 0 {
			g.saveRecording(pictures, "switch")
		}
		g.leave()
	}

	// The blue key, in one of them
	g.record(pictures, 2)
	g.chase(roRobotOf[g.room(roKey)])
	g.pick(roKey)
	g.leave()
	g.saveRecording(pictures, "key")
	g.press(" ")

	// The orange and blue robots into the white one
	g.record(pictures, 3)
	for _, robot := range []struct{ obj, x, y int }{{roOrange, 35, 130}, {roBlue, 85, 130}} {
		g.pickRobot(robot.obj)
		g.chase(roWhite)
		g.putRobot(robot.obj, robot.x, robot.y)
		g.leave()
	}
	g.saveRecording(pictures, "pack")

	// The door of the City Sewer, opened with the key, and the key into the
	// white robot
	g.pick(roKey)
	g.record(pictures, 2)
	g.goTo(roTile{0x03, 2, 10})
	g.unlock(17, 16, 0x03, 4, 9)
	g.saveRecording(pictures, "door")
	g.chase(roWhite)
	g.press(" ")
	shot("white")
	g.leave()

	// Round the Sewer, a picture of each room for the map, but the one under
	// the grate and the two past it, taken later
	for _, room := range []int{0x03, 0x04, 0x05, 0x07, 0x06, 0x08, 0x0C, 0x17, 0x16, 0x18, 0x0F,
		0x21, 0x15, 0x12, 0x11, 0x23, 0x10, 0x22, 0x13, 0x14, 0x0E, 0x2B, 0x2C, 0x1E} {
		g.goToRoom(room)
		o.Run(30)
		rooms[room] = pictures.Screen(o)
	}

	// The chip and the token sensor, in the maze
	g.pickRobot(roWhite)
	g.goTo(roTile{0x16, 10, 9})
	g.press(" ")
	g.record(pictures, 4)
	g.pick(roChip)
	g.saveRecording(pictures, "chip")
	g.chase(roWhite)
	g.dropAt(45, 100)
	g.leave()
	g.pick(roTokenSensor)
	shot("token")
	g.chase(roWhite)
	g.dropAt(100, 100)
	g.leave()

	// The orange and blue robots out in the room before the guarded ones,
	// and the white one two rooms away, where its crystal sensor does not
	// see the crystal
	g.pickRobot(roWhite)
	g.goTo(roTile{0x11, 9, 2})
	g.press(" ")
	for _, robot := range []struct{ obj, col int }{{roOrange, 7}, {roBlue, 12}} {
		g.chase(roWhite)
		g.pickRobot(robot.obj)
		g.leave()
		g.goTo(roTile{0x11, robot.col, 6})
		g.press(" ")
	}
	g.pickRobot(roWhite)
	g.goTo(roTile{0x2B, 4, 7})
	g.press(" ")

	// The orange robot follows the walls: from the left wall of the room of
	// the guard, it goes round, takes what is there and goes up to the room
	// above, where the remote control stops it
	g.goToRoom(0x11)
	g.pickRobot(roOrange)
	g.goToRoom(0x10)
	g.goToRoom(0x22)
	g.goTo(roTile{0x22, 9, 3})
	g.putRobot(roOrange, 28, 140)
	shot("orange")
	g.chase(roOrange)
	g.switchRobot()
	g.leave()
	item := g.fetched(0x22)
	orange := pictures.Record(o).Faster(2)
	orange.Capture(10)
	g.waitFor("the orange robot taking the object", 120, orange, func() bool { return g.peek(0x600+item) == roOrange })
	g.waitFor("the orange robot leaving", 120, orange, func() bool { return g.room(roOrange) == 0x10 })
	must(t, pictures.SaveRecording(orange, "orange", 200))
	g.record(pictures, 3)
	g.offAndHome(roOrange, 35, 130)
	g.saveRecording(pictures, "home")

	// The blue robot goes right and back: lined up with what is in the room
	// of the other guard, it takes it on the way
	g.goToRoom(0x11)
	g.pickRobot(roBlue)
	g.goToRoom(0x23)
	g.putRobot(roBlue, 14, 92)
	shot("blue")
	g.chase(roBlue)
	g.switchRobot()
	g.leave()
	item = g.fetched(0x23)
	blue := pictures.Record(o).Faster(2)
	blue.Capture(10)
	g.waitFor("the blue robot taking the object", 120, blue, func() bool { return g.peek(0x600+item) == roBlue })
	g.waitFor("the blue robot coming back", 120, blue, func() bool { return g.room(roBlue) == 0x11 })
	must(t, pictures.SaveRecording(blue, "blue", 200))
	g.offAndHome(roBlue, 85, 130)
	g.collect()

	// The directional sensor of the token
	g.pick(roDirection)
	shot("direction")
	g.chase(roWhite)
	g.dropAt(100, 60)
	shot("packed")
	g.leave()

	// The sewer grate, which only a robot gets through: the white one goes
	// up and down its lanes, and right whenever it can, with the player in
	// it, watching through its eye
	g.pickRobot(roWhite)
	g.goToRoom(0x1E)
	g.putRobot(roWhite, 16, 40)
	shot("grate")
	g.chase(roWhite)
	g.switchRobot()
	g.goTo(roTile{0x0A, 17, 2})
	ride := pictures.Record(o).Faster(4)
	ride.Capture(10)
	g.waitFor("the white robot through the grate", 300, ride, func() bool {
		// The room under the grate, for the map, through the eye
		if _, ok := rooms[0x20]; !ok && g.room(roWhite) == 0x20 && g.x(roWhite) > 30 {
			rooms[0x20] = pictures.Screen(o)
		}
		return g.room(roWhite) == 0x1F
	})
	ride.Run(2*60, 6)
	must(t, pictures.SaveRecording(ride, "ride", 200))
	g.switchRobot()

	// The last door, its lock past the grate, and the transporter
	g.pick(roKey)
	g.leaveBy("K")
	o.Run(30)
	rooms[0x1F] = pictures.Screen(o)
	g.record(pictures, 2)
	g.goTo(roTile{0x1F, 11, 6})
	g.regrip(1)
	g.unlock(86, 81, 0x1F, 5, 2)
	g.saveRecording(pictures, "lock")
	g.chase(roWhite)
	g.press(" ")
	g.leave()
	g.pickRobot(roWhite)
	g.goToRoom(0x2A)
	shot("transporter")
	rooms[0x2A] = pictures.Screen(o)
	must(t, roMap(pictures, rooms))
	g.record(pictures, 1)
	g.goTo(roTile{0x2A, 4, 9})
	g.stand(38, 24)
	g.run(3 * operator.FramesPerSecond)
	g.saveRecording(pictures, "transport")
	must(t, o.Key("Escape"))
	o.RunSeconds(30)
	shot("subway")
	for _, obj := range []int{roMagnet, roCrystal, roChip, roTokenSensor, roDirection, roKey, roOrange, roBlue} {
		if roRobotOf[g.room(obj)] == 0 {
			t.Errorf("%02X is in room %02X, not in a robot, in the Subway", obj, g.room(obj))
		}
	}
}

/*
offAndHome stops a robot that is going round with what it fetched, with the
remote control, takes it to the middle of the room before the guards, lets
the robots go on and switches it off from inside, and takes it into the white
robot, at x,y in there. A stopped robot can be carried but not gone into.
*/
func (g *roGame) offAndHome(robot, x, y int) {
	g.remote()
	g.pickRobot(robot)
	g.goToRoom(0x11)
	g.goTo(roTile{0x11, 8, 5})
	g.press(" ")
	g.remote()
	g.chase(robot)
	g.switchRobot()
	g.leave()
	g.pickRobot(robot)
	g.chase(roWhite)
	g.putRobot(robot, x, y)
	g.leave()
}

// roMapPlaces are the rooms of the Sewer on its map, by column and row, as
// the map of ASchultz's walkthrough has them: the exits of the maze do not
// all lead where a map on paper would
var roMapPlaces = map[int][2]int{
	0x28: {0, 0}, 0x02: {1, 0}, 0x03: {2, 0}, 0x0E: {3, 0}, 0x14: {4, 0}, 0x13: {5, 0}, 0x2B: {6, 0}, 0x1E: {7, 0}, 0x1F: {8, 0}, 0x2A: {9, 0},
	0x04: {0, 1}, 0x06: {1, 1}, 0x21: {3, 1}, 0x15: {4, 1}, 0x12: {5, 1}, 0x2C: {6, 1}, 0x20: {7, 1},
	0x05: {0, 2}, 0x07: {1, 2}, 0x0F: {3, 2}, 0x10: {4, 2}, 0x11: {5, 2}, 0x23: {6, 2},
	0x08: {1, 3}, 0x17: {2, 3}, 0x22: {4, 3},
	0x0C: {1, 4}, 0x16: {2, 4},
	0x18: {2, 5},
}

// roMap puts the pictures of the rooms together, each a third of its size,
// in their places on the map of the Sewer, and writes it as map.png
func roMap(pictures *album.Album, rooms map[int]image.Image) error {
	const w, h, gap = 560 / 3, 384 / 3, 4
	m := image.NewRGBA(image.Rect(0, 0, 10*(w+gap)+gap, 6*(h+gap)+gap))
	xdraw.Draw(m, m.Bounds(), image.NewUniform(color.RGBA{0x40, 0x40, 0x40, 0xFF}), image.Point{}, xdraw.Src)
	for room, place := range roMapPlaces {
		screen, ok := rooms[room]
		if !ok {
			continue
		}
		x, y := gap+place[0]*(w+gap), gap+place[1]*(h+gap)
		xdraw.ApproxBiLinear.Scale(m, image.Rect(x, y, x+w, y+h), screen, screen.Bounds(), xdraw.Src, nil)
	}
	f, err := os.Create(pictures.Path("map.png"))
	if err != nil {
		return err
	}
	if err := png.Encode(f, m); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}
