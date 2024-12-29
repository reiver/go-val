package val_test

import (
	"testing"

	"github.com/reiver/go-val"
)

func TestValues_GoString_string(t *testing.T) {
	tests := []struct{
		Values val.Values[string]
		Expected string
	}{
		{
			Expected: `val.EmptyValues[string]()`,
		},



		{
			Values:    val.EmptyValues[string](),
			Expected: `val.EmptyValues[string]()`,
		},



		{
			Values:    val.SomeValues[string]("apple"),
			Expected: `val.SomeValues[string]("apple")`,
		},
		{
			Values:    val.SomeValues[string]("apple","Banana"),
			Expected: `val.SomeValues[string]("apple","Banana")`,
		},
		{
			Values:    val.SomeValues[string]("apple","Banana","CHERRY"),
			Expected: `val.SomeValues[string]("apple","Banana","CHERRY")`,
		},
	}

	for testNumber, test := range tests {

		actual := test.Values.GoString()
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual .GoString() result is not what was expected", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
