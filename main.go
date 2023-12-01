package main

import (
	"fmt"

	"github.com/kevslinger/advent-of-code-2023/day1"
)

func main() {
	day1InputPath := "./day1/data/day1.txt"
	sum, err := day1.Part1(day1InputPath)
	if err != nil {
		fmt.Printf("Error in Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Part 1 is %d\n", sum)
	}

	sum, err = day1.Part2(day1InputPath)
	if err != nil {
		fmt.Printf("Error in Part 2: %s\n", err)
	} else {
		fmt.Printf("Answer to Part 2 is %d\n", sum)
	}
}
