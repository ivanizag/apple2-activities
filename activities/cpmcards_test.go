package activities

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
	"github.com/ivanizag/apple2-activities/operator"
	"github.com/ivanizag/izapple2/screen"
)

// The Apple ][+, without and with a Language Card in slot 0, and with a Videx
// Videoterm in slot 3 as well, as the start of a command line of izapple2
const (
	cpmCardsPlusBare = `izapple2 -model none -board 2plus -cpu 6502 -screen green \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps`
	cpmCardsPlusLC = cpmCardsPlusBare + ` \
    -s0 language`
	cpmCardsPlusVidex = cpmCardsPlusLC + ` \
    -s3 videx`
)

// The Apple //e, with its 128 KB, with a Videx Videoterm in slot 3 as well,
// with a RamWorks of 1 MB in its auxiliary slot instead of the 80 column card,
// and the enhanced //e, as the start of a command line of izapple2
const (
	cpmCardsIIeBase = `izapple2 -model none -board 2e -cpu 6502 -screen green \
    -rom "<internal>/Apple2e.rom" \
    -charrom "<internal>/Apple IIe Video Unenhanced.bin" \
    -s0 language`
	cpmCardsIIeVidex = cpmCardsIIeBase + ` \
    -s3 videx`
	cpmCardsIIeRamWorks = `izapple2 -model none -board 2e -cpu 6502 -screen green \
    -rom "<internal>/Apple2e.rom" \
    -charrom "<internal>/Apple IIe Video Unenhanced.bin" \
    -ramworks 1024 \
    -s0 language`
	cpmCardsIIeEnhanced = `izapple2 -model none -board 2e -cpu 65c02 -screen green \
    -rom "<internal>/Apple2e_Enhanced.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language`
)

// cpmCardsOn is a machine with a Z80 card in slot 4, with a disk in drive 1
// and CP/M 2.23 in drive 2, as its command line of izapple2
func cpmCardsOn(machine string, card string, disk string) string {
	return machine + ` \
    -s4 ` + card + ` \
    -s6 'diskii,disk1=disks/` + disk + `,disk2=disks/CPM60K.DSK'`
}

// The disks of the versions of CP/M of the guide
const (
	cpmCardsDisk220     = "Softcard 16-sector disk (Microsoft 1980).dsk"
	cpmCardsDisk220B    = "CPM1.PO"
	cpmCardsDisk223     = "CPM60K.DSK"
	cpmCardsDisk223B    = "CPM2.23(60k).dsk"
	cpmCardsDisk30      = "CPM3.1_Disk_1.dsk"
	cpmCardsDiskCPAM40  = "CPAM40B.dsk"
	cpmCardsDiskCPAM51  = "A.E. CPM Plus card.dsk"
	cpmCardsDiskCPAM511 = "CPAM51A.dsk"
	cpmCardsDisk228B    = "Microsoft Softcard II CPM 2.28B.DSK"
	cpmCardsDiskPCPI    = "pcpiboot.do"
)

// The versions of CP/M of the guide, each the machine that starts it
var (
	cpmCards220      = cpmCardsOn(cpmCardsPlusLC, "z80softcard", cpmCardsDisk220)
	cpmCards220B     = cpmCardsOn(cpmCardsPlusLC, "z80softcard", cpmCardsDisk220B)
	cpmCards223      = cpmCardsOn(cpmCardsIIeBase, "z80softcard", cpmCardsDisk223)
	cpmCards223B     = cpmCardsOn(cpmCardsPlusLC, "z80softcard", cpmCardsDisk223B)
	cpmCards30       = cpmCardsOn(cpmCardsIIeBase, "z80softcard", cpmCardsDisk30)
	cpmCardsCPAM40   = cpmCardsOn(cpmCardsIIeBase, "z80softcard", cpmCardsDiskCPAM40)
	cpmCardsCPAM51   = cpmCardsOn(cpmCardsIIeBase, "z80softcard", cpmCardsDiskCPAM51)
	cpmCardsCPAM511  = cpmCardsOn(cpmCardsIIeBase, "z80softcard", cpmCardsDiskCPAM511)
	cpmCards228B     = cpmCardsOn(cpmCardsIIeBase, "z80softcard2", cpmCardsDisk228B)
	cpmCardsPCPI     = cpmCardsOn(cpmCardsIIeBase, "applicard", cpmCardsDiskPCPI)
	cpmCardsPCPIPlus = cpmCardsOn(cpmCardsPlusLC, "applicard", cpmCardsDiskPCPI)
	cpmCards30Videx  = cpmCardsOn(cpmCardsPlusVidex, "z80softcard", cpmCardsDisk30)
	cpmCardsRamWorks = cpmCardsOn(cpmCardsIIeRamWorks, "z80softcard", cpmCardsDiskCPAM511)
	cpmCardsMachines = []string{cpmCards220, cpmCards220B, cpmCards223, cpmCards223B, cpmCards30, cpmCardsCPAM40, cpmCardsCPAM51, cpmCardsCPAM511, cpmCards228B, cpmCardsPCPI, cpmCardsPCPIPlus, cpmCards30Videx, cpmCardsRamWorks}
)

