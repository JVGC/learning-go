package game

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type Problem struct {
	Question string
	Answer   int
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

	for _, line := range data {
		intAnswer, _ := strconv.Atoi(line[1])

		g.Problems = append(g.Problems, Problem{
			Question: line[0],
			Answer:   intAnswer,
		})
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
	var userAnswer int
	var correctAnswers int
	for index, problem := range g.Problems {
		select {
		case <-ctx.Done():
			return correctAnswers, ctx.Err()
		default:
			fmt.Printf("Problem #%d: %s = ", index+1, problem.Question)
			fmt.Scan(&userAnswer)

			if userAnswer == problem.Answer {
				correctAnswers += 1
			}
		}

	}

	return correctAnswers, nil
}
