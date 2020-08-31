package board

import "fmt"

// FullColumnError defines an error used when attempting to add pieces to an already-full column
type FullColumnError int

func (e FullColumnError) Error() string {
	return fmt.Sprintf("column %v is full and cannot accept any more pieces", int(e))
}

// TurnValidityError defines an error used when attempting to make a move during the opposite player's turn
type TurnValidityError Type

func (e TurnValidityError) Error() string {
	return fmt.Sprintf("attempted to move %v out of turn", Type(e))
}

// HistoryValidityError defines an error used when History Validity is violated, i.e. when state doesn't match history
type HistoryValidityError Board

func (e HistoryValidityError) Error() string {
	return fmt.Sprintf("board history does not match board state\nState:\n%v\nHistory:\n%v",
		e.state, e.history)
}

// EmptyBoardError defines an error used when operations requiring a non-empty board are attempted on an empty board
type EmptyBoardError struct{}

func (e EmptyBoardError) Error() string {
	return "board is empty"
}
