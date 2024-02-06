package day23

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/kevslinger/advent-of-code-2023/runner"
)

type direction byte

const (
	north direction = '^'
	south direction = 'v'
	east  direction = '>'
	west  direction = '<'
	empty direction = '.'
)

type position struct {
	x int
	y int
}

var (
	northPos position = position{-1, 0}
	southPos position = position{1, 0}
	eastPos  position = position{0, 1}
	westPos  position = position{0, -1}
	EmptyPos position = position{0, 0}
)

var directionPos = []position{northPos, southPos, eastPos, westPos}

func (p position) add(p2 position) position {
	return position{x: p.x + p2.x, y: p.y + p2.y}
}

type traversalState struct {
	CurPos  position
	Steps   int
	Visited map[position]bool
}

func RunDay23(path string) {
	steps, err := runner.RunPart(path, part1)
	if err != nil {
		fmt.Printf("Error with Day 23 Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 23 Part 1 is: %d\n", steps)
	}
	// steps, err = runner.RunPart(path, Part2)
	// if err != nil {
	// 	fmt.Printf("Error with Day 23 Part 2: %s\n", err)
	// } else {
	// 	fmt.Printf("Answer to Day 23 Part 2 is: %d\n", steps)
	// }
}

func part1(file io.Reader) (int, error) {
	maze := parseMaze(bufio.NewScanner(file))
	var dirMap map[direction]position = map[direction]position{
		north: northPos,
		south: southPos,
		east:  eastPos,
		west:  westPos,
		empty: EmptyPos,
	}
	return traverseMazeDFS(maze, dirMap), nil
}

func part2(file io.Reader) (int, error) {
	maze := parseMaze(bufio.NewScanner(file))
	var dirMap map[direction]position = map[direction]position{
		north: EmptyPos,
		south: EmptyPos,
		east:  EmptyPos,
		west:  EmptyPos,
		empty: EmptyPos,
	}
	return traverseMazeDFS(maze, dirMap), nil
}

func parseMaze(scanner *bufio.Scanner) []string {
	var maze []string = make([]string, 0)
	for scanner.Scan() {
		maze = append(maze, scanner.Text())
	}
	return maze
}

func findOpenPos(rowNum int, row string) position {
	return position{x: rowNum, y: strings.Index(row, ".")}
}

// Performing a recursive DFS proves to be rather expensive...
// Perhaps some form of a bottom-up approach would be more efficient
func traverseMazeDFS(maze []string, dirMap map[direction]position) int {
	endPos := findOpenPos(len(maze)-1, maze[len(maze)-1])
	startState := traversalState{CurPos: findOpenPos(0, maze[0]), Visited: make(map[position]bool)}
	return recursiveTraverse(maze, startState, endPos, dirMap)
}

func recursiveTraverse(maze []string, curState traversalState, endPos position, dirMap map[direction]position) int {
	if !IsInBounds(curState, maze) || maze[curState.CurPos.x][curState.CurPos.y] == '#' {
		return 0
	}
	if curState.CurPos == endPos {
		return curState.Steps
	}
	seen, ok := curState.Visited[curState.CurPos]
	// Cannot go somewhere we've already been
	if ok && seen {
		return 0
	}

	curState.Visited[curState.CurPos] = true
	var longestTrip int
	switch dirMap[direction(maze[curState.CurPos.x][curState.CurPos.y])] {
	case northPos:
		longestTrip = recursiveTraverse(maze, traversalState{CurPos: curState.CurPos.add(northPos), Steps: curState.Steps + 1, Visited: curState.Visited}, endPos, dirMap)
	case southPos:
		longestTrip = recursiveTraverse(maze, traversalState{CurPos: curState.CurPos.add(southPos), Steps: curState.Steps + 1, Visited: curState.Visited}, endPos, dirMap)
	case eastPos:
		longestTrip = recursiveTraverse(maze, traversalState{CurPos: curState.CurPos.add(eastPos), Steps: curState.Steps + 1, Visited: curState.Visited}, endPos, dirMap)
	case westPos:
		longestTrip = recursiveTraverse(maze, traversalState{CurPos: curState.CurPos.add(westPos), Steps: curState.Steps + 1, Visited: curState.Visited}, endPos, dirMap)
	case EmptyPos:
		for _, dir := range directionPos {
			steps := recursiveTraverse(maze, traversalState{CurPos: curState.CurPos.add(dir), Steps: curState.Steps + 1, Visited: curState.Visited}, endPos, dirMap)
			if steps > longestTrip {
				longestTrip = steps
			}
		}
	default:
		return 0
	}
	curState.Visited[curState.CurPos] = false
	return longestTrip
}

func IsInBounds(state traversalState, maze []string) bool {
	// Check bounds
	if state.CurPos.x < 0 || state.CurPos.x >= len(maze) || state.CurPos.y < 0 || state.CurPos.y >= len(maze) {
		return false
	}
	return true
}
