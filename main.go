package decstree

import (
	"fmt"
	"strings"
)

type (
	Question struct {
		Title   string  `json:"t"`
		Key     string  `json:"k"`
		Answers Answers `json:"a"`
	}

	Answers []*Answer

	Answer struct {
		Title  string    `json:"t"`
		Value  string    `json:"v"`
		Next   *Question `json:"n"`
		Result string    `json:"r"`
	}

	Data map[string]string
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

	return "", fmt.Errorf("data '%s' in '%s' cannot been answer with '%s'", val, q.Key, q.Answers)
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
