# Mystery House, from the start to the end

[Back to the activities](../README.md)

**Mystery House**, *Hi-Res Adventure #1* of On-Line Systems, came out in 1980
for the Apple \]\[, written by Ken and Roberta Williams. It was one of the
first adventures with pictures: each place of an old Victorian house drawn in
lines on the high resolution screen, with four lines of text under it. As
its instructions tell, you come into the house with seven other people, who
spread through it and start turning up dead: you have to find the killer
before the killer finds you, and there are jewels hidden in the house.

You type one or two words, `GO STAIRS`, `OPEN DOOR`, `TAKE NOTE`, and the
game answers under the picture. The whole game is on one side of a diskette.

This page plays it to its end, all 135 commands, with a picture of the screen
after each one: 136 pictures. The pictures are of a colour television, with
the four lines of text under them drawn white and sharp, as a monochrome
monitor showed them: the television blurred them into fringes of colour,
hard to read.

## What you need

**Mystery House (4am crack)**, from
[its page on the Internet Archive](https://archive.org/details/MysteryHouse4amCrack):
the zip `Mystery House (4am crack).zip`, with the disk
`Mystery House (4am crack).dsk`. `./fetch-disks.sh` in this repository
downloads it into `disks/` and checks it.

The commands of this page are also in this repository, a line each,
[listings/mysteryhouse.txt](listings/mysteryhouse.txt). They follow the
walkthrough of
[The Sierra Help Pages](https://sierrahelp.com/Walkthroughs/MysteryHouseWalkthrough.html),
without its `SAVE`s, which need a diskette of their own.

The [manual of Mystery House](https://archive.org/details/Mystery_House_Hi-Res_Adventure_1)
is on the Internet Archive.

## The machine

An Apple \]\[+:

- the 6502 processor at 1 MHz and 48 KB of memory;
- a Disk II controller card in slot 6, with Mystery House in drive 1.

```bash
izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s6 'diskii,disk1=disks/Mystery House (4am crack).dsk'
```

## Playing it

1. **Start izapple2** with the command above, and **type `G`** and Return,
   for the game; `I` would show its instructions.

   ![The title](images/mysteryhouse/title.png)

2. **Type the commands**, each followed by Return. Each step below is what
   the game shows, and the command to type then. When the picture stays the
   same and only the text goes on, the page shows the picture once, and
   under it the lines of text that came after.

## The walkthrough

<!-- The walkthrough, written by the generator -->

- [1. Into the house](#1-into-the-house)
- [2. A candle and matches](#2-a-candle-and-matches)
- [3. Water from the kitchen](#3-water-from-the-kitchen)
- [4. The library, and Sam the mechanic](#4-the-library-and-sam-the-mechanic)
- [5. Fire in the dining room](#5-fire-in-the-dining-room)
- [6. The bedrooms upstairs](#6-the-bedrooms-upstairs)
- [7. The chest](#7-the-chest)
- [8. The button in the study](#8-the-button-in-the-study)
- [9. Through the wall](#9-through-the-wall)
- [10. Daisy, in the tower](#10-daisy-in-the-tower)
- [11. The jewels](#11-the-jewels)
- [The end](#the-end)

### 1. Into the house

<img src="images/mysteryhouse/0001.png" width="400" alt="YOU ARE IN THE FRONT YARD OF A LARGE ABANDONED VICTORIAN HOUSE. STONE STEPS LEAD UP TO A WIDE PORCH">

**Type `GO STAIRS`**

<img src="images/mysteryhouse/0002.png" width="400" alt="--------------- ENTER COMMAND?GO STAIRS YOU ARE ON THE PORCH. STONE STEPS LEAD DOWN TO THE FRONT YARD">

**Type `OPEN DOOR`**

<img src="images/mysteryhouse/0003.png" width="400" alt="--------------- ENTER COMMAND?OPEN DOOR YOU ARE ON THE PORCH. STONE STEPS LEAD DOWN TO THE FRONT YARD">

**Type `GO DOOR`**

<img src="images/mysteryhouse/0004.png" width="400" alt="YOU ARE IN AN ENTRY HALL. DOORWAYS GO EAST, WEST AND SOUTH. A STAIRWAY GOES UP">

**Type `TAKE NOTE`**

<img src="images/mysteryhouse/0005.png" width="400" alt="YOU ARE IN AN ENTRY HALL. DOORWAYS GO EAST, WEST AND SOUTH. A STAIRWAY GOES UP">

**Type `READ NOTE`**

<img src="images/mysteryhouse/0006.png" width="400" alt="YOU ARE IN AN ENTRY HALL. DOORWAYS GO EAST, WEST AND SOUTH. A STAIRWAY GOES UP">

**Type `DROP NOTE`**

### 2. A candle and matches

<img src="images/mysteryhouse/0007.png" width="400" alt="YOU ARE IN AN ENTRY HALL. DOORWAYS GO EAST, WEST AND SOUTH. A STAIRWAY GOES UP">

**Type `SOUTH`**

<img src="images/mysteryhouse/0008.png" width="400" alt="UP --------------- ENTER COMMAND?SOUTH YOU ARE IN THE DINING ROOM">

**Type `GET CANDLE`**

<img src="images/mysteryhouse/0009.png" width="400" alt="--------------- ENTER COMMAND?GET CANDLE YOU ARE IN THE DINING ROOM">

**Type `NORTH`**

<img src="images/mysteryhouse/0010.png" width="400" alt="YOU ARE IN AN ENTRY HALL. DOORWAYS GO EAST, WEST AND SOUTH. A STAIRWAY GOES UP">

**Type `WEST`**

<img src="images/mysteryhouse/0011.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN THE KITCHEN. THERE IS A REFRIGERATOR, STOVE AND CABINET">

**Type `OPEN CABINET`**

<img src="images/mysteryhouse/0012.png" width="400" alt="ET YOU ARE IN THE KITCHEN. THERE IS A REFRIGERATOR, STOVE AND CABINET">

**Type `LOOK CABINET`**

<img src="images/mysteryhouse/0013.png" width="400" alt="ET YOU ARE IN THE KITCHEN. THERE IS A REFRIGERATOR, STOVE AND CABINET">

**Type `GET MATCHES`**

<img src="images/mysteryhouse/0014.png" width="400" alt="S YOU ARE IN THE KITCHEN. THERE IS A REFRIGERATOR, STOVE AND CABINET">

**Type `CLOSE CABINET`**

<img src="images/mysteryhouse/0015.png" width="400" alt="NET YOU ARE IN THE KITCHEN. THERE IS A REFRIGERATOR, STOVE AND CABINET">

**Type `LIGHT CANDLE`**

<img src="images/mysteryhouse/0016.png" width="400" alt="OK YOU ARE IN THE KITCHEN. THERE IS A REFRIGERATOR, STOVE AND CABINET">

**Type `OPEN REFRIGERATOR`**

<img src="images/mysteryhouse/0017.png" width="400" alt="GERATOR YOU ARE IN THE KITCHEN. THERE IS A REFRIGERATOR, STOVE AND CABINET">

**Type `LOOK REFRIGERATOR`**

<img src="images/mysteryhouse/0018.png" width="400" alt="GERATOR YOU ARE IN THE KITCHEN. THERE IS A REFRIGERATOR, STOVE AND CABINET">

**Type `TAKE PITCHER`**

<img src="images/mysteryhouse/0019.png" width="400" alt="ER YOU ARE IN THE KITCHEN. THERE IS A REFRIGERATOR, STOVE AND CABINET">

**Type `CLOSE REFRIGERATOR`**

### 3. Water from the kitchen

<img src="images/mysteryhouse/0020.png" width="400" alt="IGERATOR YOU ARE IN THE KITCHEN. THERE IS A REFRIGERATOR, STOVE AND CABINET">

**Type `LOOK SINK`**

<img src="images/mysteryhouse/0021.png" width="400" alt="THERE IS A BUTTERKNIFE HERE YOU ARE IN THE KITCHEN. THERE IS A REFRIGERATOR, STOVE AND CABINET">

**Type `GET BUTTERKNIFE`**

<img src="images/mysteryhouse/0022.png" width="400" alt="IT IS GETTING DARK YOU ARE IN THE KITCHEN. THERE IS A REFRIGERATOR, STOVE AND CABINET">

**Type `WATER ON`**

<img src="images/mysteryhouse/0023.png" width="400" alt="IT IS GETTING DARK YOU ARE IN THE KITCHEN. THERE IS A REFRIGERATOR, STOVE AND CABINET">

**Type `GET WATER`**

<img src="images/mysteryhouse/0024.png" width="400" alt="IT IS GETTING DARK YOU ARE IN THE KITCHEN. THERE IS A REFRIGERATOR, STOVE AND CABINET">

**Type `EAST`**

### 4. The library, and Sam the mechanic

<img src="images/mysteryhouse/0025.png" width="400" alt="YOU ARE IN AN ENTRY HALL. DOORWAYS GO EAST, WEST AND SOUTH. A STAIRWAY GOES UP">

**Type `EAST`**

<img src="images/mysteryhouse/0026.png" width="400" alt="--------------- ENTER COMMAND?EAST IT IS GETTING DARK YOU ARE IN THE OLD, DUSTY LIBRARY">

**Type `TAKE NOTE`**

<img src="images/mysteryhouse/0027.png" width="400" alt="--------------- ENTER COMMAND?TAKE NOTE IT IS GETTING DARK YOU ARE IN THE OLD, DUSTY LIBRARY">

**Type `READ NOTE`**

<img src="images/mysteryhouse/0028.png" width="400" alt="--------------- ENTER COMMAND?READ NOTE IT IS GETTING DARK YOU ARE IN THE OLD, DUSTY LIBRARY">

**Type `DROP NOTE`**

<img src="images/mysteryhouse/0029.png" width="400" alt="--------------- ENTER COMMAND?DROP NOTE IT IS GETTING DARK YOU ARE IN THE OLD, DUSTY LIBRARY">

**Type `OPEN DOOR`**

<img src="images/mysteryhouse/0030.png" width="400" alt="--------------- ENTER COMMAND?OPEN DOOR IT IS GETTING DARK YOU ARE IN THE OLD, DUSTY LIBRARY">

**Type `EAST`**

<img src="images/mysteryhouse/0031.png" width="400" alt="IT IS GETTING DARK YOU ARE IN THE SIDE YARD. YOU CAN FOLLOW THE FENCE TO THE SOUTH">

**Type `SOUTH`**

<img src="images/mysteryhouse/0032.png" width="400" alt="YOU ARE IN THE FENCED BACK YARD. THE FENCE FOLLOWS THE SIDE OF THE HOUSE TO THE NORTH. THERE IS A DEAD BODY HERE">

**Type `LOOK BODY`**

### 5. Fire in the dining room

<img src="images/mysteryhouse/0033.png" width="400" alt="YOU ARE IN THE FENCED BACK YARD. THE FENCE FOLLOWS THE SIDE OF THE HOUSE TO THE NORTH. THERE IS A DEAD BODY HERE">

**Type `OPEN DOOR`**

<img src="images/mysteryhouse/0034.png" width="400" alt="YOU ARE IN THE FENCED BACK YARD. THE FENCE FOLLOWS THE SIDE OF THE HOUSE TO THE NORTH. THERE IS A DEAD BODY HERE">

**Type `GO DOOR`**

<img src="images/mysteryhouse/0035.png" width="400" alt="--------------- ENTER COMMAND?GO DOOR IT IS GETTING DARK YOU ARE IN THE DINING ROOM">

**Type `LOOK ROOM`**

<img src="images/mysteryhouse/0036.png" width="400" alt="STARTED A FIRE WITH YOUR CANDLE! IT IS GETTING DARK YOU ARE IN THE DINING ROOM">

**Type `POUR WATER`**

<img src="images/mysteryhouse/0037.png" width="400" alt="THE PITCHER IS EMPTY IT IS GETTING DARK YOU ARE IN THE DINING ROOM">

**Type `LOOK HOLE`**

<img src="images/mysteryhouse/0038.png" width="400" alt="YOU ARE IN THE DINING ROOM --------------- ENTER COMMAND?LOOK HOLE YOU ARE IN THE DINING ROOM">

**Type `TAKE KEY`**

<img src="images/mysteryhouse/0039.png" width="400" alt="YOU ARE IN THE DINING ROOM --------------- ENTER COMMAND?TAKE KEY YOU ARE IN THE DINING ROOM">

**Type `NORTH`**

### 6. The bedrooms upstairs

<img src="images/mysteryhouse/0040.png" width="400" alt="YOU ARE IN AN ENTRY HALL. DOORWAYS GO EAST, WEST AND SOUTH. A STAIRWAY GOES UP">

**Type `GO STAIRS`**

<img src="images/mysteryhouse/0041.png" width="400" alt="--------------- ENTER COMMAND?GO STAIRS YOU ARE AT THE JUNCTION OF AN EAST/WEST HALLWAY AND A NORTH/SOUTH HALLWAY">

**Type `WEST`**

<img src="images/mysteryhouse/0042.png" width="400" alt="HALLWAY AND A NORTH/SOUTH HALLWAY --------------- ENTER COMMAND?WEST THERE IS A DOORWAY HERE">

**Type `GO DOOR`**

<img src="images/mysteryhouse/0043.png" width="400" alt="--------------- ENTER COMMAND?GO DOOR YOU ARE IN AN OLD NURSERY. THERE IS A DEAD BODY HERE.">

**Type `LOOK BODY`**

<img src="images/mysteryhouse/0044.png" width="400" alt="BEEN STABBED YOU ARE IN AN OLD NURSERY. THERE IS A DEAD BODY HERE.">

**Type `GO DOOR`**

<img src="images/mysteryhouse/0045.png" width="400" alt="DEAD BODY HERE. --------------- ENTER COMMAND?GO DOOR THERE IS A DOORWAY HERE">

**Type `WEST`**

<img src="images/mysteryhouse/0046.png" width="400" alt="THERE IS A DOORWAY HERE --------------- ENTER COMMAND?WEST THERE IS A DOORWAY HERE">

**Type `GO DOOR`**

<img src="images/mysteryhouse/0047.png" width="400" alt="THERE IS A DOORWAY HERE --------------- ENTER COMMAND?GO DOOR YOU ARE IN A BOYS BEDROOM">

**Type `TAKE NOTE`**

<img src="images/mysteryhouse/0048.png" width="400" alt="YOU ARE IN A BOYS BEDROOM --------------- ENTER COMMAND?TAKE NOTE YOU ARE IN A BOYS BEDROOM">

**Type `READ NOTE`**

<img src="images/mysteryhouse/0049.png" width="400" alt="YOU ARE IN A BOYS BEDROOM --------------- ENTER COMMAND?READ NOTE YOU ARE IN A BOYS BEDROOM">

**Type `DROP NOTE`**

<img src="images/mysteryhouse/0050.png" width="400" alt="YOU ARE IN A BOYS BEDROOM --------------- ENTER COMMAND?DROP NOTE YOU ARE IN A BOYS BEDROOM">

**Type `GO DOOR`**

<img src="images/mysteryhouse/0051.png" width="400" alt="YOU ARE IN A BOYS BEDROOM --------------- ENTER COMMAND?GO DOOR THERE IS A DOORWAY HERE">

**Type `EAST`**

<img src="images/mysteryhouse/0052.png" width="400" alt="THERE IS A DOORWAY HERE --------------- ENTER COMMAND?EAST THERE IS A DOORWAY HERE">

**Type `EAST`**

<img src="images/mysteryhouse/0053.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE AT THE JUNCTION OF AN EAST/WEST HALLWAY AND A NORTH/SOUTH HALLWAY">

**Type `EAST`**

<img src="images/mysteryhouse/0054.png" width="400" alt="HALLWAY AND A NORTH/SOUTH HALLWAY --------------- ENTER COMMAND?EAST THERE IS A DOORWAY HERE">

**Type `GO DOOR`**

<img src="images/mysteryhouse/0055.png" width="400" alt="A DAGGER IS THROWN AT YOU FROM OUTSIDE THE ROOM. IT MISSES! YOU ARE IN A LARGE BEDROOM">

**Type `TAKE DAGGER`**

<img src="images/mysteryhouse/0056.png" width="400" alt="--------------- ENTER COMMAND?TAKE DAGGE R YOU ARE IN A LARGE BEDROOM">

**Type `GO DOOR`**

<img src="images/mysteryhouse/0057.png" width="400" alt="YOU ARE IN A LARGE BEDROOM --------------- ENTER COMMAND?GO DOOR THERE IS A DOORWAY HERE">

**Type `EAST`**

<img src="images/mysteryhouse/0058.png" width="400" alt="THERE IS A DOORWAY HERE --------------- ENTER COMMAND?EAST THERE IS A DOORWAY HERE">

**Type `GO DOOR`**

<img src="images/mysteryhouse/0059.png" width="400" alt="--------------- ENTER COMMAND?GO DOOR YOU ARE IN A SMALL BEDROOM. THERE IS A DEAD BODY HERE.">

**Type `LOOK BODY`**

<img src="images/mysteryhouse/0060.png" width="400" alt="BLOND HAIR ON HER DRESS YOU ARE IN A SMALL BEDROOM. THERE IS A DEAD BODY HERE.">

**Type `GO DOOR`**

<img src="images/mysteryhouse/0061.png" width="400" alt="DEAD BODY HERE. --------------- ENTER COMMAND?GO DOOR THERE IS A DOORWAY HERE">

**Type `WEST`**

<img src="images/mysteryhouse/0062.png" width="400" alt="THERE IS A DOORWAY HERE --------------- ENTER COMMAND?WEST THERE IS A DOORWAY HERE">

**Type `WEST`**

### 7. The chest

<img src="images/mysteryhouse/0063.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE AT THE JUNCTION OF AN EAST/WEST HALLWAY AND A NORTH/SOUTH HALLWAY">

**Type `NORTH`**

<img src="images/mysteryhouse/0064.png" width="400" alt="HALLWAY AND A NORTH/SOUTH HALLWAY --------------- ENTER COMMAND?NORTH YOU ARE AT A STAIRWAY">

**Type `GO STAIRS`**

<img src="images/mysteryhouse/0065.png" width="400" alt="YOU ARE AT A STAIRWAY --------------- ENTER COMMAND?GO STAIRS YOU ARE IN THE ATTIC">

**Type `TAKE HAMMER`**

<img src="images/mysteryhouse/0066.png" width="400" alt="--------------- ENTER COMMAND?TAKE HAMME R YOU ARE IN THE ATTIC">

**Type `GO DOOR`**

<img src="images/mysteryhouse/0067.png" width="400" alt="YOU ARE IN THE ATTIC --------------- ENTER COMMAND?GO DOOR YOU ARE IN A STORAGE ROOM">

**Type `UNLOCK CHEST`**

<img src="images/mysteryhouse/0068.png" width="400" alt="ST OK YOU ARE IN A STORAGE ROOM">

**Type `OPEN CHEST`**

<img src="images/mysteryhouse/0069.png" width="400" alt="--------------- ENTER COMMAND?OPEN CHEST YOU ARE IN A STORAGE ROOM">

**Type `LOOK CHEST`**

<img src="images/mysteryhouse/0070.png" width="400" alt="--------------- ENTER COMMAND?LOOK CHEST YOU ARE IN A STORAGE ROOM">

**Type `GET GUN`**

### 8. The button in the study

<img src="images/mysteryhouse/0071.png" width="400" alt="YOU ARE IN A STORAGE ROOM --------------- ENTER COMMAND?GET GUN YOU ARE IN A STORAGE ROOM">

**Type `GO DOOR`**

<img src="images/mysteryhouse/0072.png" width="400" alt="YOU ARE IN A STORAGE ROOM --------------- ENTER COMMAND?GO DOOR YOU ARE IN THE ATTIC">

**Type `GO STAIRS`**

<img src="images/mysteryhouse/0073.png" width="400" alt="YOU ARE IN THE ATTIC --------------- ENTER COMMAND?GO STAIRS YOU ARE AT A STAIRWAY">

**Type `NORTH`**

<img src="images/mysteryhouse/0074.png" width="400" alt="YOU ARE AT A STAIRWAY --------------- ENTER COMMAND?NORTH THERE IS A DOORWAY HERE">

**Type `GO DOOR`**

<img src="images/mysteryhouse/0075.png" width="400" alt="THERE IS A DOORWAY HERE --------------- ENTER COMMAND?GO DOOR YOU ARE IN THE STUDY">

**Type `USE BUTTERKNIFE`**

<img src="images/mysteryhouse/0076.png" width="400" alt="KNIFE THE PICTURE IS LOOSE YOU ARE IN THE STUDY">

**Type `TAKE PICTURE`**

<img src="images/mysteryhouse/0077.png" width="400" alt="RE THERE IS A BUTTON ON THE WALL YOU ARE IN THE STUDY">

**Type `PRESS BUTTON`**

<img src="images/mysteryhouse/0078.png" width="400" alt="ON PART OF THE WALL OPENS YOU ARE IN THE STUDY">

**Type `EAST`**

<img src="images/mysteryhouse/0079.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN THE BATHROOM. THERE IS A DEAD BODY HERE.">

**Type `LOOK BODY`**

<img src="images/mysteryhouse/0080.png" width="400" alt="STRANGLED WITH A PAIR OF PANTYHOSE YOU ARE IN THE BATHROOM. THERE IS A DEAD BODY HERE.">

**Type `TAKE TOWEL`**

<img src="images/mysteryhouse/0081.png" width="400" alt="YOU ARE IN THE BATHROOM. THERE IS A DEAD BODY HERE.">

**Type `GO DOOR`**

<img src="images/mysteryhouse/0082.png" width="400" alt="DEAD BODY HERE. --------------- ENTER COMMAND?GO DOOR YOU ARE IN THE STUDY">

**Type `NORTH`**

<img src="images/mysteryhouse/0083.png" width="400" alt="YOU ARE IN THE STUDY --------------- ENTER COMMAND?NORTH THERE IS A DOORWAY HERE">

**Type `SOUTH`**

<img src="images/mysteryhouse/0084.png" width="400" alt="THERE IS A DOORWAY HERE --------------- ENTER COMMAND?SOUTH YOU ARE AT A STAIRWAY">

**Type `SOUTH`**

<img src="images/mysteryhouse/0085.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE AT THE JUNCTION OF AN EAST/WEST HALLWAY AND A NORTH/SOUTH HALLWAY">

**Type `GO STAIRS`**

### 9. Through the wall

<img src="images/mysteryhouse/0086.png" width="400" alt="YOU ARE IN AN ENTRY HALL. DOORWAYS GO EAST, WEST AND SOUTH. A STAIRWAY GOES UP">

**Type `WEST`**

<img src="images/mysteryhouse/0087.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN THE KITCHEN. THERE IS A REFRIGERATOR, STOVE AND CABINET">

**Type `MOVE CABINET`**

<img src="images/mysteryhouse/0088.png" width="400" alt="THE WALL IS BRICKED UP BEHIND IT YOU ARE IN THE KITCHEN. THERE IS A REFRIGERATOR, STOVE AND CABINET">

**Type `BREAK WALL`**

<img src="images/mysteryhouse/0089.png" width="400" alt="HOLE YOU ARE IN THE KITCHEN. THERE IS A REFRIGERATOR, STOVE AND CABINET">

**Type `GO HOLE`**

<img src="images/mysteryhouse/0090.png" width="400" alt="REFRIGERATOR, STOVE AND CABINET --------------- ENTER COMMAND?GO HOLE YOU ARE IN A SMALL PANTRY">

**Type `GO STAIRS`**

<img src="images/mysteryhouse/0091.png" width="400" alt="--------------- ENTER COMMAND?GO STAIRS YOU ARE AT THE NORTH END OF A NARROW NORTH/SOUTH PASSAGEWAY">

**Type `SOUTH`**

<img src="images/mysteryhouse/0092.png" width="400" alt="YOU ARE IN A MOIST BASEMENT. ALGAE COVERS THE WALLS. THERE IS A DEAD BODY HERE.">

**Type `LOOK BODY`**

<img src="images/mysteryhouse/0093.png" width="400" alt="YOU ARE IN A MOIST BASEMENT. ALGAE COVERS THE WALLS. THERE IS A DEAD BODY HERE.">

**Type `TAKE KEY`**

<img src="images/mysteryhouse/0094.png" width="400" alt="YOU ARE IN A MOIST BASEMENT. ALGAE COVERS THE WALLS. THERE IS A DEAD BODY HERE.">

**Type `GO HOLE`**

<img src="images/mysteryhouse/0095.png" width="400" alt="--------------- ENTER COMMAND?GO HOLE YOU ARE AT THE SOUTH END OF A NORTH/SOUTH TUNNEL">

**Type `NORTH`**

<img src="images/mysteryhouse/0096.png" width="400" alt="--------------- ENTER COMMAND?NORTH THERE IS A VERY TALL PINE TREE IN FRONT OF YOU">

**Type `UP`**

<img src="images/mysteryhouse/0097.png" width="400" alt="--------------- ENTER COMMAND?UP YOU ARE AT THE TOP OF A VERY TALL PINE TREE">

**Type `LOOK TELESCOPE`**

<img src="images/mysteryhouse/0098.png" width="400" alt="ATTIC CEILING YOU ARE AT THE TOP OF A VERY TALL PINE TREE">

**Type `DOWN`**

<img src="images/mysteryhouse/0099.png" width="400" alt="--------------- ENTER COMMAND?DOWN THERE IS A VERY TALL PINE TREE IN FRONT OF YOU">

**Type `DOWN`**

<img src="images/mysteryhouse/0100.png" width="400" alt="OF YOU --------------- ENTER COMMAND?DOWN YOU ARE IN A FOREST">

**Type `NORTH`**

<img src="images/mysteryhouse/0101.png" width="400" alt="YOU ARE IN A FOREST --------------- ENTER COMMAND?NORTH YOU ARE IN A FOREST">

**Type `NORTH`**

<img src="images/mysteryhouse/0102.png" width="400" alt="YOU ARE IN A FOREST --------------- ENTER COMMAND?NORTH YOU ARE IN A FOREST">

**Type `NORTH`**

<img src="images/mysteryhouse/0103.png" width="400" alt="YOU ARE IN A FOREST --------------- ENTER COMMAND?NORTH YOU ARE IN A FOREST">

**Type `NORTH`**

<img src="images/mysteryhouse/0104.png" width="400" alt="YOU ARE IN A FOREST --------------- ENTER COMMAND?NORTH YOU ARE IN A FOREST">

**Type `UP`**

<img src="images/mysteryhouse/0105.png" width="400" alt="--------------- ENTER COMMAND?UP THE KITCHEN DOOR IS CLOSED YOU ARE IN A FOREST">

**Type `OPEN DOOR`**

<img src="images/mysteryhouse/0106.png" width="400" alt="--------------- ENTER COMMAND?OPEN DOOR THE KITCHEN DOOR IS OPEN YOU ARE IN A FOREST">

**Type `UP`**

### 10. Daisy, in the tower

<img src="images/mysteryhouse/0107.png" width="400" alt="--------------- ENTER COMMAND?UP YOU ARE IN THE KITCHEN. THERE IS A REFRIGERATOR, STOVE AND CABINET">

**Type `EAST`**

<img src="images/mysteryhouse/0108.png" width="400" alt="YOU ARE IN AN ENTRY HALL. DOORWAYS GO EAST, WEST AND SOUTH. A STAIRWAY GOES UP">

**Type `GO STAIRS`**

<img src="images/mysteryhouse/0109.png" width="400" alt="--------------- ENTER COMMAND?GO STAIRS YOU ARE AT THE JUNCTION OF AN EAST/WEST HALLWAY AND A NORTH/SOUTH HALLWAY">

**Type `NORTH`**

<img src="images/mysteryhouse/0110.png" width="400" alt="HALLWAY AND A NORTH/SOUTH HALLWAY --------------- ENTER COMMAND?NORTH YOU ARE AT A STAIRWAY">

**Type `GO STAIRS`**

<img src="images/mysteryhouse/0111.png" width="400" alt="YOU ARE AT A STAIRWAY --------------- ENTER COMMAND?GO STAIRS YOU ARE IN THE ATTIC">

**Type `OPEN TRAPDOOR`**

<img src="images/mysteryhouse/0112.png" width="400" alt="--------------- ENTER COMMAND?OPEN TRAPD OOR YOU ARE IN THE ATTIC">

**Type `CLIMB LADDER`**

<img src="images/mysteryhouse/0113.png" width="400" alt="--------------- ENTER COMMAND?CLIMB LADD ER YOU ARE IN THE TOWER">

**Type `KILL DAISY`**

<img src="images/mysteryhouse/0114.png" width="400" alt="WITH WHAT? YOU ARE IN THE TOWER">

**Type `WITH GUN`**

<img src="images/mysteryhouse/0115.png" width="400" alt="YOUR GUN IS EMPTY DAISY IS NOW DEAD YOU ARE IN THE TOWER">

**Type `TAKE NOTE`**

<img src="images/mysteryhouse/0116.png" width="400" alt="--------------- ENTER COMMAND?TAKE NOTE OK YOU ARE IN THE TOWER">

**Type `READ NOTE`**

<img src="images/mysteryhouse/0117.png" width="400" alt="YOU ARE IN THE TOWER --------------- ENTER COMMAND?READ NOTE YOU ARE IN THE TOWER">

**Type `DROP NOTE`**

### 11. The jewels

<img src="images/mysteryhouse/0118.png" width="400" alt="YOU ARE IN THE TOWER --------------- ENTER COMMAND?DROP NOTE YOU ARE IN THE TOWER">

**Type `GO TRAPDOOR`**

<img src="images/mysteryhouse/0119.png" width="400" alt="--------------- ENTER COMMAND?GO TRAPDOO R YOU ARE IN THE ATTIC">

**Type `GO STAIRS`**

<img src="images/mysteryhouse/0120.png" width="400" alt="YOU ARE IN THE ATTIC --------------- ENTER COMMAND?GO STAIRS YOU ARE AT A STAIRWAY">

**Type `NORTH`**

<img src="images/mysteryhouse/0121.png" width="400" alt="YOU ARE AT A STAIRWAY --------------- ENTER COMMAND?NORTH THERE IS A DOORWAY HERE">

**Type `GO DOOR`**

<img src="images/mysteryhouse/0122.png" width="400" alt="THERE IS A DOORWAY HERE --------------- ENTER COMMAND?GO DOOR YOU ARE IN THE STUDY">

**Type `GO WALL`**

<img src="images/mysteryhouse/0123.png" width="400" alt="--------------- ENTER COMMAND?GO WALL THE WALL CLOSES BEHIND YOU WITH A BANG YOU ARE IN A MUSTY CRAWLSPACE">

**Type `DOWN`**

<img src="images/mysteryhouse/0124.png" width="400" alt="YOU ARE IN A MUSTY CRAWLSPACE --------------- ENTER COMMAND?DOWN YOU ARE ON A STAIRWAY">

**Type `DOWN`**

<img src="images/mysteryhouse/0125.png" width="400" alt="YOU ARE IN A MOIST BASEMENT. ALGAE COVERS THE WALLS. THERE IS A DEAD BODY HERE.">

**Type `REMOVE ALGAE`**

<img src="images/mysteryhouse/0126.png" width="400" alt="YOU ARE IN A MOIST BASEMENT. ALGAE COVERS THE WALLS. THERE IS A DEAD BODY HERE.">

**Type `TAKE BRICK`**

<img src="images/mysteryhouse/0127.png" width="400" alt="YOU ARE IN A MOIST BASEMENT. ALGAE COVERS THE WALLS. THERE IS A DEAD BODY HERE.">

**Type `GET JEWELS`**

<img src="images/mysteryhouse/0128.png" width="400" alt="YOU ARE IN A MOIST BASEMENT. ALGAE COVERS THE WALLS. THERE IS A DEAD BODY HERE.">

**Type `NORTH`**

<img src="images/mysteryhouse/0129.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE AT THE NORTH END OF A NARROW NORTH/SOUTH PASSAGEWAY">

**Type `GO STAIRS`**

<img src="images/mysteryhouse/0130.png" width="400" alt="NORTH/SOUTH PASSAGEWAY --------------- ENTER COMMAND?GO STAIRS YOU ARE IN A SMALL PANTRY">

**Type `GO HOLE`**

<img src="images/mysteryhouse/0131.png" width="400" alt="--------------- ENTER COMMAND?GO HOLE YOU ARE IN THE KITCHEN. THERE IS A REFRIGERATOR, STOVE AND CABINET">

**Type `EAST`**

<img src="images/mysteryhouse/0132.png" width="400" alt="YOU ARE IN AN ENTRY HALL. DOORWAYS GO EAST, WEST AND SOUTH. A STAIRWAY GOES UP">

**Type `UNLOCK DOOR`**

<img src="images/mysteryhouse/0133.png" width="400" alt="YOU ARE IN AN ENTRY HALL. DOORWAYS GO EAST, WEST AND SOUTH. A STAIRWAY GOES UP">

**Type `OPEN DOOR`**

<img src="images/mysteryhouse/0134.png" width="400" alt="YOU ARE IN AN ENTRY HALL. DOORWAYS GO EAST, WEST AND SOUTH. A STAIRWAY GOES UP">

**Type `NORTH`**

<img src="images/mysteryhouse/0135.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE ON THE PORCH. STONE STEPS LEAD DOWN TO THE FRONT YARD">

**Type `GO STAIRS`**

### The end

<img src="images/mysteryhouse/0136.png" width="400" alt="CONGRATULATIONS YOU HAVE BEATEN ADVENTURE AND ARE DECLARED A GURU WIZARD WOULD YOU LIKE TO PLAY AGAIN">

<!-- End of the walkthrough -->

## What next

[Time Zone](timezone.md) is the fifth of the Hi-Res Adventures, the biggest,
on twelve sides of diskettes.
