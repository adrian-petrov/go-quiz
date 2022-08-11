package quiz

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	qst "github.com/adrian-petrov/go-quiz/question"
)

const QuestionPrompt = "Question number"

type Quiz struct {
	score        int
	currQuestion int
	questions    []qst.Question
}

func NewQuiz(filePath string) *Quiz {
	questions := qst.ReadFile(filePath)

	return &Quiz{score: 0, currQuestion: 0, questions: questions}
}

func (q *Quiz) incrementScore() {
	q.score++
}

func (q *Quiz) moveToNextQuestion() {
	q.currQuestion++
}

func (q Quiz) validateAnswer(answer string) (bool, error) {
	intAnswer, err := strconv.Atoi(answer)
	if err != nil {
		return false, err
	}

	curr := q.questions[q.currQuestion]

	isCorrect := intAnswer == curr.Result()
	return isCorrect, nil
}

func (q Quiz) readInput(reader *bufio.Reader) bool {
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.Trim(input, "\n")
	isCorrect, err := q.validateAnswer(input)
	if err != nil {
		fmt.Println("Invalid input. Please input a number")
		q.readInput(reader)
	}

	return isCorrect
}

func (q Quiz) Run() {
	for i, qs := range q.questions {
		fmt.Printf("%s %d: %s\n", QuestionPrompt, i+1, qs.Operation())
		reader := bufio.NewReader(os.Stdin)

		if isCorrect := q.readInput(reader); isCorrect {
			q.incrementScore()
		}
		q.moveToNextQuestion()
	}

	fmt.Printf("Correct answers: %d out of %d questions", q.score, len(q.questions))
}
