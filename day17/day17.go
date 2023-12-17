package day17

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"github.com/kevslinger/advent-of-code-2023/runner"
)

func Runday17(path string) {
	heatLoss, err := runner.RunPart(path, Part1)
	if err != nil {
		fmt.Printf("Error with Day 17 Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 17 Part 1 is: %d\n", heatLoss)
	}
	heatLoss, err = runner.RunPart(path, Part2)
	if err != nil {
		fmt.Printf("Error with Day 17 Part 2: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 17 Part 2 is: %d\n", heatLoss)
	}
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

var dirMap map[Direction][]int = map[Direction][]int{
	North: {-1, 0},
	South: {1, 0},
	East:  {0, 1},
	West:  {0, -1},
}

type TraversalState struct {
	X            int
	Y            int
	HeatLost     int
	Dir          Direction
	TimesDirUsed int
}

func Part1(file io.Reader) (int, error) {
	blocks, err := ParseCityBlocks(file)
	if err != nil {
		return -1, err
	}
	return CalculateHeatLost(blocks, 0, 3), nil
}

func Part2(file io.Reader) (int, error) {
	blocks, err := ParseCityBlocks(file)
	if err != nil {
		return -1, err
	}
	return CalculateHeatLost(blocks, 4, 10), nil
}

func ParseCityBlocks(file io.Reader) ([][]int, error) {
	blocks := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		block := make([]int, len(line))
		for idx, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				return make([][]int, 0), err
			}
			block[idx] = num
		}
		blocks = append(blocks, block)
	}
	return blocks, nil
}

func CalculateHeatLost(blocks [][]int, minMovement int, maxMovement int) int {
	var heatLossMap [][]int = TraverseCityBlocks(blocks, minMovement, maxMovement)
	// Starting block shouldn't count
	return heatLossMap[0][0] - blocks[0][0]
}

func TraverseCityBlocks(blocks [][]int, minMovements int, maxMovements int) [][]int {
	var heatLossMap [][]int = CreateHeatLossMap(blocks)
	heatTracker := CreateHeatMap(blocks)
	// Create a queue to hold the traversal states
	var queue []TraversalState = make([]TraversalState, 0)
	// Start at the bottom corner and work our way up
	queue = append(queue, TraversalState{X: len(heatLossMap) - 1, Y: len(heatLossMap[0]) - 1, HeatLost: blocks[len(blocks)-1][len(blocks[0])-1], Dir: Dummy, TimesDirUsed: maxMovements})
	for len(queue) > 0 {
		curState := queue[0]
		// Pop
		queue = queue[1:]
		// Don't continue once we reach the start
		if curState.X == 0 && curState.Y == 0 {
			continue
		}
		// Check all directions
		for _, dir := range directions {
			newState := ProcessState(blocks, heatLossMap, heatTracker, curState, dir, minMovements, maxMovements)
			// Use HeatLost as a pseudo nil checker since heatLost is monotonically increasing
			if newState.HeatLost > 0 {
				queue = append(queue, newState)
			}
		}
	}
	return heatLossMap
}

func TraverseCityBlocksPart2(blocks [][]int) [][]int {
	var heatLossMap [][]int = CreateHeatLossMap(blocks)
	heatTracker := CreateHeatMap(blocks)
	// Create a queue to hold the traversal states
	var queue []TraversalState = make([]TraversalState, 0)
	// Start at the bottom corner and work our way up
	queue = append(queue, TraversalState{X: len(heatLossMap) - 1, Y: len(heatLossMap[0]) - 1, HeatLost: blocks[len(blocks)-1][len(blocks[0])-1], Dir: Dummy, TimesDirUsed: 4})
	for len(queue) > 0 {
		curState := queue[0]
		//fmt.Printf("Current state is %#v\n", curState)
		// Pop
		queue = queue[1:]
		if curState.X == 0 && curState.Y == 0 {
			continue
		}
		// Check all directions
		for _, dir := range directions {
			newState := ProcessState(blocks, heatLossMap, heatTracker, curState, dir, 4, 10)
			// Use HeatLost as a pseudo nil checker since heatLost is monotonically increasing
			if newState.HeatLost > 0 {
				queue = append(queue, newState)
			}
		}

	}
	return heatLossMap
}

