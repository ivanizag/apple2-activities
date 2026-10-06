package activities

import (
	"image"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// mockingboardMachine is the machine of the guide, as its command line of izapple2
const mockingboardMachine = `izapple2 -model _base -board 2plus -cpu 6502 \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s4 mockingboard \
    -s6 diskii,disk1=disks/Mockingboard_Sound_and_Speech_I_Demo_Disk_Apple_II_Plus_Sweet_Micro_Systems_1982.dsk`

/*
mockingboardScreenshots is the demonstration disk of the Mockingboard, the
sound card of Sweet Micro Systems, on an Apple ][+: its title, its menus, and
its sound effects, recorded.
*/
func mockingboardScreenshots(t *testing.T) {
	pictures := newAlbum("mockingboard", album.Color)
	o := start(t, mockingboardMachine, nil)
	sound := album.Listen(o)

	// The title page, drawn a few lines at a time, kept as it is last
	// before the menu comes
	var title *image.RGBA
	for !(o.InTextMode() && o.HasText("SELECT A LETTER")) {
		if o.Frames() > 300*60 {
			t.Fatal("the menu did not come")
		}
		if !o.InTextMode() {
			title = pictures.Screen(o)
		}
		o.Run(30)
	}
	if title == nil {
		t.Fatal("there was no title page")
	}
	must(t, pictures.Write(title, "title"))

	// The main menu
	must(t, o.WaitForPrompt(10))
	must(t, pictures.Screenshot(o, "menu"))

	// The sound effects, three seconds each
	must(t, o.Type("A"))
	must(t, o.WaitForShownText("SOUND EFFECTS MENU", 30))
	must(t, o.WaitForPrompt(10))
	must(t, pictures.Screenshot(o, "effects-menu"))
	from := sound.Now()
	for _, effect := range "ABDEK" {
		must(t, o.Type(string(effect)))
		o.RunSeconds(3)
	}
	sound.Stop()
	must(t, pictures.SaveSound(sound.Clip(from, sound.Now()), "effects"))
}
