package board

import (
	"errors"
	"sync"
	"testing"
)

func TestBoard_Move(t *testing.T) {
	table := []struct {
		name          string
		move          int
		before, after Board
		typ           Type
		row           int
		err           error
	}{
		{"red move",
			5,
			Board{
				history: History{4, 5},
				state:   State{{}, {}, {}, {}, {RED}, {BLUE}, {}},
				mutex:   &sync.RWMutex{},
			}, Board{
				history: History{4, 5, 5},
				state:   State{{}, {}, {}, {}, {RED}, {BLUE, RED}, {}},
				mutex:   &sync.RWMutex{},
			}, RED, 1, nil,
		}, {"blue move",
			5,
			Board{
				history: History{4},
				state:   State{{}, {}, {}, {}, {RED}, {}, {}},
				mutex:   &sync.RWMutex{},
			}, Board{
				history: History{4, 5},
				state:   State{{}, {}, {}, {}, {RED}, {BLUE}, {}},
				mutex:   &sync.RWMutex{},
			}, BLUE, 0, nil,
		}, {"red move on empty board",
			6,
			Board{
				history: History{},
				state:   State{{}, {}, {}, {}, {}, {}, {}},
				mutex:   &sync.RWMutex{},
			}, Board{
				history: History{6},
				state:   State{{}, {}, {}, {}, {}, {}, {RED}},
				mutex:   &sync.RWMutex{},
			}, RED, 0, nil,
		}, {"blue move on empty column",
			5,
			Board{
				history: History{0},
				state:   State{{RED}, {}, {}, {}, {}, {}, {}},
				mutex:   &sync.RWMutex{},
			}, Board{
				history: History{0, 5},
				state:   State{{RED}, {}, {}, {}, {}, {BLUE}, {}},
				mutex:   &sync.RWMutex{},
			}, BLUE, 0, nil,
		}, {"red move on full column",
			5,
			Board{
				history: History{4, 5, 5, 4, 4, 5, 5, 4, 4, 5, 5, 4},
				state:   State{{}, {}, {}, {}, {RED, BLUE, RED, BLUE, RED, BLUE}, {BLUE, RED, BLUE, RED, BLUE, RED}},
				mutex:   &sync.RWMutex{},
			}, New(), NONE, 0, FullColumnError(5),
		}, {"blue move on full column",
			0,
			Board{
				history: History{0, 1, 0, 1, 0, 1, 1, 3, 1, 0, 1, 0, 0},
				state:   State{{RED, RED, RED, BLUE, BLUE, RED}, {BLUE, BLUE, BLUE, RED, RED, RED}, {}, {BLUE}},
				mutex:   &sync.RWMutex{},
			}, New(), NONE, 0, FullColumnError(0),
		},
	}

	for _, elem := range table {
		t.Run(elem.name, func(t *testing.T) {
			// Make move, hopefully turning "before" board into expected "after" board
			typ, row, err := elem.before.Move(elem.move)

			// Board is undefined by spec on returning an error, and so doesn't need to be checked in that case
			if !elem.before.Equals(elem.after) && err == nil {
				t.Errorf("%v into column %v produced unexpected board state. Expected:\n%v\nObserved:\n%v",
					elem.name, elem.move, elem.after, elem.before)
			}

			if typ != elem.typ {
				t.Errorf("%v into column %v produced unexpected move type. Expected: %v, Observed: %v",
					elem.name, elem.move, elem.typ, typ)
			}

			if row != elem.row {
				t.Errorf("%v into column %v produced unexpected move row. Expected: %v, Observed: %v",
					elem.name, elem.move, elem.row, row)
			}

			if !errors.Is(err, elem.err) {
				t.Errorf("%v into column %v produced unexpected error. Expected:\n\t%v\nObserved:\n\t%v",
					elem.name, elem.move, elem.err, err)
			}
		})
	}
}

