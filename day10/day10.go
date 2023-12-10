package day10

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kevslinger/advent-of-code-2023/runner"
)

type Coord struct {
	X int
	Y int
}

var (
	North = Coord{-1, 0}
	South = Coord{1, 0}
	East  = Coord{0, 1}
	West  = Coord{0, -1}
)

var tiles = map[byte][]Coord{
	'|': {North, South},
	'-': {East, West},
	'L': {North, East},
	'J': {North, West},
	'7': {South, West},
	'F': {South, East},
}

func RunDay10(path string) {
	steps, err := runner.RunPart(path, Part1)
	if err != nil {
		fmt.Printf("Error with Day 10 Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 10 Part 1: %d\n", steps)
	}
}

func Part1(file io.Reader) (int, error) {
	maze := ParseFileIntoMaze(bufio.NewScanner(file))
	startingCoord := FindStartingPosition(maze)
	if startingCoord.X == -1 || startingCoord.Y == -1 {
		return -1, fmt.Errorf("Unable to find starting position S in maze")
	}
	furthestSteps := BreadthFirstSearch(maze, startingCoord)
	return furthestSteps, nil
}

func ParseFileIntoMaze(scanner *bufio.Scanner) []string {
	maze := make([]string, 0)
	for scanner.Scan() {
		maze = append(maze, scanner.Text())
	}
	return maze
}

func FindStartingPosition(maze []string) Coord {
	for i, row := range maze {
		for j := 0; j < len(row); j++ {
			if row[j] == 'S' {
				return Coord{i, j}
			}
		}
	}
	return Coord{-1, -1}
}

func BreadthFirstSearch(maze []string, start Coord) int {
	visited := map[Coord]bool{
		start: true,
	}
	queue := FindStartingConnectors(maze, start)
	steps := 1
	for queue[0] != queue[1] {
		visited[queue[0]] = true
		visited[queue[1]] = true
		steps++
		queue[0] = MoveOneStep(maze, queue[0], visited)
		queue[1] = MoveOneStep(maze, queue[1], visited)
	}
	return steps
}

func FindStartingConnectors(maze []string, start Coord) []Coord {
	nextLocations := make([]Coord, 0)
	for _, dir := range []Coord{North, South, East, West} {
		newX, newY := start.X+dir.X, start.Y+dir.Y
		// Bounds-checking
		if BoundsCheckOk(maze, newX, newY) {
			// Check that the next piece can connect to the current piece
			conns, ok := tiles[maze[newX][newY]]
			if !ok {
				continue
			}
			// Check if either of the connections connects to us
			for _, con := range conns {
				if newX+con.X == start.X && newY+con.Y == start.Y {
					nextLocations = append(nextLocations, Coord{newX, newY})
					break
				}
			}
		}
	}
	return nextLocations
}

func BoundsCheckOk(maze []string, x, y int) bool {
	return x >= 0 && x < len(maze) && y >= 0 && y < len(maze[x])
}

func MoveOneStep(maze []string, pos Coord, visited map[Coord]bool) Coord {
	// Assume that one position can always be moved to
	// Assume that the connections all exist and do not go out of bounds
	dirs := tiles[maze[pos.X][pos.Y]]
	for _, dir := range dirs {
		potentialStep := Coord{pos.X + dir.X, pos.Y + dir.Y}
		_, ok := visited[potentialStep]
		if !ok {
			return potentialStep
		}
	}
	return Coord{}
}
