package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello, world")
	// read the csv file
	// for each line print a question
	// register user input
	// update score
	// when reached the end output the final score

	type QuizQuestion struct {
		question string
		answer   string
	}

	var result []QuizQuestion

	f, err := os.Open("problems.csv")
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

		var tempArr []string

		for _, val := range line {
			tempArr = append(tempArr, val)
		}

		result = append(result, QuizQuestion{question: tempArr[0], answer: tempArr[1]})
	}

	for _, line := range result {
		fmt.Printf("The line is %#v", line)
	}
}
