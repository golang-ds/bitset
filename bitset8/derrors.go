package bitset

import "errors"

var (
	ErrNumberNotFit8 = errors.New("the given number doesn't fit 8-bit set. It must be in range of 0-7")
)
