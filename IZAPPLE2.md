# Open findings in izapple2

What making the activities found in [izapple2](https://github.com/ivanizag/izapple2),
the state of each, and what to change here when it is fixed. When one is done
there and done here, take it out of this list.

## Changes waiting to reach izapple2's master

### The library calls the operator uses

`ScreenText`, `Peek`, `GetPC` and `LoadDisk` on `Apple2`, in `inspect.go`, and
the text of a Videx card read only while it is shown. Branch `activities-api`.

**Here, once on master:** `GOWORK=off go get github.com/ivanizag/izapple2@master`,
so that `go.mod` names a version that has them and the repository builds
without a `go.work`. Drop the workspace paragraph of [AGENTS.md](AGENTS.md)
if nothing else needs it.

### INIT of a blank diskette hangs the emulator

`nibDecodeTrack` in `storage/fileNib.go` went around the track forever when
an address prolog started on its last bytes, which a track being written can
have. Fixed on `activities-api`, with a test. It is in the released izapple2
too: a reader following [Life with DOS 3.3](guides/dos33.md) with a release
from before the fix sees the machine stop at `INIT HELLO,D2`.

**Here, once released:** say in [dos33.md](guides/dos33.md) which release of
izapple2 the page needs, at least.

### The Disk II is seven times too slow with DSK, PO and NIB images

DOS checks whether the disk still turns at the start of each access, and
izapple2 stopped the motor at once and gave a new nibble on every read, so
DOS waited for the motor to come up to speed on every access: 36 million
cycles to boot DOS 3.3 instead of about 5. WOZ images, on the sequencer card,
were not affected. Branch `disk-speed`, with tests.

**Here, once on master:** make the pictures again and look at them, as the
timing of everything that reads a 5¼ diskette changes:

- [dos33.md](guides/dos33.md): the boot is recorded `Faster(10)` and the page
  says it is ten times faster; see if real speed is fine now, and fix the
  text.
- [prodos.md](guides/prodos.md): "it takes a while to load in izapple2" for
  BASIC.SYSTEM, about forty seconds now; reword.
- [pascal.md](guides/pascal.md), [cpm.md](guides/cpm.md),
  [mockingboard.md](guides/mockingboard.md) (its title is taken at a fixed
  110 seconds), [ultraterm.md](guides/ultraterm.md) and
  [card-cat.md](guides/card-cat.md): check the pictures still show what the
  text says.
- [lode-runner.md](guides/lode-runner.md) is a WOZ and takes its pictures at
  fixed times; check it is unchanged.

## Not fixed

### Pictures change from run to run: host time

The machine reads the clock of the host in a few places, so the same steps do
not always make the same pictures:

- the flashing characters and the cursor, `renderText` in `screen/text.go`;
- the cursor of the Videx cards, `cardVidexVideotern.go` and
  `cardVidexUltraterm.go`;
- the clocks: the No-Slot Clock (`noSlotClockDS1216.go`), on by default on
  the //e models, the ThunderClock Plus (`component/microPD1990ac.go`) and
  the FujiNet clock (`smartPortFujinetClock.go`).

izmac has a `StartTime` in its configuration, and counts the time of the
machine from it, by cycles.

**Here, once fixed:**

- Set the start time in `start` in `activities/setup_test.go`, the same for
  every machine, as izmac's `activityStart`.
- [ultraterm_test.go](activities/ultraterm_test.go) waits for each page of the
  demonstration by its pixels, as the pages did not come at the same frame;
  fixed times may do then.
- [desktop.md](guides/desktop.md) shows the time of the host in the menu
  bar of every picture; it will be the start time.
- Take the note on determinism out of [EDITORIAL.md](EDITORIAL.md), and add a
  check that making the pictures again changes nothing.

### Mockingboard presents Holiday Music plays nothing

The disk of the [Internet Archive](https://archive.org/details/sweet-micro-systems-mockingboard-presents-holiday-music)
(`00playable.dsk`) shows its menu, takes the number of a song, and stops
reading the keyboard as if playing, but no sound source changes its level:
the same with the Mockingboard in slots 1 to 5, on an Apple \]\[+ and on the
enhanced //e. The Sound/Speech I demonstration disk does play on the same
card, so the difference is in what this player asks of it, perhaps the
timers of the 6522 and their interrupts.

**Here, once fixed:** an activity of music on the Mockingboard, with this
disk: its menu, and a song recorded to a WAV file. Add the disk to
`disks.tsv` again.

### The music keys of the Mockingboard demonstration play late

In the *music demonstration* of the Sound/Speech I disk, keys A to K are
notes. Typed by the operator, a key at a time with the latch read in between,
the first notes come out about eight seconds late and together. The program
may wait for the key to be released, which the Apple II tells by bit 7 of
`$C010`, *any key down*, on the //e, and which izapple2's `KeyboardProvider`
has no way to say: it has no key up.

**Here, once fixed:** the operator would need to hold and release keys, as
izmac's `HoldKey` and `ReleaseKey`; then add the music demonstration to
[mockingboard.md](guides/mockingboard.md), a tune recorded.

### The SwyftCard does not see Solid-Apple

The tutorial of the SwyftCard (`-model swyft`) asks to hold Solid-Apple and
tap `=` three times. With button 1 of the game port held, the `=` are typed
as text, and with button 0, or both, the same. Possibly the same cause as the
previous one, the card looking at keys held down.

**Here, once fixed:** an activity of the SwyftCard, Jef Raskin's work before
the Canon Cat, through the first pages of its tutorial.

### Card Cat finds no printer on the parallel card

*Print* in [Card Cat](guides/card-cat.md) says `PRINT: Printer is not
available` with `-s1 parallel`, though Card Cat names the card in slot 1 as
an Apple Parallel Interface Card.

**Here, once fixed:** print the list of the cards from Card Cat, and show the
file the parallel card writes, `printer.out`, in the page.

### The text of the Ultraterm is not read

`ScreenText` gives nothing while the Ultraterm shows its demonstration, so the
operator can't wait for its text, and [ultraterm_test.go](activities/ultraterm_test.go)
counts lit dots instead.

**Here, once fixed:** wait for the text of each page.

### The headless frontend crashes when the machine can't be built

`frontend/headless/main.go` prints the error of `CreateConfiguredApple` and
goes on with a nil machine, which panics on `SetKeyboardProvider`. Seen with
`-showConfig` after a file name: Go stops reading options at the first file,
so the option is taken for one, and the machine fails to build.

**Here:** nothing; `-showConfig` before any file name works, as
[EDITORIAL.md](EDITORIAL.md) says.
