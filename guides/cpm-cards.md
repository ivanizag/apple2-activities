# CP/M on three Z80 cards

[Activities](../README.md) › [Other systems](../sections/other-systems.md)

CP/M ran on the Z80 and the Apple II had a 6502, so to run CP/M an Apple II
owner bought a Z80 on a card. Microsoft made the first, the **Z80 SoftCard**,
in 1980. Its Z80 takes the machine over while the 6502 waits, and uses the
Apple's own memory, so the CP/M it runs has what the Apple has left. Others
followed, and some worked the other way round. On the **PCPI Applicard** and
on Microsoft's own **SoftCard II** of 1984, the Z80 has 64 KB of memory of
its own and runs CP/M there. The 6502 stays in charge of the screen, the
keyboard and the disks, and the two pass bytes to each other through a port.

Each card came with its own CP/M, and the SoftCard went through several
versions, from Microsoft and from others. This page starts every version
found for the three cards izapple2 has, one after the other, and compares
them: what each says when it starts, what came on its disk, and how much
memory the same Microsoft BASIC finds free on each.

## What you need

The disks are in the CP/M folder of the [Asimov
archive](https://mirrors.apple2.org.za/ftp.apple.asimov.net/images/cpm/os/),
and `./fetch-disks.sh` in this repository downloads them all into `disks/`
and checks them. In that folder:

- **`Softcard 16-sector disk (Microsoft 1980).dsk`**: CP/M 2.20 of the
  SoftCard, for 44 KB.
- **`CPM1.PO`**, in `softcard.zip`: CP/M 2.20B of the SoftCard, for 56 KB.
- **`CPM2.23(60k).dsk`**: a CP/M 2.23B of 63 KB, marked as Datasoft's.
- **`CPM60K.DSK`**: CP/M 2.23 of the SoftCard, for 60 KB.
- **`CPM3.1_Disk_1.dsk`**, in `CPM3.1_Z80_Softcard.zip`: CP/M 3.0 for the
  SoftCard.
- **`CPAM40B.dsk`**, **`A.E. CPM Plus card.dsk`** and **`CPAM51A.dsk`**:
  CP/AM 4.0B, 5.1 and 5.1.1, the CP/M of Applied Engineering.
- **`Microsoft Softcard II CPM 2.28B.DSK`**: CP/M 2.28B of the SoftCard II.
- **`pcpiboot.do`**, in the `pcpi` folder: CP/M 2.2 of the Applicard, the
  boot disk of PCPI.

The manuals are on the Internet Archive: the
[SoftCard CP/M Reference Manual](https://archive.org/details/Microsoft_SoftCard_CPM_Reference_Manual)
of Microsoft, the
[SoftCard II Installation and Operations Manual and Programmer's Manual](https://archive.org/details/softcard-ii-installation-and-operations-manual-and-programmers-manual-1),
and the
[Appli-Card 2.0 Installation and User's Manual](https://archive.org/details/appli-card-2.0-installation-and-users-manual)
of PCPI.

## The machines

Four machines, each with one disk drive, a Disk II controller card in slot 6
and a Z80 card in slot 4:

- an Apple \]\[+, its 6502 at 1 MHz, 48 KB of memory, Applesoft BASIC in its
  ROM and the keyboard that types only capitals, with a 16 KB Language Card
  in slot 0, which makes the memory 64 KB, and the **Microsoft Z80
  SoftCard**;
- an Apple //e, its 6502 at 1 MHz, 128 KB of memory with its extended 80
  column card, with the **SoftCard**;
- the same //e with the **Microsoft SoftCard II**;
- the same //e with the **PCPI Applicard**.

Each version starts with its disk in drive 1; the command of each is in its
step. The SoftCard II and the Applicard need an izapple2 built from its
master, the releases don't have them yet.

## The SoftCard

1. **Start CP/M 2.20** on the \]\[+, the CP/M of the SoftCard of 1980:

   ```bash
   izapple2 -model none -board 2plus -cpu 6502 -screen green \
       -rom "<internal>/Apple2_Plus.rom" \
       -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
       -s0 language \
       -s4 z80softcard \
       -s6 'diskii,disk1=disks/Softcard 16-sector disk (Microsoft 1980).dsk'
   ```

   ![CP/M 2.20 of 44K](images/cpm-cards/softcard-2.20.png)

   *44K*: the Z80 shares the 48 KB of the \]\[+. The Apple keeps the
   bottom of its memory, the text screen among it, and the SoftCard moves the
   Z80's addresses so that CP/M still finds its memory where it wants it.
   This version is for a \]\[+ without the Language Card, and starts on one
   with it too.

2. **Start CP/M 2.20B**, the version that came after, for 56 KB:

   ```bash
   izapple2 -model none -board 2plus -cpu 6502 -screen green \
       -rom "<internal>/Apple2_Plus.rom" \
       -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
       -s0 language \
       -s4 z80softcard \
       -s6 'diskii,disk1=disks/CPM1.PO'
   ```

   ![CP/M 2.20B of 56K](images/cpm-cards/softcard-2.20b.png)

   *56K*, for a \]\[+ with the Language Card. [CP/M on the Z80
   SoftCard](cpm.md) uses this disk. In izapple2, 2.20 and 2.20B stop on the //e after the drive, so the
   page starts them on the \]\[+, the machine of their time.

3. **Start the CP/M 2.23B of Datasoft**, also for the \]\[+, which says it
   has 63 KB:

   ```bash
   izapple2 -model none -board 2plus -cpu 6502 -screen green \
       -rom "<internal>/Apple2_Plus.rom" \
       -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
       -s0 language \
       -s4 z80softcard \
       -s6 'diskii,disk1=disks/CPM2.23(60k).dsk'
   ```

   ![CP/M 2.23B of Datasoft](images/cpm-cards/datasoft-2.23b.png)

   Its GBASIC is Microsoft's of 2.20B, and leaves as much memory as theirs,
   below.

4. **Start CP/M 2.23** on the //e, Microsoft's version of 1982, for 60 KB:

   ```bash
   izapple2 -model none -board 2e -cpu 6502 -screen green \
       -rom "<internal>/Apple2e.rom" \
       -charrom "<internal>/Apple IIe Video Unenhanced.bin" \
       -s0 language \
       -s4 z80softcard \
       -s6 'diskii,disk1=disks/CPM60K.DSK'
   ```

   ![CP/M 2.23 of 60K](images/cpm-cards/softcard-2.23.png)

   It knows the //e: 80 columns and lower case, with the 80 column card of
   the //e. **Type `DIR`** for the files of its disk:

   ![The disk of CP/M 2.23](images/cpm-cards/softcard-2.23-dir.png)

   The tools of Digital Research, `PIP`, `ED`, `ASM`, `DDT`, `STAT`,
   `SUBMIT`, with Microsoft's BASIC-80, `MBASIC` and `GBASIC`, and the
   programs of the SoftCard, `CONFIGIO.BAS`, `APDOS`, `COPY`, `CPM60` and
   others.

5. **Start CP/M 3.0**, a port of Digital Research's CP/M Plus to the
   SoftCard, by Werner Muenchheimer in Germany in 1989:

   ```bash
   izapple2 -model none -board 2e -cpu 6502 -screen green \
       -rom "<internal>/Apple2e.rom" \
       -charrom "<internal>/Apple IIe Video Unenhanced.bin" \
       -s0 language \
       -s4 z80softcard \
       -s6 'diskii,disk1=disks/CPM3.1_Disk_1.dsk'
   ```

   ![CP/M 3.0](images/cpm-cards/softcard-3.0.png)

   Its loader lists the parts of the system it loads, the BIOS and the BDOS,
   and where, and leaves *46K TPA*, the memory for the programs. **Type
   `DIR`**: some of the programs of CP/M 3 are new, `DATE`, `SET`,
   `SETDEF` and `SHOW` among them.

   ![The disk of CP/M 3.0](images/cpm-cards/softcard-3.0-dir.png)

6. **Start CP/AM 4.0B**, the CP/M of Applied Engineering, of 1985:

   ```bash
   izapple2 -model none -board 2e -cpu 6502 -screen green \
       -rom "<internal>/Apple2e.rom" \
       -charrom "<internal>/Apple IIe Video Unenhanced.bin" \
       -s0 language \
       -s4 z80softcard \
       -s6 'diskii,disk1=disks/CPAM40B.dsk'
   ```

   ![CP/AM 4.0B](images/cpm-cards/cpam-4.0b.png)

   Applied Engineering sold a Z80 card of its own, the Z-80 Plus, which
   works as the SoftCard does, and its CP/AM starts on the SoftCard of
   izapple2. On the Applicard and the SoftCard II it says *NO A.E. Z-80 PLUS
   CARD FOUND.*

7. **Start CP/AM 5.1** and **CP/AM 5.1.1**, of 1986, for 60 KB:

   ```bash
   izapple2 -model none -board 2e -cpu 6502 -screen green \
       -rom "<internal>/Apple2e.rom" \
       -charrom "<internal>/Apple IIe Video Unenhanced.bin" \
       -s0 language \
       -s4 z80softcard \
       -s6 'diskii,disk1=disks/A.E. CPM Plus card.dsk'
   ```

   ![CP/AM 5.1](images/cpm-cards/cpam-5.1.png)

   ```bash
   izapple2 -model none -board 2e -cpu 6502 -screen green \
       -rom "<internal>/Apple2e.rom" \
       -charrom "<internal>/Apple IIe Video Unenhanced.bin" \
       -s0 language \
       -s4 z80softcard \
       -s6 'diskii,disk1=disks/CPAM51A.dsk'
   ```

   ![CP/AM 5.1.1](images/cpm-cards/cpam-5.1.1.png)

   **Type `DIR`** on 5.1.1. Beside the programs of CP/M are Applied
   Engineering's own, `RAMBOOT`, `MEGDRIVE`, `CPAM60F` and `CPAM60S` among
   them. A file named `-10/7/87 R21` looks like the label of someone's own
   copy, not a master.

   ![The disk of CP/AM 5.1.1](images/cpm-cards/cpam-5.1.1-dir.png)

## The SoftCard II

8. **Start CP/M 2.28B** on the //e with the SoftCard II:

   ```bash
   izapple2 -model none -board 2e -cpu 6502 -screen green \
       -rom "<internal>/Apple2e.rom" \
       -charrom "<internal>/Apple IIe Video Unenhanced.bin" \
       -s0 language \
       -s4 z80softcard2 \
       -s6 'diskii,disk1=disks/Microsoft Softcard II CPM 2.28B.DSK'
   ```

   ![CP/M 2.28B of the SoftCard II](images/cpm-cards/softcard-ii-2.28b.png)

   *64K*: all of the Z80's memory is its own, none of it the Apple's.
   **Type `DIR`**:

   ![The disk of the SoftCard II](images/cpm-cards/softcard-ii-2.28b-dir.png)

   The programs of the SoftCard's 2.23, but for `MBASIC`, `CPM60` and
   `DOWNLOAD`: of the two BASICs, Microsoft kept GBASIC, the one with the
   graphics. On the \]\[+ this CP/M prints its
   words in lower case, which the \]\[+ shows as other signs.

## The Applicard

9. **Start CP/M 2.2** on the //e with the Applicard:

   ```bash
   izapple2 -model none -board 2e -cpu 6502 -screen green \
       -rom "<internal>/Apple2e.rom" \
       -charrom "<internal>/Apple IIe Video Unenhanced.bin" \
       -s0 language \
       -s4 applicard \
       -s6 'diskii,disk1=disks/pcpiboot.do'
   ```

   ![CP/M 2.2 of the Applicard](images/cpm-cards/applicard-2.0.png)

   This is the CP/M of Digital Research as it came, version 2.2, with PCPI's
   part, version 2.0, under it. Licensed to MicroPro, the makers of
   WordStar. The Z80 has its own memory, so the Apple's is free, and PCPI
   uses it as a buffer for the printer: *83.5K* on the //e.

   **Type `DIR`**:

   ![The disk of the Applicard](images/cpm-cards/applicard-2.0-dir.png)

   The tools of Digital Research again, and PCPI's own, `DRIVERS`,
   `ADOSXFER`, `COPYFRMT`, `CLRPBUF`, `DOSRDSK` and `NRDRDSK`. No BASIC.

10. **Start it on the \]\[+**, with the Applicard in slot 4:

    ```bash
    izapple2 -model none -board 2plus -cpu 6502 -screen green \
        -rom "<internal>/Apple2_Plus.rom" \
        -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
        -s0 language \
        -s4 applicard \
        -s6 'diskii,disk1=disks/pcpiboot.do'
    ```

    ![CP/M 2.2 of the Applicard on the \]\[+](images/cpm-cards/applicard-2.0-plus.png)

    The same CP/M, in capitals and 40 columns, and a buffer of *35.5K*: the
    \]\[+ has 64 KB with its Language Card, the //e has 128 KB with its 80
    column card.

## The same BASIC on each

GBASIC, Microsoft's BASIC-80 with the graphics of the Apple II, is on the
disks of the SoftCard and of the SoftCard II. The number it says when it
starts, the memory free for a program, compares the versions.

11. **Type `GBASIC`** on CP/M 2.20, 2.23 and 2.28B:

    ![GBASIC on CP/M 2.20](images/cpm-cards/softcard-2.20-gbasic.png)

    ![GBASIC on CP/M 2.23](images/cpm-cards/softcard-2.23-gbasic.png)

    ![GBASIC on CP/M 2.28B](images/cpm-cards/softcard-ii-2.28b-gbasic.png)

    On 2.20B and on Datasoft's 2.23B it says *17393 BYTES FREE*; both have
    the GBASIC of November 1980. The Applicard, CP/AM and CP/M 3 have no
    GBASIC. Type `SYSTEM` to go back to CP/M.

| Card | CP/M | Of | Machine of the page | Its memory | GBASIC leaves |
|---|---|---|---|---|---|
| SoftCard | 2.20 | Microsoft, 1980 | \]\[+ | 44K | 5,105 bytes |
| SoftCard | 2.20B | Microsoft, 1980 | \]\[+ | 56K | 17,393 bytes |
| SoftCard | 2.23B | Datasoft, 1982 | \]\[+ | 63K | 17,393 bytes |
| SoftCard | 2.23 | Microsoft, 1982 | //e | 60K | 21,489 bytes |
| SoftCard | 3.0 | Werner Muenchheimer, 1989 | //e | 46K TPA | |
| SoftCard, Z-80 Plus | CP/AM 4.0B | Applied Engineering, 1985 | //e | 44K | |
| SoftCard, Z-80 Plus | CP/AM 5.1 and 5.1.1 | Applied Engineering, 1986 | //e | 60K | |
| SoftCard II | 2.28B | Microsoft, 1984 | //e | 64K | 32,807 bytes |
| Applicard | 2.2, PCPI 2.0 | Digital Research and PCPI | //e and \]\[+ | its own 64 KB | |

The years and the sizes are those the systems say when they start.

On the SoftCard II, GBASIC's `HPLOT` doesn't end in izapple2 yet, so the
page doesn't draw with it; see [CP/M on the Z80 SoftCard](cpm.md) for a
drawing on the SoftCard.

## Other disks found

These were found too, and are not on the page:

- 44 KB versions of 2.20B, the *Original Master* of the [RIAG
  collection](https://archive.org/details/riag_008_049_-_Microsoft_-_Softcard_-_Original_-_Master),
  and of 2.23, `CPMV233.DSK` on Asimov: the same versions, for a \]\[+
  without a Language Card.
- The Applicard's disk as sold with the StarCard, `Starcard - CPM Version 2.2 CP2-681-10107.DSK` on Asimov: the
  same CP/M 2.2 and PCPI 2.0.
- The Premium SoftCard IIe of Microsoft, a Z80 for the auxiliary slot of the
  //e, with CP/M 2.25 and 2.26: izapple2 doesn't have this card.
- The CP/M card of ALS, Advanced Logic Systems: izapple2 doesn't have this
  card either.
- *XEBEC APPLE CP/M V2.20B*, on the [Internet
  Archive](https://archive.org/details/xebec-apple-cpm-v-2.20-b), for the
  SoftCard: in izapple2 it finds the card and then fills the screen with
  `@`.

## What next

**Next:** [CP/M-65: CP/M on the 6502](cpm65.md), CP/M written again for the
6502 itself.

[CP/M on the Z80 SoftCard](cpm.md) does more with CP/M 2.20B: an assembler
source, a program in BASIC-80 saved as a file, and a drawing in GBASIC.
[Turbo Pascal on CP/M](turbo-pascal.md) programs on the SoftCard.
