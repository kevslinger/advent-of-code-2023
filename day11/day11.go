package day11

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"

	"github.com/kevslinger/advent-of-code-2023/runner"
)

func RunDay11(path string) {
	sum, err := runner.RunPart(path, Part1)
	if err != nil {
		fmt.Printf("Error with Day 11 Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 11 Part 1 is: %d\n", sum)
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error reading file in Day 11 Part 2: %s\n", err)
		return
	}
	defer file.Close()
	sum = Part2(file, 1_000_000)
	fmt.Printf("Answer to Day 11 Part 2 is: %d\n", sum)
}

type Galaxy struct {
	X int
	Y int
}

func Part1(file io.Reader) (int, error) {
	universe := ScanUniverse(file)
	galaxies := ParseGalaxies(universe, 2)
	return GetSumGalaxyDistances(galaxies), nil
}

func Part2(file io.Reader, emptyDistance int) int {
	universe := ScanUniverse(file)
	galaxies := ParseGalaxies(universe, emptyDistance)
	return GetSumGalaxyDistances(galaxies)
}

func ScanUniverse(file io.Reader) []string {
	scanner := bufio.NewScanner(file)
	universe := make([]string, 0)
	for scanner.Scan() {
		universe = append(universe, scanner.Text())
	}
	return universe
}

func ParseGalaxies(universe []string, emptyDistance int) []*Galaxy {
	galaxies := make([]*Galaxy, 0)

	galaxyMap := make(map[int][]*Galaxy)

	emptyRowCounter := 0
	for idx, row := range universe {
		if !strings.Contains(row, "#") {
			emptyRowCounter += emptyDistance - 1
			continue
		}
		for j := 0; j < len(row); j++ {
			if row[j] == '#' {
				_, ok := galaxyMap[j]
				if !ok {
					galaxyMap[j] = make([]*Galaxy, 0)
				}
				galaxyMap[j] = append(galaxyMap[j], &Galaxy{X: idx + emptyRowCounter})
			}
		}
	}
	emptyColCounter := 0
	for j := 0; j < len(universe[0]); j++ {
		sawGalaxy := false
		for i := 0; i < len(universe); i++ {
			if universe[i][j] == '#' {
				sawGalaxy = true
				galaxyMap[j][0].Y = j + emptyColCounter
				galaxies = append(galaxies, galaxyMap[j][0])
				galaxyMap[j] = galaxyMap[j][1:]
			}
		}
		if !sawGalaxy {
			emptyColCounter += emptyDistance - 1
		}
	}
	return galaxies
}

func GetSumGalaxyDistances(galaxies []*Galaxy) int {
	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += ManhattanDistance(galaxies[i], galaxies[j])
		}
	}
	return sum
}

func ManhattanDistance(g1, g2 *Galaxy) int {
	return int(math.Abs(float64(g2.X)-float64(g1.X))) + int(math.Abs(float64(g2.Y)-float64(g1.Y)))
}
