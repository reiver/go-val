package val_test

import (
	"testing"

	"github.com/reiver/go-val"
)

func TestValues_IsEmpty_string(t *testing.T) {
	tests := []struct{
		Values val.Values[string]
		Expected bool
	}{
		{
			Expected: true,
		},



		{
			Values: val.EmptyValues[string](),
			Expected: true,
		},



		{
			Values: val.SomeValues[string]("apple"),
			Expected: false,
		},
		{
			Values: val.SomeValues[string]("apple","Banana"),
			Expected: false,
		},
		{
			Values: val.SomeValues[string]("apple","Banana","CHERRY"),
			Expected: false,
		},
	}

	for testNumber, test := range tests {

		actual := test.Values.IsEmpty()
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual .IsEmpty() result is not what was expected", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			continue
		}
	}
}
