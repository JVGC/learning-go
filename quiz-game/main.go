package main

import (
	"quiz-game/game"
	"strconv"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "Quiz App",
	Run: func(cmd *cobra.Command, args []string) {
		filename := cmd.Flag("file").Value.String()
		timeLimit, _ := strconv.Atoi(cmd.Flag("limit").Value.String())
		game.RunGame(filename, timeLimit)
	},
}

func main() {
	var csv string
	var limit int
	rootCmd.PersistentFlags().StringVarP(&csv, "file", "f", "problems.csv", "CSV File containing the problems")
	rootCmd.PersistentFlags().IntVarP(&limit, "limit", "l", 30, "Time limit for the quiz in seconds")
	rootCmd.Execute()
}
