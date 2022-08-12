package quiz

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	qst "github.com/adrian-petrov/go-quiz/question"
)

const QuestionPrompt = "Question number"

type Quiz struct {
	score        int
	currQuestion int
	questions    []qst.Question
}

func NewQuiz() *Quiz {
	return &Quiz{score: 0, currQuestion: 0}
}

func (q Quiz) Run() {
	fileName, timeLimit := q.readFlags()
	questions := qst.ReadFile(&fileName)
	q.SetQuestions(questions)
	timer := q.startTimer(&timeLimit)

problemLoop:
	for i, qs := range q.questions {
		fmt.Printf("%s %d: %s\n", QuestionPrompt, i+1, qs.Operation())
		answerCh := make(chan bool)

		go func() {
			reader := bufio.NewReader(os.Stdin)
			var (
				input int
				err   error
			)

			for {
				input, err = q.readAndSanitiseInput(reader)
				if err != nil {
					fmt.Println("Invalid input. Please input a number")
				} else {
					break
				}
			}

			isCorrect := q.validateAnswer(input)
			answerCh <- isCorrect
		}()

		select {
		case <-timer.C:
			break problemLoop
		case isCorrect := <-answerCh:
			if isCorrect {
				q.incrementScore()
			}
		}
		q.moveToNextQuestion()
	}
	fmt.Printf("Correct answers: %d out of %d questions", q.score, len(q.questions))
}

func (q *Quiz) SetQuestions(questions []qst.Question) {
	q.questions = questions
}

func (q Quiz) readAndSanitiseInput(reader *bufio.Reader) (int, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.Trim(input, "\r\n")

	sanitised, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return sanitised, nil
}

func (q *Quiz) incrementScore() {
	q.score++
}

func (q *Quiz) moveToNextQuestion() {
	q.currQuestion++
}

func (q Quiz) validateAnswer(answer int) bool {
	curr := q.questions[q.currQuestion]
	return answer == curr.Result()
}

func (q Quiz) readFlags() (string, int) {
	filenameFlag := flag.String(
		"csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimitFlag := flag.Int("limit", 30, "an int value to determine the quiz time limit in seconds")
	flag.Parse()

	return *filenameFlag, *timeLimitFlag
}

func (q Quiz) startTimer(limit *int) *time.Timer {
	timer := time.NewTimer(time.Duration(*limit) * time.Second)
	return timer
}
