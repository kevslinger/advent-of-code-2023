package day10

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kevslinger/advent-of-code-2023/runner"
)

type coord struct {
	X int
	Y int
}

var (
	north = coord{-1, 0}
	south = coord{1, 0}
	east  = coord{0, 1}
	west  = coord{0, -1}
)

var tiles = map[byte][]coord{
	'|': {north, south},
	'-': {east, west},
	'L': {north, east},
	'J': {north, west},
	'7': {south, west},
	'F': {south, east},
}

func RunDay10(path string) {
	steps, err := runner.RunPart(path, part1)
	if err != nil {
		fmt.Printf("Error with Day 10 Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 10 Part 1: %d\n", steps)
	}
}

func part1(file io.Reader) (int, error) {
	maze := parseFileIntoMaze(bufio.NewScanner(file))
	startingcoord := findStartingPosition(maze)
	if startingcoord.X == -1 || startingcoord.Y == -1 {
		return -1, fmt.Errorf("unable to find starting position S in maze")
	}
	furthestSteps := breadthFirstSearch(maze, startingcoord)
	return furthestSteps, nil
}

func parseFileIntoMaze(scanner *bufio.Scanner) []string {
	maze := make([]string, 0)
	for scanner.Scan() {
		maze = append(maze, scanner.Text())
	}
	return maze
}

func findStartingPosition(maze []string) coord {
	for i, row := range maze {
		for j := 0; j < len(row); j++ {
			if row[j] == 'S' {
				return coord{i, j}
			}
		}
	}
	return coord{-1, -1}
}

func breadthFirstSearch(maze []string, start coord) int {
	visited := map[coord]bool{
		start: true,
	}
	queue := findStartingConnectors(maze, start)
	steps := 1
	for queue[0] != queue[1] {
		visited[queue[0]] = true
		visited[queue[1]] = true
		steps++
		queue[0] = moveOneStep(maze, queue[0], visited)
		queue[1] = moveOneStep(maze, queue[1], visited)
	}
	return steps
}

func findStartingConnectors(maze []string, start coord) []coord {
	nextLocations := make([]coord, 0)
	for _, dir := range []coord{north, south, east, west} {
		newX, newY := start.X+dir.X, start.Y+dir.Y
		// Bounds-checking
		if boundsCheckOk(maze, newX, newY) {
			// Check that the next piece can connect to the current piece
			conns, ok := tiles[maze[newX][newY]]
			if !ok {
				continue
			}
			// Check if either of the connections connects to us
			for _, con := range conns {
				if newX+con.X == start.X && newY+con.Y == start.Y {
					nextLocations = append(nextLocations, coord{newX, newY})
					break
				}
			}
		}
	}
	return nextLocations
}

func boundsCheckOk(maze []string, x, y int) bool {
	return x >= 0 && x < len(maze) && y >= 0 && y < len(maze[x])
}

func moveOneStep(maze []string, pos coord, visited map[coord]bool) coord {
	// Assume that one position can always be moved to
	// Assume that the connections all exist and do not go out of bounds
	dirs := tiles[maze[pos.X][pos.Y]]
	for _, dir := range dirs {
		potentialStep := coord{pos.X + dir.X, pos.Y + dir.Y}
		_, ok := visited[potentialStep]
		if !ok {
			return potentialStep
		}
	}
	return coord{}
}
