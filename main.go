package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/transmitt0r/aoc2024/aoc"
)

func main() {
	day := flag.Int("day", 1, "which solution to run")
	inputs := flag.String("inputs", "./inputs", "directory containing puzzle inputs")
	flag.Parse()

	c, err := aoc.NewCalendar(os.DirFS(*inputs))
	if err != nil {
		fmt.Printf("error while loading inputs: %v\n", err)
		os.Exit(1)
	}
	defer c.Close()

	sol, err := c.GetDay(*day)
	if err != nil {
		fmt.Printf("error while fetching solution: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Day %d\n", *day)
	sol.Solve(os.Stdout)
}
