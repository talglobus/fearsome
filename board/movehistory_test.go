package board

import (
	"fmt"
	"testing"
)

func ExampleHistory_String() {
	fmt.Println(History{1, 2, 3})
	// Output: RED 1 ⇨ BLUE 2 ⇨ RED 3
}

func ExampleHistory_String_longer() {
	fmt.Println(History{0, 1, 0, 2, 0, 3, 0})
	// Output: RED 0 ⇨ BLUE 1 ⇨ RED 0 ⇨ BLUE 2 ⇨ RED 0 ⇨ BLUE 3 ⇨ RED 0
}

func ExampleHistory_String_empty() {
	fmt.Println(History{})
	// Output: NO MOVES
}

func TestHistory_Equals(t *testing.T) {
	table := []struct {
		h1, h2 History
		want   bool
	}{
		{
			History{1, 6, 2, 5, 3},
			History{1, 6, 2, 5, 3},
			true,
		}, {
			History{3},
			History{3},
			true,
		}, {
			History{},
			History{},
			true,
		}, {
			History{7, 6, 5, 4},
			History{7, 6, 5, 3},
			false,
		}, {
			History{},
			History{1},
			false,
		}, {
			History{4},
			History{},
			false,
		}, {
			History{6, 2, 5, 3, 1},
			History{1, 6, 2, 5, 3},
			false,
		},
	}

	for _, r := range table {
		if got := r.h1.Equals(r.h2); got != r.want {
			t.Errorf("Equals output incorrect for histories %v and %v. Expected %v, observed %v",
				r.h1, r.h2, r.want, got)
		}
	}
}
