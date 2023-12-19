package day19

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay19Part1(t *testing.T) {
	file, err := os.Open("./data/day19_test.txt")
	if err != nil {
		t.Fatalf("Error opening file in Day 19 Part 1: %s\n", err)
	}
	defer file.Close()

	rating, err := Part1(file)
	if err != nil {
		t.Fatalf("Error processing Day 19 Part 1: %s\n", err)
	}
	expected := 19114
	assert.Equal(t, expected, rating, "Expected %d but got %d\n", expected, rating)
}
