package main

import (
	"fmt"

	"github.com/kevslinger/advent-of-code-2023/day1"
	"github.com/kevslinger/advent-of-code-2023/day2"
	"github.com/kevslinger/advent-of-code-2023/day3"
	"github.com/kevslinger/advent-of-code-2023/day4"
	"github.com/kevslinger/advent-of-code-2023/day5"
)

func main() {
	day1.RunDay1(GetInputDay(1))
	day2.RunDay2(GetInputDay(2))
	day3.RunDay3(GetInputDay(3))
	day4.RunDay4(GetInputDay(4))
	day5.RunDay5(GetInputDay(5))
}

func GetInputDay(day int) string {
	return fmt.Sprintf("./day%d/data/day%d.txt", day, day)
}
