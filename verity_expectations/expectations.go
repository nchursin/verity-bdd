package verity_expectations

import (
	internalexpectations "github.com/nchursin/verity-bdd/internal/expectations"
	"github.com/nchursin/verity-bdd/verity_expectations/ensure"
)

var Contains = internalexpectations.Contains
var ContainsKey = internalexpectations.ContainsKey
var IsEmpty = internalexpectations.IsEmpty
var ArrayLengthEquals = internalexpectations.ArrayLengthEquals
var IsGreaterThan = internalexpectations.IsGreaterThan
var IsLessThan = internalexpectations.IsLessThan

func Equals[T any](expected T) ensure.Expectation[T] {
	return internalexpectations.Equals(expected)
}

func Satisfies[T any](description string, fn func(T) error) ensure.Expectation[T] {
	return internalexpectations.Satisfies(description, fn)
}

func Not[T any](inner ensure.Expectation[T]) ensure.Expectation[T] {
	return internalexpectations.Not(inner)
}
