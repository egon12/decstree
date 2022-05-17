package learn

import (
	"encoding/json"
	"os"
	"testing"
)

func TestToQuestion(t *testing.T) {
	q := ToQuestion(dataTest)
	b, err := json.Marshal(q)
	if err != nil {
		t.Error(err)
	}
	os.WriteFile("tmp.json", b, os.ModePerm)
}
