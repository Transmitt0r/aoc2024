package calendar

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

type Report []int

func (r Report) IsSafe() bool {
	isIncreasing := r[0] < r[len(r)-1]

	for i := 0; i < len(r)-1; i++ {
		lhs := r[i]
		rhs := r[i+1]

		if lhs == rhs {
			return false
		}
		if isIncreasing && lhs >= rhs {
			return false
		}
		if !isIncreasing && lhs <= rhs {
			return false
		}

		diff := lhs - rhs
		if diff < 0 {
			diff *= -1
		}

		if diff > 3 {
			return false
		}
	}

	return true
}

func (r Report) IsSafeWithDampener() bool {
	isIncreasing := r[0] < r[len(r)-1]

	for i := 0; i < len(r)-1; i++ {
		lhs := r[i]
		rhs := r[i+1]

		diff := lhs - rhs
		if diff < 0 {
			diff *= -1
		}

		if lhs == rhs || isIncreasing && lhs >= rhs || !isIncreasing && lhs <= rhs || diff > 3 {
			clonedRhs := slices.Clone(r)
			clonedRhs = append(clonedRhs[:i], clonedRhs[i+1:]...)
			clonedLhs := slices.Clone(r)
			clonedLhs = append(clonedLhs[:i+1], clonedLhs[i+2:]...)
			return clonedLhs.IsSafe() || clonedRhs.IsSafe()
		}
	}

	return true
}

type Day02 struct {
	reports []Report
}

func NewDay02() *Day02 {
	return &Day02{
		reports: []Report{},
	}
}

func (sol *Day02) Load(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	i := 0
	for scanner.Scan() {
		report := []int{}
		i++
		line := scanner.Text()
		for j, field := range strings.Fields(line) {
			num, err := strconv.Atoi(field)
			if err != nil {
				return fmt.Errorf("error parsing report %d field %d", i, j)
			}
			report = append(report, num)
		}

		sol.reports = append(sol.reports, report)
	}

	return nil
}

func (d *Day02) SafeReports(dampenerEnabled bool) int {
	numSafeReports := 0
	for _, report := range d.reports {
		if !dampenerEnabled && report.IsSafe() {
			numSafeReports++
		}

		if dampenerEnabled && report.IsSafeWithDampener() {
			numSafeReports++
		}
	}
	return numSafeReports
}

func (d *Day02) Print(w io.Writer) {
	fmt.Fprintf(w, "Safe Reports: %d\n", d.SafeReports(false))
	fmt.Fprintf(w, "Safe Reports (with dampener): %d\n", d.SafeReports(true))
}
