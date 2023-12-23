package day21

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func RunDay21(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error opening file in Day 21 Part 1: %s\n", err)
	}
	defer file.Close()

	plots, err := Part1(file, 64)
	if err != nil {
		fmt.Printf("Error with Day 21 Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 21 Part 1: %d\n", plots)
	}
}

type Position struct {
	X int
	Y int
}

func (p Position) add(p2 Position) Position {
	return Position{p.X + p2.X, p.Y + p2.Y}
}

type Direction int

const (
	Dummy Direction = iota
	North
	South
	East
	West
)

var directions = []Direction{North, South, East, West}

var dirMap map[Direction]Position = map[Direction]Position{
	North: {-1, 0},
	South: {1, 0},
	East:  {0, 1},
	West:  {0, -1},
}

func Part1(file io.Reader, steps int) (int, error) {
	var garden []string = ParseGarden(file)
	startPosition, err := GetStartPos(garden)
	if err != nil {
		return -1, err
	}
	return CalculateGardenSpots(garden, startPosition, steps), nil
}

func ParseGarden(file io.Reader) []string {
	garden := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		garden = append(garden, scanner.Text())
	}
	return garden
}

func GetStartPos(garden []string) (Position, error) {
	for i, row := range garden {
		for j, char := range row {
			if char == 'S' {
				return Position{X: i, Y: j}, nil
			}
		}
	}
	return Position{}, fmt.Errorf("Cannot find starting position in garden")
}

func CalculateGardenSpots(garden []string, startPos Position, steps int) int {
	queue := make([]Position, 0)
	queue = append(queue, startPos)
	for i := 0; i < steps; i++ {
		queueLen := len(queue)
		newCoveredSpaces := make(map[Position]bool)
		for j := 0; j < queueLen; j++ {
			curSpace := queue[0]
			queue = queue[1:]

			// Check all possible directions of movement
			for _, dir := range dirMap {
				newPos := curSpace.add(dir)

				// Don't want to add 2 of the same space
				_, ok := newCoveredSpaces[newPos]
				if !IsInBounds(newPos, garden) || garden[newPos.X][newPos.Y] == '#' || ok {
					continue
				}
				queue = append(queue, newPos)
				newCoveredSpaces[newPos] = true
			}
		}
	}
	return len(queue)
}

func IsInBounds(pos Position, garden []string) bool {
	return pos.X >= 0 && pos.X < len(garden) && pos.Y >= 0 && pos.Y <= len(garden[pos.X])
}
