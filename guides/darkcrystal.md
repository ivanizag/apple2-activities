# The Dark Crystal, from the start to the end

[Back to the activities](../README.md)

**The Dark Crystal**, of Sierra On-Line, came out in 1982 for the Apple \]\[,
by Roberta Williams, programmed by Ken Williams and Chris Iden, after the
film of Jim Henson of the same year. Jen, a Gelfling, has to find the shard of
the Dark Crystal and heal it before the three suns meet, with Kira, Aughra
and the dying Mystic Ursu, against the Skeksis and their Garthim. The game
came on two diskettes, both sides used.

You type one or two words, `SPEAK URSU`, `PLAY FLUTE`, `RIDE LANDSTRIDER`, and
the game answers under the picture, after its own prompt, `----> ENTER
COMMAND?`. It asks for each side of its diskettes when it needs it.

This page plays it to its end, all 112 commands, with a picture of the screen
after each one: 218 pictures. The pictures are of a colour television, with
the four lines of text under them drawn white and sharp, as a monochrome
monitor showed them: the television blurred them into fringes of colour,
hard to read.

## What you need

**The Dark Crystal (4am and san inc crack)**, from
[its page on the Internet Archive](https://archive.org/details/TheDarkCrystal4amCrack):
the zip `The Dark Crystal (4am and san inc crack).zip`, with a file for each
side, from `The Dark Crystal (4am and san inc crack) disk 1A.dsk` to
`The Dark Crystal (4am and san inc crack) disk 2B.dsk`. `./fetch-disks.sh`
in this repository downloads them into `disks/` and checks them.

The commands of this page are also in this repository, a line each,
[listings/darkcrystal.txt](listings/darkcrystal.txt). They follow the
walkthrough of
[The Sierra Help Pages](https://sierrahelp.com/Walkthroughs/DarkCrystalWalkthrough.html),
with what the machine showed it needs where it is not exact: four tries to
move in the vines before the eyeball of Aughra shows, three `LOOK`s in the Pod
Village until the Garthim come, `RIDE LANDSTRIDER`, and the closet where Jen
hides from the Skeksis before the curtain of the dining room.

The [manual of The Dark Crystal](https://archive.org/details/vgmuseum_sierra_darkcrystal)
is on the Internet Archive.

## The machine

An Apple \]\[+:

- the 6502 processor at 1 MHz and 48 KB of memory;
- a Disk II controller card in slot 6, with disk 1, side A, in drive 1.

```bash
izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s6 'diskii,disk1=disks/The Dark Crystal (4am and san inc crack) disk 1A.dsk'
```

Shorter, the model `2plus` of izapple2 plays it too: the same Apple \]\[+,
with the 16 KB of a language card in slot 0 and a Videx 80 column card in
slot 3 more, on a colour monitor with its scan lines.

```bash
izapple2 -model 2plus 'disks/The Dark Crystal (4am and san inc crack) disk 1A.dsk'
```

## Playing it

1. **Start izapple2** with the command above. The game asks for disk 1,
   side B: drop its file on the area of drive 1 of the window of izapple2,
   and press Return. The page says each time it asks for another.

2. **Type the commands**, each followed by Return. Each step below is what
   the game shows, and the command to type then. When the picture stays the
   same and only the text goes on, the page shows the picture once, and
   under it the lines of text that came after. When the game stops in the
   middle of a longer text, **press Return** to go on.

## The walkthrough

<!-- The walkthrough, written by the generator -->

- [1. The Mystics and Ursu](#1-the-mystics-and-ursu)
- [2. The flute and the lily pad](#2-the-flute-and-the-lily-pad)
- [3. Aughra](#3-aughra)
- [4. Kira and the Pod Village](#4-kira-and-the-pod-village)
- [5. The ruins and the Landstriders](#5-the-ruins-and-the-landstriders)
- [6. Into the castle of the Skeksis](#6-into-the-castle-of-the-skeksis)
- [7. The scepter and the tower](#7-the-scepter-and-the-tower)
- [8. The Great Conjunction](#8-the-great-conjunction)
- [The end](#the-end)

### 1. The Mystics and Ursu

<img src="images/darkcrystal/0001-1.png" width="400" alt="PLEASE INSERT DISK #1, SIDE &#34;B&#34; AND PRESS RETURN."><br>
<img src="images/darkcrystal/0001.png" width="400" alt="JEN IS IN A BEAUTIFUL MOUNTAIN VALLEY. THE MYSTICS HAVE A SPECIAL NAME FOR IT, &#34;THE VALLEY OF THE STONES.&#34;">

*Disk 1, side B: drop `The Dark Crystal (4am and san inc crack) disk 1B.dsk` on drive 1, and press Return.*

**Type `EAST`**

<img src="images/darkcrystal/0002-1.png" width="400" alt="BEFORE JEN CAN ACT, A MYSTIC APPROACHES AND SAYS, &#34;URSU, WISEST OF OUR RACE, IS DYING. HE HAS SENT FOR YOU. COME QUICKLY!&#34; THEN THE MYSTIC WALKS AWAY."><br>
<img src="images/darkcrystal/0002.png" width="400" alt="DYING. HE HAS SENT FOR YOU. COME QUICKLY!&#34; THEN THE MYSTIC WALKS AWAY. JEN IS IN THE VALLEY OF THE STONES.">

**Type `EAST`**

<img src="images/darkcrystal/0003.png" width="400" alt="JEN IS STANDING ON A MOUNTAINSIDE COVERED WITH LOOSE, AND EXTREMELY SHARP SHALE.">

**Type `GET SHALE`**

<img src="images/darkcrystal/0004.png" width="400" alt="----&gt; ENTER COMMAND?GET SHALE OKAY. JEN IS IN THE MOUNTAINS.">

**Type `WEST`**

<img src="images/darkcrystal/0005.png" width="400" alt="JEN IS IN THE MOUNTAINS. ----&gt; ENTER COMMAND?WEST JEN IS IN THE VALLEY OF THE STONES.">

**Type `WEST`**

<img src="images/darkcrystal/0006.png" width="400" alt="JEN IS IN THE VALLEY OF THE STONES, SO NAMED FOR THE CIRCULAR FORMATIONS OF STANDING STONES THAT LIE WITHIN IT.">

**Type `WEST`**

<img src="images/darkcrystal/0007-1.png" width="400" alt="JEN IS IN THE VALLEY OF THE STONES. TOWERING ABOVE HIM TO THE WEST IS A STEEP CLIFF. A CAVE OPENING BECKONS FROM THE SIDE OF THE CLIFF."><br>
<img src="images/darkcrystal/0007.png" width="400" alt="TOWERING ABOVE HIM TO THE WEST IS A STEEP CLIFF. A CAVE OPENING BECKONS FROM THE SIDE OF THE CLIFF.">

**Type `WEST`**

<img src="images/darkcrystal/0008.png" width="400" alt="JEN IS SLOWLY WALKING WITHIN A DIMLY LIT CAVE. THE PASSAGE WINDS TO THE NORTH AND TO THE EAST.">

**Type `NORTH`**

<img src="images/darkcrystal/0009-1.png" width="400" alt="JEN IS IN A LARGE, CANDLE-LIT CAVERN. URSU, WISEST OF THE MYSTICS, LIES HERE DYING. BESIDE HIM ON THE FLOOR IS A BOWL OF LIQUID."><br>
<img src="images/darkcrystal/0009.png" width="400" alt="URSU, WISEST OF THE MYSTICS, LIES HERE DYING. BESIDE HIM ON THE FLOOR IS A BOWL OF LIQUID.">

**Type `SPEAK URSU`**

<img src="images/darkcrystal/0010-1.png" width="400" alt="A GREAT CONJUNCTION. YOU MUST FIND AUGHRA, KEEPER OF SECRETS AND WATCHER OF THE HEAVENS. SHE MAY HAVE THE SHARD YOU SEEK."><br>
<img src="images/darkcrystal/0010-2.png" width="400" alt="LIFE...BUT NOT AGAIN IN THIS ONE....&#34; WITH THESE WORDS, URSU DIES, AND HIS LIFELESS BODY VANISHES FROM THE SLEEPFRAME."><br>
<img src="images/darkcrystal/0010.png" width="400" alt="JEN IS IN A LARGE CAVERN.">

**Type `LOOK BOWL`**

### 2. The flute and the lily pad

<img src="images/darkcrystal/0011.png" width="400" alt="LOOKING INTO THE BOWL OF POTION, JEN SEES THE IMAGE OF A CRYSTAL SHARD. JEN IS IN A LARGE CAVERN.">

**Type `SOUTH`**

<img src="images/darkcrystal/0012.png" width="400" alt="JEN IS IN A LARGE CAVERN. ----&gt; ENTER COMMAND?SOUTH JEN IS IN A CAVE PASSAGE.">

**Type `EAST`**

<img src="images/darkcrystal/0013.png" width="400" alt="JEN IS IN A CAVE PASSAGE. ----&gt; ENTER COMMAND?EAST JEN IS IN THE VALLEY OF THE STONES.">

**Type `EAST`**

<img src="images/darkcrystal/0014.png" width="400" alt="JEN IS IN THE VALLEY OF THE STONES. ----&gt; ENTER COMMAND?EAST JEN IS IN THE VALLEY OF THE STONES.">

**Type `NORTH`**

<img src="images/darkcrystal/0015.png" width="400" alt="JEN IS MAKING HIS WAY ALONG A SHADOWY PATH THAT SNAKES THROUGH THE HILLS ABOVE THE VALLEY OF THE STONES.">

**Type `DIG`**

<img src="images/darkcrystal/0016-1.png" width="400" alt="USING THE SHALE, JEN DIGS IN THE GROUND AT THE BASE OF THE TREE AND UNCOVERS A STRANGE LOOKING FLUTE. JEN IS IN THE HILLS."><br>
<img src="images/darkcrystal/0016.png" width="400" alt="AT THE BASE OF THE TREE AND UNCOVERS A STRANGE LOOKING FLUTE. JEN IS IN THE HILLS.">

**Type `GET FLUTE`**

<img src="images/darkcrystal/0017.png" width="400" alt="JEN IS IN THE HILLS. ----&gt; ENTER COMMAND?GET FLUTE JEN IS IN THE HILLS.">

**Type `NORTH`**

<img src="images/darkcrystal/0018-1.png" width="400" alt="JEN FALLS HEAD OVER HEELS DOWN A STEEP SLOPE. JEN IS TRAVERSING A WILDERNESS OF TANGLED VINES, CHATTERING BLOSSOMS, AND"><br>
<img src="images/darkcrystal/0018.png" width="400" alt="JEN IS TRAVERSING A WILDERNESS OF TANGLED VINES, CHATTERING BLOSSOMS, AND WARY CREATURES.">

**Type `EAST`**

<img src="images/darkcrystal/0019-1.png" width="400" alt="JEN IS MAKING HIS WAY THROUGH A DENSE WILDERNESS WITH CHATTERING FLOWERS, TANGLED VINES, AND CREATURES PEEKING FROM EVERYWHERE."><br>
<img src="images/darkcrystal/0019.png" width="400" alt="WILDERNESS WITH CHATTERING FLOWERS, TANGLED VINES, AND CREATURES PEEKING FROM EVERYWHERE.">

**Type `NORTH`**

<img src="images/darkcrystal/0020-1.png" width="400" alt="THERE IS A BABBLING BROOK SPLASHING THROUGH THE WILDERNESS HERE. CHATTERING FLOWERS AND TALL GRASSES LINE ITS BANKS."><br>
<img src="images/darkcrystal/0020.png" width="400" alt="THROUGH THE WILDERNESS HERE. CHATTERING FLOWERS AND TALL GRASSES LINE ITS BANKS.">

**Type `LISTEN BROOK`**

<img src="images/darkcrystal/0021-1.png" width="400" alt="JEN LISTENS CAREFULLY TO THE BABBLING BROOK. THE BROOK, WHICH SEEMS TO HAVE A SLIGHT STUTTER, SAYS, &#34;E-E-EN, N-N-NEN.&#34;"><br>
<img src="images/darkcrystal/0021.png" width="400" alt="SLIGHT STUTTER, SAYS, &#34;E-E-EN, N-N-NEN.&#34; JEN IS IN THE WILDERNESS.">

**Type `WEST`**

<img src="images/darkcrystal/0022-1.png" width="400" alt="JEN IS ALONE IN THE WILDERNESS. HAPPILY, THERE IS A BEAUTIFUL POND SPARKLING LIKE A GEM AMONG THE CHATTERING FLOWERS TO BRIGHTEN UP JEN&#39;S"><br>
<img src="images/darkcrystal/0022.png" width="400" alt="LONELINESS. CROAKING, FROG-LIKE CREATURES ABOUND ON HUGE LILY PADS FLOATING ON THE POND.">

**Type `CUT PAD`**

<img src="images/darkcrystal/0023-1.png" width="400" alt="USING THE SHARP SHALE, JEN CUTS THE LILY PAD AWAY FROM ITS THICK STEM AND TAKES IT WITH HIM. JEN IS IN THE WILDERNESS."><br>
<img src="images/darkcrystal/0023.png" width="400" alt="LILY PAD AWAY FROM ITS THICK STEM AND TAKES IT WITH HIM. JEN IS IN THE WILDERNESS.">

**Type `EAST`**

<img src="images/darkcrystal/0024.png" width="400" alt="JEN IS IN THE WILDERNESS. ----&gt; ENTER COMMAND?EAST JEN IS IN THE WILDERNESS.">

**Type `EAST`**

<img src="images/darkcrystal/0025-1.png" width="400" alt="CREATURES SCURRY OUT OF JEN&#39;S WAY AS HE WALKS THROUGH A WILDERNESS OF THICK FOLIAGE AND CHATTERING FLOWERS. TO HIS HORROR, JEN SEES A CRYSTAL BAT"><br>
<img src="images/darkcrystal/0025.png" width="400" alt="TO HIS HORROR, JEN SEES A CRYSTAL BAT HOVERING OVERHEAD. ITS CRYSTAL &#34;EYE&#34; IS STARING DIRECTLY AT HIM.">

**Type `EAST`**

<img src="images/darkcrystal/0026-1.png" width="400" alt="JEN IS WALKING ON AN EAST-WEST PATH THROUGH WHAT SEEMS LIKE AN ENDLESS WILDERNESS. THE SOUNDS OF THE WILDERNESS HAVE GIVEN WAY TO AN EERIE"><br>
<img src="images/darkcrystal/0026-2.png" width="400" alt="HUSH. TO HIS HORROR, JEN SEES A CRYSTAL BAT HOVERING OVERHEAD. ITS CRYSTAL &#34;EYE&#34; IS STARING DIRECTLY AT HIM."><br>
<img src="images/darkcrystal/0026.png" width="400" alt="TO HIS HORROR, JEN SEES A CRYSTAL BAT HOVERING OVERHEAD. ITS CRYSTAL &#34;EYE&#34; IS STARING DIRECTLY AT HIM.">

**Type `EAST`**

<img src="images/darkcrystal/0027-1.png" width="400" alt="JEN IS ON A PATH IN THE WILDERNESS. THE PLANTS AND ANIMALS HERE SEEM STRANGELY QUIET. THE PATH HEADS WEST AND NORTH. TO HIS HORROR, JEN SEES A CRYSTAL BAT"><br>
<img src="images/darkcrystal/0027.png" width="400" alt="TO HIS HORROR, JEN SEES A CRYSTAL BAT HOVERING OVERHEAD. ITS CRYSTAL &#34;EYE&#34; IS STARING DIRECTLY AT HIM.">

**Type `NORTH`**

<img src="images/darkcrystal/0028.png" width="400" alt="JEN HAS ARRIVED IN A SWAMPLAND WHERE THE GROUND IS MUCKY AND VINES HANG EVERYWHERE FROM ENORMOUS TREES.">

**Type `NORTH`**

<img src="images/darkcrystal/0029-1.png" width="400" alt="JEN IS AT THE SOUTHERN EDGE OF A VAST SWAMPLAND THAT EXTENDS FOR MILES TO THE EAST AND WEST. LOOKING FAR TO THE NORTH, HE CAN BARELY MAKE OUT WHAT"><br>
<img src="images/darkcrystal/0029.png" width="400" alt="EAST AND WEST. LOOKING FAR TO THE NORTH, HE CAN BARELY MAKE OUT WHAT MIGHT BE THE SWAMP&#39;S BOUNDARY.">

**Type `USE PAD`**

### 3. Aughra

<img src="images/darkcrystal/0030-1.png" width="400" alt="FLOATING ATOP THE PAD HE CUT FROM THE WATER LILY, JEN PADDLES NORTH UNTIL HE REACHES A SHALLOW PORTION OF THE SWAMP. AS HE GETS OFF THE PAD, HOWEVER, HE"><br>
<img src="images/darkcrystal/0030-2.png" width="400" alt="FORGETS TO GRAB HOLD OF IT, AND IT FLOATS AWAY, HOPELESSLY OUT OF REACH. JEN IS TRUDGING THROUGH THE NORTHERN EDGE OF A VAST SWAMPLAND. VINES"><br>
<img src="images/darkcrystal/0030-3.png" width="400" alt="DANGLING FROM THE HUGE TREES HAMPER HIS EVERY STEP. TO HIS HORROR, JEN SEES A CRYSTAL BAT HOVERING OVERHEAD. ITS CRYSTAL &#34;EYE&#34; IS"><br>
<img src="images/darkcrystal/0030.png" width="400" alt="TO HIS HORROR, JEN SEES A CRYSTAL BAT HOVERING OVERHEAD. ITS CRYSTAL &#34;EYE&#34; IS STARING DIRECTLY AT HIM.">

**Type `NORTH`**

<img src="images/darkcrystal/0031.png" width="400" alt="JEN IS WADING THROUGH A MURKY SWAMPLAND. THE TREES ARE DRAPED WITH VINES AND SLIME COVERS EVERYTHING.">

**Type `EAST`**

<img src="images/darkcrystal/0032-1.png" width="400" alt="JEN IS IN AN EERIE SWAMPLAND. VINES HAVE COME DOWN FROM THE TREES AND WRAPPED THEMSELVES AROUND HIM. THE VINES WILL NOT LET GO OF JEN AS HE"><br>
<img src="images/darkcrystal/0032.png" width="400" alt="WRAPPED THEMSELVES AROUND HIM. THE VINES WILL NOT LET GO OF JEN AS HE STRUGGLES TO FREE HIMSELF.">

**Type `NORTH`**

<img src="images/darkcrystal/0033.png" width="400" alt="JEN CAN&#39;T. THE VINES ARE BINDING HIM TOO TIGHTLY. JEN IS IN A SWAMPLAND.">

**Type `NORTH`**

<img src="images/darkcrystal/0034.png" width="400" alt="JEN CAN&#39;T. THE VINES ARE BINDING HIM TOO TIGHTLY. JEN IS IN A SWAMPLAND.">

**Type `NORTH`**

<img src="images/darkcrystal/0035.png" width="400" alt="JEN CAN&#39;T. THE VINES ARE BINDING HIM TOO TIGHTLY. JEN IS IN A SWAMPLAND.">

**Type `NORTH`**

<img src="images/darkcrystal/0036-1.png" width="400" alt="JEN CAN&#39;T. THE VINES ARE BINDING HIM TOO TIGHTLY. LOOKING OUT FROM HIS VINE-PRISON, JEN FINDS HIMSELF RETURNING THE GAZE OF A"><br>
<img src="images/darkcrystal/0036.png" width="400" alt="FINDS HIMSELF RETURNING THE GAZE OF A SINGLE EYEBALL, THRUST UP AMONG THE TENDRILS BY A WITHERED HAND.">

**Type `SPEAK BEING`**

<img src="images/darkcrystal/0037.png" width="400" alt="AUGHRA SAYS, &#34;GELFLING, YOU KNOW ANSWER TO RIDDLE? HUH?&#34; JEN IS IN A SWAMPLAND.">

**Type `YES`**

<img src="images/darkcrystal/0038.png" width="400" alt="&#34;TELL ME ANSWER THEN, GELFLING.&#34; AUGHRA SNAPS IMPATIENTLY. JEN IS IN A SWAMPLAND.">

**Type `MOON DAUGHTERS`**

<img src="images/darkcrystal/0039.png" width="400" alt="----&gt; ENTER COMMAND?MOON DAUGHTERS &#34;VERY GOOD!&#34; CACKLES AUGHRA. SHE ORDERS THE VINES TO LET GO OF JEN.">

**Type `EAST`**

<img src="images/darkcrystal/0040-1.png" width="400" alt="BEFORE JEN CAN EVEN CATCH HIS BREATH, AUGHRA GOES NORTH TO HER OBSERVATORY. JEN, SENSING THAT SOMETHING IMPORTANT IS ABOUT TO HAPPEN, FOLLOWS CLOSE"><br>
<img src="images/darkcrystal/0040-2.png" width="400" alt="BEHIND. JEN IS IN THE OBSERVATORY OF AUGHRA, WATCHER OF THE HEAVENS AND KEEPER OF SECRETS."><br>
<img src="images/darkcrystal/0040.png" width="400" alt="SECRETS. AUGHRA TURNS TO JEN AND SAYS, &#34;WHAT YOU WANT?&#34;">

**Type `CRYSTAL SHARD`**

<img src="images/darkcrystal/0041-1.png" width="400" alt="AUGHRA CACKLES, &#34;THAT ALL? WHY NOT SAY SO?&#34; SHE SETS FOUR SHARDS ON THE TABLE: A BLUE ONE, A GREEN ONE, A VIOLET ONE AND AN ORANGE ONE. &#34;I NEVER COULD"><br>
<img src="images/darkcrystal/0041-2.png" width="400" alt="FIGURE OUT WHICH SHARD BELONG TO CRYSTAL,&#34; AUGHRA SAYS. &#34;MAYBE YOU CAN. ONLY CAN TAKE ONE SHARD WITH YOU, THOUGH. AND ONCE YOU PICK, NOT ALLOWED"><br>
<img src="images/darkcrystal/0041.png" width="400" alt="TO CHANGE MIND. SO.... CHOOSE CAREFULLY!&#34; JEN IS IN AUGHRA&#39;S OBSERVATORY.">

**Type `PLAY FLUTE`**

<img src="images/darkcrystal/0042.png" width="400" alt="THE BLUE SHARD BEGINS TO GLOW AND SOUNDS THE SAME CHORD BACK. JEN IS IN AUGHRA&#39;S OBSERVATORY.">

**Type `GET BLUE`**

### 4. Kira and the Pod Village

<img src="images/darkcrystal/0043.png" width="400" alt="JEN IS IN AUGHRA&#39;S OBSERVATORY. ----&gt; ENTER COMMAND?GET BLUE JEN IS IN AUGHRA&#39;S OBSERVATORY.">

**Type `SOUTH`**

<img src="images/darkcrystal/0044-1.png" width="400" alt="JEN FREEZES WITH FEAR AS A VICIOUS HORDE OF GARTHIM, WARRIORS OF THE EVIL SKEKSIS, LAUNCH A SURPRISE ATTACK ON THE OBSERVATORY."><br>
<img src="images/darkcrystal/0044.png" width="400" alt="HORDE OF GARTHIM, WARRIORS OF THE EVIL SKEKSIS, LAUNCH A SURPRISE ATTACK ON THE OBSERVATORY.">

**Type `GO WINDOW`**

<img src="images/darkcrystal/0045-1.png" width="400" alt="JEN JUMPS TO SAFETY, AND NONE TOO SOON! GLANCING OVER HIS SHOULDER, HE SEES THAT THE GARTHIM HAVE CAPTURED AUGHRA AND SET HER OBSERVATORY ON FIRE."><br>
<img src="images/darkcrystal/0045-2.png" width="400" alt="JEN FINDS HIMSELF MIRED UP TO HIS KNEES IN AN EERIE BOG. LUCKILY, THE SHALLOWNESS OF THE SWAMP HERE OFFERS HIM AT LEAST THE POSSIBILITY OF ESCAPE."><br>
<img src="images/darkcrystal/0045.png" width="400" alt="IN AN EERIE BOG. LUCKILY, THE SHALLOWNESS OF THE SWAMP HERE OFFERS HIM AT LEAST THE POSSIBILITY OF ESCAPE.">

**Type `SOUTH`**

<img src="images/darkcrystal/0046.png" width="400" alt="HIM AT LEAST THE POSSIBILITY OF ESCAPE. ----&gt; ENTER COMMAND?SOUTH JEN IS IN A SWAMPLAND.">

**Type `WEST`**

<img src="images/darkcrystal/0047-1.png" width="400" alt="SINKING TO HIS WAIST, JEN HAS BECOME HOPELESSLY MIRED IN A BOGGY SECTION OF THE GREAT SWAMP. WITH EACH PASSING MOMENT, HE SLIPS FURTHER INTO THE"><br>
<img src="images/darkcrystal/0047.png" width="400" alt="THE GREAT SWAMP. WITH EACH PASSING MOMENT, HE SLIPS FURTHER INTO THE ALL-CONSUMING MUCK.">

**Type `HELP`**

<img src="images/darkcrystal/0048-1.png" width="400" alt="----&gt; ENTER COMMAND?HELP JEN CRIES, &#34;HELP!&#34; THE GIRL GRABS A LONG BRANCH AND HELPS JEN OUT OF THE BOG."><br>
<img src="images/darkcrystal/0048-2.png" width="400" alt="BOG. PLEASE INSERT DISK #2, SIDE &#34;A&#34; AND PRESS RETURN."><br>
<img src="images/darkcrystal/0048-3.png" width="400" alt="THE GIRL SMILES AND SAYS, &#34;I AM KIRA. AND THIS FUZZY CREATURE IS MY PET, FIZZGIG. I THOUGHT I WAS THE ONLY LIVING GELFLING. BUT THEN, I GUESS YOU"><br>
<img src="images/darkcrystal/0048-4.png" width="400" alt="MUST HAVE THOUGHT THE SAME THING!&#34; JEN AND KIRA HAVE ENTERED A GREAT FOREST. THEY ARE ON THE BANK OF A WIDE RIVER WHICH FLOWS SOUTH. BY THE SIDE OF"><br>
<img src="images/darkcrystal/0048.png" width="400" alt="RIVER WHICH FLOWS SOUTH. BY THE SIDE OF THE RIVER IS THE DISCARDED SHELL OF A GIANT BEETLE.">

*Disk 2, side A: drop `The Dark Crystal (4am and san inc crack) disk 2A.dsk` on drive 1, and press Return.*

**Type `TURN SHELL`**

<img src="images/darkcrystal/0049-1.png" width="400" alt="WITH A MIGHTY HEAVE, THEY TURN OVER THE BEETLESHELL, EXPOSING A SMALL POUCH WHICH HAD BEEN HIDDEN UNDER THE SHELL&#39;S HOLLOW UNDERSIDE."><br>
<img src="images/darkcrystal/0049.png" width="400" alt="WHICH HAD BEEN HIDDEN UNDER THE SHELL&#39;S HOLLOW UNDERSIDE. JEN AND KIRA ARE IN A FOREST.">

**Type `GET POUCH`**

<img src="images/darkcrystal/0050.png" width="400" alt="JEN AND KIRA ARE IN A FOREST. ----&gt; ENTER COMMAND?GET POUCH JEN AND KIRA ARE IN A FOREST.">

**Type `GO SHELL`**

<img src="images/darkcrystal/0051-1.png" width="400" alt="JEN, KIRA AND FIZZGIG CLIMB INTO THE OVERTURNED BEETLESHELL AND CAST OFF INTO THE CURRENT. DOWNSTREAM THEY BEACH THE SHELL AND MAKE THEIR WAY, ON FOOT,"><br>
<img src="images/darkcrystal/0051-2.png" width="400" alt="OVERTURNED BEETLESHELL AND CAST OFF INTO THE CURRENT. DOWNSTREAM THEY BEACH THE SHELL AND MAKE THEIR WAY, ON FOOT, TO THE POD VILLAGE."><br>
<img src="images/darkcrystal/0051-3.png" width="400" alt="JEN AND KIRA HAVE ENTERED THE POD VILLAGE, THE HOME OF THE GENTLE FOREST FOLK WHO HAVE SHELTERED KIRA SINCE HER EARLY CHILDHOOD."><br>
<img src="images/darkcrystal/0051.png" width="400" alt="TO HIS HORROR, JEN SEES A CRYSTAL BAT HOVERING OVERHEAD. ITS CRYSTAL &#34;EYE&#34; IS STARING DIRECTLY AT HIM.">

**Type `LOOK`**

<img src="images/darkcrystal/0052-1.png" width="400" alt="JEN AND KIRA HAVE ENTERED THE POD VILLAGE, THE HOME OF THE GENTLE FOREST FOLK WHO HAVE SHELTERED KIRA SINCE HER EARLY CHILDHOOD."><br>
<img src="images/darkcrystal/0052.png" width="400" alt="VILLAGE, THE HOME OF THE GENTLE FOREST FOLK WHO HAVE SHELTERED KIRA SINCE HER EARLY CHILDHOOD.">

**Type `LOOK`**

<img src="images/darkcrystal/0053-1.png" width="400" alt="JEN AND KIRA HAVE ENTERED THE POD VILLAGE, THE HOME OF THE GENTLE FOREST FOLK WHO HAVE SHELTERED KIRA SINCE HER EARLY CHILDHOOD."><br>
<img src="images/darkcrystal/0053.png" width="400" alt="VILLAGE, THE HOME OF THE GENTLE FOREST FOLK WHO HAVE SHELTERED KIRA SINCE HER EARLY CHILDHOOD.">

**Type `LOOK`**

### 5. The ruins and the Landstriders

<img src="images/darkcrystal/0054-1.png" width="400" alt="JEN AND KIRA HAVE ENTERED THE POD VILLAGE, THE HOME OF THE GENTLE FOREST FOLK WHO HAVE SHELTERED KIRA SINCE HER EARLY CHILDHOOD."><br>
<img src="images/darkcrystal/0054-2.png" width="400" alt="JEN AND KIRA HEAR A FEARFUL, CLATTERING SOUND. BEFORE THEM LOOMS A GARTHIM, ONE OF THE MENACING, BEETLE-LIKE WARRIORS WHO SERVE THE SKEKSIS."><br>
<img src="images/darkcrystal/0054.png" width="400" alt="SOUND. BEFORE THEM LOOMS A GARTHIM, ONE OF THE MENACING, BEETLE-LIKE WARRIORS WHO SERVE THE SKEKSIS.">

**Type `SOUTH`**

<img src="images/darkcrystal/0055-1.png" width="400" alt="RUNNING AS FAST AS THEY CAN, JEN AND KIRA ESCAPE THE UNBELIEVABLY STRONG, BUT INCREDIBLY SLOW-WITTED, GARTHIM...THIS TIME!"><br>
<img src="images/darkcrystal/0055-2.png" width="400" alt="GARTHIM...THIS TIME! PLEASE INSERT DISK #1, SIDE &#34;B&#34; AND PRESS RETURN."><br>
<img src="images/darkcrystal/0055.png" width="400" alt="THE GELFLING PAIR ARE WANDERING THROUGH A PRIMEVAL FOREST. STRANGE CREATURES PEER DOWN AT THEM FROM THEIR PERCHES.">

*Disk 1, side B: drop `The Dark Crystal (4am and san inc crack) disk 1B.dsk` on drive 1, and press Return.*

**Type `SOUTH`**

<img src="images/darkcrystal/0056-1.png" width="400" alt="THERE IS A SLING LYING HERE AMONGST THE TREES. JEN AND KIRA ARE MAKING THEIR WAY THROUGH A SERENE FOREST OF TOWERING"><br>
<img src="images/darkcrystal/0056.png" width="400" alt="JEN AND KIRA ARE MAKING THEIR WAY THROUGH A SERENE FOREST OF TOWERING PLANTS.">

**Type `WEST`**

<img src="images/darkcrystal/0057.png" width="400" alt="JEN AND KIRA HAVE WANDERED INTO A VIRGIN FOREST OF GNARLED TREES AND VINE-LIKE CREEPERS.">

**Type `NORTH`**

<img src="images/darkcrystal/0058-1.png" width="400" alt="JEN AND KIRA HAVE WANDERED INTO THE RUINS OF WHAT APPEARS TO HAVE BEEN A GELFLING VILLAGE. THEY ARE STANDING IN FRONT OF A LARGE WALL. NEARBY ARE TWO"><br>
<img src="images/darkcrystal/0058.png" width="400" alt="GELFLING VILLAGE. THEY ARE STANDING IN FRONT OF A LARGE WALL. NEARBY ARE TWO FLAT STONES.">

**Type `SIT DOWN`**

<img src="images/darkcrystal/0059-1.png" width="400" alt="A MOMENTARY TREMOR SHAKES THE RUINS, AND THE WALL EMITS A MYSTERIOUS RUMBLING SOUND. JEN AND KIRA ARE IN SOME RUINS."><br>
<img src="images/darkcrystal/0059.png" width="400" alt="AND THE WALL EMITS A MYSTERIOUS RUMBLING SOUND. JEN AND KIRA ARE IN SOME RUINS.">

**Type `LOOK WALL`**

<img src="images/darkcrystal/0060.png" width="400" alt="THERE ARE SOME UNUSUAL HIEROGLYPHICS ON THE WALL. JEN AND KIRA ARE IN SOME RUINS.">

**Type `NORTH`**

<img src="images/darkcrystal/0061-1.png" width="400" alt="A BECKONING CREATURE DRIFTS INTO THE RUINS AND SAYS, &#34;ONCE WAS CHAMBERLAIN OF SKEKSIS. THEN NEW RULER GAIN POWER, BANISH ME FROM CASTLE. AM TIRED OF"><br>
<img src="images/darkcrystal/0061-2.png" width="400" alt="BANISH ME FROM CASTLE. AM TIRED OF KILLING. WANT PEACE. HELP YOU HEAL CRYSTAL! FOLLOW ME.&#34; THE CREATURE THEN HEADS SOUTH."><br>
<img src="images/darkcrystal/0061.png" width="400" alt="CRYSTAL! FOLLOW ME.&#34; THE CREATURE THEN HEADS SOUTH. JEN AND KIRA ARE IN SOME RUINS.">

**Type `NORTH`**

<img src="images/darkcrystal/0062-1.png" width="400" alt="----&gt; ENTER COMMAND?NORTH PLEASE INSERT DISK #2, SIDE &#34;A&#34; AND PRESS RETURN."><br>
<img src="images/darkcrystal/0062-2.png" width="400" alt="JEN AND KIRA ARE STROLLING IN A FOREST OF TOWERING TREES. BIRDS ARE SINGING AND A CREATURE STARES AT THEM FROM BEHIND A LARGE ROCK UPON WHICH A SPIRAL"><br>
<img src="images/darkcrystal/0062.png" width="400" alt="AND A CREATURE STARES AT THEM FROM BEHIND A LARGE ROCK UPON WHICH A SPIRAL HAS BEEN CARVED.">

*Disk 2, side A: drop `The Dark Crystal (4am and san inc crack) disk 2A.dsk` on drive 1, and press Return.*

**Type `NORTH`**

<img src="images/darkcrystal/0063-1.png" width="400" alt="JEN AND KIRA ARE IN A WOODED AREA OF TWISTED TREES. TO THE EAST, A GREAT RIVER CAN BE SEEN, FLOWING SOUTHWARD THROUGH THE FOREST."><br>
<img src="images/darkcrystal/0063.png" width="400" alt="TO HIS HORROR, JEN SEES A CRYSTAL BAT HOVERING OVERHEAD. ITS CRYSTAL &#34;EYE&#34; IS STARING DIRECTLY AT HIM.">

**Type `WEST`**

<img src="images/darkcrystal/0064-1.png" width="400" alt="JEN AND KIRA ARE WALKING THROUGH A QUIET THICKET OF GNARLED TREES. THE CRACKLE OF LEAVES BENEATH THEIR FEET IS THE ONLY SOUND THEY HEAR."><br>
<img src="images/darkcrystal/0064.png" width="400" alt="TO HIS HORROR, JEN SEES A CRYSTAL BAT HOVERING OVERHEAD. ITS CRYSTAL &#34;EYE&#34; IS STARING DIRECTLY AT HIM.">

**Type `SOUTH`**

<img src="images/darkcrystal/0065-1.png" width="400" alt="JEN AND KIRA ARE ON THE HILL OF THE LANDSTRIDERS. TWO LONG-LEGGED BEASTS ARE GRAZING HERE. TO HIS HORROR, JEN SEES A CRYSTAL BAT"><br>
<img src="images/darkcrystal/0065.png" width="400" alt="TO HIS HORROR, JEN SEES A CRYSTAL BAT HOVERING OVERHEAD. ITS CRYSTAL &#34;EYE&#34; IS STARING DIRECTLY AT HIM.">

**Type `RIDE LANDSTRIDER`**

<img src="images/darkcrystal/0066-1.png" width="400" alt="JEN AND KIRA, STILL CARRYING HER FAITHFUL PET, FIZZGIG, CLIMB ONTO THE BACKS OF THE TWO LANDSTRIDERS. JEN AND KIRA ARE ON THE HILL OF THE"><br>
<img src="images/darkcrystal/0066.png" width="400" alt="BACKS OF THE TWO LANDSTRIDERS. JEN AND KIRA ARE ON THE HILL OF THE LANDSTRIDERS.">

**Type `WEST`**

<img src="images/darkcrystal/0067-1.png" width="400" alt="JEN AND KIRA ARE TRAVERSING A CRAGGY BRUSHLAND, WITH LOW BUSHES AND SCRUBBY VEGETATION. A HOT, DRY WIND BLOWS OVER THE LANDSCAPE FROM THE WEST."><br>
<img src="images/darkcrystal/0067.png" width="400" alt="BRUSHLAND, WITH LOW BUSHES AND SCRUBBY VEGETATION. A HOT, DRY WIND BLOWS OVER THE LANDSCAPE FROM THE WEST.">

**Type `WEST`**

<img src="images/darkcrystal/0068.png" width="400" alt="JEN AND KIRA HAVE REACHED THE EASTERN EDGE OF A DEEP CHASM, ON THE FAR SIDE OF WHICH LIES A VAST, BARREN DESERT.">

**Type `WEST`**

<img src="images/darkcrystal/0069-1.png" width="400" alt="WITH THEIR LONG LEGS, THE LANDSTRIDERS HAVE NO TROUBLE CARRYING JEN, KIRA AND FIZZGIG SAFELY TO THE OTHER SIDE OF THE CHASM."><br>
<img src="images/darkcrystal/0069-2.png" width="400" alt="JEN AND KIRA ARE NOW ON THE WEST SIDE OF THE TREACHEROUS CHASM. AN ARID DESERT BEGINS HERE AND EXTENDS BEYOND SIGHT TO THE WEST."><br>
<img src="images/darkcrystal/0069.png" width="400" alt="OF THE TREACHEROUS CHASM. AN ARID DESERT BEGINS HERE AND EXTENDS BEYOND SIGHT TO THE WEST.">

**Type `WEST`**

<img src="images/darkcrystal/0070-1.png" width="400" alt="JEN AND KIRA ARE CROSSING A GREAT DESERT ON THE BACKS OF THE LANDSTRIDERS. THE HEAT GROWS MORE OPPRESSIVE BY THE MINUTE."><br>
<img src="images/darkcrystal/0070.png" width="400" alt="DESERT ON THE BACKS OF THE LANDSTRIDERS. THE HEAT GROWS MORE OPPRESSIVE BY THE MINUTE.">

**Type `WEST`**

<img src="images/darkcrystal/0071-1.png" width="400" alt="JEN AND KIRA HAVE ENTERED A BLEAK, BARREN LAND. THERE IS NOT A SIGN OF LIFE ANYWHERE. THE AIR HERE IS STILL AND SINISTER."><br>
<img src="images/darkcrystal/0071.png" width="400" alt="BARREN LAND. THERE IS NOT A SIGN OF LIFE ANYWHERE. THE AIR HERE IS STILL AND SINISTER.">

**Type `SOUTH`**

<img src="images/darkcrystal/0072.png" width="400" alt="----&gt; ENTER COMMAND?SOUTH THE GELFLINGS ARE RIDING THROUGH A DRY, PARCHED, EMPTY LAND.">

**Type `WEST`**

<img src="images/darkcrystal/0073-1.png" width="400" alt="LOOKING TO THE SOUTH ACROSS THE DRY, CRACKED LANDSCAPE, JEN AND KIRA CAN MAKE OUT THE TURRETS OF A DECAYED CASTLE."><br>
<img src="images/darkcrystal/0073.png" width="400" alt="CRACKED LANDSCAPE, JEN AND KIRA CAN MAKE OUT THE TURRETS OF A DECAYED CASTLE.">

**Type `SOUTH`**

### 6. Into the castle of the Skeksis

<img src="images/darkcrystal/0074-1.png" width="400" alt="JEN AND KIRA HAVE ARRIVED AT A RAVINE SURROUNDING THE CASTLE. NUMEROUS GARTHIM, DREAD WARRIORS OF DESTRUCTION, ARE APPROACHING...."><br>
<img src="images/darkcrystal/0074-2.png" width="400" alt="THE GARTHIM ATTACK! JEN, KIRA AND FIZZGIG ARE THROWN TO THE GROUND AT THE EDGE OF THE RAVINE. ONE LANDSTRIDER IS DEAD; THE OTHER IS FIGHTING A LOSING"><br>
<img src="images/darkcrystal/0074.png" width="400" alt="BATTLE. A GARTHIM LOOMS OVER THE GELFLINGS, READY TO POUNCE. THE SITUATION APPEARS HOPELESS.">

**Type `JUMP`**

<img src="images/darkcrystal/0075-1.png" width="400" alt="----&gt; ENTER COMMAND?JUMP PLEASE INSERT DISK #2, SIDE &#34;B&#34; AND PRESS RETURN."><br>
<img src="images/darkcrystal/0075-2.png" width="400" alt="JEN AND KIRA, WITH FIZZGIG IN HER ARMS, ARE PLUMMETING INTO THE RAVINE THAT SURROUNDS THE CASTLE. AMAZINGLY, KIRA SPREADS A PAIR OF WINGS, SLOWING HER"><br>
<img src="images/darkcrystal/0075.png" width="400" alt="SURROUNDS THE CASTLE. AMAZINGLY, KIRA SPREADS A PAIR OF WINGS, SLOWING HER FALL.">

*Disk 2, side B: drop `The Dark Crystal (4am and san inc crack) disk 2B.dsk` on drive 1, and press Return.*

**Type `GRAB KIRA`**

<img src="images/darkcrystal/0076-1.png" width="400" alt="----&gt; ENTER COMMAND?GRAB KIRA JEN HAS GRABBED HOLD OF KIRA, AND TOGETHER THEY ARE GLIDING SLOWLY TOWARD THE BOTTOM OF THE RAVINE."><br>
<img src="images/darkcrystal/0076-2.png" width="400" alt="JEN, KIRA AND FIZZGIG ARE AT THE BOTTOM OF A RAVINE THAT SURROUNDS THE CASTLE. THE RAVINE CURVES TO THE EAST AND WEST. BEFORE THEM, CARVED INTO THE ROCK, IS A"><br>
<img src="images/darkcrystal/0076.png" width="400" alt="THE RAVINE CURVES TO THE EAST AND WEST. BEFORE THEM, CARVED INTO THE ROCK, IS A STONE FACE.">

**Type `WEST`**

<img src="images/darkcrystal/0077-1.png" width="400" alt="JEN, KIRA AND FIZZGIG ARE AT THE BOTTOM OF A RAVINE THAT SURROUNDS THE CASTLE. THE RAVINE CURVES TO THE EAST AND WEST. BEFORE THEM, CARVED INTO THE ROCK, IS A"><br>
<img src="images/darkcrystal/0077.png" width="400" alt="THE RAVINE CURVES TO THE EAST AND WEST. BEFORE THEM, CARVED INTO THE ROCK, IS A STONE FACE.">

**Type `SEND FIZZGIG`**

<img src="images/darkcrystal/0078.png" width="400" alt="WHERE DOES JEN WANT TO SEND HIM? JEN, KIRA AND FIZZGIG ARE IN THE RAVINE SURROUNDING THE CASTLE.">

**Type `SEND BARS`**

<img src="images/darkcrystal/0079-1.png" width="400" alt="JEN AND KIRA SEND FIZZGIG THROUGH THE BARS OF THE GATE. JEN AND KIRA ARE IN THE RAVINE SURROUNDING THE CASTLE. FIZZGIG HAS"><br>
<img src="images/darkcrystal/0079.png" width="400" alt="SURROUNDING THE CASTLE. FIZZGIG HAS DISAPPEARED THROUGH THE JAWS OF THE STONE FACE.">

**Type `EAST`**

<img src="images/darkcrystal/0080-1.png" width="400" alt="BEFORE JEN CAN ACT, FIZZGIG COMES BOUNDING BACK THROUGH THE BARS OF THE GATE WITH A KEY IN HIS MOUTH. JEN, KIRA AND FIZZGIG ARE IN THE RAVINE"><br>
<img src="images/darkcrystal/0080.png" width="400" alt="GATE WITH A KEY IN HIS MOUTH. JEN, KIRA AND FIZZGIG ARE IN THE RAVINE SURROUNDING THE CASTLE.">

**Type `USE KEY`**

<img src="images/darkcrystal/0081.png" width="400" alt="OKAY. JEN, KIRA AND FIZZGIG ARE IN THE RAVINE SURROUNDING THE CASTLE.">

**Type `OPEN BARS`**

<img src="images/darkcrystal/0082-1.png" width="400" alt="WHEN JEN OPENS THE GATE, THE DOOR BEHIND IT AUTOMATICALLY SWINGS OPEN. JEN, KIRA AND FIZZGIG ARE IN THE RAVINE SURROUNDING THE CASTLE."><br>
<img src="images/darkcrystal/0082.png" width="400" alt="BEHIND IT AUTOMATICALLY SWINGS OPEN. JEN, KIRA AND FIZZGIG ARE IN THE RAVINE SURROUNDING THE CASTLE.">

**Type `SOUTH`**

<img src="images/darkcrystal/0083-1.png" width="400" alt="JEN, KIRA AND FIZZGIG ARE IN A FOUL-SMELLING UNDERGROUND SEWER SYSTEM. TUNNELS VEER OFF TO THE EAST, WEST, AND SOUTH. TO THE NORTH IS A DOOR ON WHICH"><br>
<img src="images/darkcrystal/0083.png" width="400" alt="SOUTH. TO THE NORTH IS A DOOR ON WHICH THE IMAGE OF A SERPENT CHASING ITS TAIL HAS BEEN PAINTED.">

**Type `SOUTH`**

<img src="images/darkcrystal/0084-1.png" width="400" alt="JEN AND KIRA ARE TRYING TO FIND THEIR WAY THROUGH AN UNDERGROUND SEWER SYSTEM. HERE, THE TUNNEL RUNS NORTH AND SOUTH."><br>
<img src="images/darkcrystal/0084.png" width="400" alt="WAY THROUGH AN UNDERGROUND SEWER SYSTEM. HERE, THE TUNNEL RUNS NORTH AND SOUTH.">

**Type `SOUTH`**

<img src="images/darkcrystal/0085-1.png" width="400" alt="JEN, KIRA AND FIZZGIG ARE IN A FOUL-SMELLING UNDERGROUND SEWER SYSTEM. TUNNELS VEER OFF TO THE EAST, WEST AND NORTH. TO THE SOUTH IS A DOOR ON WHICH"><br>
<img src="images/darkcrystal/0085.png" width="400" alt="NORTH. TO THE SOUTH IS A DOOR ON WHICH THE IMAGE OF A SERPENT CHASING ITS TAIL HAS BEEN PAINTED.">

**Type `WEST`**

<img src="images/darkcrystal/0086-1.png" width="400" alt="JEN AND KIRA ARE LOST IN A MAZE OF UNDERGROUND SEWER TUNNELS. THE TUNNELS HERE TRAVEL TO THE NORTH, EAST AND WEST."><br>
<img src="images/darkcrystal/0086.png" width="400" alt="UNDERGROUND SEWER TUNNELS. THE TUNNELS HERE TRAVEL TO THE NORTH, EAST AND WEST.">

**Type `WEST`**

<img src="images/darkcrystal/0087-1.png" width="400" alt="JEN, KIRA AND FIZZGIG ARE IN THE SEWER SYSTEM. ONE PASSAGEWAY RUNS EAST, AND STRANGE SOUNDS ARE COMING FROM THE PASSAGEWAY TO THE SOUTH."><br>
<img src="images/darkcrystal/0087.png" width="400" alt="SYSTEM. ONE PASSAGEWAY RUNS EAST, AND STRANGE SOUNDS ARE COMING FROM THE PASSAGEWAY TO THE SOUTH.">

**Type `SOUTH`**

<img src="images/darkcrystal/0088-1.png" width="400" alt="THE CHAMBERLAIN SUDDENLY RUSHES IN FROM THE DARK REACHES OF THE SEWERS, AND GRABS KIRA AND FIZZGIG. AS HE RUNS OFF WITH THEM, HE TOUCHES AN UNSEEN LEVER"><br>
<img src="images/darkcrystal/0088-2.png" width="400" alt="ON THE TUNNEL WALL, AND A SHOWER OF BOULDERS CASCADES FROM THE CEILING. JEN IS ALONE IN THE SEWER SYSTEM. PASSAGE TO THE EAST HAS BEEN BLOCKED BY"><br>
<img src="images/darkcrystal/0088.png" width="400" alt="HUGE BOULDERS, AND STRANGE SOUNDS ARE COMING FROM THE PASSAGEWAY TO THE SOUTH.">

**Type `SOUTH`**

<img src="images/darkcrystal/0089-1.png" width="400" alt="JEN HAS FALLEN TO THE BOTTOM OF A HUGE PIT, STARTLING SEVERAL SLEEPING GARTHIM. THEY EYE HIM WARILY, THEN ADVANCE MENACINGLY. IT LOOKS LIKE JEN"><br>
<img src="images/darkcrystal/0089.png" width="400" alt="GARTHIM. THEY EYE HIM WARILY, THEN ADVANCE MENACINGLY. IT LOOKS LIKE JEN IS IN TROUBLE!">

**Type `RUN`**

<img src="images/darkcrystal/0090-1.png" width="400" alt="POWERFUL CLAWS SLASH TOWARD JEN, NARROWLY MISSING THE TOP OF HIS HEAD, AND SMASHING A HOLE IN THE WALL BESIDE HIM."><br>
<img src="images/darkcrystal/0090.png" width="400" alt="AND SMASHING A HOLE IN THE WALL BESIDE HIM. JEN IS IN A DEEP PIT.">

**Type `GO HOLE`**

<img src="images/darkcrystal/0091-1.png" width="400" alt="JEN IS PERCHED ON A PRECARIOUS LEDGE HALFWAY UP A STEEP SHAFT. BELOW HIM IS A FLAMING LAKE OF FIRE. FLOATING HIGH ABOVE HIM IS THE DARK CRYSTAL. THERE IS"><br>
<img src="images/darkcrystal/0091.png" width="400" alt="A RAGGED HOLE IN THE SHAFT THROUGH WHICH FEARFUL, CLATTERING SOUNDS CAN BE HEARD.">

**Type `CLIMB`**

### 7. The scepter and the tower

<img src="images/darkcrystal/0092.png" width="400" alt="JEN HAS CLIMBED UP THE SHAFT INTO THE CHAMBER OF LIFE. A STEEP STAIRWAY LEADS UP TO THE EAST.">

**Type `EAST`**

<img src="images/darkcrystal/0093-1.png" width="400" alt="JEN IS IN A NORTH-SOUTH HALLWAY. THERE ARE CRACKS IN THE WALLS AND CEILING. TO THE WEST IS A PASSAGE WHICH LEADS TO A STEEPLY DESCENDING STAIRWAY."><br>
<img src="images/darkcrystal/0093.png" width="400" alt="STEEPLY DESCENDING STAIRWAY. JEN HEARS THE SOUND OF FOOTSTEPS APPROACHING FROM THE NORTH.">

**Type `SOUTH`**

<img src="images/darkcrystal/0094-1.png" width="400" alt="JEN IS AT THE SOUTHERN END OF A FILTH-ENCRUSTED NORTH-SOUTH HALLWAY. THERE ARE DOORWAYS TO THE WEST AND TO THE EAST."><br>
<img src="images/darkcrystal/0094.png" width="400" alt="THE EAST. JEN HEARS THE SOUND OF FOOTSTEPS APPROACHING FROM THE NORTH.">

**Type `WEST`**

<img src="images/darkcrystal/0095-1.png" width="400" alt="AS JEN DARTS THROUGH THE DOORWAY, HE HEARS A GROUP OF SKEKSIS COME DOWN THE HALL AND GO INTO THE ROOM DIRECTLY ACROSS FROM HIM."><br>
<img src="images/darkcrystal/0095.png" width="400" alt="ACROSS FROM HIM. JEN IS IN A DANK AND DUSTY CLOSET. THERE IS A DOORWAY TO THE EAST.">

**Type `EAST`**

<img src="images/darkcrystal/0096.png" width="400" alt="THERE IS A DOORWAY TO THE EAST. ----&gt; ENTER COMMAND?EAST JEN IS IN A HALLWAY.">

**Type `EAST`**

<img src="images/darkcrystal/0097-1.png" width="400" alt="JEN IS STANDING IN THE DOORWAY OF THE DINING ROOM. SEATED AT A LONG TABLE, NEAR A CURTAIN-COVERED WALL, ARE SEVERAL SKEKSIS. UNAWARE OF JEN&#39;S"><br>
<img src="images/darkcrystal/0097.png" width="400" alt="SEVERAL SKEKSIS. UNAWARE OF JEN&#39;S PRESENCE, THEY ARE ARGUING LOUDLY WITH EACH OTHER.">

**Type `GO CURTAIN`**

<img src="images/darkcrystal/0098-1.png" width="400" alt="JEN GETS BEHIND THE CURTAIN AND INCHES TOWARD THE SKEKSIS. WHEN HE IS CLOSER, HE HEARS THEM ARGUING ABOUT A SECRET PANEL SOMEWHERE IN THE TOWER. HE WAITS"><br>
<img src="images/darkcrystal/0098-2.png" width="400" alt="TILL NO ONE IS LOOKING, THEN, QUICKLY, HE SLIPS OUT FROM BEHIND THE CURTAIN AND LEAVES THE ROOM. JEN IS IN A HALLWAY."><br>
<img src="images/darkcrystal/0098.png" width="400" alt="HE SLIPS OUT FROM BEHIND THE CURTAIN AND LEAVES THE ROOM. JEN IS IN A HALLWAY.">

**Type `NORTH`**

<img src="images/darkcrystal/0099.png" width="400" alt="JEN IS IN A HALLWAY. ----&gt; ENTER COMMAND?NORTH JEN IS IN A HALLWAY.">

**Type `NORTH`**

<img src="images/darkcrystal/0100-1.png" width="400" alt="JEN IS AT A JUNCTION OF HALLWAYS, ONE RUNNING EAST AND WEST, THE OTHER SOUTH. SIGNS OF FILTH AND DECAY ARE EVERYWHERE."><br>
<img src="images/darkcrystal/0100.png" width="400" alt="RUNNING EAST AND WEST, THE OTHER SOUTH. SIGNS OF FILTH AND DECAY ARE EVERYWHERE.">

**Type `WEST`**

<img src="images/darkcrystal/0101-1.png" width="400" alt="THERE IS A SCEPTER HERE. JEN HAS WANDERED INTO THE THRONE ROOM. A LARGE THRONE SITS AMID THE FILTH THAT COATS THE DESERTED CHAMBER."><br>
<img src="images/darkcrystal/0101.png" width="400" alt="JEN HAS WANDERED INTO THE THRONE ROOM. A LARGE THRONE SITS AMID THE FILTH THAT COATS THE DESERTED CHAMBER.">

**Type `GET SCEPTER`**

<img src="images/darkcrystal/0102.png" width="400" alt="COATS THE DESERTED CHAMBER. ----&gt; ENTER COMMAND?GET SCEPTER JEN IS IN THE THRONE ROOM.">

**Type `EAST`**

<img src="images/darkcrystal/0103.png" width="400" alt="JEN IS IN THE THRONE ROOM. ----&gt; ENTER COMMAND?EAST JEN IS IN A HALLWAY.">

**Type `EAST`**

<img src="images/darkcrystal/0104-1.png" width="400" alt="JEN HAS REACHED THE EAST END OF AN EAST-WEST PASSAGE. THROUGH A DOORWAY AT THE END OF THE PASSAGE, HE CAN SEE A NARROW RAMP WINDING ITS WAY UPWARD."><br>
<img src="images/darkcrystal/0104.png" width="400" alt="EAST-WEST PASSAGE. THROUGH A DOORWAY AT THE END OF THE PASSAGE, HE CAN SEE A NARROW RAMP WINDING ITS WAY UPWARD.">

**Type `EAST`**

<img src="images/darkcrystal/0105.png" width="400" alt="JEN IS STANDING ON THE ROTTING FLOOR OF A DESERTED TOWER. FROM HERE A NARROW RAMP WINDS DOWN TO THE HALLWAY BELOW.">

**Type `USE HOOK`**

<img src="images/darkcrystal/0106-1.png" width="400" alt="USING THE HOOK AT THE END OF THE SCEPTER, JEN PULLS ON THE LATCH AND THE SECRET PANEL OPENS. JEN IS IN A TOWER OF THE CASTLE."><br>
<img src="images/darkcrystal/0106.png" width="400" alt="SCEPTER, JEN PULLS ON THE LATCH AND THE SECRET PANEL OPENS. JEN IS IN A TOWER OF THE CASTLE.">

**Type `EAST`**

<img src="images/darkcrystal/0107-1.png" width="400" alt="JEN IS AT THE BOTTOM OF A NARROW STAIRWAY. GAZING UPWARD, HE IS SEIZED BY A FEELING OF GRAVE APPREHENSION. THERE IS AN OPENING IN THE WALL TO THE"><br>
<img src="images/darkcrystal/0107.png" width="400" alt="BY A FEELING OF GRAVE APPREHENSION. THERE IS AN OPENING IN THE WALL TO THE WEST.">

**Type `UP`**

<img src="images/darkcrystal/0108.png" width="400" alt="JEN HAS REACHED THE TOP OF THE NARROW STAIRWAY. STRANGE NOISES ARE COMING THROUGH AN OPEN DOORWAY TO THE EAST.">

**Type `EAST`**

### 8. The Great Conjunction

<img src="images/darkcrystal/0109-1.png" width="400" alt="JEN IS ON A BALCONY HIGH ABOVE THE CRYSTAL CHAMBER. THE SKEKSIS HAVE GATHERED IN A CIRCLE BENEATH THE DARK CRYSTAL. THEIR POWER CEREMONY, TIMED TO"><br>
<img src="images/darkcrystal/0109.png" width="400" alt="COINCIDE WITH THE GREAT CONJUNCTION, IS BEGINNING. KIRA AND FIZZGIG ARE WITH THE SKEKSIS.">

**Type `JUMP CRYSTAL`**

<img src="images/darkcrystal/0110-1.png" width="400" alt="JEN HAS LANDED ATOP THE DARK CRYSTAL! BUT THE IMPACT HAS CAUSED HIM TO DROP THE SHARD, WHICH NOW LIES PRECARIOUSLY ON THE BRINK OF THE SHAFT BENEATH THE"><br>
<img src="images/darkcrystal/0110-2.png" width="400" alt="ON THE BRINK OF THE SHAFT BENEATH THE CRYSTAL. JEN LOOKS UP THROUGH AN OPEN PORTAL IN THE CEILING. THE THREE SUNS ARE TOUCHING...."><br>
<img src="images/darkcrystal/0110-3.png" width="400" alt="QUICKLY, KIRA PICKS UP THE CRYSTAL SHARD. SHE IS ABOUT TO THROW IT UP TO JEN, BUT THE RITUAL-MASTER, DRAWN DAGGER IN HAND, WARNS HER, &#34;GIVE ME"><br>
<img src="images/darkcrystal/0110.png" width="400" alt="SHARD, GELFLING, AND YOU GO IN PEACE. OTHERWISE, NO CHOICE BUT TO KILL YOU!&#34; DOES JEN WANT TO SAVE KIRA?">

**Type `NO`**

<img src="images/darkcrystal/0111-1.png" width="400" alt="JEN SITS ASTRIDE THE CRYSTAL. KIRA HAS THROWN THE SHARD UP TO HIM AND HE HOLDS IT IN HIS HAND. THE EDGES OF THE THREE SUNS ARE OVERLAPPING...."><br>
<img src="images/darkcrystal/0111.png" width="400" alt="THROWN THE SHARD UP TO HIM AND HE HOLDS IT IN HIS HAND. THE EDGES OF THE THREE SUNS ARE OVERLAPPING....">

**Type `INSERT SHARD`**

<img src="images/darkcrystal/0112-1.png" width="400" alt="AS THE THREE SUNS BECOME ONE, AND A FLOOD OF BLINDING LIGHT WASHES OVER HIM, JEN PLUNGES THE SHARD DEEP INTO THE WOUND IN THE CRYSTAL."><br>
<img src="images/darkcrystal/0112-2.png" width="400" alt="JEN HAS HEALED THE WOUND IN THE CRYSTAL AND RESTORED IT TO ITS ORIGINAL BRILLIANCE! HE IS NOW ON THE FLOOR OF THE CRYSTAL CHAMBER, WHICH IS BATHED IN"><br>
<img src="images/darkcrystal/0112-3.png" width="400" alt="RADIANT LIGHT. THE GARTHIM ARE CRACKING AND FALLING APART. THE FILTHY WALLS OF THE CASTLE ARE CRUMBLING, REVEALING ITS ORIGINAL CRYSTALLINE PURITY AND BEAUTY."><br>
<img src="images/darkcrystal/0112-4.png" width="400" alt="THE CASTLE ARE CRUMBLING, REVEALING ITS ORIGINAL CRYSTALLINE PURITY AND BEAUTY. THE EVIL REIGN OF THE SKEKSIS IS OVER. HOWEVER...."><br>
<img src="images/darkcrystal/0112-5.png" width="400" alt="KIRA, STABBED AT THE MOMENT OF THE GREAT CONJUNCTION, LIES IN JEN&#39;S ARMS. THE GELFLING SOBS UNCONTROLLABLY AS HE CRADLES HER LIFELESS BODY."><br>
<img src="images/darkcrystal/0112.png" width="400" alt="GREAT CONJUNCTION, LIES IN JEN&#39;S ARMS. THE GELFLING SOBS UNCONTROLLABLY AS HE CRADLES HER LIFELESS BODY.">

**Type `KISS KIRA`**

### The end

<img src="images/darkcrystal/0113-1.png" width="400" alt="WHEN JEN KISSES KIRA, SHE OPENS HER EYES, AND THE LIFE THAT THE SKEKSIS TOOK FROM HER BEFORE JEN HEALED THE CRYSTAL IS REKINDLED. FINALLY, THE TWO"><br>
<img src="images/darkcrystal/0113-2.png" width="400" alt="GELFLINGS, AND ALL CREATURES, CAN LIVE PEACEFULLY TOGETHER IN A WORLD TO WHICH HARMONY, AFTER A THOUSAND YEARS OF DARKNESS, HAS BEEN RESTORED."><br>
<img src="images/darkcrystal/0113.png" width="400" alt="PEACEFULLY TOGETHER IN A WORLD TO WHICH HARMONY, AFTER A THOUSAND YEARS OF DARKNESS, HAS BEEN RESTORED. THANKS FOR PLAYING &#34;THE DARK CRYSTAL!&#34;">

<!-- End of the walkthrough -->

## What next

[Time Zone](timezone.md) is the biggest of the adventures of Sierra On-Line
on the Apple \]\[, and [Mystery House](mysteryhouse.md) the first.
