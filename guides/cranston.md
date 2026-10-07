# Cranston Manor, from the start to the end

[Back to the activities](../README.md)

**Cranston Manor**, *Hi-Res Adventure #3* of On-Line Systems, came out in
1981 for the Apple \]\[. It starts on Main Street of a small town, and goes
through the gate of a manor, its rooms, its tower, a cistern under it, a
subway and caves, picking up its treasures: the inventory of the game marks
each one with asterisks, `*A NEST OF GOLDEN EGGS*`. A mouse in a cage
frightens off the suits of armour that follow you, and the way out is the
main gate, locked.

You type one or two words, `GET LANTERN`, `PLAY ORGAN`, `PRIME PUMP`, and the
game answers under the picture. The whole game is on one side of a diskette.

This page plays it to its end, all 213 commands, with a picture of the screen
after each one: 245 pictures. The pictures are of a colour television, with
the four lines of text under them drawn white and sharp, as a monochrome
monitor showed them: the television blurred them into fringes of colour,
hard to read.

## What you need

**Cranston Manor (4am and san inc crack)**, from
[its page on the Internet Archive](https://archive.org/details/CranstonManor4amCrack):
the zip `Cranston Manor (4am and san inc crack).zip`, with the disk
`Cranston Manor (4am and san inc crack).dsk`. `./fetch-disks.sh` in this
repository downloads it into `disks/` and checks it.

The commands of this page are also in this repository, a line each,
[listings/cranston.txt](listings/cranston.txt). They follow the walkthrough of
[The Sierra Help Pages](https://sierrahelp.com/Walkthroughs/CranstonManorWalkthrough.html),
with what the machine showed it needs: the key is in the desk of the smoking
room, `OPEN DESK` before `GET KEY`, and the walkthrough stops before the
end, which is to unlock the main gate and go out through it.

The [manual of Cranston Manor](https://archive.org/details/cranston-manor) is
on the Internet Archive.

## The machine

An Apple \]\[+:

- the 6502 processor at 1 MHz and 48 KB of memory;
- a Disk II controller card in slot 6, with the game in drive 1.

```bash
izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s6 'diskii,disk1=disks/Cranston Manor (4am and san inc crack).dsk'
```

Shorter, the model `2plus` of izapple2 plays it too: the same Apple \]\[+,
with the 16 KB of a language card in slot 0 and a Videx 80 column card in
slot 3 more, on a colour monitor with its scan lines.

```bash
izapple2 -model 2plus 'disks/Cranston Manor (4am and san inc crack).dsk'
```

## Playing it

1. **Start izapple2** with the command above. The game starts on Main
   Street.

2. **Type the commands**, each followed by Return. Each step below is what
   the game shows, and the command to type then. When the picture stays the
   same and only the text goes on, the page shows the picture once, and
   under it the lines of text that came after. When the game stops in the
   middle of a longer text, **press Return** to go on.

3. **At the end**, after its last words, the game stops in the Monitor of
   the ROM, with the registers of the processor on the screen and its
   prompt, `*`.

## The walkthrough

<!-- The walkthrough, written by the generator -->

- [1. From Main Street to the iron gate](#1-from-main-street-to-the-iron-gate)
- [2. The hedge maze and the attic](#2-the-hedge-maze-and-the-attic)
- [3. The mouse](#3-the-mouse)
- [4. The organ and the secret room](#4-the-organ-and-the-secret-room)
- [5. The desk, the tower and the spyglass](#5-the-desk-the-tower-and-the-spyglass)
- [6. The candlestick, the pot and the bills on the rope](#6-the-candlestick-the-pot-and-the-bills-on-the-rope)
- [7. The cistern](#7-the-cistern)
- [8. The subway and the computer vault](#8-the-subway-and-the-computer-vault)
- [9. The caves](#9-the-caves)
- [10. The bedrooms and the oak tree](#10-the-bedrooms-and-the-oak-tree)
- [11. The emeralds of the fountain, and out by the gate](#11-the-emeralds-of-the-fountain-and-out-by-the-gate)
- [The end](#the-end)

### 1. From Main Street to the iron gate

<img src="images/cranston/0001.png" width="400" alt="YOU ARE STANDING ON MAIN STREET LOOKING UP FIRST STREET.">

**Type `EAST`**

<img src="images/cranston/0002.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE STANDING ON MAIN STREET, LOOKING UP 2ND STREET.">

**Type `NORTH`**

<img src="images/cranston/0003.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE FACING WEST, IN FRONT OF THE GENERAL STORE.">

**Type `WEST`**

<img src="images/cranston/0004.png" width="400" alt="THERE IS A SMALL LANTERN HERE. YOU ARE IN A DUSTY STORE. THERE IS A SIGN HERE.">

**Type `GET LANTERN`**

<img src="images/cranston/0005.png" width="400" alt="N YOU ARE IN A DUSTY STORE. THERE IS A SIGN HERE.">

**Type `EAST`**

<img src="images/cranston/0006.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE FACING WEST, IN FRONT OF THE GENERAL STORE.">

**Type `SOUTH`**

<img src="images/cranston/0007.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE STANDING ON MAIN STREET, LOOKING UP 2ND STREET.">

**Type `SOUTH`**

<img src="images/cranston/0008.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE STANDING ON PIERCE LOOKING UP 2ND STREET.">

**Type `SOUTH`**

<img src="images/cranston/0009.png" width="400" alt="2ND STREET. --------------- ENTER COMMAND?SOUTH YOU&#39;RE AT THE SOUTH END OF 2ND STREET.">

**Type `EAST`**

<img src="images/cranston/0010.png" width="400" alt="--------------- ENTER COMMAND?EAST THERE IS A CROWBAR HERE. YOU ARE IN THE JUNK YARD.">

**Type `GET CROWBAR`**

<img src="images/cranston/0011.png" width="400" alt="--------------- ENTER COMMAND?GET CROWBA R YOU ARE IN THE JUNK YARD.">

**Type `NORTH`**

<img src="images/cranston/0012.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE STANDING ON PIERCE LOOKING UP CRANSTON BLVD.">

**Type `NORTH`**

<img src="images/cranston/0013.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE STANDING ON MAIN STREET, LOOKING UP CRANSTON BLVD.">

**Type `NORTH`**

<img src="images/cranston/0014.png" width="400" alt="A PLAQUE READS &#34;CRANSTON MANOR&#34; YOU ARE AT THE MAIN GATE TO CRANSTON MANOR.">

**Type `WEST`**

<img src="images/cranston/0015.png" width="400" alt="MANOR. --------------- ENTER COMMAND?WEST THE WALL CUTS NORTH OFF CRANSTON.">

**Type `NORTH`**

<img src="images/cranston/0016.png" width="400" alt="THE WALL CUTS NORTH OFF CRANSTON. --------------- ENTER COMMAND?NORTH YOU ARE AT A SMALL IRON GATE.">

**Type `USE CROWBAR`**

### 2. The hedge maze and the attic

<img src="images/cranston/0017.png" width="400" alt="R OK. YOU ARE AT A SMALL IRON GATE.">

**Type `EAST`**

<img src="images/cranston/0018.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE INSIDE THE STONE WALL, NEXT TO A SMALL OPEN GATE.">

**Type `EAST`**

<img src="images/cranston/0019.png" width="400" alt="A SMALL OPEN GATE. --------------- ENTER COMMAND?EAST YOU ARE SURROUNDED BY TALL HEDGES.">

**Type `EAST`**

<img src="images/cranston/0020.png" width="400" alt="YOU ARE SURROUNDED BY TALL HEDGES. --------------- ENTER COMMAND?EAST YOU ARE SURROUNDED BY TALL HEDGES.">

**Type `EAST`**

<img src="images/cranston/0021.png" width="400" alt="YOU ARE SURROUNDED BY TALL HEDGES. --------------- ENTER COMMAND?EAST YOU ARE SURROUNDED BY TALL HEDGES.">

**Type `NORTH`**

<img src="images/cranston/0022-1.png" width="400" alt="THERE IS A BAG OF JEWELS HERE. YOU ARE AT THE GAZEBO IN THE HEDGE MAZE. THE PATH LEADS SOUTH, EAST AND WEST."><br>
<img src="images/cranston/0022.png" width="400" alt="YOU ARE AT THE GAZEBO IN THE HEDGE MAZE. THE PATH LEADS SOUTH, EAST AND WEST.">

**Type `GET JEWELS`**

<img src="images/cranston/0023.png" width="400" alt="YOU ARE AT THE GAZEBO IN THE HEDGE MAZE. THE PATH LEADS SOUTH, EAST AND WEST.">

**Type `SOUTH`**

<img src="images/cranston/0024.png" width="400" alt="WEST. --------------- ENTER COMMAND?SOUTH YOU ARE SURROUNDED BY TALL HEDGES.">

**Type `EAST`**

<img src="images/cranston/0025.png" width="400" alt="YOU ARE SURROUNDED BY TALL HEDGES. --------------- ENTER COMMAND?EAST YOU ARE SURROUNDED BY TALL HEDGES.">

**Type `EAST`**

<img src="images/cranston/0026.png" width="400" alt="YOU ARE SURROUNDED BY TALL HEDGES. --------------- ENTER COMMAND?EAST YOU ARE SURROUNDED BY TALL HEDGES.">

**Type `EAST`**

<img src="images/cranston/0027.png" width="400" alt="YOU ARE SURROUNDED BY TALL HEDGES. --------------- ENTER COMMAND?EAST YOU ARE ON A GRAVEL PATH.">

**Type `EAST`**

<img src="images/cranston/0028.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE INSIDE A GARAGE. LIGHT IS SHINING THROUGH A DOOR TO THE SOUTH.">

**Type `UP`**

<img src="images/cranston/0029.png" width="400" alt="--------------- ENTER COMMAND?UP THERE IS A GOLD COIN HERE. YOU ARE IN A DUSTY ATTIC.">

**Type `GET COIN`**

### 3. The mouse

<img src="images/cranston/0030.png" width="400" alt="YOU ARE IN A DUSTY ATTIC. --------------- ENTER COMMAND?GET COIN YOU ARE IN A DUSTY ATTIC.">

**Type `DOWN`**

<img src="images/cranston/0031.png" width="400" alt="--------------- ENTER COMMAND?DOWN YOU ARE INSIDE A GARAGE. LIGHT IS SHINING THROUGH A DOOR TO THE SOUTH.">

**Type `EAST`**

<img src="images/cranston/0032.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN A BREEZEWAY BETWEEN THE HOUSE AND GARAGE.">

**Type `EAST`**

<img src="images/cranston/0033.png" width="400" alt="THERE IS A SUIT OF ARMOR HERE. YOU ARE AT THE WEST END OF THE E/W HALL. THERE IS DOOR TO THE SOUTH.">

**Type `EAST`**

<img src="images/cranston/0034.png" width="400" alt="THERE IS A SUIT OF ARMOR HERE. YOU&#39;RE AT THE INTERSECTION OF THE N/S AND E/W HALLS.">

**Type `NORTH`**

<img src="images/cranston/0035.png" width="400" alt="--------------- ENTER COMMAND?NORTH THERE IS A SUIT OF ARMOR HERE. THIS IS THE NORTH END OF THE N/S HALL.">

**Type `WEST`**

<img src="images/cranston/0036.png" width="400" alt="THERE IS A SMALL CAGE HERE. THERE IS A SUIT OF ARMOR HERE. YOU ARE IN THE BIRD ATRIUM.">

**Type `GET CAGE`**

<img src="images/cranston/0037.png" width="400" alt="YOU ARE IN THE BIRD ATRIUM. --------------- ENTER COMMAND?GET CAGE YOU ARE IN THE BIRD ATRIUM.">

**Type `EAST`**

<img src="images/cranston/0038.png" width="400" alt="--------------- ENTER COMMAND?EAST THERE IS A SUIT OF ARMOR HERE. THIS IS THE NORTH END OF THE N/S HALL.">

**Type `SOUTH`**

<img src="images/cranston/0039.png" width="400" alt="THERE IS A SUIT OF ARMOR HERE. YOU&#39;RE AT THE INTERSECTION OF THE N/S AND E/W HALLS.">

**Type `EAST`**

<img src="images/cranston/0040.png" width="400" alt="THERE IS A SUIT OF ARMOR HERE. YOU ARE AT THE EAST END OF THE E/W HALL. THERE IS A DOOR TO THE SOUTH.">

**Type `SOUTH`**

<img src="images/cranston/0041-1.png" width="400" alt="THERE IS A HEAVY WOODEN CHAIR HERE. THERE IS A SUIT OF ARMOR HERE. YOU ARE IN THE SERVANTS DINING ROOM. THERE IS A DOOR TO THE SOUTH."><br>
<img src="images/cranston/0041.png" width="400" alt="THERE IS A SUIT OF ARMOR HERE. YOU ARE IN THE SERVANTS DINING ROOM. THERE IS A DOOR TO THE SOUTH.">

**Type `SOUTH`**

<img src="images/cranston/0042-1.png" width="400" alt="THERE IS A POT ON THE STOVE AND SOME MOLDY CHEESE ON THE FLOOR. THERE IS A SUIT OF ARMOR HERE. THIS IS THE KITCHEN. THERE IS A DOOR TO"><br>
<img src="images/cranston/0042.png" width="400" alt="THERE IS A SUIT OF ARMOR HERE. THIS IS THE KITCHEN. THERE IS A DOOR TO THE SOUTH.">

**Type `GET CHEESE`**

<img src="images/cranston/0043.png" width="400" alt="THIS IS THE KITCHEN. THERE IS A DOOR TO THE SOUTH.">

**Type `NORTH`**

<img src="images/cranston/0044-1.png" width="400" alt="THERE IS A HEAVY WOODEN CHAIR HERE. THERE IS A SUIT OF ARMOR HERE. YOU ARE IN THE SERVANTS DINING ROOM. THERE IS A DOOR TO THE SOUTH."><br>
<img src="images/cranston/0044.png" width="400" alt="THERE IS A SUIT OF ARMOR HERE. YOU ARE IN THE SERVANTS DINING ROOM. THERE IS A DOOR TO THE SOUTH.">

**Type `NORTH`**

<img src="images/cranston/0045.png" width="400" alt="THERE IS A SUIT OF ARMOR HERE. YOU ARE AT THE EAST END OF THE E/W HALL. THERE IS A DOOR TO THE SOUTH.">

**Type `NORTH`**

<img src="images/cranston/0046.png" width="400" alt="THERE IS A SUIT OF ARMOR HERE. THIS ROOM HAS A HOLE IN THE WALL AND THE FLOOR AND A DOOR TO THE SOUTH.">

**Type `EAST`**

<img src="images/cranston/0047.png" width="400" alt="--------------- ENTER COMMAND?EAST THERE IS A MOUSE SITTING HERE. YOU ARE IN THE MOUSE ROOM.">

**Type `DROP CHEESE`**

<img src="images/cranston/0048-1.png" width="400" alt="THE MOUSE SNIFFS CAUTIOUSLY AND THEN BEGINS EATING THE CHEESE. THERE IS A MOUSE SITTING HERE. YOU ARE IN THE MOUSE ROOM."><br>
<img src="images/cranston/0048.png" width="400" alt="BEGINS EATING THE CHEESE. THERE IS A MOUSE SITTING HERE. YOU ARE IN THE MOUSE ROOM.">

**Type `GET MOUSE`**

<img src="images/cranston/0049.png" width="400" alt="--------------- ENTER COMMAND?GET MOUSE YOU CAUGHT THE LITTLE RASCAL! YOU ARE IN THE MOUSE ROOM.">

**Type `WEST`**

<img src="images/cranston/0050.png" width="400" alt="--------------- ENTER COMMAND?WEST THIS ROOM HAS A HOLE IN THE WALL AND THE FLOOR AND A DOOR TO THE SOUTH.">

**Type `WEST`**

<img src="images/cranston/0051.png" width="400" alt="THERE IS AN EXPENSIVE TEAPOT HERE. THERE IS A SUIT OF ARMOR HERE. YOU ARE IN THE SITTING ROOM.">

**Type `DROP MOUSE`**

<img src="images/cranston/0052-1.png" width="400" alt="THE MOUSE LEAPS FROM THE CAGE. THERE IS AN EXPENSIVE TEAPOT HERE. THE ARMOR IS FRIGHTENED BY THE MOUSE AND RUNS OUT OF THE DOOR."><br>
<img src="images/cranston/0052.png" width="400" alt="THE ARMOR IS FRIGHTENED BY THE MOUSE AND RUNS OUT OF THE DOOR. YOU ARE IN THE SITTING ROOM.">

**Type `TAKE TEAPOT`**

<img src="images/cranston/0053.png" width="400" alt="--------------- ENTER COMMAND?TAKE TEAPO T YOU ARE IN THE SITTING ROOM.">

**Type `GET MOUSE`**

### 4. The organ and the secret room

<img src="images/cranston/0054.png" width="400" alt="YOU CAUGHT THE LITTLE RASCAL! THERE IS A SUIT OF ARMOR HERE. YOU ARE IN THE SITTING ROOM.">

**Type `WEST`**

<img src="images/cranston/0055.png" width="400" alt="--------------- ENTER COMMAND?WEST THERE IS A SUIT OF ARMOR HERE. THIS IS THE NORTH END OF THE N/S HALL.">

**Type `SOUTH`**

<img src="images/cranston/0056.png" width="400" alt="THERE IS A SUIT OF ARMOR HERE. YOU&#39;RE AT THE INTERSECTION OF THE N/S AND E/W HALLS.">

**Type `SOUTH`**

<img src="images/cranston/0057.png" width="400" alt="--------------- ENTER COMMAND?SOUTH THERE IS A SUIT OF ARMOR HERE. YOU ARE IN THE N/S HALL.">

**Type `SOUTH`**

<img src="images/cranston/0058.png" width="400" alt="--------------- ENTER COMMAND?SOUTH THERE IS A SUIT OF ARMOR HERE. YOU ARE IN THE MIDDLE OF A N/S HALL.">

**Type `WEST`**

<img src="images/cranston/0059.png" width="400" alt="THERE IS A SUIT OF ARMOR HERE. YOU ARE IN THE ORGAN ROOM. THERE IS A DOOR TO THE SOUTH.">

**Type `PLAY ORGAN`**

<img src="images/cranston/0060-1.png" width="400" alt="A FEW SQUEEKS AND WHEEZES AND THE BACK OF THE FIREPLACE OPENS! YOU ARE IN THE ORGAN ROOM. THERE IS A DOOR TO THE SOUTH."><br>
<img src="images/cranston/0060.png" width="400" alt="OF THE FIREPLACE OPENS! YOU ARE IN THE ORGAN ROOM. THERE IS A DOOR TO THE SOUTH.">

**Type `NORTH`**

<img src="images/cranston/0061-1.png" width="400" alt="OK. THERE IS A DAGGER AND A CRYSTAL TRIANGLE HERE. YOU ARE IN THE SECRET ROOM."><br>
<img src="images/cranston/0061.png" width="400" alt="THERE IS A DAGGER AND A CRYSTAL TRIANGLE HERE. YOU ARE IN THE SECRET ROOM.">

**Type `GET TRIANGLE`**

### 5. The desk, the tower and the spyglass

<img src="images/cranston/0062.png" width="400" alt="--------------- ENTER COMMAND?GET TRIANG LE YOU ARE IN THE SECRET ROOM.">

**Type `SOUTH`**

<img src="images/cranston/0063.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN THE ORGAN ROOM. THERE IS A DOOR TO THE SOUTH.">

**Type `SOUTH`**

<img src="images/cranston/0064.png" width="400" alt="THERE IS A SUIT OF ARMOR HERE. YOU ARE IN THE PICTURE HALL. THERE IS A DOOR TO THE SOUTH.">

**Type `SOUTH`**

<img src="images/cranston/0065.png" width="400" alt="--------------- ENTER COMMAND?SOUTH THERE IS A SUIT OF ARMOR HERE. YOU ARE IN THE SMOKING ROOM.">

**Type `DROP MOUSE`**

<img src="images/cranston/0066-1.png" width="400" alt="THE MOUSE LEAPS FROM THE CAGE. THE ARMOR IS FRIGHTENED BY THE MOUSE AND RUNS OUT OF THE DOOR. YOU ARE IN THE SMOKING ROOM."><br>
<img src="images/cranston/0066.png" width="400" alt="THE ARMOR IS FRIGHTENED BY THE MOUSE AND RUNS OUT OF THE DOOR. YOU ARE IN THE SMOKING ROOM.">

**Type `OPEN DESK`**

<img src="images/cranston/0067.png" width="400" alt="OK. THERE IS A KEY HERE. YOU ARE IN THE SMOKING ROOM.">

**Type `GET KEY`**

<img src="images/cranston/0068.png" width="400" alt="YOU ARE IN THE SMOKING ROOM. --------------- ENTER COMMAND?GET KEY YOU ARE IN THE SMOKING ROOM.">

**Type `GET MOUSE`**

<img src="images/cranston/0069.png" width="400" alt="--------------- ENTER COMMAND?GET MOUSE YOU CAUGHT THE LITTLE RASCAL! YOU ARE IN THE SMOKING ROOM.">

**Type `EAST`**

<img src="images/cranston/0070.png" width="400" alt="YOU ARE IN THE SMOKING ROOM. --------------- ENTER COMMAND?EAST YOU ARE IN THE LIBRARY.">

**Type `EMASES`**

<img src="images/cranston/0071.png" width="400" alt="--------------- ENTER COMMAND?EMASES THE BOOK CASE SLIDES OPEN. YOU ARE IN THE BASE OF THE TOWER.">

**Type `UP`**

<img src="images/cranston/0072.png" width="400" alt="--------------- ENTER COMMAND?UP THERE IS A GOLD SPYGLASS HERE. YOU ARE IN THE LOOKOUT.">

**Type `GET SPYGLASS`**

### 6. The candlestick, the pot and the bills on the rope

<img src="images/cranston/0073.png" width="400" alt="--------------- ENTER COMMAND?GET SPYGLA SS YOU ARE IN THE LOOKOUT.">

**Type `DOWN`**

<img src="images/cranston/0074.png" width="400" alt="YOU ARE IN THE LOOKOUT. --------------- ENTER COMMAND?DOWN YOU ARE IN THE BASE OF THE TOWER.">

**Type `WEST`**

<img src="images/cranston/0075.png" width="400" alt="THE WALL CLOSED BEHIND YOU. THERE IS A SUIT OF ARMOR HERE. YOU ARE IN THE LIBRARY.">

**Type `NORTH`**

<img src="images/cranston/0076.png" width="400" alt="THERE IS A SUIT OF ARMOR HERE. YOU ARE AT THE SOUTH END OF A LONG HALL. THERE IS A DOOR TO THE SOUTH.">

**Type `EAST`**

<img src="images/cranston/0077.png" width="400" alt="THERE IS A SILVER CANDLESTICK HERE. THERE IS A SUIT OF ARMOR HERE. YOU ARE IN THE DINING ROOM.">

**Type `DROP MOUSE`**

<img src="images/cranston/0078-1.png" width="400" alt="THE MOUSE LEAPS FROM THE CAGE. THERE IS A SILVER CANDLESTICK HERE. THE ARMOR IS FRIGHTENED BY THE MOUSE AND RUNS OUT OF THE DOOR."><br>
<img src="images/cranston/0078.png" width="400" alt="THE ARMOR IS FRIGHTENED BY THE MOUSE AND RUNS OUT OF THE DOOR. YOU ARE IN THE DINING ROOM.">

**Type `GET CANDLESTICK`**

<img src="images/cranston/0079.png" width="400" alt="--------------- ENTER COMMAND?GET CANDLE STICK YOU ARE IN THE DINING ROOM.">

**Type `GET MOUSE`**

<img src="images/cranston/0080.png" width="400" alt="YOU CAUGHT THE LITTLE RASCAL! THERE IS A SUIT OF ARMOR HERE. YOU ARE IN THE DINING ROOM.">

**Type `NORTH`**

<img src="images/cranston/0081.png" width="400" alt="THERE IS A SUIT OF ARMOR HERE. THIS IS THE KITCHEN. THERE IS A DOOR TO THE SOUTH.">

**Type `DROP MOUSE`**

<img src="images/cranston/0082-1.png" width="400" alt="THE MOUSE LEAPS FROM THE CAGE. THE ARMOR IS FRIGHTENED BY THE MOUSE AND RUNS OUT OF THE DOOR. THIS IS THE KITCHEN. THERE IS A DOOR TO"><br>
<img src="images/cranston/0082.png" width="400" alt="AND RUNS OUT OF THE DOOR. THIS IS THE KITCHEN. THERE IS A DOOR TO THE SOUTH.">

**Type `GET POT`**

<img src="images/cranston/0083.png" width="400" alt="--------------- ENTER COMMAND?GET POT THIS IS THE KITCHEN. THERE IS A DOOR TO THE SOUTH.">

**Type `NORTH`**

<img src="images/cranston/0084-1.png" width="400" alt="THERE IS A HEAVY WOODEN CHAIR HERE. THERE IS A SUIT OF ARMOR HERE. YOU ARE IN THE SERVANTS DINING ROOM. THERE IS A DOOR TO THE SOUTH."><br>
<img src="images/cranston/0084.png" width="400" alt="THERE IS A SUIT OF ARMOR HERE. YOU ARE IN THE SERVANTS DINING ROOM. THERE IS A DOOR TO THE SOUTH.">

**Type `NORTH`**

<img src="images/cranston/0085.png" width="400" alt="THERE IS A SUIT OF ARMOR HERE. YOU ARE AT THE EAST END OF THE E/W HALL. THERE IS A DOOR TO THE SOUTH.">

**Type `WEST`**

<img src="images/cranston/0086.png" width="400" alt="THERE IS A SUIT OF ARMOR HERE. YOU&#39;RE AT THE INTERSECTION OF THE N/S AND E/W HALLS.">

**Type `SOUTH`**

<img src="images/cranston/0087.png" width="400" alt="--------------- ENTER COMMAND?SOUTH THERE IS A SUIT OF ARMOR HERE. YOU ARE IN THE N/S HALL.">

**Type `EAST`**

<img src="images/cranston/0088.png" width="400" alt="--------------- ENTER COMMAND?EAST THERE IS A SUIT OF ARMOR HERE. YOU ARE IN THE HUNTING ROOM.">

**Type `USE KEY`**

<img src="images/cranston/0089.png" width="400" alt="--------------- ENTER COMMAND?USE KEY OK. YOU ARE IN THE HUNTING ROOM.">

**Type `EAST`**

<img src="images/cranston/0090.png" width="400" alt="--------------- ENTER COMMAND?EAST THERE IS A SUIT OF ARMOR HERE. YOU ARE IN A CLOSET WITH A ROPE.">

**Type `CLIMB ROPE`**

<img src="images/cranston/0091.png" width="400" alt="THERE IS A STACK OF 50 DOLLAR BILLS ON THE SHELF. YOU ARE HANGING AT THE TOP OF A ROPE.">

**Type `SWING`**

<img src="images/cranston/0092-1.png" width="400" alt="THE SWINGING ROPE TAKES YOU CLOSER TO THE WALL. THERE IS A STACK OF 50 DOLLAR BILLS ON THE SHELF."><br>
<img src="images/cranston/0092.png" width="400" alt="THERE IS A STACK OF 50 DOLLAR BILLS ON THE SHELF. YOU ARE HANGING AT THE TOP OF A ROPE.">

**Type `GET BILLS`**

### 7. The cistern

<img src="images/cranston/0093.png" width="400" alt="YOU ARE HANGING AT THE TOP OF A ROPE. --------------- ENTER COMMAND?GET BILLS YOU ARE HANGING AT THE TOP OF A ROPE.">

**Type `DOWN`**

<img src="images/cranston/0094.png" width="400" alt="YOU ARE HANGING AT THE TOP OF A ROPE. --------------- ENTER COMMAND?DOWN YOU ARE IN A CLOSET WITH A ROPE.">

**Type `WEST`**

<img src="images/cranston/0095.png" width="400" alt="--------------- ENTER COMMAND?WEST THERE IS A SUIT OF ARMOR HERE. YOU ARE IN THE HUNTING ROOM.">

**Type `WEST`**

<img src="images/cranston/0096.png" width="400" alt="--------------- ENTER COMMAND?WEST THERE IS A SUIT OF ARMOR HERE. YOU ARE IN THE N/S HALL.">

**Type `SOUTH`**

<img src="images/cranston/0097.png" width="400" alt="--------------- ENTER COMMAND?SOUTH THERE IS A SUIT OF ARMOR HERE. YOU ARE IN THE MIDDLE OF A N/S HALL.">

**Type `SOUTH`**

<img src="images/cranston/0098.png" width="400" alt="THERE IS A SUIT OF ARMOR HERE. YOU ARE AT THE SOUTH END OF A LONG HALL. THERE IS A DOOR TO THE SOUTH.">

**Type `WEST`**

<img src="images/cranston/0099.png" width="400" alt="THERE IS A SUIT OF ARMOR HERE. YOU ARE IN THE PICTURE HALL. THERE IS A DOOR TO THE SOUTH.">

**Type `WEST`**

<img src="images/cranston/0100.png" width="400" alt="--------------- ENTER COMMAND?WEST THE DOOR CLOSED BEHIND YOU. YOU ARE FACING EAST, ON A STONE PORCH.">

**Type `WEST`**

<img src="images/cranston/0101.png" width="400" alt="YOU ARE FACING EAST, ON A STONE PORCH. --------------- ENTER COMMAND?WEST YOU ARE IN THE DRIVE WAY.">

**Type `WEST`**

<img src="images/cranston/0102.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN FRONT OF THE CAT FOUNTAIN. THERE IS A SIGN HERE.">

**Type `GET WATER`**

<img src="images/cranston/0103.png" width="400" alt="THE POT IS FULL OF WATER. YOU ARE IN FRONT OF THE CAT FOUNTAIN. THERE IS A SIGN HERE.">

**Type `EAST`**

<img src="images/cranston/0104.png" width="400" alt="THERE IS A SIGN HERE. --------------- ENTER COMMAND?EAST YOU ARE IN THE DRIVE WAY.">

**Type `EAST`**

<img src="images/cranston/0105.png" width="400" alt="YOU ARE IN THE DRIVE WAY. --------------- ENTER COMMAND?EAST YOU ARE FACING EAST, ON A STONE PORCH.">

**Type `OPEN DOOR`**

<img src="images/cranston/0106.png" width="400" alt="--------------- ENTER COMMAND?OPEN DOOR OK. YOU ARE FACING EAST, ON A STONE PORCH.">

**Type `EAST`**

<img src="images/cranston/0107.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN THE PICTURE HALL. THERE IS A DOOR TO THE SOUTH.">

**Type `EAST`**

<img src="images/cranston/0108.png" width="400" alt="THERE IS A SUIT OF ARMOR HERE. YOU ARE AT THE SOUTH END OF A LONG HALL. THERE IS A DOOR TO THE SOUTH.">

**Type `EAST`**

<img src="images/cranston/0109.png" width="400" alt="--------------- ENTER COMMAND?EAST THERE IS A SUIT OF ARMOR HERE. YOU ARE IN THE DINING ROOM.">

**Type `NORTH`**

<img src="images/cranston/0110.png" width="400" alt="--------------- ENTER COMMAND?NORTH THIS IS THE KITCHEN. THERE IS A DOOR TO THE SOUTH.">

**Type `DOWN`**

<img src="images/cranston/0111-1.png" width="400" alt="IT IS PITCH BLACK. THIS IS THE CISTERN ROOM. THERE IS A PUMP WITH A SIGN ON IT AND A DOOR TO THE SOUTH."><br>
<img src="images/cranston/0111.png" width="400" alt="THIS IS THE CISTERN ROOM. THERE IS A PUMP WITH A SIGN ON IT AND A DOOR TO THE SOUTH.">

**Type `LANTERN ON`**

<img src="images/cranston/0112-1.png" width="400" alt="OK. THIS IS THE CISTERN ROOM. THERE IS A PUMP WITH A SIGN ON IT AND A DOOR TO THE SOUTH."><br>
<img src="images/cranston/0112.png" width="400" alt="THIS IS THE CISTERN ROOM. THERE IS A PUMP WITH A SIGN ON IT AND A DOOR TO THE SOUTH.">

**Type `PRIME PUMP`**

<img src="images/cranston/0113-1.png" width="400" alt="OK. THERE IS A BOTTLE IN THE BOTTOM OF CISTERN. THE PUMP HAS A SIGN ON IT. THIS IS THE CISTERN ROOM. THERE IS A"><br>
<img src="images/cranston/0113.png" width="400" alt="THIS IS THE CISTERN ROOM. THERE IS A PUMP WITH A SIGN ON IT AND A DOOR TO THE SOUTH.">

**Type `PUSH BUTTON`**

<img src="images/cranston/0114-1.png" width="400" alt="OK. THE PUMP STARTS AND THE CISTERN FILLS WITH WATER. THIS IS THE CISTERN ROOM. THERE IS A"><br>
<img src="images/cranston/0114.png" width="400" alt="THIS IS THE CISTERN ROOM. THERE IS A PUMP WITH A SIGN ON IT AND A DOOR TO THE SOUTH.">

**Type `GET BOTTLE`**

<img src="images/cranston/0115-1.png" width="400" alt="OK. THE BOTTLE IS FULL OF DIAMONDS!! THIS IS THE CISTERN ROOM. THERE IS A PUMP WITH A SIGN ON IT AND A DOOR TO"><br>
<img src="images/cranston/0115.png" width="400" alt="THIS IS THE CISTERN ROOM. THERE IS A PUMP WITH A SIGN ON IT AND A DOOR TO THE SOUTH.">

**Type `GET WATER`**

### 8. The subway and the computer vault

<img src="images/cranston/0116-1.png" width="400" alt="THE POT IS FULL OF WATER. THIS IS THE CISTERN ROOM. THERE IS A PUMP WITH A SIGN ON IT AND A DOOR TO THE SOUTH."><br>
<img src="images/cranston/0116.png" width="400" alt="THIS IS THE CISTERN ROOM. THERE IS A PUMP WITH A SIGN ON IT AND A DOOR TO THE SOUTH.">

**Type `UP`**

<img src="images/cranston/0117.png" width="400" alt="--------------- ENTER COMMAND?UP THIS IS THE KITCHEN. THERE IS A DOOR TO THE SOUTH.">

**Type `SOUTH`**

<img src="images/cranston/0118.png" width="400" alt="THE SOUTH. --------------- ENTER COMMAND?SOUTH YOU ARE IN THE DINING ROOM.">

**Type `WEST`**

<img src="images/cranston/0119.png" width="400" alt="THERE IS A SUIT OF ARMOR HERE. YOU ARE AT THE SOUTH END OF A LONG HALL. THERE IS A DOOR TO THE SOUTH.">

**Type `NORTH`**

<img src="images/cranston/0120.png" width="400" alt="--------------- ENTER COMMAND?NORTH THERE IS A SUIT OF ARMOR HERE. YOU ARE IN THE MIDDLE OF A N/S HALL.">

**Type `EAST`**

<img src="images/cranston/0121.png" width="400" alt="--------------- ENTER COMMAND?EAST THERE IS A SUIT OF ARMOR HERE. YOU ARE IN A CLOSET.">

**Type `DOWN`**

<img src="images/cranston/0122.png" width="400" alt="--------------- ENTER COMMAND?DOWN YOU ARE AT THE BOTTOM OF STAIRS. THE DOOR HAS A STRANGE KEYHOLE.">

**Type `USE TRIANGLE`**

<img src="images/cranston/0123.png" width="400" alt="OK. YOU ARE AT THE BOTTOM OF STAIRS. THE DOOR HAS A STRANGE KEYHOLE.">

**Type `WEST`**

<img src="images/cranston/0124.png" width="400" alt="DOOR HAS A STRANGE KEYHOLE. --------------- ENTER COMMAND?WEST YOU ARE AT A 3 WAY INTERSECTION.">

**Type `SOUTH`**

<img src="images/cranston/0125.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE AT AN INTERSECTION WITH MUSHROOMS ON THE CEILING.">

**Type `WEST`**

<img src="images/cranston/0126.png" width="400" alt="--------------- ENTER COMMAND?WEST THIS IS THE WEST END OF A LARGE HALL.THERE IS A DOOR TO THE SOUTH.">

**Type `EAST`**

<img src="images/cranston/0127.png" width="400" alt="--------------- ENTER COMMAND?EAST THIS IS THE EAST END OF A LARGE HALL. THERE IS A DOOR TO THE SOUTH.">

**Type `NORTH`**

<img src="images/cranston/0128.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE ON CONCRETE STAIRS BY A CRAWLWAY.">

**Type `DOWN`**

<img src="images/cranston/0129-1.png" width="400" alt="THERE IS A VENDING MACHINE HERE. A SHORT TIN SOLDIER JUST MARCHED AROUND A CORNER. YOU ARE ON A DESERTED SUBWAY PLATFORM."><br>
<img src="images/cranston/0129.png" width="400" alt="A SHORT TIN SOLDIER JUST MARCHED AROUND A CORNER. YOU ARE ON A DESERTED SUBWAY PLATFORM.">

**Type `USE COIN`**

<img src="images/cranston/0130-1.png" width="400" alt="OK. A SHINEY CARD SLIPS OUT OF ANOTHER SLOT AND FALLS TO THE FLOOR. YOU ARE ON A DESERTED SUBWAY PLATFORM."><br>
<img src="images/cranston/0130.png" width="400" alt="A SHINEY CARD SLIPS OUT OF ANOTHER SLOT AND FALLS TO THE FLOOR. YOU ARE ON A DESERTED SUBWAY PLATFORM.">

**Type `GET CARD`**

<img src="images/cranston/0131.png" width="400" alt="YOU ARE ON A DESERTED SUBWAY PLATFORM. --------------- ENTER COMMAND?GET CARD YOU ARE ON A DESERTED SUBWAY PLATFORM.">

**Type `UP`**

<img src="images/cranston/0132-1.png" width="400" alt="A SHORT TIN SOLDIER JUST MARCHED AROUND A CORNER. YOU ARE ON CONCRETE STAIRS BY A CRAWLWAY."><br>
<img src="images/cranston/0132.png" width="400" alt="A CORNER. YOU ARE ON CONCRETE STAIRS BY A CRAWLWAY.">

**Type `EAST`**

<img src="images/cranston/0133.png" width="400" alt="--------------- ENTER COMMAND?EAST THIS IS THE EAST END OF A LARGE HALL. THERE IS A DOOR TO THE SOUTH.">

**Type `SOUTH`**

<img src="images/cranston/0134.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN THE SOUTH CHAMBER OF THE E/W HALL. THERE&#39;S A DOOR TO THE SOUTH.">

**Type `SOUTH`**

<img src="images/cranston/0135.png" width="400" alt="THERE IS A SIGN HERE. YOU ARE FACING SOUTH IN FRONT OF A HEAVY BLAST DOOR.">

**Type `USE CARD`**

<img src="images/cranston/0136-1.png" width="400" alt="OK. THERE&#39;S A WHIRRING SOUND. THE HEAVY DOOR SLIDES OPEN. THERE IS A PLATINUM SPHERE HERE."><br>
<img src="images/cranston/0136.png" width="400" alt="DOOR SLIDES OPEN. THERE IS A PLATINUM SPHERE HERE. YOU ARE IN THE COMPUTER VAULT.">

**Type `THROW WATER`**

<img src="images/cranston/0137-1.png" width="400" alt="OK. THE COMPUTER IS DEMOLISHED!! THERE IS A PLATINUM SPHERE HERE. YOU ARE IN THE COMPUTER VAULT."><br>
<img src="images/cranston/0137.png" width="400" alt="THE COMPUTER IS DEMOLISHED!! THERE IS A PLATINUM SPHERE HERE. YOU ARE IN THE COMPUTER VAULT.">

**Type `GET SPHERE`**

### 9. The caves

<img src="images/cranston/0138.png" width="400" alt="--------------- ENTER COMMAND?GET SPHERE YOU ARE IN THE COMPUTER VAULT.">

**Type `NORTH`**

<img src="images/cranston/0139.png" width="400" alt="THE DOOR CLOSED BEHIND YOU. YOU ARE FACING SOUTH IN FRONT OF A HEAVY BLAST DOOR.">

**Type `NORTH`**

<img src="images/cranston/0140.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN THE SOUTH CHAMBER OF THE E/W HALL. THERE&#39;S A DOOR TO THE SOUTH.">

**Type `NORTH`**

<img src="images/cranston/0141.png" width="400" alt="--------------- ENTER COMMAND?NORTH THIS IS THE EAST END OF A LARGE HALL. THERE IS A DOOR TO THE SOUTH.">

**Type `WEST`**

<img src="images/cranston/0142.png" width="400" alt="--------------- ENTER COMMAND?WEST THIS IS THE WEST END OF A LARGE HALL.THERE IS A DOOR TO THE SOUTH.">

**Type `NORTH`**

<img src="images/cranston/0143.png" width="400" alt="--------------- ENTER COMMAND?NORTH THERE ARE SOME GOLD BARS HERE. YOU ARE IN A SLOPING N/S CRAWLWAY.">

**Type `GET GOLD`**

<img src="images/cranston/0144.png" width="400" alt="YOU ARE IN A SLOPING N/S CRAWLWAY. --------------- ENTER COMMAND?GET GOLD YOU ARE IN A SLOPING N/S CRAWLWAY.">

**Type `NORTH`**

<img src="images/cranston/0145.png" width="400" alt="--------------- ENTER COMMAND?NORTH THIS ROOM HAS WALLS OF BLUE ROCK. THERE IS A PASSAGE TO THE SOUTH.">

**Type `EAST`**

<img src="images/cranston/0146.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN A ROUND ROOM WITH HOLES IN ALL DIRECTIONS.">

**Type `EAST`**

<img src="images/cranston/0147.png" width="400" alt="THERE IS A SAPPHIRE PENDANT HERE. YOU ARE AT AN INTERSECTION WITH SCRATCHING ON THE WALL.">

**Type `GET PENDANT`**

<img src="images/cranston/0148.png" width="400" alt="T YOU ARE AT AN INTERSECTION WITH SCRATCHING ON THE WALL.">

**Type `EAST`**

<img src="images/cranston/0149.png" width="400" alt="SUDDENLY AN ENORMOUS PINK BULL APPEARS. HE IS LOOKING RIGHT AT YOU! YOU ARE IN THE STALAGMITE CAVERN.">

**Type `LANTERN OFF`**

<img src="images/cranston/0150-1.png" width="400" alt="A WIZARD HAS PUT THE BULL IN A TEMPORARY STASIS FIELD WHICH WILL BE RELEASED AFTER YOUR NEXT COMMAND. IT IS PITCH BLACK."><br>
<img src="images/cranston/0150.png" width="400" alt="RELEASED AFTER YOUR NEXT COMMAND. IT IS PITCH BLACK. YOU ARE IN THE STALAGMITE CAVERN.">

**Type `EAST`**

<img src="images/cranston/0151.png" width="400" alt="IT IS PITCH BLACK. YOU ARE AT THE BASE OF A STEEP UPHILL SLOPE.">

**Type `LANTERN ON`**

<img src="images/cranston/0152.png" width="400" alt="OK. YOU ARE AT THE BASE OF A STEEP UPHILL SLOPE.">

**Type `EAST`**

<img src="images/cranston/0153.png" width="400" alt="--------------- ENTER COMMAND?EAST THERE IS HUGE GOLD NUGGET HERE. YOU ARE IN THE NUGGET ROOM.">

**Type `GET NUGGET`**

<img src="images/cranston/0154.png" width="400" alt="--------------- ENTER COMMAND?GET NUGGET YOU ARE IN THE NUGGET ROOM.">

**Type `WEST`**

<img src="images/cranston/0155.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE AT THE BASE OF A STEEP UPHILL SLOPE.">

**Type `DOWN`**

<img src="images/cranston/0156.png" width="400" alt="--------------- ENTER COMMAND?DOWN YOU ARE IN A ROOM WITH SNAKE HOLES.THERE IS A PASSAGE TO THE SOUTH.">

**Type `SOUTH`**

<img src="images/cranston/0157.png" width="400" alt="--------------- ENTER COMMAND?SOUTH THERE IS A NEST OF GOLD EGGS HERE. YOU&#39;RE IN A LARGE CHAMBER WITH A POOL.">

**Type `GET EGGS`**

<img src="images/cranston/0158.png" width="400" alt="YOU&#39;RE IN A LARGE CHAMBER WITH A POOL. --------------- ENTER COMMAND?GET EGGS YOU&#39;RE IN A LARGE CHAMBER WITH A POOL.">

**Type `EAST`**

<img src="images/cranston/0159.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN A ROOM WITH DEBRIS BLOCKING ONE PASSAGE.">

**Type `EAST`**

<img src="images/cranston/0160.png" width="400" alt="THE ROCKS MOVED OUT OF YOUR WAY! THERE IS A BEAUTIFUL JADE BUDDHA HERE. YOU ARE IN A WIDE N/S PASSAGE.">

**Type `GET JADE`**

<img src="images/cranston/0161.png" width="400" alt="YOU ARE IN A WIDE N/S PASSAGE. --------------- ENTER COMMAND?GET JADE YOU ARE IN A WIDE N/S PASSAGE.">

**Type `NORTH`**

<img src="images/cranston/0162.png" width="400" alt="YOU ARE IN A WIDE N/S PASSAGE. --------------- ENTER COMMAND?NORTH THIS ROOM HAS AN ENORMOUS STALACTITE.">

**Type `DOWN`**

<img src="images/cranston/0163.png" width="400" alt="--------------- ENTER COMMAND?DOWN YOU ARE IN A ROOM WITH SNAKE HOLES.THERE IS A PASSAGE TO THE SOUTH.">

**Type `WEST`**

<img src="images/cranston/0164.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE AT AN INTERSECTION WITH MUSHROOMS ON THE CEILING.">

**Type `WEST`**

<img src="images/cranston/0165.png" width="400" alt="--------------- ENTER COMMAND?WEST THIS IS THE WEST END OF A LARGE HALL.THERE IS A DOOR TO THE SOUTH.">

**Type `NORTH`**

<img src="images/cranston/0166.png" width="400" alt="HALL.THERE IS A DOOR TO THE SOUTH. --------------- ENTER COMMAND?NORTH YOU ARE IN A SLOPING N/S CRAWLWAY.">

**Type `NORTH`**

<img src="images/cranston/0167.png" width="400" alt="--------------- ENTER COMMAND?NORTH THIS ROOM HAS WALLS OF BLUE ROCK. THERE IS A PASSAGE TO THE SOUTH.">

**Type `WEST`**

<img src="images/cranston/0168.png" width="400" alt="YOU ARE IN THE LIFT CHAMBER. A COLUMN OF GREEN LIGHT RISES AND DISAPPEARS THROUGH THE CEILING.">

**Type `DROP NUGGET`**

<img src="images/cranston/0169.png" width="400" alt="YOU ARE IN THE LIFT CHAMBER. A COLUMN OF GREEN LIGHT RISES AND DISAPPEARS THROUGH THE CEILING.">

**Type `LIFT`**

### 10. The bedrooms and the oak tree

<img src="images/cranston/0170-1.png" width="400" alt="THERE IS A SUDDEN BRIGHTING OF GREEN LIGHT AND IT DISAPPEARS. YOU ARE IN THE LIFT CHAMBER. A COLUMN OF GREEN LIGHT RISES AND DISAPPEARS"><br>
<img src="images/cranston/0170.png" width="400" alt="YOU ARE IN THE LIFT CHAMBER. A COLUMN OF GREEN LIGHT RISES AND DISAPPEARS THROUGH THE CEILING.">

**Type `EAST`**

<img src="images/cranston/0171.png" width="400" alt="--------------- ENTER COMMAND?EAST THIS ROOM HAS WALLS OF BLUE ROCK. THERE IS A PASSAGE TO THE SOUTH.">

**Type `EAST`**

<img src="images/cranston/0172.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN A ROUND ROOM WITH HOLES IN ALL DIRECTIONS.">

**Type `UP`**

<img src="images/cranston/0173.png" width="400" alt="ALL DIRECTIONS. --------------- ENTER COMMAND?UP YOU ARE IN A WINDING E/W CORRIDOR.">

**Type `EAST`**

<img src="images/cranston/0174.png" width="400" alt="YOU ARE IN A WINDING E/W CORRIDOR. --------------- ENTER COMMAND?EAST YOU ARE AT THE EAST END OF A CORRIDOR.">

**Type `DOWN`**

<img src="images/cranston/0175.png" width="400" alt="YOU ARE AT THE EAST END OF A CORRIDOR. --------------- ENTER COMMAND?DOWN YOU ARE IN A LARGE DIRT ROOM.">

**Type `UP`**

<img src="images/cranston/0176.png" width="400" alt="--------------- ENTER COMMAND?UP THIS ROOM HAS A HOLE IN THE WALL AND THE FLOOR AND A DOOR TO THE SOUTH.">

**Type `SOUTH`**

<img src="images/cranston/0177.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE AT THE EAST END OF THE E/W HALL. THERE IS A DOOR TO THE SOUTH.">

**Type `SOUTH`**

<img src="images/cranston/0178.png" width="400" alt="THERE IS A HEAVY WOODEN CHAIR HERE. YOU ARE IN THE SERVANTS DINING ROOM. THERE IS A DOOR TO THE SOUTH.">

**Type `EAST`**

<img src="images/cranston/0179.png" width="400" alt="THERE IS A DOOR TO THE SOUTH. --------------- ENTER COMMAND?EAST YOU ARE IN AN EMPTY ROOM.">

**Type `UP`**

<img src="images/cranston/0180.png" width="400" alt="YOU ARE IN AN EMPTY ROOM. --------------- ENTER COMMAND?UP YOU ARE IN A SHABBY BED ROOM.">

**Type `WEST`**

<img src="images/cranston/0181.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN A LARGE SUN LIT ROOM. THERE IS A TORCH ON THE WALL.">

**Type `PULL TORCH`**

<img src="images/cranston/0182-1.png" width="400" alt="THE TORCH MOVES SLIGHTLY AND A PANEL SLIDES OPEN SILENTLY. THERE IS A SIGN HERE. A COLUMN OF GREEN LIGHT RISES FROM A"><br>
<img src="images/cranston/0182.png" width="400" alt="A COLUMN OF GREEN LIGHT RISES FROM A HOLE IN THE FLOOR. THERE IS A SIGN ON THE WALL AND AN EXIT TO THE SOUTH.">

**Type `GET NUGGET`**

<img src="images/cranston/0183.png" width="400" alt="A COLUMN OF GREEN LIGHT RISES FROM A HOLE IN THE FLOOR. THERE IS A SIGN ON THE WALL AND AN EXIT TO THE SOUTH.">

**Type `SOUTH`**

<img src="images/cranston/0184-1.png" width="400" alt="THE WALL CLOSED BEHIND YOU. THERE IS A PEARL RING HERE. YOU ARE IN THE MASTER BEDROOM. THERE IS A DOOR TO THE SOUTH AND A TORCH ON THE"><br>
<img src="images/cranston/0184.png" width="400" alt="YOU ARE IN THE MASTER BEDROOM. THERE IS A DOOR TO THE SOUTH AND A TORCH ON THE WALL.">

**Type `GET RING`**

<img src="images/cranston/0185.png" width="400" alt="YOU ARE IN THE MASTER BEDROOM. THERE IS A DOOR TO THE SOUTH AND A TORCH ON THE WALL.">

**Type `SOUTH`**

<img src="images/cranston/0186.png" width="400" alt="WALL. --------------- ENTER COMMAND?SOUTH YOU ARE IN THE CHILDREN&#39;S BEDROOM.">

**Type `EAST`**

<img src="images/cranston/0187.png" width="400" alt="YOU ARE IN THE CHILDREN&#39;S BEDROOM. --------------- ENTER COMMAND?EAST YOU ARE IN THE MIDDLE OF A N/S HALL.">

**Type `SOUTH`**

<img src="images/cranston/0188.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE AT THE SOUTH END OF THE UPSTAIRS HALL, STAIRS LEAD DOWN.">

**Type `WEST`**

<img src="images/cranston/0189.png" width="400" alt="--------------- ENTER COMMAND?WEST THERE IS AN INFLATABLE RAFT HERE. THIS IS THE CHILDREN&#39;S PLAYROOM.">

**Type `GET RAFT`**

<img src="images/cranston/0190.png" width="400" alt="THIS IS THE CHILDREN&#39;S PLAYROOM. --------------- ENTER COMMAND?GET RAFT THIS IS THE CHILDREN&#39;S PLAYROOM.">

**Type `EAST`**

<img src="images/cranston/0191.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE AT THE SOUTH END OF THE UPSTAIRS HALL, STAIRS LEAD DOWN.">

**Type `EAST`**

<img src="images/cranston/0192.png" width="400" alt="THERE ARE MIRRORS ON THE CEILING. THIS IS THE BRIDAL SUITE. THERE IS A DOOR TO THE SOUTH.">

**Type `OPEN DRAWER`**

<img src="images/cranston/0193-1.png" width="400" alt="OK. THERE IS A NECKLACE HERE. THIS IS THE BRIDAL SUITE. THERE IS A DOOR TO THE SOUTH."><br>
<img src="images/cranston/0193.png" width="400" alt="THERE IS A NECKLACE HERE. THIS IS THE BRIDAL SUITE. THERE IS A DOOR TO THE SOUTH.">

**Type `GET NECKLACE`**

<img src="images/cranston/0194.png" width="400" alt="CE THIS IS THE BRIDAL SUITE. THERE IS A DOOR TO THE SOUTH.">

**Type `SOUTH`**

<img src="images/cranston/0195.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE ON A BALCONY, FACING SOUTH. THERE IS A DOOR TO THE NORTH.">

**Type `JUMP`**

### 11. The emeralds of the fountain, and out by the gate

<img src="images/cranston/0196.png" width="400" alt="WITH AN INCREDIBLE LEAP, YOU CLEAR THE RAILING AND YOU ARE HANGING IN AN OAK TREE.">

**Type `DOWN`**

<img src="images/cranston/0197.png" width="400" alt="YOU ARE HANGING IN AN OAK TREE. --------------- ENTER COMMAND?DOWN YOU&#39;RE AT THE BASE OF A TALL OAK TREE.">

**Type `WEST`**

<img src="images/cranston/0198.png" width="400" alt="YOU&#39;RE AT THE BASE OF A TALL OAK TREE. --------------- ENTER COMMAND?WEST YOU ARE IN FRONT OF A SHED.">

**Type `NORTH`**

<img src="images/cranston/0199.png" width="400" alt="--------------- ENTER COMMAND?NORTH THERE IS A SMALL SCREWDRIVER HERE. YOU ARE INSIDE THE SHED.">

**Type `GET SCREWDRIVER`**

<img src="images/cranston/0200.png" width="400" alt="--------------- ENTER COMMAND?GET SCREWD RIVER YOU ARE INSIDE THE SHED.">

**Type `SOUTH`**

<img src="images/cranston/0201.png" width="400" alt="YOU ARE INSIDE THE SHED. --------------- ENTER COMMAND?SOUTH YOU ARE IN FRONT OF A SHED.">

**Type `WEST`**

<img src="images/cranston/0202.png" width="400" alt="YOU ARE IN FRONT OF A SHED. --------------- ENTER COMMAND?WEST YOU ARE IN THE WOODS.">

**Type `NORTH`**

<img src="images/cranston/0203.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN THE DRIVE WAY, A SMALL PATH LEADS SOUTH INTO A WOODS.">

**Type `NORTH`**

<img src="images/cranston/0204.png" width="400" alt="LEADS SOUTH INTO A WOODS. --------------- ENTER COMMAND?NORTH YOU ARE IN THE DRIVE WAY.">

**Type `WEST`**

<img src="images/cranston/0205.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN FRONT OF THE CAT FOUNTAIN. THERE IS A SIGN HERE.">

**Type `INFLATE RAFT`**

<img src="images/cranston/0206.png" width="400" alt="THE INFLATED RAFT IS TO BIG TO CARRY. YOU ARE IN FRONT OF THE CAT FOUNTAIN. THERE IS A SIGN HERE.">

**Type `NORTH`**

<img src="images/cranston/0207.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE ON THE PEDESTAL IN THE MIDDLE OF THE FOUNTAIN.">

**Type `USE SCREWDRIVER`**

<img src="images/cranston/0208-1.png" width="400" alt="THE SCREWDRIVER EASILY PRYS OUT THE EYES! THE EYES ARE EMERALDS!!! YOU ARE ON THE PEDESTAL IN THE MIDDLE OF THE FOUNTAIN."><br>
<img src="images/cranston/0208.png" width="400" alt="EYES! THE EYES ARE EMERALDS!!! YOU ARE ON THE PEDESTAL IN THE MIDDLE OF THE FOUNTAIN.">

**Type `SOUTH`**

<img src="images/cranston/0209.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN FRONT OF THE CAT FOUNTAIN. THERE IS A SIGN HERE.">

**Type `WEST`**

<img src="images/cranston/0210.png" width="400" alt="THERE IS A SIGN HERE. --------------- ENTER COMMAND?WEST YOU ARE IN THE GARDEN.">

**Type `SOUTH`**

<img src="images/cranston/0211.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN THE DRIVE WAY, NORTH OF THE MAIN GATE.">

**Type `UNLOCK GATE`**

<img src="images/cranston/0212.png" width="400" alt="OK. YOU ARE IN THE DRIVE WAY, NORTH OF THE MAIN GATE.">

**Type `OPEN GATE`**

<img src="images/cranston/0213.png" width="400" alt="OK. YOU ARE IN THE DRIVE WAY, NORTH OF THE MAIN GATE.">

**Type `SOUTH`**

### The end

<img src="images/cranston/0214-1.png" width="400" alt="CONGRATULATIONS!!!!!! YOU HAVE SUCCESSFULLY COMPLETED YOUR MISSION, AND ARE HEREBY DECLARED A LEVEL 3 ADVENTURER."><br>
<img src="images/cranston/0214.png" width="400" alt="CRANSTON MANOR. 03D9- A=04 X=FF Y=34 P=32 S=F3 *">

<!-- End of the walkthrough -->

## What next

[Mystery House](mysteryhouse.md) is another house to search, the first of
the Hi-Res Adventures.
