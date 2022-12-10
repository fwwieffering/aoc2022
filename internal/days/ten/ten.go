package ten

import (
	"fmt"
	"strings"

	puzzleinput "github.com/fwwieffering/aoc2022/internal/puzzle-input"
)

type Cpu struct {
	cyclenum int
	x        int

	rows [][]rune

	curCmd      string
	curArgs     []int
	cmdEndCycle int
}

func NewCpu() *Cpu {
	// hardcoding the screen size
	return &Cpu{
		cyclenum: 1,
		x:        1,
		rows: [][]rune{
			make([]rune, 40),
			make([]rune, 40),
			make([]rune, 40),
			make([]rune, 40),
			make([]rune, 40),
			make([]rune, 40),
		},
	}
}

func (c *Cpu) SetCmd(cmd string, args ...int) {
	c.curCmd = cmd
	c.curArgs = args

	switch cmd {
	case "noop":
		c.cmdEndCycle = c.cyclenum + 1
	case "addx":
		c.cmdEndCycle = c.cyclenum + 2
	}
}

func (c *Cpu) draw() {
	rowIdx := (c.cyclenum - 1) / 40
	colIdx := (c.cyclenum - 1) % 40

	spriteCenter := c.x

	char := '.'
	if colIdx >= spriteCenter-1 && colIdx <= spriteCenter+1 {
		char = '#'
	}
	c.rows[rowIdx][colIdx] = char
}

func (c *Cpu) Cycle() (int, bool) {
	// first draw
	c.draw()
	// then cycle
	c.cyclenum++
	if c.cyclenum == c.cmdEndCycle {
		switch c.curCmd {
		case "noop":

		case "addx":
			c.x += c.curArgs[0]
		}
		c.curArgs = nil
		c.curCmd = ""
		c.cmdEndCycle = -1

		return c.cyclenum, true
	}
	return c.cyclenum, false
}

func process(in string) int {
	lines := strings.Split(in, "\n")

	c := NewCpu()
	count := 0

	var cycleNum int
	for _, l := range lines {
		if len(l) > 0 {
			var cmd string
			var arg int
			fmt.Sscanf(l, "%s%d", &cmd, &arg)
			c.SetCmd(cmd, arg)
			cmdDone := false
			for !cmdDone {
				cycleNum, cmdDone = c.Cycle()
				if (cycleNum-20)%40 == 0 {
					count += c.x * cycleNum

				}
			}
		}
	}
	// print terminal
	for _, r := range c.rows {
		for _, c := range r {
			fmt.Printf("%c", c)
		}
		fmt.Printf("\n")
	}
	return count
}

func Solve() error {
	in := string(puzzleinput.Day(10))

	part1 := process(in)
	fmt.Printf("part 1: %d\n", part1)
	return nil
}
