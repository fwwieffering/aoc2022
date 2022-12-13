package twelve

import (
	"bytes"
	"testing"
)

var example = []byte(`Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`)

func TestGetShortestPath(t *testing.T) {
	in := bytes.Split(example, []byte("\n"))

	bd := shortestPathBidirectional(in, coord{row: 0, col: 0}, coord{row: 2, col: 5})
	if bd != 31 {
		t.Fatalf("%d should be 31", bd)
	}

}
