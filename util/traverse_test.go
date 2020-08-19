package util

import (
	"testing"
	"time"
)

func TestTraverse(t *testing.T) {
	const TIMEOUT = 100 * time.Millisecond

	tables := []struct {
		from  int
		to    int
		steps []int
		name  string
	}{
		{1, 7, []int{1, 2, 3, 4, 5, 6, 7}, "increase between positive values"},
		{10, 3, []int{10, 9, 8, 7, 6, 5, 4, 3}, "decrease between positive values"},
		{-7, -7, []int{-7}, "single negative value"},
		{330002, 330002, []int{330002}, "single large positive value"},
		{11, -4, []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4},
			"decrease from positive to negative values"},
		{-5, 2, []int{-5, -4, -3, -2, -1, 0, 1, 2}, "increase from negative to positive values"},
	}

	for _, table := range tables {
		t.Run(table.name, func(t *testing.T) {
			c := Traverse(table.from, table.to)
			var wants, gots []int
			for _, want := range table.steps {
				select {
				case got, ok := <-c:
					wants = append(wants, want)

					if !ok {
						t.Fatalf("Traverse from %v to %v ended prematurely. Expected %v..., observed %v...",
							table.from, table.to, wants, gots)
					}

					gots = append(gots, got)

					if want != got {
						t.Fatalf("Traverse from %v to %v was incorrect. Expected %v..., observed %v...",
							table.from, table.to, wants, gots)
					}
				case <-time.After(TIMEOUT):
					t.Fatalf("Traverse from %v to %v hit %v timeout. Expected %v..., observed %v...",
						table.from, table.to, TIMEOUT, wants, gots)
				}
			}

			// See if there are any more elements on the channel
			select {
			case got, ok := <-c:
				// Check if the output channel is exhausted when the "want" list is exhausted
				if ok {
					gots = append(gots, got)
					t.Fatalf("Traverse from %v to %v was incorrect. Expected %v..., observed %v...",
						table.from, table.to, wants, gots)
				}

			// Check if the output channel was left open
			case <-time.After(TIMEOUT):
				t.Fatalf("Traverse from %v to %v failed to close output channel after %v elems, expected %v",
					table.from, table.to, len(gots), len(table.steps))
			}
		})
	}
}