func CreateHeatLossMap(blocks [][]int) [][]int {
	var heatLossMap [][]int = make([][]int, len(blocks))
	for i := 0; i < len(heatLossMap); i++ {
		heatLossMap[i] = make([]int, len(blocks[i]))
	}
	// Initialise the bottom right with the correct value
	heatLossMap[len(heatLossMap)-1][len(heatLossMap[0])-1] = blocks[len(heatLossMap)-1][len(heatLossMap[0])-1]
	return heatLossMap
}

func CreateHeatMap(blocks [][]int) [][]map[Direction][]int {
	var heatLossMap [][]map[Direction][]int = make([][]map[Direction][]int, len(blocks))
	for i := 0; i < len(heatLossMap); i++ {
		heatLossMap[i] = make([]map[Direction][]int, len(blocks[i]))
		for j := 0; j < len(heatLossMap[i]); j++ {
			heatLossMap[i][j] = make(map[Direction][]int)
			for _, dir := range directions {
				heatLossMap[i][j][dir] = make([]int, 11)
			}
		}
	}
	return heatLossMap
}

func ProcessState(blocks [][]int, heatLossMap [][]int, heatTracker [][]map[Direction][]int, state TraversalState, dir Direction, minMovements int, maxMovements int) TraversalState {
	// The crucible has restrictions on which way it can turn
	if !IsValidDirection(state, dir, blocks, minMovements, maxMovements) {
		return TraversalState{}
	}

	// Continue in same direction
	newX := state.X + dirMap[dir][0]
	newY := state.Y + dirMap[dir][1]
	if !IsInBounds(blocks, newX, newY) {
		return TraversalState{}
	}
	var timesDirUsed int
	if state.Dir == dir {
		timesDirUsed = state.TimesDirUsed + 1
	} else {
		timesDirUsed = 1
	}
	// If going in this direction would create a new best score, continue doing it
	newHeatLoss := state.HeatLost + blocks[newX][newY]
	// Update global score
	if heatLossMap[newX][newY] == 0 || newHeatLoss < heatLossMap[newX][newY] {
		heatLossMap[newX][newY] = newHeatLoss
	}
	for i := maxMovements; i >= 0; i-- {
		if i < timesDirUsed {
			break
		}
		if heatTracker[newX][newY][dir][i] == 0 || newHeatLoss < heatTracker[newX][newY][dir][i] && i == timesDirUsed {
			heatTracker[newX][newY][dir][i] = newHeatLoss
			return TraversalState{X: newX, Y: newY, HeatLost: newHeatLoss, Dir: dir, TimesDirUsed: timesDirUsed}
		}
	}
	return TraversalState{}
}

func IsValidDirection(state TraversalState, nextDir Direction, board [][]int, minDir int, maxDir int) bool {
	if IsOppositeDirection(state.Dir, nextDir) {
		return false
	}
	if nextDir == state.Dir {
		return state.TimesDirUsed < maxDir
	} else {
		switch nextDir {
		case North:
			if state.X <= minDir-1 {
				return false
			}
		case South:
			if state.X >= len(board)-minDir {
				return false
			}
		case East:
			if state.Y >= len(board[0])-minDir {
				return false
			}
		case West:
			if state.Y <= minDir-1 {
				return false
			}
		}
		return state.TimesDirUsed >= minDir
	}
}

func IsOppositeDirection(dir1, dir2 Direction) bool {
	switch dir1 {
	case North:
		return dir2 == South
	case South:
		return dir2 == North
	case East:
		return dir2 == West
	case West:
		return dir2 == East
	default:
		return false
	}
}

func IsInBounds(blocks [][]int, x, y int) bool {
	return x >= 0 && x < len(blocks) && y >= 0 && y < len(blocks[x])
}
