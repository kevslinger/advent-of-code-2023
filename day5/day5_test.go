package day5

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay5Part1(t *testing.T) {
	file, err := os.Open("./data/day5_test.txt")
	if err != nil {
		t.Fatalf("Error reading file: %s", err)
	}
	locationNum, err := Part1(file)
	if err != nil {
		t.Fatalf("Error processing Day 5 Part 1: %s", err)
	}
	expected := 35
	assert.Equal(t, expected, locationNum, "Expected %d but got %d", expected, locationNum)
}
