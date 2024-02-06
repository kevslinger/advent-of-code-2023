package day8

import (
	"bufio"
	"fmt"
	"io"
	"regexp"

	"github.com/kevslinger/advent-of-code-2023/runner"
)

func RunDay8(path string) {
	steps, err := runner.RunPart(path, part1)
	if err != nil {
		fmt.Printf("Error processing Day 8 Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 8 Part 1 is: %d\n", steps)
	}
}

func part1(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	directions := scanner.Text()

	nodeMap := parseNodeMap(scanner)
	curNode := nodeMap["AAA"]
	numSteps := 0
	for curNode.Name != "ZZZ" {
		dir := directions[numSteps%len(directions)]
		numSteps++
		if dir == 'L' {
			curNode = curNode.Left
		} else {
			curNode = curNode.Right
		}
	}
	return numSteps, nil
}

func parseNodeMap(scanner *bufio.Scanner) map[string]*node {
	nodeMatcher := regexp.MustCompile("[A-Z]{3}")
	nodeMap := make(map[string]*node, 0)
	for scanner.Scan() {
		// Skip area between directions and grid
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		nodeNames := nodeMatcher.FindAllString(line, -1)
		n, ok := nodeMap[nodeNames[0]]
		if !ok {
			n = &node{Name: nodeNames[0]}
			nodeMap[nodeNames[0]] = n
		}
		left, ok := nodeMap[nodeNames[1]]
		if !ok {
			left = &node{Name: nodeNames[1]}
			nodeMap[nodeNames[1]] = left
		}
		right, ok := nodeMap[nodeNames[2]]
		if !ok {
			right = &node{Name: nodeNames[2]}
			nodeMap[nodeNames[2]] = right
		}
		n.Left = left
		n.Right = right
	}
	return nodeMap
}

type node struct {
	Name  string
	Left  *node
	Right *node
}

func (n node) String() string {
	left := ""
	if n.Left != nil {
		left = n.Left.Name
	}
	right := ""
	if n.Right != nil {
		right = n.Right.Name
	}
	return fmt.Sprintf("%s, Left: %s, Right: %s", n.Name, left, right)
}