func TestBoard_MoveRed(t *testing.T) {
	// Note that the restricted testcases used here assume MoveRed wraps Move, so Move tests cover move functionality

	table := []struct {
		name          string
		move          int
		before, after Board
		row           int
		err           error
	}{
		{"red move",
			5,
			Board{
				history: History{4, 5},
				state:   State{{}, {}, {}, {}, {RED}, {BLUE}, {}},
				mutex:   &sync.RWMutex{},
			}, Board{
				history: History{4, 5, 5},
				state:   State{{}, {}, {}, {}, {RED}, {BLUE, RED}, {}},
				mutex:   &sync.RWMutex{},
			}, 1, nil,
		}, {"blue move",
			5,
			Board{
				history: History{4},
				state:   State{{}, {}, {}, {}, {RED}, {}, {}},
				mutex:   &sync.RWMutex{},
			}, New(), 0, TurnValidityError(RED),
		}, {"red move on full column",
			5,
			Board{
				history: History{4, 5, 5, 4, 4, 5, 5, 4, 4, 5, 5, 4},
				state:   State{{}, {}, {}, {}, {RED, BLUE, RED, BLUE, RED, BLUE}, {BLUE, RED, BLUE, RED, BLUE, RED}},
				mutex:   &sync.RWMutex{},
			}, New(), 0, FullColumnError(5),
		}, {"blue move on full column",
			0,
			Board{
				history: History{0, 1, 0, 1, 0, 1, 1, 3, 1, 0, 1, 0, 0},
				state:   State{{RED, RED, RED, BLUE, BLUE, RED}, {BLUE, BLUE, BLUE, RED, RED, RED}, {}, {BLUE}},
				mutex:   &sync.RWMutex{},
			}, New(), 0, TurnValidityError(RED),
		},
	}

	for _, elem := range table {
		t.Run(elem.name, func(t *testing.T) {
			// Make move, hopefully turning "before" board into expected "after" board
			row, err := elem.before.MoveRed(elem.move)

			// Board is undefined by spec on returning an error, and so doesn't need to be checked in that case
			if !elem.before.Equals(elem.after) && err == nil {
				t.Errorf("%v into column %v produced unexpected board state. Expected:\n%v\nObserved:\n%v",
					elem.name, elem.move, elem.after, elem.before)
			}

			if row != elem.row {
				t.Errorf("%v into column %v produced unexpected move row. Expected: %v, Observed: %v",
					elem.name, elem.move, elem.row, row)
			}

			if !errors.Is(err, elem.err) {
				t.Errorf("%v into column %v produced unexpected error. Expected:\n\t%v\nObserved:\n\t%v",
					elem.name, elem.move, elem.err, err)
			}
		})
	}
}

func TestBoard_MoveBlue(t *testing.T) {
	// Note that the restricted testcases used here assume MoveRed wraps Move, so Move tests cover move functionality

	table := []struct {
		name          string
		move          int
		before, after Board
		row           int
		err           error
	}{
		{"red move",
			5,
			Board{
				history: History{4, 5},
				state:   State{{}, {}, {}, {}, {RED}, {BLUE}, {}},
				mutex:   &sync.RWMutex{},
			}, New(), 0, TurnValidityError(BLUE),
		}, {"blue move",
			5,
			Board{
				history: History{4},
				state:   State{{}, {}, {}, {}, {RED}, {}, {}},
				mutex:   &sync.RWMutex{},
			}, Board{
				history: History{4, 5},
				state:   State{{}, {}, {}, {}, {RED}, {BLUE}, {}},
				mutex:   &sync.RWMutex{},
			}, 0, nil,
		}, {"red move on full column",
			5,
			Board{
				history: History{4, 5, 5, 4, 4, 5, 5, 4, 4, 5, 5, 4},
				state:   State{{}, {}, {}, {}, {RED, BLUE, RED, BLUE, RED, BLUE}, {BLUE, RED, BLUE, RED, BLUE, RED}},
				mutex:   &sync.RWMutex{},
			}, New(), 0, TurnValidityError(BLUE),
		}, {"blue move on full column",
			0,
			Board{
				history: History{0, 1, 0, 1, 0, 1, 1, 3, 1, 0, 1, 0, 0},
				state:   State{{RED, RED, RED, BLUE, BLUE, RED}, {BLUE, BLUE, BLUE, RED, RED, RED}, {}, {BLUE}},
				mutex:   &sync.RWMutex{},
			}, New(), 0, FullColumnError(0),
		},
	}

	for _, elem := range table {
		t.Run(elem.name, func(t *testing.T) {
			// Make move, hopefully turning "before" board into expected "after" board
			row, err := elem.before.MoveBlue(elem.move)

			// Board is undefined by spec on returning an error, and so doesn't need to be checked in that case
			if !elem.before.Equals(elem.after) && err == nil {
				t.Errorf("%v into column %v produced unexpected board state. Expected:\n%v\nObserved:\n%v",
					elem.name, elem.move, elem.after, elem.before)
			}

			if row != elem.row {
				t.Errorf("%v into column %v produced unexpected move row. Expected: %v, Observed: %v",
					elem.name, elem.move, elem.row, row)
			}

			if !errors.Is(err, elem.err) {
				t.Errorf("%v into column %v produced unexpected error. Expected:\n\t%v\nObserved:\n\t%v",
					elem.name, elem.move, elem.err, err)
			}
		})
	}
}

