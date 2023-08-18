package game

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

type Problem struct {
	Question string
	Answer   string
}

type Game struct {
	Problems []Problem
	Limit    int
}

func (g *Game) ReadProblemsFromCSVFile(filePath string) {
	data, err := ReadCSVFile(filePath)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Failed to parse the provided CSV file: %s\n", err))
	}

	g.Problems = make([]Problem, len(data))

	for i, line := range data {

		g.Problems[i] = Problem{
			Question: line[0],
			Answer:   strings.TrimSpace(line[1]),
		}
	}
}

func (g *Game) ShuffleProblems() {
	rand.Shuffle(len(g.Problems), func(i, j int) {
		g.Problems[i], g.Problems[j] = g.Problems[j], g.Problems[i]
	})
}
func (g *Game) Run() {

	correctAnswers := g.askQuestions()

	fmt.Printf("\nYou scored %d out of %d.\n", correctAnswers, len(g.Problems))
}

func (g *Game) askQuestions() int {

	timer := time.NewTimer(time.Duration(g.Limit) * time.Second)

	correctAnswers := 0
	for index, problem := range g.Problems {
		fmt.Printf("Problem #%d: %s = ", index+1, problem.Question)

		answerChannel := make(chan string)
		go func() {
			var userAnswer string
			fmt.Scan(&userAnswer)
			answerChannel <- userAnswer
		}()

		select {
		case <-timer.C:
			return correctAnswers
		case userAnswer := <-answerChannel:
			if userAnswer == problem.Answer {
				correctAnswers++
			}
		}

	}

	return correctAnswers
}
