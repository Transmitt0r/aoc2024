package calendar

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type Day03 struct {
	input string
}

func NewDay03() *Day03 {
	return &Day03{}
}

func (sol *Day03) Load(r io.Reader) error {
	body, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	sol.input = string(body)
	return nil
}

func (d *Day03) MultiplicationResult(conditional bool) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|don't\(\)|do\(\)`)

	matches := re.FindAllStringSubmatch(d.input, -1)
	mult := 0
	multIsEnabled := true
	for _, match := range matches {
		switch {
		case multIsEnabled && strings.HasPrefix(match[0], "mul"):
			lhs, err := strconv.Atoi(match[1])
			if err != nil {
				fmt.Printf("error while parsing %v", match)
				continue
			}
			rhs, err := strconv.Atoi(match[2])
			if err != nil {
				fmt.Printf("error while parsing %v", match)
				continue
			}
			mult += lhs * rhs
		case conditional && match[0] == "don't()":
			multIsEnabled = false
		case conditional && match[0] == "do()":
			multIsEnabled = true
		}
	}

	return mult
}

func (d *Day03) Print(w io.Writer) {
	fmt.Fprintf(w, "Multiplication Result: %d\n", d.MultiplicationResult(false))
	fmt.Fprintf(w, "Multiplication Result (Conditional): %d\n", d.MultiplicationResult(true))
}
