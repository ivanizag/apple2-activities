# Writing activities

How the pages of `guides/` are made, and what makes them good. This is for
whoever writes new activities, people and language models, and for whoever
reviews them. It follows the rules of the activities of
[izmac](https://github.com/ivanizag/izmac), the Macintosh emulator, and adds
what is different about the Apple II.

## What an activity is

An activity is one thing an Apple II owner really did, walked through step by
step, with what the machine showed on the way. Its parts are:

- a page, `guides/<name>.md`;
- an entry in [README.md](README.md), with a title, a picture that links to
  the page, and a paragraph;
- its pictures, in `guides/images/<name>/`, made by running the machine, never
  drawn or edited by hand;
- its generator, `activities/<name>_test.go`, run from `TestActivities`;
- its disks, in `disks.tsv`.

## How to make one

1. **Propose before writing.** List candidate activities with a line each, and
   let the user choose.
2. **Do the activity on the emulator first**, with a throwaway test that runs
   the steps and prints the screen text and takes screenshots. Write the text
   from what the machine does, not from memory of what it did.
3. **Write the generator.** `<name>Screenshots(t *testing.T)` starts its
   machine with `start(t, model, overrides, disks...)`, where the model and the
   overrides are the ones of the command line the page gives, and drives it
   with the operator: `TypeLines`, `Key`, `WaitForText`, `WaitForKeyboard`,
   `TurnPaddle`, `MouseTo`, `Click`. Pictures go into an album:
   `Screenshot`, or a `Record` saved with `SaveRecording`.
4. **Generate the pictures** and **look at every one**, every frame of the
   GIFs included. A contact sheet of the frames is the quickest way.
5. **Commit the generator, the page, the pictures and the disk list
   together.**

## Disks

- **No disk image is committed.** Every disk a page uses is a line of
  `disks.tsv`: its name in `disks/`, its SHA-256 and the URL it is downloaded
  from. `fetch-disks.sh` downloads it, and the generator uses that same file,
  so the reader runs what made the pictures.
- **Where to look:** [Asimov](https://mirrors.apple2.org.za/ftp.apple.asimov.net/images/)
  first, for the masters of Apple's own software and the classics, and the
  [Internet Archive](https://archive.org/), for the WOZ images of the
  [wozaday](https://archive.org/details/wozaday) collection and much else.
  Prefer the original disks to the cracked ones, and WOZ when the original
  needs it.
- **Use what izapple2 already has** when it fits: it carries DOS 3.3, ProDOS
  2.4.3, Apple Pascal 1.3, CP/M, Apple II DeskTop and others inside it, behind
  its models.
- **A blank diskette** is `blank.dsk`, made by the script; the page tells the
  reader to copy it.
- **Check it runs on the model of the page**, a ][+ or a //e, and with the
  cards it has.

## The page

- **Start with a link back**, `[Back to the activities](../README.md)`.
- **Then the history.** One or two paragraphs on what this was in its day and
  why people did it, then one on what the page does.
- **"What you need"** says what to download, from where, and what to rename it
  to: the archive's page linked, the file named as it is there. Say what
  `fetch-disks.sh` does for it. Say when nothing is needed but izapple2.
- **"The machine"** comes after "What you need": the configuration of the
  machine the page uses, as a list, the model and its processor and memory,
  then what is in each slot that matters, and the izapple2 command that starts
  it, in a `bash` block, exactly the one the generator builds. Do not explain
  the options of the command: the list says what the machine is. A page that
  starts the machine twice gives both commands there. `-showConfig`, before
  any file named, prints what a command builds. The reader runs `izapple2`,
  the frontend of the releases.
- **Number the steps once for the whole page**, across its sections.
- **Each step starts with the action in bold**, then what happens, the
  picture, and what it shows and why it mattered.
- **Write for someone who has never seen an Apple II.** Name the keys of the
  Apple II and what they are in izapple2: Reset is Control-F2, the Open Apple
  and Closed Apple keys are the left and right Alt or Option keys.
- **Dates, names and numbers are claims.** Say only what is sure, and check on
  the machine everything that can be checked: sizes, names, what is on the
  screen.
- **Say when a recording is faster than the machine**, and how much.
- **End with "What next"**, linking to another page.

## Pictures

- **Choose the monitor of the time** for the album: green for text and work,
  `album.Green`, and colour, `album.Color`, for graphics and games. A page can
  show both when the difference is the point.
- **The frame is part of the picture**, the black of the tube around the
  screen, from `album.Frame`.
- **Take the picture once the screen is finished**, by what is on it rather
  than by a number of frames: a program prints its last line, its prompt
  appears, its graphics are drawn. Fixed times break when the emulator gets
  faster or slower, as when the Disk II was fixed. See *Driving the machine*
  for which wait to use.
- **Show the result of each important step**, not every step.
- **Use a GIF when the movement is the point**, typing, a program running, a
  game; a PNG when the result is. The first and the last frames must both be
  worth looking at.
- **Real speed by default.** Speed up only waits that would be tedious, with
  `Faster`, and say so in the text, with the real time. Stills are capped at
  `LongestStill`.
- **Determinism is not complete yet**: the flashing cursor follows the time
  of the host, so pictures can differ from one run to the next in the cursor.
- **Look at every picture before committing**, every frame of the GIFs too,
  with `tools/sheet`. A picture taken too early, of a menu half printed or a
  title half drawn, is the most common fault.
- **Sound** goes in WAV files, linked from the page with 🔊, as GitHub does
  not play it inline: `album.Listen`, `Sound.Clip` and `SaveSound`. Keep a
  clip to what the step needs, 15 seconds is 700 KB.
- **Keep GIFs small**: the recorder stores what changed and leaves the rest
  transparent, but a screen that changes everywhere at every frame, as the
  self test of the //e, makes a huge one; use stills for those.

## Driving the machine

The operator types, turns paddles, presses buttons and moves the mouse, and
waits for what the machine does. What it has taught:

- **Which wait.**
  - `WaitForText` is true as soon as the text is in the text memory, which
    may not be what is shown, or may be there from before: a prompt `]` is
    on the screen while a command still works, and `APPLE ][` has a `]` too.
  - `WaitForShownText` also needs the screen to be in text mode: a program
    can prepare its menu in the text page while its title is shown in
    graphics.
  - `WaitForKeyboard` waits until the program looks for a key. CP/M looks for
    one while it prints, to stop, and a key typed then aborts the command.
  - `WaitForPrompt` waits until the screen is still and the keyboard is read.
    It never comes where the cursor blinks in the text memory itself, as in
    ProDOS BASIC: use `WaitForText` and `WaitForKeyboard` there.
  - For graphics, wait by the screen: the video mode, with `waitForMode`, and
    what is drawn, counting dots, as Total Replay and the Ultraterm do. Look
    at the screen every 30 frames, not every frame: rendering it is slow.
- **Keys.** `Type` waits for each key to be read before the next, as a
  person watching the screen; the Apple \]\[ and \]\[+ get capitals. Keys
  are pressed, not held: izapple2 has no key up, so programs that wait for a
  key to be released do not work (see [IZAPPLE2.md](IZAPPLE2.md)).
- **The Apple keys** of the //e are the buttons of the game port: Open Apple
  is button 0 and Closed Apple, or Solid Apple, button 1, `HoldButton`. Held
  through `Reset` they start the self test.
- **Reset** is Control-F2 in izapple2, `o.Reset()` for the operator. Applesoft
  stops on Control-C between statements by itself, any program of its own.
- **The Monitor** of the Apple \]\[: Escape and `@` clear the screen, and a
  Return on an empty line then would dump memory.
- **Sound** has to be listened to from the start, `album.Listen` before the
  machine runs: the Mockingboard works out its sound only while listened to,
  and when listened to late it has all that time to catch up with first.
- **The mouse** of the mouse card takes a place as a part of the range the
  program set, so a place on the screen is a fraction of it: `atDot` in the
  DeskTop generator for its 560 by 192 dots.
- **Paddles** for a game written in BASIC: the operator can find the ball in
  the low resolution page with `Peek`, as `loResFind` does, and turn the
  paddle under it.
- **Nothing on the machine is checked by memory.** Explore first; then check
  each claim of a page on the machine: what a key does, what a screen says,
  how long it takes.

## Markdown

- **Escape the brackets of Apple \]\[** outside code, `Apple \]\[+`, or they
  break links and pictures. `gh api markdown -f text="$(cat page.md)" -f
  mode=gfm` renders a page as GitHub will.

