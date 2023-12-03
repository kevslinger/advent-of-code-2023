package day2

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func RunDay2(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error reading file for Day 2: %s\n", err)
		return
	}
	defer file.Close()

	matchers := GetMatchers()
	sum, err := Part1(file, matchers)
	if err != nil {
		fmt.Printf("Error in Day 2 Part 1: %s\n", err)
	} else {
		fmt.Printf("The answer to Day 2 Part 1 is: %d\n", sum)
	}

	file2, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error reading file for Day 2: %s\n", err)
	}
	defer file2.Close()

	sum, err = Part2(file2, matchers)
	if err != nil {
		fmt.Printf("Error in Day 2 Part 2: %s\n", err)
	} else {
		fmt.Printf("The answer to Day2 Part 2 is: %d\n", sum)
	}
}

func GetMatchers() [3]*regexp.Regexp {
	return [3]*regexp.Regexp{
		regexp.MustCompile("[0-9]+ red"),
		regexp.MustCompile("[0-9]+ green"),
		regexp.MustCompile("[0-9]+ blue"),
	}
}

func Part1(file io.Reader, matchers [3]*regexp.Regexp) (int, error) {
	var sum int
	gameIdMatcher := regexp.MustCompile("Game [0-9]+")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		isCompliant := true
		for idx, matcher := range matchers {
			allBalls := matcher.FindAllString(line, -1)
			// Red is 12, Green is 13, Blue is 14
			isWithinCapacity, err := CheckCapacity(allBalls, 12+idx)
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
		id, err := GetGameId(string(gameIdMatcher.Find([]byte(line))))
		if err != nil {
			return -1, err
		}
		sum += id
	}
	return sum, nil
}

func CheckCapacity(balls []string, max int) (bool, error) {
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

func GetGameId(line string) (int, error) {
	id, err := strconv.Atoi(strings.Split(line, " ")[1])
	if err != nil {
		return -1, err
	}
	return id, nil
}

func Part2(file io.Reader, matchers [3]*regexp.Regexp) (int, error) {
	var sum int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		minBallsArray := make([]int, 3)
		for idx, matcher := range matchers {
			allBalls := matcher.FindAllString(line, -1)
			// Red is 12, Green is 13, Blue is 14
			minBallsRequired, err := FindMinBallsRequired(allBalls)
			if err != nil {
				return -1, err
			}
			minBallsArray[idx] = minBallsRequired
		}
		sum += minBallsArray[0] * minBallsArray[1] * minBallsArray[2]
	}
	return sum, nil
}

func FindMinBallsRequired(balls []string) (int, error) {
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