// Note that Unmove should leave board unchanged on error
func TestBoard_Unmove(t *testing.T) {
	table := []struct {
		name          string
		before, after Board
		err           error
	}{
		{"unmove empty board",
			New(),
			New(),
			EmptyBoardError{},
		}, {"unmove empty column",
			Board{
				history: History{0, 1, 0, 1, 0, 1, 1, 3, 1, 0, 1, 0, 2},
				state:   State{{RED, RED, RED, BLUE, BLUE}, {BLUE, BLUE, BLUE, RED, RED, RED}, {}, {BLUE}},
				mutex:   &sync.RWMutex{},
			},
			Board{
				history: History{0, 1, 0, 1, 0, 1, 1, 3, 1, 0, 1, 0, 2},
				state:   State{{RED, RED, RED, BLUE, BLUE}, {BLUE, BLUE, BLUE, RED, RED, RED}, {}, {BLUE}},
				mutex:   &sync.RWMutex{},
			},
			HistoryValidityError(Board{
				history: History{0, 1, 0, 1, 0, 1, 1, 3, 1, 0, 1, 0, 2},
				state:   State{{RED, RED, RED, BLUE, BLUE}, {BLUE, BLUE, BLUE, RED, RED, RED}, {}, {BLUE}},
				mutex:   &sync.RWMutex{},
			}),
		}, {"unmove board with substituted top element (red)",
			Board{
				history: History{4, 5, 5},
				state:   State{{}, {}, {}, {}, {RED}, {BLUE, BLUE}, {}},
				mutex:   &sync.RWMutex{},
			},
			Board{
				history: History{4, 5, 5},
				state:   State{{}, {}, {}, {}, {RED}, {BLUE, BLUE}, {}},
				mutex:   &sync.RWMutex{},
			},
			HistoryValidityError(Board{
				history: History{4, 5, 5},
				state:   State{{}, {}, {}, {}, {RED}, {BLUE, BLUE}, {}},
				mutex:   &sync.RWMutex{},
			}),
		}, {"unmove board with substituted top element (blue)",
			Board{
				history: History{4, 5, 5, 4, 4, 5, 5, 4, 4, 5, 5, 4},
				state:   State{{}, {}, {}, {}, {RED, BLUE, RED, BLUE, RED, RED}, {BLUE, RED, BLUE, RED, BLUE, RED}},
				mutex:   &sync.RWMutex{},
			},
			Board{
				history: History{4, 5, 5, 4, 4, 5, 5, 4, 4, 5, 5, 4},
				state:   State{{}, {}, {}, {}, {RED, BLUE, RED, BLUE, RED, RED}, {BLUE, RED, BLUE, RED, BLUE, RED}},
				mutex:   &sync.RWMutex{},
			},
			HistoryValidityError(Board{
				history: History{4, 5, 5, 4, 4, 5, 5, 4, 4, 5, 5, 4},
				state:   State{{}, {}, {}, {}, {RED, BLUE, RED, BLUE, RED, RED}, {BLUE, RED, BLUE, RED, BLUE, RED}},
				mutex:   &sync.RWMutex{},
			}),
		}, {"unmove non-empty column (red)",
			Board{
				history: History{4, 5, 5},
				state:   State{{}, {}, {}, {}, {RED}, {BLUE, RED}, {}},
				mutex:   &sync.RWMutex{},
			},
			Board{
				history: History{4, 5},
				state:   State{{}, {}, {}, {}, {RED}, {BLUE}, {}},
				mutex:   &sync.RWMutex{},
			},
			nil,
		}, {"unmove non-empty column (red) to empty board",
			Board{
				history: History{4},
				state:   State{{}, {}, {}, {}, {RED}, {}, {}},
				mutex:   &sync.RWMutex{},
			},
			New(),
			nil,
		}, {"unmove non-empty column (blue)",
			Board{
				history: History{0, 1},
				state:   State{{RED}, {BLUE}, {}, {}, {}, {}, {}},
				mutex:   &sync.RWMutex{},
			},
			Board{
				history: History{0},
				state:   State{{RED}, {}, {}, {}, {}, {}, {}},
				mutex:   &sync.RWMutex{},
			},
			nil,
		},
	}

	for _, elem := range table {
		// Perform unmove, hopefully turning "before" board into expected "after" board
		err := elem.before.Unmove()

		if !elem.before.Equals(elem.after) {
			t.Errorf("%v produced unexpected board state. Expected:\n%v\nObserved:\n%v",
				elem.name, elem.after, elem.before)
		}

		if !errors.Is(err, elem.err) {
			t.Errorf("%v produced unexpected error. Expected:\n\t%v\nObserved:\n\t%v",
				elem.name, elem.err, err)
		}
	}
}
