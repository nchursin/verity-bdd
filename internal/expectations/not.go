package expectations

import (
	"fmt"

	"github.com/nchursin/verity-bdd/internal/expectations/ensure"
)

// NotExpectation inverts any other expectation
type NotExpectation[T any] struct {
	inner ensure.Expectation[T]
}

// Evaluate returns nil if the inner expectation fails, and an error if it passes
func (n NotExpectation[T]) Evaluate(actual T) error {
	if err := n.inner.Evaluate(actual); err == nil {
		return fmt.Errorf("not %s: got %v", n.inner.Description(), actual)
	}
	return nil
}

// Description returns the expectation description
func (n NotExpectation[T]) Description() string {
	return fmt.Sprintf("not %s", n.inner.Description())
}

// Not wraps an expectation and inverts its result
func Not[T any](inner ensure.Expectation[T]) ensure.Expectation[T] {
	return NotExpectation[T]{inner: inner}
}
