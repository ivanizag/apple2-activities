# Open findings in izapple2

What making the activities found in [izapple2](https://github.com/ivanizag/izapple2),
the state of each, and what to change here when it is fixed. When one is done
there and done here, take it out of this list.

## Fixed on master, not released

### INIT of a blank diskette hangs the emulator

`nibDecodeTrack` in `storage/fileNib.go` went around the track forever when
an address prolog started on its last bytes, which a track being written can
have. Fixed on master, with a test, but the releases of izapple2 before the
fix still have it: a reader following [Life with DOS 3.3](guides/dos33.md)
with one sees the machine stop at `INIT HELLO,D2`.

**Here, once released:** say in [dos33.md](guides/dos33.md) which release of
izapple2 the page needs, at least.

## Not fixed

### No -board on the command line

The pages give the whole machine on the command line, starting from the
model `_base`, which has nothing, as `izapple2 -model _base -board 2plus
-cpu 6502 -rom ...`. The board can only come from a model today: izapple2 has
the option for each of the rest, but none for `board`. The library takes it
in the configuration, and that is how the generators build the machines, so
the pictures are right; a reader's izapple2 refuses the command.

**Here, once done:** say in [README.md](README.md) which release of izapple2
the commands of the pages need, at least. Until then, the preconfigured model
each page gives after its command runs the activity.

### Writing to a WOZ disk stops the emulator

When a program writes to a disk in WOZ format, izapple2 panics with `Write
not implemented on woz disk`, from `disketteWoz.Write` in `storage/`. Apple
Pascal writes the work file to its disk, so it cannot run from the original
disks of the [woz-a-day collection](https://archive.org/details/wozaday_Apple_Pascal_v13).

**Here, once fixed:** [pascal.md](guides/pascal.md) uses the disks inside
izapple2; it could use the woz-a-day ones, downloaded, but see the next one.

### Apple Pascal 1.3 from its original disks stops at the start with four drives

With `APPLE1` and `APPLE2` in slot 6 and `APPLE3` and `APPLE0` in slot 5, as
the `pascal` model has them, the woz-a-day disks, or the DSK images of the
Internet Archive (`211_Apple_II_Pascal_1.3_Apple0` to `214`), stop with a
blank screen after `Apple //e`, the same on every run. The woz-a-day
`APPLE1` in slot 6 and `APPLE3` alone in drive 2 of slot 5 are enough for
it, and in drive 1 of slot 5 it starts. The disks inside izapple2, which have
been used (their date is 10-Apr-90), start in all four drives. Not known
whether a real //e does the same.

**Here, once understood:** [pascal.md](guides/pascal.md) uses two drives; it
would only need four for the programs on `APPLE3` and `APPLE0`.

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

### CP/M cannot format a blank diskette

`FORMAT` of the CP/M 2.20B diskette of the SoftCard, asked for drive B with a
blank `.dsk` there, all zeros, answers `DISK I/O ERROR` after `CONTINUE
(Y/N)? Y`. DOS 3.3 initializes the same blank diskette with `INIT`.

**Here, once fixed:** [cpm.md](guides/cpm.md) could add a second drive, a
diskette formatted by CP/M, and files copied to it with `PIP`.

### The headless frontend crashes when the machine can't be built

`frontend/headless/main.go` prints the error of `CreateConfiguredApple` and
goes on with a nil machine, which panics on `SetKeyboardProvider`. Seen with
`-showConfig` after a file name: Go stops reading options at the first file,
so the option is taken for one, and the machine fails to build.

**Here:** nothing; `-showConfig` before any file name works, as
[EDITORIAL.md](EDITORIAL.md) says.
