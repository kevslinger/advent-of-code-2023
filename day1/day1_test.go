package day1

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1Part1(t *testing.T) {
	file, err := os.Open("./data/day1_test.txt")
	if err != nil {
		t.Fatalf("Cannot open file in Day 1 Part 1: %s", err)
	}
	sum, err := Part1(file)
	if err != nil {
		t.Fatalf("Error processing part 1: %s", err)
	}
	expected := 142
	assert.Equal(t, expected, sum, "Sum is %d but should be %d", sum, expected)
}

func TestDay1Part2(t *testing.T) {
	file, err := os.Open("./data/day1_part2_test.txt")
	if err != nil {
		t.Fatalf("Cannot open file in Day 1 Part 2: %s", err)
	}
	sum, err := Part2(file)
	if err != nil {
		log.Fatalf("Error processing part 2: %s", err)
	}
	expected := 281
	assert.Equal(t, expected, sum, "Sum is %d but should be %d", sum, expected)
}
