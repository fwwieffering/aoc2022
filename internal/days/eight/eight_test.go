package eight

import "testing"

var sample = []byte(`30373
25512
65332
33549
35390
`)

func TestCountVisibleTrees(t *testing.T) {

	res, _ := countVisibleTrees(intArrayrIze(sample))
	if res != 21 {
		t.Fatalf("expected %d visible but got %d", 21, res)
	}
}
