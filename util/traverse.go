package util

// Traverse is an abstraction over numeric ranges, filling a channel with values from `from` to `to`, inclusive
func Traverse(from int, to int) <-chan int {
	// As Traverse should be faster than its calling function, buffering should guarantee it's not a bottleneck
	out := make(chan int, 1)
	if from < to {
		go func() {
			for i := from; i <= to; i++ {
				out <- i
			}
			close(out)
		}()
		return out
	} else if from > to {
		go func() {
			for i := from; i >= to; i-- {
				out <- i
			}
			close(out)
		}()
		return out
	} else {
		// If the `from` and `to` values are equal, just send that value, which won't be blocking because of the buffer
		out <- from
		close(out) // TODO: Make sure this operation isn't blocking
		return out
	}
}
