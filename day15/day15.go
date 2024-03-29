package day15

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/kevslinger/advent-of-code-2023/runner"
)

func RunDay15(path string) {
	sum, err := runner.RunPart(path, part1)
	if err != nil {
		fmt.Printf("Error with Day 15 Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 15 Part 1 is: %d\n", sum)
	}
	power, err := runner.RunPart(path, part2)
	if err != nil {
		fmt.Printf("Error with Day 15 Part 2: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 15 Part 2 is: %d\n", power)
	}
}

type node struct {
	Val  Step
	Next *node
	Prev *node
}

func (n *node) Remove() {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
}

// Add a node to the end of the list
func (n *node) Add(newNode *node) {
	newNode.Prev = n.Prev
	n.Prev.Next = newNode
	n.Prev = newNode
	newNode.Next = n
}

func newNode() *node {
	return &node{}
}

type Step struct {
	Label       string
	Operation   rune
	FocalLength int
}

func part1(file io.Reader) (int, error) {
	var steps []string = parseInitializationSequence(file)
	var sum int
	for _, step := range steps {
		sum += hashInstruction(step)
	}
	return sum, nil
}

func part2(file io.Reader) (int, error) {
	var stepStrings []string = parseInitializationSequence(file)
	steps, err := parseSteps(stepStrings)
	if err != nil {
		return -1, err
	}
	return calculatePowerFromSteps(steps), nil
}

func parseInitializationSequence(file io.Reader) []string {
	scanner := bufio.NewScanner(file)
	var steps []string = make([]string, 0)
	scanner.Scan()
	steps = strings.Split(scanner.Text(), ",")
	return steps
}

func hashInstruction(step string) int {
	var value int
	for _, r := range step {
		value += int(r)
		value *= 17
		value %= 256
	}
	return value
}

func parseSteps(steps []string) ([]Step, error) {
	var stepArr []Step = make([]Step, len(steps))
	for idx, step := range steps {
		switch step[len(step)-1] {
		case '-':
			stepArr[idx] = Step{Label: step[:len(step)-1], Operation: '-'}
		// Catches "="
		default:
			strToks := strings.Split(step, "=")
			focalLength, err := strconv.Atoi(strToks[1])
			if err != nil {
				return make([]Step, 0), err
			}
			stepArr[idx] = Step{Label: strToks[0], Operation: '=', FocalLength: focalLength}
		}
	}
	return stepArr, nil
}

func calculatePowerFromSteps(steps []Step) int {
	var labelMap map[string]*node = make(map[string]*node)
	var boxMap map[int]*node = make(map[int]*node)
	for i := 0; i < 256; i++ {
		boxMap[i] = newNode()
		boxMap[i].Next = boxMap[i]
		boxMap[i].Prev = boxMap[i]
	}
	for _, step := range steps {
		n, ok := labelMap[step.Label]
		// Switch based on instruction
		switch step.Operation {
		case '-':
			if !ok {
				continue
			}
			labelMap[step.Label].Remove()
			delete(labelMap, step.Label)
		case '=':
			if ok {
				n.Val.FocalLength = step.FocalLength
			} else {
				newNode := node{Val: step}
				boxMap[hashInstruction(step.Label)].Add(&newNode)
				labelMap[step.Label] = &newNode
			}
		}
	}
	return calculateFocusingPower(boxMap)
}

func calculateFocusingPower(boxMap map[int]*node) int {
	var focusingPower int
	for i := 0; i < 256; i++ {
		// This will point to a dummy node
		node := boxMap[i]
		node = node.Next
		var slot int
		if node.Val.Label == "" {
			continue
		}
		for node.Val.Label != "" {
			focusingPower += (1 + i) * (1 + slot) * node.Val.FocalLength
			slot++
			node = node.Next
		}
	}
	return focusingPower
}
