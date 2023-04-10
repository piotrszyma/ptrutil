package ptrutil

func Deref[T any](val *T) T {
	if val != nil {
		return *val
	}

	var zero T
	return zero
}

func Ref[T any](val T) *T {
	// Returns ref to val
	return &val
}

func RefOfCopy[T any](val T) *T {
	// Copies val and returns ref
	copiedValue := val
	return &copiedValue
}
