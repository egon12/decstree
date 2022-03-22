package decstree

import (
	"encoding/json"
	"io"
	"strconv"
)

// Parse will read the data that has been serialize in json format
func Parse(r io.Reader) (*Question, error) {
	d := json.NewDecoder(r)
	q := &Question{}
	err := d.Decode(q)
	if err != nil {
		return q, err
	}

	SetID(q)
	SetLabel(q)

	return q, err
}

func SetID(q *Question) {
	setQID(q, 1)
}

func setQID(q *Question, usedQID int) int {
	if q.ID == "" {
		q.ID = "q" + strconv.Itoa(usedQID+1)
		usedQID += 1
	}

	for i, a := range q.Answers {
		if a.ID == "" {
			a.ID = q.ID + "a" + strconv.Itoa(i)
		}
		if a.Next != nil {
			usedQID = setQID(a.Next, usedQID)
		}
	}

	return usedQID
}

func SetLabel(q *Question) {
	if q.Label == "" {
		q.Label = q.Key
	}

	for _, a := range q.Answers {
		if a.Label == "" {
			a.Label = a.Value
		}
		if a.Next != nil {
			SetLabel(a.Next)
		}
	}
}
