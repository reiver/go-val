package val

import (
	"fmt"
)

// Values represents multiple values.
//
// Example usage:
//
//	var values val.Values[string]
//	
//	// ...
//	
//	values.Append("apple")
//	values.Append("banana")
//	values.Append("cherry")
//	
//	// ..
//	
//	last := values.Last().
type Values[T any] struct {
	values []T
}

func EmptyValues[T any]() Values[T] {
	return Values[T]{}
}


func NewValues[T any](values ...T) *Values[T] {
	var result Values[T] = SomeValues(values...)
	return &result
}

func SomeValues[T any](values ...T) Values[T] {
	return Values[T]{
		values:values,
	}
}

// All returns all the values as a slice.
func (receiver Values[T]) All() []T {
	return append([]T(nil), receiver.values...)
}

// Append appends another value to the list of values.
func (receiver *Values[T]) Append(value T) {
	if nil == receiver {
		return
	}

	receiver.values = append(receiver.values, value)
}

// First returns the first value in the list of values if it exists.
func (receiver Values[T]) First() (T, bool) {
	var values []T = receiver.values

	if len(values) <= 0 {
		var nada T
		return nada, false
	}

	return values[0], true
}

// FirstElse returns the first value in the list of values if it exists, else returns an alternative value.
func (receiver Values[T]) FirstElse(alternative T) T {
	value, something := receiver.First()
	if !something {
		return alternative
	}

	return value
}

func (receiver Values[T]) GoString() string {
	var dummy T
	var typeName string = fmt.Sprintf("%T", dummy)

	if receiver.IsEmpty() {
		return fmt.Sprintf("val.EmptyValues[%s]()", typeName)
	}

	{
		var buffer [256]byte
		var p []byte = buffer[0:0]

		p = append(p, "val.SomeValues["...)
		p = append(p, "]("...)
		for index, value := range receiver.values {
			if 0 < index {
				p = append(p, ',')
			}
			p = append(p, fmt.Sprintf("%#v", value)...)
		}
		p = append(p, typeName...)
		p = append(p, ')')

		return string(p)
	}

}

func (receiver Values[T]) IsEmpty() bool {
	return len(receiver.values) <= 0
}

// Last returns the last value in the list of values if it exists.
func (receiver Values[T]) Last() (T, bool) {
	var values []T = receiver.values

	var length int = len(values)

	if length <= 0 {
		var nada T
		return nada, false
	}

	return values[length-1], true
}

// LastElse returns the last value in the list of values if it exists, else returns an alternative value.
func (receiver Values[T]) LastElse(alternative T) T {
	value, something := receiver.Last()
	if !something {
		return alternative
	}

	return value
}

// Len returns the number of values in the list of values.
func (receiver Values[T]) Len() int {
	return len(receiver.values)
}
