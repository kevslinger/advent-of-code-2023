package day4

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/kevslinger/advent-of-code-2023/runner"
)

func RunDay4(path string) {
	points, err := runner.RunPart(path, part1)
	if err != nil {
		fmt.Printf("Error processing Day 4 Part 1: %s\n", err)
	} else {
		fmt.Printf("The answer to Day 4 Part 1 is: %d\n", points)
	}

	scratchCards, err := runner.RunPart(path, part2)
	if err != nil {
		fmt.Printf("Error processing Day 4 Part 2: %s\n", err)
	} else {
		fmt.Printf("The answer to Day 4 Part 2 is: %d\n", scratchCards)
	}
}

func part1(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)
	totalPoints := 0
	for scanner.Scan() {
		line := scanner.Text()

		gameStrings := parseGameCard(line)

		winningNumbers, err := getNumbersFromString(gameStrings[0])
		if err != nil {
			return -1, nil
		}
		winningMap := createWinningMap(winningNumbers)
		ourNumbers, err := getNumbersFromString(gameStrings[1])
		if err != nil {
			return -1, nil
		}
		points := 0
		for _, val := range ourNumbers {
			_, ok := winningMap[val]
			if ok {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
		totalPoints += points
	}
	return totalPoints, nil
}

func part2(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)
	totalPoints := 0
	wonScratchcards := make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()

		gameId, err := getGameId(line)
		wonScratchcards[gameId]++
		if err != nil {
			return -1, nil
		}
		gameStrings := parseGameCard(line)

		winningNumbers, err := getNumbersFromString(gameStrings[0])
		if err != nil {
			return -1, nil
		}
		winningMap := createWinningMap(winningNumbers)
		ourNumbers, err := getNumbersFromString(gameStrings[1])
		if err != nil {
			return -1, nil
		}
		points := 0
		for _, val := range ourNumbers {
			_, ok := winningMap[val]
			if ok {
				points++
			}
		}
		for i := 1; i <= points; i++ {
			wonScratchcards[gameId+i] += wonScratchcards[gameId]
		}
		totalPoints += wonScratchcards[gameId]
	}
	return totalPoints, nil
}

func getGameId(card string) (int, error) {
	gameIdMatcher := regexp.MustCompile("[0-9]+:")
	gameId := gameIdMatcher.FindAllString(card, -1)[0]
	return strconv.Atoi(gameId[:len(gameId)-1])
}

func parseGameCard(card string) [2][]string {
	afterCard := strings.Split(card, ": ")[1]
	numberSplit := strings.Split(afterCard, "|")

	numberMatcher := regexp.MustCompile("[0-9]+")
	winningStrings := numberMatcher.FindAllString(numberSplit[0], -1)
	ourStrings := numberMatcher.FindAllString(numberSplit[1], -1)
	return [2][]string{winningStrings, ourStrings}
}

func getNumbersFromString(numStr []string) ([]int, error) {
	nums := make([]int, len(numStr))
	for idx, str := range numStr {
		num, err := strconv.Atoi(str)
		if err != nil {
			return make([]int, 0), err
		}
		nums[idx] = num
	}
	return nums, nil
}

func createWinningMap(winningNumbers []int) map[int]bool {
	winningMap := make(map[int]bool)
	for _, num := range winningNumbers {
		winningMap[num] = true
	}
	return winningMap
}
