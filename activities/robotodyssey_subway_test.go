package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// The rooms of the Subway on its map: a grid that goes round, its bottom
// row the line of the train, and the room of the exit, up from Jack in the
// 'Bot, under it
var roSubwayPlaces = map[int][2]int{
	0x38: {0, 0}, 0x34: {1, 0}, 0x30: {2, 0}, 0x2C: {3, 0}, 0x28: {4, 0}, 0x24: {5, 0}, 0x20: {6, 0}, 0x3C: {7, 0},
	0x37: {0, 1}, 0x33: {1, 1}, 0x2F: {2, 1}, 0x2B: {3, 1}, 0x27: {4, 1}, 0x23: {5, 1}, 0x1F: {6, 1}, 0x3B: {7, 1},
	0x36: {0, 2}, 0x32: {1, 2}, 0x2E: {2, 2}, 0x2A: {3, 2}, 0x26: {4, 2}, 0x22: {5, 2}, 0x1E: {6, 2}, 0x3A: {7, 2},
	0x39: {0, 3}, 0x35: {1, 3}, 0x31: {2, 3}, 0x2D: {3, 3}, 0x29: {4, 3}, 0x25: {5, 3}, 0x21: {6, 3}, 0x3D: {7, 3},
	0x02: {4, 4},
}

// The pins of the inside of a robot, where the Solderpen glows for them
var (
	roPinThrusterUp    = [2]int{44, 160} // the input of the thruster at the top, which pushes down
	roPinThrusterRight = [2]int{119, 117}
	roPinThrusterDown  = [2]int{89, 16}
	roPinThrusterLeft  = [2]int{16, 63}
	roPinBumperUp      = [2]int{99, 160} // the output of the bumper at the top
	roPinBumperDown    = [2]int{32, 18}
	roPinGrabberIn     = [2]int{36, 144}
	roPinGrabberOut    = [2]int{30, 144}
)

// The pins of the blue robot: its flip-flop, two halves 27 and 28, each an
// input under it and an output over it, its NOT gate, and the NOT gate and
// the node brought from the others, where the generator puts them
var (
	roPinFlipIn27   = [2]int{66, 85}
	roPinFlipOut27  = [2]int{66, 97}
	roPinFlipIn28   = [2]int{73, 86}
	roPinFlipOut28  = [2]int{73, 98}
	roPinNotIn      = [2]int{36, 115}
	roPinNotOut     = [2]int{35, 133}
	roPinNot2In     = [2]int{78, 31}
	roPinNot2Out    = [2]int{78, 48}
	roPinNodeIn     = [2]int{100, 41}
	roPinNodeOut    = [2]int{100, 47}
	roPinNodeOut2   = [2]int{100, 35}
	roPinTokenLeft  = [2]int{93, 104} // the output of the sensor of directions for left, where it is put
	roNode          = 0xD1
	roNot2          = 0x76
	roToken         = 0x04
	roExitTicket    = 0xF7
	roDepotHandle   = 0x24
	roSubwayDepot   = 0x3D
	roSubwayLaundry = 0x21
	roSubwayJack    = 0x29
)

