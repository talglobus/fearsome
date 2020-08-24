package game

import "math/rand"

// Type is an enumerated triad type consisting of red piece, blue piece, or no piece.
// The wisdom of naming a type `Type` is generally questionable, but here is actually quite reasonable
type Type int

// NONE, RED, and BLUE are used to indicate the piece type of a square, with NONE also acting as a valid nil value
const (
	NONE Type = iota
	RED
	BLUE
)

// String enables for a `Type` to be serialized to string format
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

// enumStatic is an interface for creating singletons exposing static methods on enumerated types
type enumStatic interface {
	Rand() Type
	Count() int
}

type typeStatic struct{}

// Count returns the number of valid values in the Type enum, allowing for simple and easy reflection-like behavior
func (typeStatic) Count() int {
	return 3
}

// Rand randomly constructs a new Type
func (typeStatic) Rand() Type {
	return Type(rand.Intn(typeStatic{}.Count()))
}

// TYPE is a singleton exposing static methods on Type
var TYPE enumStatic = typeStatic{}
