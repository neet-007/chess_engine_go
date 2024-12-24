package board

import (
	"fmt"
	"math/bits"
	"strings"
)

type Bitboard uint64

func (b Bitboard) Count() int {
	return bits.OnesCount64(uint64(b))
}

func (b *Bitboard) Set(pos uint) {
	*b |= Bitboard(uint64(1) << pos)
}

func (b Bitboard) IsNotZero(pos uint) bool {
	return (b & Bitboard(uint64(1)<<pos)) != 0
}

func (b *Bitboard) Clr(pos uint) {
	*b &= Bitboard(^uint64(1) << pos)
}

func (b *Bitboard) FirstOne() int {
	bit := bits.TrailingZeros64(uint64(*b))

	if bit == 64 {
		return bit
	}

	*b = (*b >> uint(bit+1)) << uint(bit+1)
	return bit
}

func (b Bitboard) String() string {
	zeroes := ""
	for ix := 0; ix < 64; ix++ {
		zeroes = zeroes + "0"
	}

	bits := zeroes + fmt.Sprintf("%b", b)
	return bits[len(bits)-64:]
}

func (b Bitboard) Stringln() string {
	s := b.String()
	row := [8]string{}
	row[0] = s[0:8]
	row[1] = s[8:16]
	row[2] = s[16:24]
	row[3] = s[24:32]
	row[4] = s[32:40]
	row[5] = s[40:48]
	row[6] = s[48:56]
	row[7] = s[56:]
	for ix, r := range row {
		row[ix] = fmt.Sprintf("%v%v%v%v%v%v%v%v\n", r[7:8], r[6:7], r[5:6], r[4:5], r[3:4], r[2:3], r[1:2], r[0:1])
	}

	s = strings.Join(row[:], "")
	s = strings.Replace(s, "1", "1 ", -1)
	s = strings.Replace(s, "0", "0 ", -1)
	return s
}

/*
	squareIndex = 8*rankIndex + fileIndex
	FileIndex   = squareIndex modulo 8  = squareIndex & 7
	RankIndex   = squareIndex div    8  = squareIndex >> 3
*/

// Bitboard Constants for Files, Ranks, and Patterns
const (
	AFile      Bitboard = 0x0101010101010101
	HFile      Bitboard = 0x8080808080808080
	FirstRank  Bitboard = 0x00000000000000FF
	EighthRank Bitboard = 0xFF00000000000000
)

// Little endian rank-file (LERF)
const (
	A1 int = iota
	B1
	C1
	D1
	E1
	F1
	G1
	H1
	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A8
	B8
	C8
	D8
	E8
	F8
	G8
	H8
)

/*
	Some bitboard constants with LERF-mapping:

	a-file             0x0101010101010101
	h-file             0x8080808080808080
	1st rank           0x00000000000000FF
	8th rank           0xFF00000000000000
	a1-h8 diagonal     0x8040201008040201
	h1-a8 antidiagonal 0x0102040810204080
	light squares      0x55AA55AA55AA55AA
	dark squares       0xAA55AA55AA55AA55
*/
