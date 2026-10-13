package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// The rooms of the Streets on their map: the crossroads of the start, the
// Sensor Editor at its left, the maze of the storms in a column, with the
// coin at its bottom, the elevator up to the vending machine at its right,
// and the end, three rows of five rooms, at the bottom right
var roStreetsPlaces = map[int][2]int{
	0x12: {5, 0}, 0x11: {5, 1}, 0x10: {5, 2}, 0x0F: {5, 3}, 0x0E: {4, 3},
	0x02: {1, 3}, 0x03: {1, 4}, 0x04: {1, 5}, 0x05: {0, 4}, 0x31: {2, 4},
	0x06: {3, 3}, 0x07: {3, 4}, 0x19: {3, 5}, 0x1C: {4, 5}, 0x1A: {3, 6}, 0x08: {3, 7}, 0x1B: {3, 8}, 0x3C: {3, 9},
	0x28: {4, 7},
	0x2B: {5, 6}, 0x1E: {6, 6}, 0x1F: {7, 6}, 0x17: {8, 6}, 0x20: {9, 6},
	0x2C: {5, 7}, 0x21: {6, 7}, 0x22: {7, 7}, 0x16: {8, 7}, 0x2A: {9, 7},
	0x2D: {5, 8}, 0x24: {6, 8}, 0x25: {7, 8}, 0x18: {8, 8}, 0x26: {9, 8},
}

// The objects of the Streets
const (
	roShield        = 0x17 // the Magnetic Storm Shield, its input
	roElevatorKey   = 0xF8
	roCrystal3      = 0xF9 // the energy crystal of the Streets
	roForm          = 0x07 // the Form-12 of the vending machine
	roButtonSensor  = 0x0D // the sensor of the directions of the button
	roElevatorFloor = 0x8F0C
	roSentryDone    = 0x8F01
	roSentryTop     = 0x8F03
	roRoomDoor      = 0x8F19 // 1 while the door between the practice room and the final room is shut
)

