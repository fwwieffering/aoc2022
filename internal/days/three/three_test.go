package three

import (
	"fmt"
	"testing"
)

var examples = map[string]int{
	"vJrwpWtwJgWrhcsFMMfFFhFp":         16,
	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL": 38,
	"PmmdzqPrVvPwwTWBwg":               42,
	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn":   22,
	"ttgJtRGJQctTZtZT":                 20,
	"CrZsJsPPZsGzwwsLwLmpwMDw":         19,
}

var groupPrios = []map[string]int{
	{"r": 18},
	{"Z": 52},
}

func TestPriority(t *testing.T) {
	items := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	for idx, i := range items {
		p := priority(i)
		if p != idx+1 {
			t.Fatalf("%s priority should be %d was %d", string(i), idx+1, p)
		}
	}
}

func TestParseRussack(t *testing.T) {
	var groupOverlaps map[string]int
	bagNum := 1
	for s, p := range examples {
		var overlapPrio int
		overlapPrio, groupOverlaps = parseRussack([]byte(s), groupOverlaps)
		fmt.Printf("%d %+v\n", bagNum, groupOverlaps)
		if overlapPrio != p {
			t.Fatalf("%s expected %d got %d", s, p, overlapPrio)
		}
		if bagNum > 0 && bagNum%3 == 0 {
			expectedOverlaps := groupPrios[bagNum/3-1]
			if len(groupOverlaps) != 1 {
				t.Fatalf("%d expected 1 group overlap but got %+v", bagNum, groupOverlaps)
			}
			for k, v := range expectedOverlaps {
				actual := groupOverlaps[k]
				if actual != v {
					t.Fatalf("%d %+v != %+v", bagNum, groupOverlaps, expectedOverlaps)
				}
			}
			groupOverlaps = nil
		}
		bagNum++
	}
}
