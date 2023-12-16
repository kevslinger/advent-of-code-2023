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
	sum, err := runner.RunPart(path, Part1)
	if err != nil {
		fmt.Printf("Error with Day 15 Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 15 Part 1 is: %d\n", sum)
	}
	power, err := runner.RunPart(path, Part2)
	if err != nil {
		fmt.Printf("Error with Day 15 Part 2: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 15 Part 2 is: %d\n", power)
	}
}

type Node struct {
	Val  Step
	Next *Node
	Prev *Node
}

func (n *Node) Remove() {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
}

// Add a node to the end of the list
func (n *Node) Add(newNode *Node) {
	newNode.Prev = n.Prev
	n.Prev.Next = newNode
	n.Prev = newNode
	newNode.Next = n
}

func NewNode() *Node {
	return &Node{}
}

type Step struct {
	Label       string
	Operation   rune
	FocalLength int
}

func Part1(file io.Reader) (int, error) {
	var steps []string = ParseInitializationSequence(file)
	var sum int
	for _, step := range steps {
		sum += HashInstruction(step)
	}
	return sum, nil
}

func Part2(file io.Reader) (int, error) {
	var stepStrings []string = ParseInitializationSequence(file)
	steps, err := ParseSteps(stepStrings)
	if err != nil {
		return -1, err
	}
	return CalculatePowerFromSteps(steps), nil
}

func ParseInitializationSequence(file io.Reader) []string {
	scanner := bufio.NewScanner(file)
	var steps []string = make([]string, 0)
	scanner.Scan()
	steps = strings.Split(scanner.Text(), ",")
	return steps
}

func HashInstruction(step string) int {
	var value int
	for _, r := range step {
		value += int(r)
		value *= 17
		value %= 256
	}
	return value
}

func ParseSteps(steps []string) ([]Step, error) {
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

func CalculatePowerFromSteps(steps []Step) int {
	var labelMap map[string]*Node = make(map[string]*Node)
	var boxMap map[int]*Node = make(map[int]*Node)
	for i := 0; i < 256; i++ {
		boxMap[i] = NewNode()
		boxMap[i].Next = boxMap[i]
		boxMap[i].Prev = boxMap[i]
	}
	for _, step := range steps {
		node, ok := labelMap[step.Label]
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
				node.Val.FocalLength = step.FocalLength
			} else {
				newNode := Node{Val: step}
				boxMap[HashInstruction(step.Label)].Add(&newNode)
				labelMap[step.Label] = &newNode
			}
		}
	}
	return CalculateFocusingPower(boxMap)
}

func CalculateFocusingPower(boxMap map[int]*Node) int {
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
