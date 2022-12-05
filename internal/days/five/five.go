package five

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"

	puzzleinput "github.com/fwwieffering/aoc2022/internal/puzzle-input"
)

type litem struct {
	val  string
	next *litem
}

type ll struct {
	head *litem
	tail *litem
}

// append adds to end of list
func (l *ll) append(item *litem) {
	if l.head == nil {
		l.head = item
		l.tail = item
	} else {
		l.tail.next = item
		l.tail = item
	}
}

// prepend adds to start of list
func (l *ll) prepend(item *litem) {
	if l.head == nil {
		l.head = item
		l.tail = item
	} else {
		item.next = l.head
		l.head = item
	}
}

func (l *ll) popHead() *litem {
	res := l.head
	if l.head != nil {
		l.head = l.head.next
		res.next = nil
	}
	return res
}

func (l *ll) popN(count int) (*litem, *litem) {
	stackHead := l.head
	stackTail := l.head
	for i := 1; i < count; i++ {
		if stackTail != nil {
			stackTail = stackTail.next
		}
	}
	if stackTail != nil {
		l.head = stackTail.next
	} else {
		l.head = nil
	}
	return stackHead, stackTail
}

func (l *ll) shiftStack(stackHead, stackTail *litem) {
	stackTail.next = l.head
	l.head = stackHead
}

func parseCrates(crates []*ll, line []byte) []*ll {
	for i := 0; i < len(line); i += 4 {
		crateIdx := i / 4

		if crateIdx > len(crates)-1 {
			crates = append(crates, &ll{})
		}
		var curSlice []byte
		if i+3 >= len(line) {
			curSlice = line[i:]
		} else {
			curSlice = line[i : i+3]
		}
		if string(curSlice) != "   " {
			crates[crateIdx].append(&litem{val: string(curSlice[1])})
		}
	}
	return crates
}

var instructionRegex = regexp.MustCompile("move ([0-9]+) from ([0-9]+) to ([0-9]+)")

func parseInstruction(in []byte) (int, int, int) {
	instructionMatches := instructionRegex.FindAllStringSubmatch(string(in), -1)
	if len(instructionMatches) < 1 || len(instructionMatches[0]) != 4 {
		fmt.Printf("%+v\n", instructionMatches)
		panic(fmt.Sprintf("couldn't parse instruction '%s'", string(in)))
	}
	var res = []int{0, 0, 0}
	for i := 1; i < 4; i++ {
		p, _ := strconv.ParseInt(instructionMatches[0][i], 10, 64)
		res[i-1] = int(p)
	}
	return res[0], res[1], res[2]
}

func processStacks(in []byte) ([]*ll, []*ll) {
	lines := bytes.Split(in, []byte("\n"))

	stackDone := false
	var cratesPart1 []*ll
	var cratesPart2 []*ll
	for _, l := range lines {
		if len(l) >= 3 {
			if !stackDone {
				// the type of line can be determined from the first three characters
				isCrate, _ := regexp.Match("(   |\\[[A-Z]\\])", l[0:3])
				if isCrate {
					cratesPart1 = parseCrates(cratesPart1, l)
					cratesPart2 = parseCrates(cratesPart2, l)
				}
			} else {
				count, source, dest := parseInstruction(l)
				// part 1 processing
				sourceStack1 := cratesPart1[source-1]
				destStack1 := cratesPart1[dest-1]
				for i := 0; i < count; i++ {
					item := sourceStack1.popHead()
					destStack1.prepend(item)
				}

				// part 2 processing
				sourceStack2 := cratesPart2[source-1]
				destStack2 := cratesPart2[dest-1]
				sH, sT := sourceStack2.popN(count)
				destStack2.shiftStack(sH, sT)
			}
		} else if !stackDone {
			// the stack of crates is separated from the instructions by a blank line
			stackDone = true
		}
	}
	return cratesPart1, cratesPart2
}

func Solve() error {
	part1FinalCrates, part2FinalCrates := processStacks(puzzleinput.Day(5))

	part1 := ""
	for _, c := range part1FinalCrates {
		part1 += c.head.val
	}
	fmt.Printf("Part 1: %s\n", part1)

	part2 := ""
	for _, c := range part2FinalCrates {
		part2 += c.head.val
	}
	fmt.Printf("Part 2: %s\n", part2)
	return nil
}
