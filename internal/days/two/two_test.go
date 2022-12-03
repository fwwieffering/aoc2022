package two

import (
	"testing"
)

func TestScoreGame(t *testing.T) {
	cases := []struct {
		in       []byte
		expected int
	}{
		{
			in:       []byte("A Y"),
			expected: 8,
		},
		{
			in:       []byte("B X"),
			expected: 1,
		},
		{
			in:       []byte("C Z"),
			expected: 6,
		},
	}

	for _, c := range cases {
		res := scoreGamePart1(c.in)
		if res != c.expected {
			t.Fatalf("%s should be %d was %d", c.in, c.expected, res)
		}
	}
}

func TestScoreGamePart2(t *testing.T) {
	cases := []struct {
		in       []byte
		expected int
	}{
		{
			in:       []byte("A X"),
			expected: 3,
		},
		{
			in:       []byte("A Y"),
			expected: 4,
		},
		{
			in:       []byte("A Z"),
			expected: 8,
		},
		{
			in:       []byte("B X"),
			expected: 1,
		},
		{
			in:       []byte("C X"),
			expected: 2,
		},
		{
			in:       []byte("C Y"),
			expected: 6,
		},
		{
			in:       []byte("C Z"),
			expected: 7,
		},
	}

	for _, c := range cases {
		res := scoreGamePart2(c.in)
		if res != c.expected {
			t.Fatalf("%s should be %d was %d", c.in, c.expected, res)
		}
	}
}
