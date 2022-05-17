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
	"l": "level",
	"k": "level",
	"a": [
		{ "i": "a1_4", "v": "4", "l": "4", "r": "alber" },
		{ "i": "a1_5", "v": "5", "l": "5", "r": "bobby" },
		{ "i": "a1_6", "v": "6", "l": "6", "n": {
			"i": "job_1",
			"l": "job",
			"k": "job",
			"a": [
				{ "i": "a2_tech", "l": "tech", "v": "tech", "r": "chris" },
				{ "i": "a2_pm",   "l": "pm",   "v": "pm", "r": "dan" },
				{ "i": "a2_lead", "l": "lead", "v": "lead", "r": "eugine" }
			]
		}},
		{ "i": "a1_7", "l": "7", "v": "7", "n": {
			"i": "job_2",
			"l": "job",
			"k": "job",
			"a": [
				{ "i": "a3_tech", "l": "tech", "v": "tech", "r": "farah" },
				{ "i": "a3_pm",   "l": "pm",   "v": "pm", "r": "dan" },
				{ "i": "a3_lead", "l": "lead", "v": "lead", "r": "harlye" }
			]

		}},
		{ "i": "a1_8", "l": "8", "v": "8", "r": "myvp" }
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

func TestSVG2(t *testing.T) {
	f, err := os.Open("../output.json")
	assert.NoError(t, err)

	q, err := decstree.Parse(f)
	assert.NoError(t, err)

	traces, _, _ := q.AnswerWithTrace(map[string]string{
		"Weather":               "Windy",
		"Parental Availibility": "No",
		"Wealthy":               "Poor",
	})
	f, err = os.Create("qq2.svg")
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

	decstree.SetLabel(q)

	ids, _, err := q.AnswerWithTrace(data)

	got := ToDotWithTrace(q, ids)
	want := `digraph Q {
	node [style="rounded,filled"]
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
	"level" [fontname=helvetica shape=box style="filled" fillcolor=steelblue label="level"]
	"a1_4" [fontname=helvetica shape=box label="4"]
	"a1_5" [fontname=helvetica shape=box label="5"]
	"a1_6" [fontname=helvetica shape=box style="rounded,filled" fillcolor=steelblue label="6"]
	"job_1" [fontname=helvetica shape=box style="filled" fillcolor=steelblue label="job"]
	"a2_tech" [fontname=helvetica shape=box label="tech"]
	"a2_pm" [fontname=helvetica shape=box style="rounded,filled" fillcolor=steelblue label="pm"]
	"a2_lead" [fontname=helvetica shape=box label="lead"]
	"a1_7" [fontname=helvetica shape=box label="7"]
	"job_2" [fontname=helvetica shape=box style="filled" label="job"]
	"a3_tech" [fontname=helvetica shape=box label="tech"]
	"a3_pm" [fontname=helvetica shape=box label="pm"]
	"a3_lead" [fontname=helvetica shape=box label="lead"]
	"a1_8" [fontname=helvetica shape=box label="8"]
	"alber" [fontname=helvetica shape=oval]
	"bobby" [fontname=helvetica shape=oval]
	"chris" [fontname=helvetica shape=oval]
	"dan" [fontname=helvetica shape=oval, style="rounded,filled", fillcolor=steelblue]
	"eugine" [fontname=helvetica shape=oval]
	"farah" [fontname=helvetica shape=oval]
	"harlye" [fontname=helvetica shape=oval]
	"myvp" [fontname=helvetica shape=oval]
}`
	//err = os.WriteFile("qq.gv", []byte(got), os.ModePerm)
	assert.Equal(t, want, got)
}
