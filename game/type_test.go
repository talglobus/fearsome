package game

import (
	"github.com/talglobus/fearsome/test"
	"math"
	"testing"
)

func TestType_String(t *testing.T) {
	r, b, n := RED, BLUE, NONE
	if !(test.IsStringer(r) && test.IsStringer(b) && test.IsStringer(n)) {
		t.Fatal("Type must implement `Stringer` interface")
	}

	if r.String() == b.String() || b.String() == n.String() || n.String() == r.String() {
		t.Fatalf("Type string output must be unique for `RED` (%q), `BLUE` (%q), and `NONE` (%q)", r, b, n)
	}
}

func TestType_newRandType(t *testing.T) {
	const (
		ROUNDS = 10000 // Number of test iterations to confirm random variation below tolerance
		// TOLERANCE sets a threshold, for variation in random output. Given 10,000 rounds, and a 5% tolerance,
		// the probability of a random output being falsely identified as non-random (false failure) is around 0.002%.
		// This can be calculated, rather tediously, as follows, given N is the number of trials and T is the chosen
		// tolerance: P_false_negative(N, T)
		// = sum from k=0 to floor(N/3 * (1-T)) of (nCr(n, k) * 1/3^floor(N/3 * (1-T)) * (2/3)^(N-k))
		TOLERANCE = 0.05
	)
	var red, blue, none int
	for i := 0; i < ROUNDS; i++ {
		switch newRandType() {
		case RED:
			red++
		case BLUE:
			blue++
		default:
			none++
		}
	}

	if minPass := int(math.Floor(ROUNDS / 3 * (1 - TOLERANCE))); red < minPass || blue < minPass || none < minPass {
		t.Fatal("newRandType must give random mix of `RED`s, `BLUE`s, and `NONE`s. Probabilistic, try test again")
	}
}
