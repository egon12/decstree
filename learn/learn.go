package learn

import (
	"math"

	"github.com/egon12/cols"
)

type (
	FieldName    string
	FieldContent string
	Result       string

	ID3Processor struct {
		Fields []FieldName

		DB Data
	}

	Data struct {
		Header []FieldName
		Rows   Rows
	}

	Row struct {
		Input  map[FieldName]FieldContent
		Result Result
	}

	Rows []Row
)

func NextFieldName(data Data) FieldName {
	type Gains struct {
		FieldName
		Gain float64
	}

	g := cols.Map(data.Header, func(field FieldName) Gains {
		return Gains{
			Gain:      Gain(data.Rows, field),
			FieldName: field,
		}
	})

	sg := cols.MaxStruct(g, func(g Gains) float64 {
		return g.Gain
	})

	return sg.FieldName
}

func ShouldReturnResult(rows Rows) bool {
	return rows.Entropy() == 0.0
}

// ES get entropy from S (Result)
func ES(rows Rows) float64 {
	return rows.Entropy()
}

// Gain get information Gain from field
func Gain(rows Rows, field FieldName) float64 {
	fData := rows.GetDataFromField(field)
	fCount := cols.CountBy(fData, itself[FieldContent])

	var res float64 = 0.0
	for fieldContent, count := range fCount {
		entropy := ESContent(rows, field, fieldContent)
		probability := float64(count) / float64(len(rows))
		res += probability * entropy
	}
	return rows.Entropy() - res
}

// ESContent get entropy from S(Result) and FieldContent
func ESContent(rows Rows, field FieldName, content FieldContent) float64 {
	return rows.FilterByContent(field, content).Entropy()
}

func (i *ID3Processor) Load(db Data) {
	i.DB = db

	i.Fields = db.Header
}

func (r Rows) Entropy() float64 {
	rowCount := len(r)
	results := cols.Map(r, getResultFromRow)
	resultCounts := cols.CountBy(results, itself[Result])

	var res float64 = 0.0
	for _, c := range resultCounts {
		res += singleEntropy(float64(c), float64(rowCount))
	}
	return res * -1
}

func (r Rows) GetDataFromField(field FieldName) []FieldContent {
	return cols.Map(r, func(row Row) FieldContent {
		return FieldContent(row.Input[field])
	})
}

func (r Rows) FilterByContent(field FieldName, content FieldContent) Rows {
	filterFunc := isRowFieldHasContent(field, content)
	return cols.Filter(r, filterFunc)
}

func itself[T any](t T) T {
	return t
}

func singleEntropy(count, total float64) float64 {
	p := count / total
	return p * math.Log2(p)
}

func getResultFromRow(r Row) Result {
	return r.Result
}

func isRowFieldHasContent(f FieldName, c FieldContent) func(Row) bool {
	return func(r Row) bool {
		return r.Input[f] == c
	}
}
