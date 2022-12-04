package four

import "testing"

var examples = map[string][]bool{
	"2-4,6-8": {false, false},
	"2-3,4-5": {false, false},
	"5-7,7-9": {false, true},
	"2-8,3-7": {true, true},
	"6-6,4-6": {true, true},
	"2-6,4-8": {false, true},
}

func TestSectionsOverlap(t *testing.T) {
	for input, shouldOverlap := range examples {
		section1, section2 := parseSections(input)
		full := sectionsFullyOverlap(section1, section2)
		partial := sectionsPartiallyOverlap(section1, section2)
		if full != shouldOverlap[0] {
			t.Fatalf("%s overlap should be %t was %t", input, shouldOverlap[0], full)
		}
		if partial != shouldOverlap[1] {
			t.Fatalf("%s partial overlap should be %t was %t", input, shouldOverlap[1], partial)
		}
	}
}
