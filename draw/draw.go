package draw

import (
	"fmt"
	"io"
	"sort"

	"github.com/egon12/cols"
	"github.com/egon12/decstree"
	"github.com/goccy/go-graphviz"
)

func SVG(w io.Writer, q *decstree.Question, t decstree.Traces) error {
	dotContent := ToDotWithTrace(q, t)
	g, err := graphviz.ParseBytes([]byte(dotContent))
	if err != nil {
		return err
	}
	return graphviz.New().Render(g, graphviz.SVG, w)
}

func ToDotWithTrace(q *decstree.Question, traces []string) string {
	result := NewSet()
	first := "digraph Q {\n\tnode [style=\"rounded,filled\"]\n" + questionToDot(q) + "\n" + printShapeQWithColor(q, traces, result)
	return first + printShapeResult(result, traces) + "\n}"
}

func questionToDot(q *decstree.Question) string {
	return fmt.Sprintf("\t\"%s\" -> {rank=same; %s}\n",
		q.ID,
		cols.JoinString(q.Answers, func(a *decstree.Answer) string { return `"` + a.ID + `"` }),
	) + answersToDot(q.Answers)
}

func answerToDot(a *decstree.Answer) string {
	qRes := ""
	res := a.Result
	if a.Next != nil {
		res = a.Next.ID
		qRes += "\n" + questionToDot(a.Next)
	}
	return fmt.Sprintf("\t\"%s\" -> \"%s\"", a.ID, res) + qRes
}

func answersToDot(a decstree.Answers) string {
	return cols.JoinStringWithSep(
		a,
		answerToDot,
		"\n",
	)
}

func printShapeQWithColor(q *decstree.Question, traces []string, result set) (qa string) {
	res := ""
	if isIn(traces, q.ID) {
		res = fmt.Sprintf("\t\"%s\" [fontname=helvetica shape=box style=\"filled\" fillcolor=steelblue label=\"%s\"]\n", q.ID, q.Label)
	} else {
		res = fmt.Sprintf("\t\"%s\" [fontname=helvetica shape=box style=\"filled\" label=\"%s\"]\n", q.ID, q.Label)
	}
	for _, a := range q.Answers {
		if isIn(traces, a.ID) {
			res += fmt.Sprintf("\t\"%s\" [fontname=helvetica shape=box style=\"rounded,filled\" fillcolor=steelblue label=\"%s\"]\n", a.ID, a.Label)
		} else {
			res += fmt.Sprintf("\t\"%s\" [fontname=helvetica shape=box label=\"%s\"]\n", a.ID, a.Label)
		}
		if a.Next != nil {
			res += printShapeQWithColor(a.Next, traces, result)
		} else {
			result.Add(a.Result)
		}
	}
	return res
}

func printShapeResult(result set, traces []string) string {
	return cols.JoinStringWithSep(result.ToList(), func(r string) string {
		if isIn(traces, r) {
			return fmt.Sprintf("\t\"%s\" [fontname=helvetica shape=oval, style=\"rounded,filled\", fillcolor=steelblue]", r)
		}
		return fmt.Sprintf("\t\"%s\" [fontname=helvetica shape=oval]", r)
	}, "\n")
}

func isIn(haystack []string, needle string) bool {
	return cols.Any(haystack, func(i string) bool {
		return i == needle
	})
}

type set map[string]struct{}

func NewSet() set {
	return make(map[string]struct{})
}

func (s set) Add(i string) {
	s[i] = struct{}{}
}

func (s set) ToList() []string {
	res := make([]string, len(s))
	i := 0
	for k, _ := range s {
		res[i] = k
		i += 1
	}
	sort.Strings(res)
	return res
}
