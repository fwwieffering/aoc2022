package fourteen

import (
	"fmt"
	"strings"

	puzzleinput "github.com/fwwieffering/aoc2022/internal/puzzle-input"
)

type Coord struct {
	x int
	y int
}

func ParseCoord(in string) Coord {
	var x int
	var y int
	fmt.Sscanf(in, "%d,%d", &x, &y)
	return Coord{x: x, y: y}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func GetLine(a, b Coord) []Coord {
	final := []Coord{}

	xmin := min(a.x, b.x)
	xmax := max(a.x, b.x)
	for x := xmin; x <= xmax; x++ {
		// using a.y because this should be a straight line
		final = append(final, Coord{x: x, y: a.y})
	}

	ymin := min(a.y, b.y)
	ymax := max(a.y, b.y)
	for y := ymin; y <= ymax; y++ {
		// using a.x because this should be a straight line
		final = append(final, Coord{x: a.x, y: y})
	}
	return final
}

type Cave struct {
	// minY is 0 so no need to track that
	minX int
	maxX int
	maxY int

	walls map[Coord]bool
	sand  map[Coord]bool
}

func NewCave() *Cave {
	return &Cave{
		minX:  -1,
		maxX:  -1,
		maxY:  -1,
		walls: make(map[Coord]bool),
		sand:  make(map[Coord]bool),
	}
}

func (c *Cave) Draw(withFloor bool) {
	xDiff := c.maxX - c.minX
	yMax := c.maxY
	if withFloor {
		yMax += 2
	}
	for y := 0; y <= yMax; y++ {
		line := ""
		for x := 0; x <= xDiff; x++ {
			if y == yMax {
				line += "#"
			} else if c.walls[Coord{x: c.minX + x, y: y}] {
				line += "#"
			} else if c.sand[Coord{x: c.minX + x, y: y}] {
				line += "o"
			} else {
				line += "."
			}
		}
		fmt.Printf("%s\n", line)
	}
}

func (c *Cave) AddWalls(in string) {
	splitWall := strings.Split(in, " -> ")
	for i := 0; i < len(splitWall); i++ {
		if i+1 < len(splitWall) {
			a := ParseCoord(splitWall[i])
			b := ParseCoord(splitWall[i+1])
			// set global max/mins if applicable
			localMinX := min(a.x, b.x)
			if c.minX == -1 || localMinX < c.minX {
				c.minX = localMinX
			}
			localMaxX := max(a.x, b.x)
			if c.maxX == -1 || localMaxX > c.maxX {
				c.maxX = localMaxX
			}
			localMaxY := max(a.y, b.y)
			if c.maxY == -1 || localMaxY > c.maxY {
				c.maxY = localMaxY
			}

			for _, wallSeg := range GetLine(a, b) {
				c.walls[wallSeg] = true
			}
		}
	}
}

// this could be faster if we used the position of the last unit of sand
// returns a bool indicating whether the unit of sand came to rest
func (c *Cave) ProcessSandUnit(withFloor bool) bool {
	// sand starts at 500,0
	cur := Coord{x: 500, y: 0}
	atRest := false

	for !atRest {
		// assume the floor is an infinite horizontal line with a y coordinate equal to two plus the highest y coordinate of any point in your scan
		atFloor := withFloor && cur.y+1 >= c.maxY+1

		// if we're over the maxY then this is done
		if !atFloor && cur.y > c.maxY {
			return false
		}

		// A unit of sand always falls down one step if possible
		down1 := Coord{x: cur.x, y: cur.y + 1}
		// If the tile immediately below is blocked (by rock or sand), the unit of sand attempts to instead move diagonally one step down and to the left
		diagLeft := Coord{x: cur.x - 1, y: cur.y + 1}
		// If that tile is blocked, the unit of sand attempts to instead move diagonally one step down and to the right
		diagRight := Coord{x: cur.x + 1, y: cur.y + 1}

		if !c.walls[down1] && !c.sand[down1] {
			cur = down1
		} else if !c.walls[diagLeft] && !c.sand[diagLeft] {
			cur = diagLeft
		} else if !c.walls[diagRight] && !c.sand[diagRight] {
			cur = diagRight
		} else {
			atRest = true
		}
		if atFloor {
			break
		}
	}
	c.sand[cur] = true
	if cur.x > c.maxX {
		c.maxX = cur.x
	}
	if cur.x < c.minX {
		c.minX = cur.x
	}
	return !(cur.x == 500 && cur.y == 0)
}

func (c *Cave) SandFall(withFloor bool) int {
	sandStillFalling := true
	for sandStillFalling {
		sandStillFalling = c.ProcessSandUnit(withFloor)
	}
	return len(c.sand)
}

func ParseCave(in []byte) *Cave {
	cave := NewCave()
	for _, line := range strings.Split(string(in), "\n") {
		cave.AddWalls(line)
	}
	return cave
}

func Solve() error {
	cave := ParseCave(puzzleinput.Day(14))
	part1 := cave.SandFall(false)
	fmt.Printf("Part 1: %d\n", part1)
	part2 := cave.SandFall(true)
	fmt.Printf("Part 2: %d\n", part2)
	return nil
}
