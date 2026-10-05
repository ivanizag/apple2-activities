package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

/*
dos33Screenshots is life with DOS 3.3 on an Apple ][+ with two drives: the
System Master started, its catalog, its greeting program listed, a blank
diskette initialized with a program of our own, and the machine started from
it.

	cp disks/blank.dsk my-disk.dsk
	izapple2 -model 2plus disks/dos33-master.dsk my-disk.dsk
	izapple2 -model 2plus my-disk.dsk
*/
func dos33Screenshots(t *testing.T) {
	pictures := newAlbum("dos33", album.Green)
	myDisk := disk(t, "blank.dsk")
	o := start(t, "2plus", nil, disk(t, "dos33-master.dsk"), myDisk)

	// Started, recorded at ten times the speed until the prompt
	booting := pictures.Record(o).Faster(10)
	booting.Capture(20)
	for !o.HasText("COPYRIGHT") {
		if o.Frames() > 200*60 {
			t.Fatal("the System Master did not start")
		}
		booting.Run(60, 30)
	}
	booting.Run(300, 30)
	must(t, pictures.SaveRecording(booting, "boot", 300))

	// The catalog, a screen at a time
	must(t, o.TypeLines("HOME", "CATALOG"))
	o.Run(240)
	must(t, pictures.Screenshot(o, "catalog"))
	must(t, o.Key("Space"))
	o.Run(120)

	// The greeting program, loaded and listed
	must(t, o.TypeLines("HOME", "LOAD HELLO", "HOME", "LIST 10,100"))
	o.Run(60)
	must(t, pictures.Screenshot(o, "hello"))

	// A program of our own, and a blank diskette in drive 2 made a disk of
	// DOS 3.3 that starts it
	must(t, o.TypeLines(
		"NEW",
		"HOME",
		`10 PRINT "THIS IS ADA'S DISK"`,
		`20 PRINT "MADE ON AN APPLE ][+"`,
		"INIT HELLO,D2",
	))
	must(t, o.WaitForKeyboard(120))
	o.Run(60)
	initializing := pictures.Record(o)
	initializing.Capture(50)
	must(t, initializing.TypeLines("CATALOG,D2"))
	initializing.Run(240, 30)
	must(t, pictures.SaveRecording(initializing, "init", 300))

	// The machine started again with our disk alone in drive 1
	again := start(t, "2plus", nil, myDisk)
	must(t, again.WaitForText("ADA'S DISK", 120))
	again.Run(60)
	must(t, pictures.Screenshot(again, "own-disk"))
}
