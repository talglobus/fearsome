package util

import (
	"testing"
	"time"
)

func TestTraverse(t *testing.T) {
	tables := []struct {
		from int
		to int
		steps []int
	}{
		{1, 7, []int{1, 2, 3, 4, 5, 6, 7}},
		{10, 3, []int{10, 9, 8, 7, 6, 5, 4, 3}},
		{-7, -7, []int{-7}},
		{330002, 330002, []int{330002}},
		{11, -4, []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4}},
		{-5, 2, []int{-5, -4, -3, -2, -1, 0, 1, 2}},
	}

	TraverseCases:							// Label to jump to next test case
	for _, table := range tables {
		c := Traverse(table.from, table.to)
		var expecteds, observeds []int
		for _, expected := range table.steps {
			select {
				case observed, ok := <-c:
					expecteds = append(expecteds, expected)

					if !ok {
						t.Errorf("Traverse from %v to %v ended prematurely. Expected %v..., observed %v...",
							table.from, table.to, expecteds, observeds)
						continue TraverseCases
					}

					observeds = append(observeds, observed)

					if expected != observed {
						t.Errorf("Traverse from %v to %v was incorrect. Expected %v..., observed %v...",
							table.from, table.to, expecteds, observeds)
						continue TraverseCases
					}
				case <-time.After(100 * time.Millisecond):
					t.Errorf("Traverse from %v to %v hit 100ms timeout. Expected %v..., observed %v...",
						table.from, table.to, expecteds, observeds)
					continue TraverseCases
			}
		}

		// See if there are any more elements on the channel
		select {
			case observed, ok := <-c:
				// Check if the output channel is exhausted when the "expected" list is exhausted
				if ok {
					observeds = append(observeds, observed)
					t.Errorf("Traverse from %v to %v was incorrect. Expected %v..., observed %v...",
						table.from, table.to, expecteds, observeds)
					continue TraverseCases
				}

			// Check if the output channel was left open
			case <-time.After(100 * time.Millisecond):
				t.Errorf("Traverse from %v to %v failed to close output channel after %v elems, expected %v",
					table.from, table.to, len(observeds), len(table.steps))
				continue TraverseCases
			}
	}
}