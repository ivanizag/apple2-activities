# The Apple \]\[ and its cassettes

[Activities](../README.md) › [First steps](../sections/first-steps.md)

The first Apple \]\[s went on sale on June 10, 1977: a board of Steve
Wozniak's in a plastic case, with its keyboard, two paddles, 4 KB of memory,
and Integer BASIC in its ROM, ready when switched on. There was no disk
drive. To keep a program, or to buy one, there was the cassette: a jack on
the back of the machine for the earphone of an ordinary tape recorder, and
another for its microphone. The recorder was not in the box, nor was the
television; the programs Apple sold were on cassettes, and the first ones
showed what the machine could do: a game, and its colours.

Getting a program in took a minute and some patience: the volume and the
tone of the recorder had to be right, or the Apple answered `ERR`.

This page loads three of Apple's tapes into a first Apple \]\[, plays
Breakout, and runs the demonstrations of the colour graphics and of the high
resolution graphics.

## What you need

The recordings of three of Apple's cassettes, from the
[Apple cassettes](https://www.brutaldeluxe.fr/projects/cassettes/apple/index.html)
of the Brutal Deluxe cassette project, which keeps the tapes of the Apple \]\[
as sound files:

- **`k7_apple_002000101_breakout.wav`** and
  **`k7_apple_002000101_colorgraphics.wav`**, the two programs of the tape
  002-0001-01, *Breakout / Color Graphics*;
- **`k7_apple_600201600_highresolutiongraphics.wav`**, the tape 600-2016-00,
  *High-Resolution Graphics*.

Each comes in a `.zip` of the same name. `./fetch-disks.sh` in this repository
downloads them into `disks/`, takes the recordings out and checks them.

Apple's manual of its tapes for the Apple \]\[,
[The Apple Tapes](https://www.brutaldeluxe.fr/projects/cassettes/apple/apple_theappletapes_appleii_manual.pdf),
of 1979, is on Brutal Deluxe.

## The machine

The first Apple \]\[, as it came:

- the 6502 processor at 1 MHz and 48 KB of memory;
- the ROM with Integer BASIC and the Monitor;
- the board of Revision 0, with four colours in the high resolution graphics,
  and a keyboard of capitals only;
- the cassette input, with a tape in the recorder;
- paddles in the game port, and no cards in its slots.

The machine with the tape of Breakout:

```bash
izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/341-000x_integer.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -mods four-colors \
    -tape disks/k7_apple_002000101_breakout.wav
```

With the tape of Color Graphics:

```bash
izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/341-000x_integer.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -mods four-colors \
    -tape disks/k7_apple_002000101_colorgraphics.wav
```

With the tape of High-Resolution Graphics:

```bash
izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/341-000x_integer.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -mods four-colors \
    -tape disks/k7_apple_600201600_highresolutiongraphics.wav
```

izapple2 plays the tape when the Apple starts listening to the cassette
input, and stops it when the Apple stops: there is no Play button to press,
and it runs the machine as fast as it can while the tape plays, so a program
that took half a minute to load comes in a moment.

## Breakout

1. **Start izapple2** with the first command. The Apple \]\[ starts in the
   Monitor, with its prompt, `*`. **Press Control-B and Return** for Integer
   BASIC, with its prompt, `>`.

2. **Type `LOAD`.** The Apple listens to the cassette: a tone first, to set
   the level, then the program, as tones of two pitches, one for each value
   of a bit. When it has all of it, the prompt comes back. **Type `RUN`.**

   ![Breakout asks for a name](images/apple-ii-cassettes/breakout-name.png)

   Breakout was a game of Atari, of 1976, which Wozniak had designed. Here
   it is in Integer BASIC, with the bricks in columns at the right and the
   bat at the left.

3. **Type a name**, and **`Y`** for the standard colours. **Move the mouse
   left and right** over the izapple2 window to move the bat up and down:
   without a joystick, izapple2 makes the mouse paddle 0, from 0 at the left
   of the window to 255 at the right.

   ![Breakout played](images/apple-ii-cassettes/breakout.gif)

   The bricks of the column at the right count more than the ones at the
   left. This is the speed of the machine: Integer BASIC works out every
   line of the game, the ball, the bat, the bricks and the score.

## Color Graphics

4. **Quit izapple2 and start it with the second command.** Control-B,
   Return, `LOAD` and `RUN`, as before. The menu of the demonstrations:

   ![The demonstrations of the colour graphics](images/apple-ii-cassettes/color-menu.png)

5. **Type `3` and Return**, the kaleidoscope: the colours of the low
   resolution graphics, drawn in patterns that mirror each other, as in a
   kaleidoscope.

   ![The kaleidoscope](images/apple-ii-cassettes/kaleidoscope.gif)

   The colour of the Apple \]\[ was the point: the Commodore PET and the
   TRS-80 of the same year had none. The Apple made its colours from the way a colour television decodes
   its signal, with very few chips.

## High-Resolution Graphics

6. **Quit izapple2 and start it with the third command.** This tape has
   two parts: routines in machine code, that draw in the high resolution
   graphics, and a program in Integer BASIC that calls them. The machine
   code comes first, and the Monitor reads it: **type `C00.FFFR`**, *read
   the tape into memory from `C00` to `FFF`*. Then **Control-B and Return**,
   and **`LOAD`** for the BASIC that comes next on the tape, and **`RUN`**.

   ![The demonstrations of the high resolution graphics](images/apple-ii-cassettes/hires-menu.png)

7. **Type `6` and Return**, the donut, rings drawn one after the other: a
   minute of the machine.

   ![The donut](images/apple-ii-cassettes/donut.gif)

   Each dot of the high resolution graphics is a bit of memory, 280 across
   and 192 down. All of them are violet and green, as everything on a board
   of Revision 0: [From four colours to six](six-colours.md) is how the later
   boards got two more.

## What next

**Next:** [From four colours to six](six-colours.md), the colours of the high
resolution graphics on the first board and on the later ones.

[The Apple \]\[ of 1977](apple-ii.md) goes through the Monitor, the
Mini-Assembler and Integer BASIC of the same machine, and
[A game in Integer BASIC](paddle-game.md) is a game of bat and ball typed
into the same Apple \]\[.
