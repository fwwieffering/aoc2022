package thirteen

import (
	"fmt"
	"testing"

	puzzleinput "github.com/fwwieffering/aoc2022/internal/puzzle-input"
	"github.com/google/go-cmp/cmp"
)

var example = []byte(`[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]
`)

var expected = []*Packet{
	{
		Left:  []interface{}{1, 1, 3, 1, 1},
		Right: []interface{}{1, 1, 5, 1, 1},
	},
	{
		Left:  []interface{}{[]interface{}{1}, []interface{}{2, 3, 4}},
		Right: []interface{}{[]interface{}{1}, 4},
	},
	{
		Left:  []interface{}{9},
		Right: []interface{}{[]interface{}{8, 7, 6}},
	},
	{
		Left:  []interface{}{[]interface{}{4, 4}, 4, 4},
		Right: []interface{}{[]interface{}{4, 4}, 4, 4, 4},
	},
	{
		Left:  []interface{}{7, 7, 7, 7},
		Right: []interface{}{7, 7, 7},
	},
	{
		Left:  []interface{}{},
		Right: []interface{}{3},
	},
	{
		Left:  []interface{}{[]interface{}{[]interface{}{}}},
		Right: []interface{}{[]interface{}{}},
	},
	{
		Left:  []interface{}{1, []interface{}{2, []interface{}{3, []interface{}{4, []interface{}{5, 6, 7}}}}, 8, 9},
		Right: []interface{}{1, []interface{}{2, []interface{}{3, []interface{}{4, []interface{}{5, 6, 0}}}}, 8, 9},
	},
}

var validationExpectations = []bool{
	true,
	true,
	false,
	true,
	false,
	true,
	false,
	false,
}

func TestParsePackets(t *testing.T) {
	result := parsePackets(example)
	for idx, exp := range expected {
		if !cmp.Equal(exp, result[idx]) {
			fmt.Printf("left: %+v\nright:%+v\n", result[idx].Left, result[idx].Right)
			t.Fatalf("idx: %d\n%s", idx, cmp.Diff(exp, result[idx]))
		}
	}

	for idx, expectValid := range validationExpectations {
		fmt.Printf("----------- %d -----------\n", idx)
		isValid := result[idx].Validate(false)
		if isValid != expectValid {
			t.Fatalf("idx %d expected %t but got %t", idx, expectValid, isValid)
		}
	}
	p1 := part1(result)
	if p1 != 13 {
		t.Fatalf("expected 13 got %d", p1)
	}
}

var realResults = []bool{true, false, false, false, false, false, true, true, false, true, true, false, true, false, true, false, true, false, false, true, true, false, true, false, true, false, true, false, true, true, false, true, false, true, false, false, true, false, false, true, true, true, true, true, false, false, false, true, false, true, true, false, true, true, true, false, false, true, true, true, true, false, true, false, true, false, false, true, false, true, true, false, true, true, true, true, false, false, true, true, true, false, false, true, true, true, true, true, false, false, false, false, false, true, true, true, false, false, false, true, false, true, true, true, false, true, true, true, true, true, true, true, false, true, false, false, false, true, false, true, false, true, true, true, true, false, true, false, false, true, false, false, true, true, false, true, false, true, false, true, false, true, false, false, false, true, false, true, false, false}

func TestRealResults(t *testing.T) {
	pairs := parsePackets(puzzleinput.Day(13))
	for idx, p := range pairs {
		// fmt.Printf("----------- %d -----------\n", idx)
		res := p.Validate(false)
		if res != realResults[idx] {
			t.Fatalf("error at idx %d. expected %t got %t", idx, realResults[idx], res)
		}
	}
}
