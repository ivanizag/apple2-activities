# Mission: Asteroid, from the start to the end

[Back to the activities](../README.md)

**Mission: Asteroid**, *Hi-Res Adventure #0* of On-Line Systems, came out in
1980 for the Apple \]\[, written by Ken and Roberta Williams: a short
adventure, numbered before [Mystery House](mysteryhouse.md). An asteroid is going to hit the Earth at 07:15 tonight;
you are the astronaut who has to fly to it and blow it up.

You type one or two words, `PUSH SWITCH`, `TAKE SHOWER`, `PULL THROTTLE`, and
the game answers under the picture. Each command takes five minutes of the
watch of the game. The whole game is on one side of a diskette.

This page plays it to its end, all 80 commands, with a picture of the screen
after each one: 114 pictures. The pictures are of a colour television, with
the four lines of text under them drawn white and sharp, as a monochrome
monitor showed them: the television blurred them into fringes of colour,
hard to read.

## What you need

**Mission Asteroid (4am and san inc crack)**, from
[the Asimov archive](https://mirrors.apple2.org.za/ftp.apple.asimov.net/images/games/adventure/):
the zip `Mission Asteroid (4am and san inc crack).zip`, with the disk
`Mission Asteroid (4am and san inc crack).dsk`. It is also on
[its page on the Internet Archive](https://archive.org/details/MissionAsteroid4amCrack).
`./fetch-disks.sh` in this repository downloads it into `disks/` and checks
it.

The commands of this page are also in this repository, a line each,
[listings/missionasteroid.txt](listings/missionasteroid.txt). They follow the
walkthrough of
[The Sierra Help Pages](https://sierrahelp.com/Walkthroughs/MissionAsteroidWalkthrough.html),
with what the machine showed it needs: the diskette goes `INTO DRIVE` and the
explosives `IN PIT`, the answers to the questions the game asks, and two more
`LOOK WATCH` at the end, for the time the explosives take.

The [manual of Mission: Asteroid](https://archive.org/details/Mission_Asteroid_Hi-Res_Adventure_0)
is on the Internet Archive.

## The machine

An Apple \]\[+:

- the 6502 processor at 1 MHz and 48 KB of memory;
- a Disk II controller card in slot 6, with the game in drive 1.

```bash
izapple2 -model none -board 2plus -cpu 6502 -screen color \
    -rom "<internal>/Apple2_Plus.rom" \
    -charrom "<internal>/Apple2rev7CharGen.rom" -forceCaps \
    -s6 'diskii,disk1=disks/Mission Asteroid (4am and san inc crack).dsk'
```

## Playing it

1. **Start izapple2** with the command above. The game starts in front of a
   building.

2. **Type the commands**, each followed by Return. Each step below is what
   the game shows, and the command to type then. When the picture stays the
   same and only the text goes on, the page shows the picture once, and
   under it the lines of text that came after. When the game stops in the
   middle of a longer text, **press Return** to go on.

## The walkthrough

<!-- The walkthrough, written by the generator -->

- [1. At the space center](#1-at-the-space-center)
- [2. Into space](#2-into-space)
- [3. The asteroid](#3-the-asteroid)
- [4. Back to Earth](#4-back-to-earth)
- [The end](#the-end)

### 1. At the space center

<img src="images/missionasteroid/0001.png" width="400" alt="YOU ARE IN FRONT OF A BUILDING.">

**Type `INV`**

<img src="images/missionasteroid/0002.png" width="400" alt="YOU HEAR A NOISE, &#34;BEEP, BEEP, BEEP!&#34; A WATCH. YOU ARE IN FRONT OF A BUILDING.">

**Type `LOOK WATCH`**

<img src="images/missionasteroid/0003-1.png" width="400" alt="THE BEEPER WATCH HAS A SWITCH ON IT. THE TIME IS 12:05 YOU HEAR A NOISE, &#34;BEEP, BEEP, BEEP!&#34;"><br>
<img src="images/missionasteroid/0003.png" width="400" alt="YOU HEAR A NOISE, &#34;BEEP, BEEP, BEEP!&#34; YOU ARE IN FRONT OF A BUILDING.">

**Type `PUSH SWITCH`**

<img src="images/missionasteroid/0004-1.png" width="400" alt="A VOICE FROM THE BEEPER WATCH SAYS, &#34;THIS IS MISSION CONTROL, YOU ARE TO REPORT TO THE BRIEFING ROOM AT ONCE. THE PASSWORD IS STARSTRUCK.&#34;"><br>
<img src="images/missionasteroid/0004.png" width="400" alt="REPORT TO THE BRIEFING ROOM AT ONCE. THE PASSWORD IS STARSTRUCK.&#34; YOU ARE IN FRONT OF A BUILDING.">

**Type `OPEN DOOR`**

<img src="images/missionasteroid/0005.png" width="400" alt="--------------- ENTER COMMAND?OPEN DOOR OK YOU ARE IN FRONT OF A BUILDING.">

**Type `GO DOOR`**

<img src="images/missionasteroid/0006-1.png" width="400" alt="YOU ARE IN THE FRONT OFFICE OF THE BUILDING. THERE IS A SECRETARY HERE. THERE IS A DOORWAY TO THE NORTH AND A DOORWAY TO THE EAST."><br>
<img src="images/missionasteroid/0006.png" width="400" alt="BUILDING. THERE IS A SECRETARY HERE. THERE IS A DOORWAY TO THE NORTH AND A DOORWAY TO THE EAST.">

**Type `TALK SECRETARY`**

<img src="images/missionasteroid/0007-1.png" width="400" alt="THE SECRETARY SAYS, &#34;HELLO&#34;. YOU ARE IN THE FRONT OFFICE OF THE BUILDING. THERE IS A SECRETARY HERE. THERE IS A DOORWAY TO THE NORTH AND A"><br>
<img src="images/missionasteroid/0007.png" width="400" alt="BUILDING. THERE IS A SECRETARY HERE. THERE IS A DOORWAY TO THE NORTH AND A DOORWAY TO THE EAST.">

**Type `SAY STARSTRUCK`**

<img src="images/missionasteroid/0008-1.png" width="400" alt="THE SECRETARY SAYS &#34;YOU MAY NOW PASS.&#34;. YOU ARE IN THE FRONT OFFICE OF THE BUILDING. THERE IS A SECRETARY HERE. THERE IS A DOORWAY TO THE NORTH AND A"><br>
<img src="images/missionasteroid/0008.png" width="400" alt="BUILDING. THERE IS A SECRETARY HERE. THERE IS A DOORWAY TO THE NORTH AND A DOORWAY TO THE EAST.">

**Type `NORTH`**

<img src="images/missionasteroid/0009.png" width="400" alt="YOU ARE AT A JUNCTION OF HALLWAYS. ONE HALLWAY GOES EAST AND WEST, THE OTHER GOES SOUTH.">

**Type `WEST`**

<img src="images/missionasteroid/0010.png" width="400" alt="THERE IS A GENERAL HERE. I THINK YOU HAD BETTER SALUTE. YOU ARE IN THE BRIEFING ROOM.">

**Type `SALUTE`**

<img src="images/missionasteroid/0011-1.png" width="400" alt="THE GENERAL SALUTES BACK AND SAYS &#34;AN ASTEROID IS ABOUT TO HIT THE EARTH. YOU MUST FLY TO THE ASTEROID AND BLOW IT UP. OF COURSE, THIS INFORMATION IS TOP"><br>
<img src="images/missionasteroid/0011-2.png" width="400" alt="SECRET. THE ASTEROID IS PROJECTED TO STRIKE THE EARTH AT 07:15 TONIGHT.&#34; HE THEN LEAVES. YOU ARE IN THE BRIEFING ROOM."><br>
<img src="images/missionasteroid/0011.png" width="400" alt="STRIKE THE EARTH AT 07:15 TONIGHT.&#34; HE THEN LEAVES. YOU ARE IN THE BRIEFING ROOM.">

**Type `EAST`**

<img src="images/missionasteroid/0012.png" width="400" alt="YOU ARE AT A JUNCTION OF HALLWAYS. ONE HALLWAY GOES EAST AND WEST, THE OTHER GOES SOUTH.">

**Type `EAST`**

<img src="images/missionasteroid/0013-1.png" width="400" alt="THERE IS A DISKETTE ON THE TABLE. YOU ARE IN THE COMPUTER ROOM. THERE IS AN APPLE COMPUTER HERE. THERE ARE DOORWAYS TO THE EAST AND WEST."><br>
<img src="images/missionasteroid/0013.png" width="400" alt="YOU ARE IN THE COMPUTER ROOM. THERE IS AN APPLE COMPUTER HERE. THERE ARE DOORWAYS TO THE EAST AND WEST.">

**Type `TAKE DISK`**

<img src="images/missionasteroid/0014.png" width="400" alt="YOU ARE IN THE COMPUTER ROOM. THERE IS AN APPLE COMPUTER HERE. THERE ARE DOORWAYS TO THE EAST AND WEST.">

**Type `INSERT DISK`**

<img src="images/missionasteroid/0015-1.png" width="400" alt="PUT DISKETTE INTO WHAT? (ANSWER QUESTION). YOU ARE IN THE COMPUTER ROOM. THERE IS AN APPLE COMPUTER HERE. THERE ARE"><br>
<img src="images/missionasteroid/0015.png" width="400" alt="YOU ARE IN THE COMPUTER ROOM. THERE IS AN APPLE COMPUTER HERE. THERE ARE DOORWAYS TO THE EAST AND WEST.">

**Type `INTO DRIVE`**

<img src="images/missionasteroid/0016-1.png" width="400" alt="UP ON THE MONITOR COMES YOUR FLIGHT PLAN: GO RIGHT FOR 10 MINUTES, UP FOR 5 MINUTES, LEFT FOR 15 MINUTES, DOWN FOR 5 MINUTES, LEFT FOR 5 MINUTES AND UP"><br>
<img src="images/missionasteroid/0016-2.png" width="400" alt="FOR 10 MINUTES. YOU ARE IN THE COMPUTER ROOM. THERE IS AN APPLE COMPUTER HERE. THERE ARE DOORWAYS TO THE EAST AND WEST."><br>
<img src="images/missionasteroid/0016.png" width="400" alt="YOU ARE IN THE COMPUTER ROOM. THERE IS AN APPLE COMPUTER HERE. THERE ARE DOORWAYS TO THE EAST AND WEST.">

**Type `EAST`**

<img src="images/missionasteroid/0017.png" width="400" alt="THERE ARE SOME EXPLOSIVES HERE. YOU ARE IN THE SUPPLY ROOM. THERE ARE DOORWAYS TO THE NORTH, EAST AND WEST.">

**Type `TAKE EXPLOSIVES`**

<img src="images/missionasteroid/0018.png" width="400" alt="SIVES YOU ARE IN THE SUPPLY ROOM. THERE ARE DOORWAYS TO THE NORTH, EAST AND WEST.">

**Type `LOOK EXPLOSIVES`**

<img src="images/missionasteroid/0019.png" width="400" alt="I SEE NOTHING SPECIAL. YOU ARE IN THE SUPPLY ROOM. THERE ARE DOORWAYS TO THE NORTH, EAST AND WEST.">

**Type `EAST`**

<img src="images/missionasteroid/0020.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN THE GYM. THERE ARE DOORWAYS TO THE NORTH, EAST AND WEST.">

**Type `EXERCISE`**

<img src="images/missionasteroid/0021.png" width="400" alt="OK. YOU NEEDED A GOOD WORKOUT. YOU ARE IN THE GYM. THERE ARE DOORWAYS TO THE NORTH, EAST AND WEST.">

**Type `EAST`**

<img src="images/missionasteroid/0022.png" width="400" alt="--------------- ENTER COMMAND?EAST YOU ARE IN THE SHOWER ROOM. THERE IS A DOORWAY TO THE WEST.">

**Type `TAKE SHOWER`**

<img src="images/missionasteroid/0023.png" width="400" alt="AHHHHH. THAT FEELS REFRESHING. YOU ARE IN THE SHOWER ROOM. THERE IS A DOORWAY TO THE WEST.">

**Type `WEST`**

<img src="images/missionasteroid/0024.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE IN THE GYM. THERE ARE DOORWAYS TO THE NORTH, EAST AND WEST.">

**Type `NORTH`**

<img src="images/missionasteroid/0025.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN A LONG TUNNEL. THERE IS A DOORWAY TO THE EAST.">

**Type `NORTH`**

<img src="images/missionasteroid/0026.png" width="400" alt="YOU ARE IN THE PRE-FLIGHT CHECKOUT ROOM. THERE ARE DOORWAYS TO THE WEST AND SOUTH (BEHIND YOU).">

**Type `TALK DOCTOR`**

<img src="images/missionasteroid/0027-1.png" width="400" alt="THE DOCTOR ASKS YOU HOW YOU ARE FEELING. YOU ARE IN THE PRE-FLIGHT CHECKOUT ROOM. THERE ARE DOORWAYS TO THE WEST"><br>
<img src="images/missionasteroid/0027.png" width="400" alt="YOU ARE IN THE PRE-FLIGHT CHECKOUT ROOM. THERE ARE DOORWAYS TO THE WEST AND SOUTH (BEHIND YOU).">

**Type `GOOD`**

<img src="images/missionasteroid/0028-1.png" width="400" alt="THATS GOOD. YOU ARE IN THE PRE-FLIGHT CHECKOUT ROOM. THERE ARE DOORWAYS TO THE WEST AND SOUTH (BEHIND YOU)."><br>
<img src="images/missionasteroid/0028.png" width="400" alt="YOU ARE IN THE PRE-FLIGHT CHECKOUT ROOM. THERE ARE DOORWAYS TO THE WEST AND SOUTH (BEHIND YOU).">

**Type `OPEN DOOR`**

<img src="images/missionasteroid/0029-1.png" width="400" alt="THE DOOR OPENS. YOU ARE IN THE PRE-FLIGHT CHECKOUT ROOM. THERE ARE DOORWAYS TO THE WEST AND SOUTH (BEHIND YOU)."><br>
<img src="images/missionasteroid/0029.png" width="400" alt="YOU ARE IN THE PRE-FLIGHT CHECKOUT ROOM. THERE ARE DOORWAYS TO THE WEST AND SOUTH (BEHIND YOU).">

**Type `WEST`**

### 2. Into space

<img src="images/missionasteroid/0030.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE ON THE AIRFIELD. YOU SEE A ROCKET IN THE DISTANCE.">

**Type `NORTH`**

<img src="images/missionasteroid/0031.png" width="400" alt="YOU ARE BELOW THE ROCKET. THERE IS A LADDER GOING UP. THERE IS A ROAD GOING SOUTH (BEHIND YOU).">

**Type `UP`**

<img src="images/missionasteroid/0032.png" width="400" alt="YOU ARE AT THE DOOR OF THE ROCKET. A LADDER IS GOING DOWN. THERE IS A BUTTON BY THE DOOR.">

**Type `PUSH BUTTON`**

<img src="images/missionasteroid/0033-1.png" width="400" alt="THE DOOR OPENS. YOU ARE AT THE DOOR OF THE ROCKET. A LADDER IS GOING DOWN. THERE IS A BUTTON BY THE DOOR."><br>
<img src="images/missionasteroid/0033.png" width="400" alt="YOU ARE AT THE DOOR OF THE ROCKET. A LADDER IS GOING DOWN. THERE IS A BUTTON BY THE DOOR.">

**Type `GO DOOR`**

<img src="images/missionasteroid/0034.png" width="400" alt="YOU ARE IN THE VACUUM LOCK ROOM OF THE ROCKET. THERE IS A BLUE AND AN ORANGE BUTTON ON THE WALL.">

**Type `PUSH BLUE`**

<img src="images/missionasteroid/0035-1.png" width="400" alt="THE NORTH DOOR OPENS AND THE SOUTH DOOR CLOSES AND THE LADDER ROLLS UP INTO THE ROCKET AND YOU HEAR A WHOOSHING SOUND. YOU ARE IN THE VACUUM LOCK ROOM OF THE"><br>
<img src="images/missionasteroid/0035.png" width="400" alt="YOU ARE IN THE VACUUM LOCK ROOM OF THE ROCKET. THERE IS A BLUE AND AN ORANGE BUTTON ON THE WALL.">

**Type `NORTH`**

<img src="images/missionasteroid/0036-1.png" width="400" alt="YOU ARE IN THE CONTROL ROOM OF THE ROCKET. THERE IS A CONSOLE AND A WINDOW HERE. THERE IS A VIOLET BUTTON ON THE WALL BESIDE YOU."><br>
<img src="images/missionasteroid/0036.png" width="400" alt="ROCKET. THERE IS A CONSOLE AND A WINDOW HERE. THERE IS A VIOLET BUTTON ON THE WALL BESIDE YOU.">

**Type `PUSH VIOLET`**

<img src="images/missionasteroid/0037-1.png" width="400" alt="THE DOOR CLOSES. YOU ARE IN THE CONTROL ROOM OF THE ROCKET. THERE IS A CONSOLE AND A WINDOW HERE. THERE IS A VIOLET BUTTON ON THE"><br>
<img src="images/missionasteroid/0037.png" width="400" alt="ROCKET. THERE IS A CONSOLE AND A WINDOW HERE. THERE IS A VIOLET BUTTON ON THE WALL BESIDE YOU.">

**Type `READ SIGN`**

<img src="images/missionasteroid/0038-1.png" width="400" alt="THE SIGN SAYS, &#34;TO OPERATE THIS ROCKET: PUSH THE THROTTLE FOR LIFTOFF, PULL THE THROTTLE TO LAND. PUSHING THE WHITE, BLACK, ORANGE AND BLUE BUTTONS WILL"><br>
<img src="images/missionasteroid/0038-2.png" width="400" alt="IGNITE THE BOOSTER ROCKETS TO MOVE THE ROCKET LEFT, RIGHT, UP OR DOWN RESPECTIVELY. YOU ARE IN THE CONTROL ROOM OF THE"><br>
<img src="images/missionasteroid/0038.png" width="400" alt="ROCKET. THERE IS A CONSOLE AND A WINDOW HERE. THERE IS A VIOLET BUTTON ON THE WALL BESIDE YOU.">

**Type `PUSH THROTTLE`**

<img src="images/missionasteroid/0039.png" width="400" alt="--------------- ENTER COMMAND?PUSH THROT TLE YOU ARE ORBITING THE EARTH.">

**Type `PUSH WHITE`**

<img src="images/missionasteroid/0040.png" width="400" alt="--------------- ENTER COMMAND?PUSH WHITE YOU ARE IN SPACE.">

**Type `PUSH WHITE`**

<img src="images/missionasteroid/0041.png" width="400" alt="--------------- ENTER COMMAND?PUSH WHITE YOU ARE IN SPACE.">

**Type `PUSH ORANGE`**

<img src="images/missionasteroid/0042.png" width="400" alt="--------------- ENTER COMMAND?PUSH ORANG E YOU ARE IN SPACE.">

**Type `PUSH ORANGE`**

<img src="images/missionasteroid/0043.png" width="400" alt="--------------- ENTER COMMAND?PUSH ORANG E YOU ARE ORBITTING THE ASTEROID.">

**Type `PULL THROTTLE`**

<img src="images/missionasteroid/0044.png" width="400" alt="--------------- ENTER COMMAND?PULL THROT TLE YOU ARE ON THE ASTEROID.">

**Type `WEST`**

<img src="images/missionasteroid/0045.png" width="400" alt="THERE IS A SPACESUIT HERE. YOU ARE IN THE STORAGE ROOM OF THE ROCKET. THERE IS A DOORWAY TO THE EAST.">

**Type `GET SUIT`**

<img src="images/missionasteroid/0046.png" width="400" alt="--------------- ENTER COMMAND?GET SUIT YOU ARE IN THE STORAGE ROOM OF THE ROCKET. THERE IS A DOORWAY TO THE EAST.">

**Type `WEAR SUIT`**

<img src="images/missionasteroid/0047.png" width="400" alt="YOU ARE WEARING THE SUIT. YOU ARE IN THE STORAGE ROOM OF THE ROCKET. THERE IS A DOORWAY TO THE EAST.">

**Type `EAST`**

<img src="images/missionasteroid/0048.png" width="400" alt="ROCKET. THERE IS A DOORWAY TO THE EAST. --------------- ENTER COMMAND?EAST YOU ARE ON THE ASTEROID.">

**Type `PUSH VIOLET`**

### 3. The asteroid

<img src="images/missionasteroid/0049.png" width="400" alt="T THE DOOR OPENS. YOU ARE ON THE ASTEROID.">

**Type `SOUTH`**

<img src="images/missionasteroid/0050.png" width="400" alt="YOU ARE IN THE VACUUM LOCK ROOM OF THE ROCKET. THERE IS A BLUE AND AN ORANGE BUTTON ON THE WALL.">

**Type `SET TIMER`**

<img src="images/missionasteroid/0051-1.png" width="400" alt="TO HOW MANY MINUTES? (90, 120, 150 OR 180)? YOU ARE IN THE VACUUM LOCK ROOM OF THE ROCKET. THERE IS A BLUE AND AN ORANGE"><br>
<img src="images/missionasteroid/0051.png" width="400" alt="YOU ARE IN THE VACUUM LOCK ROOM OF THE ROCKET. THERE IS A BLUE AND AN ORANGE BUTTON ON THE WALL.">

**Type `150`**

<img src="images/missionasteroid/0052-1.png" width="400" alt="OK YOU ARE IN THE VACUUM LOCK ROOM OF THE ROCKET. THERE IS A BLUE AND AN ORANGE BUTTON ON THE WALL."><br>
<img src="images/missionasteroid/0052.png" width="400" alt="YOU ARE IN THE VACUUM LOCK ROOM OF THE ROCKET. THERE IS A BLUE AND AN ORANGE BUTTON ON THE WALL.">

**Type `TURN DIAL`**

<img src="images/missionasteroid/0053-1.png" width="400" alt="THE AIR IS ON. YOU ARE IN THE VACUUM LOCK ROOM OF THE ROCKET. THERE IS A BLUE AND AN ORANGE BUTTON ON THE WALL."><br>
<img src="images/missionasteroid/0053.png" width="400" alt="YOU ARE IN THE VACUUM LOCK ROOM OF THE ROCKET. THERE IS A BLUE AND AN ORANGE BUTTON ON THE WALL.">

**Type `PUSH ORANGE`**

<img src="images/missionasteroid/0054-1.png" width="400" alt="THE NORTH DOOR CLOSES AND THE SOUTH DOOR OPENS AND THE LADDER UNROLLS TO THE GROUND AND YOU HEAR A WHOOSHING SOUND."><br>
<img src="images/missionasteroid/0054.png" width="400" alt="YOU ARE IN THE VACUUM LOCK ROOM OF THE ROCKET. THERE IS A BLUE AND AN ORANGE BUTTON ON THE WALL.">

**Type `SOUTH`**

<img src="images/missionasteroid/0055.png" width="400" alt="YOU ARE AT THE DOOR OF THE ROCKET. A LADDER IS GOING DOWN. THERE IS A BUTTON BY THE DOOR.">

**Type `DOWN`**

<img src="images/missionasteroid/0056.png" width="400" alt="--------------- ENTER COMMAND?DOWN YOU ARE BELOW THE ROCKET. THERE IS A LADDER GOING UP.">

**Type `SOUTH`**

<img src="images/missionasteroid/0057.png" width="400" alt="LADDER GOING UP. --------------- ENTER COMMAND?SOUTH YOU ARE ON A SMALL ASTEROID.">

**Type `WEST`**

<img src="images/missionasteroid/0058.png" width="400" alt="--------------- ENTER COMMAND?WEST YOU ARE ON A SMALL ASTEROID. THERE IS A CAVE HERE.">

**Type `GO CAVE`**

<img src="images/missionasteroid/0059.png" width="400" alt="--------------- ENTER COMMAND?GO CAVE YOU ARE IN A CAVE. A PASSAGEWAY GOES SOUTH.">

**Type `SOUTH`**

<img src="images/missionasteroid/0060.png" width="400" alt="YOU ARE IN A CAVE. IT ENDS RIGHT HERE. THERE IS A VERY DEEP PIT HERE. A PASSAGEWAY GOES NORTH.">

**Type `DROP EXPLOSIVES`**

<img src="images/missionasteroid/0061-1.png" width="400" alt="WHERE? IN THE PIT OR ON THE FLOOR? YOU ARE IN A CAVE. IT ENDS RIGHT HERE. THERE IS A VERY DEEP PIT HERE. A PASSAGEWAY GOES NORTH."><br>
<img src="images/missionasteroid/0061.png" width="400" alt="YOU ARE IN A CAVE. IT ENDS RIGHT HERE. THERE IS A VERY DEEP PIT HERE. A PASSAGEWAY GOES NORTH.">

**Type `IN PIT`**

<img src="images/missionasteroid/0062-1.png" width="400" alt="OK. YOU HAVE DROPPED THE EXPLOSIVES INTO THE PIT. YOU ARE IN A CAVE. IT ENDS RIGHT HERE. THERE IS A VERY DEEP PIT HERE. A"><br>
<img src="images/missionasteroid/0062.png" width="400" alt="YOU ARE IN A CAVE. IT ENDS RIGHT HERE. THERE IS A VERY DEEP PIT HERE. A PASSAGEWAY GOES NORTH.">

**Type `NORTH`**

<img src="images/missionasteroid/0063.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE IN A CAVE. A PASSAGEWAY GOES SOUTH.">

**Type `NORTH`**

<img src="images/missionasteroid/0064.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE ON A SMALL ASTEROID. THERE IS A CAVE HERE.">

**Type `SOUTH`**

<img src="images/missionasteroid/0065.png" width="400" alt="CAVE HERE. --------------- ENTER COMMAND?SOUTH YOU ARE ON A SMALL ASTEROID.">

**Type `NORTH`**

<img src="images/missionasteroid/0066.png" width="400" alt="--------------- ENTER COMMAND?NORTH YOU ARE BELOW THE ROCKET. THERE IS A LADDER GOING UP.">

**Type `UP`**

<img src="images/missionasteroid/0067.png" width="400" alt="YOU ARE AT THE DOOR OF THE ROCKET. A LADDER IS GOING DOWN. THERE IS A BUTTON BY THE DOOR.">

**Type `GO DOOR`**

### 4. Back to Earth

<img src="images/missionasteroid/0068.png" width="400" alt="YOU ARE IN THE VACUUM LOCK ROOM OF THE ROCKET. THERE IS A BLUE AND AN ORANGE BUTTON ON THE WALL.">

**Type `PUSH BLUE`**

<img src="images/missionasteroid/0069-1.png" width="400" alt="THE NORTH DOOR OPENS AND THE SOUTH DOOR CLOSES AND THE LADDER ROLLS UP INTO THE ROCKET AND YOU HEAR A WHOOSHING SOUND. YOU ARE IN THE VACUUM LOCK ROOM OF THE"><br>
<img src="images/missionasteroid/0069.png" width="400" alt="YOU ARE IN THE VACUUM LOCK ROOM OF THE ROCKET. THERE IS A BLUE AND AN ORANGE BUTTON ON THE WALL.">

**Type `TURN DIAL`**

<img src="images/missionasteroid/0070-1.png" width="400" alt="THE AIR IS OFF. YOU ARE IN THE VACUUM LOCK ROOM OF THE ROCKET. THERE IS A BLUE AND AN ORANGE BUTTON ON THE WALL."><br>
<img src="images/missionasteroid/0070.png" width="400" alt="YOU ARE IN THE VACUUM LOCK ROOM OF THE ROCKET. THERE IS A BLUE AND AN ORANGE BUTTON ON THE WALL.">

**Type `NORTH`**

<img src="images/missionasteroid/0071.png" width="400" alt="BUTTON ON THE WALL. --------------- ENTER COMMAND?NORTH YOU ARE ON THE ASTEROID.">

**Type `PUSH THROTTLE`**

<img src="images/missionasteroid/0072.png" width="400" alt="--------------- ENTER COMMAND?PUSH THROT TLE YOU ARE ORBITTING THE ASTEROID.">

**Type `PUSH BLUE`**

<img src="images/missionasteroid/0073.png" width="400" alt="YOU ARE ORBITTING THE ASTEROID. --------------- ENTER COMMAND?PUSH BLUE YOU ARE IN SPACE.">

**Type `PUSH BLUE`**

<img src="images/missionasteroid/0074.png" width="400" alt="YOU ARE IN SPACE. --------------- ENTER COMMAND?PUSH BLUE YOU ARE IN SPACE.">

**Type `PUSH BLACK`**

<img src="images/missionasteroid/0075.png" width="400" alt="--------------- ENTER COMMAND?PUSH BLACK YOU ARE IN SPACE.">

**Type `PUSH BLACK`**

<img src="images/missionasteroid/0076.png" width="400" alt="--------------- ENTER COMMAND?PUSH BLACK YOU ARE IN SPACE.">

**Type `PUSH WHITE`**

<img src="images/missionasteroid/0077.png" width="400" alt="--------------- ENTER COMMAND?PUSH WHITE YOU ARE ORBITING THE EARTH.">

**Type `PULL THROTTLE`**

<img src="images/missionasteroid/0078-1.png" width="400" alt="YOU ARE IN THE CONTROL ROOM OF THE ROCKET. THERE IS A CONSOLE AND A WINDOW HERE. THERE IS A VIOLET BUTTON ON THE WALL BESIDE YOU."><br>
<img src="images/missionasteroid/0078.png" width="400" alt="ROCKET. THERE IS A CONSOLE AND A WINDOW HERE. THERE IS A VIOLET BUTTON ON THE WALL BESIDE YOU.">

**Type `LOOK WATCH`**

<img src="images/missionasteroid/0079-1.png" width="400" alt="THE TIME IS 06:25 YOU ARE IN THE CONTROL ROOM OF THE ROCKET. THERE IS A CONSOLE AND A WINDOW"><br>
<img src="images/missionasteroid/0079.png" width="400" alt="ROCKET. THERE IS A CONSOLE AND A WINDOW HERE. THERE IS A VIOLET BUTTON ON THE WALL BESIDE YOU.">

**Type `LOOK WATCH`**

<img src="images/missionasteroid/0080-1.png" width="400" alt="THE TIME IS 06:30 YOU ARE IN THE CONTROL ROOM OF THE ROCKET. THERE IS A CONSOLE AND A WINDOW"><br>
<img src="images/missionasteroid/0080.png" width="400" alt="ROCKET. THERE IS A CONSOLE AND A WINDOW HERE. THERE IS A VIOLET BUTTON ON THE WALL BESIDE YOU.">

**Type `LOOK WATCH`**

### The end

<img src="images/missionasteroid/0081-1.png" width="400" alt="THE TIME IS 06:35 THE ASTEROID HAS EXPLODED! YOU HAVE SAVED THE EARTH!!!"><br>
<img src="images/missionasteroid/0081-2.png" width="400" alt="KEN AND ROBERTA WILLIAMS THANK YOU FOR PLAYING MISSION: ASTEROID. GOOD BYE.... YOU ARE IN THE CONTROL ROOM OF THE ROCKET. THERE IS A CONSOLE AND A WINDOW"><br>
<img src="images/missionasteroid/0081.png" width="400" alt="ROCKET. THERE IS A CONSOLE AND A WINDOW HERE. THERE IS A VIOLET BUTTON ON THE WALL BESIDE YOU.">

<!-- End of the walkthrough -->

## What next

[Mystery House](mysteryhouse.md) is the first of the Hi-Res Adventures, and
[Time Zone](timezone.md) the fifth, the biggest.
