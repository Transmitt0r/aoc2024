package aoc

import (
	"fmt"
	"io"
	"io/fs"

	"github.com/transmitt0r/aoc2024/calendar"
)

const days = 25

type Solution interface {
	Load(io.Reader) error
	Print(io.Writer)
}

func Solve[S Solution](s S, r io.Reader, w io.Writer) {
	s.Load(r)
	s.Print(w)
}

type Day struct {
	Input    io.ReadCloser
	Solution Solution
}

func (d *Day) Solve(w io.Writer) {
	if d.Solution != nil && d.Input != nil {
		Solve(d.Solution, d.Input, w)
	}
}

type Calendar struct {
	days []Day
}

func (c *Calendar) GetDay(day int) (*Day, error) {
	if day <= 0 || day > days {
		return nil, fmt.Errorf("unabile to fetch solution for day %d", day)
	}
	return &c.days[day-1], nil
}

func NewCalendar(inputs fs.FS) (*Calendar, error) {
	c := Calendar{
		days: make([]Day, days),
	}

	for i := 1; i <= days; i++ {
		filename := fmt.Sprintf("day%02d.txt", i)
		file, err := inputs.Open(filename)
		if err != nil {
			// If the file doesn't exist, continue without breaking the calendar setup.
			if _, ok := err.(*fs.PathError); ok {
				continue
			}
			return nil, fmt.Errorf("failed to open file %s: %w", filename, err)
		}

		// Save an io.Reader for the file
		c.days[i-1] = Day{
			Input: file,
		}
	}

	c.days[0].Solution = calendar.NewDay01()
	c.days[1].Solution = calendar.NewDay02()
	c.days[2].Solution = calendar.NewDay03()
	c.days[3].Solution = calendar.NewDay04()
	return &c, nil
}

func (c *Calendar) Close() {
	for _, d := range c.days {
		if d.Input != nil {
			d.Input.Close()
		}
	}
}
