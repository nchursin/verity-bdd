package verity_expectations

import (
	internalexpectations "github.com/nchursin/verity-bdd/internal/expectations"
	"github.com/nchursin/verity-bdd/verity_expectations/ensure"
)

type ContainsExpectation = internalexpectations.ContainsExpectation
type ContainsKeyExpectation = internalexpectations.ContainsKeyExpectation
type IsEmptyExpectation = internalexpectations.IsEmptyExpectation
type ArrayLengthEqualsExpectation = internalexpectations.ArrayLengthEqualsExpectation
type IsGreaterThanExpectation = internalexpectations.IsGreaterThanExpectation
type IsLessThanExpectation = internalexpectations.IsLessThanExpectation

var NewContains = internalexpectations.NewContains
var Contains = internalexpectations.Contains
var NewContainsKey = internalexpectations.NewContainsKey
var ContainsKey = internalexpectations.ContainsKey

var NewIsEmpty = internalexpectations.NewIsEmpty
var IsEmpty = internalexpectations.IsEmpty
var NewArrayLengthEquals = internalexpectations.NewArrayLengthEquals
var ArrayLengthEquals = internalexpectations.ArrayLengthEquals

var NewIsGreaterThan = internalexpectations.NewIsGreaterThan
var IsGreaterThan = internalexpectations.IsGreaterThan
var NewIsLessThan = internalexpectations.NewIsLessThan
var IsLessThan = internalexpectations.IsLessThan

func NewEquals[T any](expected T) ensure.Expectation[T] {
	return internalexpectations.NewEquals(expected)
}

func Equals[T any](expected T) ensure.Expectation[T] {
	return internalexpectations.Equals(expected)
}

func Satisfies[T any](description string, fn func(T) error) ensure.Expectation[T] {
	return internalexpectations.Satisfies(description, fn)
}
