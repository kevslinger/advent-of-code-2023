package day3

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"unicode"

	"github.com/kevslinger/advent-of-code-2023/runner"
)

func RunDay3(path string) {
	sum, err := runner.RunPart(path, part1)
	if err != nil {
		fmt.Printf("Error with Day 3 Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 3 Part 1: %d\n", sum)
	}

	sum, err = runner.RunPart(path, part2)
	if err != nil {
		fmt.Printf("Error with Day 3 Part 2: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 3 Part 2: %d\n", sum)
	}
}

func part1(file io.Reader) (int, error) {
	schem, err := parseTable(file)
	if err != nil {
		return -1, err
	}

	return schem.calculateSumOfPartNumbers(), nil
}

func part2(file io.Reader) (int, error) {
	schem, err := parseTable(file)
	if err != nil {
		return -1, err
	}

	return schem.calculateSumOfGearRatios(), nil
}

type schematic struct {
	table  [][]int
	lookup map[int]int
}

func parseTable(file io.Reader) (schematic, error) {
	scanner := bufio.NewScanner(file)
	table := make([][]int, 0)
	lookup := make(map[int]int)
	// Start at 1 for assigning keys to our table
	numParts := 1
	for scanner.Scan() {
		line := scanner.Text()
		gridLine := make([]int, len(line))
		for idx, char := range line {
			switch char {
			case '.':
				gridLine[idx] = 0
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				// Already initialised: skip
				if gridLine[idx] != 0 {
					continue
				}
				// Traverse the line to get the full number
				idx2 := idx + 1
				for idx2 < len(line) && unicode.IsDigit(rune(line[idx2])) {
					idx2++
				}
				partNumber, err := strconv.Atoi(line[idx:idx2])
				if err != nil {
					return schematic{}, err
				}
				// Generate ID for this part
				partNumberId := numParts
				lookup[partNumberId] = partNumber
				numParts++
				for i := idx; i < idx2; i++ {
					gridLine[i] = partNumberId
				}
			// Stars are special because they are gears
			case '*':
				gridLine[idx] = -2
			default:
				gridLine[idx] = -1
			}
		}
		table = append(table, gridLine)
	}
	return schematic{table, lookup}, nil
}

func (s *schematic) calculateSumOfPartNumbers() int {
	directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	sum := 0
	for i, row := range s.table {
		for j, val := range row {
			// 0 represents blank, < 0 represents a symbol, >0 represents parts
			if val < 0 {
				// Check the 8 neighbours
				for _, dir := range directions {
					newI, newJ := i+dir[0], j+dir[1]
					// Check boundaries
					if newI >= len(s.table) || newI < 0 || newJ >= len(row) || newJ < 0 {
						continue
					}
					// Found a part!
					if s.table[newI][newJ] > 0 {
						sum += s.lookup[s.table[newI][newJ]]
						// Set to 0 to avoid double counting
						s.lookup[s.table[newI][newJ]] = 0
					}
				}
			}
		}
	}
	return sum
}

func (s *schematic) calculateSumOfGearRatios() int {
	directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	sum := 0
	for i, row := range s.table {
		for j, val := range row {
			// -2 represents a gear ('*')
			if val == -2 {
				numPartsFound, gearRatio := 0, 1
				// Check the 8 neighbours
				for _, dir := range directions {
					newI, newJ := i+dir[0], j+dir[1]
					// Check boundaries
					if newI >= len(s.table) || newI < 0 || newJ >= len(row) || newJ < 0 {
						continue
					}
					// Found a part!
					if s.table[newI][newJ] > 0 && s.lookup[s.table[newI][newJ]] > 0 {
						numPartsFound++
						gearRatio *= s.lookup[s.table[newI][newJ]]
						// Set to 0 to avoid double counting
						s.lookup[s.table[newI][newJ]] = 0
					}
				}
				// Only include gears who have exactly 2 parts
				if numPartsFound == 2 {
					sum += gearRatio
				}
			}
		}
	}
	return sum
}
