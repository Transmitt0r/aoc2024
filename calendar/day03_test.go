package calendar_test

import (
	"bytes"
	"testing"

	"github.com/transmitt0r/aoc2024/calendar"
)

var testInputDay03 = []byte(`xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`)
var secondTestInputDay03 = []byte(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`)

func TestDay03(t *testing.T) {
	sol := calendar.NewDay03()
	err := sol.Load(bytes.NewBuffer(testInputDay03))
	if err != nil {
		t.Fatal("error while loading input", err)
	}

	if result := sol.MultiplicationResult(false); result != 161 {
		t.Errorf("wrong result, wanted %d, got %d", 161, result)
	}

	err = sol.Load(bytes.NewBuffer(secondTestInputDay03))
	if err != nil {
		t.Fatal("error while loading input", err)
	}

	if result := sol.MultiplicationResult(true); result != 48 {
		t.Errorf("wrong result, wanted %d, got %d", 48, result)
	}
}
