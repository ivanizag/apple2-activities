package activities

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// dos32Machine is an Apple ][+ with the Disk II controller of 13 sectors and
// the System Master of DOS 3.2
const dos32Machine = `izapple2 -model none -board 2plus -cpu 6502 -screen green \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s6 'diskii,sectors13,disk1="disks/Apple DOS 3.2 Plus.nib"'`

// dos32Upgraded is the same Apple ][+ with the controller of 16 sectors, the
// System Master of DOS 3.3 in drive 1 and the one of DOS 3.2 in drive 2. What
// it writes is kept in the folder changes, and the images left as they are.
const dos32Upgraded = `izapple2 -model none -board 2plus -cpu 6502 -screen green \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -saveDir changes \
    -s6 'diskii,disk1=disks/DOS 3.3 System Master - 680-0210-A (1982).dsk,disk2="disks/Apple DOS 3.2 Plus.nib"'`

/*
dos32Screenshots is the move from 13 sectors to 16: DOS 3.2 on the
controller of 13 sectors, then DOS 3.3 on the one of 16, which can't read
the old diskette, the old diskette started anyway with START13, and a
program moved from it to the new format with MUFFIN, and run there.
*/
func dos32Screenshots(t *testing.T) {
	pictures := newAlbum("dos32", album.Green)

	// DOS 3.2, and the catalog of its System Master
	o := start(t, dos32Machine, nil)
	must(t, o.WaitForText("MASTER DISKETTE VERISON 3.2 PLUS", 30))
	must(t, o.WaitForKeyboard(30))
	must(t, o.TypeLines("CATALOG"))
	must(t, o.WaitForText("HOPALONG CASSIDY", 30))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "dos32"))

	// DOS 3.3 on the controller of 16 sectors: the old diskette can't be read
	changes := filepath.Join(t.TempDir(), "changes")
	must(t, os.Mkdir(changes, 0o755))
	o = start(t, dos32Upgraded, map[string]string{"changes": changes})
	must(t, o.WaitForText("SYSTEM MASTER", 30))
	must(t, o.WaitForKeyboard(30))
	must(t, o.TypeLines("CATALOG,D2"))
	must(t, o.WaitForText("I/O ERROR", 30))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "io-error"))

	// MUFFIN: Little Brick Out, from the diskette of DOS 3.2 in drive 2 to
	// the one of DOS 3.3 in drive 1
	must(t, o.TypeLines("BRUN MUFFIN,D1"))
	must(t, o.WaitForText("WHICH WOULD YOU LIKE?", 30))
	must(t, o.WaitForKeyboard(10))
	must(t, pictures.Screenshot(o, "muffin"))
	must(t, o.TypeLines("1"))
	must(t, o.WaitForText("SOURCE SLOT?", 10))
	must(t, o.TypeLines("6", "2", "6", "1", "LITTLE BRICK OUT"))
	must(t, o.WaitForText("ANY OTHER KEY TO BEGIN", 10))
	must(t, o.WaitForKeyboard(10))
	must(t, o.Key("Return"))
	must(t, o.WaitForText("PRESS ANY KEY TO CONTINUE", 60))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "converted"))
	if kept, _ := os.ReadDir(changes); len(kept) == 0 {
		t.Fatal("what MUFFIN wrote is not kept in the folder of -saveDir")
	}

	// Back in DOS 3.3, the program on the new diskette
	must(t, o.Key("Return"))
	must(t, o.WaitForText("WHICH WOULD YOU LIKE?", 10))
	must(t, o.TypeLines("2"))
	must(t, o.WaitForKeyboard(10))
	must(t, o.TypeLines("CATALOG,D1"))
	must(t, o.WaitForText("BOOT13", 30))
	must(t, o.WaitForKeyboard(10))
	must(t, o.Key("Space"))
	must(t, o.WaitForText("LITTLE BRICK OUT", 30))
	must(t, o.WaitForKeyboard(10))
	o.Run(30)
	must(t, pictures.Screenshot(o, "catalog-33"))

	// The old diskette started on the new controller, with START13: it
	// boots drive 1, so the old diskette goes there
	must(t, o.TypeLines("RUN START13"))
	must(t, o.WaitForText("SLOT TO BOOT FROM", 30))
	must(t, o.WaitForKeyboard(10))
	must(t, pictures.Screenshot(o, "start13"))
	must(t, o.InsertDisk(0, disk(t, "Apple DOS 3.2 Plus.nib")))
	must(t, o.TypeLines(""))
	must(t, o.WaitForText("MASTER DISKETTE VERISON 3.2 PLUS", 30))
	must(t, o.WaitForKeyboard(30))
	o.Run(30)
	must(t, pictures.Screenshot(o, "boot13"))
}
