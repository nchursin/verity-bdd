package verity_wait

import (
	internalcore "github.com/verity-bdd/verity-bdd/internal/core"
	internalensure "github.com/verity-bdd/verity-bdd/internal/expectations/ensure"
	internalwait "github.com/verity-bdd/verity-bdd/internal/wait"
)

// Until creates a wait condition that polls question until expectation is met.
// Default timeout: 5s. Default polling interval: 500ms.
// Configure with .For() and .CheckingEvery().
//
// Example:
//
//	actor.AttemptsTo(
//	    wait.Until(someQuestion, expectations.Equals("ready")),
//	    wait.Until(someQuestion, expectations.Equals("ready")).For(30*time.Second),
//	    wait.Until(someQuestion, expectations.Equals("ready")).For(30*time.Second).CheckingEvery(1*time.Second),
//	)
func Until[T any](question internalcore.Question[T], expectation internalensure.Expectation[T]) *internalwait.ConditionActivity[T] {
	return internalwait.Until(question, expectation)
}
