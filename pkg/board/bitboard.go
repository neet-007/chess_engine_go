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

// Little endian rank-file (LERF)
const (
	a1 Bitboard = iota
	b1
	c1
	d1
	e1
	f1
	g1
	h1
	a2
	b2
	c2
	d2
	e2
	f2
	g2
	h2
	a3
	b3
	c3
	d3
	e3
	f3
	g3
	h3
	a4
	b4
	c4
	d4
	e4
	f4
	g4
	h4
	a5
	b5
	c5
	d5
	e5
	f5
	g5
	h5
	a6
	b6
	c6
	d6
	e6
	f6
	g6
	h6
	a7
	b7
	c7
	d7
	e7
	f7
	g7
	h7
	a8
	b8
	c8
	d8
	e8
	f8
	g8
	h8
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
