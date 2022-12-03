package three

import (
	"bytes"
	"fmt"

	puzzleinput "github.com/fwwieffering/aoc2022/internal/puzzle-input"
)

// convert byte to priority where
// Lowercase item types a through z have priorities 1 through 26.
// Uppercase item types A through Z have priorities 27 through 52.
// also tried implementing it as a map lookup but it was slower than this
func priority(b byte) int {
	// http://web.cecs.pdx.edu/~harry/compilers/ASCIIChart.pdf
	// normalized means A = 1 a = 33
	normalized := int(b) - 64
	isCap := normalized/32 == 0
	if isCap {
		return normalized + 26
	}
	return normalized % 32
}

func parseRussack(r []byte, potentialOverlaps map[string]int) (int, map[string]int) {
	mid := len(r) / 2

	pSet1 := make(map[byte]int)
	pset2 := make(map[byte]int)
	overlapPrio := 0

	newOverlaps := make(map[string]int)
	for idx, b := range r {
		// part 1 stuff
		if idx < mid {
			pSet1[b] = 0
		} else {
			if _, ok := pSet1[b]; ok {
				if _, in2 := pset2[b]; !in2 {
					overlapPrio += priority(b)
				}
			}
			pset2[b] = 0
		}
		// part 2 stuff
		if _, isOverlap := potentialOverlaps[string(b)]; isOverlap || len(potentialOverlaps) == 0 {
			newOverlaps[string(b)] = priority(b)
		}
	}
	return overlapPrio, newOverlaps
}

func Solve() error {
	input := puzzleinput.Day(3)
	sacks := bytes.Split(input, []byte("\n"))

	overlapPrio := 0
	groupPrio := 0
	groupidx := 0
	var groupOverlaps map[string]int
	for _, s := range sacks {
		var p int
		p, groupOverlaps = parseRussack(s, groupOverlaps)
		overlapPrio += p

		if groupidx == 2 {
			if len(groupOverlaps) != 1 {
				return fmt.Errorf("should only be one overlap but there were %+v", groupOverlaps)
			}
			for _, p := range groupOverlaps {
				groupPrio += p
			}
			groupOverlaps = nil
			groupidx = 0
		} else {
			groupidx++
		}
	}

	fmt.Printf("part 1: total priority %d\n", overlapPrio)
	fmt.Printf("part 2: per group overlap prio %d\n", groupPrio)
	return nil
}
