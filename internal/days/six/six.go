package six

import (
	"fmt"

	puzzleinput "github.com/fwwieffering/aoc2022/internal/puzzle-input"
)

type AlphaBitSet struct {
	bits uint64
	size int
}

func findMarker(in []byte, distinctChars int) int {
	size := len(in)

	currentChars := map[byte]int{}

	leftCursor := 0
	rightCursor := 0
	currentChars[in[0]]++

	for leftCursor < size && rightCursor <= size {
		// base case, we have the number of distinct chars we are looking for
		if len(currentChars) == distinctChars {
			// it wants the nth char which is index +1
			return rightCursor + 1
		}
		// move right cursor out and record char count
		if rightCursor-leftCursor < distinctChars-1 {
			rightCursor++
			currentChars[in[rightCursor]]++
		} else {
			// move left cursor in and remove char count
			currentChars[in[leftCursor]]--
			if currentChars[in[leftCursor]] < 1 {
				delete(currentChars, in[leftCursor])
			}
			leftCursor++
		}
	}

	return -1
}

func Solve() error {
	input := puzzleinput.Day(6)
	part1 := findMarker(input, 4)
	part2 := findMarker(input, 14)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
	return nil
}
