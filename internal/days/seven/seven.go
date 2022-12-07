package seven

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	puzzleinput "github.com/fwwieffering/aoc2022/internal/puzzle-input"
)

type File struct {
	Name string
	Size int64
}

type Dir struct {
	Path          string
	Parent        *Dir
	Children      []*Dir
	RecursiveSize int64
}

func (d Dir) Print(level int) {
	msg := ""
	for i := 0; i < level*2; i++ {
		msg += " "
	}
	msg += fmt.Sprintf("- %s (dir) size=%d", d.Path, d.RecursiveSize)
	fmt.Println(msg)
	for _, c := range d.Children {
		c.Print(level + 1)
	}
}

func (d *Dir) AddRecursiveSize(amt int64) {
	d.RecursiveSize = d.RecursiveSize + amt
	if d.Parent != nil {
		d.Parent.AddRecursiveSize(amt)
	}
}

func processCommand(terminal []string, currentDir *Dir) (int, *Dir) {
	lineNum := 0
	// parse command
	splitLine := strings.Split(terminal[lineNum], " ")
	// new command
	var command = splitLine[1]
	var params []string
	if len(splitLine) > 2 {
		params = splitLine[2:]
	} else {
		params = nil
	}
	lineNum++
	switch command {
	case "cd":
		switch params[0] {
		case "..":
			currentDir = currentDir.Parent
		default:
			// todo : fully qualify path?
			path := params[0]
			newDir := &Dir{Path: path, Parent: currentDir}
			if currentDir != nil {
				currentDir.Children = append(currentDir.Children, newDir)
			}
			currentDir = newDir
		}
	case "ls":
		for lineNum < len(terminal) && len(terminal[lineNum]) > 0 {
			if terminal[lineNum][0] == '$' {
				break
			}
			splitLine := strings.Split(terminal[lineNum], " ")
			fileSize, _ := strconv.ParseInt(splitLine[0], 10, 64)
			currentDir.AddRecursiveSize(fileSize)
			lineNum++
		}
	default:
		panic(fmt.Sprintf("unknown command %s", command))
	}
	return lineNum, currentDir
}

func processFileSystem(in string) (*Dir, []*Dir) {
	lines := strings.Split(in, "\n")
	lineIdx := 0

	var uniqueDirs = map[*Dir]struct{}{}
	var rootDir *Dir
	var currentDir *Dir

	for lineIdx < len(lines) && len(lines[lineIdx]) > 0 {
		var incrLineNumBy int
		incrLineNumBy, currentDir = processCommand(lines[lineIdx:], currentDir)
		uniqueDirs[currentDir] = struct{}{}
		lineIdx = lineIdx + incrLineNumBy
		if rootDir == nil {
			rootDir = currentDir
		}
	}
	// process dirs into a directory
	var allDirs = make([]*Dir, len(uniqueDirs))
	idx := 0
	for d, _ := range uniqueDirs {
		allDirs[idx] = d
		idx++
	}

	return rootDir, allDirs
}

func Solve() error {
	input := string(puzzleinput.Day(7))

	rootDir, allDirs := processFileSystem(input)

	// sort dirs by size increasing
	sort.Slice(allDirs, func(i, j int) bool {
		return allDirs[i].RecursiveSize < allDirs[j].RecursiveSize
	})

	var part1 int64 = 0
	var part2 int64 = 0
	var currentSpace = 70000000 - rootDir.RecursiveSize
	var neededSpace = 30000000
	for _, d := range allDirs {
		if d.RecursiveSize <= 100000 {
			part1 += d.RecursiveSize
		}
		if d.RecursiveSize+currentSpace >= int64(neededSpace) {
			part2 = d.RecursiveSize
			break
		}
	}
	fmt.Printf("Part 1: Sum of all dir sizes <= 100000 %d\n", part1)
	fmt.Printf("Part 2: Smallest possible dir that will still provide the needed space %d\n", part2)
	return nil
}
