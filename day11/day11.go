package day11

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"

	"github.com/kevslinger/advent-of-code-2023/runner"
)

func RunDay11(path string) {
	sum, err := runner.RunPart(path, part1)
	if err != nil {
		fmt.Printf("Error with Day 11 Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 11 Part 1 is: %d\n", sum)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error with Day 11 Part 2: %s\n", err)
	}
	defer file.Close()
	sum = part2(file, 1_000_000)
	if err != nil {
		fmt.Printf("Error with Day 11 Part 2: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 11 Part 2 is: %d\n", sum)
	}
}

type galaxy struct {
	x int
	y int
}

func part1(file io.Reader) (int, error) {
	universe := ScanUniverse(file)
	galaxies := parseGalaxies(universe, 2)
	return getSumGalaxyDistances(galaxies), nil
}

func part2(file io.Reader, emptyDistance int) int {
	universe := ScanUniverse(file)
	galaxies := parseGalaxies(universe, emptyDistance)
	return getSumGalaxyDistances(galaxies)
}

func ScanUniverse(file io.Reader) []string {
	scanner := bufio.NewScanner(file)
	universe := make([]string, 0)
	for scanner.Scan() {
		universe = append(universe, scanner.Text())
	}
	return universe
}

func parseGalaxies(universe []string, emptyDistance int) []*galaxy {
	galaxyMap := processUniverseRows(universe, emptyDistance)
	galaxies := make([]*galaxy, 0)
	emptyColCounter := 0
	for j := 0; j < len(universe[0]); j++ {
		sawgalaxy := false
		for i := 0; i < len(universe); i++ {
			if universe[i][j] == '#' {
				sawgalaxy = true
				galaxyMap[j][0].y = j + emptyColCounter
				galaxies = append(galaxies, galaxyMap[j][0])
				galaxyMap[j] = galaxyMap[j][1:]
			}
		}
		if !sawgalaxy {
			emptyColCounter += emptyDistance - 1
		}
	}
	return galaxies
}

func processUniverseRows(universe []string, emptyDistance int) map[int][]*galaxy {
	galaxyMap := make(map[int][]*galaxy)
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
					galaxyMap[j] = make([]*galaxy, 0)
				}
				galaxyMap[j] = append(galaxyMap[j], &galaxy{x: idx + emptyRowCounter})
			}
		}
	}
	return galaxyMap
}

func getSumGalaxyDistances(galaxies []*galaxy) int {
	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += manhattanDistance(galaxies[i], galaxies[j])
		}
	}
	return sum
}

func manhattanDistance(g1, g2 *galaxy) int {
	return int(math.Abs(float64(g2.x)-float64(g1.x))) + int(math.Abs(float64(g2.y)-float64(g1.y)))
}
