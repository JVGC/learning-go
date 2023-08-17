package game

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type problem struct {
	Question string
	Answer   int
}

func readProblems(file *os.File) []problem {
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()

	if err != nil {
		log.Fatalf(fmt.Sprintf("Failed to parse the provided CSV file: %s\n", err))
	}

	var problems []problem

	for _, line := range data {

		intAnswer, _ := strconv.Atoi(line[1])
		problems = append(problems, problem{
			Question: line[0],
			Answer:   intAnswer,
		})
	}

	return problems
}

func RunGame(fileName string, limit int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Failed to open the CSV file: %s\n", err))
	}

	defer file.Close()

	problems := readProblems(file)

	var userAnswer int
	var correctAnswers int

	for index, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", index+1, problem.Question)
		fmt.Scan(&userAnswer)

		if userAnswer == problem.Answer {
			correctAnswers += 1
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correctAnswers, len(problems))

}