/*
robotOdysseySubwayScreenshots plays the second level of Robot Odyssey, the
Subway, after the Sewer, played first without pictures: the blue robot rewired
three times, for the token of the subway, for the exit ticket and for the
token again, the turnstile, the train called at the Depot, and the exit from
Jack in the 'Bot to the Streets, with the three robots and all that is in
them.
*/
func robotOdysseySubwayScreenshots(t *testing.T) {
	o := start(t, robotOdysseyMachine, nil)
	g := roSewer(t, o, album.New(t.TempDir(), album.Color))
	pictures := newAlbum("robotodyssey-subway", album.Color)
	g.mapAlbum = pictures
	shot := func(name string) {
		o.Run(30)
		must(t, pictures.Screenshot(o, name))
	}
	wire := func(pins ...[2]int) {
		for i := 0; i < len(pins); i += 2 {
			g.wire(pins[i], pins[i+1])
		}
	}
	unwire := func(pins ...[2]int) {
		for _, p := range pins {
			g.solder(p[0], p[1])
		}
	}
	waitFor := func(what string, rec *album.Recording, done func() bool) {
		g.waitFor(what, 120, rec, done)
	}

	// Into the Subway with the white robot, the others in it. The blue one
	// out, and into it a node of the white one and the NOT gate of the
	// orange one, and the sensor of the directions of the token
	shot("arrive")
	g.press(" ")
	g.chase(roWhite)
	g.pickRobot(roBlue)
	g.leave()
	g.press(" ")
	g.chase(roWhite)
	g.pick(roNode)
	g.leave()
	g.chase(roBlue)
	g.dropAt(100, 40)
	g.leave()
	g.chase(roWhite)
	g.chase(roOrange)
	g.pick(roNot2)
	g.leave()
	g.leave()
	g.chase(roBlue)
	g.dropAt(78, 40)
	g.leave()
	g.chase(roWhite)
	g.pick(roDirection)
	g.leave()
	g.chase(roBlue)
	g.dropAt(100, 100)

	// The battery of the blue robot, dead, charged with the energy crystal
	shot("dead")
	g.pick(roCrystal)
	g.recharge()
	g.dropAt(58, 48)
	shot("parts")

	// Wired to go up and down, flip-flop and bumpers; to the left when the
	// sensor of directions finds the token there; and back to the right
	// once the grabber holds it
	g.record(pictures, 6)
	g.pen()
	unwire(roPinThrusterLeft, roPinThrusterRight, roPinFlipIn27, roPinFlipIn28, roPinNotIn)
	wire(roPinBumperDown, roPinFlipIn28, roPinFlipOut28, roPinThrusterDown,
		roPinBumperUp, roPinFlipIn27, roPinFlipOut27, roPinThrusterUp,
		roPinTokenLeft, roPinThrusterRight, roPinGrabberOut, roPinThrusterLeft)
	g.cursor()
	g.saveRecording(pictures, "wire-token")
	shot("token-wired")
	g.leave()

	// The key put down out of its way, and the robot set going in the
	// corridor of the room of the token
	g.chase(roWhite)
	g.pick(roKey)
	g.leave()
	g.goTo(roTile{0x2B, 1, 10})
	g.press(" ")
	g.pickRobot(roBlue)
	g.goTo(roTile{0x2B, 15, 1})
	g.goToRoom(0x2C)
	g.goTo(roTile{0x2C, 15, 10})
	g.putRobot(roBlue, 110, 24)
	g.chase(roBlue)
	g.setSwitch(true)
	g.leave()
	g.goTo(roTile{0x2B, 18, 3})
	g.pick(roKey)
	g.regrip(8)
	g.goTo(roTile{0x2B, 18, 3})

	// The key in its lock while the robot is above it: the wall of the
	// token opens, the robot goes in, takes it, and comes out; back in the
	// corridor, it takes the player in
	waitFor("the blue robot up the corridor", nil, func() bool { return g.room(roBlue) == 0x2C && g.y(roBlue) > 100 })
	g.goTo(roTile{0x2C, 16, 10})
	g.moveCarried(114, 32)
	token := pictures.Record(o).Faster(2)
	token.Capture(10)
	waitFor("the token taken", token, func() bool { return g.peek(0x600+roToken) == roBlue })
	waitFor("the player taken in", token, func() bool { return g.room(0) == roInside[roBlue] })
	must(t, pictures.SaveRecording(token, "token", 200))

	// Inside: off, the grabber let go, and the token out
	g.setSwitch(false)
	g.pen()
	unwire(roPinGrabberIn)
	g.cursor()
	g.remote()
	g.remote()
	g.leave()
	g.cursor()
	shot("token-out")

	// The robot by the turnstile, and the token into its slot: it turns for
	// a few seconds, and the player goes through with the robot
	g.pickRobot(roBlue)
	g.goToRoom(0x38)
	g.goTo(roTile{0x38, 6, 9})
	g.putRobot(roBlue, 42, 30)
	g.pick(roToken)
	g.goToRoom(0x38)
	g.goTo(roTile{0x38, 10, 9})
	g.regripAt(-1, 2)
	g.record(pictures, 1)
	g.turnstile(roBlue)
	g.saveRecording(pictures, "turnstile")

	// The Depot: the robot, bouncing up and down at the left, goes up into
	// the room of the button when the handle is pulled, and the train comes
	depot := func(first bool) {
		g.goToRoom(0x3C)
		g.goToRoom(roSubwayDepot)
		g.goTo(roTile{roSubwayDepot, 8, 7})
		g.putRobot(roBlue, 60, 60)
		g.chase(roBlue)
		g.setSwitch(true)
		g.leave()
		if first {
			// Running: stopped to be carried
			g.remote()
		}
		g.pickRobot(roBlue)
		g.putRobot(roBlue, 8, 80)
		g.remote()
		g.goTo(roTile{roSubwayDepot, 13, 7})
		g.stand(92, 104)
		if first {
			shot("depot")
		}
		g.press(" ")
		var train *album.Recording
		if first {
			train = pictures.Record(o).Faster(2)
			train.Capture(10)
		}
		for range 5 {
			g.press("K")
		}
		for range 50 {
			if train != nil {
				train.Run(operator.FramesPerSecond/2, 6)
			} else {
				g.o.RunSeconds(0.5)
			}
		}
		if first {
			must(t, pictures.SaveRecording(train, "train", 200))
		}
		g.press(" ")
		waitFor("the blue robot down", nil, func() bool { return g.room(roBlue) == roSubwayDepot && g.y(roBlue) < 90 })
		g.remote()
		g.pickRobot(roBlue)
		g.goTo(roTile{roSubwayDepot, 4, 7})
		g.stand(30, 72)
		g.press(" ")
	}
	depot(true)

	// The train goes round the stations; off at the Laundrobot
	ride := pictures.Record(o).Faster(3)
	ride.Capture(10)
	waitFor("the Laundrobot", ride, func() bool { return g.room(0) == roSubwayLaundry })
	ride.Run(3*operator.FramesPerSecond, 6)
	must(t, pictures.SaveRecording(ride, "ride", 200))
	g.press("M")
	g.press(" ")

	// Wired to go right until it holds something, and then left, up and
	// down all the time: the grabber always on, and its output, by the node,
	// to the right thruster and by the other NOT gate to the left one
	g.chase(roBlue)
	g.setSwitch(false)
	g.record(pictures, 6)
	g.pen()
	unwire(roPinThrusterRight, roPinThrusterLeft)
	wire(roPinNotOut, roPinGrabberIn, roPinNodeIn, roPinGrabberOut, roPinNodeOut, roPinThrusterRight,
		roPinNodeOut2, roPinNot2In, roPinNot2Out, roPinThrusterLeft)
	g.cursor()
	g.saveRecording(pictures, "wire-ticket")
	shot("ticket-wired")
	g.setSwitch(true)
	g.leaveBy("J")

	// Into the room of the exit ticket from its left, where its guard does
	// not see the player, and the robot let go: under the wall, to the
	// ticket, and back
	g.pickRobot(roBlue)
	g.goToRoom(0x24)
	g.goTo(roTile{0x20, 2, 6})
	g.press(" ")
	ticket := pictures.Record(o).Faster(2)
	ticket.Capture(10)
	g.remote()
	waitFor("the exit ticket taken", ticket, func() bool { return g.peek(0x600+roExitTicket) == roBlue })
	waitFor("the blue robot back", ticket, func() bool { return g.room(roBlue) == 0x20 && g.x(roBlue) < 16 })
	must(t, pictures.SaveRecording(ticket, "ticket", 200))
	g.remote()
	g.pickRobotAbove(roBlue)

	// Back on foot to the turnstile, the ticket let go and put in the robot
	for _, room := range []int{0x24, 0x23, 0x27, 0x38} {
		g.goToRoom(room)
	}
	g.goTo(roTile{0x38, 12, 6})
	g.press(" ")
	g.chase(roBlue)
	g.setSwitch(false)
	g.pen()
	unwire(roPinGrabberIn)
	g.cursor()
	g.remote()
	g.remote()
	g.leave()
	g.pick(roExitTicket)
	g.chase(roBlue)
	g.dropAt(100, 70)

	// The token again, which the turnstile kept in a room under it: wired
	// to go up and down, and left whenever it can, the grabber always on
	g.pen()
	unwire(roPinThrusterLeft, roPinThrusterRight, roPinNot2In, roPinNodeIn)
	wire(roPinNotOut, roPinGrabberIn, roPinNot2Out, roPinThrusterRight)
	g.cursor()
	shot("token2-wired")
	g.setSwitch(true)
	g.leave()
	g.pickRobot(roBlue)
	g.goToRoom(0x33)
	g.goTo(roTile{0x33, 2, 6})
	g.press(" ")
	token2 := pictures.Record(o).Faster(3)
	token2.Capture(10)
	g.remote()
	waitFor("the token taken again", token2, func() bool { return g.peek(0x600+roToken) == roBlue })
	waitFor("the blue robot out with it", token2, func() bool { return g.room(roBlue) == 0x3B })
	token2.Run(3*operator.FramesPerSecond, 6)
	must(t, pictures.SaveRecording(token2, "token2", 200))
	g.remote()

	// Round the guards to the robot, and back with it to the white one
	for _, room := range []int{0x34, 0x38, 0x39, 0x36, 0x32, 0x2E, 0x2A, 0x2B, 0x27, 0x23, 0x1F, 0x3B} {
		g.goToRoom(room)
	}
	g.pickRobot(roBlue)
	for _, room := range []int{0x1F, 0x23, 0x27, 0x2B, 0x2A} {
		g.goToRoom(room)
	}
	g.goTo(roTile{0x2A, 8, 2})
	g.press(" ")
	g.chase(roBlue)
	g.setSwitch(false)
	g.pen()
	unwire(roPinGrabberIn)
	g.cursor()
	g.remote()
	g.remote()
	g.leave()

	// The white robot into the blue one, and all of them through the
	// turnstile, to the Depot and the train again
	g.pickRobot(roWhite)
	g.chase(roBlue)
	g.putRobot(roWhite, 90, 140)
	shot("all-in")
	g.leave()
	g.pickRobot(roBlue)
	g.goToRoom(0x38)
	g.goTo(roTile{0x38, 6, 9})
	g.putRobot(roBlue, 42, 30)
	g.pick(roToken)
	g.goToRoom(0x38)
	g.goTo(roTile{0x38, 10, 9})
	g.regripAt(-1, 2)
	g.turnstile(roBlue)
	depot(false)
	waitFor("Jack in the 'Bot", nil, func() bool { return g.room(0) == roSubwayJack })
	g.o.RunSeconds(3)
	g.press("I")
	shot("jack")

	// The exit ticket into its slot, up from the station: the steps of the
	// escalator come and go, and up them with the robot to the transporter
	g.press(" ")
	g.chase(roBlue)
	g.pick(roExitTicket)
	g.leave()
	g.goToRoom(0x02)
	g.goTo(roTile{0x02, 9, 10})
	g.regripAt(-6, 4)
	g.moveCarried(46, 20)
	g.press(" ")
	shot("slot")
	g.goToRoom(roSubwayJack)
	g.pickRobot(roBlue)
	g.goToRoom(0x02)
	g.goTo(roTile{0x02, 9, 10})
	g.record(pictures, 2)
	g.climb(20, 140)
	g.saveRecording(pictures, "escalator")
	shot("transporter")
	g.transport(118, 140, 134, 170)
	must(t, roMap(pictures, g.rooms, roSubwayPlaces, 8, 5))
	o.RunSeconds(3)
	shot("disk")
	must(t, o.Key("Escape"))
	o.RunSeconds(30)
	shot("streets")
	for _, obj := range []int{roMagnet, roCrystal, roChip, roTokenSensor, roDirection, roKey, roOrange, roWhite} {
		if roRobotOf[g.room(obj)] == 0 {
			t.Errorf("%02X is in room %02X, not in a robot, in the Streets", obj, g.room(obj))
		}
	}
}
