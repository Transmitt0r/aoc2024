package calendar

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

type Day01 struct {
	lhs []int
	rhs []int
}

func NewDay01() *Day01 {
	return &Day01{
		lhs: []int{},
		rhs: []int{},
	}
}

func (sol *Day01) Load(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	i := 0
	for scanner.Scan() {
		i++
		line := scanner.Text()
		fields := strings.Fields(line)

		lhs, err := strconv.Atoi(fields[0])
		if err != nil {
			return fmt.Errorf("error parsing lhs at line %v", i)
		}

		rhs, err := strconv.Atoi(fields[1])
		if err != nil {
			return fmt.Errorf("error parsing rhs at line %v", i)
		}

		sol.lhs = append(sol.lhs, lhs)
		sol.rhs = append(sol.rhs, rhs)
	}

	return nil
}

func (d *Day01) TotalDistance() int {
	lhs := slices.Clone(d.lhs)
	rhs := slices.Clone(d.rhs)
	slices.Sort(lhs)
	slices.Sort(rhs)

	distanceSum := 0

	for i := 0; i < len(lhs); i++ {
		dif := lhs[i] - rhs[i]
		if dif < 0 {
			dif *= -1
		}
		distanceSum += dif
	}

	return distanceSum
}

func (d *Day01) SimilarityScore() int {
	rhsCount := map[int]int{}

	similarity := 0

	for _, num := range d.rhs {
		if _, ok := rhsCount[num]; !ok {
			rhsCount[num] = 0
		}
		rhsCount[num]++
	}

	for _, num := range d.lhs {
		similarity += num * rhsCount[num]
	}

	return similarity
}

func (d *Day01) Print(w io.Writer) {
	fmt.Fprintf(w, "Total Distance: %d\n", d.TotalDistance())
	fmt.Fprintf(w, "Similarity: %d\n", d.SimilarityScore())
}
