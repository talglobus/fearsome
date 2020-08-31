package board

import "sync"

// COLS and ROWS define the size of the game board, and are largely manipulable for different game dynamics
const (
	COLS = 7
	ROWS = 6
)

// Board holds the current state of the board, the history that got it there, and an rwmutex for thread safety.
// Note that rows are filled from bottom to top, indices 0 to ROWS-1, respectively
type Board struct {
	history History
	state   State
	mutex   *sync.RWMutex
}

// New constructs a new Board
func New() Board {
	return Board{
		History{},
		State{},
		&sync.RWMutex{},
	}
}

// RLock exposes the mutex's RLock functionality without exposing the mutex itself
func (b Board) RLock() {
	b.mutex.RLock()
}

// Lock exposes the mutex's Lock functionality without exposing the mutex itself
func (b Board) Lock() {
	b.mutex.Lock()
}

// RUnlock exposes the mutex's RUnlock functionality without exposing the mutex itself
func (b Board) RUnlock() {
	b.mutex.RUnlock()
}

// Unlock exposes the mutex's Unlock functionality without exposing the mutex itself
func (b Board) Unlock() {
	b.mutex.Unlock()
}

func (b Board) String() string {
	b.RLock()
	defer b.RUnlock()
	return b.state.String() + "\n" + b.history.String()
}

// Equals tests equality between two boards
func (b Board) Equals(b2 Board) bool {
	// Note that in a naively-instantiated board, the mutexes could be nil pointers, causing a panic
	// Given that this is always programmer error, however, this panic is intentionally left unguarded
	b.RLock()
	b2.RLock()
	defer b.RUnlock()
	defer b2.RUnlock()
	return b.state == b2.state && b.history.Equals(b2.history)
}

// nextTurn determines type of next move based on the board's history, even-indexed moves being red, odd-indexed blue
func (b Board) nextTurn() Type {
	b.RLock()
	defer b.RUnlock()
	var t Type
	if turnParity := len(b.history) % 2; turnParity == 0 {
		t = RED
	} else {
		t = BLUE
	}
	return t
}

// Move allows for by-column, "drop-style" move making, as in real-life Connect Four
// Note: Returns the resulting row coordinate, and an error if attempting to make a move in an already-full column
func (b *Board) Move(colNum int) (Type, int, error) {
	t := b.nextTurn()

	b.Lock()
	defer b.Unlock()

	// Iterate over chosen column from bottom to top, filling the first empty square, or returning error if none found
	for rowNum := range b.state[colNum] {
		if b.state[colNum][rowNum] == NONE {
			b.state[colNum][rowNum] = t
			b.history = append(b.history, Move(colNum))
			return t, rowNum, nil
		}
	}

	return NONE, 0, FullColumnError(colNum)
}

// MoveRed wraps Move for a slightly safer, `red`-specific way of making moves
// Note that there are two failure modes, namely an attempt to move out of turn, and failure of the move itself.
// In case of both, the error caused by the out-of-turn move takes precedence, thus only that error will be shown.
func (b *Board) MoveRed(colNum int) (int, error) {
	if b.nextTurn() != RED {
		return 0, TurnValidityError(RED)
	}

	_, r, e := b.Move(colNum)

	return r, e
}

// MoveBlue wraps Move for a slightly safer, `blue`-specific way of making moves.
// Note that there are two failure modes, namely an attempt to move out of turn, and failure of the move itself.
// In case of both, the error caused by the out-of-turn move takes precedence, thus only that error will be shown.
func (b *Board) MoveBlue(colNum int) (int, error) {
	if b.nextTurn() != BLUE {
		return 0, TurnValidityError(BLUE)
	}

	_, r, e := b.Move(colNum)

	return r, e
}
