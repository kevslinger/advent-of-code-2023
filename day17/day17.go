package day17

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"github.com/kevslinger/advent-of-code-2023/runner"
)

func Runday17(path string) {
	heatLoss, err := runner.RunPart(path, part1)
	if err != nil {
		fmt.Printf("Error with Day 17 Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 17 Part 1 is: %d\n", heatLoss)
	}
	heatLoss, err = runner.RunPart(path, part2)
	if err != nil {
		fmt.Printf("Error with Day 17 Part 2: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 17 Part 2 is: %d\n", heatLoss)
	}
}

type direction int

const (
	dummy direction = iota
	north
	south
	east
	west
)

var directions = []direction{north, south, east, west}

var dirMap map[direction][]int = map[direction][]int{
	north: {-1, 0},
	south: {1, 0},
	east:  {0, 1},
	west:  {0, -1},
}

type traversalState struct {
	x            int
	y            int
	heatLost     int
	dir          direction
	timesDirUsed int
}

func part1(file io.Reader) (int, error) {
	blocks, err := parseCityBlocks(file)
	if err != nil {
		return -1, err
	}
	return calculateHeatLost(blocks, 0, 3), nil
}

func part2(file io.Reader) (int, error) {
	blocks, err := parseCityBlocks(file)
	if err != nil {
		return -1, err
	}
	return calculateHeatLost(blocks, 4, 10), nil
}

func parseCityBlocks(file io.Reader) ([][]int, error) {
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

func calculateHeatLost(blocks [][]int, minMovement int, maxMovement int) int {
	var heatLossMap [][]int = traverseCityBlocks(blocks, minMovement, maxMovement)
	// Starting block shouldn't count
	return heatLossMap[0][0] - blocks[0][0]
}

func traverseCityBlocks(blocks [][]int, minMovements int, maxMovements int) [][]int {
	var heatLossMap [][]int = createHeatLossMap(blocks)
	heatTracker := createHeatMap(blocks)
	// Create a queue to hold the traversal states
	var queue []traversalState = make([]traversalState, 0)
	// Start at the bottom corner and work our way up
	queue = append(queue, traversalState{x: len(heatLossMap) - 1, y: len(heatLossMap[0]) - 1, heatLost: blocks[len(blocks)-1][len(blocks[0])-1], dir: dummy, timesDirUsed: maxMovements})
	for len(queue) > 0 {
		curState := queue[0]
		// Pop
		queue = queue[1:]
		// Don't continue once we reach the start
		if curState.x == 0 && curState.y == 0 {
			continue
		}
		// Check all directions
		for _, dir := range directions {
			newState := processState(blocks, heatLossMap, heatTracker, curState, dir, minMovements, maxMovements)
			// Use heatLost as a pseudo nil checker since heatLost is monotonically increasing
			if newState.heatLost > 0 {
				queue = append(queue, newState)
			}
		}
	}
	return heatLossMap
}

func createHeatLossMap(blocks [][]int) [][]int {
	var heatLossMap [][]int = make([][]int, len(blocks))
	for i := 0; i < len(heatLossMap); i++ {
		heatLossMap[i] = make([]int, len(blocks[i]))
	}
	// Initialise the bottom right with the correct value
	heatLossMap[len(heatLossMap)-1][len(heatLossMap[0])-1] = blocks[len(heatLossMap)-1][len(heatLossMap[0])-1]
	return heatLossMap
}

func createHeatMap(blocks [][]int) [][]map[direction][]int {
	var heatLossMap [][]map[direction][]int = make([][]map[direction][]int, len(blocks))
	for i := 0; i < len(heatLossMap); i++ {
		heatLossMap[i] = make([]map[direction][]int, len(blocks[i]))
		for j := 0; j < len(heatLossMap[i]); j++ {
			heatLossMap[i][j] = make(map[direction][]int)
			for _, dir := range directions {
				heatLossMap[i][j][dir] = make([]int, 11)
			}
		}
	}
	return heatLossMap
}

func processState(blocks [][]int, heatLossMap [][]int, heatTracker [][]map[direction][]int, state traversalState, dir direction, minMovements int, maxMovements int) traversalState {
	// The crucible has restrictions on which way it can turn
	if !isValidDirection(state, dir, blocks, minMovements, maxMovements) {
		return traversalState{}
	}

	// Continue in same direction
	newX := state.x + dirMap[dir][0]
	newY := state.y + dirMap[dir][1]
	if !isInBounds(blocks, newX, newY) {
		return traversalState{}
	}
	var timesDirUsed int
	if state.dir == dir {
		timesDirUsed = state.timesDirUsed + 1
	} else {
		timesDirUsed = 1
	}
	// If going in this direction would create a new best score, continue doing it
	newHeatLoss := state.heatLost + blocks[newX][newY]
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
			return traversalState{x: newX, y: newY, heatLost: newHeatLoss, dir: dir, timesDirUsed: timesDirUsed}
		}
	}
	return traversalState{}
}

func isValidDirection(state traversalState, nextdir direction, board [][]int, mindir int, maxdir int) bool {
	if isOppositeDirection(state.dir, nextdir) {
		return false
	}
	if nextdir == state.dir {
		return state.timesDirUsed < maxdir
	} else {
		switch nextdir {
		case north:
			if state.x <= mindir-1 {
				return false
			}
		case south:
			if state.x >= len(board)-mindir {
				return false
			}
		case east:
			if state.y >= len(board[0])-mindir {
				return false
			}
		case west:
			if state.y <= mindir-1 {
				return false
			}
		}
		return state.timesDirUsed >= mindir
	}
}

func isOppositeDirection(dir1, dir2 direction) bool {
	switch dir1 {
	case north:
		return dir2 == south
	case south:
		return dir2 == north
	case east:
		return dir2 == west
	case west:
		return dir2 == east
	default:
		return false
	}
}

func isInBounds(blocks [][]int, x, y int) bool {
	return x >= 0 && x < len(blocks) && y >= 0 && y < len(blocks[x])
}
