package main

import (
	"github.com/neet-007/chess_engine_go/internal/board"
	"testing"
)

func TestBoard(t *testing.T) {
	cases := []struct {
		name string
		fen  string
	}{
		{
			name: "Starting Position",
			fen:  "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		},
		{
			name: "Midgame - No Castling",
			fen:  "r1bk3r/p2pBpNp/n5p1/1ppNP2P/6P1/3P4/P1P1K3/q5b1 b - - 0 1",
		},
		{
			name: "En Passant Active",
			fen:  "rnbqkbnr/pppppp1p/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1",
		},
		{
			name: "All Pieces on Board",
			fen:  "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		},
		{
			name: "Empty Board (Illegal but good for testing)",
			fen:  "8/8/8/8/8/8/8/8 w - - 0 1",
		},
		{
			name: "Promotion & Endgame",
			fen:  "8/P7/8/1k6/8/8/5K2/8 w - - 0 1",
		},
		{
			name: "Kiwipete (Standard Move Gen Test)",
			fen:  "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1",
		},
	}

	for _, c := range cases {
		b := board.NewBoard()
		if err := parseFEN(b, c.fen); err != nil {
			t.Errorf("Name: %s\nFEN: %s\nError: %s", c.name, c.fen, err)
			continue
		}

		for rank := range uint8(8) {
			rank = 7 - rank

			for file := range uint8(8) {
				index := board.GetSquareIndex(file, rank)
				found := false
				var foundPiece board.Piece

				for i := range 12 {
					if b.Bitboards[i]&(1<<index) != 0 {
						if found {
							t.Errorf("Name: %s\nFEN: %s\nFound %s first and then %s", c.name, c.fen, foundPiece, board.Piece(i))
							return
						}

						found = true
						foundPiece = board.Piece(i)
					}
				}
			}
		}

		if b.Occupied != ^b.Empty {
			t.Errorf("Name: %s\nFEN: %s\nOccupied: %b\nEmpty: %b", c.name, c.fen, b.Occupied, b.Empty)
			continue
		}

		if b.Occupied != (b.Bitboards[board.White_all] | b.Bitboards[board.Black_all]) {
			t.Errorf("Name: %s\nFEN: %s\nOccupied: %b\nAll: %b", c.name, c.fen, b.Occupied, b.Bitboards[board.White_all]|b.Bitboards[board.Black_all])
			continue
		}

		serialized := serializeFEN(b)
		if serialized != c.fen {
			t.Errorf("Name: %s\nFEN: %s\nSerialized: %s", c.name, c.fen, serialized)
			continue
		}
	}
}

func TestMoveEncode(t *testing.T) {
	cases := []struct {
		from  board.Square
		to    board.Square
		flags uint16
	}{
		{
			board.A1, board.B1, 0,
		},
		{
			board.A1, board.B1, board.KnightPromoCaptureFlag,
		},
		{
			board.H8, board.A8, board.EpCaptureFlag,
		},
	}

	for _, c := range cases {
		m := board.NewMove(c.from, c.to, c.flags)
		if m.From() != int(c.from) {
			t.Errorf("in Move %b Expected from %b found %b", m, c.from, m.From())
		}

		if m.To() != int(c.to) {
			t.Errorf("in Move %b Expected to %b found %b", m, c.to, m.To())
		}

		if m.Flags() != int(c.flags) {
			t.Errorf("in Move %b Expected flags %b found %b", m, c.flags, m.Flags())
		}
	}
}
