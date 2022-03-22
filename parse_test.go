package decstree

import (
	"encoding/json"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	file, err := os.Open("input.json")
	if err != nil {
		t.Fatal("test file 'input.json' not exists")
	}

	q, err := Parse(file)
	if err != nil {
		t.Error(err)
	}

	got, _ := json.Marshal(q)

	want, err := os.ReadFile("output.json")
	if err != nil {
		t.Fatal("test file 'output.json' not exists")
	}

	if string(got) != string(want) {
		t.Error("different")
	}
}
