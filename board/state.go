package board

// State holds the full state of the board at a given point in time
type State [COLS][ROWS]Type

// String() provides fancy output to show the state in an easily readable format
func (s State) String() string {
	str := ""
	for rowNum := ROWS - 1; rowNum >= 0; rowNum-- {
		str += "|"
		for colNum := 0; colNum < COLS; colNum++ {
			str += " " + s[colNum][rowNum].String() + " |"
		}
		str += "\n"
	}

	return str[:len(str)-1]
}
