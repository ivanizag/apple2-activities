package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
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

	// The title, the menu, and the game, which starts with the player
	// falling into the Sewer
	o.RunSeconds(20)
	shot("menu")
	must(t, o.Type("\n"))
	o.RunSeconds(40)
	shot("ready")
	must(t, o.Type("\n"))
	o.RunSeconds(90)
	shot("start")

	// The three robots, switched off, so that they stay where they are
	g.goTo(roTile{0x02, 0, 9})
	shot("robots")
	for i, robot := range []int{roOrange, roBlue, roWhite} {
		g.chase(robot)
		if i == 0 {
			shot("inside")
		}
		g.switchRobot()
		g.leave()
	}

	// The blue key, in one of them, and the door of the City Sewer
	g.chase(roRobotOf[g.room(roKey)])
	g.pick(roKey)
	g.leave()
	shot("key")
	g.press(" ")
	for _, robot := range []struct{ obj, x, y int }{{roOrange, 35, 130}, {roBlue, 85, 130}} {
		g.pickRobot(robot.obj)
		g.chase(roWhite)
		g.putRobot(robot.obj, robot.x, robot.y)
		g.leave()
	}
	g.pick(roKey)
	g.goTo(roTile{0x03, 2, 10})
	g.unlock(17, 16, 0x03, 4, 9)
	shot("door")
	g.chase(roWhite)
	g.press(" ")
	shot("white")
	g.leave()

	// The chip and the token sensor, in the maze
	g.pickRobot(roWhite)
	g.goTo(roTile{0x16, 10, 9})
	g.press(" ")
	g.pick(roChip)
	shot("chip")
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
	g.offAndHome(roOrange, 35, 130)

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
	g.waitFor("the white robot through the grate", 300, ride, func() bool { return g.room(roWhite) == 0x1F })
	ride.Run(2*60, 6)
	must(t, pictures.SaveRecording(ride, "ride", 200))
	g.switchRobot()

	// The last door, its lock past the grate, and the transporter
	g.pick(roKey)
	g.leaveBy("K")
	g.goTo(roTile{0x1F, 11, 6})
	g.regrip(1)
	g.unlock(86, 81, 0x1F, 5, 2)
	shot("open")
	g.chase(roWhite)
	g.press(" ")
	g.leave()
	g.pickRobot(roWhite)
	g.goToRoom(0x2A)
	shot("transporter")
	g.goTo(roTile{0x2A, 4, 9})
	g.stand(38, 24)
	o.RunSeconds(3)
	shot("disk")
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
