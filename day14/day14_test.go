package day14

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay14Part1(t *testing.T) {
	file, err := os.Open("./data/day14_test.txt")
	if err != nil {
		t.Fatalf("Error opening file in Day 14 Part 1: %s\n", err)
	}
	defer file.Close()

	load, err := Part1(file)
	if err != nil {
		t.Fatalf("Error processing Day 14 Part 1: %s\n", err)
	}
	expected := 136
	assert.Equal(t, expected, load, "Expected %d but got %d", expected, load)
}

// func TestDay14Part2(t *testing.T) {
// 	file, err := os.Open("./data/day14_test.txt")
// 	if err != nil {
// 		t.Fatalf("Error opening file in Day 14 Part 2: %s\n", err)
// 	}
// 	defer file.Close()

// 	load, err := Part2(file)
// 	if err != nil {
// 		t.Fatalf("Error processing Day 14 Part 2: %s\n", err)
// 	}
// 	expected := 64
// 	assert.Equal(t, expected, load, "Expected %d but got %d", expected, load)
// }
