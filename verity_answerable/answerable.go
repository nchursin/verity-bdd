package verity_answerable

import (
	"context"

	verity "github.com/nchursin/verity-bdd"
	internalanswerable "github.com/nchursin/verity-bdd/internal/answerable"
)

// ValueOf creates a Question[T] that returns the provided static value
// when answered by any actor.
//
// The value is treated as-is, even if it's an error type. This means that
// error values are passed through as the answer rather than being treated
// as failure conditions.
//
// Parameters:
//   - value: The static value to be wrapped as a Question
//
// Returns:
//   - Question[T]: A question that always returns the provided value
//
// Example:
//
//	q := ValueOf(42)
//	result, err := q.AnsweredBy(actor) // result = 42, err = nil
func ValueOf[T any](value T) verity.Question[T] {
	return internalanswerable.ValueOf(value)
}

// ResultOf creates a Question[T] from a function with a custom description.
//
// The function is executed when the question is answered by an actor.
// This allows for dynamic calculations and complex operations in assertions.
//
// Parameters:
//   - description: Human-readable description for test reports
//   - fn: Function that takes a context and actor, returns (value, error)
//
// Returns:
//   - Question[T]: A question that executes the function when answered
//
// Example:
//
//	ensure.That(
//		ResultOf("user count", func(ctx context.Context, actor verity.Actor) (int, error) {
//			db := actor.AbilityTo(DatabaseAbility{}).(DatabaseAbility)
//			return db.CountUsers(), nil
//		}),
//		expectations.IsGreaterThan(0),
//	)
func ResultOf[T any](description string, fn func(context.Context, verity.Actor) (T, error)) verity.Question[T] {
	return internalanswerable.ResultOf(description, fn)
}