/*
cpmCardsVersions are the versions in the order of the guide, each with the
name of its pictures, whether the files of its disk are shown, where its
BDOS starts, which is the top of the memory for the programs, where the
warm start of its BIOS is, and what GBASIC says it leaves free when the disk
has it
*/
var cpmCardsVersions = []struct {
	picture string
	machine string
	dir     bool
	bdos    string
	bios    string
	gbasic  string
}{
	{"softcard-2.20", cpmCards220, false, "9C06", "AA03", "5105 BYTES FREE"},
	{"softcard-2.20b", cpmCards220B, false, "CC06", "DA03", "17393 BYTES FREE"},
	{"softcard-2.23", cpmCards223, true, "DC06", "FA03", "21489 Bytes free"},
	{"datasoft-2.23b", cpmCards223B, false, "CC06", "DA03", "17393 BYTES FREE"},
	{"softcard-3.0", cpmCards30, true, "BA06", "D903", ""},
	{"cpam-4.0b", cpmCardsCPAM40, false, "9C06", "AA03", ""},
	{"cpam-5.1", cpmCardsCPAM51, false, "CD06", "DB03", ""},
	{"cpam-5.1.1", cpmCardsCPAM511, true, "CD06", "DB03", ""},
	{"softcard-ii-2.28b", cpmCards228B, true, "EB06", "F903", "32807 Bytes free"},
	{"applicard-2.0", cpmCardsPCPI, true, "E506", "F303", ""},
	{"applicard-2.0-plus", cpmCardsPCPIPlus, false, "E506", "F303", ""},
}

// cpmCardsGBASIC are the versions whose GBASIC is pictured
var cpmCardsGBASIC = []string{"softcard-2.20", "softcard-2.23", "softcard-ii-2.28b"}

/*
cpmCardsScreenshots is CP/M started on each of the Z80 cards of izapple2, a
version after another: the screen it starts with, for some the files of its
disk, where its BDOS and its BIOS are, and the memory GBASIC finds free,
where the disk has it. Then CP/M 3 on a ][+ with a Videx card, the RAM disk
of CP/AM on a //e with a RamWorks, and each version on six machines.
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
		cpmCardsTPA(t, o, v.picture == "softcard-3.0", v.bdos, v.bios)
		if v.picture == cpmCardsTPAPicture {
			must(t, pictures.Screenshot(o, "tpa"))
		}
		must(t, o.TypeLines("G0"))
		must(t, o.WaitForText("\nA>", 30))
		must(t, o.WaitForPrompt(30))
		if v.gbasic != "" {
			must(t, o.TypeLines("GBASIC"))
			must(t, o.WaitForText(v.gbasic, 60))
			must(t, o.WaitForPrompt(30))
			if slices.Contains(cpmCardsGBASIC, v.picture) {
				must(t, pictures.Screenshot(o, v.picture+"-gbasic"))
			}
		}
	}

	o := start(t, cpmCards30Videx, nil)
	must(t, o.WaitForText("A>", 120))
	must(t, o.WaitForPrompt(30))
	must(t, pictures.Screenshot(o, "softcard-3.0-videx"))

	o = start(t, cpmCardsRamWorks, nil)
	must(t, o.WaitForText("A>", 120))
	must(t, o.WaitForPrompt(30))
	must(t, o.TypeLines("MEGDRIVE"))
	must(t, o.WaitForText("Available storage space is 1010k", 30))
	must(t, o.WaitForPrompt(30))
	must(t, o.TypeLines("PIP C:=A:STAT.COM"))
	must(t, o.WaitForPrompt(60))
	must(t, o.TypeLines("C:", "DIR"))
	must(t, o.WaitForText("C: STAT     COM", 30))
	must(t, o.WaitForPrompt(30))
	must(t, o.TypeLines("STAT C:"))
	must(t, o.WaitForText("C: R/W, Space: 1004k", 30))
	must(t, o.WaitForPrompt(30))
	must(t, pictures.Screenshot(o, "cpam-5.1.1-megdrive"))

	cpmCardsCompare(t)
}

// cpmCardsTPAPicture is the version whose DDT is pictured reading the top of
// the memory for the programs
const cpmCardsTPAPicture = "softcard-ii-2.28b"

/*
cpmCardsTPA finds where the BDOS of CP/M starts, the top of the memory left
to the programs, the TPA, and where the BIOS is. The jump at 5, which the
programs call CP/M by, goes to the BDOS, and the jump at 0 to the warm start
of the BIOS, but under DDT the jump at 5 goes to DDT, which puts itself
under the BDOS. So DDT, from the disk of drive B, assembles a program of five
instructions that copies the addresses of the two jumps to 3000 and 3002, and
the program is saved and run without DDT; then DDT shows what it copied, and
waits. CP/M 3 saves with a SAVE of its own, started before the program to
save.
*/
func cpmCardsTPA(t *testing.T, o *operator.Operator, cpm3 bool, bdos string, bios string) {
	if cpm3 {
		must(t, o.TypeLines("SAVE"))
		must(t, o.WaitForPrompt(30))
	}
	must(t, o.TypeLines("B:DDT"))
	must(t, o.WaitForPrompt(30))
	must(t, o.TypeLines("A100", "LHLD 6", "SHLD 3000", "LHLD 1", "SHLD 3002", "JMP 0", "", "G0"))
	if cpm3 {
		must(t, o.WaitForText("Enter file", 30))
		must(t, o.TypeLines("TPA.COM", "100", "10E"))
	} else {
		must(t, o.WaitForPrompt(30))
		must(t, o.TypeLines("SAVE 1 TPA.COM"))
	}
	must(t, o.WaitForPrompt(30))
	must(t, o.TypeLines("TPA"))
	must(t, o.WaitForPrompt(30))
	must(t, o.TypeLines("B:DDT"))
	must(t, o.WaitForPrompt(30))
	must(t, o.TypeLines("D3000,3003"))
	must(t, o.WaitForText("3000 "+bdos[2:]+" "+bdos[:2]+" "+bios[2:]+" "+bios[:2], 30))
	must(t, o.WaitForPrompt(30))
}

