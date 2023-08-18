package game

import (
	"context"
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

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(g.Limit)*time.Second)
	defer cancelFunc()

	correctAnswers, _ := g.askQuestions(ctx)

	fmt.Printf("You scored %d out of %d.\n", correctAnswers, len(g.Problems))
}

func (g *Game) askQuestions(ctx context.Context) (int, error) {
	var userAnswer string
	var correctAnswers int
	for index, problem := range g.Problems {
		select {
		case <-ctx.Done():
			return correctAnswers, ctx.Err()
		default:
			fmt.Printf("Problem #%d: %s = ", index+1, problem.Question)
			fmt.Scan(&userAnswer)

			if userAnswer == problem.Answer {
				correctAnswers++
			}
		}

	}

	return correctAnswers, nil
}
