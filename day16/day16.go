package day16

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kevslinger/advent-of-code-2023/runner"
)

func RunDay16(path string) {
	tiles, err := runner.RunPart(path, Part1)
	if err != nil {
		fmt.Printf("Error in Day 16 Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 16 Part 1 is: %d\n", tiles)
	}
	tiles, err = runner.RunPart(path, Part2)
	if err != nil {
		fmt.Printf("Error in Day 16 Part 2: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 16 Part 2 is: %d\n", tiles)
	}
}

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

var dirMap map[Direction][]int = map[Direction][]int{
	North: {-1, 0},
	South: {1, 0},
	East:  {0, 1},
	West:  {0, -1},
}

type Position struct {
	X int
	Y int
}

type Beam struct {
	Pos Position
	Dir Direction
}

func Part1(file io.Reader) (int, error) {
	var cave []string = ParseCave(file)
	var energizedTileMap [][]int = SimulateLight(cave, Beam{Pos: Position{X: 0, Y: 0}, Dir: East})
	return CountEnergizedTiles(energizedTileMap), nil
}

func Part2(file io.Reader) (int, error) {
	var cave []string = ParseCave(file)
	var highestEnergizedTiles int
	// Can start from any outside edge pointing in
	var energizedTileMap [][]int
	var energizedTiles int
	for idx := 0; idx < len(cave); idx++ {
		// EAST
		energizedTileMap = SimulateLight(cave, Beam{Pos: Position{X: idx, Y: 0}, Dir: East})
		energizedTiles = CountEnergizedTiles(energizedTileMap)
		if energizedTiles > highestEnergizedTiles {
			highestEnergizedTiles = energizedTiles
		}
		// WEST
		energizedTileMap = SimulateLight(cave, Beam{Pos: Position{X: idx, Y: len(cave[idx]) - 1}, Dir: West})
		energizedTiles = CountEnergizedTiles(energizedTileMap)
		if energizedTiles > highestEnergizedTiles {
			highestEnergizedTiles = energizedTiles
		}
	}
	for idx := 0; idx < len(cave[0]); idx++ {
		// SOUTH
		energizedTileMap = SimulateLight(cave, Beam{Pos: Position{X: 0, Y: idx}, Dir: South})
		energizedTiles = CountEnergizedTiles(energizedTileMap)
		if energizedTiles > highestEnergizedTiles {
			highestEnergizedTiles = energizedTiles
		}
		// NORTH
		energizedTileMap = SimulateLight(cave, Beam{Pos: Position{X: len(cave) - 1, Y: idx}, Dir: North})
		energizedTiles = CountEnergizedTiles(energizedTileMap)
		if energizedTiles > highestEnergizedTiles {
			highestEnergizedTiles = energizedTiles
		}
	}

	return highestEnergizedTiles, nil
}

func ParseCave(file io.Reader) []string {
	scanner := bufio.NewScanner(file)
	cave := make([]string, 0)
	for scanner.Scan() {
		cave = append(cave, scanner.Text())
	}
	return cave
}

func SimulateLight(cave []string, startingBeam Beam) [][]int {
	energizedTileMap := MakeTileMap(cave)
	beamCache := make(map[Beam]bool)
	queue := make([]Beam, 0)
	queue = append(queue, startingBeam)
	for len(queue) > 0 {
		curBeam := queue[0]
		// Pop
		queue = queue[1:]
		// Prevent simulating duplicate beams
		_, ok := beamCache[curBeam]
		if !IsInBounds(curBeam.Pos, cave) || ok {
			continue
		}
		beamCache[curBeam] = true
		energizedTileMap[curBeam.Pos.X][curBeam.Pos.Y]++
		// Process next step
		queue = append(queue, CalculateNextBeams(curBeam, cave)...)

	}
	return energizedTileMap
}

func MakeTileMap(cave []string) [][]int {
	tileMap := make([][]int, len(cave))
	for idx, row := range cave {
		tileMap[idx] = make([]int, len(row))
	}
	return tileMap
}

func CalculateNextBeams(curBeam Beam, cave []string) []Beam {
	nextBeams := make([]Beam, 0)
	switch cave[curBeam.Pos.X][curBeam.Pos.Y] {
	case '/':
		switch curBeam.Dir {
		case North:
			nextBeams = append(nextBeams, StepBeam(curBeam, East))
		case South:
			nextBeams = append(nextBeams, StepBeam(curBeam, West))
		case East:
			nextBeams = append(nextBeams, StepBeam(curBeam, North))
		case West:
			nextBeams = append(nextBeams, StepBeam(curBeam, South))
		}
	case '\\':
		switch curBeam.Dir {
		case North:
			nextBeams = append(nextBeams, StepBeam(curBeam, West))
		case South:
			nextBeams = append(nextBeams, StepBeam(curBeam, East))
		case East:
			nextBeams = append(nextBeams, StepBeam(curBeam, South))
		case West:
			nextBeams = append(nextBeams, StepBeam(curBeam, North))
		}
	case '-':
		if curBeam.Dir == West || curBeam.Dir == East {
			nextBeams = append(nextBeams, StepBeam(curBeam, curBeam.Dir))
		} else {
			nextBeams = append(nextBeams, StepBeam(curBeam, West))
			nextBeams = append(nextBeams, StepBeam(curBeam, East))
		}
	case '|':
		if curBeam.Dir == North || curBeam.Dir == South {
			nextBeams = append(nextBeams, StepBeam(curBeam, curBeam.Dir))
		} else {
			nextBeams = append(nextBeams, StepBeam(curBeam, North))
			nextBeams = append(nextBeams, StepBeam(curBeam, South))
		}
	default:
		nextBeams = append(nextBeams, StepBeam(curBeam, curBeam.Dir))
	}
	return nextBeams
}

func StepBeam(curBeam Beam, dir Direction) Beam {
	change := dirMap[dir]
	newPos := Position{curBeam.Pos.X + change[0], curBeam.Pos.Y + change[1]}
	return Beam{Pos: newPos, Dir: dir}
}

func CountEnergizedTiles(cave [][]int) int {
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

func IsInBounds(pos Position, cave []string) bool {
	return pos.X >= 0 && pos.X < len(cave) && pos.Y >= 0 && pos.Y < len(cave[pos.X])
}
