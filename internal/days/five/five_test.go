package five

import (
	"bytes"
	"testing"
)

var testInput = []byte(`    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`)

func TestParseCrates(t *testing.T) {
	splitInput := bytes.Split(testInput, []byte("\n"))

	var crates []*ll
	for i := 0; i < 3; i++ {
		crates = parseCrates(crates, splitInput[i])
	}
	expectedStacks := [][]string{{"N", "Z"}, {"D", "C", "M"}, {"P"}}
	for idx, crate := range crates {
		expected := expectedStacks[idx]
		cur := crate.head
		curIdx := 0
		for i := 0; i < len(expected); i++ {
			if cur == nil {
				t.Fatalf("%+v index %d did not match expected %s != actual nil", expected, i, expected[i])
			}
			if cur.val != expected[i] {
				t.Fatalf("%+v index %d did not match expected %s != actual %s", expected, i, expected[i], cur.val)
			}
			cur = cur.next
			curIdx++
		}
	}
}

func TestProcessStacks(t *testing.T) {
	stacks1, stacks2 := processStacks(testInput)
	expectedHeads1 := "CMZ"
	actualHeads1 := ""
	for _, s := range stacks1 {
		actualHeads1 += s.head.val
	}
	if expectedHeads1 != actualHeads1 {
		t.Fatalf("2: expected %s != actual %s", expectedHeads1, actualHeads1)
	}

	expectedHeads2 := "MCD"
	actualHeads2 := ""
	for _, s := range stacks2 {
		actualHeads2 += s.head.val
	}
	if expectedHeads2 != actualHeads2 {
		t.Fatalf("1: expected %s != actual %s", expectedHeads2, actualHeads2)
	}
}
