package twelve

import (
	"bytes"
	"fmt"

	puzzleinput "github.com/fwwieffering/aoc2022/internal/puzzle-input"
)

type coord struct {
	row int
	col int
}

type queueItem struct {
	coord    coord
	distance int
}

func getNeighbors(in [][]byte, visited map[coord]int, c coord, compFunc func(int, int) bool) []coord {
	potentialNeighbors := []coord{
		{row: c.row - 1, col: c.col},
		{row: c.row + 1, col: c.col},
		{row: c.row, col: c.col + 1},
		{row: c.row, col: c.col - 1},
	}
	finalNeighbors := []coord{}
	for _, n := range potentialNeighbors {
		if n.row >= 0 && n.col >= 0 && n.row < len(in) && n.col < len(in[0]) {
			if _, ok := visited[n]; !ok {
				if compFunc(getAdjustedHeight(in[n.row][n.col]), getAdjustedHeight(in[c.row][c.col])) {
					finalNeighbors = append(finalNeighbors, n)
				}
			}
		}
	}
	return finalNeighbors
}

func getAdjustedHeight(in byte) int {
	switch in {
	case 'S':
		return int('a')
	case 'E':
		return int('z')
	default:
		return int(in)
	}
}

func bfs(in [][]byte, queue []queueItem, myVisited map[coord]int, otherVisited map[coord]int, forward bool) ([]queueItem, int) {
	if len(queue) > 0 {
		item := queue[0]
		if len(queue) > 1 {
			queue = queue[1:]
		} else {
			queue = []queueItem{}
		}
		// check if intersecting
		if ov, ok := otherVisited[item.coord]; ok {
			return nil, ov + item.distance
		}
		// for forward - compfunc returns true if the neighbors height is at most one more
		// for backward - compfunc returns true if the neighbors height is at most one less
		var compFunc func(int, int) bool
		if forward {
			compFunc = func(i1, i2 int) bool { return i1-i2 <= 1 }
		} else {
			compFunc = func(i1, i2 int) bool { return i2-i1 <= 1 }
		}
		neighbors := getNeighbors(in, myVisited, item.coord, compFunc)
		for _, n := range neighbors {
			myVisited[n] = item.distance + 1
			queue = append(queue, queueItem{coord: n, distance: item.distance + 1})
		}
	}
	return queue, -1
}

func shortestPathBidirectional(in [][]byte, start coord, end coord) int {
	visitedFwd := map[coord]int{start: 0}
	visitedBwd := map[coord]int{end: 0}

	queueFwd := []queueItem{{coord: start, distance: 0}}
	queueBwd := []queueItem{{coord: end, distance: 0}}

	for len(queueFwd) > 0 || len(queueBwd) > 0 {
		var fwdHit, bwdHit int
		queueFwd, fwdHit = bfs(in, queueFwd, visitedFwd, visitedBwd, true)
		if fwdHit != -1 {
			return fwdHit
		}
		queueBwd, bwdHit = bfs(in, queueBwd, visitedBwd, visitedFwd, false)
		if bwdHit != -1 {
			return bwdHit
		}
	}
	return -1
}

func allDistancesTo(in [][]byte, start coord, compFxn func(int, int) bool, endCon func(b byte) bool) int {
	queue := []queueItem{{coord: start, distance: 0}}
	visited := map[coord]int{start: 0}

	final := map[coord]int{}
	for len(queue) > 0 {
		item := queue[0]
		if len(queue) > 1 {
			queue = queue[1:]
		} else {
			queue = []queueItem{}
		}
		// store the distances we care about
		if endCon(in[item.coord.row][item.coord.col]) {
			final[item.coord] = item.distance
		}
		neighbors := getNeighbors(in, visited, item.coord, compFxn)
		for _, n := range neighbors {
			visited[n] = item.distance + 1
			queue = append(queue, queueItem{coord: n, distance: item.distance + 1})
		}
	}
	// get the lowest distance
	min := -1
	for _, v := range final {
		if min == -1 || v < min {
			min = v
		}
	}
	return min
}

func Solve() error {
	in := bytes.Split(bytes.TrimRight(puzzleinput.Day(12), "\n"), []byte("\n"))
	var start coord
	var end coord
	for ridx := range in {
		for cidx := range in[ridx] {
			if in[ridx][cidx] == 'S' {
				start.row = ridx
				start.col = cidx
			} else if in[ridx][cidx] == 'E' {
				end.row = ridx
				end.col = cidx
			}
		}
	}

	part1 := shortestPathBidirectional(in, start, end)
	fmt.Printf("Part 1: %d\n", part1)

	part2 := allDistancesTo(in, end, func(i1, i2 int) bool { return i2-i1 <= 1 }, func(b byte) bool { return b == 'a' })
	fmt.Printf("Part 2: %d\n", part2)
	return nil
}
