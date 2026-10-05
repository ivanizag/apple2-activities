package activities

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

/*
What the activities run on: the disks fetch-disks.sh downloads into disks/,
which a reader gets the same of, copied for each machine so that what it
writes does not change them, and the folder the pictures of each guide go to.
*/

const (
	// guideImages is where the pictures of the guides go, a folder each
	guideImages = "../guides/images"

	// disks is where fetch-disks.sh leaves the disks
	disks = "../disks"
)

// newAlbum is the album of the pictures of a guide
func newAlbum(guide string, monitor album.Monitor) *album.Album {
	return album.New(filepath.Join(guideImages, guide), monitor)
}

// disk is a copy of a disk of disks/, in the test's own directory
func disk(t testing.TB, name string) string {
	t.Helper()

	source, err := os.Open(filepath.Join(disks, name))
	if err != nil {
		t.Fatalf("the disk %v is missing, run ./fetch-disks.sh: %v", name, err)
	}
	defer source.Close()

	copied := filepath.Join(t.TempDir(), name)
	target, err := os.Create(copied)
	if err != nil {
		t.Fatal(err)
	}
	defer target.Close()

	if _, err := io.Copy(target, source); err != nil {
		t.Fatal(err)
	}
	return copied
}

// start builds a machine of a model and sits an operator at it
func start(t testing.TB, model string, overrides map[string]string, files ...string) *operator.Operator {
	t.Helper()
	o, err := operator.Start(model, overrides, files...)
	if err != nil {
		t.Fatal(err)
	}
	return o
}

// must stops the activity when the operator could not do what it was asked
func must(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
