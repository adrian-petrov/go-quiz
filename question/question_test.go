package question

import (
	"testing"
)

func TestQuestion(t *testing.T) {
	t.Run("it should have 12 lines", func(t *testing.T) {
		got := len(ReadFile("../problems.csv"))
		want := 12

		if got != want {
			t.Errorf("got %d, but want %d", got, want)
		}
	})

	t.Run("operation should equal the result", func(t *testing.T) {
		questions := ReadFile("../problems.csv")

		for _, qst := range questions {
			opResult := qst.Mathify()
			if opResult != qst.result {
				t.Errorf("got %d, but want %d", opResult, qst.result)
			}
		}
	})
}
