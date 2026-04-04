package verity_answerable

import (
	"context"

	verity "github.com/nchursin/verity-bdd"
	internalanswerable "github.com/nchursin/verity-bdd/internal/answerable"
)

func ValueOf[T any](value T) verity.Question[T] {
	return internalanswerable.ValueOf(value)
}

func ResultOf[T any](description string, fn func(context.Context, verity.Actor) (T, error)) verity.Question[T] {
	return internalanswerable.ResultOf(description, fn)
}
