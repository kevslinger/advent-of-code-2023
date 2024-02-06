package day16

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kevslinger/advent-of-code-2023/runner"
)

func RunDay16(path string) {
	tiles, err := runner.RunPart(path, part1)
	if err != nil {
		fmt.Printf("Error in Day 16 Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 16 Part 1 is: %d\n", tiles)
	}
	tiles, err = runner.RunPart(path, part2)
	if err != nil {
		fmt.Printf("Error in Day 16 Part 2: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 16 Part 2 is: %d\n", tiles)
	}
}

type direction int

const (
	north direction = iota
	south
	east
	west
)

var dirMap map[direction][]int = map[direction][]int{
	north: {-1, 0},
	south: {1, 0},
	east:  {0, 1},
	west:  {0, -1},
}

type position struct {
	X int
	Y int
}

type beam struct {
	pos position
	dir direction
}

func part1(file io.Reader) (int, error) {
	var cave []string = parseCave(file)
	var energizedTileMap [][]int = simulateLight(cave, beam{pos: position{X: 0, Y: 0}, dir: east})
	return countEnergizedTiles(energizedTileMap), nil
}

func part2(file io.Reader) (int, error) {
	var cave []string = parseCave(file)
	var highestEnergizedTiles int
	// Can start from any outside edge pointing in
	var energizedTileMap [][]int
	var energizedTiles int
	for idx := 0; idx < len(cave); idx++ {
		// east
		energizedTileMap = simulateLight(cave, beam{pos: position{X: idx, Y: 0}, dir: east})
		energizedTiles = countEnergizedTiles(energizedTileMap)
		if energizedTiles > highestEnergizedTiles {
			highestEnergizedTiles = energizedTiles
		}
		// west
		energizedTileMap = simulateLight(cave, beam{pos: position{X: idx, Y: len(cave[idx]) - 1}, dir: west})
		energizedTiles = countEnergizedTiles(energizedTileMap)
		if energizedTiles > highestEnergizedTiles {
			highestEnergizedTiles = energizedTiles
		}
	}
	for idx := 0; idx < len(cave[0]); idx++ {
		// south
		energizedTileMap = simulateLight(cave, beam{pos: position{X: 0, Y: idx}, dir: south})
		energizedTiles = countEnergizedTiles(energizedTileMap)
		if energizedTiles > highestEnergizedTiles {
			highestEnergizedTiles = energizedTiles
		}
		// north
		energizedTileMap = simulateLight(cave, beam{pos: position{X: len(cave) - 1, Y: idx}, dir: north})
		energizedTiles = countEnergizedTiles(energizedTileMap)
		if energizedTiles > highestEnergizedTiles {
			highestEnergizedTiles = energizedTiles
		}
	}

	return highestEnergizedTiles, nil
}

func parseCave(file io.Reader) []string {
	scanner := bufio.NewScanner(file)
	cave := make([]string, 0)
	for scanner.Scan() {
		cave = append(cave, scanner.Text())
	}
	return cave
}

func simulateLight(cave []string, startingbeam beam) [][]int {
	energizedTileMap := makeTileMap(cave)
	beamCache := make(map[beam]bool)
	queue := make([]beam, 0)
	queue = append(queue, startingbeam)
	for len(queue) > 0 {
		curbeam := queue[0]
		// Pop
		queue = queue[1:]
		// Prevent simulating duplicate beams
		_, ok := beamCache[curbeam]
		if !isInBounds(curbeam.pos, cave) || ok {
			continue
		}
		beamCache[curbeam] = true
		energizedTileMap[curbeam.pos.X][curbeam.pos.Y]++
		// Process next step
		queue = append(queue, calculateNextBeams(curbeam, cave)...)

	}
	return energizedTileMap
}

func makeTileMap(cave []string) [][]int {
	tileMap := make([][]int, len(cave))
	for idx, row := range cave {
		tileMap[idx] = make([]int, len(row))
	}
	return tileMap
}

func calculateNextBeams(curbeam beam, cave []string) []beam {
	nextbeams := make([]beam, 0)
	switch cave[curbeam.pos.X][curbeam.pos.Y] {
	case '/':
		switch curbeam.dir {
		case north:
			nextbeams = append(nextbeams, stepBeam(curbeam, east))
		case south:
			nextbeams = append(nextbeams, stepBeam(curbeam, west))
		case east:
			nextbeams = append(nextbeams, stepBeam(curbeam, north))
		case west:
			nextbeams = append(nextbeams, stepBeam(curbeam, south))
		}
	case '\\':
		switch curbeam.dir {
		case north:
			nextbeams = append(nextbeams, stepBeam(curbeam, west))
		case south:
			nextbeams = append(nextbeams, stepBeam(curbeam, east))
		case east:
			nextbeams = append(nextbeams, stepBeam(curbeam, south))
		case west:
			nextbeams = append(nextbeams, stepBeam(curbeam, north))
		}
	case '-':
		if curbeam.dir == west || curbeam.dir == east {
			nextbeams = append(nextbeams, stepBeam(curbeam, curbeam.dir))
		} else {
			nextbeams = append(nextbeams, stepBeam(curbeam, west))
			nextbeams = append(nextbeams, stepBeam(curbeam, east))
		}
	case '|':
		if curbeam.dir == north || curbeam.dir == south {
			nextbeams = append(nextbeams, stepBeam(curbeam, curbeam.dir))
		} else {
			nextbeams = append(nextbeams, stepBeam(curbeam, north))
			nextbeams = append(nextbeams, stepBeam(curbeam, south))
		}
	default:
		nextbeams = append(nextbeams, stepBeam(curbeam, curbeam.dir))
	}
	return nextbeams
}

func stepBeam(curbeam beam, dir direction) beam {
	change := dirMap[dir]
	newpos := position{curbeam.pos.X + change[0], curbeam.pos.Y + change[1]}
	return beam{pos: newpos, dir: dir}
}

func countEnergizedTiles(cave [][]int) int {
	var energizedTiles int
	for _, row := range cave {
		for _, cell := range row {
			if cell > 0 {
				energizedTiles++
			}
		}
	}
	return energizedTiles
}

func isInBounds(pos position, cave []string) bool {
	return pos.X >= 0 && pos.X < len(cave) && pos.Y >= 0 && pos.Y < len(cave[pos.X])
}
