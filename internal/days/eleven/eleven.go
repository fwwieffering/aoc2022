package eleven

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	puzzleinput "github.com/fwwieffering/aoc2022/internal/puzzle-input"
)

type Monkey struct {
	Idx        int
	Items      []int64
	Divisor    int64
	Operation  func(int64) int64
	MonkeyTest func(int64) int64
}

func makeOperatorFunc(operator string, val string) func(
	int64) int64 {
	if val == "old" {
		return func(x int64) int64 {
			return x * x
		}
	}
	iVal, _ := strconv.ParseInt(val, 10, 64)
	return map[string]func(int64) func(int64) int64{
		"+": func(i int64) func(int64) int64 {
			return func(x int64) int64 {
				return x + i
			}
		},
		"*": func(i int64) func(int64) int64 {
			return func(x int64) int64 {
				return x * i
			}
		},
	}[operator](iVal)
}

func makeTestFunc(divisor int64, trueMonkey int64, falseMonkey int64) func(int64) int64 {
	return func(i int64) int64 {
		if i%divisor == 0 {
			return trueMonkey
		}
		return falseMonkey
	}
}

func parseMonkeys(in string) []*Monkey {
	lines := strings.Split(in, "\n")
	var res []*Monkey
	var curMonkey *Monkey

	var testDivisor int64
	var trueMonkeyIdx int64
	var falseMonkeyIdx int64

	for _, l := range lines {
		if len(l) == 0 {
			if curMonkey != nil {
				curMonkey.MonkeyTest = makeTestFunc(testDivisor, trueMonkeyIdx, falseMonkeyIdx)
				res = append(res, curMonkey)
				curMonkey = nil
			}
		} else {
			yamlLine := strings.Split(l, ":")
			key := yamlLine[0]
			switch key[0:6] {
			case "Monkey":
				var monkeyNum int
				fmt.Sscanf(key, "Monkey %d", &monkeyNum)
				curMonkey = &Monkey{Idx: monkeyNum}
			case "  Star":
				vals := strings.Split(yamlLine[1], ",")
				items := make([]int64, len(vals))

				for idx, rv := range vals {
					var v int64
					fmt.Sscanf(rv, "%d", &v)
					items[idx] = v
				}
				curMonkey.Items = items
			case "  Oper":
				var v string
				var op string
				fmt.Sscanf(yamlLine[1], " new = old %s %s", &op, &v)
				curMonkey.Operation = makeOperatorFunc(op, v)
			case "  Test":
				var v int64
				fmt.Sscanf(yamlLine[1], " divisible by %d", &v)
				curMonkey.Divisor = v
				testDivisor = v
			case "    If":
				var c bool
				fmt.Sscanf(key, "    If %t", &c)
				var monkeyIdx int64
				fmt.Sscanf(yamlLine[1], " throw to monkey %d", &monkeyIdx)
				if c {
					trueMonkeyIdx = monkeyIdx
				} else {
					falseMonkeyIdx = monkeyIdx
				}
			}
		}

	}

	return res
}

func doMonkeyBusiness(monkeys []*Monkey, numRounds int, divideByThree bool) int64 {
	var monkeyInspections = make([]int, len(monkeys))

	// least common multiple - I cheated to figure this out. Did not know it was an option I tried to use BigInts
	lcm := int64(1)
	for _, m := range monkeys {
		lcm *= m.Divisor
	}

	for i := 0; i < numRounds; i++ {
		for mIdx := 0; mIdx < len(monkeys); mIdx++ {
			// fmt.Printf("Monkey %d\n", mIdx)
			for len(monkeys[mIdx].Items) > 0 {
				// tick up the inspection counter
				monkeyInspections[mIdx]++
				// inspect item (then divide by three)
				item := monkeys[mIdx].Items[0]
				// fmt.Printf("  Monkey inspects an item with a worry level of %v\n", item)
				updatedLevel := monkeys[mIdx].Operation(item)
				// fmt.Printf("    Worry level is updated to %v\n", updatedLevel)
				if divideByThree {
					updatedLevel = updatedLevel / 3
					// fmt.Printf("    Monkey gets bored with item. Worry level is divided by 3 to %v\n", updatedLevel)
				}
				// move item to new monkey
				newMonkeyIdx := monkeys[mIdx].MonkeyTest(updatedLevel)
				// reduce the worry level to the lcm + remainder. This makes all our tests return the same result but keeps the number
				// manageable. thanks internet for the help
				if updatedLevel > lcm {
					updatedLevel = lcm + (updatedLevel % lcm)
				}

				// fmt.Printf("    Item with worry level %v is thrown to monkey %d\n", updatedLevel, newMonkeyIdx)
				monkeys[newMonkeyIdx].Items = append(monkeys[newMonkeyIdx].Items, updatedLevel)
				if len(monkeys[mIdx].Items) == 1 {
					monkeys[mIdx].Items = []int64{}
				} else {
					monkeys[mIdx].Items = monkeys[mIdx].Items[1:]
				}
			}
		}
	}

	sort.Slice(monkeyInspections, func(i, j int) bool { return monkeyInspections[i] > monkeyInspections[j] })
	return int64(monkeyInspections[0]) * int64(monkeyInspections[1])
}
func Solve() error {
	in := puzzleinput.Day(11)
	monkeys := parseMonkeys(string(in))
	part1 := doMonkeyBusiness(monkeys, 20, true)
	fmt.Printf("Part 1: %v\n", part1)
	// double parsing rather than copying :\
	monkeys2 := parseMonkeys(string(in))
	part2 := doMonkeyBusiness(monkeys2, 10000, false)
	fmt.Printf("Part 2: %v\n", part2)
	return nil
}
