# Ulysses and the Golden Fleece, from the start to the end

[Back to the activities](../README.md)

**Ulysses and the Golden Fleece**, *Hi-Res Adventure #4* of On-Line Systems,
by Bob Davis and Ken Williams, here in its version 1.1 of 1982. The king of
a small town of ancient Greece sends Ulysses for a legendary fleece of gold,
far to the north, with a ship and a crew: on the way, the Island of Storms
and Pluto, the god of the underworld, under it, Neptune, the sirens, the Cyclops, harpies and a
band of skeletons, and Pegasus.

You type one or two words, `HIRE CREW`, `CAST OFF`, `KILL SHEEP`, and the
game answers under the picture. The game is on the two sides of a diskette:
it starts from the first and asks for the other.

This page plays it to its end, all 216 commands, with a picture of the screen
after each one: 315 pictures. The pictures are of a colour television, with
the four lines of text under them drawn white and sharp, as a monochrome
monitor showed them: the television blurred them into fringes of colour,
hard to read.

## What you need

**Ulysses and the Golden Fleece v1.1 (4am crack)**, from
[its page on the Internet Archive](https://archive.org/details/UlyssesAndTheGoldenFleece_v11_4amCrack):
the zip `Ulysses and the Golden Fleece v1.1 (4am crack).zip`, with a file for
each side, `Ulysses and the Golden Fleece v1.1 (4am crack) side A.dsk` and
`Ulysses and the Golden Fleece v1.1 (4am crack) side B.dsk`.
`./fetch-disks.sh` in this repository downloads them into `disks/` and
checks them.

The commands of this page are also in this repository, a line each,
[listings/ulysses.txt](listings/ulysses.txt). They follow the walkthrough of
[The Sierra Help Pages](https://sierrahelp.com/Walkthroughs/UlyssesWalkthrough.html),
with what the machine showed it needs: `TALK GUARD` and `YES` to the guard;
`YES` to the storekeeper; the chest taken back from the tavern, where the
walkthrough leaves it; the note taken from the bottle with `GET BOTTLE`;
`TIE ROPE` and `TO ME` at the sirens; and `SAY SEVENSEAS` to scare off the
harpies, which
[a solution on AtariArchives.org](https://www.atariarchives.org/cfn/12/02/0068.php)
gives.

The [manual of Ulysses and the Golden Fleece](https://archive.org/details/ulyssesandthegoldenfleeceonlinesystems)
is on the Internet Archive.

## The machine

An Apple \]\[+:

- the 6502 processor at 1 MHz and 48 KB of memory;
- a Disk II controller card in slot 6, with side A in drive 1.

```bash
izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s6 'diskii,disk1=disks/Ulysses and the Golden Fleece v1.1 (4am crack) side A.dsk'
```

## Playing it

1. **Start izapple2** with the command above. The game asks to flip the
   diskette over: drop the file of side B on the area of drive 1 of the
   window of izapple2, and press a key. The page says each time it asks
   again.

2. **Type the commands**, each followed by Return. Each step below is what
   the game shows, and the command to type then. When the picture stays the
   same and only the text goes on, the page shows the picture once, and
   under it the lines of text that came after. When the game stops in the
   middle of a longer text, **press Return** to go on.

## The walkthrough

<!-- The walkthrough, written by the generator -->

- [1. A chest in the forest, and the king](#1-a-chest-in-the-forest-and-the-king)
- [2. The store](#2-the-store)
- [3. The tavern and the docks](#3-the-tavern-and-the-docks)
- [4. At sea](#4-at-sea)
- [5. The Island of Storms](#5-the-island-of-storms)
- [6. Under the island: the fjord, the dragon and the wall of fire](#6-under-the-island-the-fjord-the-dragon-and-the-wall-of-fire)
- [7. Neptune and the sirens](#7-neptune-and-the-sirens)
- [8. The island of the Cyclops](#8-the-island-of-the-cyclops)
- [9. The skeletons, the cliff and Pegasus](#9-the-skeletons-the-cliff-and-pegasus)
- [The end](#the-end)

### 1. A chest in the forest, and the king

<img src="images/ulysses/0001-1.png" width="400" alt="PLEASE FLIP DISK OVER AND PRESS ANY KEY"><br>
<img src="images/ulysses/0001.png" width="400" alt="YOU ARE ON A 3-WAY ROAD IN A SMALL TOWN IN ANCIENT GREECE. THERE IS A STORE TO THE WEST WITH A FENCE NEXT TO IT.">

*Side B: drop `Ulysses and the Golden Fleece v1.1 (4am crack) side B.dsk` on drive 1, and press Return.*

**Type `EAST`**

<img src="images/ulysses/0002.png" width="400" alt="THE WEST WITH A FENCE NEXT TO IT. --------------- ENTER COMMAND?EAST YOU ARE IN A LUSH GREEN FOREST.">

**Type `SOUTH`**

<img src="images/ulysses/0003.png" width="400" alt="YOU ARE IN A LUSH GREEN FOREST. --------------- ENTER COMMAND?SOUTH YOU ARE IN A FOREST.">

**Type `SOUTH`**

<img src="images/ulysses/0004.png" width="400" alt="YOU ARE IN A FOREST. --------------- ENTER COMMAND?SOUTH YOU ARE IN A FOREST.">

**Type `EAST`**

<img src="images/ulysses/0005.png" width="400" alt="YOU ARE IN A FOREST. --------------- ENTER COMMAND?EAST YOU ARE IN A LUSH GREEN FOREST.">

**Type `EAST`**

<img src="images/ulysses/0006.png" width="400" alt="YOU ARE IN A LUSH GREEN FOREST. --------------- ENTER COMMAND?EAST YOU ARE WANDERING IN A FOREST.">

**Type `NORTH`**

<img src="images/ulysses/0007.png" width="400" alt="--------------- ENTER COMMAND?NORTH THERE IS A CHEST HERE. YOU ARE LOST IN A FOREST.">

**Type `GET CHEST`**

<img src="images/ulysses/0008.png" width="400" alt="YOU ARE LOST IN A FOREST. --------------- ENTER COMMAND?GET CHEST YOU ARE LOST IN A FOREST.">

**Type `WEST`**

<img src="images/ulysses/0009.png" width="400" alt="YOU ARE LOST IN A FOREST. --------------- ENTER COMMAND?WEST YOU ARE LOST IN A FOREST.">

**Type `SOUTH`**

<img src="images/ulysses/0010.png" width="400" alt="YOU ARE LOST IN A FOREST. --------------- ENTER COMMAND?SOUTH YOU ARE IN A LUSH GREEN FOREST.">

**Type `WEST`**

<img src="images/ulysses/0011.png" width="400" alt="YOU ARE IN A LUSH GREEN FOREST. --------------- ENTER COMMAND?WEST YOU ARE IN A FOREST.">

**Type `NORTH`**

<img src="images/ulysses/0012.png" width="400" alt="YOU ARE IN A FOREST. --------------- ENTER COMMAND?NORTH YOU ARE IN A FOREST.">

**Type `NORTH`**

<img src="images/ulysses/0013.png" width="400" alt="YOU ARE IN A FOREST. --------------- ENTER COMMAND?NORTH YOU ARE IN A LUSH GREEN FOREST.">

**Type `WEST`**

<img src="images/ulysses/0014.png" width="400" alt="YOU ARE ON A 3-WAY ROAD IN A SMALL TOWN IN ANCIENT GREECE. THERE IS A STORE TO THE WEST WITH A FENCE NEXT TO IT.">

**Type `SOUTH`**

<img src="images/ulysses/0015-1.png" width="400" alt="YOU ARE ON THE OUT-SKIRTS OF A SMALL TOWN. A ROAD LEADS NORTH AND WEST. YOU SEE A CASTLE IN THE DISTANCE TO THE WEST."><br>
<img src="images/ulysses/0015.png" width="400" alt="TOWN. A ROAD LEADS NORTH AND WEST. YOU SEE A CASTLE IN THE DISTANCE TO THE WEST.">

**Type `WEST`**

<img src="images/ulysses/0016.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN THE KING&#39;S VINYARD. THE ROAD LEADS EAST AND WEST.">

**Type `WEST`**

<img src="images/ulysses/0017.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE OUTSIDE THE KING&#39;S CASTLE FACING WEST.">

**Type `WEST`**

<img src="images/ulysses/0018.png" width="400" alt="YOU ARE IN THE ENTRY-HALL OF THE KING&#39;S CASTLE. THERE IS A GUARD HERE. AN EXIT IS TO THE EAST.">

**Type `TALK GUARD`**

<img src="images/ulysses/0019-1.png" width="400" alt="THE GUARD ASKS,&#34;ARE YOU ULYSSES?&#34; YOU ARE IN THE ENTRY-HALL OF THE KING&#39;S CASTLE. THERE IS A GUARD HERE. AN EXIT IS TO THE EAST."><br>
<img src="images/ulysses/0019.png" width="400" alt="YOU ARE IN THE ENTRY-HALL OF THE KING&#39;S CASTLE. THERE IS A GUARD HERE. AN EXIT IS TO THE EAST.">

**Type `YES`**

<img src="images/ulysses/0020-1.png" width="400" alt="&#34;GOOD. THE KING IS EXPECTING YOU. I WILL ESCORT YOU.&#34; I THINK YOU HAD BETTER BOW. YOU ARE IN THE KING&#39;S THRONE ROOM"><br>
<img src="images/ulysses/0020.png" width="400" alt="I THINK YOU HAD BETTER BOW. YOU ARE IN THE KING&#39;S THRONE ROOM STANDING BEFORE HIS MAJESTY.">

**Type `BOW`**

### 2. The store

<img src="images/ulysses/0021-1.png" width="400" alt="THE KING ACCEPTS THIS HUMBLE GESTURE AND SAYS,&#34;THERE IS A LEGENDARY FLEECE OF GOLD, FAR OFF TO THE NORTH. I GIVE YOU SOME GOLD AND SILVER, AND A SHIP TO"><br>
<img src="images/ulysses/0021-2.png" width="400" alt="COMPLETE YOUR VOYAGE AND RETURN HOME WITH THE FLEECE. YOU TAKE THE GOLD AND SILVER, THEN THE GUARD ESCORTS YOU BACK OUT OF THE"><br>
<img src="images/ulysses/0021.png" width="400" alt="PALACE. YOU ARE OUTSIDE THE KING&#39;S CASTLE FACING WEST.">

**Type `EAST`**

<img src="images/ulysses/0022.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN THE KING&#39;S VINYARD. THE ROAD LEADS EAST AND WEST.">

**Type `EAST`**

<img src="images/ulysses/0023-1.png" width="400" alt="YOU ARE ON THE OUT-SKIRTS OF A SMALL TOWN. A ROAD LEADS NORTH AND WEST. YOU SEE A CASTLE IN THE DISTANCE TO THE WEST."><br>
<img src="images/ulysses/0023.png" width="400" alt="TOWN. A ROAD LEADS NORTH AND WEST. YOU SEE A CASTLE IN THE DISTANCE TO THE WEST.">

**Type `NORTH`**

<img src="images/ulysses/0024.png" width="400" alt="YOU ARE ON A 3-WAY ROAD IN A SMALL TOWN IN ANCIENT GREECE. THERE IS A STORE TO THE WEST WITH A FENCE NEXT TO IT.">

**Type `WEST`**

<img src="images/ulysses/0025.png" width="400" alt="YOU ARE IN A STORE. THERE ARE DOORS TO THE EAST AND NORTH. THERE IS A SIGN HERE.">

**Type `BUY WINE`**

<img src="images/ulysses/0026-1.png" width="400" alt="THE STOREKEEPER ASKS,&#34;WILL THAT BE ALL?&#34; YOU ARE IN A STORE. THERE ARE DOORS TO THE EAST AND NORTH. THERE IS A SIGN"><br>
<img src="images/ulysses/0026.png" width="400" alt="YOU ARE IN A STORE. THERE ARE DOORS TO THE EAST AND NORTH. THERE IS A SIGN HERE.">

**Type `BUY ROPE`**

<img src="images/ulysses/0027-1.png" width="400" alt="THE STOREKEEPER ASKS,&#34;WILL THAT BE ALL?&#34; YOU ARE IN A STORE. THERE ARE DOORS TO THE EAST AND NORTH. THERE IS A SIGN"><br>
<img src="images/ulysses/0027.png" width="400" alt="YOU ARE IN A STORE. THERE ARE DOORS TO THE EAST AND NORTH. THERE IS A SIGN HERE.">

**Type `BUY WAX`**

<img src="images/ulysses/0028-1.png" width="400" alt="THE STOREKEEPER ASKS,&#34;WILL THAT BE ALL?&#34; YOU ARE IN A STORE. THERE ARE DOORS TO THE EAST AND NORTH. THERE IS A SIGN"><br>
<img src="images/ulysses/0028.png" width="400" alt="YOU ARE IN A STORE. THERE ARE DOORS TO THE EAST AND NORTH. THERE IS A SIGN HERE.">

**Type `BUY WOOD`**

<img src="images/ulysses/0029-1.png" width="400" alt="THE STOREKEEPER ASKS,&#34;WILL THAT BE ALL?&#34; YOU ARE IN A STORE. THERE ARE DOORS TO THE EAST AND NORTH. THERE IS A SIGN"><br>
<img src="images/ulysses/0029.png" width="400" alt="YOU ARE IN A STORE. THERE ARE DOORS TO THE EAST AND NORTH. THERE IS A SIGN HERE.">

**Type `BUY SWORD`**

<img src="images/ulysses/0030-1.png" width="400" alt="THE STOREKEEPER ASKS,&#34;WILL THAT BE ALL?&#34; YOU ARE IN A STORE. THERE ARE DOORS TO THE EAST AND NORTH. THERE IS A SIGN"><br>
<img src="images/ulysses/0030.png" width="400" alt="YOU ARE IN A STORE. THERE ARE DOORS TO THE EAST AND NORTH. THERE IS A SIGN HERE.">

**Type `BUY LEATHER`**

<img src="images/ulysses/0031-1.png" width="400" alt="THE STOREKEEPER ASKS,&#34;WILL THAT BE ALL?&#34; YOU ARE IN A STORE. THERE ARE DOORS TO THE EAST AND NORTH. THERE IS A SIGN"><br>
<img src="images/ulysses/0031.png" width="400" alt="YOU ARE IN A STORE. THERE ARE DOORS TO THE EAST AND NORTH. THERE IS A SIGN HERE.">

**Type `BUY FLINT`**

<img src="images/ulysses/0032-1.png" width="400" alt="THE STOREKEEPER ASKS,&#34;WILL THAT BE ALL?&#34; YOU ARE IN A STORE. THERE ARE DOORS TO THE EAST AND NORTH. THERE IS A SIGN"><br>
<img src="images/ulysses/0032.png" width="400" alt="YOU ARE IN A STORE. THERE ARE DOORS TO THE EAST AND NORTH. THERE IS A SIGN HERE.">

**Type `YES`**

### 3. The tavern and the docks

<img src="images/ulysses/0033-1.png" width="400" alt="THE STOREKEEPER THANKS YOU AND TAKES A BAG OF SILVER. YOU ARE IN A STORE. THERE ARE DOORS TO THE EAST AND NORTH. THERE IS A SIGN"><br>
<img src="images/ulysses/0033.png" width="400" alt="YOU ARE IN A STORE. THERE ARE DOORS TO THE EAST AND NORTH. THERE IS A SIGN HERE.">

**Type `EAST`**

<img src="images/ulysses/0034.png" width="400" alt="YOU ARE ON A 3-WAY ROAD IN A SMALL TOWN IN ANCIENT GREECE. THERE IS A STORE TO THE WEST WITH A FENCE NEXT TO IT.">

**Type `NORTH`**

<img src="images/ulysses/0035-1.png" width="400" alt="THERE SEEMS TO BE A BOTTLE FLOATING HERE. YOU ARE ON SOME DOCKS. YOU SEE A TAVERN TO THE WEST AND AN OCEAN TO THE NORTH."><br>
<img src="images/ulysses/0035.png" width="400" alt="TO THE WEST AND AN OCEAN TO THE NORTH. A ROAD LEADS TO THE SOUTH. THERE IS A GUARD HERE.">

**Type `WEST`**

<img src="images/ulysses/0036.png" width="400" alt="YOU ARE IN A SMALL TAVERN. SOMETIMES SAILORS WAIT HERE HOPING FOR WORK. THERE ARE DOORS TO THE SOUTH AND EAST.">

**Type `DROP CHEST`**

<img src="images/ulysses/0037.png" width="400" alt="YOU ARE IN A SMALL TAVERN. SOMETIMES SAILORS WAIT HERE HOPING FOR WORK. THERE ARE DOORS TO THE SOUTH AND EAST.">

**Type `SOUTH`**

<img src="images/ulysses/0038-1.png" width="400" alt="THERE IS A GOLD COIN HERE. YOU ARE IN AN ALLEY IN A SMALL VILLAGE.THERE IS A TAVERN TO THE NORTH AND A STORE TO THE SOUTH."><br>
<img src="images/ulysses/0038.png" width="400" alt="YOU ARE IN AN ALLEY IN A SMALL VILLAGE.THERE IS A TAVERN TO THE NORTH AND A STORE TO THE SOUTH.">

**Type `GET COIN`**

<img src="images/ulysses/0039.png" width="400" alt="YOU ARE IN AN ALLEY IN A SMALL VILLAGE.THERE IS A TAVERN TO THE NORTH AND A STORE TO THE SOUTH.">

**Type `NORTH`**

<img src="images/ulysses/0040.png" width="400" alt="YOU ARE IN A SMALL TAVERN. SOMETIMES SAILORS WAIT HERE HOPING FOR WORK. THERE ARE DOORS TO THE SOUTH AND EAST.">

**Type `GET CHEST`**

<img src="images/ulysses/0041.png" width="400" alt="YOU ARE IN A SMALL TAVERN. SOMETIMES SAILORS WAIT HERE HOPING FOR WORK. THERE ARE DOORS TO THE SOUTH AND EAST.">

**Type `HIRE CREW`**

<img src="images/ulysses/0042-1.png" width="400" alt="SEVERAL MEN RUSH FORWARD,(I THINK ONE OF THEM IS HERCULES) AND TAKE YOUR BAG OF GOLD. YOU NOW HAVE A FULL CREW. YOUR CREW IS FOLLOWING YOU."><br>
<img src="images/ulysses/0042.png" width="400" alt="YOU ARE IN A SMALL TAVERN. SOMETIMES SAILORS WAIT HERE HOPING FOR WORK. THERE ARE DOORS TO THE SOUTH AND EAST.">

**Type `EAST`**

<img src="images/ulysses/0043-1.png" width="400" alt="THERE SEEMS TO BE A BOTTLE FLOATING HERE. YOUR CREW IS FOLLOWING YOU. YOU ARE ON SOME DOCKS. YOU SEE A TAVERN"><br>
<img src="images/ulysses/0043.png" width="400" alt="TO THE WEST AND AN OCEAN TO THE NORTH. A ROAD LEADS TO THE SOUTH. THERE IS A GUARD HERE.">

**Type `GIVE COIN`**

<img src="images/ulysses/0044-1.png" width="400" alt="ARE YOU ATTEMPTING TO BRIBE A GUARD? THERE SEEMS TO BE A BOTTLE FLOATING HERE. YOUR CREW IS FOLLOWING YOU."><br>
<img src="images/ulysses/0044-2.png" width="400" alt="YOU ARE ON SOME DOCKS. YOU SEE A TAVERN TO THE WEST AND AN OCEAN TO THE NORTH. A ROAD LEADS TO THE SOUTH. THERE IS A GUARD HERE."><br>
<img src="images/ulysses/0044.png" width="400" alt="TO THE WEST AND AN OCEAN TO THE NORTH. A ROAD LEADS TO THE SOUTH. THERE IS A GUARD HERE.">

**Type `YES`**

<img src="images/ulysses/0045-1.png" width="400" alt="THE GUARD ACCEPTS THE BRIBE AND GIVES YOU A MAP. THERE SEEMS TO BE A BOTTLE FLOATING HERE."><br>
<img src="images/ulysses/0045-2.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE ON SOME DOCKS. YOU SEE A TAVERN TO THE WEST AND AN OCEAN TO THE NORTH. A ROAD LEADS TO THE SOUTH. THERE IS A"><br>
<img src="images/ulysses/0045.png" width="400" alt="TO THE WEST AND AN OCEAN TO THE NORTH. A ROAD LEADS TO THE SOUTH. THERE IS A GUARD HERE.">

**Type `GET BOTTLE`**

<img src="images/ulysses/0046-1.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE ON SOME DOCKS. YOU SEE A TAVERN TO THE WEST AND AN OCEAN TO THE NORTH. A ROAD LEADS TO THE SOUTH. THERE IS A"><br>
<img src="images/ulysses/0046.png" width="400" alt="TO THE WEST AND AN OCEAN TO THE NORTH. A ROAD LEADS TO THE SOUTH. THERE IS A GUARD HERE.">

**Type `GET NOTE`**

<img src="images/ulysses/0047-1.png" width="400" alt="O.K. YOUR CREW IS FOLLOWING YOU. YOU ARE ON SOME DOCKS. YOU SEE A TAVERN TO THE WEST AND AN OCEAN TO THE NORTH."><br>
<img src="images/ulysses/0047.png" width="400" alt="TO THE WEST AND AN OCEAN TO THE NORTH. A ROAD LEADS TO THE SOUTH. THERE IS A GUARD HERE.">

**Type `READ NOTE`**

### 4. At sea

<img src="images/ulysses/0048-1.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE ON SOME DOCKS. YOU SEE A TAVERN TO THE WEST AND AN OCEAN TO THE NORTH. A ROAD LEADS TO THE SOUTH. THERE IS A"><br>
<img src="images/ulysses/0048.png" width="400" alt="TO THE WEST AND AN OCEAN TO THE NORTH. A ROAD LEADS TO THE SOUTH. THERE IS A GUARD HERE.">

**Type `EAST`**

<img src="images/ulysses/0049.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE ON THE KING&#39;S PRIVATE DOCK. THERE IS A GREAT SHIP HERE.">

**Type `NORTH`**

<img src="images/ulysses/0050.png" width="400" alt="THERE IS A GREAT SHIP HERE. --------------- ENTER COMMAND?NORTH YOU ARE ON THE DECK OF A SHIP.">

**Type `CAST OFF`**

<img src="images/ulysses/0051.png" width="400" alt="YOU ARE ON THE DECK OF A SHIP. --------------- ENTER COMMAND?CAST OFF YOU ARE IN THE OCEAN.">

**Type `NORTH`**

<img src="images/ulysses/0052.png" width="400" alt="YOU ARE IN THE OCEAN. --------------- ENTER COMMAND?NORTH YOU ARE IN THE OCEAN.">

**Type `WEST`**

<img src="images/ulysses/0053.png" width="400" alt="AN ALBATROSS SWOOPS DOWN CLOSE TO YOUR SHIP AND DROPS A BAG. YOU ARE LOST IN THE OCEAN.">

**Type `GET BAG`**

<img src="images/ulysses/0054.png" width="400" alt="YOU ARE LOST IN THE OCEAN. --------------- ENTER COMMAND?GET BAG YOU ARE LOST IN THE OCEAN.">

**Type `EAST`**

<img src="images/ulysses/0055.png" width="400" alt="YOU ARE LOST IN THE OCEAN. --------------- ENTER COMMAND?EAST YOU ARE IN THE OCEAN.">

**Type `SOUTH`**

<img src="images/ulysses/0056.png" width="400" alt="YOU ARE IN THE OCEAN. --------------- ENTER COMMAND?SOUTH YOU ARE IN THE OCEAN.">

**Type `EAST`**

<img src="images/ulysses/0057-1.png" width="400" alt="A GIANT CONDOR HAS FLOWN OFF COURSE AND SLAMMED INTO YOUR MAST, KILLING ITSELF AND FALLING TO YOUR FEET. YOU ARE IN THE OCEAN."><br>
<img src="images/ulysses/0057.png" width="400" alt="SLAMMED INTO YOUR MAST, KILLING ITSELF AND FALLING TO YOUR FEET. YOU ARE IN THE OCEAN.">

**Type `GET CONDOR`**

<img src="images/ulysses/0058.png" width="400" alt="--------------- ENTER COMMAND?GET CONDOR YOU ARE IN THE OCEAN.">

**Type `EAST`**

<img src="images/ulysses/0059.png" width="400" alt="YOU ARE APPROACHING A HURRICANE, FACING EAST. THROUGH THE STORM YOU CAN BARELY MAKE OUT AN ISLAND.">

**Type `NORTH`**

<img src="images/ulysses/0060.png" width="400" alt="MAKE OUT AN ISLAND. --------------- ENTER COMMAND?NORTH YOU ARE IN THE OCEAN.">

**Type `NORTH`**

<img src="images/ulysses/0061.png" width="400" alt="YOU ARE IN THE OCEAN. --------------- ENTER COMMAND?NORTH YOU ARE IN THE OCEAN.">

**Type `EAST`**

<img src="images/ulysses/0062.png" width="400" alt="YOU ARE IN THE OCEAN. --------------- ENTER COMMAND?EAST YOU ARE IN THE OCEAN.">

**Type `SOUTH`**

<img src="images/ulysses/0063.png" width="400" alt="YOU ARE IN THE OCEAN. --------------- ENTER COMMAND?SOUTH YOU ARE IN THE OCEAN.">

**Type `WEST`**

<img src="images/ulysses/0064.png" width="400" alt="YOU ARE IN THE OCEAN. --------------- ENTER COMMAND?WEST YOU ARE IN THE OCEAN.">

**Type `SOUTH`**

<img src="images/ulysses/0065.png" width="400" alt="YOU ARE IN THE OCEAN. --------------- ENTER COMMAND?SOUTH YOU ARE IN THE OCEAN.">

**Type `WEST`**

<img src="images/ulysses/0066.png" width="400" alt="YOU ARE IN THE OCEAN. --------------- ENTER COMMAND?WEST YOU ARE IN THE OCEAN.">

**Type `SOUTH`**

<img src="images/ulysses/0067.png" width="400" alt="YOU ARE IN THE OCEAN. --------------- ENTER COMMAND?SOUTH YOU ARE IN THE OCEAN.">

**Type `EAST`**

<img src="images/ulysses/0068.png" width="400" alt="YOU ARE IN THE OCEAN. --------------- ENTER COMMAND?EAST YOU ARE IN THE OCEAN.">

**Type `NORTH`**

<img src="images/ulysses/0069.png" width="400" alt="YOU ARE IN THE OCEAN. --------------- ENTER COMMAND?NORTH YOU ARE IN THE OCEAN.">

**Type `EAST`**

<img src="images/ulysses/0070.png" width="400" alt="YOU ARE IN THE OCEAN. --------------- ENTER COMMAND?EAST YOU ARE IN THE OCEAN.">

**Type `NORTH`**

<img src="images/ulysses/0071.png" width="400" alt="YOU ARE IN THE OCEAN. --------------- ENTER COMMAND?NORTH YOU ARE IN THE OCEAN.">

**Type `EAST`**

<img src="images/ulysses/0072.png" width="400" alt="YOU ARE IN THE OCEAN. --------------- ENTER COMMAND?EAST YOU ARE IN THE OCEAN.">

**Type `EAST`**

### 5. The Island of Storms

<img src="images/ulysses/0073.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE JUST OFF-SHORE OF THE ISLAND OF STORMS.">

**Type `GO ISLAND`**

<img src="images/ulysses/0074-1.png" width="400" alt="YOUR CREW STAYS BEHIND. YOU ARE ON THE BEACH OF THE ISLAND OF STORMS. YOU SEE A CAVE IN THE DISTANCE TO THE NORTH. YOUR SHIP IS ANCHORED OFF"><br>
<img src="images/ulysses/0074.png" width="400" alt="STORMS. YOU SEE A CAVE IN THE DISTANCE TO THE NORTH. YOUR SHIP IS ANCHORED OFF SHORE.">

**Type `EAST`**

<img src="images/ulysses/0075.png" width="400" alt="SHORE. --------------- ENTER COMMAND?EAST YOU ARE IN A JUNGLE.">

**Type `SOUTH`**

<img src="images/ulysses/0076.png" width="400" alt="YOU ARE IN A JUNGLE. --------------- ENTER COMMAND?SOUTH YOU ARE WANDERING IN A JUNGLE.">

**Type `SOUTH`**

<img src="images/ulysses/0077.png" width="400" alt="--------------- ENTER COMMAND?SOUTH THERE IS A BRIDLE HERE. YOU ARE LOST IN A JUNGLE.">

**Type `GET BRIDLE`**

<img src="images/ulysses/0078.png" width="400" alt="--------------- ENTER COMMAND?GET BRIDLE YOU ARE LOST IN A JUNGLE.">

**Type `EAST`**

<img src="images/ulysses/0079.png" width="400" alt="YOU ARE LOST IN A JUNGLE. --------------- ENTER COMMAND?EAST YOU ARE IN A JUNGLE.">

**Type `SOUTH`**

<img src="images/ulysses/0080.png" width="400" alt="YOU ARE IN A JUNGLE. --------------- ENTER COMMAND?SOUTH YOU ARE IN A JUNGLE.">

**Type `LOOK HOLE`**

<img src="images/ulysses/0081.png" width="400" alt="THERE IS A SMALL PILE OF SPARKLING DUST HERE. YOU ARE IN A JUNGLE.">

**Type `GET DUST`**

<img src="images/ulysses/0082.png" width="400" alt="YOU ARE IN A JUNGLE. --------------- ENTER COMMAND?GET DUST YOU ARE IN A JUNGLE.">

**Type `NORTH`**

<img src="images/ulysses/0083.png" width="400" alt="YOU ARE IN A JUNGLE. --------------- ENTER COMMAND?NORTH YOU ARE IN A JUNGLE.">

**Type `WEST`**

<img src="images/ulysses/0084.png" width="400" alt="YOU ARE IN A JUNGLE. --------------- ENTER COMMAND?WEST YOU ARE LOST IN A JUNGLE.">

**Type `NORTH`**

<img src="images/ulysses/0085.png" width="400" alt="YOU ARE LOST IN A JUNGLE. --------------- ENTER COMMAND?NORTH YOU ARE WANDERING IN A JUNGLE.">

**Type `NORTH`**

<img src="images/ulysses/0086.png" width="400" alt="YOU ARE WANDERING IN A JUNGLE. --------------- ENTER COMMAND?NORTH YOU ARE IN A JUNGLE.">

**Type `NORTH`**

<img src="images/ulysses/0087.png" width="400" alt="YOU ARE IN A JUNGLE. --------------- ENTER COMMAND?NORTH YOU ARE LOST IN A JUNGLE.">

**Type `NORTH`**

<img src="images/ulysses/0088.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE LOST IN A JUNGLE. YOU SEE A CAVE ON A MOUNTAIN TO THE NORTH.">

**Type `NORTH`**

<img src="images/ulysses/0089.png" width="400" alt="YOU ARE AT THE BASE OF A LARGE MOUNTAIN. YOU SEE A CAVE HIGH UP ON TOP.">

**Type `UP`**

<img src="images/ulysses/0090.png" width="400" alt="--------------- ENTER COMMAND?UP YOU ARE ATOP A LARGE MOUNTAIN. THERE IS A CAVE TO THE NORTH.">

**Type `NORTH`**

<img src="images/ulysses/0091-1.png" width="400" alt="YOU HEAR A FAINT BUBBLING NOISE FROM SOMEWHERE DEEP IN THE TUNNELS. YOU ARE INSIDE A CAVE. THERE ARE DIMLY LIT TUNNELS LEADING NORTH, EAST AND"><br>
<img src="images/ulysses/0091.png" width="400" alt="YOU ARE INSIDE A CAVE. THERE ARE DIMLY LIT TUNNELS LEADING NORTH, EAST AND WEST. THERE IS AN EXIT TO THE SOUTH.">

**Type `EAST`**

<img src="images/ulysses/0092.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN A TUNNEL. PATHS RUN NORTH AND WEST.">

**Type `NORTH`**

<img src="images/ulysses/0093.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN AN OPEN ROOM. THERE IS A SPRING HERE.">

**Type `GET WATER`**

### 6. Under the island: the fjord, the dragon and the wall of fire

<img src="images/ulysses/0094.png" width="400" alt="O.K. YOU ARE IN AN OPEN ROOM. THERE IS A SPRING HERE.">

**Type `SOUTH`**

<img src="images/ulysses/0095.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN A TUNNEL. PATHS RUN NORTH SOUTH EAST AND WEST.">

**Type `NORTH`**

<img src="images/ulysses/0096.png" width="400" alt="YOU ARE IN A TUNNEL. THERE APPEARS TO BE A PIT IN THE FLOOR. PATHS LEAD SOUTH AND WEST.">

**Type `DOWN`**

<img src="images/ulysses/0097-1.png" width="400" alt="YOU ARE IN AN UNDERGROUND PASSAGEWAY. THERE IS A HOLE IN THE CEILING AND A PATH LEADING SOUTH. PHOSPHORUS LINES THE WALLS."><br>
<img src="images/ulysses/0097.png" width="400" alt="THERE IS A HOLE IN THE CEILING AND A PATH LEADING SOUTH. PHOSPHORUS LINES THE WALLS.">

**Type `SOUTH`**

<img src="images/ulysses/0098.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN AN UNDERGROUND PASSAGEWAY. PATHS LEAD NORTH, SOUTH AND EAST.">

**Type `SOUTH`**

<img src="images/ulysses/0099.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN AN UNDERGROUND PASSAGEWAY. PATHS LEAD NORTH AND EAST.">

**Type `EAST`**

<img src="images/ulysses/0100.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN AN UNDERGROUND PASSAGEWAY. PATHS LEAD NORTH, WEST AND EAST.">

**Type `EAST`**

<img src="images/ulysses/0101.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN AN UNDERGROUND PASSAGEWAY. PATHS LEAD NORTH, SOUTH AND WEST.">

**Type `SOUTH`**

<img src="images/ulysses/0102.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN AN UNDERGROUND PASSAGEWAY. PATHS LEAD NORTH, EAST AND WEST.">

**Type `EAST`**

<img src="images/ulysses/0103-1.png" width="400" alt="YOU ARE IN AN UNDERGROUND PASSAGEWAY. A PATH LEADS WEST. THERE IS A DEEP FJORD TO THE EAST. THERE ARE TWO ROCKS ON THE OTHER SIDE, CLOSE TOGETHER."><br>
<img src="images/ulysses/0103.png" width="400" alt="PATH LEADS WEST. THERE IS A DEEP FJORD TO THE EAST. THERE ARE TWO ROCKS ON THE OTHER SIDE, CLOSE TOGETHER.">

**Type `TIE LEATHER`**

<img src="images/ulysses/0104-1.png" width="400" alt="TO WHAT? YOU ARE IN AN UNDERGROUND PASSAGEWAY. A PATH LEADS WEST. THERE IS A DEEP FJORD TO THE EAST. THERE ARE TWO ROCKS ON THE"><br>
<img src="images/ulysses/0104.png" width="400" alt="PATH LEADS WEST. THERE IS A DEEP FJORD TO THE EAST. THERE ARE TWO ROCKS ON THE OTHER SIDE, CLOSE TOGETHER.">

**Type `TO LEATHER`**

<img src="images/ulysses/0105-1.png" width="400" alt="O.K. YOU BIND THE LEATHER STRAPS SECURELY TOGETHER. YOU ARE IN AN UNDERGROUND PASSAGEWAY. A PATH LEADS WEST. THERE IS A DEEP FJORD"><br>
<img src="images/ulysses/0105.png" width="400" alt="PATH LEADS WEST. THERE IS A DEEP FJORD TO THE EAST. THERE ARE TWO ROCKS ON THE OTHER SIDE, CLOSE TOGETHER.">

**Type `THROW LEATHER`**

<img src="images/ulysses/0106-1.png" width="400" alt="YOU THROW THE STRAPS AND THEY SPAN THE FJORD, LODGING SECURELY BETWEEN TWO ROCKS ON THE OTHER SIDE. YOU ARE IN AN UNDERGROUND PASSAGEWAY. A"><br>
<img src="images/ulysses/0106.png" width="400" alt="PATH LEADS WEST. THERE IS A DEEP FJORD TO THE EAST. THERE ARE TWO ROCKS ON THE OTHER SIDE, CLOSE TOGETHER.">

**Type `EAST`**

<img src="images/ulysses/0107-1.png" width="400" alt="YOU SAFELY HOIST YOURSELF ACROSS THE FJORD. AS SOON AS YOU CROSS, THE STRAPS COME LOOSE AND FALL INTO THE FJORD, LOST"><br>
<img src="images/ulysses/0107-2.png" width="400" alt="FOREVER. YOU ARE IN AN UNDERGROUND PASSAGEWAY. A PATH LEADS NORTH. THERE IS A DEEP FJORD TO THE WEST."><br>
<img src="images/ulysses/0107.png" width="400" alt="YOU ARE IN AN UNDERGROUND PASSAGEWAY. A PATH LEADS NORTH. THERE IS A DEEP FJORD TO THE WEST.">

**Type `NORTH`**

<img src="images/ulysses/0108.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN AN UNDERGROUND PASSAGEWAY. PATHS LEAD SOUTH, EAST AND WEST.">

**Type `EAST`**

<img src="images/ulysses/0109.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN AN UNDERGROUND PASSAGEWAY. PATHS LEAD NORTH, SOUTH AND WEST.">

**Type `NORTH`**

<img src="images/ulysses/0110-1.png" width="400" alt="THERE IS A DRAGON HERE. I THINK HE MEANS BUSINESS. YOU ARE IN AN UNDERGROUND PASSAGEWAY. PATHS LEAD NORTH, EAST AND SOUTH."><br>
<img src="images/ulysses/0110.png" width="400" alt="MEANS BUSINESS. YOU ARE IN AN UNDERGROUND PASSAGEWAY. PATHS LEAD NORTH, EAST AND SOUTH.">

**Type `GIVE JEWELS`**

<img src="images/ulysses/0111-1.png" width="400" alt="THE DRAGON TAKES THE GEMS AND SMILES,(DRAGONS ARE VERY GREEDY) HE LIKES THE WAY THE STONES GLISTEN. HE HAPPILY WALKS AWAY."><br>
<img src="images/ulysses/0111.png" width="400" alt="HAPPILY WALKS AWAY. YOU ARE IN AN UNDERGROUND PASSAGEWAY. PATHS LEAD NORTH, EAST AND SOUTH.">

**Type `NORTH`**

<img src="images/ulysses/0112.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN AN UNDERGROUND PASSAGEWAY. PATHS LEAD SOUTH AND EAST.">

**Type `EAST`**

<img src="images/ulysses/0113.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN AN UNDERGROUND PASSAGEWAY. PATHS LEAD SOUTH, EAST AND WEST.">

**Type `EAST`**

<img src="images/ulysses/0114.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN AN UNDERGROUND PASSAGEWAY. PATHS LEAD NORTH, SOUTH AND WEST.">

**Type `NORTH`**

<img src="images/ulysses/0115.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE STANDING BEFORE A GREAT CANYON.THERE IS A PATH LEADING SOUTH.">

**Type `PLUCK CONDOR`**

<img src="images/ulysses/0116.png" width="400" alt="O.K. YOU ARE STANDING BEFORE A GREAT CANYON.THERE IS A PATH LEADING SOUTH.">

**Type `USE WAX`**

<img src="images/ulysses/0117.png" width="400" alt="AND WHAT? YOU ARE STANDING BEFORE A GREAT CANYON.THERE IS A PATH LEADING SOUTH.">

**Type `AND FEATHERS`**

<img src="images/ulysses/0118-1.png" width="400" alt="USING SOME WAX, AND FEATHERS FROM THE GIANT CONDOR, YOU FASHION SOME BEAUTIFUL WINGS. YOU ARE STANDING BEFORE A GREAT"><br>
<img src="images/ulysses/0118.png" width="400" alt="BEAUTIFUL WINGS. YOU ARE STANDING BEFORE A GREAT CANYON.THERE IS A PATH LEADING SOUTH.">

**Type `FLY`**

<img src="images/ulysses/0119-1.png" width="400" alt="YOU LEAP OVER THE EDGE OF THE CANYON AND SOAR GRACEFULLY TO THE FAR EDGE. YOU ARE IN AN UNDERGROUND PASSAGEWAY. THERE IS A LARGE FJORD TO THE EAST."><br>
<img src="images/ulysses/0119.png" width="400" alt="THERE IS A LARGE FJORD TO THE EAST. THERE IS A HOLE HERE WITH A ROCK NEXT TO IT.">

**Type `GET ROCK`**

<img src="images/ulysses/0120-1.png" width="400" alt="O.K. THERE ARE SOME REINS HERE. YOU ARE IN AN UNDERGROUND PASSAGEWAY. THERE IS A LARGE FJORD TO THE EAST."><br>
<img src="images/ulysses/0120.png" width="400" alt="THERE IS A LARGE FJORD TO THE EAST. THERE IS A HOLE HERE WITH A ROCK NEXT TO IT.">

**Type `GET REINS`**

<img src="images/ulysses/0121-1.png" width="400" alt="YOU ARE IN AN UNDERGROUND PASSAGEWAY. THERE IS A LARGE FJORD TO THE EAST. THERE IS A HOLE HERE WITH A ROCK NEXT TO IT."><br>
<img src="images/ulysses/0121.png" width="400" alt="THERE IS A LARGE FJORD TO THE EAST. THERE IS A HOLE HERE WITH A ROCK NEXT TO IT.">

**Type `DOWN`**

<img src="images/ulysses/0122.png" width="400" alt="YOU ARE IN AN UNDERGROUND PASSAGEWAY. A PATH LEADS WEST. THERE IS A HOLE IN THE CEILING.">

**Type `WEST`**

<img src="images/ulysses/0123.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN AN UNDERGROUND PASSAGEWAY. PATHS LEAD NORTH, EAST AND WEST.">

**Type `WEST`**

<img src="images/ulysses/0124.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN AN UNDERGROUND PASSAGEWAY. PATHS LEAD SOUTH AND EAST.">

**Type `SOUTH`**

<img src="images/ulysses/0125-1.png" width="400" alt="THIS IS PLUTO. THE GOD OF THE UNDERWORLD. YOU ARE IN AN UNDERGROUND PASSAGEWAY. PATHS LEAD NORTH AND SOUTH."><br>
<img src="images/ulysses/0125.png" width="400" alt="UNDERWORLD. YOU ARE IN AN UNDERGROUND PASSAGEWAY. PATHS LEAD NORTH AND SOUTH.">

**Type `THROW DUST`**

<img src="images/ulysses/0126-1.png" width="400" alt="YOU THROW THE DUST RIGHT INTO HIS EYES. HE SCREAMS IN TERROR AND DISAPPEARS THROUGH THE TUNNELS HOWLING. YOU ARE IN AN UNDERGROUND PASSAGEWAY."><br>
<img src="images/ulysses/0126.png" width="400" alt="THROUGH THE TUNNELS HOWLING. YOU ARE IN AN UNDERGROUND PASSAGEWAY. PATHS LEAD NORTH AND SOUTH.">

**Type `SOUTH`**

<img src="images/ulysses/0127.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN AN UNDERGROUND PASSAGEWAY. PATHS LEAD NORTH AND EAST.">

**Type `EAST`**

<img src="images/ulysses/0128.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN AN UNDERGROUND PASSAGEWAY. PATHS LEAD WEST AND EAST.">

**Type `EAST`**

<img src="images/ulysses/0129.png" width="400" alt="YOU ARE IN AN UNDERGROUND PASSAGEWAY FACING SOUTH. A PATH LEADS WEST AND ONE LEADS SOUTH INTO A GREAT WALL OF FIRE.">

**Type `POUR WINE`**

<img src="images/ulysses/0130-1.png" width="400" alt="WHERE? YOU ARE IN AN UNDERGROUND PASSAGEWAY FACING SOUTH. A PATH LEADS WEST AND ONE LEADS SOUTH INTO A GREAT WALL OF FIRE."><br>
<img src="images/ulysses/0130.png" width="400" alt="YOU ARE IN AN UNDERGROUND PASSAGEWAY FACING SOUTH. A PATH LEADS WEST AND ONE LEADS SOUTH INTO A GREAT WALL OF FIRE.">

**Type `ON ME`**

<img src="images/ulysses/0131-1.png" width="400" alt="O.K. YOU ARE IN AN UNDERGROUND PASSAGEWAY FACING SOUTH. A PATH LEADS WEST AND ONE LEADS SOUTH INTO A GREAT WALL OF FIRE."><br>
<img src="images/ulysses/0131.png" width="400" alt="YOU ARE IN AN UNDERGROUND PASSAGEWAY FACING SOUTH. A PATH LEADS WEST AND ONE LEADS SOUTH INTO A GREAT WALL OF FIRE.">

**Type `SOUTH`**

<img src="images/ulysses/0132-1.png" width="400" alt="DRENCHED WITH WINE, YOU SAFELY PASS THROUGH THE FIRE. THE WINE HAS DRIED. YOU ARE IN AN UNDERGROUND PASSAGEWAY."><br>
<img src="images/ulysses/0132.png" width="400" alt="PATHS LEAD SOUTH AND EAST. THERE IS A GIANT WALL OF FIRE DIRECTLY IN FRONT OF YOU.">

**Type `EAST`**

<img src="images/ulysses/0133.png" width="400" alt="YOU ARE IN AN UNDERGROUND PASSAGEWAY. A PATH LEADS WEST. THERE IS A LARGE HOLE IN THE WALL HERE.">

**Type `GO HOLE`**

<img src="images/ulysses/0134.png" width="400" alt="IN THE WALL HERE. --------------- ENTER COMMAND?GO HOLE YOU ARE LOST IN A JUNGLE.">

**Type `SOUTH`**

<img src="images/ulysses/0135.png" width="400" alt="YOU ARE LOST IN A JUNGLE. --------------- ENTER COMMAND?SOUTH YOU ARE IN A JUNGLE.">

**Type `WEST`**

<img src="images/ulysses/0136-1.png" width="400" alt="YOU ARE ON THE BEACH OF THE ISLAND OF STORMS. YOU SEE A CAVE IN THE DISTANCE TO THE NORTH. YOUR SHIP IS ANCHORED OFF SHORE."><br>
<img src="images/ulysses/0136.png" width="400" alt="STORMS. YOU SEE A CAVE IN THE DISTANCE TO THE NORTH. YOUR SHIP IS ANCHORED OFF SHORE.">

**Type `GO SHIP`**

### 7. Neptune and the sirens

<img src="images/ulysses/0137.png" width="400" alt="--------------- ENTER COMMAND?GO SHIP YOU ARE JUST OFF-SHORE OF THE ISLAND OF STORMS.">

**Type `NORTH`**

<img src="images/ulysses/0138.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE ON THE OCEAN. THE ISLAND OF STORMS IS TO THE SOUTH OF YOU.">

**Type `NORTH`**

<img src="images/ulysses/0139.png" width="400" alt="STORMS IS TO THE SOUTH OF YOU. --------------- ENTER COMMAND?NORTH YOU ARE ON THE OCEAN.">

**Type `WEST`**

<img src="images/ulysses/0140.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE ON THE OCEAN. YOU SEE A PASSAGEWAY TO THE NORTH.">

**Type `NORTH`**

<img src="images/ulysses/0141.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN A PASS. HIGH CLIFFS ARE TO THE EAST AND WEST.">

**Type `NORTH`**

<img src="images/ulysses/0142-1.png" width="400" alt="YOU HEAR A THUNDERING AVALANCHE BEHIND YOU. SUDDENLY KING NEPTUNE RISES FROM THE OCEAN DEPTHS. YOU ARE IN THE OCEAN."><br>
<img src="images/ulysses/0142.png" width="400" alt="YOU. SUDDENLY KING NEPTUNE RISES FROM THE OCEAN DEPTHS. YOU ARE IN THE OCEAN.">

**Type `POUR WATER`**

<img src="images/ulysses/0143.png" width="400" alt="WHERE? YOU ARE IN THE OCEAN.">

**Type `IN OCEAN`**

<img src="images/ulysses/0144-1.png" width="400" alt="AS YOU POUR THE MAGIC POTION INTO THE OCEAN, NEPTUNE GASPS AND IN A GREAT UPHEAVEL, PLUNGES TO THE BOTTOM OF THE SEA."><br>
<img src="images/ulysses/0144.png" width="400" alt="UPHEAVEL, PLUNGES TO THE BOTTOM OF THE SEA. YOU ARE IN THE OCEAN.">

**Type `NORTH`**

<img src="images/ulysses/0145.png" width="400" alt="YOU ARE IN THE OCEAN. --------------- ENTER COMMAND?NORTH YOU ARE LOST IN THE OCEAN.">

**Type `EAST`**

<img src="images/ulysses/0146.png" width="400" alt="YOU ARE LOST IN THE OCEAN. --------------- ENTER COMMAND?EAST YOU ARE LOST ON THE OCEAN.">

**Type `EAST`**

<img src="images/ulysses/0147-1.png" width="400" alt="YOU HAVE SAILED INTO ENCHANTED WATERS. YOU HEAR THE WIND, IT SEEMS TO SAY SOMETHING,&#34;BEWARE THE SIRENS, E-1 N-1.&#34; YOU ARE ON AN OCEAN. YOU SEE THE ISLAND"><br>
<img src="images/ulysses/0147.png" width="400" alt="SOMETHING,&#34;BEWARE THE SIRENS, E-1 N-1.&#34; YOU ARE ON AN OCEAN. YOU SEE THE ISLAND OF THE SIRENS&#39; TO THE NORTH.">

**Type `NORTH`**

<img src="images/ulysses/0148-1.png" width="400" alt="THE SONG OF THE SIRENS&#39; IS DRAWING YOUR SHIP TOWARD THE JAGGED ROCKS ON SHORE. YOU ARE WEST OF THE ISLAND OF THE SIRENS&#39;."><br>
<img src="images/ulysses/0148.png" width="400" alt="SHIP TOWARD THE JAGGED ROCKS ON SHORE. YOU ARE WEST OF THE ISLAND OF THE SIRENS&#39;.">

**Type `TIE ROPE`**

<img src="images/ulysses/0149.png" width="400" alt="TO WHAT? YOU ARE WEST OF THE ISLAND OF THE SIRENS&#39;.">

**Type `TO ME`**

<img src="images/ulysses/0150.png" width="400" alt="DO YOU MEAN TIE YOURSELF TO THE MAST? YOU ARE WEST OF THE ISLAND OF THE SIRENS&#39;.">

**Type `YES`**

<img src="images/ulysses/0151.png" width="400" alt="O.K. YOU ARE WEST OF THE ISLAND OF THE SIRENS&#39;.">

**Type `HOLD WAX`**

<img src="images/ulysses/0152-1.png" width="400" alt="O.K. THE WAX IS NOW SOFT. WHAT DO YOU WANT TO DO WITH IT? YOU ARE WEST OF THE ISLAND OF THE SIRENS&#39;."><br>
<img src="images/ulysses/0152.png" width="400" alt="WANT TO DO WITH IT? YOU ARE WEST OF THE ISLAND OF THE SIRENS&#39;.">

**Type `WAX EARS`**

### 8. The island of the Cyclops

<img src="images/ulysses/0153-1.png" width="400" alt="YOU HAVE PUT WAX IN THE EARS OF ALL YOUR MEN. DO YOU WANT SOME? BOUND TO THE MAST YOU ARE SAFE. THE SONG OF THE SIRENS&#39; IS TERRIFYING, YET"><br>
<img src="images/ulysses/0153-2.png" width="400" alt="YOU STRUGGLE TO GO TO THEM. THEY MIMIC LOVED ONES WHO HAVE PASSED-ON AND REPEAT A RHYME. &#34;MANY SEARCH BUT NONE WILL FIND,"><br>
<img src="images/ulysses/0153-3.png" width="400" alt="PRICELESS TREASURE LEFT BEHIND. PRISON STEEP, THE KEEPER&#39;S CRUEL, THEY BUILT THE KEY OF SUPPELTUEL.&#34; YOU HAVE DRIFTED SAFELY AWAY FROM THE"><br>
<img src="images/ulysses/0153-4.png" width="400" alt="ISLAND. YOU AND YOUR MEN ASSUME YOUR NORMAL BUSINESS AND DISCARD THE WAX. YOU ARE IN THE OCEAN. THE ISLAND OF THE SIRENS&#39; IS SAFELY BEHIND YOU."><br>
<img src="images/ulysses/0153.png" width="400" alt="NORMAL BUSINESS AND DISCARD THE WAX. YOU ARE IN THE OCEAN. THE ISLAND OF THE SIRENS&#39; IS SAFELY BEHIND YOU.">

**Type `WEST`**

<img src="images/ulysses/0154.png" width="400" alt="SIRENS&#39; IS SAFELY BEHIND YOU. --------------- ENTER COMMAND?WEST YOU ARE LOST IN THE OCEAN.">

**Type `NORTH`**

<img src="images/ulysses/0155.png" width="400" alt="YOU ARE LOST IN THE OCEAN. --------------- ENTER COMMAND?NORTH YOU ARE ON AN OCEAN.">

**Type `NORTH`**

<img src="images/ulysses/0156.png" width="400" alt="YOU ARE ON AN OCEAN. --------------- ENTER COMMAND?NORTH YOU ARE LOST IN THE OCEAN.">

**Type `WEST`**

<img src="images/ulysses/0157-1.png" width="400" alt="SUDDENLY THE OCEAN GROWS UNUSUALLY CALM. YOU ARE WANDERING IN THE OCEAN. YOU SEE AN ISLAND TO THE NORTH."><br>
<img src="images/ulysses/0157.png" width="400" alt="CALM. YOU ARE WANDERING IN THE OCEAN. YOU SEE AN ISLAND TO THE NORTH.">

**Type `NORTH`**

<img src="images/ulysses/0158.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN AN OCEAN. YOU SEE AN ISLAND TO THE NORTH.">

**Type `NORTH`**

<img src="images/ulysses/0159.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE ON THE OCEAN. YOU SEE AN ISLAND TO THE NORTH.">

**Type `NORTH`**

<img src="images/ulysses/0160.png" width="400" alt="TO THE NORTH. --------------- ENTER COMMAND?NORTH YOU ARE OFF-SHORE A COLOSSAL ISLAND.">

**Type `GO ISLAND`**

<img src="images/ulysses/0161-1.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE ON A SMALL BEACH. THE OCEAN IS TO YOUR SOUTH. THE JUNGLE IS TOO OVER GROWN TO SEE A CLEAR PATH."><br>
<img src="images/ulysses/0161.png" width="400" alt="YOU ARE ON A SMALL BEACH. THE OCEAN IS TO YOUR SOUTH. THE JUNGLE IS TOO OVER GROWN TO SEE A CLEAR PATH.">

**Type `WEST`**

<img src="images/ulysses/0162.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE ON AN ISLAND BEACH. THE OCEAN IS TO YOUR SOUTH.">

**Type `WEST`**

<img src="images/ulysses/0163-1.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE ON THE EDGE OF A BEACH. THE OCEAN IS TO YOUR SOUTH. DENSE JUNGLE SEEMS TO ENGULF THE ISLAND."><br>
<img src="images/ulysses/0163.png" width="400" alt="YOU ARE ON THE EDGE OF A BEACH. THE OCEAN IS TO YOUR SOUTH. DENSE JUNGLE SEEMS TO ENGULF THE ISLAND.">

**Type `NORTH`**

<img src="images/ulysses/0164.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE IN A SMALL CLEARING OF THE JUNGLE. THERE IS A LARGE TREE HERE.">

**Type `LOOK TREE`**

<img src="images/ulysses/0165-1.png" width="400" alt="THERE IS SOMETHING CARVED ON ITS TRUNK. YOUR CREW IS FOLLOWING YOU. YOU ARE IN A SMALL CLEARING OF THE JUNGLE. THERE IS A LARGE TREE HERE."><br>
<img src="images/ulysses/0165.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE IN A SMALL CLEARING OF THE JUNGLE. THERE IS A LARGE TREE HERE.">

**Type `READ CARVING`**

<img src="images/ulysses/0166-1.png" width="400" alt="IT SAYS,&#34;SVENEESAS&#34;. YOUR CREW IS FOLLOWING YOU. YOU ARE IN A SMALL CLEARING OF THE JUNGLE. THERE IS A LARGE TREE HERE."><br>
<img src="images/ulysses/0166.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE IN A SMALL CLEARING OF THE JUNGLE. THERE IS A LARGE TREE HERE.">

**Type `EAST`**

<img src="images/ulysses/0167.png" width="400" alt="--------------- ENTER COMMAND?EAST YOUR CREW IS FOLLOWING YOU. YOU ARE IN A DENSE JUNGLE.">

**Type `EAST`**

<img src="images/ulysses/0168.png" width="400" alt="--------------- ENTER COMMAND?EAST YOUR CREW IS FOLLOWING YOU. YOU ARE IN A DENSE JUNGLE.">

**Type `NORTH`**

<img src="images/ulysses/0169.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOUR CREW IS FOLLOWING YOU. YOU ARE IN A JUNGLE.">

**Type `EAST`**

<img src="images/ulysses/0170.png" width="400" alt="--------------- ENTER COMMAND?EAST YOUR CREW IS FOLLOWING YOU. YOU ARE LOST IN A JUNGLE.">

**Type `NORTH`**

<img src="images/ulysses/0171-1.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOUR MEN ARE GETTING VERY HUNGRY AND ARE STARTING TO GRUMBLE. I THINK YOU HAD BETTER FIND SOME FOOD."><br>
<img src="images/ulysses/0171.png" width="400" alt="HAD BETTER FIND SOME FOOD. YOU ARE IN A JUNGLE. YOU SEE A CLEARING TO THE WEST.">

**Type `WEST`**

<img src="images/ulysses/0172.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE STANDING IN SOME RUINS. THERE IS A LARGE CAGE HERE.">

**Type `SAY SEVENSEAS`**

<img src="images/ulysses/0173-1.png" width="400" alt="AS YOU SPEAK THE WORD, THE SKIES THUNDER AND THE HARPES SCATTER, FEARING FOR THEIR LIVES. YOUR CREW IS FOLLOWING YOU."><br>
<img src="images/ulysses/0173.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE STANDING IN SOME RUINS. THERE IS A LARGE CAGE HERE.">

**Type `OPEN CAGE`**

<img src="images/ulysses/0174-1.png" width="400" alt="O.K. THE MAN IS VERY GRATEFUL TO YOU AND OFFERS YOU THIS ENCHANTED MALLET. YOUR CREW IS FOLLOWING YOU."><br>
<img src="images/ulysses/0174.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE STANDING IN SOME RUINS. THERE IS A LARGE CAGE HERE.">

**Type `GET MALLET`**

<img src="images/ulysses/0175.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE STANDING IN SOME RUINS. THERE IS A LARGE CAGE HERE.">

**Type `EAST`**

<img src="images/ulysses/0176.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE IN A JUNGLE. YOU SEE A CLEARING TO THE WEST.">

**Type `EAST`**

<img src="images/ulysses/0177.png" width="400" alt="--------------- ENTER COMMAND?EAST YOUR CREW IS FOLLOWING YOU. YOU ARE IN A JUNGLE.">

**Type `SOUTH`**

<img src="images/ulysses/0178.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE IN A DENSE JUNGLE. THERE IS A CAVE TO THE EAST.">

**Type `EAST`**

<img src="images/ulysses/0179-1.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU HEAR A SHEEP&#39;S CRY FROM INSIDE THE CAVE. YOU ARE STANDING IN FRONT OF A VERY"><br>
<img src="images/ulysses/0179.png" width="400" alt="CAVE. YOU ARE STANDING IN FRONT OF A VERY LARGE CAVE ENTRANCE.">

**Type `NORTH`**

<img src="images/ulysses/0180-1.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. AS YOU ENTER, YOU HEAR A LOUD LAUGH. A CYCLOPS HAS COME IN AND CAUGHT YOU PILFERING HIS SHEEP. HE BLOCKS THE"><br>
<img src="images/ulysses/0180.png" width="400" alt="OPENING WITH A GIANT BOULDER. YOU ARE INSIDE A GIGANTIC CAVE. THERE IS AN EXIT TO THE SOUTH.">

**Type `GIVE WINE`**

<img src="images/ulysses/0181-1.png" width="400" alt="THE CYCLOPS TAKES THE WINE AND SAMPLES IT CAUTIOUSLY, THEN TAKES IT DOWN IN ONE GULP AND SAYS,&#34;THIS IS GOOD!! WHAT DO YOU NEED TO MAKE ME MORE?&#34;"><br>
<img src="images/ulysses/0181.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE INSIDE A GIGANTIC CAVE. THERE IS AN EXIT TO THE SOUTH.">

**Type `GRAPES`**

<img src="images/ulysses/0182-1.png" width="400" alt="THE CYCLOPS SAYS,&#34;I WILL RETURN SHORTLY.&#34; THEN LEAVES BLOCKING THE ENTRANCE BEHIND HIM. THERE IS A SMALL TREE TRUNK ON THE"><br>
<img src="images/ulysses/0182-2.png" width="400" alt="FLOOR. YOUR CREW IS FOLLOWING YOU. YOU ARE INSIDE A GIGANTIC CAVE. THERE IS AN EXIT TO THE SOUTH."><br>
<img src="images/ulysses/0182.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE INSIDE A GIGANTIC CAVE. THERE IS AN EXIT TO THE SOUTH.">

**Type `GET TRUNK`**

<img src="images/ulysses/0183.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE INSIDE A GIGANTIC CAVE. THERE IS AN EXIT TO THE SOUTH.">

**Type `SHARPEN TRUNK`**

<img src="images/ulysses/0184-1.png" width="400" alt="USING YOUR SWORD, YOU SHARPEN THE END OF THE TREE TRUNK TO A POINT AND AWAIT THE CYCLOPS&#39; RETURN. YOUR CREW IS FOLLOWING YOU."><br>
<img src="images/ulysses/0184.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE INSIDE A GIGANTIC CAVE. THERE IS AN EXIT TO THE SOUTH.">

**Type `LOOK`**

<img src="images/ulysses/0185-1.png" width="400" alt="YOU SEE NOTHING SPECIAL. YOUR CREW IS FOLLOWING YOU. YOU ARE INSIDE A GIGANTIC CAVE. THERE IS AN EXIT TO THE SOUTH."><br>
<img src="images/ulysses/0185.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE INSIDE A GIGANTIC CAVE. THERE IS AN EXIT TO THE SOUTH.">

**Type `LOOK`**

<img src="images/ulysses/0186-1.png" width="400" alt="YOU SEE NOTHING SPECIAL. YOUR CREW IS FOLLOWING YOU. THE CYCLOPS RETURNS AND GIVES YOU GRAPES TO MAKE HIM WINE WITH, THEN"><br>
<img src="images/ulysses/0186.png" width="400" alt="STANDS THERE IMPATIENTLY. YOU ARE INSIDE A GIGANTIC CAVE. THERE IS AN EXIT TO THE SOUTH.">

**Type `MAKE WINE`**

<img src="images/ulysses/0187-1.png" width="400" alt="THE CYCLOPS GULPS IT DOWN AND ROARS, &#34;MORE WINE.&#34; YOUR CREW IS FOLLOWING YOU. YOU ARE INSIDE A GIGANTIC CAVE. THERE"><br>
<img src="images/ulysses/0187.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE INSIDE A GIGANTIC CAVE. THERE IS AN EXIT TO THE SOUTH.">

**Type `MAKE WINE`**

<img src="images/ulysses/0188-1.png" width="400" alt="THE CYCLOPS GULPS IT DOWN AND ROARS, &#34;MORE WINE.&#34; YOUR CREW IS FOLLOWING YOU. YOU ARE INSIDE A GIGANTIC CAVE. THERE"><br>
<img src="images/ulysses/0188.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE INSIDE A GIGANTIC CAVE. THERE IS AN EXIT TO THE SOUTH.">

**Type `USE TRUNK`**

<img src="images/ulysses/0189-1.png" width="400" alt="HOW? YOUR CREW IS FOLLOWING YOU. YOU ARE INSIDE A GIGANTIC CAVE. THERE IS AN EXIT TO THE SOUTH."><br>
<img src="images/ulysses/0189.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE INSIDE A GIGANTIC CAVE. THERE IS AN EXIT TO THE SOUTH.">

**Type `IN EYE`**

<img src="images/ulysses/0190-1.png" width="400" alt="YOU AND YOUR MEN PLUNGE THE SHARPENED TREE INTO THE CYCLOPS&#39; EYE. IN A BLIND RAGE HE PICKS UP THE BOULDER AND HEAVES IT AT YOU, BUT MISSING. HE THEN"><br>
<img src="images/ulysses/0190-2.png" width="400" alt="STUMBLES OUT OF THE CAVE INTO THE JUNGLE. YOUR CREW IS FOLLOWING YOU. A FEW SHEEP ARE FOLLOWING YOU."><br>
<img src="images/ulysses/0190.png" width="400" alt="A FEW SHEEP ARE FOLLOWING YOU. YOU ARE INSIDE A GIGANTIC CAVE. THERE IS AN EXIT TO THE SOUTH.">

**Type `KILL SHEEP`**

<img src="images/ulysses/0191-1.png" width="400" alt="WITH YOUR SWORD, YOU KILL THE SHEEP AND PREPARE THEM FOR EATING. THERE ARE SOME DEAD SHEEP HERE. YOUR CREW IS FOLLOWING YOU."><br>
<img src="images/ulysses/0191.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE INSIDE A GIGANTIC CAVE. THERE IS AN EXIT TO THE SOUTH.">

**Type `START FIRE`**

<img src="images/ulysses/0192-1.png" width="400" alt="USING THE FLINT AND SOME WOODEN PLANKS YOU START A ROARING FIRE. THERE ARE SOME DEAD SHEEP HERE. YOUR CREW IS FOLLOWING YOU."><br>
<img src="images/ulysses/0192.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE INSIDE A GIGANTIC CAVE. THERE IS AN EXIT TO THE SOUTH.">

**Type `COOK SHEEP`**

<img src="images/ulysses/0193-1.png" width="400" alt="YOU AND YOUR MEN ROAST THE SHEEP OVER AN OPEN FIRE. YOUR CREW IS FOLLOWING YOU. YOU ARE INSIDE A GIGANTIC CAVE. THERE"><br>
<img src="images/ulysses/0193.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE INSIDE A GIGANTIC CAVE. THERE IS AN EXIT TO THE SOUTH.">

**Type `EAT SHEEP`**

### 9. The skeletons, the cliff and Pegasus

<img src="images/ulysses/0194-1.png" width="400" alt="YOU AND YOUR MEN GORGE YOURSELVES ON ROAST MUTTON. COMPLETELY SATISFIED, YOU CONTINUE ON. YOUR CREW IS FOLLOWING YOU."><br>
<img src="images/ulysses/0194.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE INSIDE A GIGANTIC CAVE. THERE IS AN EXIT TO THE SOUTH.">

**Type `SOUTH`**

<img src="images/ulysses/0195.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE STANDING IN FRONT OF A VERY LARGE CAVE ENTRANCE.">

**Type `SOUTH`**

<img src="images/ulysses/0196.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOUR CREW IS FOLLOWING YOU. YOU ARE IN A DENSE JUNGLE.">

**Type `EAST`**

<img src="images/ulysses/0197.png" width="400" alt="--------------- ENTER COMMAND?EAST YOUR CREW IS FOLLOWING YOU. YOU ARE STUMBLING THROUGH THE JUNGLE.">

**Type `NORTH`**

<img src="images/ulysses/0198.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. THERE IS A BAND OF SKELETONS HERE. YOU ARE IN A LARGE CLEARING.">

**Type `SAY ECEELF`**

<img src="images/ulysses/0199-1.png" width="400" alt="THE CHEST VIBRATES SLIGHTLY, THEN THE LID POPS OPEN. YOUR CREW IS FOLLOWING YOU. THERE IS A BAND OF SKELETONS FOLLOWING"><br>
<img src="images/ulysses/0199.png" width="400" alt="THERE IS A BAND OF SKELETONS FOLLOWING YOU. YOU ARE IN A LARGE CLEARING.">

**Type `LOOK CHEST`**

<img src="images/ulysses/0200-1.png" width="400" alt="THERE IS AN ENCHANTED SWORD HERE. YOUR CREW IS FOLLOWING YOU. THERE IS A BAND OF SKELETONS FOLLOWING YOU."><br>
<img src="images/ulysses/0200.png" width="400" alt="THERE IS A BAND OF SKELETONS FOLLOWING YOU. YOU ARE IN A LARGE CLEARING.">

**Type `GET SWORD`**

<img src="images/ulysses/0201-1.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. THERE IS A BAND OF SKELETONS FOLLOWING YOU. YOU ARE IN A LARGE CLEARING."><br>
<img src="images/ulysses/0201.png" width="400" alt="THERE IS A BAND OF SKELETONS FOLLOWING YOU. YOU ARE IN A LARGE CLEARING.">

**Type `KILL SKELETONS`**

<img src="images/ulysses/0202-1.png" width="400" alt="WITH WHAT? YOUR CREW IS FOLLOWING YOU. THERE IS A BAND OF SKELETONS FOLLOWING YOU."><br>
<img src="images/ulysses/0202.png" width="400" alt="THERE IS A BAND OF SKELETONS FOLLOWING YOU. YOU ARE IN A LARGE CLEARING.">

**Type `WITH SWORD`**

<img src="images/ulysses/0203-1.png" width="400" alt="USING THE MAGICAL SWORD YOU QUICKLY HACK THE SKELETONS TO PIECES. THEIR BONES TURN TO DUST AND BLOW AWAY. YOUR CREW IS FOLLOWING YOU."><br>
<img src="images/ulysses/0203.png" width="400" alt="BONES TURN TO DUST AND BLOW AWAY. YOUR CREW IS FOLLOWING YOU. YOU ARE IN A LARGE CLEARING.">

**Type `EAST`**

<img src="images/ulysses/0204.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE IN A JUNGLE. THERE ARE CLIFFS FAR OFF TO THE NORTH.">

**Type `NORTH`**

<img src="images/ulysses/0205.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE IN A JUNGLE. THERE IS A CLIFF JUST AHEAD OF YOU.">

**Type `NORTH`**

<img src="images/ulysses/0206-1.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU HEAR A HORSE&#39;S NEIGH NEARBY. YOU ARE IN FRONT OF AN UNCLIMBABLE CLIFF."><br>
<img src="images/ulysses/0206.png" width="400" alt="YOU HEAR A HORSE&#39;S NEIGH NEARBY. YOU ARE IN FRONT OF AN UNCLIMBABLE CLIFF.">

**Type `SUPPELTUDE`**

<img src="images/ulysses/0207-1.png" width="400" alt="YOU HEAR A DEAFENING ROAR, THEN THE CLIFF OPENS BEFORE YOU. YOUR CREW IS FOLLOWING YOU. YOU ARE IN FRONT OF AN UNCLIMBABLE"><br>
<img src="images/ulysses/0207.png" width="400" alt="YOUR CREW IS FOLLOWING YOU. YOU ARE IN FRONT OF AN UNCLIMBABLE CLIFF.">

**Type `NORTH`**

<img src="images/ulysses/0208-1.png" width="400" alt="YOUR CREW STAYS BEHIND. THERE IS A LOUD RUMBLE, AND THE CLIFF CLOSES FOREVER. YOU ARE IN A SMALL VALLEY SURROUNDED BY"><br>
<img src="images/ulysses/0208.png" width="400" alt="YOU ARE IN A SMALL VALLEY SURROUNDED BY STEEP CLIFFS. THERE IS A TREE HERE SURROUNDED BY GIANT THORNS.">

**Type `GIVE REINS`**

<img src="images/ulysses/0209-1.png" width="400" alt="O.K. YOU ARE IN A SMALL VALLEY SURROUNDED BY STEEP CLIFFS. THERE IS A TREE HERE SURROUNDED BY GIANT THORNS."><br>
<img src="images/ulysses/0209.png" width="400" alt="YOU ARE IN A SMALL VALLEY SURROUNDED BY STEEP CLIFFS. THERE IS A TREE HERE SURROUNDED BY GIANT THORNS.">

**Type `GIVE BRIDLE`**

<img src="images/ulysses/0210-1.png" width="400" alt="O.K. YOU ARE IN A SMALL VALLEY SURROUNDED BY STEEP CLIFFS. THERE IS A TREE HERE SURROUNDED BY GIANT THORNS."><br>
<img src="images/ulysses/0210.png" width="400" alt="YOU ARE IN A SMALL VALLEY SURROUNDED BY STEEP CLIFFS. THERE IS A TREE HERE SURROUNDED BY GIANT THORNS.">

**Type `USE HAMMER`**

<img src="images/ulysses/0211-1.png" width="400" alt="AS YOU STRIKE THE CHAIN WITH THE ENCHANTED MALLET, THEY BOTH DISAPPEAR. YOU ARE IN A SMALL VALLEY SURROUNDED BY STEEP CLIFFS. THERE IS A TREE HERE"><br>
<img src="images/ulysses/0211.png" width="400" alt="YOU ARE IN A SMALL VALLEY SURROUNDED BY STEEP CLIFFS. THERE IS A TREE HERE SURROUNDED BY GIANT THORNS.">

**Type `MOUNT HORSE`**

<img src="images/ulysses/0212-1.png" width="400" alt="AS YOU MOUNT PEGASUS, HE TAKES OFF WITH A GREAT LEAP, FLYING YOU OVER THE THORNS AND SETS DOWN NEXT TO THE TREE. YOU ARE IN A SMALL VALLEY NEXT TO A"><br>
<img src="images/ulysses/0212.png" width="400" alt="THORNS AND SETS DOWN NEXT TO THE TREE. YOU ARE IN A SMALL VALLEY NEXT TO A TREE.">

**Type `GET FLEECE`**

<img src="images/ulysses/0213.png" width="400" alt="YOU ARE IN A SMALL VALLEY NEXT TO A TREE.">

**Type `MOUNT HORSE`**

<img src="images/ulysses/0214-1.png" width="400" alt="PEGASUS SWIFTLY LEAPS INTO THE AIR, CARRYING YOU SAFELY HOME TO THE KING THEN LEAVES, WHILE YOUR CREW FOLLOWS BY SEA."><br>
<img src="images/ulysses/0214.png" width="400" alt="SEA. YOU ARE OUTSIDE OF THE KING&#39;S CASTLE FACING WEST.">

**Type `WEST`**

<img src="images/ulysses/0215.png" width="400" alt="YOU ARE IN THE ENTRY HALL OF THE KING&#39;S CASTLE. THERE IS A GUARD HERE. THERE IS AN EXIT TO THE EAST.">

**Type `YES`**

<img src="images/ulysses/0216-1.png" width="400" alt="&#34;GOOD. THE KING IS EXPECTING YOU. I WILL ESCORT YOU.&#34; YOU ARE IN THE KING&#39;S THRONE ROOM STANDING BEFORE HIS MAJESTY."><br>
<img src="images/ulysses/0216.png" width="400" alt="WILL ESCORT YOU.&#34; YOU ARE IN THE KING&#39;S THRONE ROOM STANDING BEFORE HIS MAJESTY.">

**Type `GIVE FLEECE`**

### The end

<img src="images/ulysses/0217-1.png" width="400" alt="THE KING DELIGHTEDLY TAKES THE FLEECE. HE AWARDS YOU A KINGDOM OF YOUR OWN AND 300 BAGS OF GOLD. CONGRATULATIONS!!!! YOU HAVE"><br>
<img src="images/ulysses/0217-2.png" width="400" alt="SUCCESSFULLY COMPLETED &#34;ULYSSES AND THE GOLDEN FLEECE&#34; AND ARE HEREBY DECLARED A LEVEL-2 ADVENTURER. THANK YOU FOR PLAYING &#34;ULYSSES AND THE"><br>
<img src="images/ulysses/0217.png" width="400" alt="THANK YOU FOR PLAYING &#34;ULYSSES AND THE GOLDEN FLEECE.&#34; ]">

<!-- End of the walkthrough -->

## What next

[Time Zone](timezone.md) is the next of the Hi-Res Adventures, the fifth,
and the biggest.
