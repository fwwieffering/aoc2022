package two

import (
	"bytes"
	"fmt"

	puzzleinput "github.com/fwwieffering/aoc2022/internal/puzzle-input"
)

var winningCombos = map[string]string{
	"A": "Y",
	"B": "Z",
	"C": "X",
}

var drawMap = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

var lossMap = map[string]string{
	"A": "Z",
	"B": "X",
	"C": "Y",
}

var scores = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func scoreGamePart1(gamebytes []byte) int {
	score := 0
	// without modification (b-a)%3 == 0 for win, one for loss, two for draw
	// so if you -1 we get 1 for draw, 2 for win, 0 for loss which is what we want
	diff := int(gamebytes[2]-gamebytes[0]-1) % 3
	score += (diff) * 3
	// playing scores can also be discovered by the byte val % 3
	return score + int(gamebytes[2]-64)%23

}

var wrapper = []int{1, 2, 3}

func scoreGamePart2(gamebytes []byte) int {
	score := 0
	theirPlay := int(gamebytes[0] - 64)
	// 1 = loss, 2 = draw, 3 = win
	result := int(gamebytes[2]-64)%23 - 1
	score += result * 3

	// I'm SURE there is a way to do this with just math but I gave up with the tinkering and used an index
	playIdx := (theirPlay - 1 + result - 1) % 3
	if playIdx < 0 {
		playIdx = 3 + playIdx
	}
	playScore := wrapper[playIdx]
	score += playScore
	return score
}

func Solve() error {
	input := puzzleinput.Day(2)

	games := bytes.Split(input, []byte("\n"))
	part1Score := 0
	part2Score := 0
	for _, game := range games {
		if len(game) > 0 {
			part1Score += scoreGamePart1(game)
			part2Score += scoreGamePart2(game)
		}
	}

	fmt.Printf("Part 1: total score %d\n", part1Score)
	fmt.Printf("Part 2: total score %d\n", part2Score)

	return nil
}
