package board

import (
	"fmt"
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
