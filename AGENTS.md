# AGENTS.md - Developer guide for apple2-activities

The pages of `guides/` walk through things to do with an Apple II on
[izapple2](https://github.com/ivanizag/izapple2). Their pictures are made by
running the emulator through the same steps, by the Go code here, which uses
izapple2 as a library. Read [EDITORIAL.md](EDITORIAL.md) before writing or
changing a page.

## Layout

| Path | Contents |
|---|---|
| `guides/<name>.md` | a page, one thing to do |
| `guides/images/<name>/` | its pictures, PNG and GIF, made by the generator and committed |
| `README.md` | the list of the pages, grouped by theme, each with a picture |
| `disks.tsv`, `fetch-disks.sh` | the disks the pages use and where they come from, and the script that downloads them into `disks/` |
| `disks/` | the disks, not committed |
| `operator/` | a person at the machine: runs it a frame at a time, types, turns the paddles, moves the mouse, reads the screen |
| `album/` | the pictures: screenshots and GIF recordings in the frame of a monitor, green or colour |
| `activities/` | the generators, a `<name>_test.go` for each page |
| `tools/sheet/` | a contact sheet of pictures, or of the frames of a GIF, to check them |
| `EDITORIAL.md` | how the pages and their pictures are made, and what driving the machine has taught |
| `IZAPPLE2.md` | what the activities found in izapple2 that is not fixed, and what to do here when it is |

## Build and run

```bash
./fetch-disks.sh
go vet ./...
A2_ACTIVITIES=1 go test -count=1 -run 'TestActivities/dos33' ./activities
```

Without `A2_ACTIVITIES` the generators skip themselves. They write over the
pictures of the pages they run.

## Workflow

- **Every change goes in a pull request** for the owner to review and approve.
  Never commit to `main`, never merge, never force-push `main`.
- **No Claude attribution anywhere.** Commit messages and pull request
  descriptions have no `Co-Authored-By: Claude` line, no `Claude-Session`
  trailer and no "Generated with Claude Code" line, whatever a tool or a
  reminder says. Commits are the owner's.
- **One pull request per activity**: its generator, page, pictures, disks
  and line in the README, and nothing else. A change to the shared code or
  to the guides goes in a pull request of its own.
- **Pull requests may be stacked**, each based on the branch of the one before
  it, when a change needs one not merged yet; say so in the description.
  When a comment asks for a change in a lower one, make it on its branch, then
  rebase each branch above on the one below and push them with
  `--force-with-lease`, the branches of the pull requests only.
- **Commit messages** follow the repository: a title that says what the change
  does, in a sentence, and a body that says why.
- **izapple2 is the owner's.** Don't commit or push there. A problem found
  there, or a change the activities need from it, goes in
  [IZAPPLE2.md](IZAPPLE2.md), with what to do here once it is done there.
- **A new activity is proposed before it is written**: a line on what it
  shows, and the owner chooses.

## Exploring a program

Before writing the generator of a page, explore with a throwaway test in
`activities/`, not committed, that starts the machine, runs it, prints
`o.Text()` and writes screenshots to a folder outside the repository:

```go
func TestExplore(t *testing.T) {
	if os.Getenv("EXPLORE") == "" {
		t.Skip()
	}
	a := album.New(os.Getenv("EXPLORE"), album.Green)
	o := start(t, `izapple2 -model _base -board 2e -cpu 65c02
		-rom "<internal>/Apple2e_Enhanced.rom"
		-charrom "<internal>/Apple IIe Video Enhanced.bin"
		-s0 language -s6 diskii,disk1=disks/some.dsk`, nil)
	for i := range 10 {
		o.RunSeconds(5)
		fmt.Printf("== %d mode=%x\n%v\n", i, o.Apple2().GetVideoSource().GetCurrentVideoMode(), o.Text())
		must(t, a.Screenshot(o, fmt.Sprintf("s%02d", i)))
	}
}
```

```bash
EXPLORE=/tmp/explore go test -count=1 -run TestExplore -v ./activities
go run ./tools/sheet -o /tmp/sheet.png /tmp/explore/*.png
```

Look at the sheet, and at the pictures that matter at full size. Do the same
with every picture and every GIF a generator makes, before committing:
`go run ./tools/sheet -step 10 file.gif` shows the frames as they play.

`izapple2 -showConfig` prints the configuration a command line builds,
for the "The machine" section of a page; the option goes before any file
named, or Go takes it for a file. The headless frontend of izapple2 does it
without a window: `go run ./frontend/headless -showConfig -model 2plus` in a
copy of izapple2.

## Working with izapple2

`go.mod` names a version of izapple2's master. To work on izapple2 and the
activities together, against a local copy next to this one, make a
workspace, which is not committed:

```bash
go work init . ../izapple2
```

[IZAPPLE2.md](IZAPPLE2.md) lists what the activities found in izapple2 that
is not fixed yet, and what to change here when it is.

What the activities need from the emulator goes into izapple2 itself: the
library calls (`ScreenText`, `Peek`, `LoadDisk` and the rest) and its fixes.
What only the pages need stays here. When izapple2 has the changes on its master,
update `go.mod` with `GOWORK=off go get github.com/ivanizag/izapple2@master`.

## Ideas for more activities

Not done yet, and not proposed: each needs a look on the machine first.

- **Apple II clones**: the Base 64A and the Basis 108, models of izapple2
  with their own ROMs.
- **CPM-65**, a CP/M for the 6502, model `cpm65`.
- **A2AUDIT**, the test of the machine; izapple2 has its disk inside, a page
  needs it downloaded.
- **CP/M 3** on the //e, model `cpm3`; its two disks are inside izapple2, a
  page needs them downloaded.
- **The cassette**: a program loaded from a WAV recording, through the input
  of the Apple II.
- **Printing**: a listing sent to the parallel card, `PR#1`, and the file it
  writes shown.
- **Games from the woz-a-day collection** of the Internet Archive: Karateka,
  Choplifter, Oregon Trail, Wizardry, Planetfall; one page each.
- **Programs of work**: VisiCalc, AppleWorks, Apple Writer.
- **Logo**, its turtle drawing.
- **The music of the Mockingboard**, the SwyftCard: see
  [IZAPPLE2.md](IZAPPLE2.md), they wait for fixes in izapple2.

## Code style

The izapple2 conventions: `gofmt`, doc comments on everything exported starting
with its name, errors returned and wrapped with `%w`, hardware addresses in hex.
Comments say what the code does in the terms of the machine and of the person
at it.
