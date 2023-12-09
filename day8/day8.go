package day8

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func RunDay8(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Cannot open file in Day 8 Part 1: %s\n", err)
	}
	defer file.Close()

	steps, err := Part1(file)
	if err != nil {
		fmt.Printf("Error processing Day 8 Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 8 Part 1 is: %d\n", steps)
	}
}

func Part1(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	directions := scanner.Text()

	nodeMap := ParseNodeMap(scanner)
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

func ParseNodeMap(scanner *bufio.Scanner) map[string]*Node {
	nodeMatcher := regexp.MustCompile("[A-Z]{3}")
	nodeMap := make(map[string]*Node, 0)
	for scanner.Scan() {
		// Skip area between directions and grid
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		nodeNames := nodeMatcher.FindAllString(line, -1)
		node, ok := nodeMap[nodeNames[0]]
		if !ok {
			node = &Node{Name: nodeNames[0]}
			nodeMap[nodeNames[0]] = node
		}
		left, ok := nodeMap[nodeNames[1]]
		if !ok {
			left = &Node{Name: nodeNames[1]}
			nodeMap[nodeNames[1]] = left
		}
		right, ok := nodeMap[nodeNames[2]]
		if !ok {
			right = &Node{Name: nodeNames[2]}
			nodeMap[nodeNames[2]] = right
		}
		node.Left = left
		node.Right = right
	}
	return nodeMap
}

type Node struct {
	Name  string
	Left  *Node
	Right *Node
}

func (n Node) String() string {
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
