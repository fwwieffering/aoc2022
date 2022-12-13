package thirteen

import (
	"fmt"
	"strconv"

	puzzleinput "github.com/fwwieffering/aoc2022/internal/puzzle-input"
)

type Packet struct {
	Left  []interface{}
	Right []interface{}
}

func (p *Packet) SetSide(side string, data []interface{}) {
	if side == "left" {
		p.Left = data
	} else {
		p.Right = data
	}
}

// return two bools - valid, final. if valid and not final keep going
func validate(left, right []interface{}, level int, debug bool) (bool, bool) {
	buf := ""
	for i := 0; i < level; i++ {
		buf += "  "
	}
	if level == 0 && debug {
		fmt.Printf("%s- Compare %+v vs %+v\n", buf, left, right)
	}

	valid := false
	final := false

	idx := 0
	for idx < len(left) && idx < len(right) {
		if debug {
			fmt.Printf("%s  - Compare %d vs %d\n", buf, left[idx], right[idx])
		}
		leftInt, leftIsInt := left[idx].(int)
		leftArr, leftIsArr := left[idx].([]interface{})
		rightInt, rightIsInt := right[idx].(int)
		rightArr, rightIsArr := right[idx].([]interface{})
		if leftIsInt && rightIsInt {
			// If the left integer is lower than the right integer, the inputs are in the right order
			if leftInt != rightInt {
				if leftInt < rightInt && debug {
					fmt.Printf("%s    - Left side is smaller, so inputs are in the right order\n", buf)
				} else if debug {
					fmt.Printf("%s    - Right side is smaller, so inputs are *not* in the right order\n", buf)
				}
				return leftInt < rightInt, true
			}
		} else if rightIsArr || leftIsArr {
			if !rightIsArr {
				valid, final = validate(leftArr, []interface{}{rightInt}, level+1, debug)
			} else if !leftIsArr {
				valid, final = validate([]interface{}{leftInt}, rightArr, level+1, debug)
			} else {
				valid, final = validate(leftArr, rightArr, level+1, debug)
			}
			if final {
				return valid, final
			}
		}
		idx++
	}
	// If the left list runs out of items first, the inputs are in the right order
	if len(left) != len(right) {
		if len(left) < len(right) {
			if debug {
				fmt.Printf("%s    - Left side ran out of items, so inputs are in the right order\n", buf)
			}
			return true, true
		} else {
			if debug {
				fmt.Printf("%s    - Right side ran out of items, so inputs are *not* in the right order\n", buf)
			}
			return false, true

		}
	}

	return valid, final
}

func ValidatePackets(left []interface{}, right []interface{}, debug bool) bool {
	valid, _ := validate(left, right, 0, debug)
	return valid
}
func (p *Packet) Validate(debug bool) bool {
	valid, _ := validate(p.Left, p.Right, 0, debug)
	return valid
}

func parsePackets(in []byte) []*Packet {
	final := []*Packet{}

	currentPacket := &Packet{}
	var curArr *[]interface{} = nil
	parentArrays := []*[]interface{}{}
	packetSide := "left"
	intStr := ""

	for idx, b := range in {
		switch b {
		// newline means an end of a packet or pair of packets if two
		case '\n':
			if in[idx-1] == '\n' {
				final = append(final, currentPacket)
				currentPacket = &Packet{}
				packetSide = "left"
			} else {
				packetSide = "right"
			}

		case '[':
			if curArr == nil {
				curArr = &[]interface{}{}
			} else {
				newArr := []interface{}{}
				*curArr = append(*curArr, newArr)
				parentArrays = append(parentArrays, curArr)
				curArr = &newArr
			}
		case ']':
			if intStr != "" {
				i, _ := strconv.Atoi(string(intStr))
				*curArr = append(*curArr, i)
				intStr = ""
			}
			if len(parentArrays) > 0 {
				parent := parentArrays[len(parentArrays)-1]
				(*parent)[len(*parent)-1] = *curArr
				curArr = parent
				parentArrays = parentArrays[:len(parentArrays)-1]
			} else {
				currentPacket.SetSide(packetSide, *curArr)
				curArr = nil
			}
		case ',':
			if intStr != "" {
				i, _ := strconv.Atoi(string(intStr))
				*curArr = append(*curArr, i)
				intStr = ""
			}
		// a number
		default:
			intStr += string(b)
		}
	}
	// hanging packet
	if currentPacket.Left != nil && currentPacket.Right != nil {
		final = append(final, currentPacket)
	}

	return final
}

func part1(packets []*Packet) int {
	part1 := 0
	for idx, p := range packets {
		if p.Validate(false) {
			part1 += idx + 1
		}
	}
	return part1
}

func Solve() error {
	in := puzzleinput.Day(13)

	packets := parsePackets(in)
	fmt.Printf("Part1: %d\n", part1(packets))
	// add divider packets, then sort for part 2
	packets = append(packets)

	return nil
}
