package board

import (
	"strings"
	"testing"

	"github.com/talglobus/fearsome/test"
)

// TestState_String asserts that State implements stringer, and that rows separated by \n are correct in number.
// There are any number of other clever checks that could be done on States, but given the intended flexibility of the
// string format for a State, the inherent complexity in testing that allows for that flexibility, combined with the
// vanishing probability of introducing bugs into Board.String that are not immediately obvious upon the use of the
// method — as is done frequently when developing this module — this testing is deemed to be premature optimization.
//
// Future versions of this module may yet add more clever checks, but for now we'll stick with this.
func TestState_String(t *testing.T) {
	s := Board{}.state
	if !test.IsStringer(s) {
		t.Errorf("State must implement `Stringer` interface")
		return
	}

	rows := strings.Split(s.String(), "\n")
	if len(rows) != ROWS {
		t.Errorf("State has incorrect number of rows. Expected %v, observed %v", ROWS, len(rows))
	}
}
