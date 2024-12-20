package board

type Bitboard uint64

/*
	squareIndex = 8*rankIndex + fileIndex
	FileIndex   = squareIndex modulo 8  = squareIndex & 7
	RankIndex   = squareIndex div    8  = squareIndex >> 3
*/

func NewBitboard() Bitboard {
	return 0
}

// Bitboard Constants for Files, Ranks, and Patterns
const (
	AFile      Bitboard = 0x0101010101010101
	HFile      Bitboard = 0x8080808080808080
	FirstRank  Bitboard = 0x00000000000000FF
	EighthRank Bitboard = 0xFF00000000000000
)

// Little endian rank-file (LERF)
const (
	A1 Bitboard = iota
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
