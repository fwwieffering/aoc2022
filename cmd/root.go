package cmd

import (
	"errors"

	one "github.com/fwwieffering/aoc2022/internal/days/one"
	"github.com/spf13/cobra"
)

var dayFuncs = map[int]func() error{
	1: one.Solve,
}

var rootCmd = &cobra.Command{
	Use:   "advent",
	Short: "Run advent of code days",
	RunE: func(cmd *cobra.Command, args []string) error {
		if dayInput == 0 || dayInput > 31 {
			return errors.New("day must be between 1-31")
		}
		return dayFuncs[dayInput]()
	},
}

var dayInput int

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&dayInput, "day", "d", 0, "day to run")
}
