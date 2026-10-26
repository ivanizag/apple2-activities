package activities

import (
	"slices"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// cpmCardsPlus is an Apple ][+ with a Language Card and a Z80 card in slot 4,
// with a disk in its drive, as its command line of izapple2
func cpmCardsPlus(card string, disk string) string {
	return `izapple2 -model none -board 2plus -cpu 6502 -screen green \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s0 language \
    -s4 ` + card + ` \
    -s6 'diskii,disk1=disks/` + disk + `'`
}

// cpmCardsIIe is an Apple //e with a Z80 card in slot 4, with a disk in its
// drive, as its command line of izapple2
func cpmCardsIIe(card string, disk string) string {
	return `izapple2 -model none -board 2e -cpu 6502 -screen green \
    -rom "<internal>/Apple2e.rom" \
    -charrom "<internal>/Apple IIe Video Unenhanced.bin" \
    -s0 language \
    -s4 ` + card + ` \
    -s6 'diskii,disk1=disks/` + disk + `'`
}

// The versions of CP/M of the guide, each the machine that starts it
var (
	cpmCards220      = cpmCardsPlus("z80softcard", "Softcard 16-sector disk (Microsoft 1980).dsk")
	cpmCards220B     = cpmCardsPlus("z80softcard", "CPM1.PO")
	cpmCards223B     = cpmCardsPlus("z80softcard", "CPM2.23(60k).dsk")
	cpmCards223      = cpmCardsIIe("z80softcard", "CPM60K.DSK")
	cpmCards30       = cpmCardsIIe("z80softcard", "CPM3.1_Disk_1.dsk")
	cpmCardsCPAM40   = cpmCardsIIe("z80softcard", "CPAM40B.dsk")
	cpmCardsCPAM51   = cpmCardsIIe("z80softcard", "A.E. CPM Plus card.dsk")
	cpmCardsCPAM511  = cpmCardsIIe("z80softcard", "CPAM51A.dsk")
	cpmCards228B     = cpmCardsIIe("z80softcard2", "Microsoft Softcard II CPM 2.28B.DSK")
	cpmCardsPCPI     = cpmCardsIIe("applicard", "pcpiboot.do")
	cpmCardsPCPIPlus = cpmCardsPlus("applicard", "pcpiboot.do")
)

/*
cpmCardsVersions are the versions in the order of the guide, each with the
name of its pictures, whether the files of its disk are shown, and what
GBASIC says it leaves free when the disk has it
*/
var cpmCardsVersions = []struct {
	picture string
	machine string
	dir     bool
	gbasic  string
}{
	{"softcard-2.20", cpmCards220, false, "5105 BYTES FREE"},
	{"softcard-2.20b", cpmCards220B, false, "17393 BYTES FREE"},
	{"datasoft-2.23b", cpmCards223B, false, "17393 BYTES FREE"},
	{"softcard-2.23", cpmCards223, true, "21489 Bytes free"},
	{"softcard-3.0", cpmCards30, true, ""},
	{"cpam-4.0b", cpmCardsCPAM40, false, ""},
	{"cpam-5.1", cpmCardsCPAM51, false, ""},
	{"cpam-5.1.1", cpmCardsCPAM511, true, ""},
	{"softcard-ii-2.28b", cpmCards228B, true, "32807 Bytes free"},
	{"applicard-2.0", cpmCardsPCPI, true, ""},
	{"applicard-2.0-plus", cpmCardsPCPIPlus, false, ""},
}

// cpmCardsGBASIC are the versions whose GBASIC is pictured
var cpmCardsGBASIC = []string{"softcard-2.20", "softcard-2.23", "softcard-ii-2.28b"}

/*
cpmCardsScreenshots is CP/M started on each of the Z80 cards of izapple2, a
version after another: the screen it starts with, for some the files of its
disk, and the memory GBASIC finds free, where the disk has it.
*/
func cpmCardsScreenshots(t *testing.T) {
	pictures := newAlbum("cpm-cards", album.Green)
	for _, v := range cpmCardsVersions {
		o := start(t, v.machine, nil)
		must(t, o.WaitForText("A>", 120))
		must(t, o.WaitForPrompt(30))
		must(t, pictures.Screenshot(o, v.picture))
		if v.dir {
			must(t, o.TypeLines("DIR"))
			must(t, o.WaitForPrompt(30))
			must(t, pictures.Screenshot(o, v.picture+"-dir"))
		}
		if v.gbasic != "" {
			must(t, o.TypeLines("GBASIC"))
			must(t, o.WaitForText(v.gbasic, 60))
			must(t, o.WaitForPrompt(30))
			if slices.Contains(cpmCardsGBASIC, v.picture) {
				must(t, pictures.Screenshot(o, v.picture+"-gbasic"))
			}
		}
	}
}
