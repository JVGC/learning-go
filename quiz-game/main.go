package main

import (
	"quiz-game/cli"
	"quiz-game/game"
	"strconv"
)

func startGame(shell *cli.CLI) {
	filePath := shell.GetFlagValue("file")
	timeLimit, _ := strconv.Atoi(shell.GetFlagValue("limit"))
	shuffle := shell.GetFlagValue("shuffle") == "true"

	game := game.Game{Limit: timeLimit}
	game.ReadProblemsFromCSVFile(filePath)
	if shuffle {
		game.ShuffleProblems()
	}

	game.Run()
}

func main() {
	shell := cli.CLI{}

	shell.AddCommand("", "Quiz App. Answer the questions as fast as you can.", startGame)

	fileFlag := cli.StringFlag{Name: "file", ShortName: "f", DefaultValue: "problems.csv", Usage: "Path for a CSV File containing the problems"}
	limitFlag := cli.IntFlag{Name: "limit", ShortName: "l", DefaultValue: 30, Usage: "Time limit for the quiz in seconds"}
	shuffleFlag := cli.BoolFlag{Name: "shuffle", ShortName: "s", DefaultValue: false, Usage: "Shuffle the problems"}

	shell.AddStringFlag(&fileFlag)
	shell.AddIntFlag(&limitFlag)
	shell.AddBoolFlag(&shuffleFlag)

	shell.Run()
}
