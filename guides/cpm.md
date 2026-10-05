# CP/M on the Z80 SoftCard

[Back to the activities](../README.md)

At the end of the seventies the business software of small computers ran on
CP/M, the operating system of Digital Research, and CP/M ran on the Z80 and
the 8080 processors, not on the 6502 of the Apple II. In 1980 Microsoft sold
a card that changed that, the **Z80 SoftCard**: a Z80 processor on a card,
which takes over the machine while the 6502 waits, with CP/M and Microsoft's
BASIC on a diskette. The Apple II could then run the programs written for
CP/M.

This page starts CP/M on an Apple \]\[+ with the SoftCard, looks at its
disk, and writes a program in Microsoft BASIC-80.

## What you need

Only izapple2: the diskette of CP/M 2.20B for the SoftCard comes inside it.

## The machine

An Apple \]\[+ with the Microsoft Z80 SoftCard:

- the 6502 processor at 1 MHz, 48 KB of memory, and Applesoft BASIC in its ROM;
- a 16 KB Language Card in slot 0, which makes the memory of CP/M 56 KB;
- the Microsoft Z80 SoftCard in slot 4;
- a Disk II controller card in slot 6, with the CP/M diskette in drive 1.

```bash
izapple2 -model cpm
```

## Start it

1. **Start izapple2** with the command above. The Apple \]\[+ starts from
   the diskette as always, and what it loads hands the machine over to the
   Z80. CP/M shows its version, 2.20B for 56K, and its prompt, `A>`: the
   drive it works on is drive A.

   ![CP/M started](images/cpm/started.png)

## The disk

2. **Type `DIR` and then `STAT`**, each with Return. `DIR` lists the files
   of the disk, with a name of up to eight letters and a type of three:
   `COM` is a program, run by typing its name. `STAT` says how much room is
   left on it.

   ![The files of the disk](images/cpm/dir.png)

   `PIP` copies files, `FORMAT` prepares a blank diskette, and `MBASIC` is
   Microsoft BASIC.

3. **Type `TYPE DUMP.ASM`.** `TYPE` shows a text file on the screen: this
   one is the source of a program of CP/M, in the assembly language of the
   8080, which the Z80 also runs. Press Control-C to stop it.

   ![An assembler source](images/cpm/type.png)

## Microsoft BASIC-80

4. **Type `MBASIC`.** BASIC-80 starts, says how much memory it has left for
   your program, and waits with `OK`. Type the program and run it:

   ```basic
   10 FOR I = 1 TO 10
   20 PRINT USING "###  #####.##"; I, SQR(I)*1000
   30 NEXT
   RUN
   ```

   ![A program in BASIC-80](images/cpm/mbasic.gif)

   It is the BASIC of Microsoft on another processor, a cousin of Applesoft:
   the same language with more, as `PRINT USING`, which lines up the numbers
   in columns with the decimals given.

5. **Type `SYSTEM`** to leave BASIC and go back to CP/M.

## What next

[Apple Pascal](pascal.md) is another system that brought its own world to
the Apple II, the p-System of the University of California.
