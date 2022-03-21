package decstree

import "testing"

func TestQuestion_Answer(t *testing.T) {
	q := &Question{
		Title: "your gender",
		Key:   "gender",
		Answers: []*Answer{
			{Title: "male", Value: "m", Result: "tall"},
			{Title: "female", Value: "f", Next: &Question{
				Title: "workout alot",
				Key:   "workout",
				Answers: []*Answer{
					{Title: "workout alot", Value: "t", Result: "tall"},
					{Title: "rarely workout", Value: "f", Result: "short"},
				},
			}},
		},
	}

	got, err := q.Answer(map[string]string{
		"gender":  "f",
		"workout": "f",
	})

	t.Error(got)
	t.Error(err)
}