/*
cpmCardsMachinesCompared are the machines each version is started on to
compare them: the ][+ without and with a Language Card, the ][+ with a Videx
card as well, the //e, the //e with a Videx card and the enhanced //e
*/
var cpmCardsMachinesCompared = []string{cpmCardsPlusBare, cpmCardsPlusLC, cpmCardsPlusVidex, cpmCardsIIeBase, cpmCardsIIeVidex, cpmCardsIIeEnhanced}

/*
cpmCardsCompared is what each version does on each machine of
cpmCardsMachinesCompared: "" doesn't start, "40" starts in 40 columns, "80"
in the 80 columns of the //e, "videx" in those of the Videx card, and
"signs" starts in 40 columns with its lower case shown as other signs. The
Applicard says the size of its buffer for the printer, which is checked too.
*/
var cpmCardsCompared = []struct {
	card, disk string
	on         []string
	says       []string
}{
	{"z80softcard", cpmCardsDisk220, []string{"40", "40", "", "", "", ""}, nil},
	{"z80softcard", cpmCardsDisk220B, []string{"", "40", "", "", "", ""}, nil},
	{"z80softcard", cpmCardsDisk223, []string{"", "40", "videx", "80", "80", "80"}, nil},
	{"z80softcard", cpmCardsDisk223B, []string{"", "40", "", "", "", ""}, nil},
	{"z80softcard", cpmCardsDisk30, []string{"", "", "videx", "80", "80", "80"}, nil},
	{"z80softcard", cpmCardsDiskCPAM40, []string{"40", "40", "videx", "80", "80", "80"}, nil},
	{"z80softcard", cpmCardsDiskCPAM51, []string{"", "", "videx", "80", "80", "80"}, nil},
	{"z80softcard", cpmCardsDiskCPAM511, []string{"", "", "videx", "80", "80", "80"}, nil},
	{"z80softcard2", cpmCardsDisk228B, []string{"", "signs", "videx", "80", "videx", "80"}, nil},
	{"applicard", cpmCardsDiskPCPI, []string{"40", "40", "videx", "80", "80", "80"},
		[]string{"19.5K", "35.5K", "35.5K", "83.5K", "83.5K", "83.5K"}},
}

// cpmCardsCompare starts each version on each machine compared, and checks
// that it does there what the guide says
func cpmCardsCompare(t *testing.T) {
	for _, c := range cpmCardsCompared {
		for i, machine := range cpmCardsMachinesCompared {
			got, text := cpmCardsStart(t, cpmCardsOn(machine, c.card, c.disk))
			if got != c.on[i] {
				t.Errorf("%v on machine %v: got %q, want %q\n%v", c.disk, i, got, c.on[i], text)
			}
			if c.says != nil && !strings.Contains(strings.ToUpper(text), c.says[i]+" PRINTER BUFFER") {
				t.Errorf("%v on machine %v: no %v buffer\n%v", c.disk, i, c.says[i], text)
			}
		}
	}
}

/*
cpmCardsStart starts a machine and says how CP/M starts on it: "" when its
prompt doesn't come, or the 6502 stops on an instruction it doesn't have,
else how it shows its text, as cpmCardsCompared; with the text on the screen
*/
func cpmCardsStart(t *testing.T, machine string) (shown string, text string) {
	o := start(t, machine, nil)
	defer func() {
		if r := recover(); r != nil {
			shown, text = "", fmt.Sprint(r)
		}
	}()
	if o.WaitForPrompt(60) != nil || !o.WaitUntil(30, func() bool {
		return strings.Contains(o.Text(), "\nA>") || strings.Contains(o.Text(), "\n>")
	}) {
		return "", o.Text()
	}
	text = o.Text()
	switch o.Apple2().GetVideoSource().GetCurrentVideoMode() & screen.VideoBaseMask {
	case screen.VideoVidex:
		return "videx", text
	case screen.VideoText80:
		return "80", text
	}
	if !strings.Contains(text, "A>") {
		return "signs", text
	}
	return "40", text
}
