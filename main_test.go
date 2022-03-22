package decstree

import (
	"testing"
)

func TestQuestion_Answer(t *testing.T) {
	q := &Question{
		Label: "your gender",
		Key:   "gender",
		Answers: []*Answer{
			{Label: "male", Value: "m", Result: "tall"},
			{Label: "female", Value: "f", Next: &Question{
				Label: "workout alot",
				Key:   "workout",
				Answers: []*Answer{
					{Label: "workout alot", Value: "t", Result: "tall"},
					{Label: "rarely workout", Value: "f", Result: "short"},
				},
			}},
		},
	}

	got, err := q.Answer(map[string]string{
		"gender":  "f",
		"workout": "f",
	})

	if err != nil {
		t.Errorf("want no error got %v", err)
	}

	if got != "short" {
		t.Errorf("want result to 'short' got %v", got)
	}
}
