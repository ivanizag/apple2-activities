# Apple Pascal

[Back to the activities](../README.md)

Pascal was the language for teaching programming at the end of the
seventies, and the University of California, San Diego had a version of it
that ran the same on very different computers: UCSD Pascal compiled programs
not into the code of a processor but into *p-code*, for a machine of its own
that an interpreter ran. The p-System came with its own editor, filer,
compiler, assembler and linker, and Apple adapted it as **Apple Pascal**, in
1979, with libraries of its own for the machine, among them turtle graphics,
to draw by steering a turtle, as in Logo.

This page starts Apple Pascal 1.3, the last version, on an enhanced Apple //e
with four disk drives, looks at a disk with the Filer, writes a program in
the editor, and compiles and runs it.

## What you need

Only izapple2: the four disks of Apple Pascal 1.3, `APPLE0` to `APPLE3`,
come inside it.

## The machine

An enhanced Apple //e with four disk drives:

- the 65C02 processor at 1 MHz, 128 KB of memory, and a RAMWorks memory card
  with 8 MB more in its auxiliary slot, with the 80 column card;
- a No-Slot Clock under the ROM;
- a VidHD card in slot 2, a FASTChip accelerator in slot 3 and a Mockingboard
  sound card in slot 4, unused here;
- a Disk II controller card in slot 5, with `APPLE3` and `APPLE0` in its two
  drives;
- a Disk II controller card in slot 6, with `APPLE1`, the one it starts
  from, and `APPLE2` in its two drives.

```bash
izapple2 -model pascal
```

## Start it

1. **Start izapple2** with the command above. Apple Pascal starts from
   `APPLE1`, greets you in 80 columns, and shows its command line at the top
   of the screen: each command is one key, the letter before the
   parenthesis.

   ![Apple Pascal started](images/pascal/started.png)

   The date is the one kept on the disk, not the one of the clock of the
   machine.

## The Filer

2. **Press F for the Filer, then L to list a directory, and type `*` and
   Return**: `*` is the disk the system started from.

   ![The boot disk in the Filer](images/pascal/filer.png)

   The disks of the p-System are *volumes*, with names: this is `APPLE1:`.
   The sizes are in blocks of 512 bytes. `SYSTEM.PASCAL` is the system
   itself, `SYSTEM.EDITOR` the editor and `SYSTEM.FILER` the Filer, each
   loaded from the disk when its command is given. Press Q to leave the
   Filer.

## A program

3. **Press E for the editor**, and Return to start a new file. **Press I to
   insert**, type the program, and **press Control-C** to end the insertion
   (Escape would throw it away):

   ```pascal
   (*$S+*)
   program spiral;
   uses turtlegraphics;
   var i: integer;
   begin
   initturtle;
   pencolor(white);
   for i := 1 to 90 do
   begin
   move(i * 2);
   turn(89)
   end;
   readln
   end.
   ```

   ![The program in the editor](images/pascal/editor.png)

   The editor keeps the indentation of the line above, so the program is
   typed without any. The program draws ninety lines with the turtle, each
   longer than the one before, turning 89 degrees after each: a square that
   does not quite close, and turns as it grows. `(*$S+*)` is an option for
   the compiler, not part of the program: it makes the compiler swap its
   parts in and out of memory, to leave room for the turtle graphics. The
   system says it has 64 KB, and the compiler stops for lack of memory
   without it.

4. **Press Q to quit the editor**, **U to keep the program** in the work
   file, and **E** to go back to the command line.

   ![Leaving the editor](images/pascal/quit.png)

   The work file, `SYSTEM.WRK.TEXT`, is the program the commands of the
   command line work on.

## Compile and run

5. **Press R to run it.** The work file has not been compiled, so the
   compiler starts first. It asks for a listing file: press Return for none.
   It reads the turtle graphics unit, `TURTLEGR`, from the library, and the
   program, a dot for each line.

   ![The program compiled](images/pascal/compiled.png)

6. **Watch it run.** The screen changes to the high resolution graphics, and
   the turtle draws. Press F6 in izapple2 for colour.

   ![The turtle drawing](images/pascal/drawing.gif)

   Press Return to end the program, which waits for it with `readln`, and go
   back to the command line.

## What next

[The Apple \]\[ of 1977](apple-ii.md) has the Mini-Assembler, machine code
written a line at a time; the p-System has a full assembler, `A(ssem`.
