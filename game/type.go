package game

import (
	"math/rand"
)

// Type is an enumerated triad type consisting of red piece, blue piece, or no piece.
// The wisdom of naming a type `Type` is generally questionable, but here is actually quite reasonable
type Type int

const (
	NONE Type = iota
	RED
	BLUE
)

// newRandType() randomly constructs a new `Type`
func newRandType() Type {
	return Type(rand.Intn(3))
}

// String() enables for a `Type` to be serialized to string format
func (t Type) String() string {
	switch t {
	case RED:
		return "RED"
	case BLUE:
		return "BLÜ"
	default:
		return "---"
	}
}
