# Apple Presents the IIc: The Apple at Work

[Activities](../README.md) › [First steps](../sections/first-steps.md) › [Switch on an Apple IIc](apple-iic.md)

*The Apple at Work* is the diskette of **Apple Presents the IIc**, the
owner's guide of the IIc, on the programs of work: the word processor, the
data base and the spreadsheet, the three of AppleWorks, Apple's program of
1984 that had them all, and what puts them together. It does not run
AppleWorks: it imitates it, a screen at a time, and asks for the keys a
person would press.

This page goes through its introduction, a survey and a quiz, and the lesson
of the word processor, as far as a letter changed for other friends, and
then the same letter on the side of the diskette for 80 columns.

## What you need

- **`Apple Presents The IIc - The Apple At Work Side 2 40-Col 691-0014-A (Apple-1984).dsk`**
- **`Apple Presents The IIc - The Apple At Work Side 1 80-Col 691-0014-A (Apple-1984).dsk`**

The two sides of the diskette, from the folder of the masters of each model
on
[Asimov](https://mirrors.apple2.org.za/ftp.apple.asimov.net/images/masters/model_specific/).
`./fetch-disks.sh` in this repository downloads them into `disks/`.

Side 1 is for 80 columns, for a monitor, and side 2 for 40 columns, for a
television. The IIc has a switch above its keyboard, 80/40, which the
programs read when they start: this page goes through side 2, with the
switch down, and then looks at side 1 with it up.

## The machine

The Apple IIc of [Switch on an Apple IIc](apple-iic.md), on a colour
monitor, with the diskette in its drive:

```bash
izapple2 -model none -board 2c -cpu 65c02 -screen color \
    -rom "<internal>/Apple2c_ROM255.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s1 iicserial \
    -s2 iicserial \
    -s6 "iicdisk,disk1=disks/Apple Presents The IIc - The Apple At Work Side 2 40-Col 691-0014-A (Apple-1984).dsk"
```

For side 1, the same IIc with its switch up, on 80 columns:

```bash
izapple2 -model none -board 2c -cpu 65c02 -screen color \
    -rom "<internal>/Apple2c_ROM255.rom" \
    -charrom "<internal>/Apple IIe Video Enhanced.bin" \
    -s0 language \
    -s1 iicserial \
    -s2 iicserial \
    -s6 "iicdisk,disk1=disks/Apple Presents The IIc - The Apple At Work Side 1 80-Col 691-0014-A (Apple-1984).dsk" \
    -iic80ColumnSwitch
```

## The introduction

1. **Start izapple2** with the command above: the title, *writing, figuring,
   filing*, and then the menu.

   ![The title](images/iic-at-work/title.png)

   ![The menu](images/iic-at-work/menu.png)

2. **Press Return**, for the introduction, *or what's an application
   program, anyway?*, and **Return** at each page: a desk in a mess, and
   then in order, with the Apple on it.

   ![The introduction](images/iic-at-work/introduction.png)

   ![A desk in a mess](images/iic-at-work/mess.png)

   ![A desk in order](images/iic-at-work/order.png)

3. Before going on, the guide asks to complete the *Ziegfeld Organizational
   Assessment Profile*.

   ![The Ziegfeld](images/iic-at-work/ziegfeld.png)

   **Press Return** twice, and answer its five questions: **type the
   letter** and **Return**. This page answers `B` to all.

   ![A question of the survey](images/iic-at-work/survey.png)

   ![Congratulations](images/iic-at-work/congratulations.png)

4. **Press Return** four times: a quiz on the three kinds of programs. The
   answers are `B`, a word processor, `A`, a data base program, and `A`, a
   spreadsheet.

   ![The quiz](images/iic-at-work/quiz.png)

5. **Press Return** through the pages after it. Harry and Nan Walker will
   show the programs, which the guide shows as they are in AppleWorks.

   ![The Walkers](images/iic-at-work/walkers.png)

   ![AppleWorks](images/iic-at-work/appleworks.png)

## The word processor

6. **Press Return** through the pages, to the letter Harry wrote to Pat and
   Dick Peterson in the word processor.

   ![The letter](images/iic-at-work/letter.png)

7. **Press Return**, and do what the guide asks: **the down arrow** to the
   end of the letter, **Return** twice, **Open Apple-1** back to its
   beginning, **the right arrow five times** to the names, and **type
   `Frank and Leslie,`** over them.

   ![The lesson](images/iic-at-work/lesson.gif)

   The letter is for Frank and Leslie Burns now. They have no kids, and
   the lesson goes on, to change the line about them.

   ![Personalized](images/iic-at-work/personalized.png)

8. **Hold Open Apple and press Escape**, which goes back to the menu from
   anywhere in the guide. The menu marks the topic done.

   ![The menu](images/iic-at-work/done.png)

   The other topics are the data base, the spreadsheet and the integrated
   software. **Type `6`** and **Return** to quit: the guide
   asks for the next diskette.

## Side 1, in 80 columns

9. **Quit izapple2 and start it with the second command**: the same title,
   and the menu in 80 columns.

   ![The menu in 80 columns](images/iic-at-work/menu-80.png)

10. **Type `2`** and **Return**, the word processor, and **Return** at each
    page: Harry's letter, as wide as AppleWorks writes it on a monitor, the
    whole of it on the screen.

    ![The letter in 80 columns](images/iic-at-work/letter-80.png)

11. **Press Return** at each page, until the guide asks for the right
    arrow: the same lesson. **Press the right arrow five times**, and **type
    `Frank and Leslie,`** over the old names.

    ![The letter personalized](images/iic-at-work/personalized-80.png)

## What next

**Next:** [Apple Presents the IIc: The Apple at Play](iic-at-play.md).

[AppleWorks](appleworks.md) is the program itself, on a //e, with a
spreadsheet, a letter and a list of friends.
