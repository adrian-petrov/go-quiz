package main

import (
	q "github.com/adrian-petrov/go-quiz/quiz"
)

func main() {
	quiz := q.NewQuiz("problems.csv")
	quiz.Run()
}
