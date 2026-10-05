# Apple II DeskTop

[Back to the activities](../README.md)

In 1984 the Macintosh showed what a mouse, windows and icons were for, and the
Apple II got them too: Apple sold a mouse for the //e and the //c, and
programs to use it. Apple II DeskTop was the Apple II's own Finder: the disks
and their files as icons, opened with a double click, moved with the mouse,
and the menus and the desk accessories along the top of the screen, all in
the 560 dots across of the 80 column graphics of a //e with 128 KB.

It started as MouseDesk, by the French company Version Soft, which Apple then
sold as Apple II DeskTop. Today a group of Apple II enthusiasts keeps it alive
and adds to it, at [a2desktop.com](https://a2desktop.com), and izapple2
carries their version 1.4 inside it. This page starts it, opens a disk, reads
a file, asks the machine what it has inside, and uses the Calculator.

## What you need

Only izapple2. Its `desktop` model is an enhanced Apple //e with a mouse
card, and with the 800 KB disk of Apple II DeskTop 1.4, which comes inside
izapple2, on a hard disk interface card.

## Start it

1. **Start izapple2 with the `desktop` model:**

   ```bash
   izapple2 -model desktop
   ```

   The machine starts ProDOS from the disk in slot 7, and ProDOS starts
   DeskTop. There is a disk, *A2.DeskTop*, at the top right, the Trash at the
   bottom right, a menu bar, and a pointer.

   ![The desktop](images/desktop/desktop.png)

   The mouse is the mouse of your computer: wherever the pointer is on the
   izapple2 window, it is on the Apple II screen. The time on the right comes
   from the clock of the machine, a No-Slot Clock under its ROM, which
   izapple2 sets from the clock of your computer.

## Open a disk

2. **Double-click the disk icon.** A window opens out of it, with the files
   of the disk.

   ![The disk opened](images/desktop/open-disk.gif)

   The top of the window says how many items there are, how much the disk
   holds and how much of it is free. *ProDOS* is the operating system,
   *DeskTop.system* is DeskTop, and the folders hold the rest.

3. **Double-click *Read.Me*.** DeskTop shows the text in a window of its own,
   with a scroll bar. Press Escape to close it.

   ![The Read.Me file](images/desktop/read-me.png)

## What is in the machine

4. **Pull down the Apple menu**, at the left of the menu bar, holding the
   button down, and let go on *About This Apple II*.

   ![The Apple menu](images/desktop/apple-menu.gif)

   Under the two *About* items are the desk accessories, small programs that
   open over DeskTop.

5. **Read what the machine is.**

   ![About This Apple II](images/desktop/about.png)

   It is the machine izapple2 built with the `desktop` model: an enhanced
   Apple //e, with the 65C02 processor, and 8,256 KB of memory, the 128 KB of
   the //e and 8 MB more on a RAMWorks memory card. Then what is in each of
   its seven slots: the mouse card in slot 4, the Disk II controller in slot
   6, and the hard disk card that has the disk of DeskTop in slot 7. Slots 1
   and 5 are empty in this machine, which DeskTop shows as *(unknown)*. Press
   Escape to close it.

   `izapple2 -model desktop -showConfig` lists the same machine as izapple2
   sees it, without starting it.

## The Calculator

6. **Choose *Calculator* from the Apple menu**, and click its keys, *1*, *2*,
   *\**, *3* and *=*.

   ![Twelve times three](images/desktop/calculator.gif)

   The Calculator stays open over the desktop until it is closed with the box
   at the left of its title.

## What next

The other desk accessories of the Apple menu, and the *Toys* folder, are
worth a look. [Switch on an Apple ][+](switch-on.md) shows the Apple II
before all this, with nothing on the screen but a prompt.
