# Things to do with an Apple II

[izapple2](https://github.com/ivanizag/izapple2) emulates the Apple \]\[+ and
the Apple //e and runs their real ROMs and software, so it is a way of finding
out first hand what using one was like. Each page here takes one thing an
Apple II owner did, and walks through it step by step, with the command to
start izapple2 for it and what you will see on the way.

Each page says what it needs and where it comes from. Some need nothing but
izapple2; the rest use disks from public archives, which
[`fetch-disks.sh`](https://github.com/ivanizag/apple2-activities/blob/main/fetch-disks.sh) downloads into the folder `disks`. Each
page gives the whole machine it uses, card by card, as a command line of
izapple2, and after it a shorter one with a model of izapple2 that has that
machine, and often more. The whole command needs the option `-board`, which
izapple2 does not have yet; the shorter one works with izapple2 as it is.

## First steps

### [The Apple \]\[ of 1977](guides/apple-ii.md)

[<img src="guides/images/apple-ii/colours.png" width="320" alt="The sixteen colours of Integer BASIC">](guides/apple-ii.md)

Wozniak's machine and its ROM: the Monitor, machine code written a line at a
time in the Mini-Assembler, and Integer BASIC, with its whole numbers and its
sixteen colours.

### [Switch on an Apple \]\[+](guides/switch-on.md)

[<img src="guides/images/switch-on/program.png" width="320" alt="A program in Applesoft BASIC">](guides/switch-on.md)

No disk, no operating system: switched on, the Apple \]\[+ is in Applesoft BASIC
and waiting. A few commands, a program typed, listed and run, and a loop that
never ends stopped with Control-C.

### [Switch on an Apple //e](guides/apple-iie.md)

[<img src="guides/images/apple-iie/eighty-columns.png" width="320" alt="80 columns and MouseText">](guides/apple-iie.md)

The Apple II of the rest of the decade: lower case, 80 columns, the MouseText
of the enhanced //e, and the self test in its ROM, run with both Apple keys.

### [Life with DOS 3.3](guides/dos33.md)

[<img src="guides/images/dos33/catalog.png" width="320" alt="The catalog of the System Master">](guides/dos33.md)

Two Disk II drives and the System Master: the catalog of a diskette, the
program that greets you, and a blank diskette made into one that starts the
machine with a program of your own.

### [ProDOS](guides/prodos.md)

[<img src="guides/images/prodos/disk-map.png" width="320" alt="The map of a diskette in Copy II Plus">](guides/prodos.md)

The operating system of the Apple II from 1984 on, in its version of 2023:
the program selector it starts with, BASIC and a catalog in 80 columns, a
folder and a text file on the RAM disk of the //e, and the diskette mapped
file by file in Copy II Plus.

### [Apple II DeskTop](guides/desktop.md)

[<img src="guides/images/desktop/about.png" width="320" alt="About This Apple II">](guides/desktop.md)

The Apple II with a mouse, windows and icons: a disk opened with a double
click, a file read, the machine asked what it has in its slots, a picture in
double high resolution, the flying toasters, and the Calculator.

## Making things

### [A game in Applesoft](guides/paddle-game.md)

[<img src="guides/images/paddle-game/game-over.png" width="320" alt="A bat and a ball in low resolution graphics">](guides/paddle-game.md)

A bat on a paddle and a ball, in the low resolution colour graphics of the
Apple \]\[+: typed in as the listings of the magazines were, and played with the
mouse as the paddle.

### [Logo and its turtle](guides/logo.md)

[<img src="guides/images/logo/flower.gif" width="320" alt="A flower of squares drawn by the turtle">](guides/logo.md)

Apple Logo on an Apple \]\[+: words and lists, the turtle moved by hand,
and procedures that teach it new words, a flower of squares and a spiral
that calls itself, saved on the diskette.

### [Printing](guides/printing.md)

[<img src="guides/images/printing/printing.png" width="320" alt="A calendar printed">](guides/printing.md)

A program typed into an Apple \]\[+ with a parallel printer card: its
listing printed with `PR#1`, and a calendar of 1977 three months across,
wider than the screen, as it comes out on paper.

### [Apple Pascal](guides/pascal.md)

[<img src="guides/images/pascal/editor.png" width="320" alt="A program in the editor of Apple Pascal">](guides/pascal.md)

The UCSD p-System on an Apple //e with two drives: the Filer, the editor, the
compiler, and a program that draws with the turtle.

### [A game in Apple Pascal: 2048](guides/pascal-2048.md)

[<img src="guides/images/pascal-2048/end.png" width="320" alt="2048 in Apple Pascal">](guides/pascal-2048.md)

A whole game written in Apple Pascal, its listing in this repository: typed
into the editor, compiled, and played to the end, 200 moves.

### [A game in Applesoft: the snake](guides/applesoft-snake.md)

[<img src="guides/images/applesoft-snake/long.png" width="320" alt="The snake in low resolution">](guides/applesoft-snake.md)

A whole game of 35 lines of Applesoft, its listing in this repository: the
snake in the low resolution graphics, typed in, explained, and played.

### [Assembly language with Merlin](guides/merlin.md)

[<img src="guides/images/merlin/assembled.png" width="320" alt="A program assembled by Merlin">](guides/merlin.md)

A program in the language of the 6502, its source in this repository: typed
into the Merlin assembler of 1983, assembled into 120 bytes, and run, bars
of colour moving many times a second.

## Work

### [VisiCalc](guides/visicalc.md)

[<img src="guides/images/visicalc/share.png" width="320" alt="A budget in VisiCalc">](guides/visicalc.md)

The first spreadsheet, from its original disk of 13 sectors: a household
budget typed in, totalled, a formula replicated down a column, and the rent
changed to see the whole sheet worked out again.

### [AppleWorks](guides/appleworks.md)

[<img src="guides/images/appleworks/desktop.png" width="320" alt="The Desktop of AppleWorks">](guides/appleworks.md)

The word processor, spreadsheet and data base of the //e in one: the costs
of a trip worked out, a letter with that table in it, a list of friends,
and the three saved.

### [The SwyftCard](guides/swyftcard.md)

[<img src="guides/images/swyftcard/moved.png" width="320" alt="A letter edited on the SwyftCard">](guides/swyftcard.md)

Jef Raskin's editor on a card for the //e, before the Canon Cat: its
tutorial, the cursor leaping to what you type with the Apple keys, and a
letter written from nothing with a sentence moved by leaping.

## Cards

### [The Mockingboard](guides/mockingboard.md)

[<img src="guides/images/mockingboard/title.png" width="320" alt="The Mockingboard demonstration">](guides/mockingboard.md)

The sound card of the Apple II, from its demonstration disk of 1982: its menus,
and its sound effects, recorded to listen to.

### [160 columns: the Videx Ultraterm](guides/ultraterm.md)

[<img src="guides/images/ultraterm/mode-3.png" width="320" alt="160 columns by 24 lines">](guides/ultraterm.md)

The card that gave the Apple \]\[+ up to 160 columns of text: the
demonstration of its own disk of utilities, and a program that shows each
of its modes, up to 160 by 24 and 80 by 48.

### [The RGB card and its fourteen video modes](guides/rgb-card.md)

[<img src="guides/images/rgb-card/mode-9.png" width="320" alt="Double low resolution on the RGB card">](guides/rgb-card.md)

The demonstration disk of the Video-7 RGB card for the //e: each of the
fourteen video modes it lists, the six of the //e and the eight the card
adds, text in sixteen colours, 160 and 560 dots across.

### [A megabyte on a card: the Memory Expansion Card](guides/memory-expansion.md)

[<img src="guides/images/memory-expansion/copied.png" width="320" alt="A file copied to the RAM disk of the card">](guides/memory-expansion.md)

Apple's memory card for the Apple \]\[, \]\[+ and //e, a megabyte on an
Apple \]\[+: the RAM disk ProDOS finds on it, a file loaded from it in a
sixth of the time, the test in its ROM, and the card from DOS 3.3.

### [What is in the slots: Card Cat](guides/card-cat.md)

[<img src="guides/images/card-cat/slots.png" width="320" alt="Card Cat">](guides/card-cat.md)

A modern program that finds the card in each slot, run on a //e with all seven
full, and the ROM of a card read by it.

## Other systems

### [Forth in ROM](guides/forth.md)

[<img src="guides/images/forth/bars.gif" width="320" alt="The sixteen colours drawn by Forth">](guides/forth.md)

An Apple \]\[+ that starts in FORTH-79, from a card of Offete Industries:
words used and new ones defined, loops and decisions, and the machine
underneath, its memory, its speaker and its graphics, driven from Forth.


### [CP/M on the Z80 SoftCard](guides/cpm.md)

[<img src="guides/images/cpm/dir.png" width="320" alt="The disk of CP/M">](guides/cpm.md)

A Z80 on a card turns an Apple \]\[+ into a CP/M computer: its disk, an
assembler source, a program in Microsoft BASIC-80 saved among its files, and
the high resolution graphics drawn from GBASIC.

## Play

### [Karateka](guides/karateka.md)

[<img src="guides/images/karateka/castle.png" width="320" alt="The fortress of Akuma">](guides/karateka.md)

Jordan Mechner's game of 1984, from its original disk: its titles and
prologue told like a film, Akuma and the princess, and the karateka up the
cliff and past the first guard, played with the joystick.

### [Lode Runner](guides/lode-runner.md)

[<img src="guides/images/lode-runner/title.png" width="320" alt="Lode Runner">](guides/lode-runner.md)

Broderbund's game of 1983, from a copy of its original disk: the title, the
demonstration that plays itself, and a game played from the keyboard, until
a guard catches the runner.

### [Total Replay](guides/total-replay.md)

[<img src="guides/images/total-replay/box-art.png" width="320" alt="Box art in Super Hi-Res">](guides/total-replay.md)

Hundreds of games on one hard disk, with the box art in Super Hi-Res: the
launcher, its attract mode, and a game found by typing its name.

## The disks

```bash
./fetch-disks.sh
```

downloads every disk the pages use into `disks/`, from the archives
[`disks.tsv`](https://github.com/ivanizag/apple2-activities/blob/main/disks.tsv) lists, and checks each against its SHA-256. It only
downloads what is missing or has changed, so it can be run again at any time.
`./fetch-disks.sh CPM1.PO` gets only the disks named.

## How the pictures are made

Every picture is made by running the machine through the steps of its page,
with the same disks, by the Go programs in
[`activities`](https://github.com/ivanizag/apple2-activities/tree/main/activities). See
[AGENTS.md](https://github.com/ivanizag/apple2-activities/blob/main/AGENTS.md) for how, and
[EDITORIAL.md](https://github.com/ivanizag/apple2-activities/blob/main/EDITORIAL.md) for how the pages are written.
