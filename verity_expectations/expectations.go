package verity_expectations

import (
	internalexpectations "github.com/nchursin/verity-bdd/internal/expectations"
	"github.com/nchursin/verity-bdd/verity_expectations/ensure"
)

// Contains checks if a string contains the expected substring.
var Contains = internalexpectations.Contains

// ContainsKey checks if a map contains the expected key.
var ContainsKey = internalexpectations.ContainsKey

// IsGreaterThan checks if a numeric value is greater than the expected value.
var IsGreaterThan = internalexpectations.IsGreaterThan

// IsLessThan checks if a numeric value is less than the expected value.
var IsLessThan = internalexpectations.IsLessThan

// IsEmpty checks if a value is empty (string, slice, array, or map).
func IsEmpty[T any]() ensure.Expectation[T] {
	return internalexpectations.IsEmpty[T]()
}

// ArrayLengthEquals checks if an array, slice, or string has the expected length.
func ArrayLengthEquals[T any](expectedLength int) ensure.Expectation[T] {
	return internalexpectations.ArrayLengthEquals[T](expectedLength)
}

// Equals checks if the actual value equals the expected value using deep equality.
func Equals[T any](expected T) ensure.Expectation[T] {
	return internalexpectations.Equals(expected)
}

// Satisfies creates a custom expectation with a description and validation function.
//
// The function should return nil if the expectation is met, or an error describing the failure.
//
// Example:
//
//	actor.AttemptsTo(
//		ensure.That(ValueOf(value), Satisfies("is positive number", func(actual int) error {
//			if actual <= 0 {
//				return fmt.Errorf("expected positive value, but got %v", actual)
//			}
//			return nil
//		})),
//	)
func Satisfies[T any](description string, fn func(T) error) ensure.Expectation[T] {
	return internalexpectations.Satisfies(description, fn)
}

// Not wraps an expectation and inverts its result.
func Not[T any](inner ensure.Expectation[T]) ensure.Expectation[T] {
	return internalexpectations.Not(inner)
}