/*
robotOdysseyStreetsScreenshots plays the third level of Robot Odyssey, the
Streets, after the Sewer and the Subway, played first without pictures: the
coin of the maze of the storms fetched by the blue robot with the storm
shield, the elevator key from the junk, the crystal that brings the white
robot back to life, the Form-12 out of the vending machine with the two
robots, the sentry and its two buttons, the button of the final room found by
the white robot with the sensor of its direction, and the form handed to the
gatekeeper, whose transporter goes to the fourth level.
*/
func robotOdysseyStreetsScreenshots(t *testing.T) {
	o := start(t, robotOdysseyMachine, nil)
	g := roSubway(t, o, album.New(t.TempDir(), album.Color))
	pictures := newAlbum("robotodyssey-streets", album.Color)
	g.mapAlbum = pictures
	g.rooms = nil
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
	keys := func(key string, n int) {
		for range n {
			g.press(key)
		}
	}
	mem := func(addr int) int { return int(o.Apple2().Peek(uint16(addr))) }
	waitFor := func(what string, rec *album.Recording, done func() bool) {
		g.waitFor(what, 120, rec, done)
	}
	// The elevator: on the knob at its left and off again at once, it
	// goes to the other floor, and out by its left
	elevator := func(floor int, rec *album.Recording) {
		g.stand(100, 104)
		keys("K", 30)
		g.stand(14, 80)
		keys("J", 3)
		keys("K", 2)
		waitFor("the elevator", rec, func() bool { return mem(roElevatorFloor) == floor })
		g.stand(20, 50)
		keys("J", 8)
	}

	// The items of the level are put in their rooms, in an order that turns
	// with the time, when the player first leaves the start: this time, the
	// crystal goes up the elevator
	o.Run(11)

	// The storm shield of the Sensor Editor into the blue robot, on its
	// spare NOT gate, and the sensor of the directions of the coin to the
	// thrusters, the grabber always on
	g.goToRoom(0x05)
	g.putRobot(roBlue, 70, 100)
	g.pick(roShield)
	g.chase(roBlue)
	g.dropAt(65, 150)
	g.pen()
	unwire([2]int{66, 97}, [2]int{73, 98}, roPinThrusterRight)
	wire([2]int{61, 158}, roPinNot2Out,
		roPinThrusterRight, roPinTokenLeft,
		[2]int{89, 16}, [2]int{101, 120},
		roPinThrusterUp, [2]int{99, 82},
		roPinThrusterLeft, [2]int{108, 98},
		roPinGrabberIn, roPinNotOut)
	g.cursor()
	shot("shield")
	g.leave()
	g.remote()

	// The coin flees from the player: the robot, put down in its room,
	// takes it while the player waits a room away
	g.pickRobot(roBlue)
	g.goToRoom(0x3C)
	g.putRobot(roBlue, 40, 140)
	g.goToRoom(0x1B)
	waitFor("the coin taken", nil, func() bool { return g.peek(0x600+roToken) == roBlue })
	g.goToRoom(0x3C)
	shot("coin")
	g.pickRobot(roBlue)

	// The elevator key from the junk, into its slot by the door
	g.goToRoom(0x0F)
	g.putRobot(roBlue, 40, 104)
	g.remote()
	shot("junk")
	g.pick(roElevatorKey)
	g.moveCarried(60, 90)
	g.regripAt(6, 0)
	g.moveCarried(89, 80)
	g.press(" ")
	shot("door")
	g.pickRobot(roBlue)
	g.record(pictures, 1)
	elevator(0x11, g.rec)
	g.saveRecording(pictures, "elevator")

	// The crystal, up the elevator, into the white robot, dead since the
	// Subway, and the orange one out of it
	g.putRobot(roBlue, 40, 140)
	shot("crystal")
	g.pick(roCrystal3)
	g.chase(roBlue)
	g.chase(roWhite)
	g.recharge()
	g.press(" ")
	shot("charged")
	g.pickRobot(roOrange)
	g.leave()
	g.leave()
	g.putRobot(roOrange, 84, 24)

	// The white robot to go up and down, grabbing what it meets
	g.chase(roBlue)
	g.chase(roWhite)
	g.pen()
	unwire(roPinThrusterUp, [2]int{89, 17}, [2]int{16, 64}, [2]int{64, 95}, [2]int{71, 96})
	wire(roPinThrusterUp, [2]int{64, 95},
		[2]int{89, 17}, [2]int{71, 96},
		roPinGrabberIn, [2]int{28, 70})
	g.cursor()
	shot("white-form")
	g.leave()
	g.pickRobot(roWhite)
	g.leave()
	g.putRobot(roWhite, 86, 166)

	// The key out of the blue robot, for its key sensor, and the blue
	// robot to go up from the floor, and down and let go when the key is in
	// the room
	g.chase(roBlue)
	g.pick(roKey)
	g.leave()
	g.dropAt(92, 104)
	g.chase(roBlue)
	g.pen()
	unwire(roPinGrabberIn, roPinThrusterRight, [2]int{89, 16}, roPinThrusterUp, roPinThrusterLeft)
	wire(roPinThrusterUp, roPinFlipOut27,
		[2]int{89, 16}, roPinFlipOut28)
	unwire(roPinFlipIn27)
	wire([2]int{100, 41}, [2]int{52, 115},
		roPinNodeOut, roPinFlipIn27,
		roPinNodeOut2, [2]int{35, 117})
	g.cursor()
	shot("blue-coin")
	g.leave()

	// The coin let go and held again on top of the robot, to push it into
	// the slot under the shelf
	g.remote()
	g.remote()
	g.chase(roBlue)
	g.pen()
	wire(roPinGrabberIn, roPinNotOut)
	g.cursor()
	g.leave()
	g.pick(roToken)
	g.goToRoom(0x12)
	g.stand(40, 28)
	g.dropAt(30, 40)
	g.goToRoom(0x11)
	g.pickRobot(roBlue)
	g.goToRoom(0x12)
	g.stand(50, 28)
	g.putRobot(roBlue, 24, 16)
	shot("vending")
	coin := pictures.Record(o)
	coin.Capture(10)
	g.remote()
	waitFor("the coin taken", coin, func() bool { return g.peek(0x600+roToken) == roBlue })
	waitFor("the coin under the shelf", coin, func() bool { return g.y(roToken) > 100 })
	coin.Run(operator.FramesPerSecond, 6)
	must(t, pictures.SaveRecording(coin, "coin-up", 200))
	g.goToRoom(0x11)
	g.pick(roKey)
	g.goToRoom(0x12)
	g.stand(70, 28)
	slot := pictures.Record(o)
	slot.Capture(10)
	g.press(" ")
	waitFor("the coin in the slot", slot, func() bool { return g.room(roToken) == 0x3F })
	waitFor("the blue robot down", slot, func() bool { return g.y(roBlue) < 30 })
	g.remote()
	must(t, pictures.SaveRecording(slot, "slot", 200))

	// The key into the orange robot, out of the way; the blue robot to pull
	// the knob, its grabber's output to the left thruster, and the white
	// one to fetch the form, through the gap the shelf leaves
	g.pick(roKey)
	g.regripAt(0, -8)
	g.goToRoom(0x11)
	g.stand(70, 150)
	g.chase(roOrange)
	g.dropAt(60, 100)
	g.leave()
	g.goToRoom(0x12)
	g.chase(roBlue)
	g.pen()
	wire([2]int{31, 143}, roPinThrusterLeft)
	g.cursor()
	shot("blue-knob")
	g.leave()
	g.pickRobot(roBlue)
	g.putRobot(roBlue, 82, 20)
	g.goToRoom(0x11)
	g.chase(roWhite)
	g.setSwitch(true)
	g.leave()
	g.pickRobot(roWhite)
	g.goToRoom(0x12)
	g.stand(40, 28)
	g.putRobot(roWhite, 14, 16)
	pull := pictures.Record(o)
	pull.Capture(10)
	g.remote()
	waitFor("the form taken", pull, func() bool { return g.peek(0x600+roForm) == roWhite })
	waitFor("the white robot down", pull, func() bool { return g.y(roWhite) < 30 })
	g.remote()
	must(t, pictures.SaveRecording(pull, "form", 200))

	// Both robots back: the white one first, switched off; then the key in
	// the room, which brings the blue one down
	g.pickRobot(roWhite)
	g.goToRoom(0x11)
	g.putRobot(roWhite, 86, 166)
	g.chase(roWhite)
	g.setSwitch(false)
	g.leave()
	g.remote()
	g.chase(roOrange)
	g.pick(roKey)
	g.leave()
	g.goToRoom(0x12)
	g.stand(60, 28)
	g.press(" ")
	waitFor("the blue robot down", nil, func() bool { return g.y(roBlue) < 30 })
	g.remote()
	g.pick(roKey)
	g.regripAt(0, -8)
	g.goToRoom(0x11)
	g.stand(70, 150)
	g.chase(roOrange)
	g.dropAt(60, 100)
	g.leave()
	g.goToRoom(0x12)
	g.pickRobot(roBlue)
	g.goToRoom(0x11)
	g.putRobot(roBlue, 66, 130)

	// All of them in the blue robot, its thrusters off, running for its
	// shield, down the elevator and through the storms to the end
	g.pickRobot(roWhite)
	g.chase(roBlue)
	g.putRobot(roWhite, 90, 140)
	g.leave()
	g.pickRobot(roOrange)
	g.chase(roBlue)
	g.putRobot(roOrange, 60, 24)
	g.pen()
	unwire(roPinThrusterUp, [2]int{89, 16}, roPinThrusterLeft)
	g.cursor()
	g.leave()
	g.stand(36, 100)
	g.stand(80, 100)
	g.remote()
	g.pickRobot(roBlue)
	elevator(0x0F, nil)
	for _, room := range []int{0x0E, 0x06, 0x07, 0x19, 0x1A, 0x08, 0x28, 0x2C} {
		g.goToRoom(room)
	}
	g.putRobot(roBlue, 70, 100)
	g.stand(30, 100)
	g.remote()
	shot("sentry")

	// The two buttons of the sentry: the blue robot to go right, up at the
	// wall and left at the ceiling, onto the button at the top, and the
	// white one to go right, down at the wall and left at the floor, onto
	// the one at the bottom
	g.chase(roBlue)
	g.pickRobot(roOrange)
	g.leave()
	g.putRobot(roOrange, 100, 30)
	g.chase(roBlue)
	g.pickRobot(roWhite)
	g.leave()
	g.putRobot(roWhite, 56, 150)
	g.chase(roBlue)
	g.pen()
	unwire(roPinFlipIn27)
	wire([2]int{118, 56}, roPinFlipIn27,
		[2]int{89, 16}, roPinFlipOut27,
		roPinThrusterLeft, roPinFlipOut28,
		roPinThrusterRight, [2]int{100, 159})
	g.cursor()
	shot("blue-sentry")
	g.leave()
	g.chase(roWhite)
	g.pen()
	unwire([2]int{64, 83})
	wire([2]int{118, 56}, [2]int{64, 83})
	unwire([2]int{89, 17})
	wire([2]int{16, 64}, [2]int{71, 96})
	unwire([2]int{71, 84})
	wire(roPinThrusterRight, [2]int{33, 20})
	g.cursor()
	shot("white-sentry")
	g.leave()
	g.pickRobot(roBlue)
	g.stand(120, 80)
	keys("K", 3)
	g.press(" ")
	// Watched from the doorway of the sentry's room
	buttons := pictures.Record(o).Faster(2)
	buttons.Capture(10)
	g.remote()
	waitFor("the top button", buttons, func() bool { return mem(roSentryTop) == 1 })
	g.remote()
	g.chase(roWhite)
	g.setSwitch(true)
	g.leave()
	g.pickRobot(roWhite)
	g.stand(60, 100)
	g.regripAt(6, 0)
	g.stand(120, 66)
	keys("K", 3)
	g.press(" ")
	buttons.Capture(10)
	g.remote()
	waitFor("the sentry gone", buttons, func() bool { return mem(roSentryDone) == 1 })
	g.remote()
	must(t, pictures.SaveRecording(buttons, "buttons", 200))

	// The sensor of the direction of the button, in the gatekeeper's room,
	// and the node and the other NOT gate of the blue robot, into the white
	// one; the form out of it, into the practice room
	g.pick(roButtonSensor)
	g.goToRoom(0x2C)
	g.goToRoom(0x21)
	g.dropAt(30, 110)
	g.pickRobot(roWhite)
	g.putRobot(roWhite, 60, 80)
	g.pick(roButtonSensor)
	g.chase(roWhite)
	g.dropAt(50, 40)
	g.leave()
	g.pickRobot(roBlue)
	g.putRobot(roBlue, 100, 100)
	g.chase(roBlue)
	g.pick(roNode)
	g.leave()
	g.chase(roWhite)
	g.dropAt(100, 60)
	g.leave()
	g.chase(roBlue)
	g.pick(roNot2)
	g.leave()
	g.chase(roWhite)
	g.dropAt(76, 120)
	g.pen()
	unwire(roPinGrabberIn)
	g.cursor()
	g.leave()
	g.stand(30, 40)
	g.remote()
	g.remote()
	g.pick(roForm)
	g.goToRoom(0x22)
	g.dropAt(20, 150)
	g.goToRoom(0x21)

	// The white robot to go right while it is above the button or right of
	// it, up at a wall at its right, and down when the button is lower and
	// no wall at its right; through the door of the practice room, which
	// keeps the player out, it finds the button of the final room
	g.chase(roWhite)
	g.pen()
	unwire(roPinThrusterUp, [2]int{16, 64}, roPinThrusterRight, [2]int{64, 83})
	wire([2]int{100, 61}, [2]int{118, 56},
		[2]int{89, 17}, [2]int{100, 67},
		[2]int{100, 55}, [2]int{77, 110},
		[2]int{87, 49}, [2]int{28, 70},
		[2]int{92, 49}, [2]int{77, 129},
		roPinThrusterUp, [2]int{90, 33},
		[2]int{28, 89}, [2]int{52, 57},
		[2]int{58, 138}, [2]int{42, 45},
		[2]int{63, 141}, [2]int{50, 29},
		[2]int{71, 84}, [2]int{61, 157},
		[2]int{16, 64}, [2]int{64, 95},
		roPinThrusterRight, [2]int{71, 96})
	g.cursor()
	g.stand(20, 100)
	shot("white-button")
	g.leave()
	g.pickRobot(roWhite)
	g.goToRoom(0x22)
	g.stand(8, 165)
	g.stand(112, 165)
	g.stand(112, 80)
	g.press(" ")
	g.stand(112, 150)
	g.stand(8, 150)
	g.stand(8, 20)
	g.stand(60, 20)
	g.remote()
	waitFor("the door open", nil, func() bool { return mem(roRoomDoor) == 0 })
	// The eye of the practice room shows the final room
	g.stand(122, 20)
	shot("final")
	g.remote()

	// The robots, the others in the white one, into the gatekeeper's room;
	// then the form to the gatekeeper, who throws it away, and the
	// transporter appears
	g.pickRobot(roWhite)
	g.goToRoom(0x21)
	g.putRobot(roWhite, 40, 96)
	g.pickRobot(roBlue)
	g.chase(roWhite)
	g.putRobot(roBlue, 70, 30)
	g.leave()
	g.goToRoom(0x2C)
	g.pickRobot(roOrange)
	g.goToRoom(0x21)
	g.chase(roWhite)
	g.putRobot(roOrange, 100, 150)
	g.leave()
	g.pickRobot(roWhite)
	for _, room := range []int{0x22, 0x16, 0x2A} {
		g.goToRoom(room)
	}
	g.putRobot(roWhite, 60, 130)
	g.goToRoom(0x22)
	g.pick(roForm)
	g.goToRoom(0x16)
	g.goToRoom(0x2A)
	g.stand(20, 100)
	shot("gatekeeper")
	gate := pictures.Record(o)
	gate.Capture(10)
	// Free, against the hand of the gatekeeper
	g.dropAt(41, 108)
	waitFor("the form thrown away", gate, func() bool { return g.room(roForm) == 0x3F })
	gate.Run(2*operator.FramesPerSecond, 6)
	must(t, pictures.SaveRecording(gate, "gatekeeper", 200))
	g.pickRobot(roWhite)
	g.stand(30, 80)
	shot("transporter")
	g.stand(35, 80)
	must(t, roMap(pictures, g.rooms, roStreetsPlaces, 10, 10))
	o.RunSeconds(10)
	shot("disk")
	must(t, o.Key("Escape"))
	o.RunSeconds(30)
	shot("master")
	for _, obj := range []int{roOrange, roBlue} {
		if g.room(obj) != roInside[roWhite] {
			t.Errorf("%02X is in room %02X, not in the white robot, in the Master Computer Center", obj, g.room(obj))
		}
	}
}
