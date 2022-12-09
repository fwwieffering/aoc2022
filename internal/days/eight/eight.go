package eight

import (
	"bytes"
	"fmt"
	"strconv"

	puzzleinput "github.com/fwwieffering/aoc2022/internal/puzzle-input"
)

type direction int

const (
	down direction = iota + 1
	up
	left
	right
	invalid
)

type queueItem struct {
	row int
	col int
	vec func(q *queueItem) *queueItem
}

func vRight(q *queueItem) *queueItem {
	return &queueItem{
		row: q.row,
		col: q.col + 1,
		vec: vRight,
	}
}

func vLeft(q *queueItem) *queueItem {
	return &queueItem{
		row: q.row,
		col: q.col - 1,
		vec: vLeft,
	}
}

func vUp(q *queueItem) *queueItem {
	return &queueItem{
		row: q.row - 1,
		col: q.col,
		vec: vUp,
	}
}

func vDown(q *queueItem) *queueItem {
	return &queueItem{
		row: q.row + 1,
		col: q.col,
		vec: vDown,
	}
}

// this needs some memoization or something its so slow
func someSortOfBfs(items [][]int, startrow int, startcol int) (bool, int) {
	startVal := items[startrow][startcol]
	// special case - we are the edge
	if startrow == 0 || startcol == 0 || startrow == len(items)-1 || startcol == len(items[0])-1 {
		return true, 1
	}
	queue := []*queueItem{
		{
			row: startrow,
			col: startcol + 1,
			vec: vRight,
		},
		{
			row: startrow,
			col: startcol - 1,
			vec: vLeft,
		},
		{
			row: startrow - 1,
			col: startcol,
			vec: vUp,
		},
		{
			row: startrow + 1,
			col: startcol,
			vec: vDown,
		},
	}
	isVisible := false
	visibleScore := 1

	for len(queue) > 0 {
		item := queue[0]
		if len(queue) > 1 {
			queue = queue[1:]
		} else {
			queue = []*queueItem{}
		}
		// if we are on the grid
		if item.row >= 0 || item.col >= 0 || item.row < len(items) || item.col < len(items[0]) {
			var distance int
			if item.row-startrow != 0 {
				distance = item.row - startrow
			} else if item.col-startcol != 0 {
				distance = item.col - startcol
			}
			if distance < 0 {
				distance = distance * -1
			}

			// if we are on an edge, we are visible
			if item.row == 0 || item.col == 0 || item.row == len(items)-1 || item.col == len(items[0])-1 {
				if items[item.row][item.col] < startVal {
					isVisible = true
					visibleScore *= distance
				}
			} else if items[item.row][item.col] < startVal {
				queue = append(queue, item.vec(item))
			} else {
				visibleScore *= distance
			}

		}
	}

	return isVisible, visibleScore
}

// its probably way more efficient to do some sort of sorting algorithm but I could not figure it out
// so instead this iterates over it in 4n where n = num items
func countVisible(visibleItems map[string]int, items [][]int, d direction) map[string]int {
	var outerLoopStart int
	var outerLoopCond func(i int) bool
	var outerLoopFxn func(i int) int
	var innerLoopStart int
	var innerLoopCond func(j int) bool
	var innerLoopFxn func(j int) int
	switch d {
	case left:
		outerLoopStart = 0
		outerLoopCond = func(i int) bool { return i < len(items) }
		outerLoopFxn = func(i int) int { return i + 1 }

		innerLoopStart = 0
		innerLoopCond = func(j int) bool { return j < len(items[0]) }
		innerLoopFxn = func(j int) int { return j + 1 }

	case right:
		outerLoopStart = 0
		outerLoopCond = func(i int) bool { return i < len(items) }
		outerLoopFxn = func(i int) int { return i + 1 }

		innerLoopStart = len(items[0]) - 1
		innerLoopCond = func(j int) bool { return j >= 0 }
		innerLoopFxn = func(j int) int { return j - 1 }
	case down:
		outerLoopStart = 0
		outerLoopCond = func(i int) bool { return i < len(items[0]) }
		outerLoopFxn = func(i int) int { return i + 1 }

		innerLoopStart = 0
		innerLoopCond = func(j int) bool { return j < len(items) }
		innerLoopFxn = func(j int) int { return j + 1 }
	case up:
		outerLoopStart = 0
		outerLoopCond = func(i int) bool { return i < len(items[0]) }
		outerLoopFxn = func(i int) int { return i + 1 }

		innerLoopStart = len(items) - 1
		innerLoopCond = func(j int) bool { return j >= 0 }
		innerLoopFxn = func(j int) int { return j - 1 }

	}

	for i := outerLoopStart; outerLoopCond(i); i = outerLoopFxn(i) {
		var runningMax = -1
		for j := innerLoopStart; innerLoopCond(j); j = innerLoopFxn(j) {
			var score int
			var item int
			var idx string
			if d == down || d == up {
				score = j
				item = items[j][i]
				idx = fmt.Sprintf("%d,%d", j, i)
			} else {
				score = i
				item = items[i][j]
				idx = fmt.Sprintf("%d,%d", i, j)
			}
			if j == innerLoopStart || !innerLoopCond(j+1) || item > runningMax {
				visibleItems[idx] = score

				if item > runningMax {
					visibleItems[idx] = score
					runningMax = item
				}
			}
		}
	}
	return visibleItems
}

func countVisibleTrees(items [][]int) (int, int) {
	// visible := make(map[string]int)

	// for _, d := range []direction{left, down, right, up} {
	// 	visible = countVisible(visible, items, d)
	// }

	// return len(visible)
	visCount := 0
	maxScore := 0
	for i := 0; i < len(items); i++ {
		for j := 0; j < len(items[0]); j++ {
			isVis, score := someSortOfBfs(items, i, j)
			if isVis {
				// fmt.Printf("%d, %d is visible\n", i, j)
				visCount++
			}
			if score > maxScore {
				maxScore = score
			}
		}
	}
	return visCount, maxScore
}

func intArrayrIze(in []byte) [][]int {
	lines := bytes.Split(in, []byte("\n"))
	var final [][]int
	for lidx, l := range lines {
		if len(l) > 0 {
			final = append(final, make([]int, len(l)))

			for idx, char := range l {
				i, _ := strconv.ParseInt(string(char), 10, 64)
				final[lidx][idx] = int(i)
			}
		}
	}
	return final
}

func Solve() error {
	input := intArrayrIze(puzzleinput.Day(8))

	vis, maxScore := countVisibleTrees(input)
	fmt.Printf("Part 1: %d visible trees\n", vis)
	fmt.Printf("part 2: %d max score", maxScore)

	return nil
}
