package activities

import (
	"image"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

/*
mockingboardScreenshots is the demonstration disk of the Mockingboard, the
sound card of Sweet Micro Systems, on an Apple ][+: its title, its menus, and
its sound effects, recorded.

	izapple2 -model 2plus -s4 mockingboard disks/mockingboard-demo.dsk
*/
func mockingboardScreenshots(t *testing.T) {
	pictures := newAlbum("mockingboard", album.Color)
	o := start(t, "2plus", map[string]string{"s4": "mockingboard"}, disk(t, "mockingboard-demo.dsk"))
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
