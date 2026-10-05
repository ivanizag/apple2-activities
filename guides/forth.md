# Forth in ROM

[Back to the activities](../README.md)

Forth is a language made of *words*: each does something to a stack of
numbers, and a program is new words defined from the ones there are, until
there is one for the whole job. It needs little memory and runs fast, close
to the machine, and in the early eighties it had a following among owners
of small computers who found BASIC slow and assembler hard.

Offete Industries sold a card that put FORTH-79 in the ROM of an Apple \]\[+,
there when the machine is switched on, the way Applesoft is. This page
switches on that machine, uses some words and defines new ones.

## What you need

Only izapple2: the ROM of the Forth card comes inside it. The
[reference manual of this Forth](https://archive.org/details/forth79-A2) is
on the Internet Archive.

## The machine

An Apple \]\[+ with Forth in ROM:

- the 6502 processor at 1 MHz, 48 KB of memory, and Applesoft BASIC in its ROM;
- the Forth ROM card of Offete Industries in slot 0;
- a Videx Videoterm 80 column card in slot 3, unused here;
- no disk controller card.

```bash
izapple2 -model forth
```

## Switch it on

1. **Start izapple2** with the command above. The Apple \]\[+ writes its
   name, and the card starts Forth, which says its version and waits.

   ![Forth switched on](images/forth/switched-on.png)

## Words

2. **Type these lines**, each with Return:

   ```forth
   2 3 + .
   : SQUARE DUP * ;
   7 SQUARE .
   : STARS 0 DO 42 EMIT LOOP ;
   10 STARS
   ```

   ![Words used and defined](images/forth/words.gif)

   - `2 3 + .` puts 2 and 3 on the stack, `+` adds them, and `.` prints
     the result and takes it off. Forth answers `OK` after each line.
   - `: SQUARE DUP * ;` defines a new word, `SQUARE`, from `:` to `;`: it
     duplicates the number on the stack with `DUP` and multiplies it by
     itself. `7 SQUARE .` prints 49.
   - `STARS` loops from 0 to the number on the stack, and each time `EMIT`
     writes the character 42, an asterisk.

## The dictionary

3. **Type `VLIST`.** Forth lists every word it knows, the newest first: the
   two you defined, and then those of the ROM, down to the first ones, the
   few in machine code that the rest are built from, `DUP`, `SWAP`, `+`,
   `EMIT`, `KEY`.

   ![The end of the dictionary](images/forth/vlist.png)

## What next

[The Apple \]\[ of 1977](apple-ii.md) has the other way of programming the
machine close to the metal, its Mini-Assembler.
