package ensure

import (
	verity "github.com/nchursin/verity-bdd"
	internalensure "github.com/nchursin/verity-bdd/internal/expectations/ensure"
)

type Expectation[T any] interface {
	Evaluate(actual T) error
	Description() string
}

func That[T any](question verity.Question[T], expectation Expectation[T]) verity.Activity {
	return internalensure.That(question, expectation)
}
