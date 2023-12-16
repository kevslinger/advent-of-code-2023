package day14

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kevslinger/advent-of-code-2023/runner"
)

func RunDay14(path string) {
	load, err := runner.RunPart(path, Part1)
	if err != nil {
		fmt.Printf("Error with Day 14 Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 14 Part 1 is: %d\n", load)
	}
}

func Part1(file io.Reader) (int, error) {
	var board []string = ParseRockBoard(file)
	var load int = CalculateLoad(board)
	return load, nil
}

func ParseRockBoard(file io.Reader) []string {
	scanner := bufio.NewScanner(file)
	var board []string = make([]string, 0)
	for scanner.Scan() {
		board = append(board, scanner.Text())
	}
	return board
}

func CalculateLoad(board []string) int {
	var highestRocks []int = make([]int, len(board[0]))
	for i := 0; i < len(highestRocks); i++ {
		highestRocks[i] = len(board)
	}
	var totalLoad int
	for i, line := range board {
		for j, char := range line {
			switch char {
			// Square rock
			case '#':
				highestRocks[j] = len(board) - i - 1
			// Rounded rock
			case 'O':
				totalLoad += highestRocks[j]
				highestRocks[j]--
			}
		}
	}
	return totalLoad
}
