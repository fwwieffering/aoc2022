package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/fwwieffering/aoc2022/internal/days/eight"
	"github.com/fwwieffering/aoc2022/internal/days/eleven"
	"github.com/fwwieffering/aoc2022/internal/days/five"
	"github.com/fwwieffering/aoc2022/internal/days/four"
	"github.com/fwwieffering/aoc2022/internal/days/fourteen"
	"github.com/fwwieffering/aoc2022/internal/days/nine"
	"github.com/fwwieffering/aoc2022/internal/days/one"
	"github.com/fwwieffering/aoc2022/internal/days/seven"
	"github.com/fwwieffering/aoc2022/internal/days/six"
	"github.com/fwwieffering/aoc2022/internal/days/ten"
	"github.com/fwwieffering/aoc2022/internal/days/thirteen"
	"github.com/fwwieffering/aoc2022/internal/days/three"
	"github.com/fwwieffering/aoc2022/internal/days/twelve"
	"github.com/fwwieffering/aoc2022/internal/days/two"
	"github.com/spf13/cobra"
)

var dayFuncs = map[int]func() error{
	1:  one.Solve,
	2:  two.Solve,
	3:  three.Solve,
	4:  four.Solve,
	5:  five.Solve,
	6:  six.Solve,
	7:  seven.Solve,
	8:  eight.Solve,
	9:  nine.Solve,
	10: ten.Solve,
	11: eleven.Solve,
	12: twelve.Solve,
	13: thirteen.Solve,
	14: fourteen.Solve,
}

var rootCmd = &cobra.Command{
	Use:   "advent",
	Short: "Run advent of code days",
	RunE: func(cmd *cobra.Command, args []string) error {
		if dayInput == 0 || dayInput > 25 {
			return errors.New("day must be between 1-25")
		}
		f, ok := dayFuncs[dayInput]
		if !ok {
			return fmt.Errorf("day %d not yet implemented", dayInput)
		}
		start := time.Now()
		err := f()
		duration := time.Now().Sub(start)
		fmt.Printf("took: %d microseconds\n", duration.Microseconds())
		return err
	},
}

var dayInput int

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&dayInput, "day", "d", 0, "day to run")
}
