package four

import (
	"fmt"
	"strconv"
	"strings"

	puzzleinput "github.com/fwwieffering/aoc2022/internal/puzzle-input"
)

func parseSections(in string) ([]int64, []int64) {
	sections := strings.Split(in, ",")

	res := [][]int64{nil, nil}
	for i := 0; i < 2; i++ {
		splitSection := strings.Split(sections[i], "-")
		min, _ := strconv.ParseInt(splitSection[0], 10, 64)
		max, _ := strconv.ParseInt(splitSection[1], 10, 64)
		res[i] = []int64{min, max}
	}
	return res[0], res[1]
}

func sectionsFullyOverlap(a, b []int64) bool {
	return (a[0] >= b[0] && a[1] <= b[1]) || (b[0] >= a[0] && b[1] <= a[1])
}

func sectionsPartiallyOverlap(a, b []int64) bool {
	return (a[1] >= b[0] && a[1] <= b[1]) ||
		(b[1] >= a[0] && b[1] <= a[0]) ||
		(a[0] <= b[1] && a[0] >= b[0]) ||
		(b[0] <= a[1] && b[0] >= a[0])
}

func Solve() error {
	input := string(puzzleinput.Day(4))
	pairs := strings.Split(input, "\n")

	fullOverlaps := 0
	partialOverlaps := 0
	for _, p := range pairs {
		if len(p) > 0 {
			e1, e2 := parseSections(p)
			full := sectionsFullyOverlap(e1, e2)
			partial := sectionsPartiallyOverlap(e1, e2)
			if full {
				fullOverlaps++
			}
			if partial {
				partialOverlaps++
			}
		}
	}
	fmt.Printf("Part 1: number of sections that fully overlap %d\n", fullOverlaps)
	fmt.Printf("Part 2: number of sections that partially overlap %d\n", partialOverlaps)

	return nil
}
