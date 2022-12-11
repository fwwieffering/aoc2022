package eleven

import (
	"testing"
)

var example = `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1
`

func TestParseMonkey(t *testing.T) {
	monks := parseMonkeys(example)
	if monks[0].Idx != 0 {
		t.Fatalf("monkey0 idx should be 0, was %d", monks[0].Idx)
	}
	if len(monks[0].Items) != 2 {
		t.Fatalf("missing items %+v", monks[0].Items)
	}
	expected := []int64{79, 98}
	for i := 0; i < 2; i++ {
		if expected[i] != monks[0].Items[i] {
			t.Fatalf("mismatch at idx %d: expected %d != actual %d", i, expected[i], monks[0].Items[i])
		}
	}
	for _, in := range []int64{1, 2} {
		if monks[0].Operation(in) != in*19 {
			t.Fatalf("%d != %d * 19", in*19, in)
		}
	}
	if monks[0].MonkeyTest(4) != 3 {
		t.Fatalf("4 %% 23 should be false")
	}
	if monks[0].MonkeyTest(23) != 2 {
		t.Fatalf("4 %% 23 should be true is %+v", monks[0].MonkeyTest(23))
	}
}

func TestMonkeyBiz(t *testing.T) {
	monkeys := parseMonkeys(example)
	res := doMonkeyBusiness(monkeys, 20, true)

	if res != 10605 {
		t.Fatalf("should be 10605 was %d", res)
	}
	monkeys2 := parseMonkeys(example)
	res = doMonkeyBusiness(monkeys2, 10000, false)
	if res != 2713310158 {
		t.Fatalf("should be 2713310158 was %d", res)
	}
}
