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

var expected = []interface{}{
	[]interface{}{1, 1, 3, 1, 1},
	[]interface{}{1, 1, 5, 1, 1},
	[]interface{}{[]interface{}{1}, []interface{}{2, 3, 4}},
	[]interface{}{[]interface{}{1}, 4},
	[]interface{}{9},
	[]interface{}{[]interface{}{8, 7, 6}},
	[]interface{}{[]interface{}{4, 4}, 4, 4},
	[]interface{}{[]interface{}{4, 4}, 4, 4, 4},
	[]interface{}{7, 7, 7, 7},
	[]interface{}{7, 7, 7},
	[]interface{}{},
	[]interface{}{3},
	[]interface{}{[]interface{}{[]interface{}{}}},
	[]interface{}{[]interface{}{}},
	[]interface{}{1, []interface{}{2, []interface{}{3, []interface{}{4, []interface{}{5, 6, 7}}}}, 8, 9},
	[]interface{}{1, []interface{}{2, []interface{}{3, []interface{}{4, []interface{}{5, 6, 0}}}}, 8, 9},
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
	validationIdx := 0
	for idx, exp := range expected {
		if !cmp.Equal(exp, result[idx]) {
			fmt.Printf("actual: %+v\n", result[idx])
			t.Fatalf("idx: %d\n%s", idx, cmp.Diff(exp, result[idx]))
		}
		if (idx+1)%2 == 0 {
			isValid := ValidatePackets(result[idx-1], result[idx], false)
			fmt.Printf("left: %+v\nright: %+v\n", result[idx-1], result[idx])
			if isValid != validationExpectations[validationIdx] {
				t.Fatalf("idx %d expected %t but got %t", idx, validationExpectations[validationIdx], isValid)
			}
			validationIdx++
		}
	}

	// for idx, expectValid := range validationExpectations {
	// 	fmt.Printf("----------- %d -----------\n", idx)

	// 	isValid := ValidatePackets(result[idx*2-1], result[idx*2], false)
	// 	if isValid != expectValid {
	// 		t.Fatalf("idx %d expected %t but got %t", idx, expectValid, isValid)
	// 	}

	// }
	p1 := part1(result)
	if p1 != 13 {
		t.Fatalf("expected 13 got %d", p1)
	}

	p2 := part2(result)
	if p2 != 140 {
		t.Fatalf("expected 140 got %d", p2)
	}
}

// var realResults = []bool{true, false, false, false, false, false, true, true, false, true, true, false, true, false, true, false, true, false, false, true, true, false, true, false, true, false, true, false, true, true, false, true, false, true, false, false, true, false, false, true, true, true, true, true, false, false, false, true, false, true, true, false, true, true, true, false, false, true, true, true, true, false, true, false, true, false, false, true, false, true, true, false, true, true, true, true, false, false, true, true, true, false, false, true, true, true, true, true, false, false, false, false, false, true, true, true, false, false, false, true, false, true, true, true, false, true, true, true, true, true, true, true, false, true, false, false, false, true, false, true, false, true, true, true, true, false, true, false, false, true, false, false, true, true, false, true, false, true, false, true, false, true, false, false, false, true, false, true, false, false}

// func TestRealResults(t *testing.T) {
// 	pairs := parsePackets(puzzleinput.Day(13))
// 	for idx, p := range pairs {
// 		// fmt.Printf("----------- %d -----------\n", idx)
// 		res := p.Validate(false)
// 		if res != realResults[idx] {
// 			t.Fatalf("error at idx %d. expected %t got %t", idx, realResults[idx], res)
// 		}
// 	}
// }
