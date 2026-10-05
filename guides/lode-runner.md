# Lode Runner

[Back to the activities](../README.md)

Lode Runner was Doug Smith's game, published by Broderbund in 1983 for the
Apple II before any other computer. A runner collects the gold of a level of
bricks, ladders and ropes while guards chase him; he can't jump or fight, only
dig a hole in the bricks for a guard to fall into, and the hole fills again
after a while. It came with 150 levels and an editor to make more, and it
went on to most computers and consoles of the eighties.

This page starts Lode Runner from a copy of its original disk, watches the
demonstration that plays itself, and starts a game.

## What you need

**`lode-runner.woz`**, a copy of the original Lode Runner disk, from the
[woz-a-day collection](https://archive.org/details/wozaday_Lode_Runner) of
the Internet Archive: download *00playable.woz* and rename it
`lode-runner.woz`. `./fetch-disks.sh` in this repository downloads it into
`disks/`.

A WOZ file records a diskette bit by bit, its copy protection included, so
it is the disk Broderbund sold, as it was, and not a copy with the protection
taken out.

## The machine

An Apple \]\[+ with one disk drive:

- the 6502 processor at 1 MHz, 48 KB of memory, and Applesoft BASIC in its ROM;
- a 16 KB Language Card in slot 0;
- a Videx Videoterm 80 column card in slot 3, unused here;
- a Disk II controller card in slot 6, with the Lode Runner disk in drive 1.

```bash
izapple2 -model 2plus disks/lode-runner.woz
```

## Start it

1. **Start izapple2** with the command above, and press F6 until the screen
   is in colour. After a few seconds of the drive, the title comes up.

   ![The title](images/lode-runner/title.png)

## The demonstration

2. **Wait.** The game shows its first level, opening from the middle, and
   plays it by itself: the runner in white, the guards in orange, the gold
   in the small boxes.

   ![The demonstration](images/lode-runner/demo.gif)

   The runner goes for the gold and away from the guards, and digs holes in
   front of them. When all the gold is taken, a ladder appears to the top of
   the screen, the way out to the next level.

## Play

3. **Press Space** to play. The first level comes up, with five runners,
   `MEN 005` at the bottom.

   ![A game started](images/lode-runner/game.png)

   Lode Runner is played with a joystick, or with the keyboard; its manual
   has the keys. Without a joystick connected, izapple2 makes the mouse of
   your computer the joystick of the Apple II.

## What next

[A game in Applesoft](paddle-game.md) is a game you write yourself, in
BASIC, for the same machine.
