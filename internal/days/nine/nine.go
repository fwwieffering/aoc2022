package nine

import (
	"fmt"
	"strings"

	puzzleinput "github.com/fwwieffering/aoc2022/internal/puzzle-input"
)

type coord struct {
	x int
	y int
}

func magnitude(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func direction(i int) int {
	if i == 0 {
		return 0
	} else if i > 0 {
		return -1
	}
	return 1
}
func (c *coord) getNewTail(head *coord) *coord {
	var xVec = (head.x - c.x)
	var xMagnitude = magnitude(xVec)
	var yVec = (head.y - c.y)
	var yMagnitude = magnitude(yVec)
	// tail doesn't move unless its outside of the neighbors
	if xMagnitude <= 1 && yMagnitude <= 1 {
		return c
	}
	// find the greater magnitude difference to determine which direction from
	// head tail is facing
	if xVec != 0 && yVec != 0 {
		if xMagnitude > yMagnitude {
			yVec = 0
		} else if xMagnitude < yMagnitude {
			xVec = 0
		}
	}
	return &coord{x: head.x + direction(xVec), y: head.y + direction(yVec)}
}

func (c *coord) moveHead(dir string, count int) *coord {
	switch dir {
	case "L":
		return &coord{x: c.x - count, y: c.y}
	case "R":
		return &coord{x: c.x + count, y: c.y}
	case "D":
		return &coord{x: c.x, y: c.y - count}
	case "U":
		return &coord{x: c.x, y: c.y + count}
	default:
		return nil
	}
}

func doPrintGrid(rope []*coord, gridSize *coord) {
	var grid = make([][]string, gridSize.y)
	for row := 0; row < gridSize.y; row++ {
		grid[row] = make([]string, gridSize.x)
		for col := 0; col < gridSize.x; col++ {
			grid[row][col] = "."
		}
	}
	for i := len(rope) - 1; i >= 0; i-- {
		symb := fmt.Sprintf("%d", i)
		if i == 0 {
			symb = "H"
		}
		grid[rope[i].y][rope[i].x] = symb
	}

	for i := len(grid) - 1; i >= 0; i-- {
		fmt.Println(strings.Join(grid[i], " "))
	}
}

func processInstructions(in string, ropeSize int, printGrid bool, gridSize *coord) int {
	lines := strings.Split(in, "\n")

	rope := make([]*coord, ropeSize)
	for i := 0; i < ropeSize; i++ {
		rope[i] = &coord{x: 0, y: 0}
	}
	tailSpots := map[coord]bool{}
	tailSpots[coord{x: 0, y: 0}] = true

	for _, l := range lines {
		if len(l) > 0 {
			var dir string
			var count int
			fmt.Sscanf(l, "%s%d", &dir, &count)
			if printGrid {
				fmt.Printf("=== %s %d ===\n", dir, count)
			}
			for i := 0; i < count; i++ {
				for i := 0; i < ropeSize; i++ {
					if i == 0 {
						rope[i] = rope[i].moveHead(dir, 1)
					} else {
						rope[i] = rope[i].getNewTail(rope[i-1])
					}
				}
				if printGrid {
					doPrintGrid(rope, gridSize)
					fmt.Println("")
				}
				tailSpots[*rope[ropeSize-1]] = true
			}
		}
	}

	return len(tailSpots)
}
func Solve() error {
	in := string(puzzleinput.Day(9))
	part1 := processInstructions(in, 1, false, nil)
	part2 := processInstructions(in, 10, false, nil)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
	return nil
}
