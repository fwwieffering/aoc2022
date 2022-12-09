package nine

import "testing"

func TestNewTail(t *testing.T) {
	cases := []struct {
		name     string
		tail     *coord
		head     *coord
		expected *coord
	}{
		// same row
		{
			name:     "same row",
			tail:     &coord{x: 1, y: 1},
			head:     &coord{x: 3, y: 1},
			expected: &coord{x: 2, y: 1},
		},
		// same col
		{
			name:     "same col",
			tail:     &coord{x: 1, y: 3},
			head:     &coord{x: 1, y: 1},
			expected: &coord{x: 1, y: 2},
		},
		// diagonal
		{
			name:     "diagonal-1",
			tail:     &coord{x: 1, y: 1},
			head:     &coord{x: 3, y: 2},
			expected: &coord{x: 2, y: 2},
		},
		{
			name:     "diagonal-2",
			tail:     &coord{x: 1, y: 1},
			head:     &coord{x: 2, y: 3},
			expected: &coord{x: 2, y: 2},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {

			newTail := c.tail.getNewTail(c.head)
			if newTail.x != c.expected.x || newTail.y != c.expected.y {
				t.Fatalf("expected {x:%d, y: %d} got {x:%d, y:%d}", c.expected.x, c.expected.y, newTail.x, newTail.y)
			}
		})
	}
}

var testInstructions = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`

func TestProcessInstructions(t *testing.T) {
	count := processInstructions(testInstructions, 2, false, nil)
	if count != 13 {
		t.Fatalf("should be 13 was %d", count)
	}
	biggerRope := processInstructions(testInstructions, 10, true, &coord{x: 6, y: 5})
	t.Fatalf("%d", biggerRope)
}

func TestProcessInstructions2(t *testing.T) {
	biggerIn := `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`
	count := processInstructions(biggerIn, 10, false, nil)
	if count != 36 {
		t.Fatalf("should be 36 was %d", count)
	}
}
