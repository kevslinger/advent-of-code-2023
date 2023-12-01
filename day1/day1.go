package day1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func RunDay1(path string) {
	sum, err := Part1(path)
	if err != nil {
		fmt.Printf("Error in Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Part 1 is %d\n", sum)
	}

	sum, err = Part2(path)
	if err != nil {
		fmt.Printf("Error in Part 2: %s\n", err)
	} else {
		fmt.Printf("Answer to Part 2 is %d\n", sum)
	}
}

func Part1(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	var sum int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		leftDigit, rightDigit := -1, -1

		for idx, c := range line {
			if unicode.IsDigit(c) && leftDigit == -1 {
				leftDigit = int(c - '0')
			}
			if unicode.IsDigit(rune(line[len(line)-idx-1])) && rightDigit == -1 {
				rightDigit = int(rune(line[len(line)-idx-1]) - '0')
			}
		}
		sum += 10*leftDigit + rightDigit
	}

	return sum, nil
}

func Part2(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	var sum int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		leftDigit, rightDigit := -1, -1

		for idx, c := range line {
			if leftDigit == -1 {
				if unicode.IsDigit(c) {
					leftDigit = int(c - '0')
				} else {
					// returns -1 if not found
					leftDigit = ContainsNumberWord(line[:idx+1])
				}
			}
			if rightDigit == -1 {
				if unicode.IsDigit(rune(line[len(line)-idx-1])) {
					rightDigit = int(rune(line[len(line)-idx-1]) - '0')
				} else {
					// returns -1 if not found
					rightDigit = ContainsNumberWord(line[len(line)-idx-1:])
				}
			}
		}
		sum += 10*leftDigit + rightDigit
	}

	return sum, nil
}

var numberWords = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func ContainsNumberWord(word string) int {
	for idx, number := range numberWords {
		if strings.Contains(word, number) {
			return idx + 1
		}
	}
	return -1
}
