package day6

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func RunDay6(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error processing file in Day 6 Part 1: %s\n", err)
		return
	}
	defer file.Close()

	raceProduct, err := Part1(file)
	if err != nil {
		fmt.Printf("Error processing Day 6 Part 1: %s\n", err)
	} else {
		fmt.Printf("The answer to Day 6 Part 1 is: %d\n", raceProduct)
	}

	file2, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error processing file in Day 6 Part 1: %s\n", err)
		return
	}
	defer file2.Close()

	numWins, err := Part2(file2)
	if err != nil {
		fmt.Printf("Error processing Day 6 Part 2: %s\n", err)
	} else {
		fmt.Printf("The answer to Day 6 Part 2 is: %d\n", numWins)
	}
}

func Part1(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)

	matchNumbers := regexp.MustCompile("[0-9]+")
	scanner.Scan()
	times, err := ConvertStringToInt(matchNumbers.FindAllString(scanner.Text(), -1))
	if err != nil {
		return -1, err
	}
	scanner.Scan()
	distances, err := ConvertStringToInt(matchNumbers.FindAllString(scanner.Text(), -1))
	if err != nil {
		return -1, err
	}
	races := make([]Race, len(distances))
	raceProduct := 1
	for idx, _ := range races {
		races[idx] = Race{times[idx], distances[idx]}
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

func Part2(file io.Reader) (int, error) {
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
	race := Race{time, distance}
	numWins := 0
	for i := 1; i < race.Time; i++ {
		if i*(race.Time-i) > race.Distance {
			numWins++
		}
	}

	return numWins, nil
}

func ConvertStringToInt(strs []string) ([]int, error) {
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

type Race struct {
	Time     int
	Distance int
}
