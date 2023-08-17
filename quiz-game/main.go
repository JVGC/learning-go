package main

import (
	"quiz-game/game"
	"strconv"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "Quiz App. Answer the questions as fast as you can.",
	Run: func(cmd *cobra.Command, args []string) {
		startGame(cmd, args)
	},
}

func startGame(cmd *cobra.Command, args []string) {
	filePath := cmd.Flag("file").Value.String()
	timeLimit, _ := strconv.Atoi(cmd.Flag("limit").Value.String())
	shuffle := cmd.Flag("shuffle").Value.String() == "true"

	game := game.Game{Limit: timeLimit}
	game.ReadProblemsFromCSVFile(filePath)
	if shuffle {
		game.ShuffleProblems()
	}
	game.Run()
}

func main() {
	var csv string
	var limit int
	rootCmd.PersistentFlags().StringVarP(&csv, "file", "f", "problems.csv", "Path for a CSV File containing the problems")
	rootCmd.PersistentFlags().IntVarP(&limit, "limit", "l", 30, "Time limit for the quiz in seconds")
	rootCmd.PersistentFlags().BoolP("shuffle", "s", false, "Shuffle the problems")
	rootCmd.Execute()
}
