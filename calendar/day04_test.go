package calendar_test

import (
	"bytes"
	"testing"

	"github.com/transmitt0r/aoc2024/calendar"
)

var testInputDay04 = []byte(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`)

func TestDay04(t *testing.T) {
	sol := calendar.NewDay04()
	err := sol.Load(bytes.NewBuffer(testInputDay04))
	if err != nil {
		t.Fatal("error while loading input", err)
	}

	if result := sol.Search(calendar.CheckForXMAS); result != 18 {
		t.Errorf("wrong result, wanted %d, got %d", 18, result)
	}

	if result := sol.Search(calendar.CheckForCrossMAS); result != 9 {
		t.Errorf("wrong result, wanted %d, got %d", 9, result)
	}
}
