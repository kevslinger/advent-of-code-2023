package day9

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/kevslinger/advent-of-code-2023/runner"
)

func RunDay9(path string) {
	sum, err := runner.RunPart(path, Part1)
	if err != nil {
		fmt.Printf("Error processing Day 9 Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 9 Part 1 is : %d\n", sum)
	}

	sum, err = runner.RunPart(path, Part2)
	if err != nil {
		fmt.Printf("Error processing Day 9 Part 2: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 9 Part 2 is : %d\n", sum)
	}
}

func Part1(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)
	numbers, err := ParseNumbers(scanner)
	if err != nil {
		return -1, err
	}

	sum := 0
	for _, nums := range numbers {
		progression := GetProgression(nums)
		sum += GetNextNumber(progression)
	}
	return sum, nil
}

func Part2(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)
	numbers, err := ParseNumbers(scanner)
	if err != nil {
		return -1, err
	}

	sum := 0
	for _, nums := range numbers {
		progression := GetProgression(nums)
		sum += GetPreviousNumber(progression)
	}
	return sum, nil
}

func ParseNumbers(scanner *bufio.Scanner) ([][]int, error) {
	nums := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		toks := strings.Split(line, " ")
		arr := make([]int, len(toks))
		for idx, tok := range toks {
			num, err := strconv.Atoi(tok)
			if err != nil {
				return nil, err
			}
			arr[idx] = num
		}
		nums = append(nums, arr)
	}
	return nums, nil
}

// BRUTE FORCE
func GetProgression(nums []int) [][]int {
	n := len(nums)
	progression := make([][]int, 0)
	progression = append(progression, nums)
	curArr := nums
	i := 1
	for true {
		allZero := true
		newArr := make([]int, n)
		for idx := i; idx < len(curArr); idx++ {
			newVal := curArr[idx] - curArr[idx-1]
			if newVal != 0 {
				allZero = false
			}
			newArr[idx] = newVal
		}
		progression = append(progression, newArr)
		if allZero {
			return progression
		}
		curArr = newArr
		i++
	}
	return make([][]int, 0)
}

func GetNextNumber(progression [][]int) int {
	nextNum := 0
	for i := len(progression) - 1; i >= 0; i-- {
		nextNum += progression[i][len(progression[i])-1]
	}
	return nextNum
}

func GetPreviousNumber(progression [][]int) int {
	prevNum := 0
	for i := len(progression) - 1; i >= 0; i-- {
		prevNum = progression[i][i] - prevNum
	}
	return prevNum
}
