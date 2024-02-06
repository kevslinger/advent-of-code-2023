package day2

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/kevslinger/advent-of-code-2023/runner"
)

func RunDay2(path string) {
	sum, err := runner.RunPart(path, part1)
	if err != nil {
		fmt.Printf("Error in Day 2 Part 1: %s\n", err)
	} else {
		fmt.Printf("The answer to Day 2 Part 1 is: %d\n", sum)
	}

	sum, err = runner.RunPart(path, part2)
	if err != nil {
		fmt.Printf("Error in Day 2 Part 2: %s\n", err)
	} else {
		fmt.Printf("The answer to Day2 Part 2 is: %d\n", sum)
	}
}

func part1(file io.Reader) (int, error) {
	var sum int
	gameIdMatcher := regexp.MustCompile("Game [0-9]+")
	var matchers []*regexp.Regexp = getMatchers()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		isCompliant := true
		for idx, matcher := range matchers {
			allBalls := matcher.FindAllString(line, -1)
			// Red is 12, Green is 13, Blue is 14
			isWithinCapacity, err := checkCapacity(allBalls, 12+idx)
			if err != nil {
				return -1, err
			}
			if !isWithinCapacity {
				isCompliant = false
				break
			}
		}
		if !isCompliant {
			continue
		}
		id, err := getGameId(string(gameIdMatcher.Find([]byte(line))))
		if err != nil {
			return -1, err
		}
		sum += id
	}
	return sum, nil
}

func getMatchers() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile("[0-9]+ red"),
		regexp.MustCompile("[0-9]+ green"),
		regexp.MustCompile("[0-9]+ blue"),
	}
}

func checkCapacity(balls []string, max int) (bool, error) {
	for _, ball := range balls {
		num, err := strconv.Atoi(strings.Split(ball, " ")[0])
		if err != nil {
			return false, err
		}
		if num > max {
			return false, nil
		}
	}
	return true, nil
}

func getGameId(line string) (int, error) {
	id, err := strconv.Atoi(strings.Split(line, " ")[1])
	if err != nil {
		return -1, err
	}
	return id, nil
}

func part2(file io.Reader) (int, error) {
	var sum int
	var matchers []*regexp.Regexp = getMatchers()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		minBallsArray := make([]int, 3)
		for idx, matcher := range matchers {
			allBalls := matcher.FindAllString(line, -1)
			// Red is 12, Green is 13, Blue is 14
			minBallsRequired, err := findMinBallsRequired(allBalls)
			if err != nil {
				return -1, err
			}
			minBallsArray[idx] = minBallsRequired
		}
		sum += minBallsArray[0] * minBallsArray[1] * minBallsArray[2]
	}
	return sum, nil
}

func findMinBallsRequired(balls []string) (int, error) {
	minBalls := 0
	for _, ball := range balls {
		num, err := strconv.Atoi(strings.Split(ball, " ")[0])
		if err != nil {
			return -1, err
		}
		if num > minBalls {
			minBalls = num
		}
	}
	return minBalls, nil
}
