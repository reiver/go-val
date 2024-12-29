package val_test

import (
	"testing"

	"github.com/reiver/go-val"
)

func TestValues_LastElse_string(t *testing.T) {
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
			Expected:    "Banana",
		},
		{
			Values: val.SomeValues[string]("apple","Banana"),
			Alternative: "LEFT",
			Expected:    "Banana",
		},
		{
			Values: val.SomeValues[string]("apple","Banana"),
			Alternative: "RIGHT",
			Expected:    "Banana",
		},



		{
			Values: val.SomeValues[string]("apple","Banana","CHERRY"),
			Alternative: "",
			Expected:    "CHERRY",
		},
		{
			Values: val.SomeValues[string]("apple","Banana","CHERRY"),
			Alternative: "LEFT",
			Expected:    "CHERRY",
		},
		{
			Values: val.SomeValues[string]("apple","Banana","CHERRY"),
			Alternative: "RIGHT",
			Expected:    "CHERRY",
		},
	}

	for testNumber, test := range tests {

		actual := test.Values.LastElse(test.Alternative)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual .LastElse() result is not what was expected", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			t.Logf("VALUES: %#v", test.Values)
			continue
		}
	}
}
