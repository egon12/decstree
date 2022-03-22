package decstree

import (
	"fmt"
	"strings"
)

type (
	Question struct {
		ID      string  `json:"i,omitempty"`
		Title   string  `json:"t"`
		Key     string  `json:"k"`
		Answers Answers `json:"a"`
	}

	Answers []*Answer

	Answer struct {
		ID     string    `json:"i,omitempty"`
		Title  string    `json:"t"`
		Value  string    `json:"v"`
		Next   *Question `json:"n,omitempty"`
		Result string    `json:"r,omitempty"`
	}

	Data map[string]string

	Traces []string
)

func (q *Question) Answer(data Data) (string, error) {
	val := data[q.Key]

	for _, a := range q.Answers {
		if a.Value != val {
			continue
		}

		if a.Next != nil {
			return a.Next.Answer(data)
		}

		return a.Result, nil
	}

	return "", fmt.Errorf("data '%s' in '%s' don't match with any '%s'", val, q.Key, q.Answers)
}

func (q *Question) AnswerWithTrace(data Data) (ids Traces, answer string, err error) {
	val := data[q.Key]

	for _, a := range q.Answers {
		if a.Value != val {
			continue
		}

		if a.Next != nil {
			ids, answer, err = a.Next.AnswerWithTrace(data)
			ids = append(ids, a.ID, q.ID)
			return ids, answer, err
		}

		return []string{a.ID, q.ID, a.Result}, a.Result, nil
	}

	return []string{q.ID}, "", fmt.Errorf("data '%s' in '%s' cannot been answer with '%s'", val, q.Key, q.Answers)

}

func (a Answers) String() string {
	var res []string
	for _, v := range a {
		if v != nil {
			res = append(res, v.Title)
		}
	}
	return strings.Join(res, ",")
}
