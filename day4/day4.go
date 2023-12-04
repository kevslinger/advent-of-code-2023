package day4

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func RunDay4(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error reading file in Day 4: %s\n", err)
		return
	}
	points, err := Part1(file)
	if err != nil {
		fmt.Printf("Error processing Day 4 Part 1: %s\n", err)
	} else {
		fmt.Printf("The answer to Day 4 Part 1 is: %d\n", points)
	}

	file2, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error reading file in Day 4 Part 2: %s\n", err)
		return
	}
	scratchCards, err := Part2(file2)
	if err != nil {
		fmt.Printf("Error processing Day 4 Part 2: %s\n", err)
	} else {
		fmt.Printf("The answer to Day 4 Part 2 is: %d\n", scratchCards)
	}

}

func Part1(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)
	totalPoints := 0
	for scanner.Scan() {
		points := 0
		line := scanner.Text()

		gameStrings := ParseGameCard(line)

		winningNumbers, err := GetNumbersFromString(gameStrings[0])
		if err != nil {
			return -1, nil
		}
		winningMap := make(map[int]bool)
		for _, num := range winningNumbers {
			winningMap[num] = true
		}
		ourNumbers, err := GetNumbersFromString(gameStrings[1])
		if err != nil {
			return -1, nil
		}
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

func Part2(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)
	totalPoints := 0
	wonScratchcards := make(map[int]int)
	for scanner.Scan() {
		// Include the card we started with
		points := 0
		line := scanner.Text()

		gameId, err := GetGameId(line)
		wonScratchcards[gameId]++
		if err != nil {
			return -1, nil
		}
		gameStrings := ParseGameCard(line)

		winningNumbers, err := GetNumbersFromString(gameStrings[0])
		if err != nil {
			return -1, nil
		}
		winningMap := make(map[int]bool)
		for _, num := range winningNumbers {
			winningMap[num] = true
		}
		ourNumbers, err := GetNumbersFromString(gameStrings[1])
		if err != nil {
			return -1, nil
		}
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

func GetGameId(card string) (int, error) {
	gameIdMatcher := regexp.MustCompile("[0-9]+:")
	gameId := gameIdMatcher.FindAllString(card, -1)[0]
	return strconv.Atoi(gameId[:len(gameId)-1])
}

func ParseGameCard(card string) [2][]string {
	afterCard := strings.Split(card, ": ")[1]
	numberSplit := strings.Split(afterCard, "|")

	numberMatcher := regexp.MustCompile("[0-9]+")
	winningStrings := numberMatcher.FindAllString(numberSplit[0], -1)
	ourStrings := numberMatcher.FindAllString(numberSplit[1], -1)
	return [2][]string{winningStrings, ourStrings}
}

func GetNumbersFromString(numStr []string) ([]int, error) {
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
