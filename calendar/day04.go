package calendar

import (
	"bufio"
	"fmt"
	"io"
)

type Day04 struct {
	input [][]rune
}

func NewDay04() *Day04 {
	return &Day04{
		input: [][]rune{},
	}
}

func (sol *Day04) Load(r io.Reader) error {
	sol.input = [][]rune{}

	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		sol.input = append(sol.input, []rune(line))
	}
	return nil
}

func (d *Day04) Search(searchFunc func([][]rune, int, int) int) int {
	cnt := 0
	for i := range d.input {
		for j := range d.input[i] {
			cnt += searchFunc(d.input, i, j)
		}
	}

	return cnt
}

func CheckForXMAS(input [][]rune, y, x int) int {
	const word = "XMAS"
	start := rune(word[0])

	if input[y][x] != start {
		return 0
	}

	offsets := []struct {
		x int
		y int
	}{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	count := 0

	for _, o := range offsets {
		for i := 0; i < len(word); i++ {
			newX, newY := x+(i*o.x), y+(i*o.y)
			if newX < 0 || newY < 0 || newX >= len(input[y]) || newY >= len(input) {
				break
			}
			if rune(word[i]) != input[newY][newX] {
				break
			}
			if i == len(word)-1 {
				count++
			}
		}
	}

	return count
}

func CheckForCrossMAS(input [][]rune, y, x int) int {
	wordOffsets := [][]struct {
		x int
		y int
	}{
		{
			{-1, -1},
			{0, 0},
			{1, 1},
		},
		{
			{-1, 1},
			{0, 0},
			{1, -1},
		},
	}

	const word = "MAS"
	const wordReverse = "SAM"

	if input[y][x] != 'A' {
		return 0
	}

	for _, variant := range wordOffsets {
		builtWord := []rune{}
		for i, o := range variant {
			newX, newY := x+o.x, y+o.y
			if newX < 0 || newY < 0 || newX >= len(input[y]) || newY >= len(input) {
				return 0
			}
			builtWord = append(builtWord, input[newY][newX])
			if i == len(word)-1 {
				builtWordStringified := string(builtWord)
				if builtWordStringified != word && builtWordStringified != wordReverse {
					return 0
				}
			}
		}
	}

	return 1
}

func (d *Day04) Print(w io.Writer) {
	fmt.Fprintf(w, "Search Result: XMAS occured %d times\n", d.Search(CheckForXMAS))
	fmt.Fprintf(w, "Search Result: X-MAS occured %d times\n", d.Search(CheckForCrossMAS))
}
