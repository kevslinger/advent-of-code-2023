package day23

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/kevslinger/advent-of-code-2023/runner"
)

type Direction byte

const (
	North Direction = '^'
	South Direction = 'v'
	East  Direction = '>'
	West  Direction = '<'
	Empty Direction = '.'
)

var directions = []Direction{North, South, East, West}

type Position struct {
	X int
	Y int
}

var (
	NorthPos Position = Position{-1, 0}
	SouthPos Position = Position{1, 0}
	EastPos  Position = Position{0, 1}
	WestPos  Position = Position{0, -1}
	EmptyPos Position = Position{0, 0}
)

var directionPos = []Position{NorthPos, SouthPos, EastPos, WestPos}

func (p Position) add(p2 Position) Position {
	return Position{X: p.X + p2.X, Y: p.Y + p2.Y}
}

type TraversalState struct {
	CurPos  Position
	Steps   int
	Visited map[Position]bool
}

func RunDay23(path string) {
	steps, err := runner.RunPart(path, Part1)
	if err != nil {
		fmt.Printf("Error with Day 23 Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 23 Part 1 is: %d\n", steps)
	}
	steps, err = runner.RunPart(path, Part2)
	if err != nil {
		fmt.Printf("Error with Day 23 Part 2: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 23 Part 2 is: %d\n", steps)
	}
}

func Part1(file io.Reader) (int, error) {
	maze := ParseMaze(bufio.NewScanner(file))
	var dirMap map[Direction]Position = map[Direction]Position{
		North: NorthPos,
		South: SouthPos,
		East:  EastPos,
		West:  WestPos,
		Empty: EmptyPos,
	}
	return TraverseMazeDFS(maze, dirMap), nil
	//return TraverseMazeBFS(maze), nil
}

func Part2(file io.Reader) (int, error) {
	maze := ParseMaze(bufio.NewScanner(file))
	var dirMap map[Direction]Position = map[Direction]Position{
		North: EmptyPos,
		South: EmptyPos,
		East:  EmptyPos,
		West:  EmptyPos,
		Empty: EmptyPos,
	}
	return TraverseMazeDFS(maze, dirMap), nil
}

func ParseMaze(scanner *bufio.Scanner) []string {
	var maze []string = make([]string, 0)
	for scanner.Scan() {
		maze = append(maze, scanner.Text())
	}
	return maze
}

func FindOpenPos(rowNum int, row string) Position {
	return Position{X: rowNum, Y: strings.Index(row, ".")}
}

// Performing a recursive DFS proves to be rather expensive...
// Perhaps some form of a bottom-up approach would be more efficient
func TraverseMazeDFS(maze []string, dirMap map[Direction]Position) int {
	endPos := FindOpenPos(len(maze)-1, maze[len(maze)-1])
	startState := TraversalState{CurPos: FindOpenPos(0, maze[0]), Visited: make(map[Position]bool)}
	return RecursiveTraverse(maze, startState, endPos, dirMap)
}

func RecursiveTraverse(maze []string, curState TraversalState, endPos Position, dirMap map[Direction]Position) int {
	if !IsInBounds(curState, maze) || maze[curState.CurPos.X][curState.CurPos.Y] == '#' {
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
	switch dirMap[Direction(maze[curState.CurPos.X][curState.CurPos.Y])] {
	case NorthPos:
		longestTrip = RecursiveTraverse(maze, TraversalState{CurPos: curState.CurPos.add(NorthPos), Steps: curState.Steps + 1, Visited: curState.Visited}, endPos, dirMap)
	case SouthPos:
		longestTrip = RecursiveTraverse(maze, TraversalState{CurPos: curState.CurPos.add(SouthPos), Steps: curState.Steps + 1, Visited: curState.Visited}, endPos, dirMap)
	case EastPos:
		longestTrip = RecursiveTraverse(maze, TraversalState{CurPos: curState.CurPos.add(EastPos), Steps: curState.Steps + 1, Visited: curState.Visited}, endPos, dirMap)
	case WestPos:
		longestTrip = RecursiveTraverse(maze, TraversalState{CurPos: curState.CurPos.add(WestPos), Steps: curState.Steps + 1, Visited: curState.Visited}, endPos, dirMap)
	case EmptyPos:
		for _, dir := range directionPos {
			steps := RecursiveTraverse(maze, TraversalState{CurPos: curState.CurPos.add(dir), Steps: curState.Steps + 1, Visited: curState.Visited}, endPos, dirMap)
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

func IsInBounds(state TraversalState, maze []string) bool {
	// Check bounds
	if state.CurPos.X < 0 || state.CurPos.X >= len(maze) || state.CurPos.Y < 0 || state.CurPos.Y >= len(maze) {
		return false
	}
	return true
}
