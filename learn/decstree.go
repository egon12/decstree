package learn

import (
	"github.com/egon12/cols"
	"github.com/egon12/decstree"
)

func ToQuestion(data Data) *decstree.Question {
	fieldName := NextFieldName(data)
	fieldContents := data.Rows.GetDataFromField(fieldName)
	uniqueFieldContents := cols.Unique(fieldContents)
	answers := cols.Map(uniqueFieldContents, func(c FieldContent) *decstree.Answer {
		answer := &decstree.Answer{
			Value: string(c),
		}

		rows := data.Rows.FilterByContent(fieldName, c)
		if len(rows) > 0 && ShouldReturnResult(rows) {
			answer.Result = string(rows[0].Result)
			return answer
		}

		header := cols.Filter(data.Header, func(f FieldName) bool {
			return f != fieldName
		})

		answer.Next = ToQuestion(Data{Rows: rows, Header: header})

		return answer
	})

	return &decstree.Question{
		Key:     string(fieldName),
		Answers: answers,
	}
}
