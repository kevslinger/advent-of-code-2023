package day23

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay23Part1(t *testing.T) {
	file, err := os.Open("./data/day23_test.txt")
	if err != nil {
		t.Fatalf("Error opening File for Day 23 Part 1: %s\n", err)
	}
	defer file.Close()
	steps, err := Part1(file)
	if err != nil {
		t.Fatalf("Error processing Day 23 Part 1: %s\n", err)
	}
	expected := 94
	assert.Equal(t, expected, steps, "Expected %d but got %d\n", expected, steps)
}

func TestDay23Part2(t *testing.T) {
	file, err := os.Open("./data/day23_test.txt")
	if err != nil {
		t.Fatalf("Error opening File for Day 23 Part 2: %s\n", err)
	}
	defer file.Close()
	steps, err := Part2(file)
	if err != nil {
		t.Fatalf("Error processing Day 23 Part 2: %s\n", err)
	}
	expected := 154
	assert.Equal(t, expected, steps, "Expected %d but got %d\n", expected, steps)
}
