package day17

import (
	"os"
	"testing"

	"github.com/kevslinger/advent-of-code-2023/runner"
	"github.com/stretchr/testify/assert"
)

func TestDay17Part1(t *testing.T) {
	file, err := os.Open("./data/day17_test.txt")
	if err != nil {
		t.Fatalf("Error opening file for Day 17 Part 1: %s\n", err)
	}
	defer file.Close()

	heatLoss, err := Part1(file)
	if err != nil {
		t.Fatalf("Error processing Day 17 Part 1: %s\n", err)
	}
	expected := 102
	assert.Equal(t, expected, heatLoss, "Expected %d but got %d", expected, heatLoss)
}

func TestDay17Part2(t *testing.T) {
	var testCases = []struct {
		TestName string
		Path     string
		Expected int
	}{
		{"Ex1", "./data/day17_test.txt", 94},
		{"Ex2", "./data/day17_test2.txt", 71},
	}
	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			heatLoss, err := runner.RunPart(testCase.Path, Part2)
			if err != nil {
				t.Errorf("Error with test %s: %s\n", testCase.TestName, err)
			} else {
				assert.Equal(t, testCase.Expected, heatLoss, "Expected %d but got %d\n", testCase.Expected, heatLoss)
			}
		})
	}
}
