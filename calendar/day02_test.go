package calendar_test

import (
	"bytes"
	"testing"

	"github.com/transmitt0r/aoc2024/calendar"
)

var testInputDay02 = []byte(`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`)

func TestDay02(t *testing.T) {
	sol := calendar.NewDay02()
	err := sol.Load(bytes.NewBuffer(testInputDay02))
	if err != nil {
		t.Fatal("error while loading input", err)
	}

	if reports := sol.SafeReports(false); reports != 2 {
		t.Errorf("wrong number of reports, wanted %d, got %d", 2, reports)
	}

	if reports := sol.SafeReports(true); reports != 4 {
		t.Errorf("wrong number of reports, wanted %d, got %d", 4, reports)
	}
}
