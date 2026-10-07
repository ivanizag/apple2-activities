# Time Zone, from the start to the end

[Back to the activities](../README.md)

**Time Zone**, *Hi-Res Adventure #5* of On-Line Systems, came out in 1982 for
the Apple \]\[, written by Ken and Roberta Williams, produced by Bob Davis
and Jeff Stephenson, with graphics by Terry Pierce and Michelle Pritchard. An
adventure of pictures and words, as their *Mystery House* and *Wizard and the
Princess* before it, but far bigger: it came on six diskettes, both sides
used, twelve sides of pictures. A time machine in a field behind your house
takes you to six continents, from 400 million years before Christ to 2082,
and from there to the planet Neburon, whose ruler you have to stop.

You type one or two words, `NORTH`, `GET ROCK`, `GIVE HARE`, and the game
answers under the picture, four lines at a time. In the time machine, the
orange dial sets the year and the blue one the place, and `PULL LEVER` goes
there; the button next to the lever goes home. The game asks for a side of
its disks by its number and letter, `1B` to `6L`, each time it needs another.

This page plays the whole game, all 1,058 commands, with a picture of the
screen after each one: 1,464 pictures. It is long: the game was. The
pictures are of a colour television, with the four lines of text under them
drawn white and sharp, as a monochrome monitor showed them: the television
blurred them into fringes of colour, hard to read.

## What you need

**Time Zone v1.1 (4am and san inc crack)**, from
[its page on the Internet Archive](https://archive.org/details/TimeZone4amCrack):
the zip `Time Zone (4am and san inc crack).zip`, with a file for each side,
from `Time Zone (4am and san inc crack) disk A.dsk` to
`Time Zone (4am and san inc crack) disk L.dsk`. `./fetch-disks.sh` in this
repository downloads them into `disks/` and checks them. The game tells the
sides apart by the volume number of the disk, which a `.dsk` file does not
keep; this crack keeps it inside the disks instead, so the game still knows
when the wrong one is in the drive.

The commands of this page are also in this repository, a line each,
[listings/timezone.txt](listings/timezone.txt). They follow the walkthrough of
[The Sierra Help Pages](https://sierrahelp.com/Walkthroughs/TimeZoneWalkthrough.html),
with what the machine showed it needs: `GO MACHINE` where it says `MACHINE`,
`EAST` where it says `ESAT`, and the `LOOK`s counted where it says to look
until something happens.

The [manual of Time Zone](https://www.mocagh.org/sierra/timezone-manual.pdf)
is on the Museum of Computer Adventure Game History.

## The machine

An Apple \]\[+:

- the 6502 processor at 1 MHz and 48 KB of memory;
- a Disk II controller card in slot 6, with side A in drive 1.

```bash
izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s6 'diskii,disk1=disks/Time Zone (4am and san inc crack) disk A.dsk'
```

## Playing it

1. **Start izapple2** with the command above. The title waits for a key;
   **press Return**.

   ![The title](images/timezone/title.png)

2. **Type `1`** to play, and **press Return**.

   ![The menu](images/timezone/menu.png)

   `2` would make a diskette to save the game on, which this page does not
   need.

3. **Change the disk** each time the game asks for one: drop the file of
   that side on the area of drive 1 of the window of izapple2, and press
   Return. The first is side 1B, to start; the page says each one where it
   comes.

4. **Type the commands**, each followed by Return. When the game stops in
   the middle of an answer to let you read it, **press Return** to go on.
   When the picture changes on the way, as when you walk into another
   place, the page shows each picture; when it stays the same, the page
   shows it once, and under it only the lines of text that each stop
   added.

## The walkthrough

<!-- The walkthrough, written by the generator -->

### Home, 1982

**1.**

<img src="images/timezone/0001-1.png" width="400" alt="WHICH WOULD YOU LIKE? 1 PLEASE INSERT DISK NUMBER 1B AND PRESS RETURN."><br>
<img src="images/timezone/0001.png" width="400" alt="YOU ARE IN FRONT OF YOUR OWN HOUSE.">

*Side 1B: drop `Time Zone (4am and san inc crack) disk B.dsk` on drive 1, and press Return.*

**Type `NORTH`**

**2.**

<img src="images/timezone/0002-1.png" width="400" alt="AS YOU AWAKEN FROM A HEAVY NIGHT&#39;S SLEEP, YOU ARE HAUNTED BY THE MEMORY OF A STRANGE DREAM. ALTHOUGH IT IS NOT ENTIRELY CLEAR, YOU CAN REMEMBER"><br>
<img src="images/timezone/0002-2.png" width="400" alt="SIGNIFICANT PORTIONS OF IT ... A TERRESTIAL GUARDIAN OR KEEPER, OF SORTS, HAS CHOSEN YOU FOR THE TASK OF SECURING THE EARTH&#39;S FUTURE BY"><br>
<img src="images/timezone/0002-3.png" width="400" alt="DESTROYING THE EVIL RULER OF THE DISTANT PLANET NEBURON AND OFFERS YOU THE ABILITY OF TIME AND SPACE TRAVEL IN ORDER TO DO SO. SOMEWHAT"><br>
<img src="images/timezone/0002-4.png" width="400" alt="NERVOUS BY THE REALISM OF YOUR DREAM, YOU DECIDE TO TAKE A WALK TO CLEAR YOUR HEAD. YOU ARE IN FRONT OF YOUR OWN HOUSE."><br>
<img src="images/timezone/0002.png" width="400" alt="YOU DECIDE TO TAKE A WALK TO CLEAR YOUR HEAD. YOU ARE IN FRONT OF YOUR OWN HOUSE.">

**Type `NORTH`**

**3.**

<img src="images/timezone/0003.png" width="400" alt="YOU ARE IN A FIELD OF DRY GRASS. THERE IS A STRANGE LOOKING MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `GO MACHINE`**

**4.**

<img src="images/timezone/0004-1.png" width="400" alt="LOOKING AT THE STRANGE OBJECT, YOU ARE AT FIRST CURIOUS. SUDDENLY, THE DREAM COMES RUSHING BACK TO YOU. YOU STAND THERE WITH YOUR HEART LEAPING"><br>
<img src="images/timezone/0004-2.png" width="400" alt="OUT OF YOUR CHEST, TOO STUNNED TO MOVE AND FILLED WITH A STRANGE MIXTURE OF FEAR AND EXCITEMENT AS YOU REALIZE ..... IT WASN&#39;T A DREAM AT ALL."><br>
<img src="images/timezone/0004-3.png" width="400" alt="THE ADVENTURE BEGINS ..... YOU ARE IN A FIELD OF DRY GRASS. THERE IS A STRANGE LOOKING MACHINE HERE. IT APPEARS TO BE PULSATING."><br>
<img src="images/timezone/0004.png" width="400" alt="YOU ARE IN A FIELD OF DRY GRASS. THERE IS A STRANGE LOOKING MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `GO MACHINE`**

**5.**

<img src="images/timezone/0005-1.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0005.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `GET MASK`**

**6.**

<img src="images/timezone/0006-1.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0006.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EXIT MACHINE`**

**7.**

<img src="images/timezone/0007.png" width="400" alt="YOU ARE IN A FIELD OF DRY GRASS. THERE IS A STRANGE LOOKING MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `DROP MASK`**

**8.**

<img src="images/timezone/0008.png" width="400" alt="YOU ARE IN A FIELD OF DRY GRASS. THERE IS A STRANGE LOOKING MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `GO MACHINE`**

**9.**

<img src="images/timezone/0009-1.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0009.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `SIT`**

**10.**

<img src="images/timezone/0010-1.png" width="400" alt="O.K. THE CHAIR IS COMFORTABLE. AS YOU SIT DOWN, THE MACHINE SEEMS TO COME TO LIFE. YOU ARE INSIDE A TIME MACHINE. YOU HEAR"><br>
<img src="images/timezone/0010.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT. --------------- ENTER COMMAND">

**Type `SET ORANGE`**

**11.**

<img src="images/timezone/0011-1.png" width="400" alt="TO WHAT? YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0011.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `400MILBC`**

**12.**

<img src="images/timezone/0012-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0012.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `SET BLUE`**

**13.**

<img src="images/timezone/0013-1.png" width="400" alt="TO WHAT? YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0013.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EUROPE`**

**14.**

<img src="images/timezone/0014-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0014.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `PULL LEVER`**

**15.**

<img src="images/timezone/0015-1.png" width="400" alt="THE MACHINE STARTS VIBRATING VIOLENTLY AND YOU FEEL DIZZY, ALMOST TO THE POINT OF PASSING OUT. SUDDENLY YOU FEEL A TERRIBLE JOLT AND THE MACHINE IS STILL."><br>
<img src="images/timezone/0015-2.png" width="400" alt="THE CONTINENTS, AS WE KNOW THEM, DID NOT EXIST IN THAT ERA. YOU ARE IN THE PREHISTORIC AGE."><br>
<img src="images/timezone/0015-3.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0015.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EXIT MACHINE`**

### Europe, 400 million BC

**16.**

<img src="images/timezone/0016.png" width="400" alt="YOU ARE IN A PRETTY GREEN MEADOW. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `WEST`**

**17.**

<img src="images/timezone/0017.png" width="400" alt="BE PULSATING. --------------- ENTER COMMAND?WEST YOU ARE IN A PRETTY GREEN MEADOW.">

**Type `SOUTH`**

**18.**

<img src="images/timezone/0018.png" width="400" alt="YOU ARE AT THE EDGE OF A LARGE LAKE THAT EXTENDS TO THE EAST. THERE IS A BRONTOSAURUS IN THE LAKE.">

**Type `SOUTH`**

**19.**

<img src="images/timezone/0019.png" width="400" alt="BRONTOSAURUS IN THE LAKE. --------------- ENTER COMMAND?SOUTH YOU ARE IN A CREEPY JUNGLE.">

**Type `EAST`**

**20.**

<img src="images/timezone/0020.png" width="400" alt="YOU ARE IN A CREEPY JUNGLE. --------------- ENTER COMMAND?EAST YOU ARE IN A JUNGLE.">

**Type `EAST`**

**21.**

<img src="images/timezone/0021.png" width="400" alt="OH, OH. I SEE A PTERIDACTYL OVERHEAD. YOU ARE IN A JUNGLE. THERE IS A CAVE ENTRANCE TO THE EAST.">

**Type `GO CAVE`**

**22.**

<img src="images/timezone/0022.png" width="400" alt="--------------- ENTER COMMAND?GO CAVE YOU ARE INSIDE A CAVE. THE CAVE ENDS RIGHT HERE.">

**Type `EXIT CAVE`**

**23.**

<img src="images/timezone/0023.png" width="400" alt="--------------- ENTER COMMAND?EXIT CAVE YOU ARE IN A JUNGLE. THERE IS A CAVE ENTRANCE TO THE EAST.">

**Type `SOUTH`**

**24.**

<img src="images/timezone/0024.png" width="400" alt="--------------- ENTER COMMAND?SOUTH THERE IS A POINTED STICK HERE. YOU ARE IN A JUNGLE.">

**Type `GET STICK`**

**25.**

<img src="images/timezone/0025.png" width="400" alt="YOU ARE IN A JUNGLE. --------------- ENTER COMMAND?GET STICK YOU ARE IN A JUNGLE.">

**Type `WEST`**

**26.**

<img src="images/timezone/0026.png" width="400" alt="YOU ARE IN A JUNGLE. --------------- ENTER COMMAND?WEST YOU ARE IN A JUNGLE.">

**Type `NORTH`**

**27.**

<img src="images/timezone/0027.png" width="400" alt="YOU ARE IN A JUNGLE. --------------- ENTER COMMAND?NORTH YOU ARE IN A JUNGLE.">

**Type `WEST`**

**28.**

<img src="images/timezone/0028.png" width="400" alt="YOU ARE IN A JUNGLE. --------------- ENTER COMMAND?WEST YOU ARE IN A CREEPY JUNGLE.">

**Type `NORTH`**

**29.**

<img src="images/timezone/0029.png" width="400" alt="YOU ARE AT THE EDGE OF A LARGE LAKE THAT EXTENDS TO THE EAST. THERE IS A BRONTOSAURUS IN THE LAKE.">

**Type `NORTH`**

**30.**

<img src="images/timezone/0030.png" width="400" alt="BRONTOSAURUS IN THE LAKE. --------------- ENTER COMMAND?NORTH YOU ARE IN A PRETTY GREEN MEADOW.">

**Type `EAST`**

**31.**

<img src="images/timezone/0031.png" width="400" alt="YOU ARE IN A PRETTY GREEN MEADOW. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `GO MACHINE`**

**32.**

<img src="images/timezone/0032-1.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0032.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `SIT`**

**33.**

<img src="images/timezone/0033-1.png" width="400" alt="O.K. THE CHAIR IS COMFORTABLE. AS YOU SIT DOWN, THE MACHINE SEEMS TO COME TO LIFE. YOU ARE INSIDE A TIME MACHINE. YOU HEAR"><br>
<img src="images/timezone/0033.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT. --------------- ENTER COMMAND">

**Type `10000 BC`**

**34.**

<img src="images/timezone/0034-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0034.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `PULL LEVER`**

**35.**

<img src="images/timezone/0035-1.png" width="400" alt="THE MACHINE STARTS VIBRATING VIOLENTLY AND YOU FEEL DIZZY, ALMOST TO THE POINT OF PASSING OUT. SUDDENLY YOU FEEL A TERRIBLE JOLT AND THE MACHINE IS STILL."><br>
<img src="images/timezone/0035-2.png" width="400" alt="THE CONTINENTS, AS WE KNOW THEM, DID NOT EXIST IN THAT ERA. YOU ARE IN THE STONE AGE."><br>
<img src="images/timezone/0035-3.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0035.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EXIT MACHINE`**

### Europe, 10,000 BC

**36.**

<img src="images/timezone/0036-1.png" width="400" alt="YOU ARE IN A NORTH/SOUTH CANYON. STEEP CLIFFS LEAD UPWARDS TOWARDS THE EAST AND WEST. THERE IS A TIME MACHINE HERE.IT APPEARS TO BE PULSATING."><br>
<img src="images/timezone/0036.png" width="400" alt="CLIFFS LEAD UPWARDS TOWARDS THE EAST AND WEST. THERE IS A TIME MACHINE HERE.IT APPEARS TO BE PULSATING.">

**Type `NORTH`**

**37.**

<img src="images/timezone/0037-1.png" width="400" alt="THERE IS A ROCK HERE. YOU ARE STANDING IN A ROCKY CANYON. CLIMBABLE CLIFFS LEAD UP TOWARDS THE NORTH AND WEST."><br>
<img src="images/timezone/0037.png" width="400" alt="YOU ARE STANDING IN A ROCKY CANYON. CLIMBABLE CLIFFS LEAD UP TOWARDS THE NORTH AND WEST.">

**Type `GET ROCK`**

**38.**

<img src="images/timezone/0038.png" width="400" alt="YOU ARE STANDING IN A ROCKY CANYON. CLIMBABLE CLIFFS LEAD UP TOWARDS THE NORTH AND WEST.">

**Type `WEST`**

**39.**

<img src="images/timezone/0039.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE AT THE EDGE OF A CLIMBABLE CLIFF.">

**Type `WEST`**

**40.**

<img src="images/timezone/0040.png" width="400" alt="CLIFF. --------------- ENTER COMMAND?WEST YOU ARE WANDERING IN SOME GRASS LANDS.">

**Type `NORTH`**

**41.**

<img src="images/timezone/0041.png" width="400" alt="YOU ARE WANDERING IN SOME GRASS LANDS. --------------- ENTER COMMAND?NORTH YOU ARE IN SOME GRASSLANDS">

**Type `NORTH`**

**42.**

<img src="images/timezone/0042.png" width="400" alt="YOU SEE A LARGE HERD OF MASTODONS IN THE DISTANCE. YOU ARE IN THE GRASSLANDS">

**Type `NORTH`**

**43.**

<img src="images/timezone/0043.png" width="400" alt="YOU HAVE STARTLED A HERD OF MASTODONS. THEY HAVE STARTED TO STAMPEDE! YOU ARE IN SOME GRASSLANDS.">

**Type `CLIMB TREE`**

**44.**

<img src="images/timezone/0044.png" width="400" alt="YOU HEAR THE MASTODON STAMPEDE BELOW YOU. YOU ARE UP IN A LARGE TREE.">

**Type `DOWN`**

**45.**

<img src="images/timezone/0045.png" width="400" alt="YOU ARE UP IN A LARGE TREE. --------------- ENTER COMMAND?DOWN YOU ARE IN SOME GRASSLANDS.">

**Type `NORTH`**

**46.**

<img src="images/timezone/0046.png" width="400" alt="YOU ARE IN SOME GRASSLANDS. --------------- ENTER COMMAND?NORTH YOU ARE IN SOME LOVELY ROLLING HILLS.">

**Type `WEST`**

**47.**

<img src="images/timezone/0047.png" width="400" alt="YOU ARE IN SOME LOVELY ROLLING HILLS. --------------- ENTER COMMAND?WEST YOU ARE LOST IN THE MOUNTAINS.">

**Type `NORTH`**

**48.**

<img src="images/timezone/0048.png" width="400" alt="--------------- ENTER COMMAND?NORTH THERE ARE TWO SMALL STICKS HERE. YOU ARE TRAVELING THROUGH THE MOUNTAINS">

**Type `GET STICKS`**

**49.**

<img src="images/timezone/0049.png" width="400" alt="--------------- ENTER COMMAND?GET STICKS YOU ARE TRAVELING THROUGH THE MOUNTAINS">

**Type `SOUTH`**

**50.**

<img src="images/timezone/0050.png" width="400" alt="YOU ARE TRAVELING THROUGH THE MOUNTAINS --------------- ENTER COMMAND?SOUTH YOU ARE LOST IN THE MOUNTAINS.">

**Type `EAST`**

**51.**

<img src="images/timezone/0051.png" width="400" alt="YOU ARE LOST IN THE MOUNTAINS. --------------- ENTER COMMAND?EAST YOU ARE IN SOME LOVELY ROLLING HILLS.">

**Type `NORTH`**

**52.**

<img src="images/timezone/0052.png" width="400" alt="WATCH OUT!! A SABERTOOTH TIGER IS GOING TO POUNCE ON YOU. YOU ARE IN THE MOUNTAINS.">

**Type `THROW STICK`**

**53.**

<img src="images/timezone/0053-1.png" width="400" alt="AS THE SABERTOOTH TIGER LEAPS, YOU THROW THE POINTED STICK WHICH LODGES IN HIS BODY. STARTLED, THE BIG CAT RUNS AWAY WITH THE STICK STILL IN HIM."><br>
<img src="images/timezone/0053.png" width="400" alt="HIS BODY. STARTLED, THE BIG CAT RUNS AWAY WITH THE STICK STILL IN HIM. YOU ARE IN THE MOUNTAINS.">

**Type `NORTH`**

**54.**

<img src="images/timezone/0054.png" width="400" alt="YOU ARE IN THE MOUNTAINS. --------------- ENTER COMMAND?NORTH YOU ARE ROAMING THROUGH THE MOUNTAINS.">

**Type `EAST`**

**55.**

<img src="images/timezone/0055.png" width="400" alt="YOU ARE WANDERING IN THE MOUNTAINS. THERE IS A LARGE CAVE OPENING TO THE NORTH.">

**Type `SOUTH`**

**56.**

<img src="images/timezone/0056.png" width="400" alt="NORTH. --------------- ENTER COMMAND?SOUTH YOU ARE IN THE MOUNTAINS.">

**Type `SOUTH`**

**57.**

<img src="images/timezone/0057.png" width="400" alt="YOU ARE IN THE MOUNTAINS. --------------- ENTER COMMAND?SOUTH YOU ARE LOST IN THE MOUNTAINS.">

**Type `WEST`**

**58.**

<img src="images/timezone/0058.png" width="400" alt="--------------- ENTER COMMAND?WEST THERE IS A LARGE HARE HERE. YOU ARE IN THE MOUNTAINS.">

**Type `THROW ROCK`**

**59.**

<img src="images/timezone/0059-1.png" width="400" alt="YOU THROW THE ROCK AT THE HARE AND KILL IT. THERE IS A ROCK HERE. THERE IS A LARGE HARE HERE."><br>
<img src="images/timezone/0059.png" width="400" alt="THERE IS A ROCK HERE. THERE IS A LARGE HARE HERE. YOU ARE IN THE MOUNTAINS.">

**Type `GET HARE`**

**60.**

<img src="images/timezone/0060.png" width="400" alt="--------------- ENTER COMMAND?GET HARE THERE IS A ROCK HERE. YOU ARE IN THE MOUNTAINS.">

**Type `EAST`**

**61.**

<img src="images/timezone/0061.png" width="400" alt="YOU ARE IN THE MOUNTAINS. --------------- ENTER COMMAND?EAST YOU ARE LOST IN THE MOUNTAINS.">

**Type `NORTH`**

**62.**

<img src="images/timezone/0062.png" width="400" alt="YOU ARE LOST IN THE MOUNTAINS. --------------- ENTER COMMAND?NORTH YOU ARE IN THE MOUNTAINS.">

**Type `NORTH`**

**63.**

<img src="images/timezone/0063.png" width="400" alt="YOU ARE WANDERING IN THE MOUNTAINS. THERE IS A LARGE CAVE OPENING TO THE NORTH.">

**Type `GO CAVE`**

**64.**

<img src="images/timezone/0064.png" width="400" alt="THERE IS A STONE HAMMER HERE. YOU ARE INSIDE A LARGE CAVE. THERE ARE TWO MEAN-LOOKING CAVEMEN HERE.">

**Type `MAKE FIRE`**

**65.**

<img src="images/timezone/0065-1.png" width="400" alt="YOU QUICKLY RUB THE TWO STICKS TOGETHER AND A SMALL FIRE STARTS. THE CAVEMEN ARE DELIGHTED AND AMAZED. IN GRATITUDE FOR YOUR TEACHING THEM"><br>
<img src="images/timezone/0065-2.png" width="400" alt="ABOUT FIRE, THE CAVEMEN OFFER THE HAMMER TO YOU. THERE IS A STONE HAMMER HERE. YOU ARE INSIDE A LARGE CAVE. THERE ARE"><br>
<img src="images/timezone/0065.png" width="400" alt="THERE IS A STONE HAMMER HERE. YOU ARE INSIDE A LARGE CAVE. THERE ARE TWO MEAN-LOOKING CAVEMEN HERE.">

**Type `GIVE HARE`**

**66.**

<img src="images/timezone/0066-1.png" width="400" alt="THE CAVEMEN TAKE THE HARE AND GRUDGINGLY DECIDE NOT TO KILL YOU. THERE IS A STONE HAMMER HERE. YOU ARE INSIDE A LARGE CAVE. THERE ARE"><br>
<img src="images/timezone/0066.png" width="400" alt="THERE IS A STONE HAMMER HERE. YOU ARE INSIDE A LARGE CAVE. THERE ARE TWO MEAN-LOOKING CAVEMEN HERE.">

**Type `GET HAMMER`**

**67.**

<img src="images/timezone/0067.png" width="400" alt="YOU ARE INSIDE A LARGE CAVE. THERE ARE TWO MEAN-LOOKING CAVEMEN HERE.">

**Type `SOUTH`**

**68.**

<img src="images/timezone/0068.png" width="400" alt="YOU ARE WANDERING IN THE MOUNTAINS. THERE IS A LARGE CAVE OPENING TO THE NORTH.">

**Type `WEST`**

**69.**

<img src="images/timezone/0069.png" width="400" alt="NORTH. --------------- ENTER COMMAND?WEST YOU ARE ROAMING THROUGH THE MOUNTAINS.">

**Type `SOUTH`**

**70.**

<img src="images/timezone/0070.png" width="400" alt="YOU ARE ROAMING THROUGH THE MOUNTAINS. --------------- ENTER COMMAND?SOUTH YOU ARE IN THE MOUNTAINS.">

**Type `SOUTH`**

**71.**

<img src="images/timezone/0071.png" width="400" alt="YOU ARE IN THE MOUNTAINS. --------------- ENTER COMMAND?SOUTH YOU ARE IN SOME LOVELY ROLLING HILLS.">

**Type `SOUTH`**

**72.**

<img src="images/timezone/0072.png" width="400" alt="YOU ARE IN SOME LOVELY ROLLING HILLS. --------------- ENTER COMMAND?SOUTH YOU ARE IN SOME GRASSLANDS.">

**Type `SOUTH`**

**73.**

<img src="images/timezone/0073.png" width="400" alt="YOU ARE IN SOME GRASSLANDS. --------------- ENTER COMMAND?SOUTH YOU ARE IN THE GRASSLANDS">

**Type `SOUTH`**

**74.**

<img src="images/timezone/0074.png" width="400" alt="YOU ARE IN THE GRASSLANDS --------------- ENTER COMMAND?SOUTH YOU ARE IN SOME GRASSLANDS">

**Type `SOUTH`**

**75.**

<img src="images/timezone/0075.png" width="400" alt="YOU ARE IN SOME GRASSLANDS --------------- ENTER COMMAND?SOUTH YOU ARE WANDERING IN SOME GRASS LANDS.">

**Type `EAST`**

**76.**

<img src="images/timezone/0076.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE AT THE EDGE OF A CLIMBABLE CLIFF.">

**Type `DOWN`**

**77.**

<img src="images/timezone/0077.png" width="400" alt="YOU ARE STANDING IN A ROCKY CANYON. CLIMBABLE CLIFFS LEAD UP TOWARDS THE NORTH AND WEST.">

**Type `SOUTH`**

**78.**

<img src="images/timezone/0078-1.png" width="400" alt="YOU ARE IN A NORTH/SOUTH CANYON. STEEP CLIFFS LEAD UPWARDS TOWARDS THE EAST AND WEST. THERE IS A TIME MACHINE HERE.IT APPEARS TO BE PULSATING."><br>
<img src="images/timezone/0078.png" width="400" alt="CLIFFS LEAD UPWARDS TOWARDS THE EAST AND WEST. THERE IS A TIME MACHINE HERE.IT APPEARS TO BE PULSATING.">

**Type `GO MACHINE`**

**79.**

<img src="images/timezone/0079-1.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0079.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `SIT`**

**80.**

<img src="images/timezone/0080-1.png" width="400" alt="O.K. THE CHAIR IS COMFORTABLE. AS YOU SIT DOWN, THE MACHINE SEEMS TO COME TO LIFE. YOU ARE INSIDE A TIME MACHINE. YOU HEAR"><br>
<img src="images/timezone/0080.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT. --------------- ENTER COMMAND">

**Type `SA`**

**81.**

<img src="images/timezone/0081-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0081.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `1000AD`**

**82.**

<img src="images/timezone/0082-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0082.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `PULL LEVER`**

**83.**

<img src="images/timezone/0083-1.png" width="400" alt="THE MACHINE STARTS VIBRATING VIOLENTLY AND YOU FEEL DIZZY, ALMOST TO THE POINT OF PASSING OUT. SUDDENLY YOU FEEL A TERRIBLE JOLT AND THE MACHINE IS STILL."><br>
<img src="images/timezone/0083-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0083.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EXIT MACHINE`**

### South America, 1000 AD

**84.**

<img src="images/timezone/0084-1.png" width="400" alt="NE PLEASE INSERT DISK NUMBER 5I AND PRESS RETURN."><br>
<img src="images/timezone/0084.png" width="400" alt="YOU ARE ON A PLATEAU OF A MOUNTAIN IN THE ANDES. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

*Side 5I: drop `Time Zone (4am and san inc crack) disk I.dsk` on drive 1, and press Return.*

**Type `SOUTH`**

**85.**

<img src="images/timezone/0085.png" width="400" alt="YOU ARE AT THE EDGE OF A STEEP MOUNTAINSIDE THAT HAS BEEN TERRACED BY SOMEONE.">

**Type `DOWN`**

**86.**

<img src="images/timezone/0086.png" width="400" alt="--------------- ENTER COMMAND?DOWN YOU ARE ON A TERRACED MOUNTAINSIDE. POTATOES ARE GROWING HERE.">

**Type `DOWN`**

**87.**

<img src="images/timezone/0087.png" width="400" alt="--------------- ENTER COMMAND?DOWN YOU ARE LOWER DOWN THE TERRACED MOUNTAINSIDE. CORN IS GROWING HERE.">

**Type `DOWN`**

**88.**

<img src="images/timezone/0088.png" width="400" alt="MOUNTAINSIDE. CORN IS GROWING HERE. --------------- ENTER COMMAND?DOWN YOU ARE IN A GREEN VALLEY.">

**Type `SOUTH`**

**89.**

<img src="images/timezone/0089.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN A MEADOW. THERE IS A STREAM RUSHING BY HERE.">

**Type `SOUTH`**

**90.**

<img src="images/timezone/0090.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU SEE AN INCA CITY IN THE DISTANCE TO THE EAST.">

**Type `SOUTH`**

**91.**

<img src="images/timezone/0091.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE AT THE BOTTOM OF AN INCA PYRAMID.">

**Type `DROP HAMMER`**

**92.**

<img src="images/timezone/0092.png" width="400" alt="R YOU ARE AT THE BOTTOM OF AN INCA PYRAMID.">

**Type `UP`**

**93.**

<img src="images/timezone/0093.png" width="400" alt="--------------- ENTER COMMAND?UP YOU ARE AT THE TOP OF A PYRAMID. THERE IS A DOOR HERE.">

**Type `OPEN DOOR`**

**94.**

<img src="images/timezone/0094.png" width="400" alt="O.K. YOU ARE AT THE TOP OF A PYRAMID. THERE IS A DOOR HERE.">

**Type `GO DOOR`**

**95.**

<img src="images/timezone/0095.png" width="400" alt="YOU ARE INSIDE THE INCA PYRAMID. THERE IS A HOLE IN THE FLOOR, WITH A LADDER LEADING DOWN.">

**Type `DOWN`**

**96.**

<img src="images/timezone/0096-1.png" width="400" alt="YOU ARE AT THE BOTTOM OF THE LADDER. THERE IS A SMALL PASSAGE GOING SOUTH. THERE IS AN EERIE LIGHT COMING FROM UP THE LADDER."><br>
<img src="images/timezone/0096.png" width="400" alt="THERE IS A SMALL PASSAGE GOING SOUTH. THERE IS AN EERIE LIGHT COMING FROM UP THE LADDER.">

**Type `SOUTH`**

**97.**

<img src="images/timezone/0097-1.png" width="400" alt="THERE IS A TORCH HERE. YOU ARE IN AN SMALL TOMB. THERE IS AN EERIE LIGHT COMING FROM THE DOORWAY TO THE NORTH."><br>
<img src="images/timezone/0097.png" width="400" alt="YOU ARE IN AN SMALL TOMB. THERE IS AN EERIE LIGHT COMING FROM THE DOORWAY TO THE NORTH.">

**Type `GET TORCH`**

**98.**

<img src="images/timezone/0098.png" width="400" alt="YOU ARE IN AN SMALL TOMB. THERE IS AN EERIE LIGHT COMING FROM THE DOORWAY TO THE NORTH.">

**Type `NORTH`**

**99.**

<img src="images/timezone/0099-1.png" width="400" alt="YOU ARE AT THE BOTTOM OF THE LADDER. THERE IS A SMALL PASSAGE GOING SOUTH. THERE IS AN EERIE LIGHT COMING FROM UP THE LADDER."><br>
<img src="images/timezone/0099.png" width="400" alt="THERE IS A SMALL PASSAGE GOING SOUTH. THERE IS AN EERIE LIGHT COMING FROM UP THE LADDER.">

**Type `UP`**

**100.**

<img src="images/timezone/0100.png" width="400" alt="YOU ARE INSIDE THE INCA PYRAMID. THERE IS A HOLE IN THE FLOOR, WITH A LADDER LEADING DOWN.">

**Type `GO DOOR`**

**101.**

<img src="images/timezone/0101.png" width="400" alt="--------------- ENTER COMMAND?GO DOOR YOU ARE AT THE TOP OF A PYRAMID. THERE IS A DOOR HERE.">

**Type `THROW TORCH`**

**102.**

<img src="images/timezone/0102-1.png" width="400" alt="YOU THROW THE TORCH DOWN TO THE BOTTOM OF THE PYRAMID. AS IT FALLS, IT GOES OUT. YOU ARE AT THE TOP OF A PYRAMID. THERE"><br>
<img src="images/timezone/0102.png" width="400" alt="OUT. YOU ARE AT THE TOP OF A PYRAMID. THERE IS A DOOR HERE.">

**Type `DOWN`**

**103.**

<img src="images/timezone/0103.png" width="400" alt="THERE IS A TORCH HERE. YOU ARE AT THE BOTTOM OF AN INCA PYRAMID.">

**Type `GET TORCH`**

**104.**

<img src="images/timezone/0104.png" width="400" alt="--------------- ENTER COMMAND?GET TORCH YOU ARE AT THE BOTTOM OF AN INCA PYRAMID.">

**Type `GET HAMMER`**

**105.**

<img src="images/timezone/0105.png" width="400" alt="YOU ARE AT THE BOTTOM OF AN INCA PYRAMID.">

**Type `NORTH`**

**106.**

<img src="images/timezone/0106.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU SEE AN INCA CITY IN THE DISTANCE TO THE EAST.">

**Type `NORTH`**

**107.**

<img src="images/timezone/0107.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN A MEADOW. THERE IS A STREAM RUSHING BY HERE.">

**Type `NORTH`**

**108.**

<img src="images/timezone/0108.png" width="400" alt="RUSHING BY HERE. --------------- ENTER COMMAND?NORTH YOU ARE IN A GREEN VALLEY.">

**Type `UP`**

**109.**

<img src="images/timezone/0109.png" width="400" alt="--------------- ENTER COMMAND?UP YOU ARE LOWER DOWN THE TERRACED MOUNTAINSIDE. CORN IS GROWING HERE.">

**Type `UP`**

**110.**

<img src="images/timezone/0110.png" width="400" alt="--------------- ENTER COMMAND?UP YOU ARE ON A TERRACED MOUNTAINSIDE. POTATOES ARE GROWING HERE.">

**Type `UP`**

**111.**

<img src="images/timezone/0111.png" width="400" alt="YOU ARE AT THE EDGE OF A STEEP MOUNTAINSIDE THAT HAS BEEN TERRACED BY SOMEONE.">

**Type `NORTH`**

**112.**

<img src="images/timezone/0112.png" width="400" alt="YOU ARE ON A PLATEAU OF A MOUNTAIN IN THE ANDES. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `GO MACHINE`**

**113.**

<img src="images/timezone/0113-1.png" width="400" alt="PLEASE INSERT DISK NUMBER 1B AND PRESS RETURN."><br>
<img src="images/timezone/0113-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0113.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

*Side 1B: drop `Time Zone (4am and san inc crack) disk B.dsk` on drive 1, and press Return.*

**Type `SIT`**

**114.**

<img src="images/timezone/0114-1.png" width="400" alt="O.K. THE CHAIR IS COMFORTABLE. AS YOU SIT DOWN, THE MACHINE SEEMS TO COME TO LIFE. YOU ARE INSIDE A TIME MACHINE. YOU HEAR"><br>
<img src="images/timezone/0114.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT. --------------- ENTER COMMAND">

**Type `ASIA`**

**115.**

<img src="images/timezone/0115-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0115.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `50BC`**

**116.**

<img src="images/timezone/0116-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0116.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `PULL LEVER`**

**117.**

<img src="images/timezone/0117-1.png" width="400" alt="THE MACHINE STARTS VIBRATING VIOLENTLY AND YOU FEEL DIZZY, ALMOST TO THE POINT OF PASSING OUT. SUDDENLY YOU FEEL A TERRIBLE JOLT AND THE MACHINE IS STILL."><br>
<img src="images/timezone/0117-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0117.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EXIT MACHINE`**

### Asia, 50 BC

**118.**

<img src="images/timezone/0118-1.png" width="400" alt="NE PLEASE INSERT DISK NUMBER 4G AND PRESS RETURN."><br>
<img src="images/timezone/0118.png" width="400" alt="YOU ARE IN A SOGGY RICE PADDY.THERE IS A TIME MACHINE HERE IT APPEARS TO BE PULSATING.">

*Side 4G: drop `Time Zone (4am and san inc crack) disk G.dsk` on drive 1, and press Return.*

**Type `EAST`**

**119.**

<img src="images/timezone/0119.png" width="400" alt="PULSATING. --------------- ENTER COMMAND?EAST YOU ARE IN A SOGGY RICE PADDY.">

**Type `EAST`**

**120.**

<img src="images/timezone/0120.png" width="400" alt="YOU ARE IN A SOGGY RICE PADDY. --------------- ENTER COMMAND?EAST YOU ARE IN A LUSH MEADOW.">

**Type `EAST`**

**121.**

<img src="images/timezone/0121.png" width="400" alt="YOU ARE AT THE WEST EDGE OF THE YANGTZE RIVER. THE RIVER IS RUNNING NORTH AND SOUTH.">

**Type `NORTH`**

**122.**

<img src="images/timezone/0122.png" width="400" alt="YOU ARE AT THE WEST EDGE OF THE YANGTZE RIVER. THE RIVER IS RUNNING NORTH AND SOUTH.">

**Type `NORTH`**

**123.**

<img src="images/timezone/0123-1.png" width="400" alt="THERE IS A LONG POLE HERE. YOU ARE AT THE WEST EDGE OF THE YANGTZE RIVER. THE RIVER IS RUNNING NORTH AND SOUTH"><br>
<img src="images/timezone/0123.png" width="400" alt="YOU ARE AT THE WEST EDGE OF THE YANGTZE RIVER. THE RIVER IS RUNNING NORTH AND SOUTH">

**Type `GET POLE`**

**124.**

<img src="images/timezone/0124.png" width="400" alt="YOU ARE AT THE WEST EDGE OF THE YANGTZE RIVER. THE RIVER IS RUNNING NORTH AND SOUTH">

**Type `SOUTH`**

**125.**

<img src="images/timezone/0125.png" width="400" alt="YOU ARE AT THE WEST EDGE OF THE YANGTZE RIVER. THE RIVER IS RUNNING NORTH AND SOUTH.">

**Type `SOUTH`**

**126.**

<img src="images/timezone/0126.png" width="400" alt="YOU ARE AT THE WEST EDGE OF THE YANGTZE RIVER. THE RIVER IS RUNNING NORTH AND SOUTH.">

**Type `SOUTH`**

**127.**

<img src="images/timezone/0127.png" width="400" alt="YOU ARE AT THE WEST EDGE OF THE YANGTZE RIVER. THE RIVER IS RUNNING NORTH AND SOUTH.">

**Type `SOUTH`**

**128.**

<img src="images/timezone/0128.png" width="400" alt="YOU ARE AT THE WEST EDGE OF THE YANGTZE RIVER. THERE IS A SMALL CHINESE JUNK (BOAT) HERE.">

**Type `GO BOAT`**

**129.**

<img src="images/timezone/0129.png" width="400" alt="(BOAT) HERE. --------------- ENTER COMMAND?GO BOAT YOU ARE IN THE CHINESE JUNK (BOAT).">

**Type `EAST`**

**130.**

<img src="images/timezone/0130-1.png" width="400" alt="USING THE POLE, YOU GUIDE THE BOAT ACROSS THE RIVER, THROUGH THE TRICKY CURRENTS. YOU ARE IN A CHINESE JUNK (BOAT)."><br>
<img src="images/timezone/0130.png" width="400" alt="ACROSS THE RIVER, THROUGH THE TRICKY CURRENTS. YOU ARE IN A CHINESE JUNK (BOAT).">

**Type `EXIT BOAT`**

**131.**

<img src="images/timezone/0131.png" width="400" alt="YOU ARE AT THE EAST EDGE OF THE YANGTZE RIVER. THERE IS A SMALL CHINESE JUNK (BOAT) HERE.">

**Type `NORTH`**

**132.**

<img src="images/timezone/0132.png" width="400" alt="YOU ARE ON THE EAST EDGE OF THE YANGTZE RIVER. THE RIVER IS RUNNING NORTH AND SOUTH.">

**Type `NORTH`**

**133.**

<img src="images/timezone/0133.png" width="400" alt="YOU ARE ON THE EAST EDGE OF THE YANGTZE RIVER. THE RIVER IS RUNNING NORTH AND SOUTH.">

**Type `NORTH`**

**134.**

<img src="images/timezone/0134.png" width="400" alt="YOU ARE IN A SMALL CHINESE VILLAGE.THERE IS A ROAD GOING NORTH AND SOUTH.">

**Type `NORTH`**

**135.**

<img src="images/timezone/0135.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN A FIELD. A ROAD GOES EAST AND SOUTH.">

**Type `EAST`**

**136.**

<img src="images/timezone/0136.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE OUTSIDE A BUDDHIST TEMPLE. A ROAD GOES WEST FROM HERE.">

**Type `NORTH`**

**137.**

<img src="images/timezone/0137-1.png" width="400" alt="YOU ARE INSIDE THE BUDDHIST TEMPLE.YOU SEE A STATUE OF BUDDHA WITH AN EMERALD IN HIS NAVEL. THERE ARE DOORWAYS TO THE WEST AND SOUTH."><br>
<img src="images/timezone/0137.png" width="400" alt="SEE A STATUE OF BUDDHA WITH AN EMERALD IN HIS NAVEL. THERE ARE DOORWAYS TO THE WEST AND SOUTH.">

**Type `WEST`**

**138.**

<img src="images/timezone/0138.png" width="400" alt="THERE IS A SHOVEL HERE. YOU ARE IN A CHINESE ROCK GARDEN. THE GARDEN IS ALL FENCED IN.">

**Type `GET SHOVEL`**

**139.**

<img src="images/timezone/0139.png" width="400" alt="YOU ARE IN A CHINESE ROCK GARDEN. THE GARDEN IS ALL FENCED IN.">

**Type `DIG`**

**140.**

<img src="images/timezone/0140-1.png" width="400" alt="AFTER DIGGING IN THE DIRT FOR A WHILE, YOU UNCOVER A BEAUTIFUL JADE STONE! YOU ARE IN A CHINESE ROCK GARDEN. THE GARDEN IS ALL FENCED IN."><br>
<img src="images/timezone/0140.png" width="400" alt="YOU UNCOVER A BEAUTIFUL JADE STONE! YOU ARE IN A CHINESE ROCK GARDEN. THE GARDEN IS ALL FENCED IN.">

**Type `GET JADE`**

**141.**

<img src="images/timezone/0141.png" width="400" alt="--------------- ENTER COMMAND?GET JADE YOU ARE IN A CHINESE ROCK GARDEN. THE GARDEN IS ALL FENCED IN.">

**Type `DIG`**

**142.**

<img src="images/timezone/0142-1.png" width="400" alt="AFTER DIGGING IN THE DIRT FOR A WHILE, YOU UNCOVER A SECOND BEAUTIFUL JADE STONE. YOU ARE IN A CHINESE ROCK GARDEN. THE"><br>
<img src="images/timezone/0142.png" width="400" alt="STONE. YOU ARE IN A CHINESE ROCK GARDEN. THE GARDEN IS ALL FENCED IN.">

**Type `GET JADE`**

**143.**

<img src="images/timezone/0143.png" width="400" alt="--------------- ENTER COMMAND?GET JADE YOU ARE IN A CHINESE ROCK GARDEN. THE GARDEN IS ALL FENCED IN.">

**Type `DROP SHOVEL`**

**144.**

<img src="images/timezone/0144.png" width="400" alt="THERE IS A SHOVEL HERE. YOU ARE IN A CHINESE ROCK GARDEN. THE GARDEN IS ALL FENCED IN.">

**Type `EAST`**

**145.**

<img src="images/timezone/0145-1.png" width="400" alt="YOU ARE INSIDE THE BUDDHIST TEMPLE.YOU SEE A STATUE OF BUDDHA WITH AN EMERALD IN HIS NAVEL. THERE ARE DOORWAYS TO THE WEST AND SOUTH."><br>
<img src="images/timezone/0145.png" width="400" alt="SEE A STATUE OF BUDDHA WITH AN EMERALD IN HIS NAVEL. THERE ARE DOORWAYS TO THE WEST AND SOUTH.">

**Type `SOUTH`**

**146.**

<img src="images/timezone/0146.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE OUTSIDE A BUDDHIST TEMPLE. A ROAD GOES WEST FROM HERE.">

**Type `WEST`**

**147.**

<img src="images/timezone/0147.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN A FIELD. A ROAD GOES EAST AND SOUTH.">

**Type `SOUTH`**

**148.**

<img src="images/timezone/0148.png" width="400" alt="YOU ARE IN A SMALL CHINESE VILLAGE.THERE IS A ROAD GOING NORTH AND SOUTH.">

**Type `SOUTH`**

**149.**

<img src="images/timezone/0149.png" width="400" alt="YOU ARE ON THE EAST EDGE OF THE YANGTZE RIVER. THE RIVER IS RUNNING NORTH AND SOUTH.">

**Type `SOUTH`**

**150.**

<img src="images/timezone/0150.png" width="400" alt="YOU ARE ON THE EAST EDGE OF THE YANGTZE RIVER. THE RIVER IS RUNNING NORTH AND SOUTH.">

**Type `SOUTH`**

**151.**

<img src="images/timezone/0151.png" width="400" alt="YOU ARE AT THE EAST EDGE OF THE YANGTZE RIVER. THERE IS A SMALL CHINESE JUNK (BOAT) HERE.">

**Type `SOUTH`**

**152.**

<img src="images/timezone/0152.png" width="400" alt="YOU ARE AT THE EAST EDGE OF THE YANGTZE RIVER. THE RIVER IS RUNNING NORTH AND SOUTH.">

**Type `EAST`**

**153.**

<img src="images/timezone/0153.png" width="400" alt="SOUTH. --------------- ENTER COMMAND?EAST YOU ARE IN A WIDE FIELD.">

**Type `EAST`**

**154.**

<img src="images/timezone/0154-1.png" width="400" alt="THERE IS A BAG OF RICE HERE. THE PEASANT HAS A ROPE AROUND HIS WAIST. YOU ARE IN A SOGGY RICE PADDY. THERE IS"><br>
<img src="images/timezone/0154.png" width="400" alt="YOU ARE IN A SOGGY RICE PADDY. THERE IS A PEASANT HARVESTING RICE. HE HAS SEVERAL BAGS OF RICE NEAR HIM.">

**Type `BUY ROPE`**

**155.**

<img src="images/timezone/0155-1.png" width="400" alt="WITH WHAT? THERE IS A BAG OF RICE HERE. THE PEASANT HAS A ROPE AROUND HIS WAIST."><br>
<img src="images/timezone/0155.png" width="400" alt="YOU ARE IN A SOGGY RICE PADDY. THERE IS A PEASANT HARVESTING RICE. HE HAS SEVERAL BAGS OF RICE NEAR HIM.">

**Type `WITH JADE`**

**156.**

<img src="images/timezone/0156-1.png" width="400" alt="THE PEASANT IS VERY INTERESTED IN JADE.HE TAKES IT AND GIVES YOU THE ROPE. THERE IS A BAG OF RICE HERE."><br>
<img src="images/timezone/0156.png" width="400" alt="YOU ARE IN A SOGGY RICE PADDY. THERE IS A PEASANT HARVESTING RICE. HE HAS SEVERAL BAGS OF RICE NEAR HIM.">

**Type `BUY RICE`**

**157.**

<img src="images/timezone/0157-1.png" width="400" alt="WITH WHAT? THERE IS A BAG OF RICE HERE. YOU ARE IN A SOGGY RICE PADDY. THERE IS A PEASANT HARVESTING RICE. HE HAS"><br>
<img src="images/timezone/0157.png" width="400" alt="YOU ARE IN A SOGGY RICE PADDY. THERE IS A PEASANT HARVESTING RICE. HE HAS SEVERAL BAGS OF RICE NEAR HIM.">

**Type `WITH JADE`**

**158.**

<img src="images/timezone/0158-1.png" width="400" alt="THE PEASANT IS VERY INTERESTED IN JADE.HE TAKES IT AND GIVES YOU A BAG OF RICE. YOU ARE IN A SOGGY RICE PADDY. THERE IS"><br>
<img src="images/timezone/0158.png" width="400" alt="YOU ARE IN A SOGGY RICE PADDY. THERE IS A PEASANT HARVESTING RICE. HE HAS SEVERAL BAGS OF RICE NEAR HIM.">

**Type `WEST`**

**159.**

<img src="images/timezone/0159.png" width="400" alt="SEVERAL BAGS OF RICE NEAR HIM. --------------- ENTER COMMAND?WEST YOU ARE IN A WIDE FIELD.">

**Type `WEST`**

**160.**

<img src="images/timezone/0160.png" width="400" alt="YOU ARE AT THE EAST EDGE OF THE YANGTZE RIVER. THE RIVER IS RUNNING NORTH AND SOUTH.">

**Type `NORTH`**

**161.**

<img src="images/timezone/0161.png" width="400" alt="YOU ARE AT THE EAST EDGE OF THE YANGTZE RIVER. THERE IS A SMALL CHINESE JUNK (BOAT) HERE.">

**Type `GO BOAT`**

**162.**

<img src="images/timezone/0162.png" width="400" alt="(BOAT) HERE. --------------- ENTER COMMAND?GO BOAT YOU ARE IN A CHINESE JUNK (BOAT).">

**Type `WEST`**

**163.**

<img src="images/timezone/0163-1.png" width="400" alt="USING THE POLE, YOU GUIDE THE BOAT ACROSS THE RIVER, THROUGH THE TRICKY CURRENTS. YOU ARE IN THE CHINESE JUNK (BOAT)."><br>
<img src="images/timezone/0163.png" width="400" alt="ACROSS THE RIVER, THROUGH THE TRICKY CURRENTS. YOU ARE IN THE CHINESE JUNK (BOAT).">

**Type `EXIT BOAT`**

**164.**

<img src="images/timezone/0164.png" width="400" alt="YOU ARE AT THE WEST EDGE OF THE YANGTZE RIVER. THERE IS A SMALL CHINESE JUNK (BOAT) HERE.">

**Type `NORTH`**

**165.**

<img src="images/timezone/0165.png" width="400" alt="YOU ARE AT THE WEST EDGE OF THE YANGTZE RIVER. THE RIVER IS RUNNING NORTH AND SOUTH.">

**Type `NORTH`**

**166.**

<img src="images/timezone/0166.png" width="400" alt="YOU ARE AT THE WEST EDGE OF THE YANGTZE RIVER. THE RIVER IS RUNNING NORTH AND SOUTH.">

**Type `WEST`**

**167.**

<img src="images/timezone/0167.png" width="400" alt="SOUTH. --------------- ENTER COMMAND?WEST YOU ARE IN A LUSH MEADOW.">

**Type `WEST`**

**168.**

<img src="images/timezone/0168.png" width="400" alt="YOU ARE IN A LUSH MEADOW. --------------- ENTER COMMAND?WEST YOU ARE IN A SOGGY RICE PADDY.">

**Type `WEST`**

**169.**

<img src="images/timezone/0169.png" width="400" alt="YOU ARE IN A SOGGY RICE PADDY.THERE IS A TIME MACHINE HERE IT APPEARS TO BE PULSATING.">

**Type `GO MACHINE`**

**170.**

<img src="images/timezone/0170-1.png" width="400" alt="PLEASE INSERT DISK NUMBER 1B AND PRESS RETURN."><br>
<img src="images/timezone/0170-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0170.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

*Side 1B: drop `Time Zone (4am and san inc crack) disk B.dsk` on drive 1, and press Return.*

**Type `SIT`**

**171.**

<img src="images/timezone/0171-1.png" width="400" alt="O.K. THE CHAIR IS COMFORTABLE. AS YOU SIT DOWN, THE MACHINE SEEMS TO COME TO LIFE. YOU ARE INSIDE A TIME MACHINE. YOU HEAR"><br>
<img src="images/timezone/0171.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT. --------------- ENTER COMMAND">

**Type `EUROPE`**

**172.**

<img src="images/timezone/0172-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0172.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `1400AD`**

**173.**

<img src="images/timezone/0173-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0173.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `PULL LEVER`**

**174.**

<img src="images/timezone/0174-1.png" width="400" alt="THE MACHINE STARTS VIBRATING VIOLENTLY AND YOU FEEL DIZZY, ALMOST TO THE POINT OF PASSING OUT. SUDDENLY YOU FEEL A TERRIBLE JOLT AND THE MACHINE IS STILL."><br>
<img src="images/timezone/0174-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0174.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EXIT MACHINE`**

### Europe, 1400 AD

**175.**

<img src="images/timezone/0175-1.png" width="400" alt="NE PLEASE INSERT DISK NUMBER 2C AND PRESS RETURN."><br>
<img src="images/timezone/0175-2.png" width="400" alt="YOU ARE AT THE EDGE OF AN OCEAN, HILLS ARE AROUND YOU. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING."><br>
<img src="images/timezone/0175.png" width="400" alt="ARE AROUND YOU. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

*Side 2C: drop `Time Zone (4am and san inc crack) disk C.dsk` on drive 1, and press Return.*

**Type `SOUTH`**

**176.**

<img src="images/timezone/0176.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE AT THE EDGE OF THE OCEAN. HILLS SURROUND YOU.">

**Type `SOUTH`**

**177.**

<img src="images/timezone/0177.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN A LITTLE FISHING VILLAGE, ON THE EDGE OF THE OCEAN, CALLED PALOS.">

**Type `EAST`**

**178.**

<img src="images/timezone/0178.png" width="400" alt="THE EDGE OF THE OCEAN, CALLED PALOS. --------------- ENTER COMMAND?EAST YOU ARE STANDING IN FRONT OF A HOUSE.">

**Type `OPEN DOOR`**

**179.**

<img src="images/timezone/0179.png" width="400" alt="--------------- ENTER COMMAND?OPEN DOOR THE DOOR IS NOW OPEN. YOU ARE STANDING IN FRONT OF A HOUSE.">

**Type `GO DOOR`**

**180.**

<img src="images/timezone/0180.png" width="400" alt="--------------- ENTER COMMAND?GO DOOR YOU ARE INSIDE THE HOUSE. THERE IS A MAN SEATED BEHIND THE TABLE.">

**Type `SIGN UP`**

**181.**

<img src="images/timezone/0181-1.png" width="400" alt="THE MAN SIGNS YOU UP AS A CREW MEMBER FOR THE SANTA MARIA. WOULD YOU LIKE TO WORK IN THE GALLEY, CARGO HOLD, OR ON DECK HANDLING THE SAILS AND MASTS?"><br>
<img src="images/timezone/0181.png" width="400" alt="DECK HANDLING THE SAILS AND MASTS? YOU ARE INSIDE THE HOUSE. THERE IS A MAN SEATED BEHIND THE TABLE.">

**Type `ON DECK`**

**182.**

<img src="images/timezone/0182-1.png" width="400" alt="YOU ARE OFFICIALLY WORKING WITH THE SAILS AND MASTS ON DECK. THE MAN HANDS YOU A BOARDING PASS. YOU ARE INSIDE THE HOUSE. THERE IS A"><br>
<img src="images/timezone/0182.png" width="400" alt="YOU A BOARDING PASS. YOU ARE INSIDE THE HOUSE. THERE IS A MAN SEATED BEHIND THE TABLE.">

**Type `SOUTH`**

**183.**

<img src="images/timezone/0183.png" width="400" alt="MAN SEATED BEHIND THE TABLE. --------------- ENTER COMMAND?SOUTH YOU ARE STANDING IN FRONT OF A HOUSE.">

**Type `WEST`**

**184.**

<img src="images/timezone/0184.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN A LITTLE FISHING VILLAGE, ON THE EDGE OF THE OCEAN, CALLED PALOS.">

**Type `SOUTH`**

**185.**

<img src="images/timezone/0185.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE AT THE EDGE OF AN OCEAN, THERE IS A PIER TO THE WEST.">

**Type `WEST`**

**186.**

<img src="images/timezone/0186.png" width="400" alt="IS A PIER TO THE WEST. --------------- ENTER COMMAND?WEST YOU ARE AT THE END OF THE PIER.">

**Type `SHOW PASS`**

**187.**

<img src="images/timezone/0187-1.png" width="400" alt="THE SAILOR LOOKS AT YOUR BOARDING PASS AND TELLS YOU TO GO TO THE DECK AT THE FRONT OF THE SHIP. THE SANTA MARIA IS DOCKED. THERE IS A"><br>
<img src="images/timezone/0187.png" width="400" alt="THE SANTA MARIA IS DOCKED. THERE IS A SAILOR HERE. YOU ARE AT THE END OF THE PIER.">

**Type `NORTH`**

**188.**

<img src="images/timezone/0188-1.png" width="400" alt="YOU ARE ON DECK AT THE STERN. CHRISTOPHER COLUMBUS IS STANDING HERE. THERE IS A DOOR LEADING NORTH TO THE GALLEY."><br>
<img src="images/timezone/0188.png" width="400" alt="CHRISTOPHER COLUMBUS IS STANDING HERE. THERE IS A DOOR LEADING NORTH TO THE GALLEY.">

**Type `WEST`**

**189.**

<img src="images/timezone/0189.png" width="400" alt="YOU ARE AMID SHIP. THERE IS A HOLE WITH A LADDER LEADING DOWN. A SAILOR IS STANDING HERE.">

**Type `WEST`**

**190.**

<img src="images/timezone/0190.png" width="400" alt="YOU ARE ON DECK AT THE BOW OF THE SHIP. THERE IS A VERY TALL MAST HERE WITH A ROPE LADDER ATTACHED TO IT.">

**Type `UP`**

**191.**

<img src="images/timezone/0191.png" width="400" alt="YOU ARE IN A CROWSNEST AT THE TOP OF THE MAST. THERE IS A PARROT SITTING ON TOP OF A TELESCOPE HERE.">

**Type `LOOK TELESCOPE`**

**192.**

<img src="images/timezone/0192-1.png" width="400" alt="YOU SEE A FARM HOUSE IN THE DISTANCE TO THE SOUTHEAST. THE FARM HOUSE IS SURROUNDED BY HILLS. YOU ARE IN A CROWSNEST AT THE TOP OF"><br>
<img src="images/timezone/0192.png" width="400" alt="YOU ARE IN A CROWSNEST AT THE TOP OF THE MAST. THERE IS A PARROT SITTING ON TOP OF A TELESCOPE HERE.">

**Type `DOWN`**

**193.**

<img src="images/timezone/0193.png" width="400" alt="YOU ARE ON DECK AT THE BOW OF THE SHIP. THERE IS A VERY TALL MAST HERE WITH A ROPE LADDER ATTACHED TO IT.">

**Type `EAST`**

**194.**

<img src="images/timezone/0194.png" width="400" alt="YOU ARE AMID SHIP. THERE IS A HOLE WITH A LADDER LEADING DOWN. A SAILOR IS STANDING HERE.">

**Type `EAST`**

**195.**

<img src="images/timezone/0195-1.png" width="400" alt="YOU ARE ON DECK AT THE STERN. CHRISTOPHER COLUMBUS IS STANDING HERE. THERE IS A DOOR LEADING NORTH TO THE GALLEY."><br>
<img src="images/timezone/0195.png" width="400" alt="CHRISTOPHER COLUMBUS IS STANDING HERE. THERE IS A DOOR LEADING NORTH TO THE GALLEY.">

**Type `SOUTH`**

**196.**

<img src="images/timezone/0196.png" width="400" alt="--------------- ENTER COMMAND?SOUTH THE SAILOR SAYS &#34;SEE YOU AROUND!&#34; YOU ARE AT THE END OF THE PIER.">

**Type `EAST`**

**197.**

<img src="images/timezone/0197.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE AT THE EDGE OF AN OCEAN, THERE IS A PIER TO THE WEST.">

**Type `EAST`**

**198.**

<img src="images/timezone/0198.png" width="400" alt="IS A PIER TO THE WEST. --------------- ENTER COMMAND?EAST YOU ARE WANDERING AMONG THE HILLS.">

**Type `SOUTH`**

**199.**

<img src="images/timezone/0199.png" width="400" alt="YOU ARE WANDERING AMONG THE HILLS. --------------- ENTER COMMAND?SOUTH YOU ARE IN THE HILLS.">

**Type `EAST`**

**200.**

<img src="images/timezone/0200.png" width="400" alt="YOU ARE IN THE HILLS. --------------- ENTER COMMAND?EAST YOU ARE AMONG THE ROLLING HILLS.">

**Type `SOUTH`**

**201.**

<img src="images/timezone/0201.png" width="400" alt="YOU ARE AMONG THE ROLLING HILLS. --------------- ENTER COMMAND?SOUTH YOU ARE IN FRONT OF A FARM HOUSE.">

**Type `OPEN DOOR`**

**202.**

<img src="images/timezone/0202.png" width="400" alt="YOU ARE IN FRONT OF A FARM HOUSE. --------------- ENTER COMMAND?OPEN DOOR YOU ARE IN FRONT OF A FARM HOUSE.">

**Type `GO DOOR`**

**203.**

<img src="images/timezone/0203.png" width="400" alt="--------------- ENTER COMMAND?GO DOOR THERE IS AN IRON BAR HERE. YOU ARE INSIDE THE FARM HOUSE.">

**Type `GET BAR`**

**204.**

<img src="images/timezone/0204.png" width="400" alt="YOU ARE INSIDE THE FARM HOUSE. --------------- ENTER COMMAND?GET BAR YOU ARE INSIDE THE FARM HOUSE.">

**Type `WEST`**

**205.**

<img src="images/timezone/0205.png" width="400" alt="YOU ARE INSIDE THE FARM HOUSE. --------------- ENTER COMMAND?WEST YOU ARE IN FRONT OF A FARM HOUSE.">

**Type `NORTH`**

**206.**

<img src="images/timezone/0206.png" width="400" alt="YOU ARE IN FRONT OF A FARM HOUSE. --------------- ENTER COMMAND?NORTH YOU ARE AMONG THE ROLLING HILLS.">

**Type `WEST`**

**207.**

<img src="images/timezone/0207.png" width="400" alt="YOU ARE AMONG THE ROLLING HILLS. --------------- ENTER COMMAND?WEST YOU ARE IN THE HILLS.">

**Type `NORTH`**

**208.**

<img src="images/timezone/0208.png" width="400" alt="YOU ARE IN THE HILLS. --------------- ENTER COMMAND?NORTH YOU ARE WANDERING AMONG THE HILLS.">

**Type `WEST`**

**209.**

<img src="images/timezone/0209.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE AT THE EDGE OF AN OCEAN, THERE IS A PIER TO THE WEST.">

**Type `NORTH`**

**210.**

<img src="images/timezone/0210.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN A LITTLE FISHING VILLAGE, ON THE EDGE OF THE OCEAN, CALLED PALOS.">

**Type `NORTH`**

**211.**

<img src="images/timezone/0211.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE AT THE EDGE OF THE OCEAN. HILLS SURROUND YOU.">

**Type `NORTH`**

**212.**

<img src="images/timezone/0212-1.png" width="400" alt="YOU ARE AT THE EDGE OF AN OCEAN, HILLS ARE AROUND YOU. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING."><br>
<img src="images/timezone/0212.png" width="400" alt="ARE AROUND YOU. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `DROP PASS`**

**213.**

<img src="images/timezone/0213-1.png" width="400" alt="YOU ARE AT THE EDGE OF AN OCEAN, HILLS ARE AROUND YOU. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING."><br>
<img src="images/timezone/0213.png" width="400" alt="ARE AROUND YOU. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `GO MACHINE`**

**214.**

<img src="images/timezone/0214-1.png" width="400" alt="PLEASE INSERT DISK NUMBER 1B AND PRESS RETURN."><br>
<img src="images/timezone/0214-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0214.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

*Side 1B: drop `Time Zone (4am and san inc crack) disk B.dsk` on drive 1, and press Return.*

**Type `SIT`**

**215.**

<img src="images/timezone/0215-1.png" width="400" alt="O.K. THE CHAIR IS COMFORTABLE. AS YOU SIT DOWN, THE MACHINE SEEMS TO COME TO LIFE. YOU ARE INSIDE A TIME MACHINE. YOU HEAR"><br>
<img src="images/timezone/0215.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT. --------------- ENTER COMMAND">

**Type `AUSTRALIA`**

**216.**

<img src="images/timezone/0216-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0216.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `1700AD`**

**217.**

<img src="images/timezone/0217-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0217.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `PULL LEVER`**

**218.**

<img src="images/timezone/0218-1.png" width="400" alt="THE MACHINE STARTS VIBRATING VIOLENTLY AND YOU FEEL DIZZY, ALMOST TO THE POINT OF PASSING OUT. SUDDENLY YOU FEEL A TERRIBLE JOLT AND THE MACHINE IS STILL."><br>
<img src="images/timezone/0218-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0218.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EXIT MACHINE`**

### Australia, 1700 AD

**219.**

<img src="images/timezone/0219-1.png" width="400" alt="NE PLEASE INSERT DISK NUMBER 4H AND PRESS RETURN."><br>
<img src="images/timezone/0219.png" width="400" alt="YOU ARE IN A MEADOW. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

*Side 4H: drop `Time Zone (4am and san inc crack) disk H.dsk` on drive 1, and press Return.*

**Type `NORTH`**

**220.**

<img src="images/timezone/0220.png" width="400" alt="PULSATING. --------------- ENTER COMMAND?NORTH YOU ARE IN A MEADOW.">

**Type `NORTH`**

**221.**

<img src="images/timezone/0221.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN GRAZING LAND. THERE ARE SHEEP GRAZING HERE.">

**Type `NORTH`**

**222.**

<img src="images/timezone/0222.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN GRAZING LAND. THERE ARE SOME SHEEP HERE.">

**Type `NORTH`**

**223.**

<img src="images/timezone/0223.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN GRAZING LAND. THERE ARE SOME SHEEP HERE.">

**Type `WEST`**

**224.**

<img src="images/timezone/0224.png" width="400" alt="--------------- ENTER COMMAND?WEST THERE IS A PADLOCK HERE. YOU ARE IN FRONT OF AN OLD BARN.">

**Type `BREAK PADLOCK`**

**225.**

<img src="images/timezone/0225.png" width="400" alt="WITH WHAT? THERE IS A PADLOCK HERE. YOU ARE IN FRONT OF AN OLD BARN.">

**Type `WITH BAR`**

**226.**

<img src="images/timezone/0226-1.png" width="400" alt="USING THE IRON BAR, YOU BREAK THE PADLOCK. IT FALLS TO THE GROUND. THERE IS A PADLOCK HERE. YOU ARE IN FRONT OF AN OLD BARN."><br>
<img src="images/timezone/0226.png" width="400" alt="PADLOCK. IT FALLS TO THE GROUND. THERE IS A PADLOCK HERE. YOU ARE IN FRONT OF AN OLD BARN.">

**Type `OPEN DOOR`**

**227.**

<img src="images/timezone/0227.png" width="400" alt="O.K. THERE IS A PADLOCK HERE. YOU ARE IN FRONT OF AN OLD BARN.">

**Type `NORTH`**

**228.**

<img src="images/timezone/0228.png" width="400" alt="THERE IS A SAW HERE. YOU ARE INSIDE THE BARN. THE BARN DOOR IS TO THE SOUTH.">

**Type `GET SAW`**

**229.**

<img src="images/timezone/0229.png" width="400" alt="--------------- ENTER COMMAND?GET SAW YOU ARE INSIDE THE BARN. THE BARN DOOR IS TO THE SOUTH.">

**Type `SOUTH`**

**230.**

<img src="images/timezone/0230.png" width="400" alt="--------------- ENTER COMMAND?SOUTH THERE IS A PADLOCK HERE. YOU ARE IN FRONT OF AN OLD BARN.">

**Type `EAST`**

**231.**

<img src="images/timezone/0231.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN GRAZING LAND. THERE ARE SOME SHEEP HERE.">

**Type `SOUTH`**

**232.**

<img src="images/timezone/0232.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN GRAZING LAND. THERE ARE SOME SHEEP HERE.">

**Type `SOUTH`**

**233.**

<img src="images/timezone/0233.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN GRAZING LAND. THERE ARE SHEEP GRAZING HERE.">

**Type `SOUTH`**

**234.**

<img src="images/timezone/0234.png" width="400" alt="SHEEP GRAZING HERE. --------------- ENTER COMMAND?SOUTH YOU ARE IN A MEADOW.">

**Type `SOUTH`**

**235.**

<img src="images/timezone/0235.png" width="400" alt="YOU ARE IN A MEADOW. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `GO MACHINE`**

**236.**

<img src="images/timezone/0236-1.png" width="400" alt="PLEASE INSERT DISK NUMBER 1B AND PRESS RETURN."><br>
<img src="images/timezone/0236-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0236.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

*Side 1B: drop `Time Zone (4am and san inc crack) disk B.dsk` on drive 1, and press Return.*

**Type `SIT`**

**237.**

<img src="images/timezone/0237-1.png" width="400" alt="O.K. THE CHAIR IS COMFORTABLE. AS YOU SIT DOWN, THE MACHINE SEEMS TO COME TO LIFE. YOU ARE INSIDE A TIME MACHINE. YOU HEAR"><br>
<img src="images/timezone/0237.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT. --------------- ENTER COMMAND">

**Type `NA`**

**238.**

<img src="images/timezone/0238-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0238.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `1700AD`**

**239.**

<img src="images/timezone/0239-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0239.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `PULL LEVER`**

**240.**

<img src="images/timezone/0240-1.png" width="400" alt="THE MACHINE STARTS VIBRATING VIOLENTLY AND YOU FEEL DIZZY, ALMOST TO THE POINT OF PASSING OUT. SUDDENLY YOU FEEL A TERRIBLE JOLT AND THE MACHINE IS STILL."><br>
<img src="images/timezone/0240-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0240.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EXIT MACHINE`**

### North America, 1700 AD

**241.**

<img src="images/timezone/0241-1.png" width="400" alt="NE PLEASE INSERT DISK NUMBER 5J AND PRESS RETURN."><br>
<img src="images/timezone/0241.png" width="400" alt="YOU ARE IN THE WOODS. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

*Side 5J: drop `Time Zone (4am and san inc crack) disk J.dsk` on drive 1, and press Return.*

**Type `NORTH`**

**242.**

<img src="images/timezone/0242.png" width="400" alt="PULSATING. --------------- ENTER COMMAND?NORTH YOU ARE IN A PASTURE.">

**Type `NORTH`**

**243.**

<img src="images/timezone/0243.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN A PASTURE. A DIRT ROAD LEADS EAST AND WEST.">

**Type `NORTH`**

**244.**

<img src="images/timezone/0244.png" width="400" alt="LEADS EAST AND WEST. --------------- ENTER COMMAND?NORTH YOU ARE IN FRONT OF THE COURTHOUSE.">

**Type `EAST`**

**245.**

<img src="images/timezone/0245.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN PHILADELPHIA. THE DIRT ROAD GOES NORTH, SOUTH, EAST, AND WEST.">

**Type `EAST`**

**246.**

<img src="images/timezone/0246.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN FRONT OF BENJAMIN FRANKLIN&#39;S PRINT SHOP.">

**Type `BREAK WINDOW`**

**247.**

<img src="images/timezone/0247.png" width="400" alt="OK. YOU ARE IN FRONT OF BENJAMIN FRANKLIN&#39;S PRINT SHOP.">

**Type `GO WINDOW`**

**248.**

<img src="images/timezone/0248.png" width="400" alt="--------------- ENTER COMMAND?GO WINDOW YOU ARE IN THE FRONT OFFICE OF BENJAMIN FRANKLIN&#39;S PRINT SHOP.">

**Type `NORTH`**

**249.**

<img src="images/timezone/0249-1.png" width="400" alt="THERE IS A KITE HERE. YOU ARE IN THE BACKROOM OF BENJAMIN FRANKLIN&#39;S PRINT SHOP. A WOODEN TRUNK IS IN THE CORNER. THE DOORWAY IS TO"><br>
<img src="images/timezone/0249.png" width="400" alt="FRANKLIN&#39;S PRINT SHOP. A WOODEN TRUNK IS IN THE CORNER. THE DOORWAY IS TO THE SOUTH.">

**Type `USE SAW`**

**250.**

<img src="images/timezone/0250-1.png" width="400" alt="USING THE SAW, YOU CUT THROUGH THE TOP HALF OF THE TRUNK AND OPEN IT. THERE IS A KITE HERE. YOU ARE IN THE BACKROOM OF BENJAMIN"><br>
<img src="images/timezone/0250.png" width="400" alt="FRANKLIN&#39;S PRINT SHOP. A WOODEN TRUNK IS IN THE CORNER. THE DOORWAY IS TO THE SOUTH.">

**Type `LOOK CHEST`**

**251.**

<img src="images/timezone/0251-1.png" width="400" alt="THERE IS A SKELETON KEY HERE. YOU ARE IN THE BACKROOM OF BENJAMIN FRANKLIN&#39;S PRINT SHOP. A WOODEN TRUNK IS IN THE CORNER. THE DOORWAY IS TO"><br>
<img src="images/timezone/0251.png" width="400" alt="FRANKLIN&#39;S PRINT SHOP. A WOODEN TRUNK IS IN THE CORNER. THE DOORWAY IS TO THE SOUTH.">

**Type `GET KEY`**

**252.**

<img src="images/timezone/0252-1.png" width="400" alt="YOU ARE IN THE BACKROOM OF BENJAMIN FRANKLIN&#39;S PRINT SHOP. A WOODEN TRUNK IS IN THE CORNER. THE DOORWAY IS TO THE SOUTH."><br>
<img src="images/timezone/0252.png" width="400" alt="FRANKLIN&#39;S PRINT SHOP. A WOODEN TRUNK IS IN THE CORNER. THE DOORWAY IS TO THE SOUTH.">

**Type `SOUTH`**

**253.**

<img src="images/timezone/0253.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN THE FRONT OFFICE OF BENJAMIN FRANKLIN&#39;S PRINT SHOP.">

**Type `GO WINDOW`**

**254.**

<img src="images/timezone/0254.png" width="400" alt="--------------- ENTER COMMAND?GO WINDOW YOU ARE IN FRONT OF BENJAMIN FRANKLIN&#39;S PRINT SHOP.">

**Type `WEST`**

**255.**

<img src="images/timezone/0255.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN PHILADELPHIA. THE DIRT ROAD GOES NORTH, SOUTH, EAST, AND WEST.">

**Type `WEST`**

**256.**

<img src="images/timezone/0256.png" width="400" alt="GOES NORTH, SOUTH, EAST, AND WEST. --------------- ENTER COMMAND?WEST YOU ARE IN FRONT OF THE COURTHOUSE.">

**Type `SOUTH`**

**257.**

<img src="images/timezone/0257.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN A PASTURE. A DIRT ROAD LEADS EAST AND WEST.">

**Type `SOUTH`**

**258.**

<img src="images/timezone/0258.png" width="400" alt="LEADS EAST AND WEST. --------------- ENTER COMMAND?SOUTH YOU ARE IN A PASTURE.">

**Type `SOUTH`**

**259.**

<img src="images/timezone/0259.png" width="400" alt="YOU ARE IN THE WOODS. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `GO MACHINE`**

**260.**

<img src="images/timezone/0260-1.png" width="400" alt="PLEASE INSERT DISK NUMBER 1B AND PRESS RETURN."><br>
<img src="images/timezone/0260-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0260.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

*Side 1B: drop `Time Zone (4am and san inc crack) disk B.dsk` on drive 1, and press Return.*

**Type `SIT`**

**261.**

<img src="images/timezone/0261-1.png" width="400" alt="O.K. THE CHAIR IS COMFORTABLE. AS YOU SIT DOWN, THE MACHINE SEEMS TO COME TO LIFE. YOU ARE INSIDE A TIME MACHINE. YOU HEAR"><br>
<img src="images/timezone/0261.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT. --------------- ENTER COMMAND">

**Type `ASIA`**

**262.**

<img src="images/timezone/0262-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0262.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `2082AD`**

**263.**

<img src="images/timezone/0263-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0263.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `PULL LEVER`**

**264.**

<img src="images/timezone/0264-1.png" width="400" alt="THE MACHINE STARTS VIBRATING VIOLENTLY AND YOU FEEL DIZZY, ALMOST TO THE POINT OF PASSING OUT. SUDDENLY YOU FEEL A TERRIBLE JOLT AND THE MACHINE IS STILL."><br>
<img src="images/timezone/0264-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0264.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EXIT MACHINE`**

### Asia, 2082 AD

**265.**

<img src="images/timezone/0265-1.png" width="400" alt="NE PLEASE INSERT DISK NUMBER 3F AND PRESS RETURN."><br>
<img src="images/timezone/0265-2.png" width="400" alt="YOU ARE ON A CITY STREET IN TOKYO. THE STREET DEAD-ENDS TO THE SOUTH. THERE IS A TIME MACHINE HERE, IT APPEARS TO BE PULSATING."><br>
<img src="images/timezone/0265.png" width="400" alt="STREET DEAD-ENDS TO THE SOUTH. THERE IS A TIME MACHINE HERE, IT APPEARS TO BE PULSATING.">

*Side 3F: drop `Time Zone (4am and san inc crack) disk F.dsk` on drive 1, and press Return.*

**Type `NORTH`**

**266.**

<img src="images/timezone/0266.png" width="400" alt="YOU ARE ON A CITY STREET IN TOKYO. THERE ARE STAIRS GOING DOWN FROM THE SIDEWALK HERE.">

**Type `DOWN`**

**267.**

<img src="images/timezone/0267.png" width="400" alt="YOU ARE IN A SUPERSONIC SUBWAY STATION. THERE ARE STAIRS GOING UP HERE. THERE IS A SIGN HERE.">

**Type `GO TRAIN`**

**268.**

<img src="images/timezone/0268.png" width="400" alt="--------------- ENTER COMMAND?GO TRAIN WHY DON&#39;T YOU HAVE A SEAT? YOU ARE INSIDE THE SUBWAY TRAIN.">

**Type `WEST SIDE`**

**269.**

<img src="images/timezone/0269.png" width="400" alt="THE DOOR CLOSES AND QUICKLY THE TRAIN SPEEDS OFF TO ITS&#39; DESTINATION. YOU ARE INSIDE THE SUBWAY TRAIN.">

**Type `LOOK`**

**270.**

<img src="images/timezone/0270.png" width="400" alt="THE TRAIN NOW COMES TO A HALT, AND THE DOOR OPENS. YOU ARE INSIDE THE SUBWAY TRAIN.">

**Type `LEAVE TRAIN`**

**271.**

<img src="images/timezone/0271.png" width="400" alt="YOU ARE IN A SUPERSONIC SUBWAY STATION. THERE ARE STAIRS GOING UP HERE. THERE IS A SIGN HERE.">

**Type `UP`**

**272.**

<img src="images/timezone/0272.png" width="400" alt="YOU ARE ON A CITY STREET IN TOKYO. THERE ARE STAIRS GOING DOWN FROM THE SIDEWALK HERE.">

**Type `EAST`**

**273.**

<img src="images/timezone/0273.png" width="400" alt="SIDEWALK HERE. --------------- ENTER COMMAND?EAST YOU ARE ON A CITY STREET IN TOKYO.">

**Type `SOUTH`**

**274.**

<img src="images/timezone/0274.png" width="400" alt="YOU ARE ON A CITY STREET IN TOKYO. --------------- ENTER COMMAND?SOUTH YOU ARE ON A CITY STREET IN TOKYO.">

**Type `SOUTH`**

**275.**

<img src="images/timezone/0275.png" width="400" alt="YOU ARE ON A CITY STREET IN TOKYO. THE STREET DEAD-ENDS TO THE SOUTH. THERE IS A WAREHOUSE TO THE EAST.">

**Type `USE KEY`**

**276.**

<img src="images/timezone/0276-1.png" width="400" alt="O.K. YOU ARE ON A CITY STREET IN TOKYO. THE STREET DEAD-ENDS TO THE SOUTH. THERE IS A WAREHOUSE TO THE EAST."><br>
<img src="images/timezone/0276.png" width="400" alt="YOU ARE ON A CITY STREET IN TOKYO. THE STREET DEAD-ENDS TO THE SOUTH. THERE IS A WAREHOUSE TO THE EAST.">

**Type `GET PADLOCK`**

**277.**

<img src="images/timezone/0277-1.png" width="400" alt="O.K. YOU ARE ON A CITY STREET IN TOKYO. THE STREET DEAD-ENDS TO THE SOUTH. THERE IS A WAREHOUSE TO THE EAST."><br>
<img src="images/timezone/0277.png" width="400" alt="YOU ARE ON A CITY STREET IN TOKYO. THE STREET DEAD-ENDS TO THE SOUTH. THERE IS A WAREHOUSE TO THE EAST.">

**Type `DROP PADLOCK`**

**278.**

<img src="images/timezone/0278.png" width="400" alt="YOU ARE ON A CITY STREET IN TOKYO. THE STREET DEAD-ENDS TO THE SOUTH. THERE IS A WAREHOUSE TO THE EAST.">

**Type `OPEN DOOR`**

**279.**

<img src="images/timezone/0279-1.png" width="400" alt="O.K. YOU ARE ON A CITY STREET IN TOKYO. THE STREET DEAD-ENDS TO THE SOUTH. THERE IS A WAREHOUSE TO THE EAST."><br>
<img src="images/timezone/0279.png" width="400" alt="YOU ARE ON A CITY STREET IN TOKYO. THE STREET DEAD-ENDS TO THE SOUTH. THERE IS A WAREHOUSE TO THE EAST.">

**Type `EAST`**

**280.**

<img src="images/timezone/0280.png" width="400" alt="--------------- ENTER COMMAND?EAST THERE IS ONE YEN ON THE FLOOR. YOU ARE INSIDE A BIG, EMPTY WAREHOUSE.">

**Type `GET YEN`**

**281.**

<img src="images/timezone/0281.png" width="400" alt="YOU ARE INSIDE A BIG, EMPTY WAREHOUSE. --------------- ENTER COMMAND?GET YEN YOU ARE INSIDE A BIG, EMPTY WAREHOUSE.">

**Type `WEST`**

**282.**

<img src="images/timezone/0282.png" width="400" alt="YOU ARE ON A CITY STREET IN TOKYO. THE STREET DEAD-ENDS TO THE SOUTH. THERE IS A WAREHOUSE TO THE EAST.">

**Type `NORTH`**

**283.**

<img src="images/timezone/0283.png" width="400" alt="A WAREHOUSE TO THE EAST. --------------- ENTER COMMAND?NORTH YOU ARE ON A CITY STREET IN TOKYO.">

**Type `NORTH`**

**284.**

<img src="images/timezone/0284.png" width="400" alt="YOU ARE ON A CITY STREET IN TOKYO. --------------- ENTER COMMAND?NORTH YOU ARE ON A CITY STREET IN TOKYO.">

**Type `WEST`**

**285.**

<img src="images/timezone/0285.png" width="400" alt="YOU ARE ON A CITY STREET IN TOKYO. THERE ARE STAIRS GOING DOWN FROM THE SIDEWALK HERE.">

**Type `DOWN`**

**286.**

<img src="images/timezone/0286.png" width="400" alt="YOU ARE IN A SUPERSONIC SUBWAY STATION. THERE ARE STAIRS GOING UP HERE. THERE IS A SIGN HERE.">

**Type `GO TRAIN`**

**287.**

<img src="images/timezone/0287.png" width="400" alt="IS A SIGN HERE. --------------- ENTER COMMAND?GO TRAIN YOU ARE INSIDE THE SUBWAY TRAIN.">

**Type `SIT`**

**288.**

<img src="images/timezone/0288-1.png" width="400" alt="O.K. YOU ARE SITTING. AS SOON AS YOU SIT DOWN IN THE TRAIN, THE DOOR CLOSES AND QUICKLY THE TRAIN SPEEDS OFF TO ITS&#39; DESTINATION."><br>
<img src="images/timezone/0288.png" width="400" alt="THE DOOR CLOSES AND QUICKLY THE TRAIN SPEEDS OFF TO ITS&#39; DESTINATION. YOU ARE INSIDE THE SUBWAY TRAIN.">

**Type `LOOK`**

**289.**

<img src="images/timezone/0289.png" width="400" alt="THE TRAIN NOW COMES TO A HALT, AND THE DOOR OPENS. YOU ARE INSIDE THE SUBWAY TRAIN.">

**Type `LEAVE TRAIN`**

**290.**

<img src="images/timezone/0290.png" width="400" alt="YOU ARE IN A SUPERSONIC SUBWAY STATION. THERE ARE STAIRS GOING UP HERE. THERE IS A SIGN HERE.">

**Type `GO TRAIN`**

**291.**

<img src="images/timezone/0291.png" width="400" alt="IS A SIGN HERE. --------------- ENTER COMMAND?GO TRAIN YOU ARE INSIDE THE SUBWAY TRAIN.">

**Type `SIT`**

**292.**

<img src="images/timezone/0292-1.png" width="400" alt="O.K. YOU ARE SITTING. AS SOON AS YOU SIT DOWN IN THE TRAIN, A VOICE FROM AN INTERCOM ASKS IF YOU WOULD LIKE TO GO TO THE NORTH SIDE, THE"><br>
<img src="images/timezone/0292.png" width="400" alt="SOUTH SIDE, THE EAST SIDE OR THE WEST SIDE? YOU ARE INSIDE THE SUBWAY TRAIN.">

**Type `NORTH SIDE`**

**293.**

<img src="images/timezone/0293.png" width="400" alt="THE DOOR CLOSES AND QUICKLY THE TRAIN SPEEDS OFF TO ITS&#39; DESTINATION. YOU ARE INSIDE THE SUBWAY TRAIN.">

**Type `LOOK`**

**294.**

<img src="images/timezone/0294.png" width="400" alt="THE TRAIN NOW COMES TO A HALT, AND THE DOOR OPENS. YOU ARE INSIDE THE SUBWAY TRAIN.">

**Type `LEAVE TRAIN`**

**295.**

<img src="images/timezone/0295.png" width="400" alt="YOU ARE IN A SUPERSONIC SUBWAY STATION. THERE ARE STAIRS GOING UP HERE. THERE IS A SIGN HERE.">

**Type `UP`**

**296.**

<img src="images/timezone/0296.png" width="400" alt="YOU ARE ON A CITY STREET IN TOKYO. THERE ARE STAIRS GOING DOWN FROM THE SIDEWALK HERE.">

**Type `WEST`**

**297.**

<img src="images/timezone/0297.png" width="400" alt="SIDEWALK HERE. --------------- ENTER COMMAND?WEST YOU ARE ON A CITY STREET IN TOKYO.">

**Type `WEST`**

**298.**

<img src="images/timezone/0298.png" width="400" alt="YOU ARE ON A CITY STREET IN TOKYO. --------------- ENTER COMMAND?WEST YOU ARE ON A CITY STREET IN TOKYO.">

**Type `NORTH`**

**299.**

<img src="images/timezone/0299.png" width="400" alt="YOU ARE ON A CITY STREET IN TOKYO. THERE IS A JAPANESE RESTAURANT TO THE NORTH.">

**Type `OPEN DOOR`**

**300.**

<img src="images/timezone/0300-1.png" width="400" alt="O.K. YOU ARE ON A CITY STREET IN TOKYO. THERE IS A JAPANESE RESTAURANT TO THE NORTH."><br>
<img src="images/timezone/0300.png" width="400" alt="YOU ARE ON A CITY STREET IN TOKYO. THERE IS A JAPANESE RESTAURANT TO THE NORTH.">

**Type `NORTH`**

**301.**

<img src="images/timezone/0301.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE SITTING AT A TABLE IN A JAPANESE RESTAURANT.">

**Type `LOOK`**

**302.**

<img src="images/timezone/0302-1.png" width="400" alt="YOU SEE NOTHING SPECIAL. AS YOU ARE SITTING AT THE TABLE, A JAPANESE WAITER SHOWS UP AND HANDS YOU A MENU."><br>
<img src="images/timezone/0302.png" width="400" alt="A MENU. YOU ARE SITTING AT A TABLE IN A JAPANESE RESTAURANT.">

**Type `READ MENU`**

**303.**

<img src="images/timezone/0303.png" width="400" alt="--------------- ENTER COMMAND?READ MENU YOU ARE SITTING AT A TABLE IN A JAPANESE RESTAURANT.">

**Type `LOOK`**

**304.**

<img src="images/timezone/0304-1.png" width="400" alt="THE WAITER LOOKS AT YOU AND SAYS, &#34;WOULD YOU LIKE TO ORDER DINNER NUMBER 1,2 OR 3?&#34;. YOU ARE SITTING AT A TABLE IN A"><br>
<img src="images/timezone/0304.png" width="400" alt="1,2 OR 3?&#34;. YOU ARE SITTING AT A TABLE IN A JAPANESE RESTAURANT.">

**Type `2`**

**305.**

<img src="images/timezone/0305-1.png" width="400" alt="THE WAITER THANKS YOU AND TAKES BACK THE MENU. IN A FEW MINUTES, HE COMES BACK WITH A BOWL OF SUKIYAKI AND SETS IT ON THE TABLE IN FRONT OF YOU."><br>
<img src="images/timezone/0305.png" width="400" alt="IT ON THE TABLE IN FRONT OF YOU. YOU ARE SITTING AT A TABLE IN A JAPANESE RESTAURANT.">

**Type `EAT`**

**306.**

<img src="images/timezone/0306-1.png" width="400" alt="THE WAITER TAKES BACK YOUR BOWL AND SAYS,&#34;THAT WILL BE TWO YEN PLEASE&#34;. WHEN YOU GIVE HIM YOUR YEN, HE SAYS, &#34;I TOLD YOU TWO YEN, NOT ONE YEN. TO PAY"><br>
<img src="images/timezone/0306-2.png" width="400" alt="FOR YOUR MEAL, I ORDER YOU TO GO IN THE KITCHEN AND WASH DISHES!!&#34; ANGRILY HE STALKS AWAY. YOU ARE SITTING AT A TABLE IN A"><br>
<img src="images/timezone/0306.png" width="400" alt="STALKS AWAY. YOU ARE SITTING AT A TABLE IN A JAPANESE RESTAURANT.">

**Type `NORTH`**

**307.**

<img src="images/timezone/0307.png" width="400" alt="YOU ARE IN THE KITCHEN OF THE JAPANESE RESTAURANT. THERE IS A JAPANESE COOK HERE. THE DOORWAY IS TO THE SOUTH.">

**Type `WASH DISHES`**

**308.**

<img src="images/timezone/0308-1.png" width="400" alt="O.K. YOU PATIENTLY WASH A SINK FULL OF DISHES. YOU ARE IN THE KITCHEN OF THE JAPANESE"><br>
<img src="images/timezone/0308.png" width="400" alt="YOU ARE IN THE KITCHEN OF THE JAPANESE RESTAURANT. THERE IS A JAPANESE COOK HERE. THE DOORWAY IS TO THE SOUTH.">

**Type `OPEN DRAWER`**

**309.**

<img src="images/timezone/0309-1.png" width="400" alt="O.K. YOU ARE IN THE KITCHEN OF THE JAPANESE RESTAURANT. THERE IS A JAPANESE COOK HERE. THE DOORWAY IS TO THE SOUTH."><br>
<img src="images/timezone/0309.png" width="400" alt="YOU ARE IN THE KITCHEN OF THE JAPANESE RESTAURANT. THERE IS A JAPANESE COOK HERE. THE DOORWAY IS TO THE SOUTH.">

**Type `LOOK DRAWER`**

**310.**

<img src="images/timezone/0310-1.png" width="400" alt="THERE ARE SOME MATCHES HERE. YOU ARE IN THE KITCHEN OF THE JAPANESE RESTAURANT. THERE IS A JAPANESE COOK HERE. THE DOORWAY IS TO THE SOUTH."><br>
<img src="images/timezone/0310.png" width="400" alt="YOU ARE IN THE KITCHEN OF THE JAPANESE RESTAURANT. THERE IS A JAPANESE COOK HERE. THE DOORWAY IS TO THE SOUTH.">

**Type `GET MATCHES`**

**311.**

<img src="images/timezone/0311-1.png" width="400" alt="O.K. YOU ARE IN THE KITCHEN OF THE JAPANESE RESTAURANT. THERE IS A JAPANESE COOK HERE. THE DOORWAY IS TO THE SOUTH."><br>
<img src="images/timezone/0311.png" width="400" alt="YOU ARE IN THE KITCHEN OF THE JAPANESE RESTAURANT. THERE IS A JAPANESE COOK HERE. THE DOORWAY IS TO THE SOUTH.">

**Type `SOUTH`**

**312.**

<img src="images/timezone/0312.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE SITTING AT A TABLE IN A JAPANESE RESTAURANT.">

**Type `SOUTH`**

**313.**

<img src="images/timezone/0313.png" width="400" alt="YOU ARE ON A CITY STREET IN TOKYO. THERE IS A JAPANESE RESTAURANT TO THE NORTH.">

**Type `SOUTH`**

**314.**

<img src="images/timezone/0314.png" width="400" alt="NORTH. --------------- ENTER COMMAND?SOUTH YOU ARE ON A CITY STREET IN TOKYO.">

**Type `EAST`**

**315.**

<img src="images/timezone/0315.png" width="400" alt="YOU ARE ON A CITY STREET IN TOKYO. --------------- ENTER COMMAND?EAST YOU ARE ON A CITY STREET IN TOKYO.">

**Type `EAST`**

**316.**

<img src="images/timezone/0316.png" width="400" alt="YOU ARE ON A CITY STREET IN TOKYO. THERE ARE STAIRS GOING DOWN FROM THE SIDEWALK HERE.">

**Type `DOWN`**

**317.**

<img src="images/timezone/0317.png" width="400" alt="YOU ARE IN A SUPERSONIC SUBWAY STATION. THERE ARE STAIRS GOING UP HERE. THERE IS A SIGN HERE.">

**Type `GO TRAIN`**

**318.**

<img src="images/timezone/0318.png" width="400" alt="IS A SIGN HERE. --------------- ENTER COMMAND?GO TRAIN YOU ARE INSIDE THE SUBWAY TRAIN.">

**Type `SIT`**

**319.**

<img src="images/timezone/0319-1.png" width="400" alt="O.K. YOU ARE SITTING. AS SOON AS YOU SIT DOWN IN THE TRAIN, THE DOOR CLOSES AND QUICKLY THE TRAIN SPEEDS OFF TO ITS&#39; DESTINATION."><br>
<img src="images/timezone/0319.png" width="400" alt="THE DOOR CLOSES AND QUICKLY THE TRAIN SPEEDS OFF TO ITS&#39; DESTINATION. YOU ARE INSIDE THE SUBWAY TRAIN.">

**Type `LOOK`**

**320.**

<img src="images/timezone/0320.png" width="400" alt="THE TRAIN NOW COMES TO A HALT, AND THE DOOR OPENS. YOU ARE INSIDE THE SUBWAY TRAIN.">

**Type `LEAVE TRAIN`**

**321.**

<img src="images/timezone/0321.png" width="400" alt="YOU ARE IN A SUPERSONIC SUBWAY STATION. THERE ARE STAIRS GOING UP HERE. THERE IS A SIGN HERE.">

**Type `UP`**

**322.**

<img src="images/timezone/0322.png" width="400" alt="YOU ARE ON A CITY STREET IN TOKYO. THERE ARE STAIRS GOING DOWN FROM THE SIDEWALK HERE.">

**Type `SOUTH`**

**323.**

<img src="images/timezone/0323-1.png" width="400" alt="YOU ARE ON A CITY STREET IN TOKYO. THE STREET DEAD-ENDS TO THE SOUTH. THERE IS A TIME MACHINE HERE, IT APPEARS TO BE PULSATING."><br>
<img src="images/timezone/0323.png" width="400" alt="STREET DEAD-ENDS TO THE SOUTH. THERE IS A TIME MACHINE HERE, IT APPEARS TO BE PULSATING.">

**Type `GO MACHINE`**

**324.**

<img src="images/timezone/0324-1.png" width="400" alt="PLEASE INSERT DISK NUMBER 1B AND PRESS RETURN."><br>
<img src="images/timezone/0324-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0324.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

*Side 1B: drop `Time Zone (4am and san inc crack) disk B.dsk` on drive 1, and press Return.*

**Type `SIT`**

**325.**

<img src="images/timezone/0325-1.png" width="400" alt="O.K. THE CHAIR IS COMFORTABLE. AS YOU SIT DOWN, THE MACHINE SEEMS TO COME TO LIFE. YOU ARE INSIDE A TIME MACHINE. YOU HEAR"><br>
<img src="images/timezone/0325.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT. --------------- ENTER COMMAND">

**Type `PUSH BUTTON`**

**326.**

<img src="images/timezone/0326-1.png" width="400" alt="THE MACHINE STARTS VIBRATING VIOLENTLY AND YOU FEEL DIZZY, ALMOST TO THE POINT OF PASSING OUT. SUDDENLY YOU FEEL A TERRIBLE JOLT AND THE MACHINE IS STILL."><br>
<img src="images/timezone/0326-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0326.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EXIT MACHINE`**

### Home, 1982

**327.**

<img src="images/timezone/0327.png" width="400" alt="YOU ARE IN A FIELD OF DRY GRASS. THERE IS A STRANGE LOOKING MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `LIGHT TORCH`**

**328.**

<img src="images/timezone/0328.png" width="400" alt="YOU ARE IN A FIELD OF DRY GRASS. THERE IS A STRANGE LOOKING MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `DROP MATCHES`**

**329.**

<img src="images/timezone/0329.png" width="400" alt="YOU ARE IN A FIELD OF DRY GRASS. THERE IS A STRANGE LOOKING MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `DROP BAR`**

**330.**

<img src="images/timezone/0330.png" width="400" alt="YOU ARE IN A FIELD OF DRY GRASS. THERE IS A STRANGE LOOKING MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `DROP KEY`**

**331.**

<img src="images/timezone/0331.png" width="400" alt="YOU ARE IN A FIELD OF DRY GRASS. THERE IS A STRANGE LOOKING MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `DROP SAW`**

**332.**

<img src="images/timezone/0332.png" width="400" alt="YOU ARE IN A FIELD OF DRY GRASS. THERE IS A STRANGE LOOKING MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `GO MACHINE`**

**333.**

<img src="images/timezone/0333-1.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0333.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `SIT`**

**334.**

<img src="images/timezone/0334-1.png" width="400" alt="O.K. THE CHAIR IS COMFORTABLE. AS YOU SIT DOWN, THE MACHINE SEEMS TO COME TO LIFE. YOU ARE INSIDE A TIME MACHINE. YOU HEAR"><br>
<img src="images/timezone/0334.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT. --------------- ENTER COMMAND">

**Type `EUROPE`**

**335.**

<img src="images/timezone/0335-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0335.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `1700AD`**

**336.**

<img src="images/timezone/0336-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0336.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `PULL LEVER`**

**337.**

<img src="images/timezone/0337-1.png" width="400" alt="THE MACHINE STARTS VIBRATING VIOLENTLY AND YOU FEEL DIZZY, ALMOST TO THE POINT OF PASSING OUT. SUDDENLY YOU FEEL A TERRIBLE JOLT AND THE MACHINE IS STILL."><br>
<img src="images/timezone/0337-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0337.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EXIT MACHINE`**

### Europe, 1700 AD

**338.**

<img src="images/timezone/0338-1.png" width="400" alt="NE PLEASE INSERT DISK NUMBER 3F AND PRESS RETURN."><br>
<img src="images/timezone/0338.png" width="400" alt="YOU ARE IN A GENTLE GREEN VALLEY. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

*Side 3F: drop `Time Zone (4am and san inc crack) disk F.dsk` on drive 1, and press Return.*

**Type `NORTH`**

**339.**

<img src="images/timezone/0339.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN A VINEYARD. THE GRAPES LOOK DELICIOUS.">

**Type `NORTH`**

**340.**

<img src="images/timezone/0340.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE ON A COBBLESTONE ROAD GOING NORTH. THERE IS A SIGN HERE.">

**Type `NORTH`**

**341.**

<img src="images/timezone/0341.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN THE TOWN OF PARIS DURING NAPOLEON&#39;S RULE.">

**Type `NORTH`**

**342.**

<img src="images/timezone/0342.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE ON THE OUTSKIRTS OF PARIS. SMALL HOUSES LINE THE STREET.">

**Type `EAST`**

**343.**

<img src="images/timezone/0343.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE ON THE WEST SIDE OF THE STONE FENCE.">

**Type `NORTH`**

**344.**

<img src="images/timezone/0344.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE AT THE OUTSIDE NORTHWEST CORNER OF THE STONE FENCE.">

**Type `EAST`**

**345.**

<img src="images/timezone/0345.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE AT THE BACK OF THE STONE FENCE. IT GOES EAST AND WEST.">

**Type `CLIMB FENCE`**

**346.**

<img src="images/timezone/0346.png" width="400" alt="--------------- ENTER COMMAND?CLIMB FENC E YOU ARE ON TOP OF A STONEWALL.">

**Type `INSIDE`**

**347.**

<img src="images/timezone/0347.png" width="400" alt="--------------- ENTER COMMAND?INSIDE YOU ARE INSIDE THE STONE FENCE. THE FENCE GOES EAST AND WEST.">

**Type `SOUTH`**

**348.**

<img src="images/timezone/0348.png" width="400" alt="FENCE GOES EAST AND WEST. --------------- ENTER COMMAND?SOUTH YOU ARE IN A PRETTY GARDEN.">

**Type `SOUTH`**

**349.**

<img src="images/timezone/0349.png" width="400" alt="YOU ARE IN A PRETTY GARDEN. --------------- ENTER COMMAND?SOUTH YOU SEE A SIDE DOOR OF THE PALACE.">

**Type `OPEN DOOR`**

**350.**

<img src="images/timezone/0350.png" width="400" alt="YOU SEE A SIDE DOOR OF THE PALACE. --------------- ENTER COMMAND?OPEN DOOR YOU SEE A SIDE DOOR OF THE PALACE.">

**Type `EAST`**

**351.**

<img src="images/timezone/0351.png" width="400" alt="YOU SEE A SIDE DOOR OF THE PALACE. --------------- ENTER COMMAND?EAST YOU ARE IN AN EAST/WEST HALLWAY.">

**Type `EAST`**

**352.**

<img src="images/timezone/0352.png" width="400" alt="YOU ARE IN A HALLWAY. THERE IS A STAIRWAY GOING UP. AN OPEN DOORWAY IS TO THE NORTH.">

**Type `UP`**

**353.**

<img src="images/timezone/0353.png" width="400" alt="YOU ARE IN AN EAST/WEST HALLWAY AT THE THE TOP OF THE STAIRS. THERE IS A CLOSED DOOR HERE.">

**Type `EAST`**

**354.**

<img src="images/timezone/0354.png" width="400" alt="--------------- ENTER COMMAND?EAST THERE IS A CLOSED DOOR HERE. THE HALLWAY ENDS TO THE EAST.">

**Type `OPEN DOOR`**

**355.**

<img src="images/timezone/0355.png" width="400" alt="THE HALLWAY ENDS TO THE EAST. --------------- ENTER COMMAND?OPEN DOOR THE HALLWAY ENDS TO THE EAST.">

**Type `NORTH`**

**356.**

<img src="images/timezone/0356.png" width="400" alt="THERE IS PERFUME AND A PRETTY COMB ON THE DRESSER. YOU ARE IN JOSEPHINE&#39;S BEDROOM.">

**Type `GET PERFUME`**

**357.**

<img src="images/timezone/0357.png" width="400" alt="--------------- ENTER COMMAND?GET PERFUM E YOU ARE IN JOSEPHINE&#39;S BEDROOM.">

**Type `GET COMB`**

**358.**

<img src="images/timezone/0358.png" width="400" alt="YOU ARE IN JOSEPHINE&#39;S BEDROOM. --------------- ENTER COMMAND?GET COMB YOU ARE IN JOSEPHINE&#39;S BEDROOM.">

**Type `SOUTH`**

**359.**

<img src="images/timezone/0359.png" width="400" alt="YOU ARE IN JOSEPHINE&#39;S BEDROOM. --------------- ENTER COMMAND?SOUTH THE HALLWAY ENDS TO THE EAST.">

**Type `WEST`**

**360.**

<img src="images/timezone/0360.png" width="400" alt="YOU ARE IN AN EAST/WEST HALLWAY AT THE THE TOP OF THE STAIRS. THERE IS A CLOSED DOOR HERE.">

**Type `DOWN`**

**361.**

<img src="images/timezone/0361.png" width="400" alt="YOU ARE IN A HALLWAY. THERE IS A STAIRWAY GOING UP. AN OPEN DOORWAY IS TO THE NORTH.">

**Type `WEST`**

**362.**

<img src="images/timezone/0362.png" width="400" alt="TO THE NORTH. --------------- ENTER COMMAND?WEST YOU ARE IN AN EAST/WEST HALLWAY.">

**Type `WEST`**

**363.**

<img src="images/timezone/0363.png" width="400" alt="YOU ARE IN AN EAST/WEST HALLWAY. --------------- ENTER COMMAND?WEST YOU SEE A SIDE DOOR OF THE PALACE.">

**Type `NORTH`**

**364.**

<img src="images/timezone/0364.png" width="400" alt="YOU SEE A SIDE DOOR OF THE PALACE. --------------- ENTER COMMAND?NORTH YOU ARE IN A PRETTY GARDEN.">

**Type `NORTH`**

**365.**

<img src="images/timezone/0365.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE INSIDE THE STONE FENCE. THE FENCE GOES EAST AND WEST.">

**Type `CLIMB FENCE`**

**366.**

<img src="images/timezone/0366.png" width="400" alt="--------------- ENTER COMMAND?CLIMB FENC E YOU ARE ON TOP OF A STONEWALL.">

**Type `OUTSIDE`**

**367.**

<img src="images/timezone/0367.png" width="400" alt="--------------- ENTER COMMAND?OUTSIDE YOU ARE AT THE BACK OF THE STONE FENCE. IT GOES EAST AND WEST.">

**Type `WEST`**

**368.**

<img src="images/timezone/0368.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE AT THE OUTSIDE NORTHWEST CORNER OF THE STONE FENCE.">

**Type `SOUTH`**

**369.**

<img src="images/timezone/0369.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE ON THE WEST SIDE OF THE STONE FENCE.">

**Type `WEST`**

**370.**

<img src="images/timezone/0370.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE ON THE OUTSKIRTS OF PARIS. SMALL HOUSES LINE THE STREET.">

**Type `SOUTH`**

**371.**

<img src="images/timezone/0371.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN THE TOWN OF PARIS DURING NAPOLEON&#39;S RULE.">

**Type `SOUTH`**

**372.**

<img src="images/timezone/0372.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE ON A COBBLESTONE ROAD GOING NORTH. THERE IS A SIGN HERE.">

**Type `SOUTH`**

**373.**

<img src="images/timezone/0373.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN A VINEYARD. THE GRAPES LOOK DELICIOUS.">

**Type `SOUTH`**

**374.**

<img src="images/timezone/0374.png" width="400" alt="YOU ARE IN A GENTLE GREEN VALLEY. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `GO MACHINE`**

**375.**

<img src="images/timezone/0375-1.png" width="400" alt="PLEASE INSERT DISK NUMBER 1B AND PRESS RETURN."><br>
<img src="images/timezone/0375-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0375.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

*Side 1B: drop `Time Zone (4am and san inc crack) disk B.dsk` on drive 1, and press Return.*

**Type `SIT`**

**376.**

<img src="images/timezone/0376-1.png" width="400" alt="O.K. THE CHAIR IS COMFORTABLE. AS YOU SIT DOWN, THE MACHINE SEEMS TO COME TO LIFE. YOU ARE INSIDE A TIME MACHINE. YOU HEAR"><br>
<img src="images/timezone/0376.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT. --------------- ENTER COMMAND">

**Type `NA`**

**377.**

<img src="images/timezone/0377-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0377.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `1400AD`**

**378.**

<img src="images/timezone/0378-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0378.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `PULL LEVER`**

**379.**

<img src="images/timezone/0379-1.png" width="400" alt="THE MACHINE STARTS VIBRATING VIOLENTLY AND YOU FEEL DIZZY, ALMOST TO THE POINT OF PASSING OUT. SUDDENLY YOU FEEL A TERRIBLE JOLT AND THE MACHINE IS STILL."><br>
<img src="images/timezone/0379-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0379.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EXIT MACHINE`**

### North America, 1400 AD

**380.**

<img src="images/timezone/0380-1.png" width="400" alt="NE PLEASE INSERT DISK NUMBER 2C AND PRESS RETURN."><br>
<img src="images/timezone/0380.png" width="400" alt="YOU ARE ON A PRAIRIE. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

*Side 2C: drop `Time Zone (4am and san inc crack) disk C.dsk` on drive 1, and press Return.*

**Type `WEST`**

**381.**

<img src="images/timezone/0381.png" width="400" alt="YOU HEAR A RUMBLING NOISE COMING FROM THE NORTH. YOU ARE ROAMING ON A PRAIRIE.">

**Type `WEST`**

**382.**

<img src="images/timezone/0382.png" width="400" alt="YOU HEAR A RUMBLING NOISE COMING FROM THE NORTH. YOU ARE IN A DUSTY PRAIRIE.">

**Type `SOUTH`**

**383.**

<img src="images/timezone/0383-1.png" width="400" alt="THE RUMBLING NOISE HAS CHANGED TO A LOUD ROAR AND YOU SEE A HERD OF BUFFALO HEADED YOUR WAY!! YOU ARE IN A PRAIRIE. THERE IS A GULLEY"><br>
<img src="images/timezone/0383.png" width="400" alt="HEADED YOUR WAY!! YOU ARE IN A PRAIRIE. THERE IS A GULLEY HERE.">

**Type `GO GULLEY`**

**384.**

<img src="images/timezone/0384-1.png" width="400" alt="THE STAMPEDE OF BUFFALO PASSES OVER YOU AS YOU LAY IN THE GULLEY. AS SOON AS THE BUFFALO PASS BY, YOU GET OUT OF THE GULLEY AND DUST YOURSELF OFF."><br>
<img src="images/timezone/0384.png" width="400" alt="GULLEY AND DUST YOURSELF OFF. YOU ARE IN A PRAIRIE. THERE IS A GULLEY HERE.">

**Type `NORTH`**

**385.**

<img src="images/timezone/0385.png" width="400" alt="HERE. --------------- ENTER COMMAND?NORTH YOU ARE IN A DUSTY PRAIRIE.">

**Type `EAST`**

**386.**

<img src="images/timezone/0386.png" width="400" alt="YOU ARE IN A DUSTY PRAIRIE. --------------- ENTER COMMAND?EAST YOU ARE ROAMING ON A PRAIRIE.">

**Type `NORTH`**

**387.**

<img src="images/timezone/0387.png" width="400" alt="YOU ARE ROAMING ON A PRAIRIE. --------------- ENTER COMMAND?NORTH YOU ARE ON A WIDE PRAIRIE.">

**Type `NORTH`**

**388.**

<img src="images/timezone/0388.png" width="400" alt="YOU ARE ON A WIDE PRAIRIE. --------------- ENTER COMMAND?NORTH YOU ARE WALKING THROUGH A PRAIRIE.">

**Type `NORTH`**

**389.**

<img src="images/timezone/0389-1.png" width="400" alt="THERE IS AN INDIAN WITH BOW AND ARROWS DRAWN!! THERE IS A DEAD BUFFALO BESIDE HIM. YOU ARE IN A PRAIRIE. THERE IS A DEEP"><br>
<img src="images/timezone/0389.png" width="400" alt="YOU ARE IN A PRAIRIE. THERE IS A DEEP RAVINE TO THE NORTH. A BRIDGE CROSSES THE RAVINE.">

**Type `NORTH`**

**390.**

<img src="images/timezone/0390.png" width="400" alt="YOU ARE IN SOME LOW HILLS. YOU SEE SOME TEEPEES TO THE NORTH. A BRIDGE CROSSES A RAVINE TO THE SOUTH OF YOU.">

**Type `NORTH`**

**391.**

<img src="images/timezone/0391.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN AN INDIAN VILLAGE. THERE IS A BIG TEEPEE TO THE NORTH.">

**Type `GO TEEPEE`**

**392.**

<img src="images/timezone/0392.png" width="400" alt="THERE IS A BOW AND ARROWS HERE. YOU ARE INSIDE THE TEEPEE. THE INDIAN CHIEF IS SITTING HERE.">

**Type `TRADE COMB`**

**393.**

<img src="images/timezone/0393-1.png" width="400" alt="THE INDIAN CHIEF IS INTRIGUED BY THE PRETTY COMB. HE RUNS IT THROUGH HIS HAIR AND DECIDES TO KEEP IT. HE GIVES YOU THE BOW AND ARROWS."><br>
<img src="images/timezone/0393.png" width="400" alt="YOU THE BOW AND ARROWS. YOU ARE INSIDE THE TEEPEE. THE INDIAN CHIEF IS SITTING HERE.">

**Type `SOUTH`**

**394.**

<img src="images/timezone/0394.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN AN INDIAN VILLAGE. THERE IS A BIG TEEPEE TO THE NORTH.">

**Type `SOUTH`**

**395.**

<img src="images/timezone/0395.png" width="400" alt="YOU ARE IN SOME LOW HILLS. YOU SEE SOME TEEPEES TO THE NORTH. A BRIDGE CROSSES A RAVINE TO THE SOUTH OF YOU.">

**Type `SOUTH`**

**396.**

<img src="images/timezone/0396.png" width="400" alt="YOU ARE IN A PRAIRIE. THERE IS A DEEP RAVINE TO THE NORTH. A BRIDGE CROSSES THE RAVINE.">

**Type `SOUTH`**

**397.**

<img src="images/timezone/0397.png" width="400" alt="THE RAVINE. --------------- ENTER COMMAND?SOUTH YOU ARE WALKING THROUGH A PRAIRIE.">

**Type `SOUTH`**

**398.**

<img src="images/timezone/0398.png" width="400" alt="YOU ARE WALKING THROUGH A PRAIRIE. --------------- ENTER COMMAND?SOUTH YOU ARE ON A WIDE PRAIRIE.">

**Type `SOUTH`**

**399.**

<img src="images/timezone/0399.png" width="400" alt="YOU ARE ON A WIDE PRAIRIE. --------------- ENTER COMMAND?SOUTH YOU ARE ROAMING ON A PRAIRIE.">

**Type `EAST`**

**400.**

<img src="images/timezone/0400.png" width="400" alt="YOU ARE ON A PRAIRIE. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `GO MACHINE`**

**401.**

<img src="images/timezone/0401-1.png" width="400" alt="PLEASE INSERT DISK NUMBER 1B AND PRESS RETURN."><br>
<img src="images/timezone/0401-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0401.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

*Side 1B: drop `Time Zone (4am and san inc crack) disk B.dsk` on drive 1, and press Return.*

**Type `SIT`**

**402.**

<img src="images/timezone/0402-1.png" width="400" alt="O.K. THE CHAIR IS COMFORTABLE. AS YOU SIT DOWN, THE MACHINE SEEMS TO COME TO LIFE. YOU ARE INSIDE A TIME MACHINE. YOU HEAR"><br>
<img src="images/timezone/0402.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT. --------------- ENTER COMMAND">

**Type `EUROPE`**

**403.**

<img src="images/timezone/0403-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0403.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `1000AD`**

**404.**

<img src="images/timezone/0404-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0404.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `PULL LEVER`**

**405.**

<img src="images/timezone/0405-1.png" width="400" alt="THE MACHINE STARTS VIBRATING VIOLENTLY AND YOU FEEL DIZZY, ALMOST TO THE POINT OF PASSING OUT. SUDDENLY YOU FEEL A TERRIBLE JOLT AND THE MACHINE IS STILL."><br>
<img src="images/timezone/0405-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0405.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EXIT MACHINE`**

### Europe, 1000 AD

**406.**

<img src="images/timezone/0406-1.png" width="400" alt="NE PLEASE INSERT DISK NUMBER 3E AND PRESS RETURN."><br>
<img src="images/timezone/0406.png" width="400" alt="AND PRESS RETURN. YOU ARE IN THE FOREST. THERE IS A TIME MACHINE HERE.">

*Side 3E: drop `Time Zone (4am and san inc crack) disk E.dsk` on drive 1, and press Return.*

**Type `EAST`**

**407.**

<img src="images/timezone/0407.png" width="400" alt="MACHINE HERE. --------------- ENTER COMMAND?EAST YOU ARE IN THE FOREST.">

**Type `EAST`**

**408.**

<img src="images/timezone/0408.png" width="400" alt="YOU ARE IN THE FOREST. --------------- ENTER COMMAND?EAST YOU ARE IN THE FOREST.">

**Type `TALK ROBIN`**

**409.**

<img src="images/timezone/0409-1.png" width="400" alt="WHEN YOU SPEAK TO ROBIN HOOD HE ASKS IF YOU WOULD LIKE TO JOIN HIS BAND OF MERRY MEN, BUT WARNS YOU THAT THE TEST IS DANGEROUS."><br>
<img src="images/timezone/0409.png" width="400" alt="MERRY MEN, BUT WARNS YOU THAT THE TEST IS DANGEROUS. YOU ARE IN THE FOREST.">

**Type `YES`**

**410.**

<img src="images/timezone/0410-1.png" width="400" alt="ROBIN SAYS,&#34;THEN YOU MUST PROVE YOURSELF CAPABLE BY SLAYING THE EVIL DRYAD THAT HAUNTS THE BLACK FOREST. DON&#39;T RETURN UNLESS YOU SUCCEED&#34;."><br>
<img src="images/timezone/0410.png" width="400" alt="WITH THAT, YOU ARE CAST INTO THE DARK FOREST IN SEARCH OF YOUR WICKED PREY. YOU ARE LOST IN THE DARK FOREST.">

**Type `SOUTH`**

**411.**

<img src="images/timezone/0411.png" width="400" alt="YOU ARE LOST IN THE DARK FOREST. --------------- ENTER COMMAND?SOUTH YOU ARE LOST IN THE DARK FOREST.">

**Type `SOUTH`**

**412.**

<img src="images/timezone/0412.png" width="400" alt="YOU ARE LOST IN THE DARK FOREST. --------------- ENTER COMMAND?SOUTH YOU ARE LOST IN THE DARK FOREST.">

**Type `EAST`**

**413.**

<img src="images/timezone/0413-1.png" width="400" alt="YOU FEEL AN EVIL PRESENCE IN THE AIR. SUDDENLY, A HIDEOUS CREATURE STEPS FROM BEHIND A TREE. YOU ARE LOST IN THE DARK FOREST."><br>
<img src="images/timezone/0413.png" width="400" alt="SUDDENLY, A HIDEOUS CREATURE STEPS FROM BEHIND A TREE. YOU ARE LOST IN THE DARK FOREST.">

**Type `USE BOW`**

**414.**

<img src="images/timezone/0414-1.png" width="400" alt="YOU EXPERTLY LET FLY AN ARROW AND IT FINDS ITS WAY STRAIGHT TO THE DRYAD&#39;S HEART, HE STANDS THERE STUNNNED FOR A MOMENT, THEN CHANGES BACK TO HIS"><br>
<img src="images/timezone/0414.png" width="400" alt="MOMENT, THEN CHANGES BACK TO HIS ORIGINAL, WITHERED, FORM. YOU ARE LOST IN THE DARK FOREST.">

**Type `SOUTH`**

**415.**

<img src="images/timezone/0415.png" width="400" alt="YOU ARE LOST IN THE DARK FOREST. --------------- ENTER COMMAND?SOUTH YOU ARE LOST IN THE DARK FOREST.">

**Type `NORTH`**

**416.**

<img src="images/timezone/0416.png" width="400" alt="YOU ARE LOST IN THE DARK FOREST. --------------- ENTER COMMAND?NORTH YOU ARE LOST IN THE DARK FOREST.">

**Type `SOUTH`**

**417.**

<img src="images/timezone/0417.png" width="400" alt="YOU ARE LOST IN THE DARK FOREST. --------------- ENTER COMMAND?SOUTH YOU ARE LOST IN THE DARK FOREST.">

**Type `WEST`**

**418.**

<img src="images/timezone/0418-1.png" width="400" alt="YOU HAVE PROVEN YOUR BRAVERY AND SKILL!YOU ARE OFFICIALLY ONE OF ROBIN HOOD&#39;S MERRY MEN. ROBIN HOOD NOW HAS A TASK FOR YOU TO DO. THERE IS A BAG OF"><br>
<img src="images/timezone/0418-2.png" width="400" alt="MONEY IN THE BACK ROOM OF THE SHERIFF&#39;S OFFICE. GET IT, AND BRING IT BACK TO ROBIN HOOD. YOU ARE IN THE FOREST."><br>
<img src="images/timezone/0418.png" width="400" alt="OFFICE. GET IT, AND BRING IT BACK TO ROBIN HOOD. YOU ARE IN THE FOREST.">

**Type `WEST`**

**419.**

<img src="images/timezone/0419.png" width="400" alt="YOU ARE IN THE FOREST. --------------- ENTER COMMAND?WEST YOU ARE IN THE FOREST.">

**Type `NORTH`**

**420.**

<img src="images/timezone/0420.png" width="400" alt="YOU ARE IN THE FOREST. --------------- ENTER COMMAND?NORTH YOU ARE IN THE FOREST">

**Type `NORTH`**

**421.**

<img src="images/timezone/0421.png" width="400" alt="YOU ARE IN THE FOREST --------------- ENTER COMMAND?NORTH YOU ARE LOST IN THE FOREST.">

**Type `NORTH`**

**422.**

<img src="images/timezone/0422.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN A FOREST. THERE IS A SIGN HERE. A PATH LEADS TO THE NORTH.">

**Type `NORTH`**

**423.**

<img src="images/timezone/0423.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN THE VILLAGE OF NOTTINGHAM.A ROAD RUNS NORTH &amp; SOUTH.">

**Type `NORTH`**

**424.**

<img src="images/timezone/0424.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE ON A ROAD GOING EAST, WEST, &amp; SOUTH. A FOREST IS AROUND YOU.">

**Type `WEST`**

**425.**

<img src="images/timezone/0425-1.png" width="400" alt="YOU ARE ON AN EAST/WEST ROAD. AN ALLEY IS GOING NORTH FROM HERE. YOU SEE THE SOUTH-EAST CORNER OF THE SHERIFF&#39;S OFFICE."><br>
<img src="images/timezone/0425.png" width="400" alt="IS GOING NORTH FROM HERE. YOU SEE THE SOUTH-EAST CORNER OF THE SHERIFF&#39;S OFFICE.">

**Type `WEST`**

**426.**

<img src="images/timezone/0426.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE STANDING IN FRONT OF THE SHERIFF OF NOTTINGHAMS OFFICE.">

**Type `WEST`**

**427.**

<img src="images/timezone/0427-1.png" width="400" alt="YOU ARE ON AN EAST/WEST ROAD. AN ALLEY GOES NORTH FROM HERE. YOU SEE THE SOUTHWEST CORNER OF THE SHERIFF&#39;S OFFICE."><br>
<img src="images/timezone/0427.png" width="400" alt="GOES NORTH FROM HERE. YOU SEE THE SOUTHWEST CORNER OF THE SHERIFF&#39;S OFFICE.">

**Type `WEST`**

**428.**

<img src="images/timezone/0428.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE STANDING IN FRONT OF MAID MARIAN&#39;S HOUSE.">

**Type `OPEN DOOR`**

**429.**

<img src="images/timezone/0429.png" width="400" alt="THE DOOR IS NOW OPEN. YOU ARE STANDING IN FRONT OF MAID MARIAN&#39;S HOUSE.">

**Type `NORTH`**

**430.**

<img src="images/timezone/0430.png" width="400" alt="--------------- ENTER COMMAND?NORTH THERE IS A MIRROR ON THE TABLE. YOU ARE INSIDE MAID MARIAN&#39;S HOUSE.">

**Type `GET MIRROR`**

**431.**

<img src="images/timezone/0431.png" width="400" alt="--------------- ENTER COMMAND?GET MIRROR YOU ARE INSIDE MAID MARIAN&#39;S HOUSE.">

**Type `SOUTH`**

**432.**

<img src="images/timezone/0432.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE STANDING IN FRONT OF MAID MARIAN&#39;S HOUSE.">

**Type `EAST`**

**433.**

<img src="images/timezone/0433-1.png" width="400" alt="YOU ARE ON AN EAST/WEST ROAD. AN ALLEY GOES NORTH FROM HERE. YOU SEE THE SOUTHWEST CORNER OF THE SHERIFF&#39;S OFFICE."><br>
<img src="images/timezone/0433.png" width="400" alt="GOES NORTH FROM HERE. YOU SEE THE SOUTHWEST CORNER OF THE SHERIFF&#39;S OFFICE.">

**Type `NORTH`**

**434.**

<img src="images/timezone/0434.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN A NORTH/SOUTH ALLEY. THE SIDE OF THE SHERIFF&#39;S OFFICE IS HERE.">

**Type `NORTH`**

**435.**

<img src="images/timezone/0435.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN A BACK ALLEY. THE ALLEY GOES EAST &amp; SOUTH.">

**Type `EAST`**

**436.**

<img src="images/timezone/0436.png" width="400" alt="YOU ARE AT THE BACK OF THE SHERIFF&#39;S OFFICE. AN ALLEY GOES EAST &amp; WEST. THERE IS A WINDOW WITH BARS HERE.">

**Type `LOOK WINDOW`**

**437.**

<img src="images/timezone/0437.png" width="400" alt="W YOU ARE LOOKING THROUGH THE WINDOW AT THE BACK OF A BUILDING.">

**Type `TIE ROPE`**

**438.**

<img src="images/timezone/0438.png" width="400" alt="TIE THE ROPE TO WHAT? YOU ARE LOOKING THROUGH THE WINDOW AT THE BACK OF A BUILDING.">

**Type `TO ARROW`**

**439.**

<img src="images/timezone/0439.png" width="400" alt="THE ROPE IS NOW TIED TO THE ARROW. YOU ARE LOOKING THROUGH THE WINDOW AT THE BACK OF A BUILDING.">

**Type `SHOOT ARROW`**

**440.**

<img src="images/timezone/0440-1.png" width="400" alt="YOU SHOOT THE ARROW AND IT PIERCES THE BAG OF MONEY. YOU ARE LOOKING THROUGH THE WINDOW AT THE BACK OF A BUILDING."><br>
<img src="images/timezone/0440.png" width="400" alt="BAG OF MONEY. YOU ARE LOOKING THROUGH THE WINDOW AT THE BACK OF A BUILDING.">

**Type `PULL ROPE`**

**441.**

<img src="images/timezone/0441-1.png" width="400" alt="YOU PULL ON THE ROPE AND THE ARROW AND THE BAG OF MONEY COME TO YOU. YOU PULL THE ARROW OUT OF THE BAG AND UNTIE THE ROPE."><br>
<img src="images/timezone/0441.png" width="400" alt="ROPE. YOU ARE LOOKING THROUGH THE WINDOW AT THE BACK OF A BUILDING.">

**Type `DROP BOW`**

**442.**

<img src="images/timezone/0442.png" width="400" alt="--------------- ENTER COMMAND?DROP BOW YOU ARE LOOKING THROUGH THE WINDOW AT THE BACK OF A BUILDING.">

**Type `EAST`**

**443.**

<img src="images/timezone/0443.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN A BACK ALLEY. THE ALLEY GOES WEST &amp; SOUTH.">

**Type `SOUTH`**

**444.**

<img src="images/timezone/0444.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN A NORTH/SOUTH ALLEY, THE SIDE OF THE SHERIFF&#39;S OFFICE IS HERE.">

**Type `SOUTH`**

**445.**

<img src="images/timezone/0445-1.png" width="400" alt="YOU ARE ON AN EAST/WEST ROAD. AN ALLEY IS GOING NORTH FROM HERE. YOU SEE THE SOUTH-EAST CORNER OF THE SHERIFF&#39;S OFFICE."><br>
<img src="images/timezone/0445.png" width="400" alt="IS GOING NORTH FROM HERE. YOU SEE THE SOUTH-EAST CORNER OF THE SHERIFF&#39;S OFFICE.">

**Type `EAST`**

**446.**

<img src="images/timezone/0446.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE ON A ROAD GOING EAST, WEST, &amp; SOUTH. A FOREST IS AROUND YOU.">

**Type `SOUTH`**

**447.**

<img src="images/timezone/0447.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN THE VILLAGE OF NOTTINGHAM.A ROAD RUNS NORTH &amp; SOUTH.">

**Type `SOUTH`**

**448.**

<img src="images/timezone/0448.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN A FOREST. THERE IS A SIGN HERE. A PATH LEADS TO THE NORTH.">

**Type `SOUTH`**

**449.**

<img src="images/timezone/0449.png" width="400" alt="HERE. A PATH LEADS TO THE NORTH. --------------- ENTER COMMAND?SOUTH YOU ARE LOST IN THE FOREST.">

**Type `SOUTH`**

**450.**

<img src="images/timezone/0450.png" width="400" alt="YOU ARE LOST IN THE FOREST. --------------- ENTER COMMAND?SOUTH YOU ARE IN THE FOREST">

**Type `SOUTH`**

**451.**

<img src="images/timezone/0451.png" width="400" alt="YOU ARE IN THE FOREST --------------- ENTER COMMAND?SOUTH YOU ARE IN THE FOREST.">

**Type `WEST`**

**452.**

<img src="images/timezone/0452.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN THE FOREST. THERE IS A TIME MACHINE HERE.">

**Type `GO MACHINE`**

**453.**

<img src="images/timezone/0453-1.png" width="400" alt="PLEASE INSERT DISK NUMBER 1B AND PRESS RETURN."><br>
<img src="images/timezone/0453-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0453.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

*Side 1B: drop `Time Zone (4am and san inc crack) disk B.dsk` on drive 1, and press Return.*

**Type `SIT`**

**454.**

<img src="images/timezone/0454-1.png" width="400" alt="O.K. THE CHAIR IS COMFORTABLE. AS YOU SIT DOWN, THE MACHINE SEEMS TO COME TO LIFE. YOU ARE INSIDE A TIME MACHINE. YOU HEAR"><br>
<img src="images/timezone/0454.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT. --------------- ENTER COMMAND">

**Type `AUSTRALIA`**

**455.**

<img src="images/timezone/0455-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0455.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `PULL LEVER`**

**456.**

<img src="images/timezone/0456-1.png" width="400" alt="THE MACHINE STARTS VIBRATING VIOLENTLY AND YOU FEEL DIZZY, ALMOST TO THE POINT OF PASSING OUT. SUDDENLY YOU FEEL A TERRIBLE JOLT AND THE MACHINE IS STILL."><br>
<img src="images/timezone/0456-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0456.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EXIT MACHINE`**

### Australia, 1000 AD

**457.**

<img src="images/timezone/0457-1.png" width="400" alt="NE PLEASE INSERT DISK NUMBER 4H AND PRESS RETURN."><br>
<img src="images/timezone/0457.png" width="400" alt="YOU ARE IN THE GRASSLAND. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

*Side 4H: drop `Time Zone (4am and san inc crack) disk H.dsk` on drive 1, and press Return.*

**Type `NORTH`**

**458.**

<img src="images/timezone/0458.png" width="400" alt="PULSATING. --------------- ENTER COMMAND?NORTH YOU ARE IN THE GRASSLANDS.">

**Type `WEST`**

**459.**

<img src="images/timezone/0459.png" width="400" alt="YOU ARE IN THE GRASSLANDS. --------------- ENTER COMMAND?WEST YOU ARE IN THE GRASSLANDS.">

**Type `NORTH`**

**460.**

<img src="images/timezone/0460.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN THE BUSHLAND. THERE IS A DEAD KANGAROO HERE, PARTIALLY EATEN.">

**Type `EAST`**

**461.**

<img src="images/timezone/0461.png" width="400" alt="DEAD KANGAROO HERE, PARTIALLY EATEN. --------------- ENTER COMMAND?EAST YOU ARE LOST IN THE BUSHLANDS.">

**Type `NORTH`**

**462.**

<img src="images/timezone/0462.png" width="400" alt="YOU SEE A GROUP OF ABORIGINES IN THE DISTANCE. YOU ARE IN THE BUSHLANDS.">

**Type `NORTH`**

**463.**

<img src="images/timezone/0463-1.png" width="400" alt="THERE ARE A COUPLE OF ABORIGINES HERE. YOU HAD BETTER BE CAREFUL WITH THEM. THEY MIGHT KILL YOU WITH THEIR BOOMERANG."><br>
<img src="images/timezone/0463.png" width="400" alt="THEY MIGHT KILL YOU WITH THEIR BOOMERANG. YOU ARE IN THE BUSHLANDS.">

**Type `GIVE MIRROR`**

**464.**

<img src="images/timezone/0464-1.png" width="400" alt="THE ABORIGINES TAKE THE MIRROR AND LOOK AT IT CURIOUSLY. THEY LIKE THE MIRROR. THEY GIVE YOU THE BOOMERANG IN RETURN AND WALK AWAY."><br>
<img src="images/timezone/0464.png" width="400" alt="THEY GIVE YOU THE BOOMERANG IN RETURN AND WALK AWAY. YOU ARE IN THE BUSHLANDS.">

**Type `SOUTH`**

**465.**

<img src="images/timezone/0465.png" width="400" alt="YOU ARE IN THE BUSHLANDS. --------------- ENTER COMMAND?SOUTH YOU ARE IN THE BUSHLANDS.">

**Type `SOUTH`**

**466.**

<img src="images/timezone/0466.png" width="400" alt="YOU ARE IN THE BUSHLANDS. --------------- ENTER COMMAND?SOUTH YOU ARE LOST IN THE BUSHLANDS.">

**Type `WEST`**

**467.**

<img src="images/timezone/0467.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN THE BUSHLAND. THERE IS A DEAD KANGAROO HERE, PARTIALLY EATEN.">

**Type `SOUTH`**

**468.**

<img src="images/timezone/0468.png" width="400" alt="DEAD KANGAROO HERE, PARTIALLY EATEN. --------------- ENTER COMMAND?SOUTH YOU ARE IN THE GRASSLANDS.">

**Type `EAST`**

**469.**

<img src="images/timezone/0469.png" width="400" alt="YOU ARE IN THE GRASSLANDS. --------------- ENTER COMMAND?EAST YOU ARE IN THE GRASSLANDS.">

**Type `SOUTH`**

**470.**

<img src="images/timezone/0470.png" width="400" alt="YOU ARE IN THE GRASSLAND. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `GO MACHINE`**

**471.**

<img src="images/timezone/0471-1.png" width="400" alt="PLEASE INSERT DISK NUMBER 1B AND PRESS RETURN."><br>
<img src="images/timezone/0471-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0471.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

*Side 1B: drop `Time Zone (4am and san inc crack) disk B.dsk` on drive 1, and press Return.*

**Type `SIT`**

**472.**

<img src="images/timezone/0472-1.png" width="400" alt="O.K. THE CHAIR IS COMFORTABLE. AS YOU SIT DOWN, THE MACHINE SEEMS TO COME TO LIFE. YOU ARE INSIDE A TIME MACHINE. YOU HEAR"><br>
<img src="images/timezone/0472.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT. --------------- ENTER COMMAND">

**Type `EUROPE`**

**473.**

<img src="images/timezone/0473-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0473.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `PULL LEVER`**

**474.**

<img src="images/timezone/0474-1.png" width="400" alt="THE MACHINE STARTS VIBRATING VIOLENTLY AND YOU FEEL DIZZY, ALMOST TO THE POINT OF PASSING OUT. SUDDENLY YOU FEEL A TERRIBLE JOLT AND THE MACHINE IS STILL."><br>
<img src="images/timezone/0474-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0474.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EXIT MACHINE`**

### Europe, 1000 AD, again

**475.**

<img src="images/timezone/0475-1.png" width="400" alt="NE PLEASE INSERT DISK NUMBER 3E AND PRESS RETURN."><br>
<img src="images/timezone/0475.png" width="400" alt="AND PRESS RETURN. YOU ARE IN THE FOREST. THERE IS A TIME MACHINE HERE.">

*Side 3E: drop `Time Zone (4am and san inc crack) disk E.dsk` on drive 1, and press Return.*

**Type `EAST`**

**476.**

<img src="images/timezone/0476.png" width="400" alt="MACHINE HERE. --------------- ENTER COMMAND?EAST YOU ARE IN THE FOREST.">

**Type `EAST`**

**477.**

<img src="images/timezone/0477.png" width="400" alt="YOU ARE IN THE FOREST. --------------- ENTER COMMAND?EAST YOU ARE IN THE FOREST.">

**Type `GIVE MONEY`**

**478.**

<img src="images/timezone/0478-1.png" width="400" alt="ROBIN HOOD TAKES THE BAG OF MONEY, THANKS YOU, AND THEN LEAVES WITH HIS MEN. YOU ARE IN THE FOREST."><br>
<img src="images/timezone/0478.png" width="400" alt="THANKS YOU, AND THEN LEAVES WITH HIS MEN. YOU ARE IN THE FOREST.">

**Type `WEST`**

**479.**

<img src="images/timezone/0479.png" width="400" alt="ONE OF THE SHERIFF&#39;S KNIGHTS IS AFTER YOU FOR STEALING THE MONEY. YOU ARE IN THE FOREST.">

**Type `THROW BOOMERANG`**

**480.**

<img src="images/timezone/0480-1.png" width="400" alt="YOU THROW THE BOOMERANG AS HARD AS YOU CAN AND HIT THE KNIGHT ON THE HEAD AND KNOCK HIM FROM THE HORSE. THE KNIGHT IS DEAD."><br>
<img src="images/timezone/0480.png" width="400" alt="DEAD. THERE IS A LANCE LAYING HERE. YOU ARE IN THE FOREST.">

**Type `GET LANCE`**

**481.**

<img src="images/timezone/0481.png" width="400" alt="YOU ARE IN THE FOREST. --------------- ENTER COMMAND?GET LANCE YOU ARE IN THE FOREST.">

**Type `WEST`**

**482.**

<img src="images/timezone/0482.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN THE FOREST. THERE IS A TIME MACHINE HERE.">

**Type `GO MACHINE`**

**483.**

<img src="images/timezone/0483-1.png" width="400" alt="PLEASE INSERT DISK NUMBER 1B AND PRESS RETURN."><br>
<img src="images/timezone/0483-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0483.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

*Side 1B: drop `Time Zone (4am and san inc crack) disk B.dsk` on drive 1, and press Return.*

**Type `SIT`**

**484.**

<img src="images/timezone/0484-1.png" width="400" alt="O.K. THE CHAIR IS COMFORTABLE. AS YOU SIT DOWN, THE MACHINE SEEMS TO COME TO LIFE. YOU ARE INSIDE A TIME MACHINE. YOU HEAR"><br>
<img src="images/timezone/0484.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT. --------------- ENTER COMMAND">

**Type `AFRICA`**

**485.**

<img src="images/timezone/0485-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0485.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `50BC`**

**486.**

<img src="images/timezone/0486-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0486.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `PULL LEVER`**

**487.**

<img src="images/timezone/0487-1.png" width="400" alt="THE MACHINE STARTS VIBRATING VIOLENTLY AND YOU FEEL DIZZY, ALMOST TO THE POINT OF PASSING OUT. SUDDENLY YOU FEEL A TERRIBLE JOLT AND THE MACHINE IS STILL."><br>
<img src="images/timezone/0487-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0487.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EXIT MACHINE`**

### Africa, 50 BC

**488.**

<img src="images/timezone/0488-1.png" width="400" alt="NE PLEASE INSERT DISK NUMBER 2D AND PRESS RETURN."><br>
<img src="images/timezone/0488.png" width="400" alt="YOU ARE IN A DESERT ALONG THE NILE RIVER. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

*Side 2D: drop `Time Zone (4am and san inc crack) disk D.dsk` on drive 1, and press Return.*

**Type `SOUTH`**

**489.**

<img src="images/timezone/0489.png" width="400" alt="YOU ARE IN A DESERT ALONG THE NILE RIVER. THE NILE IS RUNNING NORTH AND SOUTH.">

**Type `SOUTH`**

**490.**

<img src="images/timezone/0490.png" width="400" alt="YOU ARE IN THE ANCIENT CITY OF THEBES, BY THE NILE RIVER. A ROAD GOES WEST AND SOUTH.">

**Type `SOUTH`**

**491.**

<img src="images/timezone/0491.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE ON THE OUTSKIRTS OF THEBES, ALONG THE NILE RIVER.">

**Type `SOUTH`**

**492.**

<img src="images/timezone/0492.png" width="400" alt="YOU ARE IN A DESERT ALONG THE NILE RIVER. THE NILE IS RUNNING NORTH AND SOUTH.">

**Type `WEST`**

**493.**

<img src="images/timezone/0493.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN FRONT OF CLEOPATRA&#39;S PALACE. THERE IS A GUARD STANDING HERE.">

**Type `NORTH`**

**494.**

<img src="images/timezone/0494.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN CLEOPATRA&#39;S THRONE ROOM. THERE IS A DOORWAY TO THE SOUTH.">

**Type `WEST`**

**495.**

<img src="images/timezone/0495.png" width="400" alt="YOU ARE IN CLEOPATRA&#39;S BATHING ROOM. CLEOPATRA IS TAKING A BATH IN MILK AND HONEY. SHE HAS SERVANTS HELPING HER.">

**Type `GIVE PERFUME`**

**496.**

<img src="images/timezone/0496-1.png" width="400" alt="CLEOPTRA LOVES PERFUME. SHE TAKES THE PERFUME AND GIVES YOU SOME EGYPTIAN MONEY. YOU ARE IN CLEOPATRA&#39;S BATHING ROOM."><br>
<img src="images/timezone/0496.png" width="400" alt="YOU ARE IN CLEOPATRA&#39;S BATHING ROOM. CLEOPATRA IS TAKING A BATH IN MILK AND HONEY. SHE HAS SERVANTS HELPING HER.">

**Type `EAST`**

**497.**

<img src="images/timezone/0497.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN CLEOPATRA&#39;S THRONE ROOM. THERE IS A DOORWAY TO THE SOUTH.">

**Type `SOUTH`**

**498.**

<img src="images/timezone/0498.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN FRONT OF CLEOPATRA&#39;S PALACE. THERE IS A GUARD STANDING HERE.">

**Type `EAST`**

**499.**

<img src="images/timezone/0499.png" width="400" alt="YOU ARE IN A DESERT ALONG THE NILE RIVER. THE NILE IS RUNNING NORTH AND SOUTH.">

**Type `NORTH`**

**500.**

<img src="images/timezone/0500.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE ON THE OUTSKIRTS OF THEBES, ALONG THE NILE RIVER.">

**Type `NORTH`**

**501.**

<img src="images/timezone/0501.png" width="400" alt="YOU ARE IN THE ANCIENT CITY OF THEBES, BY THE NILE RIVER. A ROAD GOES WEST AND SOUTH.">

**Type `WEST`**

**502.**

<img src="images/timezone/0502.png" width="400" alt="YOU ARE IN THE MARKETPLACE OF THEBES. A ROAD IS GOING NORTH AND EAST. A PEASANT WOMAN IS SELLING FRUIT.">

**Type `USE MONEY`**

**503.**

<img src="images/timezone/0503-1.png" width="400" alt="EGYPTIAN MONEY? YOU ARE IN THE MARKETPLACE OF THEBES. A ROAD IS GOING NORTH AND EAST. A PEASANT WOMAN IS SELLING FRUIT."><br>
<img src="images/timezone/0503.png" width="400" alt="YOU ARE IN THE MARKETPLACE OF THEBES. A ROAD IS GOING NORTH AND EAST. A PEASANT WOMAN IS SELLING FRUIT.">

**Type `YES`**

**504.**

<img src="images/timezone/0504-1.png" width="400" alt="THE PEASANT WOMAN TAKES THE EGYPTIAN MONEY AND GIVES YOU SOME DATES. YOU ARE IN THE MARKETPLACE OF THEBES. A ROAD IS GOING NORTH AND EAST. A PEASANT"><br>
<img src="images/timezone/0504.png" width="400" alt="YOU ARE IN THE MARKETPLACE OF THEBES. A ROAD IS GOING NORTH AND EAST. A PEASANT WOMAN IS SELLING FRUIT.">

**Type `EAST`**

**505.**

<img src="images/timezone/0505.png" width="400" alt="YOU ARE IN THE ANCIENT CITY OF THEBES, BY THE NILE RIVER. A ROAD GOES WEST AND SOUTH.">

**Type `NORTH`**

**506.**

<img src="images/timezone/0506.png" width="400" alt="YOU ARE IN A DESERT ALONG THE NILE RIVER. THE NILE IS RUNNING NORTH AND SOUTH.">

**Type `NORTH`**

**507.**

<img src="images/timezone/0507.png" width="400" alt="YOU ARE IN A DESERT ALONG THE NILE RIVER. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `NORTH`**

**508.**

<img src="images/timezone/0508-1.png" width="400" alt="YOU ARE GETTING VERY HUNGRY AND THIRSTY. YOU ARE IN A DESERT ALONG THE NILE RIVER. THE NILE IS RUNNING NORTH AND"><br>
<img src="images/timezone/0508.png" width="400" alt="YOU ARE IN A DESERT ALONG THE NILE RIVER. THE NILE IS RUNNING NORTH AND SOUTH.">

**Type `DRINK WATER`**

**509.**

<img src="images/timezone/0509-1.png" width="400" alt="AHHH! THAT WATER WAS REFRESHING. YOU ARE VERY HUNGRY. YOU ARE IN A DESERT ALONG THE NILE RIVER. THE NILE IS RUNNING NORTH AND"><br>
<img src="images/timezone/0509.png" width="400" alt="YOU ARE IN A DESERT ALONG THE NILE RIVER. THE NILE IS RUNNING NORTH AND SOUTH.">

**Type `EAT DATES`**

**510.**

<img src="images/timezone/0510-1.png" width="400" alt="MMMM! THOSE DATES JUST HIT THE SPOT. YOU ARE IN A DESERT ALONG THE NILE RIVER. THE NILE IS RUNNING NORTH AND SOUTH."><br>
<img src="images/timezone/0510.png" width="400" alt="YOU ARE IN A DESERT ALONG THE NILE RIVER. THE NILE IS RUNNING NORTH AND SOUTH.">

**Type `WEST`**

**511.**

<img src="images/timezone/0511.png" width="400" alt="SOUTH. --------------- ENTER COMMAND?WEST YOU ARE IN A DESERT.">

**Type `NORTH`**

**512.**

<img src="images/timezone/0512.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE WANDERING AIMLESSLY IN THE DESERT.">

**Type `NORTH`**

**513.**

<img src="images/timezone/0513.png" width="400" alt="DESERT. --------------- ENTER COMMAND?NORTH YOU ARE IN A HOT, DRY DESERT.">

**Type `WEST`**

**514.**

<img src="images/timezone/0514.png" width="400" alt="YOU ARE IN A HOT, DRY DESERT. --------------- ENTER COMMAND?WEST YOU ARE LOST IN A DESERT.">

**Type `NORTH`**

**515.**

<img src="images/timezone/0515.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE AT THE WEST SIDE OF A LARGE PYRAMID.">

**Type `NORTH`**

**516.**

<img src="images/timezone/0516.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE AT THE BACK SIDE OF A LARGE PYRAMID.">

**Type `UP`**

**517.**

<img src="images/timezone/0517.png" width="400" alt="--------------- ENTER COMMAND?UP YOU ARE HALFWAY UP A LARGE PYRAMID. THERE IS A DOOR HERE.">

**Type `OPEN DOOR`**

**518.**

<img src="images/timezone/0518.png" width="400" alt="O.K. YOU ARE HALFWAY UP A LARGE PYRAMID. THERE IS A DOOR HERE.">

**Type `GO DOOR`**

**519.**

<img src="images/timezone/0519-1.png" width="400" alt="YOU ARE IN A SMALL, CRAMPED ROOM JUST INSIDE THE PYRAMID. THERE IS A DOOR TO THE WEST AND STEEP STAIRS LEADING DOWNWARD."><br>
<img src="images/timezone/0519.png" width="400" alt="INSIDE THE PYRAMID. THERE IS A DOOR TO THE WEST AND STEEP STAIRS LEADING DOWNWARD.">

**Type `DOWN`**

**520.**

<img src="images/timezone/0520.png" width="400" alt="YOU ARE IN THE MIDDLE OF A STEEP STAIRWAY. THERE IS A SMALL DOORWAY TO THE WEST.">

**Type `DOWN`**

**521.**

<img src="images/timezone/0521.png" width="400" alt="YOU ARE IN THE MIDDLE OF A STEEP STAIRWAY. THERE IS A SMALL DOORWAY TO THE EAST.">

**Type `DOWN`**

**522.**

<img src="images/timezone/0522.png" width="400" alt="--------------- ENTER COMMAND?DOWN YOU ARE AT THE BOTTOM OF A STEEP STAIRWAY. THERE IS A LARGE STONE HERE.">

**Type `MOVE STONE`**

**523.**

<img src="images/timezone/0523-1.png" width="400" alt="BY MOVING THE STONE, YOU HAVE UNCOVERED A HOLE IN THE WALL. THERE IS A HOLE IN THE WALL. YOU ARE AT THE BOTTOM OF A STEEP"><br>
<img src="images/timezone/0523.png" width="400" alt="THERE IS A HOLE IN THE WALL. YOU ARE AT THE BOTTOM OF A STEEP STAIRWAY. THERE IS A LARGE STONE HERE.">

**Type `GO HOLE`**

**524.**

<img src="images/timezone/0524.png" width="400" alt="YOU ARE IN A SMALL, STIFLING TUNNEL. THERE IS A HOLE IN THE WALL TO THE WEST.">

**Type `EAST`**

**525.**

<img src="images/timezone/0525.png" width="400" alt="THERE IS A SHIELD HERE. YOU ARE IN AN EMPTY TOMB. IT LOOKS LIKE IT HAS BEEN ROBBED IN TIMES PAST.">

**Type `GET SHIELD`**

**526.**

<img src="images/timezone/0526.png" width="400" alt="YOU ARE IN AN EMPTY TOMB. IT LOOKS LIKE IT HAS BEEN ROBBED IN TIMES PAST.">

**Type `WEST`**

**527.**

<img src="images/timezone/0527.png" width="400" alt="YOU ARE IN A SMALL, STIFLING TUNNEL. THERE IS A HOLE IN THE WALL TO THE WEST.">

**Type `GO HOLE`**

**528.**

<img src="images/timezone/0528.png" width="400" alt="THERE IS A HOLE IN THE WALL. YOU ARE AT THE BOTTOM OF A STEEP STAIRWAY. THERE IS A LARGE STONE HERE.">

**Type `UP`**

**529.**

<img src="images/timezone/0529.png" width="400" alt="YOU ARE IN THE MIDDLE OF A STEEP STAIRWAY. THERE IS A SMALL DOORWAY TO THE EAST.">

**Type `UP`**

**530.**

<img src="images/timezone/0530.png" width="400" alt="YOU ARE IN THE MIDDLE OF A STEEP STAIRWAY. THERE IS A SMALL DOORWAY TO THE WEST.">

**Type `UP`**

**531.**

<img src="images/timezone/0531-1.png" width="400" alt="YOU ARE IN A SMALL, CRAMPED ROOM JUST INSIDE THE PYRAMID. THERE IS A DOOR TO THE WEST AND STEEP STAIRS LEADING DOWNWARD."><br>
<img src="images/timezone/0531.png" width="400" alt="INSIDE THE PYRAMID. THERE IS A DOOR TO THE WEST AND STEEP STAIRS LEADING DOWNWARD.">

**Type `GO DOOR`**

**532.**

<img src="images/timezone/0532.png" width="400" alt="--------------- ENTER COMMAND?GO DOOR YOU ARE HALFWAY UP A LARGE PYRAMID. THERE IS A DOOR HERE.">

**Type `DOWN`**

**533.**

<img src="images/timezone/0533.png" width="400" alt="--------------- ENTER COMMAND?DOWN YOU ARE AT THE BACK SIDE OF A LARGE PYRAMID.">

**Type `SOUTH`**

**534.**

<img src="images/timezone/0534.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE AT THE WEST SIDE OF A LARGE PYRAMID.">

**Type `SOUTH`**

**535.**

<img src="images/timezone/0535.png" width="400" alt="PYRAMID. --------------- ENTER COMMAND?SOUTH YOU ARE LOST IN A DESERT.">

**Type `EAST`**

**536.**

<img src="images/timezone/0536.png" width="400" alt="YOU ARE LOST IN A DESERT. --------------- ENTER COMMAND?EAST YOU ARE IN A HOT, DRY DESERT.">

**Type `SOUTH`**

**537.**

<img src="images/timezone/0537.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE WANDERING AIMLESSLY IN THE DESERT.">

**Type `SOUTH`**

**538.**

<img src="images/timezone/0538.png" width="400" alt="DESERT. --------------- ENTER COMMAND?SOUTH YOU ARE IN A DESERT.">

**Type `EAST`**

**539.**

<img src="images/timezone/0539.png" width="400" alt="YOU ARE IN A DESERT ALONG THE NILE RIVER. THE NILE IS RUNNING NORTH AND SOUTH.">

**Type `SOUTH`**

**540.**

<img src="images/timezone/0540.png" width="400" alt="YOU ARE IN A DESERT ALONG THE NILE RIVER. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `GO MACHINE`**

**541.**

<img src="images/timezone/0541-1.png" width="400" alt="PLEASE INSERT DISK NUMBER 1B AND PRESS RETURN."><br>
<img src="images/timezone/0541-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0541.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

*Side 1B: drop `Time Zone (4am and san inc crack) disk B.dsk` on drive 1, and press Return.*

**Type `SIT`**

**542.**

<img src="images/timezone/0542-1.png" width="400" alt="O.K. THE CHAIR IS COMFORTABLE. AS YOU SIT DOWN, THE MACHINE SEEMS TO COME TO LIFE. YOU ARE INSIDE A TIME MACHINE. YOU HEAR"><br>
<img src="images/timezone/0542.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT. --------------- ENTER COMMAND">

**Type `ASIA`**

**543.**

<img src="images/timezone/0543-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0543.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `1400AD`**

**544.**

<img src="images/timezone/0544-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0544.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `PULL LEVER`**

**545.**

<img src="images/timezone/0545-1.png" width="400" alt="THE MACHINE STARTS VIBRATING VIOLENTLY AND YOU FEEL DIZZY, ALMOST TO THE POINT OF PASSING OUT. SUDDENLY YOU FEEL A TERRIBLE JOLT AND THE MACHINE IS STILL."><br>
<img src="images/timezone/0545-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0545.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EXIT MACHINE`**

### Asia, 1400 AD

**546.**

<img src="images/timezone/0546-1.png" width="400" alt="NE PLEASE INSERT DISK NUMBER 4G AND PRESS RETURN."><br>
<img src="images/timezone/0546.png" width="400" alt="YOU ARE IN A FOREST. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

*Side 4G: drop `Time Zone (4am and san inc crack) disk G.dsk` on drive 1, and press Return.*

**Type `NORTH`**

**547.**

<img src="images/timezone/0547.png" width="400" alt="PULSATING. --------------- ENTER COMMAND?NORTH YOU ARE IN A CLEARING IN A FOREST.">

**Type `NORTH`**

**548.**

<img src="images/timezone/0548.png" width="400" alt="YOU ARE IN A CLEARING IN A FOREST. --------------- ENTER COMMAND?NORTH YOU ARE IN A SMALL RICE PADDY.">

**Type `NORTH`**

**549.**

<img src="images/timezone/0549.png" width="400" alt="YOU ARE IN A SMALL RICE PADDY. --------------- ENTER COMMAND?NORTH YOU ARE AT THE EDGE OF AN OCEAN.">

**Type `WEST`**

**550.**

<img src="images/timezone/0550.png" width="400" alt="OH NO! THERE IS A SAMURAI WARRIOR HERE.HE IS ABOUT TO KILL YOU. YOU ARE AT THE EDGE OF AN OCEAN.">

**Type `THROW BOOMERANG`**

**551.**

<img src="images/timezone/0551-1.png" width="400" alt="DEFTLY, YOU THROW THE BOOMERANG AT THE SAMURAI AND KILL HIM WITH A BLOW TO THE HEAD. THERE IS A DEAD SAMURAI WARRIOR HERE."><br>
<img src="images/timezone/0551.png" width="400" alt="THERE IS A DEAD SAMURAI WARRIOR HERE. THERE IS A SWORD HERE. YOU ARE AT THE EDGE OF AN OCEAN.">

**Type `GET SWORD`**

**552.**

<img src="images/timezone/0552.png" width="400" alt="--------------- ENTER COMMAND?GET SWORD THERE IS A DEAD SAMURAI WARRIOR HERE. YOU ARE AT THE EDGE OF AN OCEAN.">

**Type `EAST`**

**553.**

<img src="images/timezone/0553.png" width="400" alt="YOU ARE AT THE EDGE OF AN OCEAN. --------------- ENTER COMMAND?EAST YOU ARE AT THE EDGE OF AN OCEAN.">

**Type `EAST`**

**554.**

<img src="images/timezone/0554.png" width="400" alt="YOU ARE AT THE EDGE OF AN OCEAN. --------------- ENTER COMMAND?EAST YOU ARE AT THE EDGE OF AN OCEAN.">

**Type `EAST`**

**555.**

<img src="images/timezone/0555.png" width="400" alt="YOU ARE IN A SMALL FISHING VILLAGE ON THE EDGE OF AN OCEAN. A ROAD GOES SOUTH FROM HERE.">

**Type `SOUTH`**

**556.**

<img src="images/timezone/0556.png" width="400" alt="YOU ARE IN FRONT OF A SMALL SILK SHOP.THERE IS A JAPANESE LADY SELLING SILK HERE.">

**Type `TRADE RICE`**

**557.**

<img src="images/timezone/0557-1.png" width="400" alt="THE JAPANESE LADY LOVES RICE. SHE TAKES THE BAG OF RICE AND GIVES YOU A ROLL OF SILK. YOU ARE IN FRONT OF A SMALL SILK"><br>
<img src="images/timezone/0557.png" width="400" alt="YOU ARE IN FRONT OF A SMALL SILK SHOP.THERE IS A JAPANESE LADY SELLING SILK HERE.">

**Type `WEST`**

**558.**

<img src="images/timezone/0558.png" width="400" alt="SILK HERE. --------------- ENTER COMMAND?WEST YOU ARE IN A SMALL RICE PADDY.">

**Type `WEST`**

**559.**

<img src="images/timezone/0559.png" width="400" alt="YOU ARE IN A SMALL RICE PADDY. --------------- ENTER COMMAND?WEST YOU ARE IN A SMALL RICE PADDY.">

**Type `SOUTH`**

**560.**

<img src="images/timezone/0560.png" width="400" alt="YOU ARE IN A SMALL RICE PADDY. --------------- ENTER COMMAND?SOUTH YOU ARE IN A CLEARING IN A FOREST.">

**Type `SOUTH`**

**561.**

<img src="images/timezone/0561.png" width="400" alt="YOU ARE IN A FOREST. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `GO MACHINE`**

**562.**

<img src="images/timezone/0562-1.png" width="400" alt="PLEASE INSERT DISK NUMBER 1B AND PRESS RETURN."><br>
<img src="images/timezone/0562-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0562.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

*Side 1B: drop `Time Zone (4am and san inc crack) disk B.dsk` on drive 1, and press Return.*

**Type `SIT`**

**563.**

<img src="images/timezone/0563-1.png" width="400" alt="O.K. THE CHAIR IS COMFORTABLE. AS YOU SIT DOWN, THE MACHINE SEEMS TO COME TO LIFE. YOU ARE INSIDE A TIME MACHINE. YOU HEAR"><br>
<img src="images/timezone/0563.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT. --------------- ENTER COMMAND">

**Type `EUROPE`**

**564.**

<img src="images/timezone/0564-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0564.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `50BC`**

**565.**

<img src="images/timezone/0565-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0565.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `PULL LEVER`**

**566.**

<img src="images/timezone/0566-1.png" width="400" alt="THE MACHINE STARTS VIBRATING VIOLENTLY AND YOU FEEL DIZZY, ALMOST TO THE POINT OF PASSING OUT. SUDDENLY YOU FEEL A TERRIBLE JOLT AND THE MACHINE IS STILL."><br>
<img src="images/timezone/0566-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0566.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EXIT MACHINE`**

### Europe, 50 BC

**567.**

<img src="images/timezone/0567-1.png" width="400" alt="NE PLEASE INSERT DISK NUMBER 2C AND PRESS RETURN."><br>
<img src="images/timezone/0567.png" width="400" alt="YOUR ARE IN THE HILLS. THERE IS A TIME MACHINE HERE, IT APPEARS TO BE PULSATING.">

*Side 2C: drop `Time Zone (4am and san inc crack) disk C.dsk` on drive 1, and press Return.*

**Type `EAST`**

**568.**

<img src="images/timezone/0568.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE ON TOP OF A HILL. YOU SEE ROME IN THE DISTANCE TO THE NORTHEAST.">

**Type `NORTH`**

**569.**

<img src="images/timezone/0569.png" width="400" alt="IN THE DISTANCE TO THE NORTHEAST. --------------- ENTER COMMAND?NORTH YOU ARE IN AN OLIVE ORCHARD.">

**Type `NORTH`**

**570.**

<img src="images/timezone/0570.png" width="400" alt="YOU ARE IN AN OLIVE ORCHARD. --------------- ENTER COMMAND?NORTH YOU ARE IN AN OLIVE ORCHARD.">

**Type `EAST`**

**571.**

<img src="images/timezone/0571.png" width="400" alt="YOU ARE IN AN OLIVE ORCHARD. --------------- ENTER COMMAND?EAST YOU ARE IN AN OLIVE ORCHARD.">

**Type `NORTH`**

**572.**

<img src="images/timezone/0572.png" width="400" alt="YOU ARE IN AN OLIVE ORCHARD. --------------- ENTER COMMAND?NORTH YOU ARE IN AN OLIVE ORCHARD.">

**Type `NORTH`**

**573.**

<img src="images/timezone/0573.png" width="400" alt="YOU ARE IN AN OLIVE ORCHARD. --------------- ENTER COMMAND?NORTH YOU ARE ON A STONE ROAD LEADING NORTH.">

**Type `NORTH`**

**574.**

<img src="images/timezone/0574.png" width="400" alt="YOU ARE STANDING IN FRONT OF THE ROMAN COLOSSEUM. THERE IS AN OPEN DOORWAY TO THE NORTH.">

**Type `NORTH`**

**575.**

<img src="images/timezone/0575.png" width="400" alt="YOU ARE IN THE CORRIDOR OF THE ROMAN COLOSSEUM. THERE IS A DOORWAY TO THE NORTH.">

**Type `WEST`**

**576.**

<img src="images/timezone/0576.png" width="400" alt="YOU ARE IN THE CORRIDOR OF THE ROMAN COLOSSEUM. THERE IS A CAGE HERE WITH PRISONERS INSIDE.">

**Type `NORTH`**

**577.**

<img src="images/timezone/0577.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN THE CORRIDOR OF THE ROMAN COLOSSEUM. THERE IS A PIT GOING DOWN.">

**Type `DOWN`**

**578.**

<img src="images/timezone/0578.png" width="400" alt="--------------- ENTER COMMAND?DOWN YOU ARE IN A LABYRINTH OF TUNNELS. THERE IS A HOLE IN THE CEILING.">

**Type `WEST`**

**579.**

<img src="images/timezone/0579.png" width="400" alt="THERE IS A HOLE IN THE CEILING. --------------- ENTER COMMAND?WEST YOU ARE IN A LABYRINTH OF TUNNELS.">

**Type `SOUTH`**

**580.**

<img src="images/timezone/0580.png" width="400" alt="YOU ARE IN A LABYRINTH OF TUNNELS. --------------- ENTER COMMAND?SOUTH YOU ARE IN A LABYRINTH OF TUNNELS.">

**Type `SOUTH`**

**581.**

<img src="images/timezone/0581.png" width="400" alt="YOU ARE IN A LABYRINTH OF TUNNELS. --------------- ENTER COMMAND?SOUTH YOU ARE IN A LABYRINTH OF TUNNELS.">

**Type `SOUTH`**

**582.**

<img src="images/timezone/0582.png" width="400" alt="YOU ARE IN A LABYRINTH OF TUNNELS. --------------- ENTER COMMAND?SOUTH YOU ARE IN A LABYRINTH OF TUNNELS.">

**Type `EAST`**

**583.**

<img src="images/timezone/0583.png" width="400" alt="YOU ARE IN A LABYRINTH OF TUNNELS. --------------- ENTER COMMAND?EAST YOU ARE IN A LABYRINTH OF TUNNELS.">

**Type `SOUTH`**

**584.**

<img src="images/timezone/0584.png" width="400" alt="THERE IS A PAIR OF TWEEZERS HERE. YOU ARE IN A LABYRINTH OF TUNNELS. THE TUNNEL DEAD-ENDS TO THE SOUTH.">

**Type `GET TWEEZERS`**

**585.**

<img src="images/timezone/0585.png" width="400" alt="RS YOU ARE IN A LABYRINTH OF TUNNELS. THE TUNNEL DEAD-ENDS TO THE SOUTH.">

**Type `NORTH`**

**586.**

<img src="images/timezone/0586.png" width="400" alt="TUNNEL DEAD-ENDS TO THE SOUTH. --------------- ENTER COMMAND?NORTH YOU ARE IN A LABYRINTH OF TUNNELS.">

**Type `NORTH`**

**587.**

<img src="images/timezone/0587.png" width="400" alt="YOU ARE IN A LABYRINTH OF TUNNELS. --------------- ENTER COMMAND?NORTH YOU ARE IN A LABYRINTH OF TUNNELS.">

**Type `WEST`**

**588.**

<img src="images/timezone/0588.png" width="400" alt="YOU ARE IN A LABYRINTH OF TUNNELS. --------------- ENTER COMMAND?WEST YOU ARE IN A LABYRINTH OF TUNNELS.">

**Type `NORTH`**

**589.**

<img src="images/timezone/0589.png" width="400" alt="YOU ARE IN A LABYRINTH OF TUNNELS. --------------- ENTER COMMAND?NORTH YOU ARE IN A LABYRINTH OF TUNNELS.">

**Type `EAST`**

**590.**

<img src="images/timezone/0590.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN A LABYRINTH OF TUNNELS. THERE IS A HOLE IN THE CEILING.">

**Type `UP`**

**591.**

<img src="images/timezone/0591.png" width="400" alt="--------------- ENTER COMMAND?UP YOU ARE IN THE CORRIDOR OF THE ROMAN COLOSSEUM. THERE IS A PIT GOING DOWN.">

**Type `EAST`**

**592.**

<img src="images/timezone/0592.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN THE CORRIDOR OF THE ROMAN COLOSSEUM.">

**Type `LOOK`**

**593.**

<img src="images/timezone/0593.png" width="400" alt="THE TWO GUARDS GRAB YOU AND THROW YOU IN THE LION CAGE. YOU ARE IN A LION&#39;S CAGE.">

**Type `USE TWEEZERS`**

**594.**

<img src="images/timezone/0594-1.png" width="400" alt="YOU GRAB THE THORN WITH THE TWEEZERS AND PULL IT OUT OF THE LION&#39;S PAW,YOU THEN DISCARD IT. THE LION MOVES TO THE OTHER SIDE OF THE CAGE."><br>
<img src="images/timezone/0594.png" width="400" alt="THEN DISCARD IT. THE LION MOVES TO THE OTHER SIDE OF THE CAGE. YOU ARE IN A LION&#39;S CAGE.">

**Type `OPEN DOOR`**

**595.**

<img src="images/timezone/0595.png" width="400" alt="--------------- ENTER COMMAND?OPEN DOOR O.K. YOU ARE IN A LION&#39;S CAGE.">

**Type `WEST`**

**596.**

<img src="images/timezone/0596.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN THE ARENA OF THE ROMAN COLOSSEUM.">

**Type `DROP TWEEZERS`**

**597.**

<img src="images/timezone/0597.png" width="400" alt="THERE IS A PAIR OF TWEEZERS HERE. YOU ARE IN THE ARENA OF THE ROMAN COLOSSEUM.">

**Type `SOUTH`**

**598.**

<img src="images/timezone/0598.png" width="400" alt="YOU ARE IN THE CORRIDOR OF THE ROMAN COLOSSEUM. THERE IS A DOORWAY TO THE NORTH.">

**Type `WEST`**

**599.**

<img src="images/timezone/0599.png" width="400" alt="YOU ARE IN THE CORRIDOR OF THE ROMAN COLOSSEUM. THERE IS A CAGE HERE WITH PRISONERS INSIDE.">

**Type `NORTH`**

**600.**

<img src="images/timezone/0600.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN THE CORRIDOR OF THE ROMAN COLOSSEUM. THERE IS A PIT GOING DOWN.">

**Type `EAST`**

**601.**

<img src="images/timezone/0601.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN THE CORRIDOR OF THE ROMAN COLOSSEUM.">

**Type `EAST`**

**602.**

<img src="images/timezone/0602.png" width="400" alt="YOU ARE IN THE CORRIDOR OF THE ROMAN COLOSSEUM. THERE ARE STAIRS GOING UP TO THE EAST.">

**Type `UP`**

**603.**

<img src="images/timezone/0603.png" width="400" alt="--------------- ENTER COMMAND?UP YOU ARE IN AN UPPER LEVEL CORRIDOR OF THE ROMAN COLOSSEUM.">

**Type `EAST`**

**604.**

<img src="images/timezone/0604.png" width="400" alt="YOU ARE IN AN UPPER LEVEL CORRIDOR OF THE ROMAN COLOSSEUM. THERE ARE TWO GUARDS HERE.">

**Type `LOOK`**

**605.**

<img src="images/timezone/0605-1.png" width="400" alt="THE GUARDS GRAB YOU AND THROW YOU INTO THE ARENA TO FIGHT A GLADIATOR. WITH YOUR SWORD AND SHIELD YOU QUICKLY MAKE MINCEMEAT OF HIM. YOU ARE PROCLAIMED A"><br>
<img src="images/timezone/0605-2.png" width="400" alt="HERO AND RECEIVE AN INVITATION TO VISIT JULIUS CAESAR. THERE IS A PAIR OF TWEEZERS HERE. YOU ARE IN THE ARENA OF THE ROMAN"><br>
<img src="images/timezone/0605.png" width="400" alt="THERE IS A PAIR OF TWEEZERS HERE. YOU ARE IN THE ARENA OF THE ROMAN COLOSSEUM.">

**Type `SOUTH`**

**606.**

<img src="images/timezone/0606.png" width="400" alt="YOU ARE IN THE CORRIDOR OF THE ROMAN COLOSSEUM. THERE IS A DOORWAY TO THE NORTH.">

**Type `SOUTH`**

**607.**

<img src="images/timezone/0607.png" width="400" alt="YOU ARE STANDING IN FRONT OF THE ROMAN COLOSSEUM. THERE IS AN OPEN DOORWAY TO THE NORTH.">

**Type `EAST`**

**608.**

<img src="images/timezone/0608.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN FRONT OF THE ROMAN SENATE BUILDING. THERE ARE TWO GUARDS HERE.">

**Type `NORTH`**

**609.**

<img src="images/timezone/0609.png" width="400" alt="YOU ARE IN THE ENTRY ROOM OF THE ROMAN SENATE BUILDING. THERE ARE DOORWAYS TO THE EAST, WEST AND SOUTH.">

**Type `WEST`**

**610.**

<img src="images/timezone/0610.png" width="400" alt="THERE IS A LADDER HERE. YOU ARE IN THE LIBRARY OF THE ROMAN SENATE BUILDING.">

**Type `GET LADDER`**

**611.**

<img src="images/timezone/0611.png" width="400" alt="YOU ARE IN THE LIBRARY OF THE ROMAN SENATE BUILDING.">

**Type `EAST`**

**612.**

<img src="images/timezone/0612.png" width="400" alt="YOU ARE IN THE ENTRY ROOM OF THE ROMAN SENATE BUILDING. THERE ARE DOORWAYS TO THE EAST, WEST AND SOUTH.">

**Type `SOUTH`**

**613.**

<img src="images/timezone/0613.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN FRONT OF THE ROMAN SENATE BUILDING. THERE ARE TWO GUARDS HERE.">

**Type `WEST`**

**614.**

<img src="images/timezone/0614.png" width="400" alt="YOU ARE STANDING IN FRONT OF THE ROMAN COLOSSEUM. THERE IS AN OPEN DOORWAY TO THE NORTH.">

**Type `SOUTH`**

**615.**

<img src="images/timezone/0615.png" width="400" alt="THE NORTH. --------------- ENTER COMMAND?SOUTH YOU ARE ON A STONE ROAD LEADING NORTH.">

**Type `SOUTH`**

**616.**

<img src="images/timezone/0616.png" width="400" alt="YOU ARE ON A STONE ROAD LEADING NORTH. --------------- ENTER COMMAND?SOUTH YOU ARE IN AN OLIVE ORCHARD.">

**Type `SOUTH`**

**617.**

<img src="images/timezone/0617.png" width="400" alt="YOU ARE IN AN OLIVE ORCHARD. --------------- ENTER COMMAND?SOUTH YOU ARE IN AN OLIVE ORCHARD.">

**Type `WEST`**

**618.**

<img src="images/timezone/0618.png" width="400" alt="YOU ARE IN AN OLIVE ORCHARD. --------------- ENTER COMMAND?WEST YOU ARE IN AN OLIVE ORCHARD.">

**Type `SOUTH`**

**619.**

<img src="images/timezone/0619.png" width="400" alt="YOU ARE IN AN OLIVE ORCHARD. --------------- ENTER COMMAND?SOUTH YOU ARE IN AN OLIVE ORCHARD.">

**Type `SOUTH`**

**620.**

<img src="images/timezone/0620.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE ON TOP OF A HILL. YOU SEE ROME IN THE DISTANCE TO THE NORTHEAST.">

**Type `WEST`**

**621.**

<img src="images/timezone/0621.png" width="400" alt="YOUR ARE IN THE HILLS. THERE IS A TIME MACHINE HERE, IT APPEARS TO BE PULSATING.">

**Type `GO MACHINE`**

**622.**

<img src="images/timezone/0622-1.png" width="400" alt="PLEASE INSERT DISK NUMBER 1B AND PRESS RETURN."><br>
<img src="images/timezone/0622-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0622.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

*Side 1B: drop `Time Zone (4am and san inc crack) disk B.dsk` on drive 1, and press Return.*

**Type `SIT`**

**623.**

<img src="images/timezone/0623-1.png" width="400" alt="O.K. THE CHAIR IS COMFORTABLE. AS YOU SIT DOWN, THE MACHINE SEEMS TO COME TO LIFE. YOU ARE INSIDE A TIME MACHINE. YOU HEAR"><br>
<img src="images/timezone/0623.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT. --------------- ENTER COMMAND">

**Type `AFRICA`**

**624.**

<img src="images/timezone/0624-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0624.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `1000AD`**

**625.**

<img src="images/timezone/0625-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0625.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `PULL LEVER`**

**626.**

<img src="images/timezone/0626-1.png" width="400" alt="THE MACHINE STARTS VIBRATING VIOLENTLY AND YOU FEEL DIZZY, ALMOST TO THE POINT OF PASSING OUT. SUDDENLY YOU FEEL A TERRIBLE JOLT AND THE MACHINE IS STILL."><br>
<img src="images/timezone/0626-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0626.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EXIT MACHINE`**

### Africa, 1000 AD

**627.**

<img src="images/timezone/0627-1.png" width="400" alt="NE PLEASE INSERT DISK NUMBER 2D AND PRESS RETURN."><br>
<img src="images/timezone/0627.png" width="400" alt="YOU ARE IN A JUNGLE. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

*Side 2D: drop `Time Zone (4am and san inc crack) disk D.dsk` on drive 1, and press Return.*

**Type `NORTH`**

**628.**

<img src="images/timezone/0628.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN A JUNGLE. YOU ARE HOT AND STICKY.">

**Type `NORTH`**

**629.**

<img src="images/timezone/0629-1.png" width="400" alt="THERE IS A LOG HERE. YOU ARE ON THE SOUTH SIDE OF THE CONGO RIVER. THE RIVER IS RUNNING EAST AND WEST."><br>
<img src="images/timezone/0629.png" width="400" alt="YOU ARE ON THE SOUTH SIDE OF THE CONGO RIVER. THE RIVER IS RUNNING EAST AND WEST.">

**Type `WEST`**

**630.**

<img src="images/timezone/0630-1.png" width="400" alt="THERE IS A LOG HERE. YOU ARE ON THE SOUTH SIDE OF THE CONGO RIVER. THE RIVER IS RUNNING EAST AND WEST."><br>
<img src="images/timezone/0630.png" width="400" alt="YOU ARE ON THE SOUTH SIDE OF THE CONGO RIVER. THE RIVER IS RUNNING EAST AND WEST.">

**Type `GET LOG`**

**631.**

<img src="images/timezone/0631.png" width="400" alt="YOU ARE ON THE SOUTH SIDE OF THE CONGO RIVER. THE RIVER IS RUNNING EAST AND WEST.">

**Type `EAST`**

**632.**

<img src="images/timezone/0632-1.png" width="400" alt="THERE IS A LOG HERE. YOU ARE ON THE SOUTH SIDE OF THE CONGO RIVER. THE RIVER IS RUNNING EAST AND WEST."><br>
<img src="images/timezone/0632.png" width="400" alt="YOU ARE ON THE SOUTH SIDE OF THE CONGO RIVER. THE RIVER IS RUNNING EAST AND WEST.">

**Type `DROP LOG`**

**633.**

<img src="images/timezone/0633-1.png" width="400" alt="THERE ARE TWO LOGS HERE. YOU ARE ON THE SOUTH SIDE OF THE CONGO RIVER. THE RIVER IS RUNNING EAST AND WEST."><br>
<img src="images/timezone/0633.png" width="400" alt="YOU ARE ON THE SOUTH SIDE OF THE CONGO RIVER. THE RIVER IS RUNNING EAST AND WEST.">

**Type `EAST`**

**634.**

<img src="images/timezone/0634-1.png" width="400" alt="THERE IS A LOG HERE. YOU ARE ON THE SOUTH SIDE OF THE CONGO RIVER. THE RIVER IS RUNNING EAST AND WEST."><br>
<img src="images/timezone/0634.png" width="400" alt="YOU ARE ON THE SOUTH SIDE OF THE CONGO RIVER. THE RIVER IS RUNNING EAST AND WEST.">

**Type `GET LOG`**

**635.**

<img src="images/timezone/0635.png" width="400" alt="YOU ARE ON THE SOUTH SIDE OF THE CONGO RIVER. THE RIVER IS RUNNING EAST AND WEST.">

**Type `WEST`**

**636.**

<img src="images/timezone/0636-1.png" width="400" alt="THERE ARE TWO LOGS HERE. YOU ARE ON THE SOUTH SIDE OF THE CONGO RIVER. THE RIVER IS RUNNING EAST AND WEST."><br>
<img src="images/timezone/0636.png" width="400" alt="YOU ARE ON THE SOUTH SIDE OF THE CONGO RIVER. THE RIVER IS RUNNING EAST AND WEST.">

**Type `DROP LOG`**

**637.**

<img src="images/timezone/0637-1.png" width="400" alt="THERE ARE THREE LOGS HERE. YOU ARE ON THE SOUTH SIDE OF THE CONGO RIVER. THE RIVER IS RUNNING EAST AND WEST."><br>
<img src="images/timezone/0637.png" width="400" alt="YOU ARE ON THE SOUTH SIDE OF THE CONGO RIVER. THE RIVER IS RUNNING EAST AND WEST.">

**Type `MAKE RAFT`**

**638.**

<img src="images/timezone/0638-1.png" width="400" alt="O.K. YOU ARE ON THE SOUTH SIDE OF THE CONGO RIVER. THE RIVER IS RUNNING EAST AND WEST."><br>
<img src="images/timezone/0638.png" width="400" alt="YOU ARE ON THE SOUTH SIDE OF THE CONGO RIVER. THE RIVER IS RUNNING EAST AND WEST.">

**Type `CROSS RIVER`**

**639.**

<img src="images/timezone/0639-1.png" width="400" alt="YOU ARE IN THE MIDDLE OF THE CONGO RIVER ON A SMALL, FLIMSY RAFT. USING THE LONG POLE, YOU ARE ABLE TO GUIDE THE RAFT ACROSS THE RIVER."><br>
<img src="images/timezone/0639.png" width="400" alt="RIVER ON A SMALL, FLIMSY RAFT. USING THE LONG POLE, YOU ARE ABLE TO GUIDE THE RAFT ACROSS THE RIVER.">

**Type `NORTH`**

**640.**

<img src="images/timezone/0640.png" width="400" alt="YOU ARE NEAR THE NORTH SHORE OF THE CONGO RIVER. THE RIVER IS RUNNING EAST AND WEST.">

**Type `EAST`**

**641.**

<img src="images/timezone/0641.png" width="400" alt="YOU ARE NEAR THE NORTH SIDE OF THE CONGO RIVER. THE RIVER IS RUNNING EAST AND WEST.">

**Type `NORTH`**

**642.**

<img src="images/timezone/0642.png" width="400" alt="AND WEST. --------------- ENTER COMMAND?NORTH YOU ARE IN A STEAMY JUNGLE.">

**Type `NORTH`**

**643.**

<img src="images/timezone/0643.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN A JUNGLE. YOU ARE HOT AND STICKY.">

**Type `NORTH`**

**644.**

<img src="images/timezone/0644.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU HAVE FALLEN INTO A DEEP PIT. THERE IS THE SKELETON OF AN ELEPHANT HERE.">

**Type `GET TUSKS`**

**645.**

<img src="images/timezone/0645.png" width="400" alt="WITH WHAT? YOU HAVE FALLEN INTO A DEEP PIT. THERE IS THE SKELETON OF AN ELEPHANT HERE.">

**Type `WITH HAMMER`**

**646.**

<img src="images/timezone/0646-1.png" width="400" alt="YOU BREAK OFF THE TUSKS WITH THE STONE HAMMER AND TAKE THEM. YOU HAVE FALLEN INTO A DEEP PIT. THERE IS THE SKELETON OF AN ELEPHANT HERE."><br>
<img src="images/timezone/0646.png" width="400" alt="HAMMER AND TAKE THEM. YOU HAVE FALLEN INTO A DEEP PIT. THERE IS THE SKELETON OF AN ELEPHANT HERE.">

**Type `USE LADDER`**

**647.**

<img src="images/timezone/0647-1.png" width="400" alt="USING THE LADDER, YOU CLIMB OUT OF THE PIT YOU ARE IN A JUNGLE. YOU ARE HOT AND STICKY."><br>
<img src="images/timezone/0647.png" width="400" alt="PIT YOU ARE IN A JUNGLE. YOU ARE HOT AND STICKY.">

**Type `SOUTH`**

**648.**

<img src="images/timezone/0648.png" width="400" alt="STICKY. --------------- ENTER COMMAND?SOUTH YOU ARE IN A STEAMY JUNGLE.">

**Type `SOUTH`**

**649.**

<img src="images/timezone/0649.png" width="400" alt="YOU ARE NEAR THE NORTH SIDE OF THE CONGO RIVER. THE RIVER IS RUNNING EAST AND WEST.">

**Type `WEST`**

**650.**

<img src="images/timezone/0650.png" width="400" alt="YOU ARE NEAR THE NORTH SHORE OF THE CONGO RIVER. THE RIVER IS RUNNING EAST AND WEST.">

**Type `CROSS RIVER`**

**651.**

<img src="images/timezone/0651-1.png" width="400" alt="YOU ARE IN THE MIDDLE OF THE CONGO RIVER ON A SMALL, FLIMSY RAFT. USING THE LONG POLE, YOU ARE ABLE TO GUIDE THE RAFT ACROSS THE RIVER."><br>
<img src="images/timezone/0651.png" width="400" alt="RIVER ON A SMALL, FLIMSY RAFT. USING THE LONG POLE, YOU ARE ABLE TO GUIDE THE RAFT ACROSS THE RIVER.">

**Type `SOUTH`**

**652.**

<img src="images/timezone/0652.png" width="400" alt="YOU ARE ON THE SOUTH SIDE OF THE CONGO RIVER. THE RIVER IS RUNNING EAST AND WEST.">

**Type `GET ROPE`**

**653.**

<img src="images/timezone/0653-1.png" width="400" alt="THERE ARE THREE LOGS HERE. YOU ARE ON THE SOUTH SIDE OF THE CONGO RIVER. THE RIVER IS RUNNING EAST AND WEST."><br>
<img src="images/timezone/0653.png" width="400" alt="YOU ARE ON THE SOUTH SIDE OF THE CONGO RIVER. THE RIVER IS RUNNING EAST AND WEST.">

**Type `SOUTH`**

**654.**

<img src="images/timezone/0654.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN A JUNGLE. YOU ARE HOT AND STICKY.">

**Type `SOUTH`**

**655.**

<img src="images/timezone/0655.png" width="400" alt="YOU ARE IN A JUNGLE. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `GO MACHINE`**

**656.**

<img src="images/timezone/0656-1.png" width="400" alt="PLEASE INSERT DISK NUMBER 1B AND PRESS RETURN."><br>
<img src="images/timezone/0656-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0656.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

*Side 1B: drop `Time Zone (4am and san inc crack) disk B.dsk` on drive 1, and press Return.*

**Type `SIT`**

**657.**

<img src="images/timezone/0657-1.png" width="400" alt="O.K. THE CHAIR IS COMFORTABLE. AS YOU SIT DOWN, THE MACHINE SEEMS TO COME TO LIFE. YOU ARE INSIDE A TIME MACHINE. YOU HEAR"><br>
<img src="images/timezone/0657.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT. --------------- ENTER COMMAND">

**Type `AFRICA`**

**658.**

<img src="images/timezone/0658-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0658.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `1400AD`**

**659.**

<img src="images/timezone/0659-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0659.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `PULL LEVER`**

**660.**

<img src="images/timezone/0660-1.png" width="400" alt="THE MACHINE STARTS VIBRATING VIOLENTLY AND YOU FEEL DIZZY, ALMOST TO THE POINT OF PASSING OUT. SUDDENLY YOU FEEL A TERRIBLE JOLT AND THE MACHINE IS STILL."><br>
<img src="images/timezone/0660-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0660.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EXIT MACHINE`**

### Africa, 1400 AD

**661.**

<img src="images/timezone/0661-1.png" width="400" alt="NE PLEASE INSERT DISK NUMBER 2D AND PRESS RETURN."><br>
<img src="images/timezone/0661.png" width="400" alt="YOU ARE IN A DESERT. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

*Side 2D: drop `Time Zone (4am and san inc crack) disk D.dsk` on drive 1, and press Return.*

**Type `NORTH`**

**662.**

<img src="images/timezone/0662.png" width="400" alt="PULSATING. --------------- ENTER COMMAND?NORTH YOU ARE IN A VERY HOT DESERT.">

**Type `NORTH`**

**663.**

<img src="images/timezone/0663.png" width="400" alt="YOU ARE IN A VERY HOT DESERT. --------------- ENTER COMMAND?NORTH YOU ARE IN A HOT DESERT.">

**Type `NORTH`**

**664.**

<img src="images/timezone/0664.png" width="400" alt="YOU ARE IN A HOT DESERT. --------------- ENTER COMMAND?NORTH YOU ARE IN A HOT, DRY DESERT.">

**Type `NORTH`**

**665.**

<img src="images/timezone/0665.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN A MORROCCAN BAZAAR. YOU SEE AN ARAB MERCHANT HERE.">

**Type `TRADE TUSKS`**

**666.**

<img src="images/timezone/0666-1.png" width="400" alt="THE MERCHANT INSPECTS THE IVORY, SEEMS PLEASED WITH IT AND GIVES YOU THE KNIFE. YOU ARE IN A MORROCCAN BAZAAR. YOU SEE"><br>
<img src="images/timezone/0666.png" width="400" alt="KNIFE. YOU ARE IN A MORROCCAN BAZAAR. YOU SEE AN ARAB MERCHANT HERE.">

**Type `NORTH`**

**667.**

<img src="images/timezone/0667.png" width="400" alt="AN ARAB MERCHANT HERE. --------------- ENTER COMMAND?NORTH YOU ARE IN A PARCHED DESERT.">

**Type `SOUTH`**

**668.**

<img src="images/timezone/0668.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN A MORROCCAN BAZAAR. YOU SEE AN ARAB MERCHANT HERE.">

**Type `TRADE SILK`**

**669.**

<img src="images/timezone/0669-1.png" width="400" alt="THE MERCHANT INSPECTS THE SILK, SEEMS PLEASED WITH IT AND GIVES YOU THE PERSIAN RUG. YOU ARE IN A MORROCCAN BAZAAR. YOU SEE"><br>
<img src="images/timezone/0669.png" width="400" alt="PERSIAN RUG. YOU ARE IN A MORROCCAN BAZAAR. YOU SEE AN ARAB MERCHANT HERE.">

**Type `SOUTH`**

**670.**

<img src="images/timezone/0670.png" width="400" alt="AN ARAB MERCHANT HERE. --------------- ENTER COMMAND?SOUTH YOU ARE IN A HOT, DRY DESERT.">

**Type `SOUTH`**

**671.**

<img src="images/timezone/0671.png" width="400" alt="YOU ARE IN A HOT, DRY DESERT. --------------- ENTER COMMAND?SOUTH YOU ARE IN A HOT DESERT.">

**Type `SOUTH`**

**672.**

<img src="images/timezone/0672.png" width="400" alt="YOU ARE GETTING VERY HOT AND THIRSTY. I THINK YOU HAD BETTER DRINK SOME WATER. YOU ARE IN A VERY HOT DESERT.">

**Type `SOUTH`**

**673.**

<img src="images/timezone/0673-1.png" width="400" alt="YOU ARE GETTING VERY HOT AND THIRSTY. I THINK YOU HAD BETTER DRINK SOME WATER. YOU ARE IN A DESERT. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE"><br>
<img src="images/timezone/0673.png" width="400" alt="YOU ARE IN A DESERT. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `GO MACHINE`**

**674.**

<img src="images/timezone/0674-1.png" width="400" alt="PLEASE INSERT DISK NUMBER 1B AND PRESS RETURN."><br>
<img src="images/timezone/0674-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0674.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

*Side 1B: drop `Time Zone (4am and san inc crack) disk B.dsk` on drive 1, and press Return.*

**Type `SIT`**

**675.**

<img src="images/timezone/0675-1.png" width="400" alt="O.K. THE CHAIR IS COMFORTABLE. AS YOU SIT DOWN, THE MACHINE SEEMS TO COME TO LIFE. YOU ARE INSIDE A TIME MACHINE. YOU HEAR"><br>
<img src="images/timezone/0675.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT. --------------- ENTER COMMAND">

**Type `ASIA`**

**676.**

<img src="images/timezone/0676-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0676.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `1000AD`**

**677.**

<img src="images/timezone/0677-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0677.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `PULL LEVER`**

**678.**

<img src="images/timezone/0678-1.png" width="400" alt="THE MACHINE STARTS VIBRATING VIOLENTLY AND YOU FEEL DIZZY, ALMOST TO THE POINT OF PASSING OUT. SUDDENLY YOU FEEL A TERRIBLE JOLT AND THE MACHINE IS STILL."><br>
<img src="images/timezone/0678-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0678.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EXIT MACHINE`**

### Asia, 1000 AD

**679.**

<img src="images/timezone/0679-1.png" width="400" alt="NE PLEASE INSERT DISK NUMBER 4G AND PRESS RETURN."><br>
<img src="images/timezone/0679.png" width="400" alt="YOU ARE IN A DESERT. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

*Side 4G: drop `Time Zone (4am and san inc crack) disk G.dsk` on drive 1, and press Return.*

**Type `EAST`**

**680.**

<img src="images/timezone/0680.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN A DESERT. YOU SEE THE WALLED CITY OF BAGHDAD IN THE DISTANCE.">

**Type `NORTH`**

**681.**

<img src="images/timezone/0681.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN FRONT OF THE WALLED CITY OF BAGHDAD.">

**Type `NORTH`**

**682.**

<img src="images/timezone/0682.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE INSIDE THE WALLED CITY OF BAGHDAD.">

**Type `EAST`**

**683.**

<img src="images/timezone/0683.png" width="400" alt="YOU ARE IN THE MARKETPLACE OF BAGHDAD. THERE IS A MERCHANT SELLING CAMELS HERE.">

**Type `TRADE RUG`**

**684.**

<img src="images/timezone/0684-1.png" width="400" alt="THE MERCHANT LOOKS AT THE PERSIAN RUG AND IS VERY PLEASED. HE TAKES THE RUG AND GIVES YOU A CAMEL. THERE IS A CAMEL FOLLOWING YOU."><br>
<img src="images/timezone/0684.png" width="400" alt="YOU ARE IN THE MARKETPLACE OF BAGHDAD. THERE IS A MERCHANT SELLING CAMELS HERE.">

**Type `RIDE CAMEL`**

**685.**

<img src="images/timezone/0685-1.png" width="400" alt="YOU ARE RIDING A CAMEL. YOU ARE IN THE MARKETPLACE OF BAGHDAD. THERE IS A MERCHANT SELLING CAMELS HERE."><br>
<img src="images/timezone/0685.png" width="400" alt="YOU ARE IN THE MARKETPLACE OF BAGHDAD. THERE IS A MERCHANT SELLING CAMELS HERE.">

**Type `WEST`**

**686.**

<img src="images/timezone/0686.png" width="400" alt="YOU ARE RIDING A CAMEL. YOU ARE INSIDE THE WALLED CITY OF BAGHDAD.">

**Type `SOUTH`**

**687.**

<img src="images/timezone/0687.png" width="400" alt="YOU ARE RIDING A CAMEL. YOU ARE IN FRONT OF THE WALLED CITY OF BAGHDAD.">

**Type `WEST`**

**688.**

<img src="images/timezone/0688.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE RIDING A CAMEL. YOU ARE IN A HOT DESERT.">

**Type `WEST`**

**689.**

<img src="images/timezone/0689.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE RIDING A CAMEL. YOU ARE IN A DESERT.">

**Type `NORTH`**

**690.**

<img src="images/timezone/0690.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE RIDING A CAMEL. YOU ARE IN AN ARID DESERT.">

**Type `NORTH`**

**691.**

<img src="images/timezone/0691.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE RIDING A CAMEL. YOU ARE IN A DESERT. YOU ARE VERY HOT.">

**Type `NORTH`**

**692.**

<img src="images/timezone/0692.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE RIDING A CAMEL. YOU ARE IN A VERY HOT DESERT.">

**Type `EAST`**

**693.**

<img src="images/timezone/0693.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE RIDING A CAMEL. YOU ARE WANDERING IN A DESERT.">

**Type `EAST`**

**694.**

<img src="images/timezone/0694.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE RIDING A CAMEL. YOU ARE LOST IN A DESERT.">

**Type `NORTH`**

**695.**

<img src="images/timezone/0695-1.png" width="400" alt="YOU ARE RIDING A CAMEL. YOU ARE IN A DESERT. THERE IS A STRANGE LOOKING MOUNTAIN IN FRONT OF YOU."><br>
<img src="images/timezone/0695.png" width="400" alt="YOU ARE IN A DESERT. THERE IS A STRANGE LOOKING MOUNTAIN IN FRONT OF YOU.">

**Type `DISMOUNT CAMEL`**

**696.**

<img src="images/timezone/0696-1.png" width="400" alt="THERE IS A CAMEL FOLLOWING YOU. YOU ARE IN A DESERT. THERE IS A STRANGE LOOKING MOUNTAIN IN FRONT OF YOU."><br>
<img src="images/timezone/0696.png" width="400" alt="YOU ARE IN A DESERT. THERE IS A STRANGE LOOKING MOUNTAIN IN FRONT OF YOU.">

**Type `OPEN SESAME`**

**697.**

<img src="images/timezone/0697-1.png" width="400" alt="MAGICALLY, THE NARROW SLIT IN THE MOUNTAIN WIDENS TO A LARGE CAVE OPENING. THERE IS A CAMEL FOLLOWING YOU."><br>
<img src="images/timezone/0697.png" width="400" alt="YOU ARE IN A DESERT. THERE IS A STRANGE LOOKING MOUNTAIN IN FRONT OF YOU.">

**Type `GO CAVE`**

**698.**

<img src="images/timezone/0698-1.png" width="400" alt="AS SOON AS YOU ENTER THE CAVE, THE ENTRANCE TO THE SOUTH, (BEHIND YOU), CLOSES. YOU ARE INSIDE A CAVE. THERE IS A"><br>
<img src="images/timezone/0698.png" width="400" alt="CLOSES. YOU ARE INSIDE A CAVE. THERE IS A PASSAGE TO THE NORTH.">

**Type `NORTH`**

**699.**

<img src="images/timezone/0699.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN A NARROW PASSAGE OF A CAVE. THE PASSAGE GOES SOUTH AND EAST.">

**Type `EAST`**

**700.**

<img src="images/timezone/0700.png" width="400" alt="--------------- ENTER COMMAND?EAST THERE IS SOME GOLD HERE. YOU ARE IN A LARGE ROOM WITHIN A CAVE.">

**Type `GET GOLD`**

**701.**

<img src="images/timezone/0701.png" width="400" alt="YOU ARE IN A LARGE ROOM WITHIN A CAVE. --------------- ENTER COMMAND?GET GOLD YOU ARE IN A LARGE ROOM WITHIN A CAVE.">

**Type `WEST`**

**702.**

<img src="images/timezone/0702.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN A NARROW PASSAGE OF A CAVE. THE PASSAGE GOES SOUTH AND EAST.">

**Type `SOUTH`**

**703.**

<img src="images/timezone/0703.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE INSIDE A CAVE. THERE IS A PASSAGE TO THE NORTH.">

**Type `OPEN SESAME`**

**704.**

<img src="images/timezone/0704-1.png" width="400" alt="THE CAVE ENTRANCE TO THE SOUTH MAGICALLY OPENS! YOU ARE INSIDE A CAVE. THERE IS A PASSAGE TO THE NORTH."><br>
<img src="images/timezone/0704.png" width="400" alt="MAGICALLY OPENS! YOU ARE INSIDE A CAVE. THERE IS A PASSAGE TO THE NORTH.">

**Type `SOUTH`**

**705.**

<img src="images/timezone/0705-1.png" width="400" alt="MAGICALLY, THE CAVE OPENING IN THE MOUNTAIN CLOSES TO A NARROW SLIT. THERE IS A CAMEL FOLLOWING YOU. YOU ARE IN A DESERT. THERE IS A"><br>
<img src="images/timezone/0705.png" width="400" alt="YOU ARE IN A DESERT. THERE IS A STRANGE LOOKING MOUNTAIN IN FRONT OF YOU.">

**Type `RIDE CAMEL`**

**706.**

<img src="images/timezone/0706-1.png" width="400" alt="YOU ARE RIDING A CAMEL. YOU ARE IN A DESERT. THERE IS A STRANGE LOOKING MOUNTAIN IN FRONT OF YOU."><br>
<img src="images/timezone/0706.png" width="400" alt="YOU ARE IN A DESERT. THERE IS A STRANGE LOOKING MOUNTAIN IN FRONT OF YOU.">

**Type `SOUTH`**

**707.**

<img src="images/timezone/0707.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE RIDING A CAMEL. YOU ARE LOST IN A DESERT.">

**Type `WEST`**

**708.**

<img src="images/timezone/0708.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE RIDING A CAMEL. YOU ARE WANDERING IN A DESERT.">

**Type `WEST`**

**709.**

<img src="images/timezone/0709.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE RIDING A CAMEL. YOU ARE IN A VERY HOT DESERT.">

**Type `SOUTH`**

**710.**

<img src="images/timezone/0710.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE RIDING A CAMEL. YOU ARE IN A DESERT. YOU ARE VERY HOT.">

**Type `SOUTH`**

**711.**

<img src="images/timezone/0711.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE RIDING A CAMEL. YOU ARE IN AN ARID DESERT.">

**Type `SOUTH`**

**712.**

<img src="images/timezone/0712.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE RIDING A CAMEL. YOU ARE IN A DESERT.">

**Type `SOUTH`**

**713.**

<img src="images/timezone/0713.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE RIDING A CAMEL. YOU ARE IN AN ARID DESERT.">

**Type `EAST`**

**714.**

<img src="images/timezone/0714-1.png" width="400" alt="YOU ARE RIDING A CAMEL. YOU ARE IN A DESERT. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING."><br>
<img src="images/timezone/0714.png" width="400" alt="YOU ARE IN A DESERT. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `DISMOUNT CAMEL`**

**715.**

<img src="images/timezone/0715-1.png" width="400" alt="THERE IS A CAMEL FOLLOWING YOU. YOU ARE IN A DESERT. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING."><br>
<img src="images/timezone/0715.png" width="400" alt="YOU ARE IN A DESERT. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `GO MACHINE`**

**716.**

<img src="images/timezone/0716-1.png" width="400" alt="PLEASE INSERT DISK NUMBER 1B AND PRESS RETURN."><br>
<img src="images/timezone/0716-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0716.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

*Side 1B: drop `Time Zone (4am and san inc crack) disk B.dsk` on drive 1, and press Return.*

**Type `SIT`**

**717.**

<img src="images/timezone/0717-1.png" width="400" alt="O.K. THE CHAIR IS COMFORTABLE. AS YOU SIT DOWN, THE MACHINE SEEMS TO COME TO LIFE. YOU ARE INSIDE A TIME MACHINE. YOU HEAR"><br>
<img src="images/timezone/0717.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT. --------------- ENTER COMMAND">

**Type `AUSTRALIA`**

**718.**

<img src="images/timezone/0718-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0718.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `2082AD`**

**719.**

<img src="images/timezone/0719-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0719.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `PULL LEVER`**

**720.**

<img src="images/timezone/0720-1.png" width="400" alt="THE MACHINE STARTS VIBRATING VIOLENTLY AND YOU FEEL DIZZY, ALMOST TO THE POINT OF PASSING OUT. SUDDENLY YOU FEEL A TERRIBLE JOLT AND THE MACHINE IS STILL."><br>
<img src="images/timezone/0720-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0720.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EXIT MACHINE`**

### Australia, 2082 AD

**721.**

<img src="images/timezone/0721-1.png" width="400" alt="NE PLEASE INSERT DISK NUMBER 3E AND PRESS RETURN."><br>
<img src="images/timezone/0721.png" width="400" alt="YOU ARE IN AN OPEN FIELD. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

*Side 3E: drop `Time Zone (4am and san inc crack) disk E.dsk` on drive 1, and press Return.*

**Type `NORTH`**

**722.**

<img src="images/timezone/0722.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN A FIELD. THERE IS A ROAD GOING NORTH AND EAST.">

**Type `NORTH`**

**723.**

<img src="images/timezone/0723.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN THE SUBURB OF SIDNEY. THERE IS A HOUSE TO THE NORTH HERE.">

**Type `BREAK WINDOW`**

**724.**

<img src="images/timezone/0724-1.png" width="400" alt="AS YOU BREAK OPEN THE WINDOW, A LOUD BURGLAR ALARM IS ACTIVATED . YOU ARE IN THE SUBURB OF SIDNEY. THERE IS A HOUSE TO THE NORTH HERE."><br>
<img src="images/timezone/0724.png" width="400" alt="BURGLAR ALARM IS ACTIVATED . YOU ARE IN THE SUBURB OF SIDNEY. THERE IS A HOUSE TO THE NORTH HERE.">

**Type `GO WINDOW`**

**725.**

<img src="images/timezone/0725-1.png" width="400" alt="YOU ARE IN THE COMFORTABLE LIVING ROOM OF THE HOUSE. THE FRONT DOOR IS TO THE SOUTH. THERE IS A BROKEN WINDOW TO THE SOUTH."><br>
<img src="images/timezone/0725.png" width="400" alt="OF THE HOUSE. THE FRONT DOOR IS TO THE SOUTH. THERE IS A BROKEN WINDOW TO THE SOUTH.">

**Type `WEST`**

**726.**

<img src="images/timezone/0726.png" width="400" alt="SOUTH. --------------- ENTER COMMAND?WEST YOU ARE IN THE BEDROOM OF THE HOUSE.">

**Type `OPEN CLOSET`**

**727.**

<img src="images/timezone/0727.png" width="400" alt="O.K. THERE IS A COAT HERE. YOU ARE IN THE BEDROOM OF THE HOUSE.">

**Type `GET COAT`**

**728.**

<img src="images/timezone/0728.png" width="400" alt="YOU ARE IN THE BEDROOM OF THE HOUSE. --------------- ENTER COMMAND?GET COAT YOU ARE IN THE BEDROOM OF THE HOUSE.">

**Type `EAST`**

**729.**

<img src="images/timezone/0729-1.png" width="400" alt="YOU ARE IN THE COMFORTABLE LIVING ROOM OF THE HOUSE. THE FRONT DOOR IS TO THE SOUTH. THERE IS A BROKEN WINDOW TO THE SOUTH."><br>
<img src="images/timezone/0729.png" width="400" alt="OF THE HOUSE. THE FRONT DOOR IS TO THE SOUTH. THERE IS A BROKEN WINDOW TO THE SOUTH.">

**Type `GO WINDOW`**

**730.**

<img src="images/timezone/0730.png" width="400" alt="--------------- ENTER COMMAND?GO WINDOW YOU ARE IN THE SUBURB OF SIDNEY. THERE IS A HOUSE TO THE NORTH HERE.">

**Type `SOUTH`**

**731.**

<img src="images/timezone/0731.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN A FIELD. THERE IS A ROAD GOING NORTH AND EAST.">

**Type `SOUTH`**

**732.**

<img src="images/timezone/0732.png" width="400" alt="YOU ARE IN AN OPEN FIELD. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `GO MACHINE`**

**733.**

<img src="images/timezone/0733-1.png" width="400" alt="PLEASE INSERT DISK NUMBER 1B AND PRESS RETURN."><br>
<img src="images/timezone/0733-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0733.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

*Side 1B: drop `Time Zone (4am and san inc crack) disk B.dsk` on drive 1, and press Return.*

**Type `SIT`**

**734.**

<img src="images/timezone/0734-1.png" width="400" alt="O.K. THE CHAIR IS COMFORTABLE. AS YOU SIT DOWN, THE MACHINE SEEMS TO COME TO LIFE. YOU ARE INSIDE A TIME MACHINE. YOU HEAR"><br>
<img src="images/timezone/0734.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT. --------------- ENTER COMMAND">

**Type `ASIA`**

**735.**

<img src="images/timezone/0735-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0735.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `1700AD`**

**736.**

<img src="images/timezone/0736-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0736.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `PULL LEVER`**

**737.**

<img src="images/timezone/0737-1.png" width="400" alt="THE MACHINE STARTS VIBRATING VIOLENTLY AND YOU FEEL DIZZY, ALMOST TO THE POINT OF PASSING OUT. SUDDENLY YOU FEEL A TERRIBLE JOLT AND THE MACHINE IS STILL."><br>
<img src="images/timezone/0737-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0737.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EXIT MACHINE`**

### Asia, 1700 AD

**738.**

<img src="images/timezone/0738-1.png" width="400" alt="NE PLEASE INSERT DISK NUMBER 4H AND PRESS RETURN."><br>
<img src="images/timezone/0738.png" width="400" alt="YOU ARE IN THE SNOW AND ICE. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

*Side 4H: drop `Time Zone (4am and san inc crack) disk H.dsk` on drive 1, and press Return.*

**Type `WEAR COAT`**

**739.**

<img src="images/timezone/0739-1.png" width="400" alt="O.K. YOU ARE WEARING THE COAT. YOU ARE IN THE SNOW AND ICE. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING."><br>
<img src="images/timezone/0739.png" width="400" alt="YOU ARE IN THE SNOW AND ICE. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `SOUTH`**

**740.**

<img src="images/timezone/0740.png" width="400" alt="PULSATING. --------------- ENTER COMMAND?SOUTH YOU ARE WALKING THROUGH THE ICE.">

**Type `SOUTH`**

**741.**

<img src="images/timezone/0741.png" width="400" alt="YOU ARE IN THE ICE AND SNOW. THERE IS A WATERWAY TO THE SOUTH OF YOU. THERE IS A BRIDGE CROSSING THE WATERWAY.">

**Type `SOUTH`**

**742.**

<img src="images/timezone/0742.png" width="400" alt="YOU ARE IN THE ICE AND SNOW. THERE IS A WATERWAY TO THE NORTH. THERE IS A BRIDGE CROSSING THE WATERWAY.">

**Type `SOUTH`**

**743.**

<img src="images/timezone/0743.png" width="400" alt="BRIDGE CROSSING THE WATERWAY. --------------- ENTER COMMAND?SOUTH YOU ARE IN THE FREEZING ICE AND SNOW.">

**Type `SOUTH`**

**744.**

<img src="images/timezone/0744.png" width="400" alt="YOU ARE IN THE FREEZING ICE AND SNOW. --------------- ENTER COMMAND?SOUTH YOU ARE IN THE FREEZING SNOW AND ICE.">

**Type `EAST`**

**745.**

<img src="images/timezone/0745.png" width="400" alt="YOU ARE IN THE ICE AND SNOW. THERE IS A WATERWAY TO THE EAST. THERE IS A BRIDGE CROSSING THE WATERWAY.">

**Type `EAST`**

**746.**

<img src="images/timezone/0746-1.png" width="400" alt="THERE IS A MEAN LOOKING KOSSACK HERE. I THINK HE IS GOING TO KILL YOU. YOU ARE IN THE ICE AND SNOW. THERE IS A WATERWAY TO THE WEST. THERE IS A BRIDGE"><br>
<img src="images/timezone/0746.png" width="400" alt="YOU ARE IN THE ICE AND SNOW. THERE IS A WATERWAY TO THE WEST. THERE IS A BRIDGE CROSSING THE WATERWAY.">

**Type `KILL KOSSACK`**

**747.**

<img src="images/timezone/0747-1.png" width="400" alt="USING YOUR SWORD, YOU MAKE A BIG GASH IN THE KOSSACK&#39;S ARM. AFRAID FOR HIS LIFE, THE KOSSACK RUNS AWAY. YOU ARE IN THE ICE AND SNOW. THERE IS A"><br>
<img src="images/timezone/0747.png" width="400" alt="YOU ARE IN THE ICE AND SNOW. THERE IS A WATERWAY TO THE WEST. THERE IS A BRIDGE CROSSING THE WATERWAY.">

**Type `EAST`**

**748.**

<img src="images/timezone/0748.png" width="400" alt="YOU ARE IN THE ICE AND SNOW. YOU SEE A CITY IN THE DISTANCE TO THE EAST. THERE IS A SIGN HERE.">

**Type `EAST`**

**749.**

<img src="images/timezone/0749.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN THE TOWN OF ST.PETERSBURG, THE CITY OF MANY BRIDGES.">

**Type `NORTH`**

**750.**

<img src="images/timezone/0750.png" width="400" alt="YOU ARE IN THE ICE AND SNOW. THERE IS A CASTLE IN THE DISTANCE. THE BALTIC SEA IS TO THE EAST.">

**Type `NORTH`**

**751.**

<img src="images/timezone/0751.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN FRONT OF A CASTLE DOOR. THERE ARE TWO KOSSACKS HERE.">

**Type `LOOK`**

**752.**

<img src="images/timezone/0752.png" width="400" alt="YOU SEE NOTHING SPECIAL. YOU ARE IN FRONT OF A CASTLE DOOR. THERE ARE TWO KOSSACKS HERE.">

**Type `LOOK`**

**753.**

<img src="images/timezone/0753.png" width="400" alt="YOU SEE NOTHING SPECIAL. YOU ARE IN FRONT OF A CASTLE DOOR. THERE ARE TWO KOSSACKS HERE.">

**Type `LOOK`**

**754.**

<img src="images/timezone/0754.png" width="400" alt="YOU SEE NOTHING SPECIAL. YOU ARE IN FRONT OF A CASTLE DOOR. THERE ARE TWO KOSSACKS HERE.">

**Type `LOOK`**

**755.**

<img src="images/timezone/0755.png" width="400" alt="YOU SEE NOTHING SPECIAL. YOU ARE IN FRONT OF A CASTLE DOOR. THERE ARE TWO KOSSACKS HERE.">

**Type `LOOK`**

**756.**

<img src="images/timezone/0756-1.png" width="400" alt="YOU SEE NOTHING SPECIAL. THE DOOR OF THE CASTLE SUDDENLY OPENS AND A COACH DRAWN BY TWO HORSES LEAVES THE CASTLE. PETER THE GREAT AND"><br>
<img src="images/timezone/0756-2.png" width="400" alt="CATHERINE THE FIRST ARE INSIDE THE COACH. YOU ARE IN FRONT OF A CASTLE DOOR. THERE ARE TWO KOSSACKS HERE."><br>
<img src="images/timezone/0756.png" width="400" alt="COACH. YOU ARE IN FRONT OF A CASTLE DOOR. THERE ARE TWO KOSSACKS HERE.">

**Type `LOOK`**

**757.**

<img src="images/timezone/0757-1.png" width="400" alt="YOU SEE NOTHING SPECIAL. CATHERINE RAISES HER HAND TO WAVE TO YOU, AND WHEN SHE DOES, SOMETHING FALLS OUT OF HER HAND ONTO THE GROUND.QUICKLY"><br>
<img src="images/timezone/0757-2.png" width="400" alt="THEY RIDE OFF INTO THE COUNTRYSIDE. THERE IS A HAT PIN HERE. YOU ARE IN FRONT OF A CASTLE DOOR. THERE ARE TWO KOSSACKS HERE."><br>
<img src="images/timezone/0757.png" width="400" alt="THERE IS A HAT PIN HERE. YOU ARE IN FRONT OF A CASTLE DOOR. THERE ARE TWO KOSSACKS HERE.">

**Type `LOOK`**

**758.**

<img src="images/timezone/0758-1.png" width="400" alt="YOU SEE NOTHING SPECIAL. THERE IS A HAT PIN HERE. YOU ARE IN FRONT OF A CASTLE DOOR. THERE ARE TWO KOSSACKS HERE."><br>
<img src="images/timezone/0758.png" width="400" alt="THERE IS A HAT PIN HERE. YOU ARE IN FRONT OF A CASTLE DOOR. THERE ARE TWO KOSSACKS HERE.">

**Type `GET PIN`**

**759.**

<img src="images/timezone/0759.png" width="400" alt="--------------- ENTER COMMAND?GET PIN YOU ARE IN FRONT OF A CASTLE DOOR. THERE ARE TWO KOSSACKS HERE.">

**Type `SOUTH`**

**760.**

<img src="images/timezone/0760.png" width="400" alt="YOU ARE IN THE ICE AND SNOW. THERE IS A CASTLE IN THE DISTANCE. THE BALTIC SEA IS TO THE EAST.">

**Type `SOUTH`**

**761.**

<img src="images/timezone/0761.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN THE TOWN OF ST.PETERSBURG, THE CITY OF MANY BRIDGES.">

**Type `WEST`**

**762.**

<img src="images/timezone/0762.png" width="400" alt="YOU ARE IN THE ICE AND SNOW. YOU SEE A CITY IN THE DISTANCE TO THE EAST. THERE IS A SIGN HERE.">

**Type `WEST`**

**763.**

<img src="images/timezone/0763.png" width="400" alt="YOU ARE IN THE ICE AND SNOW. THERE IS A WATERWAY TO THE WEST. THERE IS A BRIDGE CROSSING THE WATERWAY.">

**Type `WEST`**

**764.**

<img src="images/timezone/0764.png" width="400" alt="YOU ARE IN THE ICE AND SNOW. THERE IS A WATERWAY TO THE EAST. THERE IS A BRIDGE CROSSING THE WATERWAY.">

**Type `WEST`**

**765.**

<img src="images/timezone/0765.png" width="400" alt="CROSSING THE WATERWAY. --------------- ENTER COMMAND?WEST YOU ARE IN THE FREEZING SNOW AND ICE.">

**Type `NORTH`**

**766.**

<img src="images/timezone/0766.png" width="400" alt="YOU ARE IN THE FREEZING SNOW AND ICE. --------------- ENTER COMMAND?NORTH YOU ARE IN THE FREEZING ICE AND SNOW.">

**Type `NORTH`**

**767.**

<img src="images/timezone/0767.png" width="400" alt="YOU ARE IN THE ICE AND SNOW. THERE IS A WATERWAY TO THE NORTH. THERE IS A BRIDGE CROSSING THE WATERWAY.">

**Type `NORTH`**

**768.**

<img src="images/timezone/0768.png" width="400" alt="YOU ARE IN THE ICE AND SNOW. THERE IS A WATERWAY TO THE SOUTH OF YOU. THERE IS A BRIDGE CROSSING THE WATERWAY.">

**Type `NORTH`**

**769.**

<img src="images/timezone/0769.png" width="400" alt="A BRIDGE CROSSING THE WATERWAY. --------------- ENTER COMMAND?NORTH YOU ARE WALKING THROUGH THE ICE.">

**Type `NORTH`**

**770.**

<img src="images/timezone/0770.png" width="400" alt="YOU ARE IN THE SNOW AND ICE. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `GO MACHINE`**

**771.**

<img src="images/timezone/0771-1.png" width="400" alt="PLEASE INSERT DISK NUMBER 1B AND PRESS RETURN."><br>
<img src="images/timezone/0771-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0771.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

*Side 1B: drop `Time Zone (4am and san inc crack) disk B.dsk` on drive 1, and press Return.*

**Type `SIT`**

**772.**

<img src="images/timezone/0772-1.png" width="400" alt="O.K. THE CHAIR IS COMFORTABLE. AS YOU SIT DOWN, THE MACHINE SEEMS TO COME TO LIFE. YOU ARE INSIDE A TIME MACHINE. YOU HEAR"><br>
<img src="images/timezone/0772.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT. --------------- ENTER COMMAND">

**Type `PUSH BUTTON`**

**773.**

<img src="images/timezone/0773-1.png" width="400" alt="THE MACHINE STARTS VIBRATING VIOLENTLY AND YOU FEEL DIZZY, ALMOST TO THE POINT OF PASSING OUT. SUDDENLY YOU FEEL A TERRIBLE JOLT AND THE MACHINE IS STILL."><br>
<img src="images/timezone/0773-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0773.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EXIT MACHINE`**

### Home, 1982

**774.**

<img src="images/timezone/0774.png" width="400" alt="YOU ARE IN A FIELD OF DRY GRASS. THERE IS A STRANGE LOOKING MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `DROP SWORD`**

**775.**

<img src="images/timezone/0775.png" width="400" alt="YOU ARE IN A FIELD OF DRY GRASS. THERE IS A STRANGE LOOKING MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `DROP SHIELD`**

**776.**

<img src="images/timezone/0776.png" width="400" alt="YOU ARE IN A FIELD OF DRY GRASS. THERE IS A STRANGE LOOKING MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `DROP POLE`**

**777.**

<img src="images/timezone/0777.png" width="400" alt="YOU ARE IN A FIELD OF DRY GRASS. THERE IS A STRANGE LOOKING MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `DROP COAT`**

**778.**

<img src="images/timezone/0778.png" width="400" alt="YOU ARE IN A FIELD OF DRY GRASS. THERE IS A STRANGE LOOKING MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `GET KEY`**

**779.**

<img src="images/timezone/0779.png" width="400" alt="YOU ARE IN A FIELD OF DRY GRASS. THERE IS A STRANGE LOOKING MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `GET MASK`**

**780.**

<img src="images/timezone/0780.png" width="400" alt="YOU ARE IN A FIELD OF DRY GRASS. THERE IS A STRANGE LOOKING MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `GET BAR`**

**781.**

<img src="images/timezone/0781.png" width="400" alt="YOU ARE IN A FIELD OF DRY GRASS. THERE IS A STRANGE LOOKING MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `GET SAW`**

**782.**

<img src="images/timezone/0782.png" width="400" alt="YOU ARE IN A FIELD OF DRY GRASS. THERE IS A STRANGE LOOKING MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `GO MACHINE`**

**783.**

<img src="images/timezone/0783-1.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0783.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `SIT`**

**784.**

<img src="images/timezone/0784-1.png" width="400" alt="O.K. THE CHAIR IS COMFORTABLE. AS YOU SIT DOWN, THE MACHINE SEEMS TO COME TO LIFE. YOU ARE INSIDE A TIME MACHINE. YOU HEAR"><br>
<img src="images/timezone/0784.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT. --------------- ENTER COMMAND">

**Type `NA`**

**785.**

<img src="images/timezone/0785-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0785.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `2082AD`**

**786.**

<img src="images/timezone/0786-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0786.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `PULL LEVER`**

**787.**

<img src="images/timezone/0787-1.png" width="400" alt="THE MACHINE STARTS VIBRATING VIOLENTLY AND YOU FEEL DIZZY, ALMOST TO THE POINT OF PASSING OUT. SUDDENLY YOU FEEL A TERRIBLE JOLT AND THE MACHINE IS STILL."><br>
<img src="images/timezone/0787-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0787.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EXIT MACHINE`**

### North America, 2082 AD

**788.**

<img src="images/timezone/0788-1.png" width="400" alt="NE PLEASE INSERT DISK NUMBER 5I AND PRESS RETURN."><br>
<img src="images/timezone/0788.png" width="400" alt="YOU ARE IN A VACANT LOT IN LOS ANGELES.THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

*Side 5I: drop `Time Zone (4am and san inc crack) disk I.dsk` on drive 1, and press Return.*

**Type `SOUTH`**

**789.**

<img src="images/timezone/0789.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE ON A CITY STREET IN LOS ANGELES.">

**Type `SOUTH`**

**790.**

<img src="images/timezone/0790.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE ON A CITY STREET IN LOS ANGELES.">

**Type `EAST`**

**791.**

<img src="images/timezone/0791.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE ON A CITY STREET IN LOS ANGELES.">

**Type `SOUTH`**

**792.**

<img src="images/timezone/0792.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE ON A CITY STREET IN LOS ANGELES.">

**Type `WEST`**

**793.**

<img src="images/timezone/0793.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE ON A CITY STREET IN LOS ANGELES.">

**Type `WEST`**

**794.**

<img src="images/timezone/0794.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE ON A CITY STREET IN LOS ANGELES.">

**Type `WEST`**

**795.**

<img src="images/timezone/0795.png" width="400" alt="YOU ARE ON A RESIDENTIAL STREET IN LOS ANGELES. THERE IS A HOUSE TO THE NORTH WITH A CAR PARKED IN FRONT OF IT.">

**Type `NORTH`**

**796.**

<img src="images/timezone/0796.png" width="400" alt="WITH A CAR PARKED IN FRONT OF IT. --------------- ENTER COMMAND?NORTH YOU ARE AT THE FRONT PORCH OF A HOUSE.">

**Type `GET MAT`**

**797.**

<img src="images/timezone/0797.png" width="400" alt="--------------- ENTER COMMAND?GET MAT THERE IS A SMALL KEY HERE. YOU ARE AT THE FRONT PORCH OF A HOUSE.">

**Type `DROP MAT`**

**798.**

<img src="images/timezone/0798.png" width="400" alt="--------------- ENTER COMMAND?DROP MAT THERE IS A SMALL KEY HERE. YOU ARE AT THE FRONT PORCH OF A HOUSE.">

**Type `GET KEY`**

**799.**

<img src="images/timezone/0799.png" width="400" alt="YOU ARE AT THE FRONT PORCH OF A HOUSE. --------------- ENTER COMMAND?GET KEY YOU ARE AT THE FRONT PORCH OF A HOUSE.">

**Type `SOUTH`**

**800.**

<img src="images/timezone/0800.png" width="400" alt="YOU ARE ON A RESIDENTIAL STREET IN LOS ANGELES. THERE IS A HOUSE TO THE NORTH WITH A CAR PARKED IN FRONT OF IT.">

**Type `UNLOCK TRUNK`**

**801.**

<img src="images/timezone/0801-1.png" width="400" alt="O.K. THE TRUNK IS UNLOCKED. YOU ARE ON A RESIDENTIAL STREET IN LOS ANGELES. THERE IS A HOUSE TO THE NORTH WITH A CAR PARKED IN FRONT OF IT."><br>
<img src="images/timezone/0801.png" width="400" alt="YOU ARE ON A RESIDENTIAL STREET IN LOS ANGELES. THERE IS A HOUSE TO THE NORTH WITH A CAR PARKED IN FRONT OF IT.">

**Type `OPEN TRUNK`**

**802.**

<img src="images/timezone/0802-1.png" width="400" alt="O.K. YOU ARE ON A RESIDENTIAL STREET IN LOS ANGELES. THERE IS A HOUSE TO THE NORTH WITH A CAR PARKED IN FRONT OF IT."><br>
<img src="images/timezone/0802.png" width="400" alt="YOU ARE ON A RESIDENTIAL STREET IN LOS ANGELES. THERE IS A HOUSE TO THE NORTH WITH A CAR PARKED IN FRONT OF IT.">

**Type `LOOK TRUNK`**

**803.**

<img src="images/timezone/0803-1.png" width="400" alt="THERE IS DYNAMITE HERE. YOU ARE ON A RESIDENTIAL STREET IN LOS ANGELES. THERE IS A HOUSE TO THE NORTH WITH A CAR PARKED IN FRONT OF IT."><br>
<img src="images/timezone/0803.png" width="400" alt="YOU ARE ON A RESIDENTIAL STREET IN LOS ANGELES. THERE IS A HOUSE TO THE NORTH WITH A CAR PARKED IN FRONT OF IT.">

**Type `GET DYNAMITE`**

**804.**

<img src="images/timezone/0804.png" width="400" alt="YOU ARE ON A RESIDENTIAL STREET IN LOS ANGELES. THERE IS A HOUSE TO THE NORTH WITH A CAR PARKED IN FRONT OF IT.">

**Type `EAST`**

**805.**

<img src="images/timezone/0805.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE ON A CITY STREET IN LOS ANGELES.">

**Type `EAST`**

**806.**

<img src="images/timezone/0806.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE ON A CITY STREET IN LOS ANGELES.">

**Type `EAST`**

**807.**

<img src="images/timezone/0807.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE ON A CITY STREET IN LOS ANGELES.">

**Type `NORTH`**

**808.**

<img src="images/timezone/0808.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE ON A CITY STREET IN LOS ANGELES.">

**Type `WEST`**

**809.**

<img src="images/timezone/0809.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE ON A CITY STREET IN LOS ANGELES.">

**Type `NORTH`**

**810.**

<img src="images/timezone/0810.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE ON A CITY STREET IN LOS ANGELES.">

**Type `NORTH`**

**811.**

<img src="images/timezone/0811.png" width="400" alt="YOU ARE IN A VACANT LOT IN LOS ANGELES.THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `GO MACHINE`**

**812.**

<img src="images/timezone/0812-1.png" width="400" alt="PLEASE INSERT DISK NUMBER 1B AND PRESS RETURN."><br>
<img src="images/timezone/0812-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0812.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

*Side 1B: drop `Time Zone (4am and san inc crack) disk B.dsk` on drive 1, and press Return.*

**Type `SIT`**

**813.**

<img src="images/timezone/0813-1.png" width="400" alt="O.K. THE CHAIR IS COMFORTABLE. AS YOU SIT DOWN, THE MACHINE SEEMS TO COME TO LIFE. YOU ARE INSIDE A TIME MACHINE. YOU HEAR"><br>
<img src="images/timezone/0813.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT. --------------- ENTER COMMAND">

**Type `EUROPE`**

**814.**

<img src="images/timezone/0814-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0814.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `2082AD`**

**815.**

<img src="images/timezone/0815-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0815.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `PULL LEVER`**

**816.**

<img src="images/timezone/0816-1.png" width="400" alt="THE MACHINE STARTS VIBRATING VIOLENTLY AND YOU FEEL DIZZY, ALMOST TO THE POINT OF PASSING OUT. SUDDENLY YOU FEEL A TERRIBLE JOLT AND THE MACHINE IS STILL."><br>
<img src="images/timezone/0816-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0816.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EXIT MACHINE`**

### Europe, 2082 AD

**817.**

<img src="images/timezone/0817.png" width="400" alt="YOU ARE ON A CITY STREET IN LONDON. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `NORTH`**

**818.**

<img src="images/timezone/0818.png" width="400" alt="APPEARS TO BE PULSATING. --------------- ENTER COMMAND?NORTH YOU ARE ON A CITY STREET IN LONDON.">

**Type `WEST`**

**819.**

<img src="images/timezone/0819.png" width="400" alt="YOU ARE ON A CITY STREET IN LONDON. --------------- ENTER COMMAND?WEST YOU ARE ON A CITY STREET IN LONDON.">

**Type `NORTH`**

**820.**

<img src="images/timezone/0820.png" width="400" alt="YOU ARE ON A CITY STREET IN LONDON. --------------- ENTER COMMAND?NORTH YOU ARE ON A CITY STREET IN LONDON.">

**Type `NORTH`**

**821.**

<img src="images/timezone/0821.png" width="400" alt="YOU ARE ON A CITY STREET IN LONDON. --------------- ENTER COMMAND?NORTH YOU ARE ON A CITY STREET IN LONDON.">

**Type `WEST`**

**822.**

<img src="images/timezone/0822.png" width="400" alt="YOU ARE ON A CITY STREET IN LONDON. --------------- ENTER COMMAND?WEST YOU ARE ON A CITY STREET IN LONDON.">

**Type `WEST`**

**823.**

<img src="images/timezone/0823.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE ON A CITY STREET IN LONDON. SCOTLAND YARD IS TO THE WEST OF YOU.">

**Type `WEST`**

**824.**

<img src="images/timezone/0824-1.png" width="400" alt="YOU ARE INSIDE THE SCOTLAND YARD BUILDING. THERE IS A MAN SITTING BEHIND A DESK. THERE IS A PIECE OF PAPER ON THE DESK."><br>
<img src="images/timezone/0824.png" width="400" alt="BUILDING. THERE IS A MAN SITTING BEHIND A DESK. THERE IS A PIECE OF PAPER ON THE DESK.">

**Type `GET DOG`**

**825.**

<img src="images/timezone/0825-1.png" width="400" alt="THE MAN GIVES YOU A DOG. THERE IS A POLICE DOG FOLLOWING YOU. YOU ARE INSIDE THE SCOTLAND YARD BUILDING. THERE IS A MAN SITTING BEHIND"><br>
<img src="images/timezone/0825.png" width="400" alt="BUILDING. THERE IS A MAN SITTING BEHIND A DESK. THERE IS A PIECE OF PAPER ON THE DESK.">

**Type `TIE ROPE`**

**826.**

<img src="images/timezone/0826-1.png" width="400" alt="TO WHAT? THERE IS A POLICE DOG FOLLOWING YOU. YOU ARE INSIDE THE SCOTLAND YARD BUILDING. THERE IS A MAN SITTING BEHIND"><br>
<img src="images/timezone/0826.png" width="400" alt="BUILDING. THERE IS A MAN SITTING BEHIND A DESK. THERE IS A PIECE OF PAPER ON THE DESK.">

**Type `TO DOG`**

**827.**

<img src="images/timezone/0827-1.png" width="400" alt="THE ROPE IS NOW TIED TO THE DOG. THERE IS A POLICE DOG FOLLOWING YOU. YOU ARE INSIDE THE SCOTLAND YARD BUILDING. THERE IS A MAN SITTING BEHIND"><br>
<img src="images/timezone/0827.png" width="400" alt="BUILDING. THERE IS A MAN SITTING BEHIND A DESK. THERE IS A PIECE OF PAPER ON THE DESK.">

**Type `EAST`**

**828.**

<img src="images/timezone/0828.png" width="400" alt="THERE IS A POLICE DOG FOLLOWING YOU. YOU ARE ON A CITY STREET IN LONDON. SCOTLAND YARD IS TO THE WEST OF YOU.">

**Type `EAST`**

**829.**

<img src="images/timezone/0829.png" width="400" alt="--------------- ENTER COMMAND?EAST THERE IS A POLICE DOG FOLLOWING YOU. YOU ARE ON A CITY STREET IN LONDON.">

**Type `EAST`**

**830.**

<img src="images/timezone/0830.png" width="400" alt="--------------- ENTER COMMAND?EAST THERE IS A POLICE DOG FOLLOWING YOU. YOU ARE ON A CITY STREET IN LONDON.">

**Type `SOUTH`**

**831.**

<img src="images/timezone/0831.png" width="400" alt="--------------- ENTER COMMAND?SOUTH THERE IS A POLICE DOG FOLLOWING YOU. YOU ARE ON A CITY STREET IN LONDON.">

**Type `SOUTH`**

**832.**

<img src="images/timezone/0832.png" width="400" alt="--------------- ENTER COMMAND?SOUTH THERE IS A POLICE DOG FOLLOWING YOU. YOU ARE ON A CITY STREET IN LONDON.">

**Type `EAST`**

**833.**

<img src="images/timezone/0833.png" width="400" alt="--------------- ENTER COMMAND?EAST THERE IS A POLICE DOG FOLLOWING YOU. YOU ARE ON A CITY STREET IN LONDON.">

**Type `EAST`**

**834.**

<img src="images/timezone/0834.png" width="400" alt="--------------- ENTER COMMAND?EAST THERE IS A POLICE DOG FOLLOWING YOU. YOU ARE ON A CITY STREET IN LONDON.">

**Type `EAST`**

**835.**

<img src="images/timezone/0835.png" width="400" alt="--------------- ENTER COMMAND?EAST THERE IS A POLICE DOG FOLLOWING YOU. YOU ARE ON A CITY STREET IN LONDON.">

**Type `NORTH`**

**836.**

<img src="images/timezone/0836-1.png" width="400" alt="OH,OH. THERE IS A THIEF WITH A GUN HERE. THERE IS A POLICE DOG FOLLOWING YOU. YOU ARE ON A CITY STREET IN LONDON. THE"><br>
<img src="images/timezone/0836.png" width="400" alt="YOU ARE ON A CITY STREET IN LONDON. THE STREET DEAD-ENDS TO THE NORTH AND GOES SOUTH.">

**Type `UNTIE DOG`**

**837.**

<img src="images/timezone/0837-1.png" width="400" alt="THE POLICE DOG ATTACKS THE THIEF, MAKING HIM DROP HIS GUN. HE RUNS AWAY WITH THE DOG FOLLOWING HIM. YOU ARE ON A CITY STREET IN LONDON. THE"><br>
<img src="images/timezone/0837.png" width="400" alt="YOU ARE ON A CITY STREET IN LONDON. THE STREET DEAD-ENDS TO THE NORTH AND GOES SOUTH.">

**Type `GET GUN`**

**838.**

<img src="images/timezone/0838.png" width="400" alt="YOU ARE ON A CITY STREET IN LONDON. THE STREET DEAD-ENDS TO THE NORTH AND GOES SOUTH.">

**Type `SOUTH`**

**839.**

<img src="images/timezone/0839.png" width="400" alt="SOUTH. --------------- ENTER COMMAND?SOUTH YOU ARE ON A CITY STREET IN LONDON.">

**Type `WEST`**

**840.**

<img src="images/timezone/0840.png" width="400" alt="YOU ARE ON A CITY STREET IN LONDON. --------------- ENTER COMMAND?WEST YOU ARE ON A CITY STREET IN LONDON.">

**Type `WEST`**

**841.**

<img src="images/timezone/0841.png" width="400" alt="YOU ARE ON A CITY STREET IN LONDON. --------------- ENTER COMMAND?WEST YOU ARE ON A CITY STREET IN LONDON.">

**Type `SOUTH`**

**842.**

<img src="images/timezone/0842.png" width="400" alt="YOU ARE ON A CITY STREET IN LONDON. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `GO MACHINE`**

**843.**

<img src="images/timezone/0843-1.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/0843.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `SIT`**

**844.**

<img src="images/timezone/0844-1.png" width="400" alt="O.K. THE CHAIR IS COMFORTABLE. AS YOU SIT DOWN, THE MACHINE SEEMS TO COME TO LIFE. YOU ARE INSIDE A TIME MACHINE. YOU HEAR"><br>
<img src="images/timezone/0844.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT. --------------- ENTER COMMAND">

**Type `4082AD`**

**845.**

<img src="images/timezone/0845-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0845.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `IP`**

**846.**

<img src="images/timezone/0846-1.png" width="400" alt="O.K. IT IS NOW SET. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0846.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `PULL LEVER`**

**847.**

<img src="images/timezone/0847-1.png" width="400" alt="THE MACHINE STARTS VIBRATING VIOLENTLY AND YOU FEEL DIZZY, ALMOST TO THE POINT OF PASSING OUT. SUDDENLY YOU FEEL A TERRIBLE JOLT AND THE MACHINE IS STILL."><br>
<img src="images/timezone/0847-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0847.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `WEAR MASK`**

**848.**

<img src="images/timezone/0848-1.png" width="400" alt="O.K. YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/0848.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EXIT MACHINE`**

### Neburon, 4082 AD

**849.**

<img src="images/timezone/0849-1.png" width="400" alt="NE PLEASE INSERT DISK NUMBER 6K AND PRESS RETURN."><br>
<img src="images/timezone/0849.png" width="400" alt="YOU ARE IN A BARREN WASTELAND. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

*Side 6K: drop `Time Zone (4am and san inc crack) disk K.dsk` on drive 1, and press Return.*

**Type `NORTH`**

**850.**

<img src="images/timezone/0850.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN A BARREN WASTELAND. YOU SEE A DOME IN THE DISTANCE TO THE NORTH.">

**Type `NORTH`**

**851.**

<img src="images/timezone/0851.png" width="400" alt="YOU ARE IN A BARREN WASTELAND. THERE IS A DOME TO THE NORTH. THERE IS A CITY INSIDE THE DOME.">

**Type `NORTH`**

**852.**

<img src="images/timezone/0852.png" width="400" alt="YOU ARE IN FRONT OF A HUGE DOME ON THE BARREN WASTELAND. THERE IS A CITY INSIDE THE DOME.">

**Type `EAST`**

**853.**

<img src="images/timezone/0853.png" width="400" alt="INSIDE THE DOME. --------------- ENTER COMMAND?EAST YOU ARE IN A BARREN WASTELAND.">

**Type `EAST`**

**854.**

<img src="images/timezone/0854.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN A BARREN WASTELAND. THERE IS A HOLE IN THE GROUND HERE.">

**Type `DOWN`**

**855.**

<img src="images/timezone/0855.png" width="400" alt="--------------- ENTER COMMAND?DOWN YOU ARE IN AN UNDERGROUND SEWER SYSTEM.THERE IS A HOLE GOING UP HERE.">

**Type `EAST`**

**856.**

<img src="images/timezone/0856.png" width="400" alt="SYSTEM.THERE IS A HOLE GOING UP HERE. --------------- ENTER COMMAND?EAST YOU ARE IN AN UNDERGROUND SEWER SYSTEM.">

**Type `SOUTH`**

**857.**

<img src="images/timezone/0857.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN AN UNDERGROUND SEWER SYSTEM.THERE IS A GRATE TO THE WEST.">

**Type `USE HAMMER`**

**858.**

<img src="images/timezone/0858-1.png" width="400" alt="USING THE STONE HAMMER, YOU POUND AT THE GRATE UNTIL MOST OF THE RUST HAS FALLEN OFF. YOU CAN PROBABLY OPEN IT NOW."><br>
<img src="images/timezone/0858.png" width="400" alt="NOW. YOU ARE IN AN UNDERGROUND SEWER SYSTEM.THERE IS A GRATE TO THE WEST.">

**Type `DROP HAMMER`**

**859.**

<img src="images/timezone/0859.png" width="400" alt="R YOU ARE IN AN UNDERGROUND SEWER SYSTEM.THERE IS A GRATE TO THE WEST.">

**Type `OPEN GRATE`**

**860.**

<img src="images/timezone/0860.png" width="400" alt="YOU ARE IN AN UNDERGROUND SEWER SYSTEM.THERE IS A GRATE TO THE WEST.">

**Type `WEST`**

**861.**

<img src="images/timezone/0861.png" width="400" alt="YOU ARE IN AN UNDERGROUND SEWER SYSTEM.THERE IS A GRATE HERE. THERE IS A HOLE IN THE WALL OF THE SEWER.">

**Type `WEST`**

**862.**

<img src="images/timezone/0862.png" width="400" alt="A HOLE IN THE WALL OF THE SEWER. --------------- ENTER COMMAND?WEST YOU ARE IN AN UNDERGROUND SEWER SYSTEM.">

**Type `WEST`**

**863.**

<img src="images/timezone/0863.png" width="400" alt="YOU ARE IN AN UNDERGROUND SEWER SYSTEM. --------------- ENTER COMMAND?WEST YOU ARE IN AN UNDERGROUND SEWER SYSTEM.">

**Type `NORTH`**

**864.**

<img src="images/timezone/0864.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN AN UNDERGROUND SEWER SYSTEM.THERE IS A MANHOLE ABOVE YOU.">

**Type `USE KNIFE`**

**865.**

<img src="images/timezone/0865-1.png" width="400" alt="THE KNIFE FITS IN THE CRACK BETWEEN THE COVER AND THE MANHOLE PERFECTLY. CAREFULLY, YOU SCRAPE AWAY THE RUST, FREEING THE MANHOLE COVER."><br>
<img src="images/timezone/0865.png" width="400" alt="FREEING THE MANHOLE COVER. YOU ARE IN AN UNDERGROUND SEWER SYSTEM.THERE IS A MANHOLE ABOVE YOU.">

**Type `DROP KNIFE`**

**866.**

<img src="images/timezone/0866.png" width="400" alt="YOU ARE IN AN UNDERGROUND SEWER SYSTEM.THERE IS A MANHOLE ABOVE YOU.">

**Type `OPEN COVER`**

**867.**

<img src="images/timezone/0867-1.png" width="400" alt="YOU PUSH WITH ALL YOUR STRENGTH ON THE MANHOLE COVER UNTIL IT TURNS OVER ON THE STREET ABOVE. YOU ARE IN AN UNDERGROUND SEWER"><br>
<img src="images/timezone/0867.png" width="400" alt="THE STREET ABOVE. YOU ARE IN AN UNDERGROUND SEWER SYSTEM.THERE IS A MANHOLE ABOVE YOU.">

**Type `UP`**

**868.**

<img src="images/timezone/0868-1.png" width="400" alt="THERE IS OXYGEN HERE. YOU CAN BREATHE WITHOUT THE OXYGEN MASK. THERE IS A WALLET HERE. YOU ARE ON A DEAD-END STREET. THERE IS"><br>
<img src="images/timezone/0868.png" width="400" alt="THERE IS A WALLET HERE. YOU ARE ON A DEAD-END STREET. THERE IS A MANHOLE HERE.">

**Type `GET WALLET`**

**869.**

<img src="images/timezone/0869.png" width="400" alt="YOU ARE ON A DEAD-END STREET. THERE IS A MANHOLE HERE.">

**Type `OPEN WALLET`**

**870.**

<img src="images/timezone/0870.png" width="400" alt="O.K. YOU ARE ON A DEAD-END STREET. THERE IS A MANHOLE HERE.">

**Type `GET ID`**

**871.**

<img src="images/timezone/0871.png" width="400" alt="O.K. YOU ARE ON A DEAD-END STREET. THERE IS A MANHOLE HERE.">

**Type `DROP WALLET`**

**872.**

<img src="images/timezone/0872.png" width="400" alt="T YOU ARE ON A DEAD-END STREET. THERE IS A MANHOLE HERE.">

**Type `DOWN`**

**873.**

<img src="images/timezone/0873.png" width="400" alt="--------------- ENTER COMMAND?DOWN YOU ARE IN AN UNDERGROUND SEWER SYSTEM.THERE IS A MANHOLE ABOVE YOU.">

**Type `SOUTH`**

**874.**

<img src="images/timezone/0874.png" width="400" alt="SYSTEM.THERE IS A MANHOLE ABOVE YOU. --------------- ENTER COMMAND?SOUTH YOU ARE IN AN UNDERGROUND SEWER SYSTEM.">

**Type `EAST`**

**875.**

<img src="images/timezone/0875.png" width="400" alt="YOU ARE IN AN UNDERGROUND SEWER SYSTEM. --------------- ENTER COMMAND?EAST YOU ARE IN AN UNDERGROUND SEWER SYSTEM.">

**Type `EAST`**

**876.**

<img src="images/timezone/0876.png" width="400" alt="YOU ARE IN AN UNDERGROUND SEWER SYSTEM.THERE IS A GRATE HERE. THERE IS A HOLE IN THE WALL OF THE SEWER.">

**Type `TIE ROPE`**

**877.**

<img src="images/timezone/0877-1.png" width="400" alt="TIE ROPE TO WHAT? YOU ARE IN AN UNDERGROUND SEWER SYSTEM.THERE IS A GRATE HERE. THERE IS A HOLE IN THE WALL OF THE SEWER."><br>
<img src="images/timezone/0877.png" width="400" alt="YOU ARE IN AN UNDERGROUND SEWER SYSTEM.THERE IS A GRATE HERE. THERE IS A HOLE IN THE WALL OF THE SEWER.">

**Type `TO ROCK`**

**878.**

<img src="images/timezone/0878-1.png" width="400" alt="O.K. YOU ARE IN AN UNDERGROUND SEWER SYSTEM.THERE IS A GRATE HERE. THERE IS A HOLE IN THE WALL OF THE SEWER."><br>
<img src="images/timezone/0878.png" width="400" alt="YOU ARE IN AN UNDERGROUND SEWER SYSTEM.THERE IS A GRATE HERE. THERE IS A HOLE IN THE WALL OF THE SEWER.">

**Type `DOWN`**

**879.**

<img src="images/timezone/0879-1.png" width="400" alt="YOU ARE IN A LARGE UNDERGROUND CAVERN. THERE IS A ROPE HANGING FROM A PIT WAY UP IN THE CEILING OF THE CAVERN. THERE IS A STREAM HERE."><br>
<img src="images/timezone/0879.png" width="400" alt="THERE IS A ROPE HANGING FROM A PIT WAY UP IN THE CEILING OF THE CAVERN. THERE IS A STREAM HERE.">

**Type `WEST`**

**880.**

<img src="images/timezone/0880.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN A PASSAGE. THERE IS A SLIGHT INCLINE TO THE WEST.">

**Type `WEST`**

**881.**

<img src="images/timezone/0881.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN A SMALL PASSAGE. IT IS A LITTLE STEEPER HERE.">

**Type `WEST`**

**882.**

<img src="images/timezone/0882.png" width="400" alt="LITTLE STEEPER HERE. --------------- ENTER COMMAND?WEST THE PASSAGE ENDS INTO A CLIFF HERE.">

**Type `UP`**

**883.**

<img src="images/timezone/0883.png" width="400" alt="YOU ARE AT THE EDGE OF A CLIFF GOING DOWN. THERE IS A PASSAGE GOING TO THE WEST.">

**Type `WEST`**

**884.**

<img src="images/timezone/0884.png" width="400" alt="WEST. --------------- ENTER COMMAND?WEST YOU ARE IN AN EAST/WEST PASSAGE.">

**Type `WEST`**

**885.**

<img src="images/timezone/0885.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE AT THE END OF A PASSAGE. THERE IS A GRATE TO THE NORTH HERE.">

**Type `OPEN GRATE`**

**886.**

<img src="images/timezone/0886.png" width="400" alt="YOU ARE AT THE END OF A PASSAGE. THERE IS A GRATE TO THE NORTH HERE.">

**Type `NORTH`**

**887.**

<img src="images/timezone/0887-1.png" width="400" alt="THERE IS OXYGEN HERE. YOU CAN BREATHE WITHOUT THE OXYGEN MASK. YOU ARE IN AN AIR CONDITIONING DUCT. THERE IS A GRATE TO THE SOUTH."><br>
<img src="images/timezone/0887.png" width="400" alt="WITHOUT THE OXYGEN MASK. YOU ARE IN AN AIR CONDITIONING DUCT. THERE IS A GRATE TO THE SOUTH.">

**Type `REMOVE MASK`**

**888.**

<img src="images/timezone/0888.png" width="400" alt="O.K. YOU ARE IN AN AIR CONDITIONING DUCT. THERE IS A GRATE TO THE SOUTH.">

**Type `EAST`**

**889.**

<img src="images/timezone/0889.png" width="400" alt="THERE IS A GRATE TO THE SOUTH. --------------- ENTER COMMAND?EAST YOU ARE IN AN AIR CONDITIONING DUCT.">

**Type `NORTH`**

**890.**

<img src="images/timezone/0890.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN AN AIR CONDITIONING DUCT. THERE IS A GRATE HERE.">

**Type `OPEN GRATE`**

**891.**

<img src="images/timezone/0891.png" width="400" alt="YOU ARE IN AN AIR CONDITIONING DUCT. THERE IS A GRATE HERE.">

**Type `DOWN`**

**892.**

<img src="images/timezone/0892-1.png" width="400" alt="YOU HEAR FOOTSTEPS COMING FROM THE NORTH. YOU ARE IN A NORTH/SOUTH HALLWAY. YOU SEE AN OPEN GRATE OVERHEAD."><br>
<img src="images/timezone/0892.png" width="400" alt="NORTH. YOU ARE IN A NORTH/SOUTH HALLWAY. YOU SEE AN OPEN GRATE OVERHEAD.">

**Type `SOUTH`**

**893.**

<img src="images/timezone/0893-1.png" width="400" alt="YOU STILL HEAR FOOTSTEPS COMING FROM THE NORTH. YOU ARE IN A HALLWAY. THERE IS A DOOR TO THE EAST."><br>
<img src="images/timezone/0893.png" width="400" alt="THE NORTH. YOU ARE IN A HALLWAY. THERE IS A DOOR TO THE EAST.">

**Type `EAST`**

**894.**

<img src="images/timezone/0894.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU HEAR FOOTSTEPS VERY CLOSE NOW. YOU ARE INSIDE AN EMPTY CLOSET.">

**Type `CLOSE DOOR`**

**895.**

<img src="images/timezone/0895-1.png" width="400" alt="YOU HEAR SOMEONE RIGHT OUTSIDE THE CLOSET. THERE IS A PEEPHOLE IN THE DOOR. YOU ARE INSIDE AN EMPTY CLOSET."><br>
<img src="images/timezone/0895.png" width="400" alt="CLOSET. THERE IS A PEEPHOLE IN THE DOOR. YOU ARE INSIDE AN EMPTY CLOSET.">

**Type `LOOK PEEPHOLE`**

**896.**

<img src="images/timezone/0896-1.png" width="400" alt="YOU SEE A GUARD STANDING OUTSIDE THE CLOSET. HE IS WEARING A FUNNY LOOKING UNIFORM. YOU ARE INSIDE AN EMPTY CLOSET."><br>
<img src="images/timezone/0896.png" width="400" alt="CLOSET. HE IS WEARING A FUNNY LOOKING UNIFORM. YOU ARE INSIDE AN EMPTY CLOSET.">

**Type `LOOK`**

**897.**

<img src="images/timezone/0897.png" width="400" alt="YOU HEAR THE FOOTSTEPS GOING AWAY. THERE IS A PEEPHOLE IN THE DOOR. YOU ARE INSIDE AN EMPTY CLOSET.">

**Type `OPEN DOOR`**

**898.**

<img src="images/timezone/0898.png" width="400" alt="YOU ARE INSIDE AN EMPTY CLOSET. --------------- ENTER COMMAND?OPEN DOOR YOU ARE INSIDE AN EMPTY CLOSET.">

**Type `WEST`**

**899.**

<img src="images/timezone/0899.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN A HALLWAY. THERE IS A DOOR TO THE EAST.">

**Type `NORTH`**

**900.**

<img src="images/timezone/0900.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN A NORTH/SOUTH HALLWAY. YOU SEE AN OPEN GRATE OVERHEAD.">

**Type `NORTH`**

**901.**

<img src="images/timezone/0901.png" width="400" alt="YOU SEE A GUARD COMING TOWARD YOU. HE LOOKS MENACING. YOU ARE AT A JUNCTION OF HALLWAYS.">

**Type `USE GUN`**

**902.**

<img src="images/timezone/0902.png" width="400" alt="YOU SHOOT THE GUARD WITH YOUR LASER GUN AND HE FALLS DOWN DEAD. YOU ARE AT A JUNCTION OF HALLWAYS.">

**Type `GET UNIFORM`**

**903.**

<img src="images/timezone/0903.png" width="400" alt="M YOU TAKE THE UNIFORM OFF THE GUARD. YOU ARE AT A JUNCTION OF HALLWAYS.">

**Type `WEAR UNIFORM`**

**904.**

<img src="images/timezone/0904.png" width="400" alt="RM O.K. YOU PUT THE UNIFORM ON. YOU ARE AT A JUNCTION OF HALLWAYS.">

**Type `EAST`**

**905.**

<img src="images/timezone/0905.png" width="400" alt="YOU ARE IN A HALLWAY. THERE IS AN OPEN DOORWAY TO THE NORTH. STAIRS GO UP FROM HERE.">

**Type `NORTH`**

**906.**

<img src="images/timezone/0906.png" width="400" alt="YOU ARE ON A STREET THAT GOES EAST, WEST AND NORTH. A COMMUNAL DORMITORY IS TO THE SOUTH OF YOU. THE DOOR IS OPEN.">

**Type `NORTH`**

**907.**

<img src="images/timezone/0907-1.png" width="400" alt="AS SOON AS YOU LEAVE THE COMMUNAL DORMITORY, THE GOVERNMENT POLICE ARREST YOU AND THROW YOU IN JAIL FOR THE ALLEGED MURDER OF A GUARD IN THE"><br>
<img src="images/timezone/0907.png" width="400" alt="ALLEGED MURDER OF A GUARD IN THE DORMITORY. YOU ARE IN A JAIL CELL.">

**Type `USE SAW`**

**908.**

<img src="images/timezone/0908.png" width="400" alt="YOU PATIENTLY SAW THE BARS OF THE WINDOW UNTIL THEY BREAK APART. YOU ARE IN A JAIL CELL.">

**Type `GO WINDOW`**

**909.**

<img src="images/timezone/0909.png" width="400" alt="YOU ARE IN A BACK ALLEY. A JAIL CELL WINDOW IS HERE. THE BARS HAVE BEEN SAWED APART.">

**Type `EAST`**

**910.**

<img src="images/timezone/0910-1.png" width="400" alt="OH NO!! THERE IS A THIEF WITH A RAY GUN HERE. HE SAYS,&#34;YOUR MONEY OR YOUR LIFE.&#34; YOU ARE IN A PARK. A STREET GOES NORTH"><br>
<img src="images/timezone/0910.png" width="400" alt="LIFE.&#34; YOU ARE IN A PARK. A STREET GOES NORTH AND SOUTH. AN ALLEY LEADS WEST.">

**Type `GIVE GOLD`**

**911.**

<img src="images/timezone/0911-1.png" width="400" alt="YOU HAND THE GOLD OVER TO THE THIEF, WHO GRABS IT THEN RUNS AWAY. YOU ARE IN A PARK. A STREET GOES NORTH AND SOUTH. AN ALLEY LEADS WEST."><br>
<img src="images/timezone/0911.png" width="400" alt="WHO GRABS IT THEN RUNS AWAY. YOU ARE IN A PARK. A STREET GOES NORTH AND SOUTH. AN ALLEY LEADS WEST.">

**Type `NORTH`**

**912.**

<img src="images/timezone/0912.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN FRONT OF THE GOVERNMENT BUILDING. THERE IS A GUARD HERE.">

**Type `EAST`**

**913.**

<img src="images/timezone/0913.png" width="400" alt="YOU ARE ON AN EAST/WEST STREET. AN ALLEY GOES TO THE NORTH FOLLOWING THE SIDE OF A BUILDING.">

**Type `EAST`**

**914.**

<img src="images/timezone/0914.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN A NURSERY. THERE ARE MANY PLANTS GROWING HERE.">

**Type `EAST`**

**915.**

<img src="images/timezone/0915.png" width="400" alt="THERE IS AN ESPECIALLY BEAUTIFUL FLOWER HERE. YOU ARE IN A NURSERY.">

**Type `GET FLOWER`**

**916.**

<img src="images/timezone/0916.png" width="400" alt="--------------- ENTER COMMAND?GET FLOWER YOU ARE IN A NURSERY.">

**Type `WEST`**

**917.**

<img src="images/timezone/0917.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN A NURSERY. THERE ARE MANY PLANTS GROWING HERE.">

**Type `WEST`**

**918.**

<img src="images/timezone/0918.png" width="400" alt="YOU ARE ON AN EAST/WEST STREET. AN ALLEY GOES TO THE NORTH FOLLOWING THE SIDE OF A BUILDING.">

**Type `WEST`**

**919.**

<img src="images/timezone/0919.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN FRONT OF THE GOVERNMENT BUILDING. THERE IS A GUARD HERE.">

**Type `WEST`**

**920.**

<img src="images/timezone/0920.png" width="400" alt="YOU ARE ON AN EAST/WEST STREET. AN ALLEY GOES TO THE NORTH FOLLOWING THE SIDE OF A BUILDING.">

**Type `NORTH`**

**921.**

<img src="images/timezone/0921-1.png" width="400" alt="YOU ARE ON THE WEST SIDE OF THE GOVERNMENT BUILDING. AN ALLEY GOES TO THE SOUTH. A HIGH FENCE IS HERE. THERE IS A CEMENT BLOCK HERE."><br>
<img src="images/timezone/0921.png" width="400" alt="GOVERNMENT BUILDING. AN ALLEY GOES TO THE SOUTH. A HIGH FENCE IS HERE. THERE IS A CEMENT BLOCK HERE.">

**Type `USE BAR`**

**922.**

<img src="images/timezone/0922-1.png" width="400" alt="USING THE IRON BAR, YOU CAREFULLY LIFT UP THE CEMENT BLOCK UNTIL IT TURNS OVER. THERE IS A HOLE IN THE GROUND UNDER THE BLOCK."><br>
<img src="images/timezone/0922-2.png" width="400" alt="YOU ARE ON THE WEST SIDE OF THE GOVERNMENT BUILDING. AN ALLEY GOES TO THE SOUTH. A HIGH FENCE IS HERE. THERE IS A CEMENT BLOCK HERE."><br>
<img src="images/timezone/0922.png" width="400" alt="GOVERNMENT BUILDING. AN ALLEY GOES TO THE SOUTH. A HIGH FENCE IS HERE. THERE IS A CEMENT BLOCK HERE.">

**Type `WEAR MASK`**

**923.**

<img src="images/timezone/0923-1.png" width="400" alt="O.K. YOU ARE ON THE WEST SIDE OF THE GOVERNMENT BUILDING. AN ALLEY GOES TO THE SOUTH. A HIGH FENCE IS HERE. THERE"><br>
<img src="images/timezone/0923.png" width="400" alt="GOVERNMENT BUILDING. AN ALLEY GOES TO THE SOUTH. A HIGH FENCE IS HERE. THERE IS A CEMENT BLOCK HERE.">

**Type `DOWN`**

**924.**

<img src="images/timezone/0924.png" width="400" alt="--------------- ENTER COMMAND?DOWN YOU ARE IN AN UNDERGROUND SEWER SYSTEM.THERE IS A HOLE OVERHEAD.">

**Type `NORTH`**

**925.**

<img src="images/timezone/0925.png" width="400" alt="SYSTEM.THERE IS A HOLE OVERHEAD. --------------- ENTER COMMAND?NORTH YOU ARE IN AN UNDERGROUND SEWER SYSTEM.">

**Type `EAST`**

**926.**

<img src="images/timezone/0926.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN AN UNDERGROUND SEWER SYSTEM.THERE IS A GRATE OVERHEAD.">

**Type `USE LADDER`**

**927.**

<img src="images/timezone/0927-1.png" width="400" alt="YOU LEAN THE LADDER AGAINST THE SEWER WALLS. YOU ARE IN AN UNDERGROUND SEWER SYSTEM.THERE IS A GRATE OVERHEAD."><br>
<img src="images/timezone/0927.png" width="400" alt="WALLS. YOU ARE IN AN UNDERGROUND SEWER SYSTEM.THERE IS A GRATE OVERHEAD.">

**Type `OPEN GRATE`**

**928.**

<img src="images/timezone/0928.png" width="400" alt="O.K. YOU ARE IN AN UNDERGROUND SEWER SYSTEM.THERE IS A GRATE OVERHEAD.">

**Type `UP`**

**929.**

<img src="images/timezone/0929-1.png" width="400" alt="THERE IS OXYGEN HERE. YOU CAN BREATHE WITHOUT THE OXYGEN MASK. YOU ARE IN A BASEMENT. THERE ARE STAIRS GOING UP. THERE IS A GRATE IN THE"><br>
<img src="images/timezone/0929.png" width="400" alt="YOU ARE IN A BASEMENT. THERE ARE STAIRS GOING UP. THERE IS A GRATE IN THE FLOOR.">

**Type `REMOVE MASK`**

**930.**

<img src="images/timezone/0930-1.png" width="400" alt="O.K. YOU ARE IN A BASEMENT. THERE ARE STAIRS GOING UP. THERE IS A GRATE IN THE FLOOR."><br>
<img src="images/timezone/0930.png" width="400" alt="YOU ARE IN A BASEMENT. THERE ARE STAIRS GOING UP. THERE IS A GRATE IN THE FLOOR.">

**Type `UP`**

**931.**

<img src="images/timezone/0931.png" width="400" alt="--------------- ENTER COMMAND?UP YOU ARE AT THE TOP OF SOME STAIRS. THERE IS A DOOR TO THE EAST.">

**Type `OPEN DOOR`**

**932.**

<img src="images/timezone/0932.png" width="400" alt="O.K. YOU ARE AT THE TOP OF SOME STAIRS. THERE IS A DOOR TO THE EAST.">

**Type `EAST`**

**933.**

<img src="images/timezone/0933.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN A NORTH/SOUTH HALLWAY. THERE IS A DOOR TO THE WEST.">

**Type `NORTH`**

**934.**

<img src="images/timezone/0934-1.png" width="400" alt="--------------- ENTER COMMAND?NORTH PLEASE INSERT DISK NUMBER 6L AND PRESS RETURN."><br>
<img src="images/timezone/0934.png" width="400" alt="AND PRESS RETURN. YOU ARE AT A JUNCTION OF HALLWAYS. THERE IS A DOORWAY TO THE NORTH.">

*Side 6L: drop `Time Zone (4am and san inc crack) disk L.dsk` on drive 1, and press Return.*

**Type `WEST`**

**935.**

<img src="images/timezone/0935.png" width="400" alt="THERE IS A DOORWAY TO THE NORTH. --------------- ENTER COMMAND?WEST YOU ARE IN AN OFFICE.">

**Type `OPEN DRAWER`**

**936.**

<img src="images/timezone/0936.png" width="400" alt="R O.K. YOU ARE IN AN OFFICE.">

**Type `LOOK DRAWER`**

**937.**

<img src="images/timezone/0937.png" width="400" alt="R THERE IS A MILITARY CARD (MC) HERE. YOU ARE IN AN OFFICE.">

**Type `GET MC`**

**938.**

<img src="images/timezone/0938.png" width="400" alt="YOU ARE IN AN OFFICE. --------------- ENTER COMMAND?GET MC YOU ARE IN AN OFFICE.">

**Type `EAST`**

**939.**

<img src="images/timezone/0939.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE AT A JUNCTION OF HALLWAYS. THERE IS A DOORWAY TO THE NORTH.">

**Type `EAST`**

**940.**

<img src="images/timezone/0940.png" width="400" alt="THERE IS A DOORWAY TO THE NORTH. --------------- ENTER COMMAND?EAST YOU ARE IN AN ASSEMBLY ROOM.">

**Type `EAST`**

**941.**

<img src="images/timezone/0941.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE AT A JUNCTION OF HALLWAYS. THERE IS A GUARD HERE.">

**Type `SOUTH`**

**942.**

<img src="images/timezone/0942.png" width="400" alt="--------------- ENTER COMMAND?SOUTH YOU ARE IN AN OFFICE. THERE IS A DESK, A CHAIR AND A SAFE HERE.">

**Type `USE PIN`**

**943.**

<img src="images/timezone/0943-1.png" width="400" alt="YOU POKE THE HAT PIN INTO THE LOCK AND JIGGLE IT AROUND A BIT. SUDDENLY YOU HEAR A CLICK AND THE SAFE IS UNLOCKED. YOU ARE IN AN OFFICE. THERE IS A DESK,"><br>
<img src="images/timezone/0943.png" width="400" alt="HEAR A CLICK AND THE SAFE IS UNLOCKED. YOU ARE IN AN OFFICE. THERE IS A DESK, A CHAIR AND A SAFE HERE.">

**Type `OPEN SAFE`**

**944.**

<img src="images/timezone/0944.png" width="400" alt="O.K. YOU ARE IN AN OFFICE. THERE IS A DESK, A CHAIR AND A SAFE HERE.">

**Type `LOOK SAFE`**

**945.**

<img src="images/timezone/0945.png" width="400" alt="THERE IS A NOTE HERE. YOU ARE IN AN OFFICE. THERE IS A DESK, A CHAIR AND A SAFE HERE.">

**Type `GET NOTE`**

**946.**

<img src="images/timezone/0946.png" width="400" alt="--------------- ENTER COMMAND?GET NOTE YOU ARE IN AN OFFICE. THERE IS A DESK, A CHAIR AND A SAFE HERE.">

**Type `READ NOTE`**

**947.**

<img src="images/timezone/0947-1.png" width="400" alt="THE NOTE SAYS,&#34;THE PASSWORD IS ALLEGIANCE.&#34; YOU ARE IN AN OFFICE. THERE IS A DESK, A CHAIR AND A SAFE HERE."><br>
<img src="images/timezone/0947.png" width="400" alt="ALLEGIANCE.&#34; YOU ARE IN AN OFFICE. THERE IS A DESK, A CHAIR AND A SAFE HERE.">

**Type `NORTH`**

**948.**

<img src="images/timezone/0948.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE AT A JUNCTION OF HALLWAYS. THERE IS A GUARD HERE.">

**Type `NORTH`**

**949.**

<img src="images/timezone/0949.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN A NORTH/SOUTH HALLWAY. THERE IS A DOOR TO THE EAST.">

**Type `OPEN DOOR`**

**950.**

<img src="images/timezone/0950.png" width="400" alt="O.K. YOU ARE IN A NORTH/SOUTH HALLWAY. THERE IS A DOOR TO THE EAST.">

**Type `EAST`**

**951.**

<img src="images/timezone/0951.png" width="400" alt="IS A DOOR TO THE EAST. --------------- ENTER COMMAND?EAST YOU ARE IN RAMADU&#39;S OFFICE.">

**Type `CLOSE DOOR`**

**952.**

<img src="images/timezone/0952.png" width="400" alt="O.K. YOU ARE IN RAMADU&#39;S OFFICE.">

**Type `USE GUN`**

**953.**

<img src="images/timezone/0953.png" width="400" alt="YOU POINT THE LASER GUN STRAIGHT AT RAMADU AND KILL HIM. YOU ARE IN RAMADU&#39;S OFFICE.">

**Type `OPEN DRAWER`**

**954.**

<img src="images/timezone/0954.png" width="400" alt="R O.K. YOU ARE IN RAMADU&#39;S OFFICE.">

**Type `LOOK DRAWER`**

**955.**

<img src="images/timezone/0955.png" width="400" alt="O.K. RAMADU&#39;S MILITARY CARD (MC) IS HERE. YOU ARE IN RAMADU&#39;S OFFICE.">

**Type `GET MC`**

**956.**

<img src="images/timezone/0956.png" width="400" alt="YOU ARE IN RAMADU&#39;S OFFICE. --------------- ENTER COMMAND?GET MC YOU ARE IN RAMADU&#39;S OFFICE.">

**Type `LOOK ROOM`**

**957.**

<img src="images/timezone/0957.png" width="400" alt="YOU ARE IN RAMADU&#39;S OFFICE. --------------- ENTER COMMAND?LOOK ROOM YOU ARE IN RAMADU&#39;S OFFICE.">

**Type `DROP GUN`**

**958.**

<img src="images/timezone/0958.png" width="400" alt="YOU ARE IN RAMADU&#39;S OFFICE. --------------- ENTER COMMAND?DROP GUN YOU ARE IN RAMADU&#39;S OFFICE.">

**Type `OPEN DOOR`**

**959.**

<img src="images/timezone/0959.png" width="400" alt="--------------- ENTER COMMAND?OPEN DOOR O.K. YOU ARE IN RAMADU&#39;S OFFICE.">

**Type `WEST`**

**960.**

<img src="images/timezone/0960.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN A NORTH/SOUTH HALLWAY. THERE IS A DOOR TO THE EAST.">

**Type `NORTH`**

**961.**

<img src="images/timezone/0961.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN A CONFERENCE ROOM. THERE ARE DOOWAYS TO THE WEST AND SOUTH.">

**Type `WEST`**

**962.**

<img src="images/timezone/0962.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN A HALLWAY. THERE IS A DOOR TO THE NORTH WITH A SLOT NEXT TO IT.">

**Type `INSERT MC`**

**963.**

<img src="images/timezone/0963-1.png" width="400" alt="WHEN YOU INSERT THE MILITARY CARD INTO THE SLOT, THE DOOR OPENS. YOU ARE IN A HALLWAY. THERE IS A DOOR TO THE NORTH WITH A SLOT NEXT TO IT."><br>
<img src="images/timezone/0963.png" width="400" alt="THE SLOT, THE DOOR OPENS. YOU ARE IN A HALLWAY. THERE IS A DOOR TO THE NORTH WITH A SLOT NEXT TO IT.">

**Type `NORTH`**

**964.**

<img src="images/timezone/0964-1.png" width="400" alt="YOU ARE IN A YARD. THE BACK OF THE GOVERNMENT BUILDING IS TO THE SOUTH.THERE IS A DOOR IN THE BUILDING WITH A SLOT NEXT TO IT. YOU SEE A FENCE"><br>
<img src="images/timezone/0964.png" width="400" alt="SOUTH.THERE IS A DOOR IN THE BUILDING WITH A SLOT NEXT TO IT. YOU SEE A FENCE IN THE DISTANCE.">

**Type `WEST`**

**965.**

<img src="images/timezone/0965.png" width="400" alt="--------------- ENTER COMMAND?WEST THERE IS A TALL FENCE HERE. THERE IS A SHED HERE TOO.">

**Type `OPEN DOOR`**

**966.**

<img src="images/timezone/0966.png" width="400" alt="O.K. THERE IS A TALL FENCE HERE. THERE IS A SHED HERE TOO.">

**Type `GO SHED`**

**967.**

<img src="images/timezone/0967.png" width="400" alt="SHED HERE TOO. --------------- ENTER COMMAND?GO SHED YOU ARE INSIDE AN EMPTY SHED.">

**Type `CLOSE DOOR`**

**968.**

<img src="images/timezone/0968.png" width="400" alt="O.K. YOU ARE INSIDE AN EMPTY SHED.">

**Type `LOOK`**

**969.**

<img src="images/timezone/0969.png" width="400" alt="--------------- ENTER COMMAND?LOOK YOU SEE NOTHING SPECIAL. YOU ARE INSIDE AN EMPTY SHED.">

**Type `LOOK`**

**970.**

<img src="images/timezone/0970.png" width="400" alt="--------------- ENTER COMMAND?LOOK YOU SEE NOTHING SPECIAL. YOU ARE INSIDE AN EMPTY SHED.">

**Type `LOOK`**

**971.**

<img src="images/timezone/0971.png" width="400" alt="--------------- ENTER COMMAND?LOOK YOU SEE NOTHING SPECIAL. YOU ARE INSIDE AN EMPTY SHED.">

**Type `LOOK`**

**972.**

<img src="images/timezone/0972.png" width="400" alt="--------------- ENTER COMMAND?LOOK YOU SEE NOTHING SPECIAL. YOU ARE INSIDE AN EMPTY SHED.">

**Type `LOOK`**

**973.**

<img src="images/timezone/0973.png" width="400" alt="--------------- ENTER COMMAND?LOOK YOU SEE NOTHING SPECIAL. YOU ARE INSIDE AN EMPTY SHED.">

**Type `LOOK`**

**974.**

<img src="images/timezone/0974.png" width="400" alt="--------------- ENTER COMMAND?LOOK YOU SEE NOTHING SPECIAL. YOU ARE INSIDE AN EMPTY SHED.">

**Type `LOOK`**

**975.**

<img src="images/timezone/0975.png" width="400" alt="--------------- ENTER COMMAND?LOOK YOU SEE NOTHING SPECIAL. YOU ARE INSIDE AN EMPTY SHED.">

**Type `LOOK`**

**976.**

<img src="images/timezone/0976.png" width="400" alt="--------------- ENTER COMMAND?LOOK YOU SEE NOTHING SPECIAL. YOU ARE INSIDE AN EMPTY SHED.">

**Type `LOOK`**

**977.**

<img src="images/timezone/0977.png" width="400" alt="--------------- ENTER COMMAND?LOOK YOU SEE NOTHING SPECIAL. YOU ARE INSIDE AN EMPTY SHED.">

**Type `LOOK`**

**978.**

<img src="images/timezone/0978.png" width="400" alt="YOU SEE NOTHING SPECIAL. YOU HEAR SOMEBODY COMING. YOU ARE INSIDE AN EMPTY SHED.">

**Type `LOOK`**

**979.**

<img src="images/timezone/0979.png" width="400" alt="YOU SEE NOTHING SPECIAL. YOU HEAR SOMEBODY COMING. YOU ARE INSIDE AN EMPTY SHED.">

**Type `LOOK`**

**980.**

<img src="images/timezone/0980-1.png" width="400" alt="YOU SEE NOTHING SPECIAL. YOU HEAR VOICES OUTSIDE THE SHED, BUT THEY DO NOT LOOK INSIDE. IN A MINUTE, YOU HEAR THEM GO AWAY."><br>
<img src="images/timezone/0980.png" width="400" alt="THEY DO NOT LOOK INSIDE. IN A MINUTE, YOU HEAR THEM GO AWAY. YOU ARE INSIDE AN EMPTY SHED.">

**Type `OPEN DOOR`**

**981.**

<img src="images/timezone/0981.png" width="400" alt="--------------- ENTER COMMAND?OPEN DOOR O.K. YOU ARE INSIDE AN EMPTY SHED.">

**Type `EAST`**

**982.**

<img src="images/timezone/0982.png" width="400" alt="--------------- ENTER COMMAND?EAST THERE IS A TALL FENCE HERE. THERE IS A SHED HERE TOO.">

**Type `EAST`**

**983.**

<img src="images/timezone/0983-1.png" width="400" alt="YOU ARE IN A YARD. THE BACK OF THE GOVERNMENT BUILDING IS TO THE SOUTH.THERE IS A DOOR IN THE BUILDING WITH A SLOT NEXT TO IT. YOU SEE A FENCE"><br>
<img src="images/timezone/0983.png" width="400" alt="SOUTH.THERE IS A DOOR IN THE BUILDING WITH A SLOT NEXT TO IT. YOU SEE A FENCE IN THE DISTANCE.">

**Type `NORTH`**

**984.**

<img src="images/timezone/0984.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN FRONT OF A TALL FENCE WITH A GATE.">

**Type `USE KEY`**

**985.**

<img src="images/timezone/0985-1.png" width="400" alt="USING THE SKELETON KEY, YOU MANAGE TO UNLOCK THE GATE. YOU ARE IN FRONT OF A TALL FENCE WITH A GATE."><br>
<img src="images/timezone/0985.png" width="400" alt="UNLOCK THE GATE. YOU ARE IN FRONT OF A TALL FENCE WITH A GATE.">

**Type `OPEN GATE`**

**986.**

<img src="images/timezone/0986.png" width="400" alt="O.K. YOU ARE IN FRONT OF A TALL FENCE WITH A GATE.">

**Type `NORTH`**

**987.**

<img src="images/timezone/0987.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE ON A NORTH/SOUTH PATH. THERE IS A TALL FENCE WITH A GATE BEHIND YOU.">

**Type `NORTH`**

**988.**

<img src="images/timezone/0988.png" width="400" alt="YOU ARE IN FRONT OF A TALL WIRE FENCE. THERE IS A GATE IN THE FENCE. TWO GUARDS ARE HERE.">

**Type `LOOK`**

**989.**

<img src="images/timezone/0989-1.png" width="400" alt="THE GUARDS TELL YOU THAT THEIR GREAT LEADER (RAMADU), HAS BEEN KILLED. THEY ARE CHECKING-OUT EVERYBODY WHO COMES HERE. QUICKLY THEY FRISK YOU. THEY FIND"><br>
<img src="images/timezone/0989-2.png" width="400" alt="NO GUN ON YOU. YOU ARE CLEARED. YOU ARE IN FRONT OF A TALL WIRE FENCE. THERE IS A GATE IN THE FENCE. TWO GUARDS ARE HERE."><br>
<img src="images/timezone/0989.png" width="400" alt="YOU ARE IN FRONT OF A TALL WIRE FENCE. THERE IS A GATE IN THE FENCE. TWO GUARDS ARE HERE.">

**Type `ALLEGIANCE`**

**990.**

<img src="images/timezone/0990-1.png" width="400" alt="WHEN YOU SAY THE PASSWORD, THE GUARDS OPEN THE GATES. YOU ARE IN FRONT OF A TALL WIRE FENCE. THERE IS A GATE IN THE FENCE. TWO"><br>
<img src="images/timezone/0990.png" width="400" alt="YOU ARE IN FRONT OF A TALL WIRE FENCE. THERE IS A GATE IN THE FENCE. TWO GUARDS ARE HERE.">

**Type `NORTH`**

**991.**

<img src="images/timezone/0991-1.png" width="400" alt="YOU ARE IN FRONT OF THE MILITARY INSTALLATION. THE DOOR IS WIDE OPEN. THERE IS A TALL WIRE FENCE BEHIND YOU.THERE IS A GATE IN THE FENCE. TWO"><br>
<img src="images/timezone/0991.png" width="400" alt="THERE IS A TALL WIRE FENCE BEHIND YOU.THERE IS A GATE IN THE FENCE. TWO GUARDS ARE ON THE OTHER SIDE.">

**Type `NORTH`**

**992.**

<img src="images/timezone/0992-1.png" width="400" alt="YOU ARE IN THE LOBBY OF THE MILITARY INSTALLATION. THERE IS A MAN SITTING BEHIND A DESK HERE. THERE ARE OPEN DOORWAYS TO THE WEST AND SOUTH."><br>
<img src="images/timezone/0992.png" width="400" alt="INSTALLATION. THERE IS A MAN SITTING BEHIND A DESK HERE. THERE ARE OPEN DOORWAYS TO THE WEST AND SOUTH.">

**Type `LOOK`**

**993.**

<img src="images/timezone/0993-1.png" width="400" alt="THE GUARD ASKS TO SEE YOUR ID CARD AND YOUR MILITARY CARD (MC). YOU ARE IN THE LOBBY OF THE MILITARY INSTALLATION. THERE IS A MAN SITTING"><br>
<img src="images/timezone/0993.png" width="400" alt="INSTALLATION. THERE IS A MAN SITTING BEHIND A DESK HERE. THERE ARE OPEN DOORWAYS TO THE WEST AND SOUTH.">

**Type `SHOW ID`**

**994.**

<img src="images/timezone/0994-1.png" width="400" alt="O.K. YOU ARE IN THE LOBBY OF THE MILITARY INSTALLATION. THERE IS A MAN SITTING BEHIND A DESK HERE. THERE ARE OPEN"><br>
<img src="images/timezone/0994.png" width="400" alt="INSTALLATION. THERE IS A MAN SITTING BEHIND A DESK HERE. THERE ARE OPEN DOORWAYS TO THE WEST AND SOUTH.">

**Type `SHOW MC`**

**995.**

<img src="images/timezone/0995-1.png" width="400" alt="O.K. YOU ARE IN THE LOBBY OF THE MILITARY INSTALLATION. THERE IS A MAN SITTING BEHIND A DESK HERE. THERE ARE OPEN"><br>
<img src="images/timezone/0995.png" width="400" alt="INSTALLATION. THERE IS A MAN SITTING BEHIND A DESK HERE. THERE ARE OPEN DOORWAYS TO THE WEST AND SOUTH.">

**Type `WEST`**

**996.**

<img src="images/timezone/0996.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN AN EAST/WEST HALLWAY. THERE IS AN OPEN DOORWAY TO THE NORTH.">

**Type `WEST`**

**997.**

<img src="images/timezone/0997.png" width="400" alt="IS AN OPEN DOORWAY TO THE NORTH. --------------- ENTER COMMAND?WEST YOU ARE AT A JUNCTION OF HALLWAYS.">

**Type `SOUTH`**

**998.**

<img src="images/timezone/0998.png" width="400" alt="YOU ARE AT A JUNCTION OF HALLWAYS. --------------- ENTER COMMAND?SOUTH YOU ARE AT A JUNCTION OF HALLWAYS.">

**Type `SOUTH`**

**999.**

<img src="images/timezone/0999.png" width="400" alt="YOU ARE IN A HALLWAY WHICH DEAD ENDS TO THE SOUTH. THERE IS A DOOR TO THE WEST. THERE IS AN ARMED GUARD BY THE DOOR.">

**Type `GIVE FLOWER`**

**1000.**

<img src="images/timezone/1000-1.png" width="400" alt="THIS GUY HAPPENS TO LOVE FLOWERS. HE TAKES THE FLOWER AND HOLDS IT TO HIS NOSE AND SNIFFS IT. IMMEDIATELY HE FALLS ASLEEP."><br>
<img src="images/timezone/1000.png" width="400" alt="YOU ARE IN A HALLWAY WHICH DEAD ENDS TO THE SOUTH. THERE IS A DOOR TO THE WEST. THERE IS AN ARMED GUARD BY THE DOOR.">

**Type `OPEN DOOR`**

**1001.**

<img src="images/timezone/1001-1.png" width="400" alt="O.K. YOU ARE IN A HALLWAY WHICH DEAD ENDS TO THE SOUTH. THERE IS A DOOR TO THE WEST. THERE IS AN ARMED GUARD BY THE DOOR."><br>
<img src="images/timezone/1001.png" width="400" alt="YOU ARE IN A HALLWAY WHICH DEAD ENDS TO THE SOUTH. THERE IS A DOOR TO THE WEST. THERE IS AN ARMED GUARD BY THE DOOR.">

**Type `WEST`**

**1002.**

<img src="images/timezone/1002.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN A MAZE OF HALLWAYS. THERE IS A DOOR HERE.">

**Type `CLOSE DOOR`**

**1003.**

<img src="images/timezone/1003.png" width="400" alt="O.K. YOU ARE IN A MAZE OF HALLWAYS. THERE IS A DOOR HERE.">

**Type `USE LANCE`**

**1004.**

<img src="images/timezone/1004-1.png" width="400" alt="YOU HOLD THE LANCE OUT IN FRONT OF YOU AND TAP THE GROUND AS YOU WALK. YOU ARE IN A MAZE OF HALLWAYS. THERE IS A DOOR HERE."><br>
<img src="images/timezone/1004.png" width="400" alt="AND TAP THE GROUND AS YOU WALK. YOU ARE IN A MAZE OF HALLWAYS. THERE IS A DOOR HERE.">

**Type `WEST`**

**1005.**

<img src="images/timezone/1005.png" width="400" alt="A DOOR HERE. --------------- ENTER COMMAND?WEST YOU ARE IN A MAZE OF HALLWAYS.">

**Type `WEST`**

**1006.**

<img src="images/timezone/1006.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?WEST YOU ARE IN A MAZE OF HALLWAYS.">

**Type `SOUTH`**

**1007.**

<img src="images/timezone/1007-1.png" width="400" alt="ZZZZAP!! THE TIP OF YOUR LANCE JUST TRIGGERED A LASER MINE, BUT YOU ARE OKAY. YOU ARE IN A MAZE OF HALLWAYS."><br>
<img src="images/timezone/1007.png" width="400" alt="TRIGGERED A LASER MINE, BUT YOU ARE OKAY. YOU ARE IN A MAZE OF HALLWAYS.">

**Type `SOUTH`**

**1008.**

<img src="images/timezone/1008.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?SOUTH YOU ARE IN A MAZE OF HALLWAYS.">

**Type `WEST`**

**1009.**

<img src="images/timezone/1009.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?WEST YOU ARE IN A MAZE OF HALLWAYS.">

**Type `SOUTH`**

**1010.**

<img src="images/timezone/1010.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?SOUTH YOU ARE IN A MAZE OF HALLWAYS.">

**Type `EAST`**

**1011.**

<img src="images/timezone/1011.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?EAST YOU ARE IN A MAZE OF HALLWAYS.">

**Type `SOUTH`**

**1012.**

<img src="images/timezone/1012-1.png" width="400" alt="ZZZZAP!! THE TIP OF YOUR LANCE JUST TRIGGERED A LASER MINE, BUT YOU ARE OKAY. YOU ARE IN A MAZE OF HALLWAYS."><br>
<img src="images/timezone/1012.png" width="400" alt="TRIGGERED A LASER MINE, BUT YOU ARE OKAY. YOU ARE IN A MAZE OF HALLWAYS.">

**Type `SOUTH`**

**1013.**

<img src="images/timezone/1013.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?SOUTH YOU ARE IN A MAZE OF HALLWAYS.">

**Type `SOUTH`**

**1014.**

<img src="images/timezone/1014.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?SOUTH YOU ARE IN A MAZE OF HALLWAYS.">

**Type `SOUTH`**

**1015.**

<img src="images/timezone/1015.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?SOUTH YOU ARE IN A MAZE OF HALLWAYS.">

**Type `WEST`**

**1016.**

<img src="images/timezone/1016.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?WEST YOU ARE IN A MAZE OF HALLWAYS.">

**Type `WEST`**

**1017.**

<img src="images/timezone/1017.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?WEST YOU ARE IN A MAZE OF HALLWAYS.">

**Type `WEST`**

**1018.**

<img src="images/timezone/1018.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?WEST YOU ARE IN A MAZE OF HALLWAYS.">

**Type `NORTH`**

**1019.**

<img src="images/timezone/1019.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?NORTH YOU ARE IN A MAZE OF HALLWAYS.">

**Type `WEST`**

**1020.**

<img src="images/timezone/1020.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?WEST YOU ARE IN A MAZE OF HALLWAYS.">

**Type `SOUTH`**

**1021.**

<img src="images/timezone/1021-1.png" width="400" alt="ZZZZAP!! THE TIP OF YOUR LANCE JUST TRIGGERED A LASER MINE, BUT YOU ARE OKAY. YOU ARE IN A MAZE OF HALLWAYS."><br>
<img src="images/timezone/1021.png" width="400" alt="TRIGGERED A LASER MINE, BUT YOU ARE OKAY. YOU ARE IN A MAZE OF HALLWAYS.">

**Type `SOUTH`**

**1022.**

<img src="images/timezone/1022.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?SOUTH YOU ARE IN A MAZE OF HALLWAYS.">

**Type `WEST`**

**1023.**

<img src="images/timezone/1023.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?WEST YOU ARE IN A MAZE OF HALLWAYS.">

**Type `SOUTH`**

**1024.**

<img src="images/timezone/1024.png" width="400" alt="YOU ARE IN A HALLWAY THAT DEAD-ENDS TO THE SOUTH. THERE IS A DOOR TO THE WEST.">

**Type `INSERT MC`**

**1025.**

<img src="images/timezone/1025-1.png" width="400" alt="YOU INSERT RAMADU&#39;S MILITARY CARD INTO THE SLOT AND THE DOOR OPENS. YOU ARE IN A HALLWAY THAT DEAD-ENDS TO THE SOUTH. THERE IS A DOOR TO THE WEST."><br>
<img src="images/timezone/1025.png" width="400" alt="YOU ARE IN A HALLWAY THAT DEAD-ENDS TO THE SOUTH. THERE IS A DOOR TO THE WEST.">

**Type `WEST`**

**1026.**

<img src="images/timezone/1026-1.png" width="400" alt="THERE IS A GIANT RAY MACHINE POINTING OUT INTO SPACE. YOU ARE IN A LARGE ROOM WITH A DOMED CEILING."><br>
<img src="images/timezone/1026.png" width="400" alt="OUT INTO SPACE. YOU ARE IN A LARGE ROOM WITH A DOMED CEILING.">

**Type `DROP DYNAMITE`**

**1027.**

<img src="images/timezone/1027-1.png" width="400" alt="THERE IS A GIANT RAY MACHINE POINTING OUT INTO SPACE. YOU ARE IN A LARGE ROOM WITH A DOMED CEILING."><br>
<img src="images/timezone/1027.png" width="400" alt="OUT INTO SPACE. YOU ARE IN A LARGE ROOM WITH A DOMED CEILING.">

**Type `LIGHT FUSE`**

**1028.**

<img src="images/timezone/1028-1.png" width="400" alt="YOU LIGHT THE FUSE TO THE DYNAMITE. THE FUSE IS BURNING QUICKLY. THERE IS A GIANT RAY MACHINE POINTING OUT INTO SPACE."><br>
<img src="images/timezone/1028.png" width="400" alt="OUT INTO SPACE. YOU ARE IN A LARGE ROOM WITH A DOMED CEILING.">

**Type `EAST`**

**1029.**

<img src="images/timezone/1029-1.png" width="400" alt="--------------- ENTER COMMAND?EAST BOOOM!! THE DYNAMITE EXPLODES. YOU ARE IN A HALLWAY THAT DEAD-ENDS TO THE SOUTH. THERE IS A DOOR TO THE WEST."><br>
<img src="images/timezone/1029.png" width="400" alt="YOU ARE IN A HALLWAY THAT DEAD-ENDS TO THE SOUTH. THERE IS A DOOR TO THE WEST.">

**Type `NORTH`**

**1030.**

<img src="images/timezone/1030.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN A MAZE OF HALLWAYS.">

**Type `EAST`**

**1031.**

<img src="images/timezone/1031.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?EAST YOU ARE IN A MAZE OF HALLWAYS.">

**Type `NORTH`**

**1032.**

<img src="images/timezone/1032-1.png" width="400" alt="ZZZZAP!! THE TIP OF YOUR LANCE JUST TRIGGERED A LASER MINE, BUT YOU ARE OKAY. YOU ARE IN A MAZE OF HALLWAYS."><br>
<img src="images/timezone/1032.png" width="400" alt="TRIGGERED A LASER MINE, BUT YOU ARE OKAY. YOU ARE IN A MAZE OF HALLWAYS.">

**Type `NORTH`**

**1033.**

<img src="images/timezone/1033.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?NORTH YOU ARE IN A MAZE OF HALLWAYS.">

**Type `EAST`**

**1034.**

<img src="images/timezone/1034.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?EAST YOU ARE IN A MAZE OF HALLWAYS.">

**Type `SOUTH`**

**1035.**

<img src="images/timezone/1035.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?SOUTH YOU ARE IN A MAZE OF HALLWAYS.">

**Type `EAST`**

**1036.**

<img src="images/timezone/1036.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?EAST YOU ARE IN A MAZE OF HALLWAYS.">

**Type `EAST`**

**1037.**

<img src="images/timezone/1037.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?EAST YOU ARE IN A MAZE OF HALLWAYS.">

**Type `EAST`**

**1038.**

<img src="images/timezone/1038.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?EAST YOU ARE IN A MAZE OF HALLWAYS.">

**Type `NORTH`**

**1039.**

<img src="images/timezone/1039.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?NORTH YOU ARE IN A MAZE OF HALLWAYS.">

**Type `NORTH`**

**1040.**

<img src="images/timezone/1040.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?NORTH YOU ARE IN A MAZE OF HALLWAYS.">

**Type `NORTH`**

**1041.**

<img src="images/timezone/1041-1.png" width="400" alt="ZZZZAP!! THE TIP OF YOUR LANCE JUST TRIGGERED A LASER MINE, BUT YOU ARE OKAY. YOU ARE IN A MAZE OF HALLWAYS."><br>
<img src="images/timezone/1041.png" width="400" alt="TRIGGERED A LASER MINE, BUT YOU ARE OKAY. YOU ARE IN A MAZE OF HALLWAYS.">

**Type `NORTH`**

**1042.**

<img src="images/timezone/1042.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?NORTH YOU ARE IN A MAZE OF HALLWAYS.">

**Type `WEST`**

**1043.**

<img src="images/timezone/1043.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?WEST YOU ARE IN A MAZE OF HALLWAYS.">

**Type `NORTH`**

**1044.**

<img src="images/timezone/1044.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?NORTH YOU ARE IN A MAZE OF HALLWAYS.">

**Type `NORTH`**

**1045.**

<img src="images/timezone/1045.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?NORTH YOU ARE IN A MAZE OF HALLWAYS.">

**Type `WEST`**

**1046.**

<img src="images/timezone/1046.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?WEST YOU ARE IN A MAZE OF HALLWAYS.">

**Type `NORTH`**

**1047.**

<img src="images/timezone/1047.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?NORTH YOU ARE IN A MAZE OF HALLWAYS.">

**Type `WEST`**

**1048.**

<img src="images/timezone/1048.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?WEST YOU ARE IN A MAZE OF HALLWAYS.">

**Type `SOUTH`**

**1049.**

<img src="images/timezone/1049.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?SOUTH YOU ARE IN A MAZE OF HALLWAYS.">

**Type `WEST`**

**1050.**

<img src="images/timezone/1050-1.png" width="400" alt="ZZZZAP!! THE TIP OF YOUR LANCE JUST TRIGGERED A LASER MINE, BUT YOU ARE OKAY. YOU ARE IN A MAZE OF HALLWAYS."><br>
<img src="images/timezone/1050.png" width="400" alt="TRIGGERED A LASER MINE, BUT YOU ARE OKAY. YOU ARE IN A MAZE OF HALLWAYS.">

**Type `SOUTH`**

**1051.**

<img src="images/timezone/1051.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS. --------------- ENTER COMMAND?SOUTH YOU ARE IN A MAZE OF HALLWAYS.">

**Type `WEST`**

**1052.**

<img src="images/timezone/1052.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS WITH A BUTTON ON THE WALL. THERE IS A CIRCLE ON THE FLOOR HERE.">

**Type `STAND CIRCLE`**

**1053.**

<img src="images/timezone/1053-1.png" width="400" alt="O.K. YOU ARE IN A MAZE OF HALLWAYS WITH A BUTTON ON THE WALL. THERE IS A CIRCLE ON THE FLOOR HERE."><br>
<img src="images/timezone/1053.png" width="400" alt="YOU ARE IN A MAZE OF HALLWAYS WITH A BUTTON ON THE WALL. THERE IS A CIRCLE ON THE FLOOR HERE.">

**Type `PUSH BUTTON`**

**1054.**

<img src="images/timezone/1054-1.png" width="400" alt="AS YOU STAND IN THE CIRCLE AND PRESS THE BUTTON SIMULTANEOUSLY, YOU ACTIVATE THE TELE-TRANSPORTER. YOU HEAR A FAINT HUMMING, THAT GROWS RAPIDLY LOUDER"><br>
<img src="images/timezone/1054-2.png" width="400" alt="SURFACE. PLEASE INSERT DISK NUMBER 6K AND PRESS RETURN."><br>
<img src="images/timezone/1054-3.png" width="400" alt="PLEASE INSERT DISK NUMBER 6K AND PRESS RETURN. YOU ARE IN A BARREN WASTELAND. THERE IS"><br>
<img src="images/timezone/1054.png" width="400" alt="YOU ARE IN A BARREN WASTELAND. THERE IS A TIME MACHINE HERE. IT APPEARS TO BE PULSATING.">

*Side 6K: drop `Time Zone (4am and san inc crack) disk K.dsk` on drive 1, and press Return.*

**Type `GO MACHINE`**

**1055.**

<img src="images/timezone/1055-1.png" width="400" alt="PLEASE INSERT DISK NUMBER 1B AND PRESS RETURN."><br>
<img src="images/timezone/1055-2.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT."><br>
<img src="images/timezone/1055.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

*Side 1B: drop `Time Zone (4am and san inc crack) disk B.dsk` on drive 1, and press Return.*

**Type `SIT`**

**1056.**

<img src="images/timezone/1056-1.png" width="400" alt="O.K. THE CHAIR IS COMFORTABLE. AS YOU SIT DOWN, THE MACHINE SEEMS TO COME TO LIFE. YOU ARE INSIDE A TIME MACHINE. YOU HEAR"><br>
<img src="images/timezone/1056.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT. --------------- ENTER COMMAND">

**Type `PUSH BUTTON`**

**1057.**

<img src="images/timezone/1057-1.png" width="400" alt="THE MACHINE STARTS VIBRATING VIOLENTLY AND YOU FEEL DIZZY, ALMOST TO THE POINT OF PASSING OUT. SUDDENLY YOU FEEL A TERRIBLE JOLT AND THE MACHINE IS STILL."><br>
<img src="images/timezone/1057-2.png" width="400" alt="YOU HAVE TRIED TO TAKE SOMETHING TOO FAR BACK IN TIME. IT HAS BEEN LOST."><br>
<img src="images/timezone/1057-3.png" width="400" alt="YOU ARE INSIDE A TIME MACHINE. YOU HEAR A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO"><br>
<img src="images/timezone/1057.png" width="400" alt="A FAINT HUMMING SOUND. THERE ARE TWO DIALS AND A LEVER WITH A BUTTON NEXT TO IT.">

**Type `EXIT MACHINE`**

**1058.**

<img src="images/timezone/1058.png" width="400" alt="YOU ARE IN A FIELD OF DRY GRASS. THERE IS A STRANGE LOOKING MACHINE HERE. IT APPEARS TO BE PULSATING.">

**Type `SOUTH`**

### The end

<img src="images/timezone/1059-1.png" width="400" alt="AS YOU RETURN HOME, NEWS OF YOUR DARING ADVENTURES HAS REACHED YOUR FRIENDS AND FAMILY. YOU RECIEVE A HERO&#39;S WELCOME, COMPLETE WITH FANFARE AND A KEY TO THE"><br>
<img src="images/timezone/1059-2.png" width="400" alt="COMPLETE WITH FANFARE AND A KEY TO THE CITY. YOU ARE HEREBY DECLARED AN...&#34;ULTIMATE ADVENTURER.&#34; THANK YOU FOR PLAYING TIME ZONE.">

<!-- End of the walkthrough -->

## What next

[Karateka](karateka.md) is another game of the Apple \]\[ played on this
machine, and [Life with DOS 3.3](dos33.md) shows the system the disks of
Time Zone are written in.
