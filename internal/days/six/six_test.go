package six

import "testing"

type answerKey struct {
	part1 int
	part2 int
}

var testCases = map[string]answerKey{
	"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    {part1: 7, part2: 19},
	"bvwbjplbgvbhsrlpgdmjqwftvncz":      {part1: 5, part2: 23},
	"nppdvjthqldpwncqszvftbrmjlhg":      {part1: 6, part2: 23},
	"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": {part1: 10, part2: 29},
	"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  {part1: 11, part2: 26},
}

func TestFindStartOfPacketMarker(t *testing.T) {
	for in, answer := range testCases {
		res := findMarker([]byte(in), 4)
		if res != answer.part1 {
			t.Fatalf("%s start-of-packet-marker should be %d got %d", in, answer.part1, res)
		}
		p2 := findMarker([]byte(in), 14)
		if p2 != answer.part2 {
			t.Fatalf("%s part 2 should be be %d got %d", in, answer.part2, p2)
		}
	}
}
