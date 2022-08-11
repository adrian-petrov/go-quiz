package question

import (
	"testing"
)

func TestQuestion(t *testing.T) {
	t.Run("it should have 12 lines", func(t *testing.T) {
		reader := new(QuestionsReader)
		reader.Read()
		got := len(reader.questions)
		want := 12

		if got != want {
			t.Errorf("got %d, but want %d", got, want)
		}
	})

	t.Run("operation should equal the result", func(t *testing.T) {
		reader := new(QuestionsReader)
		reader.Read()

		for _, qst := range reader.questions {
			opResult := qst.Mathify()
			if opResult != qst.result {
				t.Errorf("got %d, but want %d", opResult, qst.result)
			}
		}
	})
}
