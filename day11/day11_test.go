package day11

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay11Part1(t *testing.T) {
	file, err := os.Open("./data/day11_test.txt")
	if err != nil {
		t.Fatalf("Cannot open file for Day 11 Part 1: %s\n", err)
	}
	sum, err := Part1(file)
	if err != nil {
		t.Fatalf("Error procesing Day 11 Part 1: %s\n", err)
	}
	expected := 374
	assert.Equal(t, expected, sum, "Expected %d but got %d", expected, sum)
}

func TestDay11Part2(t *testing.T) {
	path := "./data/day11_test.txt"
	var testCases = []struct {
		TestName      string
		EmptyDistance int
		Expected      int
	}{
		{"10x", 10, 1030},
		{"100x", 100, 8410},
	}
	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			file, err := os.Open(path)
			if err != nil {
				t.Errorf("Error with test %s: %s\n", testCase.TestName, err)
			} else {
				steps := Part2(file, testCase.EmptyDistance)
				assert.Equal(t, testCase.Expected, steps, "Expected %d but got %d\n", testCase.Expected, steps)
			}
		})
	}
}
