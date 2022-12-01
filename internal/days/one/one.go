package one

import (
	"bytes"
	"fmt"
	"strconv"

	puzzleinput "github.com/fwwieffering/aoc2022/internal/puzzle-input"
)

func addToSortedArray(curArr []int64, newItem int64) []int64 {
	// pop off last item of list
	curIdx := 0
	finalIdx := 0
	var final = make([]int64, len(curArr))
	for finalIdx < len(curArr) {
		if newItem > curArr[curIdx] && finalIdx == curIdx {
			final[finalIdx] = newItem
			finalIdx++
		} else {
			final[finalIdx] = curArr[curIdx]
			finalIdx++
			curIdx++
		}
	}
	return final
}

func getMaxElves(in [][]byte, numMaxElves int) []int64 {
	maxElves := make([]int64, numMaxElves)

	var curElf int64 = 0

	for _, line := range in {
		sl := string(line)
		if sl == "" {
			if maxElves[numMaxElves-1] <= curElf {
				maxElves = addToSortedArray(maxElves, curElf)
			}
			curElf = 0
		} else {
			cals, _ := strconv.ParseInt(sl, 10, 32)
			curElf += cals
		}
	}
	// trailing uncompleted elf
	if curElf != 0 {
		if maxElves[numMaxElves-1] <= curElf {
			maxElves = addToSortedArray(maxElves, curElf)
		}
	}
	return maxElves
}

func Solve() error {
	input := puzzleinput.Day(1)

	// split bytes by line
	lines := bytes.Split(input, []byte("\n"))
	maxElves := getMaxElves(lines, 3)
	fmt.Printf("Part 1: elf with most calories has %d calories\n", maxElves[0])
	var sum int64 = 0
	for _, elf := range maxElves {
		sum += elf
	}
	fmt.Printf("Part 2: top three elves %d\n", sum)
	return nil
}
