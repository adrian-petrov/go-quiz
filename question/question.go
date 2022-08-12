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

func (qs *Question) Operation() string {
	return qs.operation
}

func (qs *Question) Result() int {
	return qs.result
}

func (qs *Question) Mathify() int {
	eval := goval.NewEvaluator()
	result, err := eval.Evaluate(qs.operation, nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	return result.(int)
}

func ReadFile(filePath *string) []Question {
	f, err := os.Open(*filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	var questions []Question

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

		questions = append(questions, Question{
			operation: questionArr[0],
			result:    intResult,
		})
	}
	return questions
}
