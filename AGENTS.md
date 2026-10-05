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

## Build and run

```bash
./fetch-disks.sh
go vet ./...
A2_ACTIVITIES=1 go test -count=1 -run 'TestActivities/dos33' ./activities
```

Without `A2_ACTIVITIES` the generators skip themselves. They write over the
pictures of the pages they run.

## Working with izapple2

The generators need izapple2 changes that may not be released yet. To build
against a local copy, next to this one, make a workspace, which is not
committed:

```bash
go work init . ../izapple2
```

What the activities need from the emulator goes into izapple2 itself: the
library calls (`ScreenText`, `Peek`, `LoadDisk` and the rest) and its fixes.
What only the pages need stays here. When izapple2 has the changes on its main
branch, update `go.mod` with `GOWORK=off go get github.com/ivanizag/izapple2@master`.

## Code style

The izapple2 conventions: `gofmt`, doc comments on everything exported starting
with its name, errors returned and wrapped with `%w`, hardware addresses in hex.
Comments say what the code does in the terms of the machine and of the person
at it.
