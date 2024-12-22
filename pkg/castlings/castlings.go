package castlings

import (
	"strings"
)

type Castlings uint

const (
	ShortW = uint(0x1)
	LongW  = uint(0x2)
	ShortB = uint(0x4)
	LongB  = uint(0x8)
)

func ParseCastlings(fenCast string) Castlings {
	c := uint(0)

	if fenCast == "-" {
		return Castlings(0)
	}

	if strings.Contains(fenCast, "K") {
		c |= ShortW
	}

	if strings.Contains(fenCast, "Q") {
		c |= LongW
	}

	if strings.Contains(fenCast, "k") {
		c |= ShortB
	}

	if strings.Contains(fenCast, "Q") {
		c |= LongB
	}

	return Castlings(c)
}
