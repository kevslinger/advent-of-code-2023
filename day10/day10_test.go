package day10

import (
	"testing"

	"github.com/kevslinger/advent-of-code-2023/runner"
	"github.com/stretchr/testify/assert"
)

func TestDay10Part1(t *testing.T) {
	var testCases = []struct {
		TestName string
		Path     string
		Expected int
	}{
		{"simple", "./data/day10_test.txt", 4},
		{"complex", "./data/day10_test2.txt", 8},
	}
	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			steps, err := runner.RunPart(testCase.Path, part1)
			if err != nil {
				t.Errorf("Error with test %s: %s\n", testCase.TestName, err)
			} else {
				assert.Equal(t, testCase.Expected, steps, "Expected %d but got %d\n", testCase.Expected, steps)
			}
		})
	}
}

// func TestDay10Part2(t *testing.T) {
// 	var testCases = []struct {
// 		TestName string
// 		Path     string
// 		Expected int
// 	}{
// 		{"simple", "./data/day10_test3.txt", 4},
// 		{"complex", "./data/day10_test4.txt", 8},
// 	}
// 	for _, testCase := range testCases {
// 		t.Run(testCase.TestName, func(t *testing.T) {
// 			tiles, err := runner.RunPart(testCase.Path, Part2)
// 			if err != nil {
// 				t.Errorf("Error with test %s: %s\n", testCase.TestName, err)
// 			} else {
// 				assert.Equal(t, testCase.Expected, tiles, "Expected %d but got %d\n", testCase.Expected, tiles)
// 			}
// 		})
// 	}
// }
