package day1

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"

	"github.com/kevslinger/advent-of-code-2023/runner"
)

func RunDay1(path string) {
	sum, err := runner.RunPart(path, part1)
	if err != nil {
		fmt.Printf("Error in Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Part 1 is %d\n", sum)
	}

	sum, err = runner.RunPart(path, part2)
	if err != nil {
		fmt.Printf("Error in Part 2: %s\n", err)
	} else {
		fmt.Printf("Answer to Part 2 is %d\n", sum)
	}
}

func part1(file io.Reader) (int, error) {
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

func part2(file io.Reader) (int, error) {
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
					leftDigit = containsNumberWord(line[:idx+1])
				}
			}
			if rightDigit == -1 {
				if unicode.IsDigit(rune(line[len(line)-idx-1])) {
					rightDigit = int(rune(line[len(line)-idx-1]) - '0')
				} else {
					// returns -1 if not found
					rightDigit = containsNumberWord(line[len(line)-idx-1:])
				}
			}
		}
		sum += 10*leftDigit + rightDigit
	}

	return sum, nil
}

var numberWords = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func containsNumberWord(word string) int {
	for idx, number := range numberWords {
		if strings.Contains(word, number) {
			return idx + 1
		}
	}
	return -1
}
