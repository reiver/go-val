package val_test

import (
	"testing"

	"github.com/reiver/go-val"
)

func TestValues_First_string(t *testing.T) {
	tests := []struct{
		Values val.Values[string]
		ExpectedValue string
		ExpectedFound bool
	}{
		{
			ExpectedValue: "",
			ExpectedFound: false,
		},



		{
			Values: val.EmptyValues[string](),
			ExpectedValue: "",
			ExpectedFound: false,
		},



		{
			Values: val.SomeValues[string]("apple"),
			ExpectedValue: "apple",
			ExpectedFound: true,
		},
		{
			Values: val.SomeValues[string]("apple","Banana"),
			ExpectedValue: "apple",
			ExpectedFound: true,
		},
		{
			Values: val.SomeValues[string]("apple","Banana","CHERRY"),
			ExpectedValue: "apple",
			ExpectedFound: true,
		},
	}

	for testNumber, test := range tests {

		actualValue, actualFound := test.Values.First()

		{
			actual := actualFound
			expected := test.ExpectedFound

			if expected != actual {
				t.Errorf("For test #%d, the actual 'found' .First() result is not what was expected", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("VALUES: %#v", test.Values)
				continue
			}
		}

		{
			actual := actualValue
			expected := test.ExpectedValue

			if expected != actual {
				t.Errorf("For test #%d, the actual 'value' .First() result is not what was expected", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("VALUES: %#v", test.Values)
				continue
			}
		}
	}
}
