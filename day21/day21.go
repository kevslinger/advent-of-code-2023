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

	plots, err := part1(file, 64)
	if err != nil {
		fmt.Printf("Error with Day 21 Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 21 Part 1: %d\n", plots)
	}
}

type position struct {
	x int
	y int
}

func (p position) add(p2 position) position {
	return position{p.x + p2.x, p.y + p2.y}
}

type direction int

const (
	north direction = iota + 1
	south
	east
	west
)

var dirMap map[direction]position = map[direction]position{
	north: {-1, 0},
	south: {1, 0},
	east:  {0, 1},
	west:  {0, -1},
}

func part1(file io.Reader, steps int) (int, error) {
	var garden []string = parseGarden(file)
	startposition, err := getStartPos(garden)
	if err != nil {
		return -1, err
	}
	return calculateGardenSpots(garden, startposition, steps), nil
}

func parseGarden(file io.Reader) []string {
	garden := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		garden = append(garden, scanner.Text())
	}
	return garden
}

func getStartPos(garden []string) (position, error) {
	for i, row := range garden {
		for j, char := range row {
			if char == 'S' {
				return position{x: i, y: j}, nil
			}
		}
	}
	return position{}, fmt.Errorf("cannot find starting position in garden")
}

func calculateGardenSpots(garden []string, startPos position, steps int) int {
	queue := make([]position, 0)
	queue = append(queue, startPos)
	for i := 0; i < steps; i++ {
		queueLen := len(queue)
		newCoveredSpaces := make(map[position]bool)
		for j := 0; j < queueLen; j++ {
			curSpace := queue[0]
			queue = queue[1:]

			// Check all possible directions of movement
			for _, dir := range dirMap {
				newPos := curSpace.add(dir)

				// Don't want to add 2 of the same space
				_, ok := newCoveredSpaces[newPos]
				if !isInBounds(newPos, garden) || garden[newPos.x][newPos.y] == '#' || ok {
					continue
				}
				queue = append(queue, newPos)
				newCoveredSpaces[newPos] = true
			}
		}
	}
	return len(queue)
}

func isInBounds(pos position, garden []string) bool {
	return pos.x >= 0 && pos.x < len(garden) && pos.y >= 0 && pos.y <= len(garden[pos.x])
}
