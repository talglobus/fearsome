package board

import "fmt"

// Move holds the choice of column for a given move
type Move int

// History holds a sequence of moves, with even-indexed moves red and odd-indexed moves blue, in sequential order
type History []Move

// String outputs a fancy format for the move history, consisting of type-labeled, arrow-separated moves in sequence
func (h History) String() string {
	str := ""
	isRed := true

	for _, m := range h {
		if isRed {
			str += fmt.Sprintf("RED %v ⇨ ", m)
		} else {
			str += fmt.Sprintf("BLUE %v ⇨ ", m)
		}
		isRed = !isRed
	}

	if len(h) == 0 {
		return "NO MOVES"
	}

	// Return move sequence without trailing spaces and arrow
	return str[:len(str)-4]
}

// Equals tests equality for two move histories
func (h History) Equals(h2 History) bool {
	if len(h) != len(h2) {
		return false
	}

	for i := range h {
		if h[i] != h2[i] {
			return false
		}
	}

	return true
}
