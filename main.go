package main

import (
	"fmt"

	"github.com/kevslinger/advent-of-code-2023/day1"
	"github.com/kevslinger/advent-of-code-2023/day10"
	"github.com/kevslinger/advent-of-code-2023/day11"
	"github.com/kevslinger/advent-of-code-2023/day13"
	"github.com/kevslinger/advent-of-code-2023/day14"
	"github.com/kevslinger/advent-of-code-2023/day15"
	"github.com/kevslinger/advent-of-code-2023/day17"
	"github.com/kevslinger/advent-of-code-2023/day19"
	"github.com/kevslinger/advent-of-code-2023/day2"
	"github.com/kevslinger/advent-of-code-2023/day21"
	"github.com/kevslinger/advent-of-code-2023/day23"
	"github.com/kevslinger/advent-of-code-2023/day3"
	"github.com/kevslinger/advent-of-code-2023/day4"
	"github.com/kevslinger/advent-of-code-2023/day5"
	"github.com/kevslinger/advent-of-code-2023/day6"
	"github.com/kevslinger/advent-of-code-2023/day7"
	"github.com/kevslinger/advent-of-code-2023/day8"
	"github.com/kevslinger/advent-of-code-2023/day9"
)

func main() {
	day1.RunDay1(GetInputDay(1))
	day2.RunDay2(GetInputDay(2))
	day3.RunDay3(GetInputDay(3))
	day4.RunDay4(GetInputDay(4))
	day5.RunDay5(GetInputDay(5))
	day6.RunDay6(GetInputDay(6))
	day7.RunDay7(GetInputDay(7))
	day8.RunDay8(GetInputDay(8))
	day9.RunDay9(GetInputDay(9))
	day10.RunDay10(GetInputDay(10))
	day11.RunDay11(GetInputDay(11))

	day13.RunDay13(GetInputDay(13))
	day14.RunDay14(GetInputDay(14))
	day15.RunDay15(GetInputDay(15))

	day17.Runday17(GetInputDay(17))

	day19.RunDay19(GetInputDay(19))

	day21.RunDay21(GetInputDay(21))

	day23.RunDay23(GetInputDay(23))
}

func GetInputDay(day int) string {
	return fmt.Sprintf("./day%d/data/day%d.txt", day, day)
}
