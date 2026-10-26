package activities

import (
	"math"
	"slices"
	"strings"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
)

// cpmCardsPlus is an Apple ][+ with a Language Card and a Z80 card in slot 4,
// with a disk in drive 1 and CP/M 2.23 in drive 2, as its command line of
// izapple2
func cpmCardsPlus(card string, disk string) string {
	return `izapple2 -model none -board 2plus -cpu 6502 -screen green \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s0 language \
    -s4 ` + card + ` \
    -s6 'diskii,disk1=disks/` + disk + `,disk2=disks/CPM60K.DSK'`
}

// cpmCardsIIe is an Apple //e with a Z80 card in slot 4, with a disk in drive
// 1 and CP/M 2.23 in drive 2, as its command line of izapple2
func cpmCardsIIe(card string, disk string) string {
	return `izapple2 -model none -board 2e -cpu 6502 -screen green \
    -rom "<internal>/Apple2e.rom" \
    -charrom "<internal>/Apple IIe Video Unenhanced.bin" \
    -s0 language \
    -s4 ` + card + ` \
    -s6 'diskii,disk1=disks/` + disk + `,disk2=disks/CPM60K.DSK'`
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
name of its pictures, whether the files of its disk are shown, where its
BDOS starts, which is the top of the memory for the programs, how many
seconds a loop of the Z80 takes in izapple2, and what GBASIC
says it leaves free when the disk has it
*/
var cpmCardsVersions = []struct {
	picture string
	machine string
	dir     bool
	bdos    string
	loop    float64
	gbasic  string
}{
	{"softcard-2.20", cpmCards220, false, "9C06", 4.1, "5105 BYTES FREE"},
	{"softcard-2.20b", cpmCards220B, false, "CC06", 4.1, "17393 BYTES FREE"},
	{"datasoft-2.23b", cpmCards223B, false, "CC06", 4.1, "17393 BYTES FREE"},
	{"softcard-2.23", cpmCards223, true, "DC06", 4.1, "21489 Bytes free"},
	{"softcard-3.0", cpmCards30, true, "BA06", 4.1, ""},
	{"cpam-4.0b", cpmCardsCPAM40, false, "9C06", 4.1, ""},
	{"cpam-5.1", cpmCardsCPAM51, false, "CD06", 4.1, ""},
	{"cpam-5.1.1", cpmCardsCPAM511, true, "CD06", 4.1, ""},
	{"softcard-ii-2.28b", cpmCards228B, true, "EB06", 5.1, "32807 Bytes free"},
	{"applicard-2.0", cpmCardsPCPI, true, "E506", 6.8, ""},
	{"applicard-2.0-plus", cpmCardsPCPIPlus, false, "E506", 6.8, ""},
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
		cpmCardsTPA(t, o, v.picture == "softcard-3.0", v.bdos)
		if v.picture == cpmCardsTPAPicture {
			must(t, pictures.Screenshot(o, "tpa"))
		}
		must(t, o.TypeLines("G0"))
		must(t, o.WaitForText("\nA>", 30))
		must(t, o.WaitForPrompt(30))
		cpmCardsLoop(t, o, v.picture == "softcard-3.0", v.loop)
		if v.picture == cpmCardsLoopPicture {
			must(t, pictures.Screenshot(o, "loop"))
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

// cpmCardsTPAPicture is the version whose DDT is pictured reading the top of
// the memory for the programs
const cpmCardsTPAPicture = "softcard-ii-2.28b"

/*
cpmCardsSave assembles a program in DDT, from the disk of drive B, at 0100,
and saves it with the name given, to the address of its last byte. The CCP
of CP/M 2 saves pages with its SAVE once DDT has ended; CP/M 3 has a SAVE
of its own instead, started before the program to save, which asks for the
name and the addresses when the program ends.
*/
func cpmCardsSave(t *testing.T, o *operator.Operator, cpm3 bool, name string, last string, program ...string) {
	if cpm3 {
		must(t, o.TypeLines("SAVE"))
		must(t, o.WaitForPrompt(30))
	}
	must(t, o.TypeLines("B:DDT"))
	must(t, o.WaitForPrompt(30))
	must(t, o.TypeLines("A100"))
	must(t, o.TypeLines(program...))
	must(t, o.TypeLines("", "G0"))
	if cpm3 {
		must(t, o.WaitForText("Enter file", 30))
		must(t, o.TypeLines(name+".COM", "100", last))
	} else {
		must(t, o.WaitForPrompt(30))
		must(t, o.TypeLines("SAVE 1 "+name+".COM"))
	}
	must(t, o.WaitForPrompt(30))
}

/*
cpmCardsTPA finds where the BDOS of CP/M starts, the top of the memory left
to the programs, the TPA. The jump at 5, which the programs call CP/M by,
goes there, but under DDT it goes to DDT, which puts itself under the BDOS.
So a program of three instructions copies the address of the jump to 3000,
run without DDT; then DDT shows what it copied, and waits.
*/
func cpmCardsTPA(t *testing.T, o *operator.Operator, cpm3 bool, bdos string) {
	cpmCardsSave(t, o, cpm3, "TPA", "108", "LHLD 6", "SHLD 3000", "JMP 0")
	must(t, o.TypeLines("TPA"))
	must(t, o.WaitForPrompt(30))
	must(t, o.TypeLines("B:DDT"))
	must(t, o.WaitForPrompt(30))
	must(t, o.TypeLines("D3000,3001"))
	must(t, o.WaitForText("3000 "+bdos[2:]+" "+bdos[:2], 30))
	must(t, o.WaitForPrompt(30))
}

// cpmCardsLoopPicture is the version whose loop is pictured
const cpmCardsLoopPicture = "applicard-2.0"

/*
cpmCardsLoop times the Z80: a program prints a star, counts BC down from
65536 to 0 sixteen times, and prints another. It is the time between the
stars, which leaves out the loading of the program, as a person with a
stopwatch would. The time is checked against the one the page gives, to
the tenth of a second.
*/
func cpmCardsLoop(t *testing.T, o *operator.Operator, cpm3 bool, seconds float64) {
	cpmCardsSave(t, o, cpm3, "LOOP", "11D",
		"MVI E,2A", "MVI C,2", "CALL 5",
		"MVI D,10", "LXI B,0", "DCX B", "MOV A,B", "ORA C", "JNZ 10C",
		"DCR D", "JNZ 109",
		"MVI E,2A", "MVI C,2", "CALL 5", "RET")
	must(t, o.TypeLines("LOOP"))
	stars := func(n int) func() bool {
		return func() bool {
			text := o.Text()
			i := strings.LastIndex(text, "A>LOOP")
			return i >= 0 && strings.Count(text[i:], "*") >= n
		}
	}
	if !o.WaitUntil(30, stars(1)) {
		t.Fatalf("LOOP printed no star:\n%v", o.Text())
	}
	started := o.Frames()
	if !o.WaitUntil(60, stars(2)) {
		t.Fatalf("LOOP did not end:\n%v", o.Text())
	}
	took := float64(o.Frames()-started) / operator.FramesPerSecond
	if math.Abs(took-seconds) > 0.1 {
		t.Errorf("LOOP took %.2f seconds, the page says %.1f", took, seconds)
	}
	must(t, o.WaitForPrompt(30))
}
