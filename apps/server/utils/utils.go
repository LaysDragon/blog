package utils

import "fmt"

type ErrorWrapped interface{ Unwrap() error }

// helper slice mapping function
func MappingFunc[S any, T any](source []S, mapper func(S) T) []T {
	result := make([]T, 0)
	for _, s := range source {
		result = append(result, mapper(s))
	}
	return result

}

// a helper return any type's zero value include interface type(nil)
func NilVal[T any]() T {
	var t T
	return t
}

// an helper to quick return fmt error msg with extra value
func ErrorWrap[T any](val T, err error) func(string) (T, error) {
	return func(msg string) (T, error) {
		if err != nil {
			return val, fmt.Errorf(msg, err)
		}
		return val, err
	}
}
