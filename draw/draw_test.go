package draw

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/egon12/decstree"
	"github.com/stretchr/testify/assert"
)

const question = `
{
	"i": "level",
	"k": "level",
	"a": [
		{ "i": "a1_4", "v": "4", "r": "alber" },
		{ "i": "a1_5", "v": "5", "r": "bobby" },
		{ "i": "a1_6", "v": "6", "n": {
			"i": "job_1",
			"k": "job",
			"a": [
				{ "i": "a2_tech", "v": "tech", "r": "chris" },
				{ "i": "a2_pm", "v": "pm", "r": "dan" },
				{ "i": "a2_lead", "v": "lead", "r": "eugine" }
			]
		}},
		{ "i": "a1_7", "v": "7", "n": {
			"i": "job_2",
			"k": "job",
			"a": [
				{ "i": "a3_tech", "v": "tech", "r": "farah" },
				{ "i": "a3_pm", "v": "pm", "r": "dan" },
				{ "i": "a3_lead", "v": "lead", "r": "harlye" }
			]

		}},
		{ "i": "a1_8", "v": "8", "r": "myvp" }
	]
}`

var data = map[string]string{
	"level": "6",
	"job":   "pm",
}

func TestSVG(t *testing.T) {
	q := &decstree.Question{}
	err := json.Unmarshal([]byte(question), q)
	assert.NoError(t, err)

	traces, _, _ := q.AnswerWithTrace(data)
	f, err := os.Create("qq.svg")
	if err != nil {
		t.Fatal(err)
	}

	SVG(f, q, traces)
	f.Close()
}

func TestToGraph(t *testing.T) {

	q := &decstree.Question{}
	err := json.Unmarshal([]byte(question), q)
	assert.NoError(t, err)

	ids, _, err := q.AnswerWithTrace(data)

	got := ToDotWithTrace(q, ids)
	want := `digraph Q {
	node [style=rounded]
	"level" -> {rank=same; "a1_4","a1_5","a1_6","a1_7","a1_8"}
	"a1_4" -> "alber"
	"a1_5" -> "bobby"
	"a1_6" -> "job_1"
	"job_1" -> {rank=same; "a2_tech","a2_pm","a2_lead"}
	"a2_tech" -> "chris"
	"a2_pm" -> "dan"
	"a2_lead" -> "eugine"
	"a1_7" -> "job_2"
	"job_2" -> {rank=same; "a3_tech","a3_pm","a3_lead"}
	"a3_tech" -> "farah"
	"a3_pm" -> "dan"
	"a3_lead" -> "harlye"
	"a1_8" -> "myvp"
	"level" [shape=diamond, style="rounded,filled", fillcolor=steelblue]
	"a1_4" [shape=box]
	"a1_5" [shape=box]
	"a1_6" [shape=box, style="rounded,filled", fillcolor=steelblue]
	"job_1" [shape=diamond, style="rounded,filled", fillcolor=steelblue]
	"a2_tech" [shape=box]
	"a2_pm" [shape=box, style="rounded,filled", fillcolor=steelblue]
	"a2_lead" [shape=box]
	"a1_7" [shape=box]
	"job_2" [shape=diamond]
	"a3_tech" [shape=box]
	"a3_pm" [shape=box]
	"a3_lead" [shape=box]
	"a1_8" [shape=box]
	"alber" [shape=oval]
	"bobby" [shape=oval]
	"chris" [shape=oval]
	"dan" [shape=oval, style="rounded,filled", fillcolor=steelblue]
	"eugine" [shape=oval]
	"farah" [shape=oval]
	"harlye" [shape=oval]
	"myvp" [shape=oval]
}`
	//err = os.WriteFile("qq.gv", []byte(got), os.ModePerm)
	assert.Equal(t, want, got)
}