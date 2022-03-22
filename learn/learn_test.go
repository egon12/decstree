package learn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var dataTest Data = Data{
	Header: []FieldName{"Weather", "Parental Availibility", "Wealthy"},
	Rows: []Row{
		{
			Input: map[FieldName]FieldContent{
				"Weather":               "Sunny",
				"Parental Availibility": "Yes",
				"Wealthy":               "Rich",
			},
			Result: "C",
		},
		{
			Input: map[FieldName]FieldContent{
				"Weather":               "Sunny",
				"Parental Availibility": "No",
				"Wealthy":               "Rich",
			},
			Result: "T",
		},
		{
			Input: map[FieldName]FieldContent{
				"Weather":               "Windy",
				"Parental Availibility": "Yes",
				"Wealthy":               "Rich",
			},
			Result: "C",
		},
		{
			Input: map[FieldName]FieldContent{
				"Weather":               "Rainy",
				"Parental Availibility": "Yes",
				"Wealthy":               "Poor",
			},
			Result: "C",
		},
		{
			Input: map[FieldName]FieldContent{
				"Weather":               "Rainy",
				"Parental Availibility": "No",
				"Wealthy":               "Rich",
			},
			Result: "H",
		},
		{
			Input: map[FieldName]FieldContent{
				"Weather":               "Rainy",
				"Parental Availibility": "Yes",
				"Wealthy":               "Poor",
			},
			Result: "C",
		},
		{
			Input: map[FieldName]FieldContent{
				"Weather":               "Windy",
				"Parental Availibility": "No",
				"Wealthy":               "Poor",
			},
			Result: "C",
		},
		{
			Input: map[FieldName]FieldContent{
				"Weather":               "Windy",
				"Parental Availibility": "No",
				"Wealthy":               "Rich",
			},
			Result: "S",
		},
		{
			Input: map[FieldName]FieldContent{
				"Weather":               "Windy",
				"Parental Availibility": "Yes",
				"Wealthy":               "Rich",
			},
			Result: "C",
		},
		{
			Input: map[FieldName]FieldContent{
				"Weather":               "Sunny",
				"Parental Availibility": "No",
				"Wealthy":               "Rich",
			},
			Result: "T",
		},
	},
}

func TestID3(t *testing.T) {
	p := &ID3Processor{}
	p.Load(dataTest)

	assert.InDelta(t, 1.5709505, ES(p.DB.Rows), 0.0001)
	assert.InDelta(t, 0.6954618, Gain(p.DB.Rows, "Weather"), 0.0001)
	assert.InDelta(t, 0.6099866, Gain(p.DB.Rows, "Parental Availibility"), 0.0001)
	assert.InDelta(t, 0.28129077, Gain(p.DB.Rows, "Wealthy"), 0.0001)
	assert.Equal(t, FieldName("Weather"), NextFieldName(p.DB))
}

func TestRows_Entropy(t *testing.T) {
	r := Rows{
		{Result: "one"},
	}

	assert.Equal(t, 0.0, r.Entropy())
}
