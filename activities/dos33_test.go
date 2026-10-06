package activities

import (
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

/*
dos33Machine is the machine of the guide, as its command line of izapple2,
with the System Master in drive 1 and the reader's diskette in drive 2
*/
const dos33Machine = `izapple2 -model _base -board 2plus -cpu 6502 \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s0 language \
    -s6 'diskii,disk1=disks/DOS 3.3 System Master - 680-0210-A (1982).dsk,disk2=my-disk.dsk'`

// dos33OwnDisk is the same machine with the reader's diskette alone
const dos33OwnDisk = `izapple2 -model _base -board 2plus -cpu 6502 \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s0 language \
    -s6 diskii,disk1=my-disk.dsk`

/*
dos33Screenshots is life with DOS 3.3 on an Apple ][+ with two drives: the
System Master started, its catalog, its greeting program listed, a blank
diskette initialized with a program of our own, and the machine started from
it.

	cp disks/blank.dsk my-disk.dsk
*/
func dos33Screenshots(t *testing.T) {
	pictures := newAlbum("dos33", album.Green)
	myDisk := disk(t, "blank.dsk")
	o := start(t, dos33Machine, map[string]string{"my-disk.dsk": myDisk})

	// Started, recorded until the prompt
	booting := pictures.Record(o)
	booting.Capture(20)
	for !o.HasText("COPYRIGHT") {
		if o.Frames() > 60*60 {
			t.Fatal("the System Master did not start")
		}
		booting.Run(12, 6)
	}
	booting.Run(60, 6)
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
	again := start(t, dos33OwnDisk, map[string]string{"my-disk.dsk": myDisk})
	must(t, again.WaitForText("ADA'S DISK", 120))
	again.Run(60)
	must(t, pictures.Screenshot(again, "own-disk"))
}
