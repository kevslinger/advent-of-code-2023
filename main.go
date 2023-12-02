package main

import (
	"fmt"

	"github.com/kevslinger/advent-of-code-2023/day1"
	"github.com/kevslinger/advent-of-code-2023/day2"
)

func main() {
	day1.RunDay1(GetInputDay(1))
	day2.RunDay2(GetInputDay(2))
}

func GetInputDay(day int) string {
	return fmt.Sprintf("./day%d/data/day%d.txt", day, day)
}
