package decstree

import (
	"github.com/egon12/cols"
)

func ToDot(q *Question) string {
	q.Answers

}

func questionToDot(q *Question) string {
	return fmt.Sprtinf("%s -> { %s }",
		q.ID,
		cols.JoinString(q.Answers, func(a Answer) string { return a.ID }),
	)

}
