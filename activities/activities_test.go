/*
Package activities makes the pictures of the guides, by taking a machine
through what each guide tells its reader to do. It is not a test: it does
nothing unless asked for with A2_ACTIVITIES, and then writes the pictures of
each guide in guides/images, over the ones there.

	A2_ACTIVITIES=1 go test -count=1 -run 'TestActivities/switch-on' ./activities
*/
package activities

import (
	"os"
	"testing"
)

func TestActivities(t *testing.T) {
	if os.Getenv("A2_ACTIVITIES") == "" {
		t.Skip("this makes the pictures of the guides, with A2_ACTIVITIES=1")
	}
	t.Run("switch-on", switchOnScreenshots)
	t.Run("dos33", dos33Screenshots)
	t.Run("paddle-game", paddleScreenshots)
	t.Run("desktop", deskTopScreenshots)
	t.Run("apple-ii", appleIIScreenshots)
	t.Run("pascal", pascalScreenshots)
	t.Run("cpm", cpmScreenshots)
	t.Run("lode-runner", lodeRunnerScreenshots)
	t.Run("mockingboard", mockingboardScreenshots)
	t.Run("apple-iie", appleIIeScreenshots)
	t.Run("card-cat", cardCatScreenshots)
}
