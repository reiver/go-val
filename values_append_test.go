package val_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-val"
)

func TestValues_Append_string(t *testing.T) {
	tests := []struct{
		Values []string
		Expected val.Values[string]
	}{
		{
			Expected: val.EmptyValues[string](),
		},



		{
			Values: []string{},
			Expected: val.EmptyValues[string](),
		},



		{
			Values:                 []string{"apple"},
			Expected: val.SomeValues[string]("apple"),
		},
		{
			Values:                 []string{"apple","Banana"},
			Expected: val.SomeValues[string]("apple","Banana"),
		},
		{
			Values:                 []string{"apple","Banana","CHERRY"},
			Expected: val.SomeValues[string]("apple","Banana","CHERRY"),
		},
	}

	for testNumber, test := range tests {

		var actual val.Values[string]
		for _, value := range test.Values {
			actual.Append(value)
		}

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual .Append()s result is not what was expected", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}
	}
}
