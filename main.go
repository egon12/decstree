package decstree

import (
	"fmt"
	"strings"
)

type (
	// Question represent one question that can be answered by data
	// We answer question by looking Data and find it's key.
	// Question only can be answered by choosing the Answer that have
	// correct value
	Question struct {
		ID      string  `json:"i,omitempty"`
		Label   string  `json:"l"`
		Key     string  `json:"k"`
		Answers Answers `json:"a"`
	}

	// Answer are represent Answer that can be choosed when we answer
	// the Question. You can think that Question like multiple-choice question
	// and the answer are the one of the choice.
	// Answer can have Result, or Next Question. If the answer that we choose,
	// contains Next, then we can ask next Question. If it doesn't contain next,
	// we can mark Result as the final Answer, and return the Result
	Answer struct {
		ID     string    `json:"i,omitempty"`
		Label  string    `json:"l"`
		Value  string    `json:"v"`
		Next   *Question `json:"n,omitempty"`
		Result string    `json:"r,omitempty"`
	}

	// Answers are type Alias for []*Answer created so we can add function to it
	Answers []*Answer

	// Data are needed to answer the question. You can access the data by
	// using key that is in question. And match the value with value in Answer
	Data map[string]string

	// Traces are collection of ID that have been visit by Data and Answer.
	// It will contains ID of the first question until the last result
	// you can use package draw to see what question and answer that has been answerd
	// by data
	Traces []string
)

// Answer is the fastest way to get the result, based on the data
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

// AnswerWithTrace will answer all the questions until we gate result, and then also return the traces
// that can be used in package draw
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

// String implementation of stringer
func (a Answers) String() string {
	var res []string
	for _, v := range a {
		if v != nil {
			res = append(res, v.Label)
		}
	}
	return strings.Join(res, ",")
}
