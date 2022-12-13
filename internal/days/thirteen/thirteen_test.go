package thirteen

import (
	"fmt"
	"testing"

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

func TestParsePackets(t *testing.T) {
	result := parsePackets(example)
	for idx, exp := range expected {
		fmt.Printf("Left: %+v Right: %+v\n", result[idx].Left, result[idx].Right)
		if !cmp.Equal(exp, result[idx]) {
			t.Fatalf(cmp.Diff(exp, result[idx]))
		}
	}
}
