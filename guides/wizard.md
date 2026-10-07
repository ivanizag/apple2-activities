# The Wizard and the Princess, from the start to the end

[Back to the activities](../README.md)

**The Wizard and the Princess**, *Hi-Res Adventure #2* of On-Line Systems,
came out in 1980 for the Apple \]\[, written by Ken and Roberta Williams.
After the lines of [Mystery House](mysteryhouse.md), its pictures are painted
in colour: a desert, a sea, an island, mountains and a castle, on the way to
bring a princess back to the village of Serenia.

You type one or two words, `GET ROCK`, `THROW ROCK`, `KISS FROG`, and the game
answers under the picture. The whole game is on one side of a diskette.

This page plays it to its end, all 176 commands, with a picture of the screen
after each one: 203 pictures. The pictures are of a colour television, with
the four lines of text under them drawn white and sharp, as a monochrome
monitor showed them: the television blurred them into fringes of colour,
hard to read.

## What you need

**`The Wizard and the Princess (1980-On-Line Systems).nib`**, from
[the Asimov archive](https://mirrors.apple2.org.za/ftp.apple.asimov.net/images/games/adventure/wizard_and_the_princess/):
the disk of On-Line Systems as it was, track by track. `./fetch-disks.sh` in
this repository downloads it into `disks/` and checks it. The same folder has
a crack of an edition of Green Valley Publishing, whose ending thanks you in
its name instead.

The commands of this page are also in this repository, a line each,
[listings/wizard.txt](listings/wizard.txt). They follow the walkthrough of
[The Sierra Help Pages](https://sierrahelp.com/Walkthroughs/WizardAndThePrincessWalkthrough.html),
with what the machine showed it needs: the rope is tied to the anchor with
`TIE ROPE` and `TO ANCHOR`, where it says `USE ROPE`.

The [manual of The Wizard and the Princess](https://archive.org/details/The_Wizard_And_The_Princess_Hi-Res_Adventure_2)
is on the Internet Archive.

## The machine

An Apple \]\[+:

- the 6502 processor at 1 MHz and 48 KB of memory;
- a Disk II controller card in slot 6, with the game in drive 1.

```bash
izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s6 'diskii,disk1=disks/The Wizard and the Princess (1980-On-Line Systems).nib'
```

Shorter, the model `2plus` of izapple2 plays it too: the same Apple \]\[+,
with the 16 KB of a language card in slot 0 and a Videx 80 column card in
slot 3 more, on a colour monitor with its scan lines.

```bash
izapple2 -model 2plus 'disks/The Wizard and the Princess (1980-On-Line Systems).nib'
```

## Playing it

1. **Start izapple2** with the command above. The game starts in the village
   of Serenia.

2. **Type the commands**, each followed by Return. Each step below is what
   the game shows, and the command to type then. When the picture stays the
   same and only the text goes on, the page shows the picture once, and
   under it the lines of text that came after. When the game stops in the
   middle of a longer text, **press Return** to go on.

## The walkthrough

<!-- The walkthrough, written by the generator -->

- [1. The desert of Serenia](#1-the-desert-of-serenia)
- [2. Over the chasm](#2-over-the-chasm)
- [3. The parrot and the lion](#3-the-parrot-and-the-lion)
- [4. The boat](#4-the-boat)
- [5. The pirate's chest](#5-the-pirates-chest)
- [6. The ring, the coin and the giant](#6-the-ring-the-coin-and-the-giant)
- [7. The horn and the maze](#7-the-horn-and-the-maze)
- [8. The castle of Harlin](#8-the-castle-of-harlin)
- [The end](#the-end)

### 1. The desert of Serenia

<img src="images/wizard/0001.png" width="400" alt="YOU ARE IN THE VILLAGE OF SERENIA. AROUND YOU IS A DESERT.">

**Type `SOUTH`**

<img src="images/wizard/0002.png" width="400" alt="AROUND YOU IS A DESERT. --------------- ENTER COMMAND?SOUTH YOU ARE LOST IN THE DESERT">

**Type `SOUTH`**

<img src="images/wizard/0003.png" width="400" alt="YOU ARE LOST IN THE DESERT --------------- ENTER COMMAND?SOUTH YOU ARE LOST IN THE DESERT">

**Type `EAST`**

<img src="images/wizard/0004.png" width="400" alt="YOU ARE LOST IN THE DESERT --------------- ENTER COMMAND?EAST YOU ARE LOST IN THE DESERT">

**Type `SOUTH`**

<img src="images/wizard/0005.png" width="400" alt="YOU ARE LOST IN THE DESERT --------------- ENTER COMMAND?SOUTH YOU ARE LOST IN THE DESERT">

**Type `EAST`**

<img src="images/wizard/0006.png" width="400" alt="YOU ARE LOST IN THE DESERT --------------- ENTER COMMAND?EAST YOU ARE LOST IN THE DESERT">

**Type `SOUTH`**

<img src="images/wizard/0007.png" width="400" alt="YOU ARE LOST IN THE DESERT --------------- ENTER COMMAND?SOUTH YOU ARE LOST IN THE DESERT">

**Type `GET ROCK`**

<img src="images/wizard/0008.png" width="400" alt="YOU ARE LOST IN THE DESERT --------------- ENTER COMMAND?GET ROCK YOU ARE LOST IN THE DESERT">

**Type `NORTH`**

<img src="images/wizard/0009.png" width="400" alt="YOU ARE LOST IN THE DESERT --------------- ENTER COMMAND?NORTH YOU ARE LOST IN THE DESERT">

**Type `WEST`**

<img src="images/wizard/0010.png" width="400" alt="YOU ARE LOST IN THE DESERT --------------- ENTER COMMAND?WEST YOU ARE LOST IN THE DESERT">

**Type `NORTH`**

<img src="images/wizard/0011.png" width="400" alt="YOU ARE LOST IN THE DESERT --------------- ENTER COMMAND?NORTH YOU ARE LOST IN THE DESERT">

**Type `NORTH`**

<img src="images/wizard/0012.png" width="400" alt="YOU ARE LOST IN THE DESERT --------------- ENTER COMMAND?NORTH YOU ARE IN THE DESERT">

**Type `THROW ROCK`**

<img src="images/wizard/0013.png" width="400" alt="YOU HIT THE SNAKE ON THE HEAD WITH A ROCK AND KILL IT. YOU ARE IN THE DESERT">

**Type `NORTH`**

<img src="images/wizard/0014.png" width="400" alt="YOU ARE HOT AND THIRSTY. I THINK YOU HAD BETTER DRINK SOME WATER. YOU ARE IN THE DESERT">

**Type `DRINK WATER`**

<img src="images/wizard/0015.png" width="400" alt="THAT IS MUCH BETTER. NOW YOUR FLASK IS EMPTY. YOU ARE IN THE DESERT">

**Type `EAST`**

<img src="images/wizard/0016.png" width="400" alt="YOU ARE IN THE DESERT --------------- ENTER COMMAND?EAST YOU ARE IN THE DESERT">

**Type `GET STICK`**

<img src="images/wizard/0017.png" width="400" alt="YOU ARE IN THE DESERT --------------- ENTER COMMAND?GET STICK YOU ARE IN THE DESERT">

**Type `WEST`**

<img src="images/wizard/0018.png" width="400" alt="YOU ARE IN THE DESERT --------------- ENTER COMMAND?WEST YOU ARE IN THE DESERT">

**Type `WEST`**

<img src="images/wizard/0019.png" width="400" alt="YOU ARE IN THE DESERT --------------- ENTER COMMAND?WEST YOU ARE IN THE DESERT">

**Type `LOOK HOLE`**

<img src="images/wizard/0020.png" width="400" alt="--------------- ENTER COMMAND?LOOK HOLE THERE IS A CRACKER HERE. YOU ARE IN THE DESERT">

**Type `GET CRACKER`**

<img src="images/wizard/0021.png" width="400" alt="--------------- ENTER COMMAND?GET CRACKE R YOU ARE IN THE DESERT">

**Type `NORTH`**

<img src="images/wizard/0022.png" width="400" alt="YOU ARE IN THE DESERT --------------- ENTER COMMAND?NORTH YOU ARE IN THE DESERT">

**Type `GET ROCK`**

<img src="images/wizard/0023-1.png" width="400" alt="WHEN YOU REMOVE THE ROCK FROM THE SNAKE&#39;S TAIL, THE SNAKE LOOKS AT YOU AND SAYS, &#34;I AM THE KING OF THE SNAKES AND TO REPAY YOU, I WILL GIVE YOU A"><br>
<img src="images/wizard/0023.png" width="400" alt="AND TO REPAY YOU, I WILL GIVE YOU A MAGIC WORD.&#34; &#34;HISS.&#34; THEN HE LEAVES. YOU ARE IN THE DESERT">

**Type `EAST`**

<img src="images/wizard/0024.png" width="400" alt="YOU ARE IN THE DESERT --------------- ENTER COMMAND?EAST YOU ARE IN THE DESERT">

**Type `SOUTH`**

<img src="images/wizard/0025.png" width="400" alt="THERE IS A RATTLESNAKE HERE. YOU HAD BETTER WATCH OUT! YOU ARE IN THE DESERT">

**Type `USE STICK`**

<img src="images/wizard/0026.png" width="400" alt="YOU HIT THE SNAKE ON THE HEAD WITH THE STICK AND DRIVE IT AWAY. YOU ARE IN THE DESERT">

**Type `LOOK HOLE`**

<img src="images/wizard/0027.png" width="400" alt="--------------- ENTER COMMAND?LOOK HOLE THERE IS A NOTE HERE. YOU ARE IN THE DESERT">

**Type `GET NOTE`**

<img src="images/wizard/0028.png" width="400" alt="YOU ARE IN THE DESERT --------------- ENTER COMMAND?GET NOTE YOU ARE IN THE DESERT">

**Type `READ NOTE`**

<img src="images/wizard/0029.png" width="400" alt="THERE IS SOME STRANGE WRITING ON THE NOTE. YOU ARE IN THE DESERT">

**Type `DROP NOTE`**

<img src="images/wizard/0030.png" width="400" alt="--------------- ENTER COMMAND?DROP NOTE THERE IS A NOTE HERE. YOU ARE IN THE DESERT">

**Type `NORTH`**

<img src="images/wizard/0031.png" width="400" alt="YOU ARE IN THE DESERT --------------- ENTER COMMAND?NORTH YOU ARE IN THE DESERT">

**Type `NORTH`**

<img src="images/wizard/0032.png" width="400" alt="--------------- ENTER COMMAND?NORTH THERE IS A LOCKET HERE. YOU ARE IN THE DESERT">

**Type `GET LOCKET`**

<img src="images/wizard/0033.png" width="400" alt="--------------- ENTER COMMAND?GET LOCKET YOU ARE IN THE DESERT">

**Type `OPEN LOCKET`**

<img src="images/wizard/0034.png" width="400" alt="T OK. YOU ARE IN THE DESERT">

**Type `LOOK LOCKET`**

<img src="images/wizard/0035.png" width="400" alt="T OK. YOU ARE IN THE DESERT">

**Type `WEST`**

<img src="images/wizard/0036.png" width="400" alt="--------------- ENTER COMMAND?WEST THERE IS A NOTE HERE. YOU ARE IN THE DESERT">

**Type `GET NOTE`**

<img src="images/wizard/0037.png" width="400" alt="YOU ARE IN THE DESERT --------------- ENTER COMMAND?GET NOTE YOU ARE IN THE DESERT">

**Type `READ NOTE`**

### 2. Over the chasm

<img src="images/wizard/0038.png" width="400" alt="THERE IS SOME STRANGE WRITING ON THE NOTE. YOU ARE IN THE DESERT">

**Type `WEST`**

<img src="images/wizard/0039.png" width="400" alt="YOU ARE IN THE DESERT --------------- ENTER COMMAND?WEST YOU ARE IN THE DESERT">

**Type `NORTH`**

<img src="images/wizard/0040.png" width="400" alt="YOU ARE AT THE SOUTH EDGE OF A DEEP CHASM. THERE IS A COTTAGE AND SOME WOODS ON THE OTHER SIDE.">

**Type `HOCUS`**

<img src="images/wizard/0041.png" width="400" alt="YOU ARE AT THE SOUTH EDGE OF A DEEP CHASM. THERE IS A COTTAGE AND SOME WOODS ON THE OTHER SIDE.">

**Type `NORTH`**

<img src="images/wizard/0042.png" width="400" alt="YOU ARE AT THE NORTH EDGE OF A DEEP CHASM. THERE IS A BRIDGE SPANNING THE CHASM.">

**Type `EAST`**

<img src="images/wizard/0043-1.png" width="400" alt="THERE IS AN APPLE HERE. YOU ARE IN THE ONE ROOM COTTAGE. IT IS ALMOST EMPTY EXCEPT FOR A COUPLE OF TABLES."><br>
<img src="images/wizard/0043.png" width="400" alt="YOU ARE IN THE ONE ROOM COTTAGE. IT IS ALMOST EMPTY EXCEPT FOR A COUPLE OF TABLES.">

**Type `GET APPLE`**

<img src="images/wizard/0044.png" width="400" alt="YOU ARE IN THE ONE ROOM COTTAGE. IT IS ALMOST EMPTY EXCEPT FOR A COUPLE OF TABLES.">

**Type `WEST`**

<img src="images/wizard/0045.png" width="400" alt="YOU ARE AT THE NORTH EDGE OF A DEEP CHASM. THERE IS A BRIDGE SPANNING THE CHASM.">

**Type `NORTH`**

<img src="images/wizard/0046.png" width="400" alt="--------------- ENTER COMMAND?NORTH THERE IS A LITTLE GNOME HERE. YOU ARE IN THE WOODS">

**Type `EAST`**

<img src="images/wizard/0047.png" width="400" alt="THE LITTLE GNOME GRABS SOME OF YOUR THINGS AND RUNS AWAY WITH THEM. YOU ARE IN THE WOODS">

**Type `EAST`**

<img src="images/wizard/0048.png" width="400" alt="YOU ARE IN THE WOODS --------------- ENTER COMMAND?EAST YOU ARE IN THE WOODS">

**Type `NORTH`**

<img src="images/wizard/0049.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN THE WOODS. THERE IS A BANK WITH A SMALL CREVICE IN IT.">

**Type `HISS`**

<img src="images/wizard/0050.png" width="400" alt="YOU SUDDENLY TURN INTO A SNAKE! YOU ARE IN THE WOODS. THERE IS A BANK WITH A SMALL CREVICE IN IT.">

**Type `GO CREVICE`**

<img src="images/wizard/0051-1.png" width="400" alt="YOU ARE IN AN UNDERGROUND TUNNEL. THERE IS A SMALL CREVICE GOING TO THE OUTSIDE HERE. SUNLIGHT IS COMING IN THROUGH THE CREVICE."><br>
<img src="images/wizard/0051.png" width="400" alt="IS A SMALL CREVICE GOING TO THE OUTSIDE HERE. SUNLIGHT IS COMING IN THROUGH THE CREVICE.">

**Type `SOUTH`**

<img src="images/wizard/0052.png" width="400" alt="YOU ARE IN AN UNDERGROUND TUNNEL. THERE IS JUST ENOUGH SUNLIGHT COMING THROUGH THE CREVICE TO SEE.">

**Type `SOUTH`**

<img src="images/wizard/0053-1.png" width="400" alt="YOU HAVE CHANGED BACK INTO YOURSELF AGAIN! YOU ARE IN AN UNDERGROUND TUNNEL. THERE IS JUST ENOUGH SUNLIGHT COMING IN"><br>
<img src="images/wizard/0053.png" width="400" alt="YOU ARE IN AN UNDERGROUND TUNNEL. THERE IS JUST ENOUGH SUNLIGHT COMING IN THROUGH THE CREVICE TO SEE.">

**Type `SOUTH`**

<img src="images/wizard/0054-1.png" width="400" alt="YOU ARE IN AN UNDERGROUND TUNNEL. THERE IS JUST ENOUGH SUNLIGHT COMING IN THROUGH THE CREVICE TO SEE. THERE IS A LITTLE DOOR HERE."><br>
<img src="images/wizard/0054.png" width="400" alt="IS JUST ENOUGH SUNLIGHT COMING IN THROUGH THE CREVICE TO SEE. THERE IS A LITTLE DOOR HERE.">

**Type `GET LOCKET`**

<img src="images/wizard/0055-1.png" width="400" alt="YOU ARE IN AN UNDERGROUND TUNNEL. THERE IS JUST ENOUGH SUNLIGHT COMING IN THROUGH THE CREVICE TO SEE. THERE IS A LITTLE DOOR HERE."><br>
<img src="images/wizard/0055.png" width="400" alt="IS JUST ENOUGH SUNLIGHT COMING IN THROUGH THE CREVICE TO SEE. THERE IS A LITTLE DOOR HERE.">

**Type `GET CRACKER`**

<img src="images/wizard/0056-1.png" width="400" alt="YOU ARE IN AN UNDERGROUND TUNNEL. THERE IS JUST ENOUGH SUNLIGHT COMING IN THROUGH THE CREVICE TO SEE. THERE IS A LITTLE DOOR HERE."><br>
<img src="images/wizard/0056.png" width="400" alt="IS JUST ENOUGH SUNLIGHT COMING IN THROUGH THE CREVICE TO SEE. THERE IS A LITTLE DOOR HERE.">

**Type `GET BREAD`**

<img src="images/wizard/0057-1.png" width="400" alt="YOU ARE IN AN UNDERGROUND TUNNEL. THERE IS JUST ENOUGH SUNLIGHT COMING IN THROUGH THE CREVICE TO SEE. THERE IS A LITTLE DOOR HERE."><br>
<img src="images/wizard/0057.png" width="400" alt="IS JUST ENOUGH SUNLIGHT COMING IN THROUGH THE CREVICE TO SEE. THERE IS A LITTLE DOOR HERE.">

**Type `UNLOCK DOOR`**

<img src="images/wizard/0058-1.png" width="400" alt="OK. THE DOOR IS UNLOCKED. YOU ARE IN AN UNDERGROUND TUNNEL. THERE IS JUST ENOUGH SUNLIGHT COMING IN THROUGH THE CREVICE TO SEE. THERE IS A"><br>
<img src="images/wizard/0058.png" width="400" alt="IS JUST ENOUGH SUNLIGHT COMING IN THROUGH THE CREVICE TO SEE. THERE IS A LITTLE DOOR HERE.">

**Type `OPEN DOOR`**

<img src="images/wizard/0059-1.png" width="400" alt="YOU ARE IN AN UNDERGROUND TUNNEL. THERE IS JUST ENOUGH SUNLIGHT COMING IN THROUGH THE CREVICE TO SEE. THERE IS A LITTLE DOOR HERE."><br>
<img src="images/wizard/0059.png" width="400" alt="IS JUST ENOUGH SUNLIGHT COMING IN THROUGH THE CREVICE TO SEE. THERE IS A LITTLE DOOR HERE.">

**Type `GO DOOR`**

<img src="images/wizard/0060.png" width="400" alt="--------------- ENTER COMMAND?GO DOOR YOU ARE AT THE BOTTOM OF THE STAIRS THERE IS A LITTLE DOOR HERE.">

**Type `UP`**

<img src="images/wizard/0061.png" width="400" alt="YOU ARE INSIDE A TREE. THERE ARE STAIRS LEADING DOWNWARD. THERE IS A HOLE LEADIUNG OUTSIDE.">

**Type `GO HOLE`**

### 3. The parrot and the lion

<img src="images/wizard/0062.png" width="400" alt="LEADIUNG OUTSIDE. --------------- ENTER COMMAND?GO HOLE YOU ARE IN THE WOODS">

**Type `EAST`**

<img src="images/wizard/0063.png" width="400" alt="YOU ARE IN THE WOODS --------------- ENTER COMMAND?EAST YOU ARE IN THE WOODS">

**Type `NORTH`**

<img src="images/wizard/0064.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN THE WOODS. THERE IS A PARROT SITTING IN A TREE.">

**Type `GIVE CRACKER`**

<img src="images/wizard/0065-1.png" width="400" alt="THE PARROT EATS THE CRACKER AND IS VERY GRATEFUL. HE SETS A VIAL OF LIQUID ON THE TREE BRANCH FOR YOU. THERE IS A VIAL HERE"><br>
<img src="images/wizard/0065.png" width="400" alt="THERE IS A VIAL HERE YOU ARE IN THE WOODS. THERE IS A PARROT SITTING IN A TREE.">

**Type `GET VIAL`**

<img src="images/wizard/0066.png" width="400" alt="--------------- ENTER COMMAND?GET VIAL YOU ARE IN THE WOODS. THERE IS A PARROT SITTING IN A TREE.">

**Type `SOUTH`**

<img src="images/wizard/0067.png" width="400" alt="SITTING IN A TREE. --------------- ENTER COMMAND?SOUTH YOU ARE IN THE WOODS">

**Type `WEST`**

<img src="images/wizard/0068.png" width="400" alt="YOU ARE IN THE WOODS --------------- ENTER COMMAND?WEST YOU ARE IN THE WOODS">

**Type `WEST`**

<img src="images/wizard/0069.png" width="400" alt="YOU ARE IN THE WOODS --------------- ENTER COMMAND?WEST YOU ARE IN THE WOODS">

**Type `NORTH`**

<img src="images/wizard/0070.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN THE WOODS. THERE IS A BABBLING BROOK HERE.">

**Type `FILL FLASK`**

<img src="images/wizard/0071.png" width="400" alt="YOUR FLASK IS NOW FULL OF WATER. YOU ARE IN THE WOODS. THERE IS A BABBLING BROOK HERE.">

**Type `NORTH`**

<img src="images/wizard/0072.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN THE WOODS. THERE IS A VERY TALL TREE IN FRONT OF YOU.">

**Type `UP`**

<img src="images/wizard/0073.png" width="400" alt="--------------- ENTER COMMAND?UP YOU ARE AT THE TOP OF A TREE. YOU SEE AN OCEAN IN THE DISTANCE.">

**Type `DOWN`**

<img src="images/wizard/0074.png" width="400" alt="--------------- ENTER COMMAND?DOWN YOU ARE IN THE WOODS. THERE IS A VERY TALL TREE IN FRONT OF YOU.">

**Type `SOUTH`**

<img src="images/wizard/0075.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN THE WOODS. THERE IS A BABBLING BROOK HERE.">

**Type `SOUTH`**

<img src="images/wizard/0076.png" width="400" alt="BABBLING BROOK HERE. --------------- ENTER COMMAND?SOUTH YOU ARE IN THE WOODS">

**Type `WEST`**

<img src="images/wizard/0077.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE AT THE EDGE OF THE WOODS. YOU SEE THE OCEAN.">

**Type `GIVE BREAD`**

### 4. The boat

<img src="images/wizard/0078-1.png" width="400" alt="THE LION WOLFS DOWN THE BREAD AND THEN WALKS AWAY. YOU ARE AT THE EDGE OF THE WOODS. YOU SEE THE OCEAN."><br>
<img src="images/wizard/0078.png" width="400" alt="WALKS AWAY. YOU ARE AT THE EDGE OF THE WOODS. YOU SEE THE OCEAN.">

**Type `NORTH`**

<img src="images/wizard/0079.png" width="400" alt="THERE IS A ROPE HERE. YOU ARE AT THE EDGE OF THE OCEAN ON A BEACH. THERE IS A ROWBOAT HERE.">

**Type `GET ROPE`**

<img src="images/wizard/0080.png" width="400" alt="--------------- ENTER COMMAND?GET ROPE YOU ARE AT THE EDGE OF THE OCEAN ON A BEACH. THERE IS A ROWBOAT HERE.">

**Type `GO BOAT`**

<img src="images/wizard/0081.png" width="400" alt="--------------- ENTER COMMAND?GO BOAT THERE IS A HOLE IN THE BOAT YOU ARE IN THE ROWBOAT ON THE BEACH.">

**Type `USE BLANKET`**

<img src="images/wizard/0082.png" width="400" alt="YOU STUFF THE BLANKET INTO THE HOLE. LET&#39;S HOPE IT WORKS. YOU ARE IN THE ROWBOAT ON THE BEACH.">

**Type `NORTH`**

<img src="images/wizard/0083.png" width="400" alt="YOU ARE IN THE ROWBOAT ON THE BEACH. --------------- ENTER COMMAND?NORTH YOU ARE IN THE MIDDLE OF THE OCEAN.">

**Type `NORTH`**

<img src="images/wizard/0084.png" width="400" alt="YOU ARE HOT AND THIRSTY. I THINK YOU HAD BETTER DRINK SOME WATER. YOU ARE IN THE MIDDLE OF THE OCEAN.">

**Type `DRINK WATER`**

<img src="images/wizard/0085.png" width="400" alt="THAT IS MUCH BETTER. NOW YOUR FLASK IS EMPTY. YOU ARE IN THE MIDDLE OF THE OCEAN.">

**Type `NORTH`**

<img src="images/wizard/0086.png" width="400" alt="YOU ARE IN THE MIDDLE OF THE OCEAN. --------------- ENTER COMMAND?NORTH YOU ARE IN THE MIDDLE OF THE OCEAN">

**Type `EAST`**

<img src="images/wizard/0087.png" width="400" alt="YOU ARE IN THE MIDDLE OF THE OCEAN --------------- ENTER COMMAND?EAST YOU ARE IN THE MIDDLE OF THE OCEAN">

**Type `EAST`**

<img src="images/wizard/0088.png" width="400" alt="YOU ARE IN THE MIDDLE OF THE OCEAN --------------- ENTER COMMAND?EAST YOU ARE IN THE MIDDLE OF THE OCEAN">

**Type `EAST`**

<img src="images/wizard/0089.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN A ROWBOAT ON THE BEACH OF AN ISLAND.">

**Type `LEAVE BOAT`**

### 5. The pirate's chest

<img src="images/wizard/0090.png" width="400" alt="--------------- ENTER COMMAND?LEAVE BOAT YOU ARE ON A BEACH ON AN ISLAND.">

**Type `EAST`**

<img src="images/wizard/0091.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN THE JUNGLE OF THE ISLAND. A PATH GOES NORTH AND A PATH GOES WEST.">

**Type `NORTH`**

<img src="images/wizard/0092.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN THE JUNGLE OF THE ISLAND. A PATH GOES NORTH, SOUTH AND WEST.">

**Type `NORTH`**

<img src="images/wizard/0093.png" width="400" alt="--------------- ENTER COMMAND?NORTH THERE IS AN ANCHOR HERE. YOU ARE ON A BEACH ON THE ISLAND.">

**Type `GET ANCHOR`**

<img src="images/wizard/0094.png" width="400" alt="--------------- ENTER COMMAND?GET ANCHOR YOU ARE ON A BEACH ON THE ISLAND.">

**Type `TIE ROPE`**

<img src="images/wizard/0095.png" width="400" alt="--------------- ENTER COMMAND?TIE ROPE TO WHAT? YOU ARE ON A BEACH ON THE ISLAND.">

**Type `TO ANCHOR`**

<img src="images/wizard/0096.png" width="400" alt="--------------- ENTER COMMAND?TO ANCHOR THE ANCHOR IS TIED TO THE ROPE. YOU ARE ON A BEACH ON THE ISLAND.">

**Type `WEST`**

<img src="images/wizard/0097.png" width="400" alt="YOU ARE IN THE JUNGLE OF THE ISLAND. THERE IS A TREE HOUSE UP IN A LARGE TREE. A PATH GOES EAST AND SOUTH.">

**Type `THROW ROPE`**

<img src="images/wizard/0098-1.png" width="400" alt="THE ROPE IS THROWN OVER THE BRANCH. YOU ARE IN THE JUNGLE OF THE ISLAND. THERE IS A TREE HOUSE UP IN A LARGE TREE. A PATH GOES EAST AND SOUTH."><br>
<img src="images/wizard/0098.png" width="400" alt="YOU ARE IN THE JUNGLE OF THE ISLAND. THERE IS A TREE HOUSE UP IN A LARGE TREE. A PATH GOES EAST AND SOUTH.">

**Type `UP`**

<img src="images/wizard/0099.png" width="400" alt="TREE. A PATH GOES EAST AND SOUTH. --------------- ENTER COMMAND?UP YOU ARE IN THE TREE HOUSE.">

**Type `GET SHOVEL`**

<img src="images/wizard/0100.png" width="400" alt="--------------- ENTER COMMAND?GET SHOVEL YOU ARE IN THE TREE HOUSE.">

**Type `DOWN`**

<img src="images/wizard/0101.png" width="400" alt="YOU ARE IN THE JUNGLE OF THE ISLAND. THERE IS A TREE HOUSE UP IN A LARGE TREE. A PATH GOES EAST AND SOUTH.">

**Type `SOUTH`**

<img src="images/wizard/0102.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN THE JUNGLE OF THE ISLAND. A PATH GOES NORTH AND SOUTH.">

**Type `SOUTH`**

<img src="images/wizard/0103.png" width="400" alt="PATH GOES NORTH AND SOUTH. --------------- ENTER COMMAND?SOUTH YOU ARE ON A BEACH ON AN ISLAND.">

**Type `DIG`**

<img src="images/wizard/0104.png" width="400" alt="--------------- ENTER COMMAND?DIG YOU HAVE UNCOVERED A TREASURE CHEST. YOU ARE ON A BEACH ON AN ISLAND.">

**Type `GET CHEST`**

<img src="images/wizard/0105-1.png" width="400" alt="A PIRATE JUMPS FROM BEHIND A TREE. &#34;SHIVER ME TIMBERS! TRYING TO STEAL MY TREASURE?&#34; HE GRABS THE CHEST AND RUNS. YOU ARE ON A BEACH ON AN ISLAND."><br>
<img src="images/wizard/0105.png" width="400" alt="&#34;SHIVER ME TIMBERS! TRYING TO STEAL MY TREASURE?&#34; HE GRABS THE CHEST AND RUNS. YOU ARE ON A BEACH ON AN ISLAND.">

**Type `EAST`**

<img src="images/wizard/0106.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN THE JUNGLE OF THE ISLAND. A PATH GOES NORTH AND A PATH GOES WEST.">

**Type `NORTH`**

<img src="images/wizard/0107.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN THE JUNGLE OF THE ISLAND. A PATH GOES NORTH, SOUTH AND WEST.">

**Type `WEST`**

<img src="images/wizard/0108.png" width="400" alt="--------------- ENTER COMMAND?WEST THERE IS A CAVE HERE IN THE MIDDLE OF THE ISLAND. A PATH GOES EAST.">

**Type `GO CAVE`**

<img src="images/wizard/0109.png" width="400" alt="--------------- ENTER COMMAND?GO CAVE YOU ARE IN THE CAVE. THE CAVE ENDS RIGHT HERE.">

**Type `OPEN CHEST`**

<img src="images/wizard/0110.png" width="400" alt="YOU ARE IN THE CAVE. THE CAVE ENDS RIGHT HERE.">

**Type `LOOK CHEST`**

<img src="images/wizard/0111.png" width="400" alt="THERE IS A SMALL HARP HERE. YOU ARE IN THE CAVE. THE CAVE ENDS RIGHT HERE.">

**Type `GET HARP`**

<img src="images/wizard/0112.png" width="400" alt="--------------- ENTER COMMAND?GET HARP YOU ARE IN THE CAVE. THE CAVE ENDS RIGHT HERE.">

**Type `LEAVE CAVE`**

### 6. The ring, the coin and the giant

<img src="images/wizard/0113.png" width="400" alt="THERE IS A CAVE HERE IN THE MIDDLE OF THE ISLAND. A PATH GOES EAST.">

**Type `EAST`**

<img src="images/wizard/0114.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN THE JUNGLE OF THE ISLAND. A PATH GOES NORTH, SOUTH AND WEST.">

**Type `NORTH`**

<img src="images/wizard/0115.png" width="400" alt="PATH GOES NORTH, SOUTH AND WEST. --------------- ENTER COMMAND?NORTH YOU ARE ON A BEACH ON THE ISLAND.">

**Type `DRINK VIAL`**

<img src="images/wizard/0116.png" width="400" alt="WHEN YOU DRINK THE LIQUID YOUR ARMS TURN INTO WINGS! YOU ARE ON A BEACH ON THE ISLAND.">

**Type `NORTH`**

<img src="images/wizard/0117.png" width="400" alt="YOU ARE ON A BEACH ON THE ISLAND. --------------- ENTER COMMAND?NORTH YOU ARE ON A BEACH.">

**Type `NORTH`**

<img src="images/wizard/0118-1.png" width="400" alt="YOU HAVE CHANGED BACK INTO YOURSELF AGAIN! THERE IS A BEAUTIFUL SAPPHIRE RING HERE."><br>
<img src="images/wizard/0118.png" width="400" alt="THERE IS A BEAUTIFUL SAPPHIRE RING HERE. YOU ARE IN THE FOOTHILLS.">

**Type `GET RING`**

<img src="images/wizard/0119.png" width="400" alt="YOU ARE IN THE FOOTHILLS. --------------- ENTER COMMAND?GET RING YOU ARE IN THE FOOTHILLS.">

**Type `NORTH`**

<img src="images/wizard/0120.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN THE MOUTAINS. THERE IS AN OLD PEASANT WOMAN HERE.">

**Type `TALK WOMAN`**

<img src="images/wizard/0121-1.png" width="400" alt="SHE WARNS YOU OF THE GIANT WHO LIVES IN THE MOUNTAINS YOU ARE IN THE MOUTAINS. THERE IS AN OLD PEASANT WOMAN HERE."><br>
<img src="images/wizard/0121.png" width="400" alt="THE MOUNTAINS YOU ARE IN THE MOUTAINS. THERE IS AN OLD PEASANT WOMAN HERE.">

**Type `WEST`**

<img src="images/wizard/0122.png" width="400" alt="--------------- ENTER COMMAND?WEST THERE IS A SLIGHT DRIZZLE. YOU ARE IN THE MOUNTAINS.">

**Type `GO RAINBOW`**

<img src="images/wizard/0123.png" width="400" alt="THERES A GOLD COIN HERE YOU ARE IN THE MOUNTAINS.">

**Type `GET COIN`**

<img src="images/wizard/0124.png" width="400" alt="--------------- ENTER COMMAND?GET COIN THE SUN CAME OUT. YOU ARE IN THE MOUNTAINS.">

**Type `NORTH`**

<img src="images/wizard/0125-1.png" width="400" alt="YOU ARE AT THE EAST EDGE OF A DEEP GORGE. THERE IS A RICKETY BRIDGE CROSSING THE GORGE. I DON&#39;T THINK IT CAN HOLD MUCH WEIGHT."><br>
<img src="images/wizard/0125.png" width="400" alt="GORGE. THERE IS A RICKETY BRIDGE CROSSING THE GORGE. I DON&#39;T THINK IT CAN HOLD MUCH WEIGHT.">

**Type `LUCY`**

<img src="images/wizard/0126-1.png" width="400" alt="EVERYTHING YOU ARE CARRYING DISAPPEARS. YOU HAVE NOTHING LEFT. YOU ARE AT THE EAST EDGE OF A DEEP GORGE. THERE IS A RICKETY BRIDGE"><br>
<img src="images/wizard/0126.png" width="400" alt="GORGE. THERE IS A RICKETY BRIDGE CROSSING THE GORGE. I DON&#39;T THINK IT CAN HOLD MUCH WEIGHT.">

**Type `WEST`**

<img src="images/wizard/0127.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE AT THE WEST END OF A DEEP GORGE.">

**Type `WEST`**

<img src="images/wizard/0128.png" width="400" alt="GORGE. --------------- ENTER COMMAND?WEST YOU ARE IN THE MOUNTAINS.">

**Type `NORTH`**

<img src="images/wizard/0129.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN THE MOUNTAINS. THERE IS A CAVE HERE.">

**Type `GO CAVE`**

<img src="images/wizard/0130.png" width="400" alt="CAVE HERE. --------------- ENTER COMMAND?GO CAVE YOU ARE IN A CAVE. IT ENDS RIGHT HERE.">

**Type `GET ALL`**

<img src="images/wizard/0131.png" width="400" alt="YOU ARE IN A CAVE. IT ENDS RIGHT HERE. --------------- ENTER COMMAND?GET ALL YOU ARE IN A CAVE. IT ENDS RIGHT HERE.">

**Type `LEAVE CAVE`**

<img src="images/wizard/0132.png" width="400" alt="YOU ARE IN THE MOUNTAINS. THERE IS A CAVE HERE.">

**Type `SOUTH`**

<img src="images/wizard/0133.png" width="400" alt="CAVE HERE. --------------- ENTER COMMAND?SOUTH YOU ARE IN THE MOUNTAINS.">

**Type `WEST`**

<img src="images/wizard/0134.png" width="400" alt="YOU ARE IN THE MOUNTAINS. --------------- ENTER COMMAND?WEST YOU ARE IN THE MOUNTAINS">

**Type `PLAY HARP`**

### 7. The horn and the maze

<img src="images/wizard/0135-1.png" width="400" alt="THE GIANT IS A GREAT LOVER OF MUSIC. HE THANKS YOU FOR THE HARP AND LEAVES WITH IT. YOU ARE IN THE MOUNTAINS"><br>
<img src="images/wizard/0135.png" width="400" alt="THANKS YOU FOR THE HARP AND LEAVES WITH IT. YOU ARE IN THE MOUNTAINS">

**Type `NORTH`**

<img src="images/wizard/0136.png" width="400" alt="YOU ARE IN THE MOUNTAINS --------------- ENTER COMMAND?NORTH YOU ARE IN THE MOUNTAINS">

**Type `NORTH`**

<img src="images/wizard/0137-1.png" width="400" alt="THERE IS A PEDDLER SELLING WARES FOR ONE GOLD COIN EACH. YOU ARE IN THE FOOTHILLS ON THE NORTH SIDE OF THE MOUNTAINS."><br>
<img src="images/wizard/0137.png" width="400" alt="ONE GOLD COIN EACH. YOU ARE IN THE FOOTHILLS ON THE NORTH SIDE OF THE MOUNTAINS.">

**Type `LOOK TABLE`**

<img src="images/wizard/0138-1.png" width="400" alt="THERE ARE BOOTS, A JUG OF WINE, A DAGGER, A HORN AND A FRYING PAN ON THE TABLE. YOU ARE IN THE FOOTHILLS ON THE NORTH"><br>
<img src="images/wizard/0138.png" width="400" alt="TABLE. YOU ARE IN THE FOOTHILLS ON THE NORTH SIDE OF THE MOUNTAINS.">

**Type `BUY HORN`**

<img src="images/wizard/0139.png" width="400" alt="--------------- ENTER COMMAND?BUY HORN YOU ARE IN THE FOOTHILLS ON THE NORTH SIDE OF THE MOUNTAINS.">

**Type `NORTH`**

<img src="images/wizard/0140.png" width="400" alt="YOU ARE IN THE FOOTHILLS ON THE NORTH SIDE OF THE MOUNTAINS. YOU SEE A CASTLE IN THE DISTANCE.">

**Type `NORTH`**

<img src="images/wizard/0141.png" width="400" alt="YOU ARE IN FRONT OF THE CASTLE. THERE IS A MOAT AROUND THE CASTLE FULL OF CROCODILES.">

**Type `PLAY HORN`**

<img src="images/wizard/0142.png" width="400" alt="YOU ARE IN FRONT OF THE CASTLE. THERE IS A MOAT AROUND THE CASTLE FULL OF CROCODILES.">

**Type `NORTH`**

<img src="images/wizard/0143.png" width="400" alt="YOU ARE IN THE ENTRY HALL OF THE CASTLE. THERE ARE DOORWAYS TO THE NORTH, WEST, AND SOUTH (BEHIND YOU).">

**Type `NORTH`**

<img src="images/wizard/0144.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN A MAZE OF PASSAGEWAYS.PASSAGES GO SOUTH AND WEST.">

**Type `WEST`**

<img src="images/wizard/0145.png" width="400" alt="YOU ARE IN A MAZE OF PASSAGEWAYS.PASSAGES GO NORTH, EAST AND WEST.">

**Type `WEST`**

<img src="images/wizard/0146.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN A MAZE OF PASSAGEWAYS. PASSAGES GO EAST AND WEST.">

**Type `WEST`**

<img src="images/wizard/0147.png" width="400" alt="YOU ARE IN A MAZE OF PASSAGEWAYS.PASSAGES GO NORTH, EAST AND WEST.">

**Type `NORTH`**

<img src="images/wizard/0148.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN A MAZE OF PASSAGEWAYS. PASSAGES GO EAST AND SOUTH.">

**Type `EAST`**

<img src="images/wizard/0149.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN A MAZE OF PASSAGEWAYS. PASSAGES GO EAST, WEST AND SOUTH.">

**Type `EAST`**

<img src="images/wizard/0150.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN A MAZE OF PASSAGEWAYS. PASSAGES GO NORTH AND WEST.">

**Type `NORTH`**

<img src="images/wizard/0151.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN A MAZE OF PASSAGEWAYS.PASSAGES GO N, S, E AND W.">

**Type `NORTH`**

<img src="images/wizard/0152.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN A MAZE OF PASSAGEWAYS. PASSAGES GO WEST AND SOUTH.">

**Type `WEST`**

<img src="images/wizard/0153.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN A MAZE OF PASSAGEWAYS. PASSAGES GO E, W, N AND S.">

**Type `NORTH`**

<img src="images/wizard/0154.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN A MAZE OF PASSAGEWAYS. PASSAGES GO NORTH, SOUTH AND EAST.">

**Type `EAST`**

<img src="images/wizard/0155.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN A MAZE OF PASSAGEWAYS. PASSAGES GO EAST AND WEST.">

**Type `EAST`**

<img src="images/wizard/0156.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN A MAZE OF PASSAGEWAYS. PASSAGES GO EAST AND WEST.">

**Type `EAST`**

### 8. The castle of Harlin

<img src="images/wizard/0157.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN A SMALL EMPTY ROOM. THERE IS A PASSAGE BEHIND YOU.">

**Type `USE KNIFE`**

<img src="images/wizard/0158.png" width="400" alt="YOU PICK THE LOCK. YOU ARE IN A SMALL EMPTY ROOM. THERE IS A PASSAGE BEHIND YOU.">

**Type `OPEN DOOR`**

<img src="images/wizard/0159.png" width="400" alt="--------------- ENTER COMMAND?OPEN DOOR YOU ARE IN A SMALL EMPTY ROOM. THERE IS A PASSAGE BEHIND YOU.">

**Type `EAST`**

<img src="images/wizard/0160.png" width="400" alt="A PASSAGE BEHIND YOU. --------------- ENTER COMMAND?EAST YOU ARE IN A SHORT HALLWAY.">

**Type `UP`**

<img src="images/wizard/0161.png" width="400" alt="--------------- ENTER COMMAND?UP YOU ARE IN A TOWER. THERE ARE STAIRS LEADING DOWN.">

**Type `DOWN`**

<img src="images/wizard/0162.png" width="400" alt="LEADING DOWN. --------------- ENTER COMMAND?DOWN YOU ARE IN A SHORT HALLWAY.">

**Type `UP`**

<img src="images/wizard/0163.png" width="400" alt="--------------- ENTER COMMAND?UP YOU ARE IN A TOWER. THERE ARE STAIRS LEADING DOWN.">

**Type `WEAR RING`**

<img src="images/wizard/0164-1.png" width="400" alt="THE RING IS ON YOUR FINGER. THERE IS A BIRD HERE YOU ARE IN A TOWER. THERE ARE STAIRS LEADING DOWN."><br>
<img src="images/wizard/0164.png" width="400" alt="THERE IS A BIRD HERE YOU ARE IN A TOWER. THERE ARE STAIRS LEADING DOWN.">

**Type `RUB RING`**

<img src="images/wizard/0165-1.png" width="400" alt="YOU TURN INTO A CAT, LEAP UP AND EAT THE BIRD. YOU ARE YOURSELF AGAIN. YOU ARE IN A TOWER. THERE ARE STAIRS"><br>
<img src="images/wizard/0165.png" width="400" alt="YOU ARE YOURSELF AGAIN. YOU ARE IN A TOWER. THERE ARE STAIRS LEADING DOWN.">

**Type `DOWN`**

<img src="images/wizard/0166.png" width="400" alt="LEADING DOWN. --------------- ENTER COMMAND?DOWN YOU ARE IN A SHORT HALLWAY.">

**Type `EAST`**

<img src="images/wizard/0167.png" width="400" alt="YOU ARE IN A SHORT HALLWAY. --------------- ENTER COMMAND?EAST YOU ARE IN A TINY ROOM.">

**Type `KISS FROG`**

<img src="images/wizard/0168.png" width="400" alt="THE FROG BECOMES A PRINCESS! THERE IS A PRINCESS FOLLOWING YOU. YOU ARE IN A TINY ROOM.">

**Type `EAST`**

<img src="images/wizard/0169.png" width="400" alt="--------------- ENTER COMMAND?EAST THERE IS A PRINCESS FOLLOWING YOU. YOU ARE IN AN EAST/WEST HALLWAY.">

**Type `EAST`**

<img src="images/wizard/0170.png" width="400" alt="--------------- ENTER COMMAND?EAST THERE IS A PRINCESS FOLLOWING YOU. YOU ARE IN A ROOM WITH A CLOSET.">

**Type `OPEN CLOSET`**

<img src="images/wizard/0171.png" width="400" alt="T THERE IS A PRINCESS FOLLOWING YOU. YOU ARE IN A ROOM WITH A CLOSET.">

**Type `LOOK CLOSET`**

<img src="images/wizard/0172.png" width="400" alt="THERE IS A PAIR OF SHOES HERE THERE IS A PRINCESS FOLLOWING YOU. YOU ARE IN A ROOM WITH A CLOSET.">

**Type `GET SHOES`**

<img src="images/wizard/0173.png" width="400" alt="--------------- ENTER COMMAND?GET SHOES THERE IS A PRINCESS FOLLOWING YOU. YOU ARE IN A ROOM WITH A CLOSET.">

**Type `LOOK SHOES`**

<img src="images/wizard/0174.png" width="400" alt="THERE IS A PRINCESS FOLLOWING YOU. YOU ARE IN A ROOM WITH A CLOSET.">

**Type `WEAR SHOES`**

<img src="images/wizard/0175.png" width="400" alt="OK. THERE IS A PRINCESS FOLLOWING YOU. YOU ARE IN A ROOM WITH A CLOSET.">

**Type `WHOOSH`**

<img src="images/wizard/0176-1.png" width="400" alt="YOU ARE TRANSPORTED TO SERENIA. THERE IS A PRINCESS FOLLOWING YOU. YOU ARE IN THE VILLAGE OF SERENIA. AROUND YOU IS A DESERT."><br>
<img src="images/wizard/0176.png" width="400" alt="THERE IS A PRINCESS FOLLOWING YOU. YOU ARE IN THE VILLAGE OF SERENIA. AROUND YOU IS A DESERT.">

**Type `LOOK`**

### The end

<img src="images/wizard/0177-1.png" width="400" alt="CONGRATULATIONS!!! YOU HAVE SAFELY RETURNED THE PRINCESS TO SERENIA. FOR THIS OUTSTANDING FEAT YOU HAVE BEEN DECLARED A JUNIOR-MASTER ADVENTURER."><br>
<img src="images/wizard/0177-2.png" width="400" alt="THANK YOU FOR PLAYING HI-RES ADVENTURE #2 &#39;THE WIZARD AND THE PRINCESS&#39; ... KEN AND ROBERTA WILLIAMS, ON-LINE SYSTEMS"><br>
<img src="images/wizard/0177.png" width="400" alt="SYSTEMS YOU ARE IN THE VILLAGE OF SERENIA. AROUND YOU IS A DESERT.">

<!-- End of the walkthrough -->

## What next

[Mystery House](mysteryhouse.md) is the first of the Hi-Res Adventures, and
[Time Zone](timezone.md) the fifth, the biggest.
