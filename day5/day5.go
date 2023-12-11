package day5

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func RunDay5(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Cannot open file in Day 5 Part 1: %s\n", err)
		return
	}
	locationNum, err := Part1(file)
	if err != nil {
		fmt.Printf("Error with processing Day 5 Part 1: %s\n", err)
	} else {
		fmt.Printf("Answer to Day 5 Part 1 is: %d\n", locationNum)
	}

}

func Part1(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	toks := strings.Split(line, " ")
	seeds := make([]int, len(toks)-1)
	for idx, tok := range toks[1:] {
		seed, err := strconv.Atoi(tok)
		if err != nil {
			return -1, err
		}
		seeds[idx] = seed
	}

	intervals := make([]ByDest, 7)
	for i := 0; i < 7; i++ {
		scanner.Scan()
		m, err := ParseIntervals(scanner)
		if err != nil {
			return -1, err
		}
		intervals[i] = m
	}
	locationNum := -1
	for _, seed := range seeds {
		val := GetSeedLocation(seed, intervals)
		if val < locationNum || locationNum < 0 {
			locationNum = val
		}
	}

	return locationNum, nil
}

func Part2(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	toks := strings.Split(line, " ")
	seeds := make([][2]int, (len(toks)-1)/2)
	for idx, tok := range toks[1:] {
		if idx%2 == 1 {
			continue
		}
		seedStart, err := strconv.Atoi(tok)
		if err != nil {
			return -1, err
		}
		//fmt.Printf("SeedStart is %d\n", seedStart)
		seedEnd, err := strconv.Atoi(toks[idx+2])
		if err != nil {
			return -1, err
		}
		//fmt.Printf("SeedEnd is %d\n", seedEnd)
		seeds[idx/2] = [2]int{seedStart, seedEnd}
	}
	//fmt.Printf("Seeds: %#v\n", seeds)

	intervals := make([]ByDest, 7)
	for i := 0; i < 7; i++ {
		scanner.Scan()
		m, err := ParseIntervals(scanner)
		if err != nil {
			return -1, err
		}
		intervals[i] = m
	}
	locationNum := -1
	// TODO: This ends up in a loop that takes too long
	// We can be smarter about this by caching the pathways we've
	// seen before
	// And even smarter by computing the next pathway which will actually change anything...
	for _, seedRange := range seeds {
		for i := seedRange[0]; i < seedRange[0]+seedRange[1]; i++ {
			val := GetSeedLocation(i, intervals)
			//fmt.Printf("Val is %d locationNum is %d\n", val, locationNum)
			if val < locationNum || locationNum < 0 {
				locationNum = val
			}
		}

	}
	return locationNum, nil
}

type DstToTgt struct {
	Dest int
	Src  int
	Num  int
}

// Implementing sort interface
type ByDest []DstToTgt

func (bd ByDest) Len() int           { return len(bd) }
func (bd ByDest) Swap(i, j int)      { bd[i], bd[j] = bd[j], bd[i] }
func (bd ByDest) Less(i, j int) bool { return bd[i].Src < bd[j].Src }

func ParseIntervals(scanner *bufio.Scanner) (ByDest, error) {
	srcToTgt := make([]DstToTgt, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		_, err := strconv.Atoi(line[:1])
		if err != nil {
			continue
		}
		toks := strings.Split(line, " ")
		nums := make([]int, 3)
		for idx, tok := range toks {
			num, err := strconv.Atoi(tok)
			if err != nil {
				return nil, err
			}
			nums[idx] = num
		}
		srcToTgt = append(srcToTgt, DstToTgt{Dest: nums[0], Src: nums[1], Num: nums[2]})
	}
	sort.Sort(ByDest(srcToTgt))
	return srcToTgt, nil
}

func BinarySearch(interval ByDest, val int) int {
	leftPtr, rightPtr := 0, len(interval)-1
	for leftPtr <= rightPtr {
		midPt := (rightPtr + leftPtr) / 2
		if val < interval[midPt].Src {
			rightPtr = midPt - 1
		} else if val > interval[midPt].Src+interval[midPt].Num-1 {
			leftPtr = midPt + 1
		} else {
			// In range
			return interval[midPt].Dest + (val - interval[midPt].Src)
		}
	}
	return val
}

func GetSeedLocation(val int, intervals []ByDest) int {
	for _, interval := range intervals {
		val = BinarySearch(interval, val)
	}
	return val
}
