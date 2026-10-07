package activities

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/ivanizag/apple2-activities/album"
)

// TestProbe runs a game through keys from PROBE_KEYS, one item a line:
// text to type (\n for Return), or "disk:FILE" to insert, and after each
// runs 15 seconds and prints the last lines and where the processor is
func TestProbe(t *testing.T) {
	machine := os.Getenv("PROBE_MACHINE")
	if machine == "" {
		t.Skip()
	}
	pictures := album.New(os.Getenv("PROBE_DIR"), album.ColorWhiteText)
	o := start(t, machine, nil)
	step := func(n int, name string) {
		o.RunSeconds(15)
		h := map[uint16]int{}
		for range 120 {
			o.Run(1)
			h[o.Apple2().GetPC()]++
		}
		keys := []int{}
		for k := range h {
			keys = append(keys, int(k))
		}
		sort.Ints(keys)
		lo, hi := keys[0], keys[len(keys)-1]
		fmt.Printf("== %02d %q pcs %04X-%04X (%d)\n%q\n", n, name, lo, hi, len(keys), o.Text())
		must(t, pictures.Screenshot(o, fmt.Sprintf("%02d", n)))
	}
	step(0, "boot")
	for i, k := range strings.Split(os.Getenv("PROBE_KEYS"), "\n") {
		if k == "" {
			continue
		}
		if strings.HasPrefix(k, "disk:") {
			must(t, o.InsertDisk(0, disk(t, k[5:])))
			o.RunSeconds(3)
			continue
		}
		must(t, o.Type(strings.ReplaceAll(k, `\n`, "\n")))
		step(i+1, k)
	}
}
