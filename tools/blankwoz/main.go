/*
Blankwoz writes a blank diskette as a WOZ image, to standard output: a 5.25"
diskette that was never formatted, its 35 tracks of 51,200 bits all zero.
The drive of izapple2 reads noise where a track has no bits, as a real drive
does, and writes on the tracks of the image, so it is a diskette for the
programs that want one never formatted, as the SwyftCard.

	go run ./tools/blankwoz > blank.woz

The image is the same on every run, so its checksum can be listed in
disks.tsv; fetch-disks.sh makes it from there.
*/
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"os"
)

const (
	tracks      = 35
	trackBlocks = 13    // 512 bytes each, 6,656 bytes for each track
	trackBits   = 51200 // the bits of a track of a 5.25" diskette
	firstBlock  = 3     // the bits of the tracks start at 3*512 in the file
)

func main() {
	if _, err := os.Stdout.Write(blankWoz()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// blankWoz is the WOZ 2 image of a blank diskette
func blankWoz() []byte {
	var chunks bytes.Buffer
	chunk := func(id string, data []byte) {
		chunks.WriteString(id)
		binary.Write(&chunks, binary.LittleEndian, uint32(len(data)))
		chunks.Write(data)
	}

	// INFO: version 2, a 5.25" diskette, not write protected, one side,
	// the bits 4 microseconds apart, the largest track of 13 blocks
	info := make([]byte, 60)
	info[0] = 2
	info[1] = 1
	copy(info[5:37], fmt.Sprintf("%-32s", "apple2-activities blank"))
	info[37] = 1
	info[39] = 32
	binary.LittleEndian.PutUint16(info[44:], trackBlocks)
	chunk("INFO", info)

	// TMAP: each track read from its quarter track and the ones beside it
	tmap := bytes.Repeat([]byte{0xFF}, 160)
	for t := range tracks {
		for q := 4*t - 1; q <= 4*t+1; q++ {
			if q >= 0 {
				tmap[q] = byte(t)
			}
		}
	}
	chunk("TMAP", tmap)

	// TRKS: where the bits of each track are, and the bits, all zero
	var trks bytes.Buffer
	for t := range 160 {
		if t < tracks {
			binary.Write(&trks, binary.LittleEndian, uint16(firstBlock+t*trackBlocks))
			binary.Write(&trks, binary.LittleEndian, uint16(trackBlocks))
			binary.Write(&trks, binary.LittleEndian, uint32(trackBits))
		} else {
			trks.Write(make([]byte, 8))
		}
	}
	trks.Write(make([]byte, tracks*trackBlocks*512))
	chunk("TRKS", trks.Bytes())

	var woz bytes.Buffer
	woz.WriteString("WOZ2\xff\n\r\n")
	binary.Write(&woz, binary.LittleEndian, crc32.ChecksumIEEE(chunks.Bytes()))
	woz.Write(chunks.Bytes())
	return woz.Bytes()
}
