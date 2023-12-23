package day21

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay21Part1(t *testing.T) {
	file, err := os.Open("./data/day21_test.txt")
	if err != nil {
		t.Fatalf("Error opening file in Day 21 Part 1: %s\n", err)
	}
	defer file.Close()
	plots, err := Part1(file, 6)
	if err != nil {
		t.Fatalf("Error processing Day 21 Part 1: %s\n", err)
	}
	expected := 16
	assert.Equal(t, expected, plots, "Expected %d but got %d", expected, plots)
}
