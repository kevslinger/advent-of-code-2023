package main

import (
	"fmt"

	"github.com/kevslinger/advent-of-code-2023/day1"
)

func main() {
	day1.RunDay1(GetInputDay(1))
}

func GetInputDay(day int) string {
	return fmt.Sprintf("./day%d/data/day%d.txt", day, day)
}
