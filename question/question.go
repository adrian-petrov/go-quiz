package question

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/maja42/goval"
)

type Question struct {
	operation string
	result    int
}

type QuestionsReader struct {
	questions []Question
}

func (qr *QuestionsReader) Read() {
	f, err := os.Open("../problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		var questionArr []string
		questionArr = append(questionArr, line...)

		intResult, err := strconv.Atoi(questionArr[1])
		if err != nil {
			log.Fatal(err)
		}

		qr.questions = append(qr.questions, Question{
			operation: questionArr[0],
			result:    intResult,
		})
	}
}

func (q *Question) Mathify() int {
	eval := goval.NewEvaluator()
	result, err := eval.Evaluate(q.operation, nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	return result.(int)
}
