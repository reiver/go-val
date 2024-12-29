package val_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-val"
)

func TestValues_All_string(t *testing.T) {
	tests := []struct{
		Values val.Values[string]
		Expected []string
	}{
		{
			Expected: []string{},
		},



		{
			Values: val.EmptyValues[string](),
			Expected: []string{},
		},



		{
			Values: val.SomeValues[string]("apple"),
			Expected:             []string{"apple"},
		},
		{
			Values: val.SomeValues[string]("apple","Banana"),
			Expected:             []string{"apple","Banana"},
		},
		{
			Values: val.SomeValues[string]("apple","Banana","CHERRY"),
			Expected:             []string{"apple","Banana","CHERRY"},
		},
	}

	for testNumber, test := range tests {

		actual := test.Values.All()
		expected := test.Expected

		if !reflect.DeepEqual(expected,actual) {
			t.Errorf("For test #%d, the actual .All() result is not what was expected", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}
	}
}
