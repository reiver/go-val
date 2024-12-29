package val_test

import (
	"testing"

	"github.com/reiver/go-val"
)

func TestValues_FirstElse_string(t *testing.T) {
	tests := []struct{
		Values val.Values[string]
		Alternative string
		Expected string
	}{
		{
			Alternative: "",
			Expected:    "",
		},
		{
			Alternative: "LEFT",
			Expected:    "LEFT",
		},
		{
			Alternative: "RIGHT",
			Expected:    "RIGHT",
		},



		{
			Values: val.EmptyValues[string](),
			Alternative: "",
			Expected:    "",
		},
		{
			Values: val.EmptyValues[string](),
			Alternative: "LEFT",
			Expected:    "LEFT",
		},
		{
			Values: val.EmptyValues[string](),
			Alternative: "RIGHT",
			Expected:    "RIGHT",
		},



		{
			Values: val.SomeValues[string]("apple"),
			Alternative: "",
			Expected:    "apple",
		},
		{
			Values: val.SomeValues[string]("apple"),
			Alternative: "LEFT",
			Expected:    "apple",
		},
		{
			Values: val.SomeValues[string]("apple"),
			Alternative: "RIGHT",
			Expected:    "apple",
		},



		{
			Values: val.SomeValues[string]("apple","Banana"),
			Alternative: "",
			Expected:    "apple",
		},
		{
			Values: val.SomeValues[string]("apple","Banana"),
			Alternative: "LEFT",
			Expected:    "apple",
		},
		{
			Values: val.SomeValues[string]("apple","Banana"),
			Alternative: "RIGHT",
			Expected:    "apple",
		},



		{
			Values: val.SomeValues[string]("apple","Banana","CHERRY"),
			Alternative: "",
			Expected:    "apple",
		},
		{
			Values: val.SomeValues[string]("apple","Banana","CHERRY"),
			Alternative: "LEFT",
			Expected:    "apple",
		},
		{
			Values: val.SomeValues[string]("apple","Banana","CHERRY"),
			Alternative: "RIGHT",
			Expected:    "apple",
		},
	}

	for testNumber, test := range tests {

		actual := test.Values.FirstElse(test.Alternative)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual .FirstElse() result is not what was expected", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			t.Logf("VALUES: %#v", test.Values)
			continue
		}
	}
}
