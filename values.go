package val

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
