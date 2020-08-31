/*
Package board implements functionality around the game Connect Four, centered around the eponymous "Board" type.

A Board consists of three parts, namely a history, a current state resulting from that history, and an rwmutex.
From a design perspective, the benefit of this dual-truth storage mechanism is that the history provides complete
knowledge of the board for the purpose of analysis — no further knowledge exists within the context of the board —
while the state avoids the need for reconstructing the current board state for any and all analysis to occur.

Because of the possibility for the state and the history to fall out of sync if not carefully managed, all moves
are made through constrained setters which update both in sequence, and to maintain thread-safety, an rwmutex
ensures that it can be read by multiple readers simultaneously across goroutines, or written by one writer in one
goroutine. The Board type, as well as all its methods, IS thread safe when passed by reference, so tinker away :)

Boards are held to various standards of validity in this package, defined as follows:

• Drop Validity - Guarantees that no occupied square sits above an empty square

• Turn Validity - Strict superset of Drop Validity. Guarantees alternating red/blue turns, starting with red

• Termination Validity - Guarantees termination of the game upon the win of either player

• History Validity - Guarantees that the sequence of moves in Board.history results in the state in Board.state

Drop Validity and Termination Validity are easy to prove from state alone. History Validity is easy to prove given
a board and history simply by rewalking the history to see if it results in the board. Turn Validity is implicit
in the history, but is not tractable to prove without the history. Therefore, any simulation generating boards can
easily guarantee all four validity types, while any process starting with a pre-generated board for which the
specific sequence of moves that originated it are unknown can only readily be guaranteed Drop and Turn Validity.
*/
package board
