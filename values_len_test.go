package val_test

import (
	"testing"

	"github.com/reiver/go-val"
)

func TestValues_Len_string(t *testing.T) {
	tests := []struct{
		Values val.Values[string]
		Expected int
	}{
		{
			Expected: 0,
		},



		{
			Values: val.EmptyValues[string](),
			Expected: 0,
		},



		{
			Values: val.SomeValues[string]("apple"),
			Expected: 1,
		},
		{
			Values: val.SomeValues[string]("apple","Banana"),
			Expected: 2,
		},
		{
			Values: val.SomeValues[string]("apple","Banana","CHERRY"),
			Expected: 3,
		},
	}

	for testNumber, test := range tests {

		actual := test.Values.Len()
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual .Len() result is not what was expected", testNumber)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			continue
		}
	}
}
