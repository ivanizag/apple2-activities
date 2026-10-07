package activities

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestDump(t *testing.T) {
	m := os.Getenv("DUMP_MACHINE")
	if m == "" {
		t.Skip()
	}
	o := start(t, m, nil)
	o.RunSeconds(20)
	a, _ := strconv.ParseUint(os.Getenv("DUMP_AT"), 16, 16)
	for pc := uint16(a); pc < uint16(a)+0x40; pc += 16 {
		fmt.Printf("%04X:", pc)
		for i := range uint16(16) {
			fmt.Printf(" %02X", o.Apple2().Peek(pc+i))
		}
		fmt.Println()
	}
	fmt.Printf("%q\n", o.Text())
}
