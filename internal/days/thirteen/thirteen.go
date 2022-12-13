package thirteen

import (
	"strconv"
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

func parsePackets(in []byte) []*Packet {
	final := []*Packet{}

	currentPacket := &Packet{}
	var curArr *[]interface{} = nil
	parentArrays := []*[]interface{}{}
	packetSide := "left"

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
			if len(parentArrays) > 0 {
				parent := parentArrays[len(parentArrays)-1]
				(*parent)[len(*parent)-1] = *curArr
				curArr = parent

				parentArrays = parentArrays[:len(parentArrays)-1]
			} else {
				currentPacket.SetSide(packetSide, *curArr)
				curArr = nil
			}
		// don't care about commas - they are delimiters
		case ',':
			continue
		// a number
		default:
			i, _ := strconv.Atoi(string(b))
			*curArr = append(*curArr, i)
		}
	}
	// hanging packet
	if currentPacket.Left != nil && currentPacket.Right != nil {
		final = append(final, currentPacket)
	}

	return final
}

func Solve() error {
	return nil
}
