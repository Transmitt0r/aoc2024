package calendar_test

import (
	"bytes"
	"testing"

	"github.com/transmitt0r/aoc2024/calendar"
)

var testInputDay01 = []byte(`3   4
4   3
2   5
1   3
3   9
3   3`)

func TestDay01(t *testing.T) {
	sol := calendar.NewDay01()
	err := sol.Load(bytes.NewBuffer(testInputDay01))
	if err != nil {
		t.Fatal("error while loading input", err)
	}

	if distance := sol.TotalDistance(); distance != 11 {
		t.Errorf("wrong distance, wanted %d, got %d", 11, distance)
	}

	if similarity := sol.SimilarityScore(); similarity != 31 {
		t.Errorf("wrong similarity, wanted %d, got %d", 31, similarity)
	}
}
