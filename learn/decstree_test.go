package learn

import (
	"encoding/json"
	"os"
	"testing"
)

func TestToQuestion(t *testing.T) {
	q := ToQuestion(dataTest)
	b, err := json.Marshal(q)
	os.WriteFile("tmp.json", b, os.ModePerm)
}
