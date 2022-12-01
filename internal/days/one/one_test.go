package one

import (
	"bytes"
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	Solve()
}

func TestParseElves(t *testing.T) {
	in := bytes.Split([]byte(`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`), []byte("\n"))
	elves := getMaxElves(in, 3)
	if elves[0] != 24000 {
		fmt.Printf("%+v\n", elves)
		t.Fatalf("%d != 24000", elves[0])
	}
}
