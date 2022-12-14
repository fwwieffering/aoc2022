package fourteen

import (
	"fmt"
	"testing"
)

func TestProcessInput(t *testing.T) {
	var example = []byte(`498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9
`)
	cave := ParseCave(example)

	res := cave.SandFall(false)
	if res != 24 {
		t.Fatalf("should be 24 was %d", res)
	}
	p2 := cave.SandFall(true)
	cave.Draw(true)
	if p2 != 93 {
		fmt.Println(len(cave.sand))
		t.Fatalf("should be 93 was %d", p2)
	}
}
