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

### No -board on the command line

The pages give the whole machine on the command line, as `izapple2 -model
none -board 2plus -cpu 6502 -screen green -rom ...`. izapple2 had no
`-board` option, no model `none` and no `-screen`, so a reader's izapple2
refused the commands of the pages. Master has them now, but the releases of
izapple2 don't yet.

**Here, once released:** say in [README.md](README.md) which release of
izapple2 the commands of the pages need, at least. Until then, the
preconfigured model each page gives after its command runs the activity.

## Not fixed

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

The keys of [Karateka](guides/karateka.md) may be another case: `Q`, `A`,
`Z`, `W`, `S` and `X`, which its manual gives to punch and kick from the
fighting stance, made no difference to the fight, which the joystick plays.
Once keys can be held, try them, and say so on the page.

### Calc of the SwyftCard stays in Applesoft

The command *Calc* of the SwyftCard, Use Front `G`, Control-G, gives the
highlighted line to Applesoft and should put the answer in the Text. On page
61 of the tutorial, with `? 5.6 + 3` typed, Applesoft prints `8.6` over the
80 column page and stays at its `]` prompt: the hooks of output and input,
`$36` to `$39`, are the ones of the ROM, `$FDF0` and `$FD1B`, and the
processor waits for a key in the firmware of the 80 column card, at `$C83D`,
so nothing takes the answer back to the card.

**Here, once fixed:** add Calc to [swyftcard.md](guides/swyftcard.md), the
sum of the tutorial worked out in the Text.

### The SwyftCard can't save its Text

The command *Disk* of the SwyftCard, Use Front `L`, saves the Text only on
a diskette that is blank, never formatted, and in a format of its own.
izapple2 has no such diskette that can be written: a `.dsk`, even of zeros,
is a formatted DOS diskette to the card, which refuses it as not blank, and
writes to it would be decoded as the 16 sectors of DOS and ProDOS
(`saveTrack` in `storage/fileNib.go`); a `.nib` is never written
(`newFileNib` does not allow it); and a `.woz` can't be written either (see
*Writing to a WOZ disk stops the emulator*). An all-zero `.nib` makes the
card wait forever at `$D745`, as a real drive gives noise there and not
nothing.

**Here, once fixed:** in [swyftcard.md](guides/swyftcard.md), save the
letter on a blank diskette with Disk, and switch on again with it in the
drive to see it come back.

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

### The shift mod does nothing

`-mods shift` is in the help of izapple2, but `setupShiftedKeyboard` is
commented out in `setup.go`: the mod is not there. On a real Apple \]\[+ it
wired the Shift key to the input of button 2, `$C063`, and word processors
in 80 columns, as Apple Writer II with a Videx Videoterm, read it there to
type capitals with Shift. The keyboard of izapple2 would have to tell when
Shift is held, as it does not now.

**Here, once done:** an activity of the shift mod, an Apple \]\[+ with a
Videoterm and a word processor of the time typing capitals and small
letters with Shift.

### A program that runs into an undefined opcode stops izapple2

On a 6502, an opcode that is not defined does something, or hangs the
processor; in izapple2 it panics, and the whole emulator stops: `panic:
Unknown opcode 0x0b`, from `iz6502` `ExecuteInstruction`. It came up when
`RUN` followed a `LOAD` that had failed, so running garbage, which on a real
machine would hang or crash it, not the computer it is emulated on.

**Here, once fixed:** nothing to change.

### The text of the Basis 108 in 80 columns is read a column out of two

In 80 columns the Basis 108 keeps its text in a memory of its own, columns
in an order of its own (`VideoText80AltOrder`). `ScreenText` gives one
column out of two of it: `PRINT "Hello"` shows on the screen as it is, and
comes back as `RN Hlo` and `Hlo`. So the operator can't wait for a text, and
[basis108_test.go](activities/basis108_test.go) waits for the keyboard to
be read and the screen to be still instead.

**Here, once fixed:** wait for the text of each step in basis108_test.go.

### CP/M cannot format a blank diskette

`FORMAT` of the CP/M 2.20B diskette of the SoftCard, asked for drive B with a
blank `.dsk` there, all zeros, answers `DISK I/O ERROR` after `CONTINUE
(Y/N)? Y`. DOS 3.3 initializes the same blank diskette with `INIT`.

**Here, once fixed:** [cpm.md](guides/cpm.md) could add a second drive, a
diskette formatted by CP/M, and files copied to it with `PIP`.

### The 132 columns of the Ultraterm are shown 160 wide

In the mode 7 of the Ultraterm, 132 columns by 24 lines, interlaced, the
firmware writes 144 to the register 1 of the 6845, the characters across,
and `MC6845.Write` in `component/mc6845.go` turns it into 160, a hack for
that mode. The lines of the firmware then run on across the screen instead
of starting at the left.

**Here, once fixed:** take the picture of mode 7 in
[ultraterm_test.go](activities/ultraterm_test.go), and show it in
[ultraterm.md](guides/ultraterm.md) with the others.

### The headless frontend crashes when the machine can't be built

`frontend/headless/main.go` prints the error of `CreateConfiguredApple` and
goes on with a nil machine, which panics on `SetKeyboardProvider`. Seen with
`-showConfig` after a file name: Go stops reading options at the first file,
so the option is taken for one, and the machine fails to build.

**Here:** nothing; `-showConfig` before any file name works, as
[EDITORIAL.md](EDITORIAL.md) says.
