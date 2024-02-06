package day6

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/kevslinger/advent-of-code-2023/runner"
)

func RunDay6(path string) {
	raceProduct, err := runner.RunPart(path, part1)
	if err != nil {
		fmt.Printf("Error processing Day 6 Part 1: %s\n", err)
	} else {
		fmt.Printf("The answer to Day 6 Part 1 is: %d\n", raceProduct)
	}

	numWins, err := runner.RunPart(path, part2)
	if err != nil {
		fmt.Printf("Error processing Day 6 Part 2: %s\n", err)
	} else {
		fmt.Printf("The answer to Day 6 Part 2 is: %d\n", numWins)
	}
}

func part1(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)

	matchNumbers := regexp.MustCompile("[0-9]+")
	scanner.Scan()
	times, err := convertStringToInt(matchNumbers.FindAllString(scanner.Text(), -1))
	if err != nil {
		return -1, err
	}
	scanner.Scan()
	distances, err := convertStringToInt(matchNumbers.FindAllString(scanner.Text(), -1))
	if err != nil {
		return -1, err
	}
	races := make([]race, len(distances))
	raceProduct := 1
	for idx := range races {
		races[idx] = race{times[idx], distances[idx]}
		numWins := 0
		for i := 1; i < races[idx].Time; i++ {
			if i*(races[idx].Time-i) > races[idx].Distance {
				numWins++
			}
		}
		raceProduct *= numWins
	}

	return raceProduct, nil
}

func part2(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)

	matchNumbers := regexp.MustCompile("[0-9 ]+")
	scanner.Scan()
	time, err := strconv.Atoi(strings.Join(strings.Fields(matchNumbers.FindString(scanner.Text())), ""))
	if err != nil {
		return -1, err
	}
	scanner.Scan()
	distance, err := strconv.Atoi(strings.Join(strings.Fields(matchNumbers.FindString(scanner.Text())), ""))
	if err != nil {
		return -1, err
	}
	race := race{time, distance}
	numWins := 0
	for i := 1; i < race.Time; i++ {
		if i*(race.Time-i) > race.Distance {
			numWins++
		}
	}

	return numWins, nil
}

func convertStringToInt(strs []string) ([]int, error) {
	nums := make([]int, len(strs))
	for idx, str := range strs {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		nums[idx] = num
	}
	return nums, nil
}

type race struct {
	Time     int
	Distance int
}
