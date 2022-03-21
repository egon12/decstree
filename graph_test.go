package decstree

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToGraph(t *testing.T) {

	input := `
{
	"i": "q1",
	"k": "k",
	"a": [
		{ "i": "a1a", "v": "a", "r": "a" },
		{ "i": "a1b", "v": "b", "r": "b" },
		{ "i": "a1c", "v": "c", "n": {
			"i": "q2",
			"k": "k2",
			"a": [
				{ "i": "a2a", "v": "a", "r": "a" },
				{ "i": "a2b", "v": "b", "r": "b" }
			]
		}}
	]
}
`
	q := &Question{}
	err := json.Unmarshal([]byte(input), q)
	assert.NoError(t, err)

	got := ToDot(q)
	want := `
digraph Q {
	q1 -> { a1a, a1b, a1c }
	a1a -> a
	a1b -> b
	a1c -> q2
	q2 -> { a2a, a2b }
	a2a -> a
	a2b -> b
}
`
	assert.Equal(t, want, got)
}
