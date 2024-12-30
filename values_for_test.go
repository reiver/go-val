package val_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-val"
)

func TestValues_For_string(t *testing.T) {

	tests := []struct{
		Values val.Values[string]
		Expected []string
	}{
		{
		},


		{
			Values: val.SomeValues("apple"),
			Expected:     []string{"apple"},
		},
		{
			Values: val.SomeValues("apple","Banana"),
			Expected:     []string{"apple","Banana"},
		},
		{
			Values: val.SomeValues("apple","Banana","CHERRY"),
			Expected:     []string{"apple","Banana","CHERRY"},
		},
	}

	for testNumber, test := range tests {

		var actual []string
		test.Values.For(func(value string){
			actual = append(actual, value)
		})

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, is actual result of .For() is not what was expected.", testNumber)
			t.Logf("EXPECTED: (%d) %#v", len(expected), expected)
			t.Logf("ACTUAL:   (%d) %#v", len(actual), actual)
			t.Logf("VALUES: %#v", test.Values)
			continue
		}
	}
}
