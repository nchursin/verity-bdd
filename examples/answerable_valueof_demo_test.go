package examples

import (
	"errors"
	"testing"

	verity "github.com/nchursin/verity-bdd"
	"github.com/nchursin/verity-bdd/verity_abilities/api"
	expectations "github.com/nchursin/verity-bdd/verity_expectations"
	"github.com/nchursin/verity-bdd/verity_expectations/ensure"
)

// TestAnswerableValueOf demonstrates the usage of verity.ValueOf() API
// for creating Question[T] objects from static values.
//
// This enables the use of static values in ensure.That() assertions:
//
//	ensure.That(verity.ValueOf(4), expectations.Equals(4))
//
// Previously, only core.Question[T] objects could be used in ensure.That(),
// but verity.ValueOf() allows any static value to be wrapped as a Question.
func TestAnswerableValueOf(t *testing.T) {
	test := verity.NewVerityTest(t, verity.Scene{})

	actor := test.ActorCalled("ValueTester")

	// Basic scalar values
	actor.AttemptsTo(
		ensure.That(verity.ValueOf(42), expectations.Equals(42)),
		ensure.That(verity.ValueOf(3.14), expectations.Equals(3.14)),
		ensure.That(verity.ValueOf("hello"), expectations.Contains("hell")),
		ensure.That(verity.ValueOf(true), expectations.Equals(true)),
	)

	// Complex types
	type Person struct {
		Name string
		Age  int
	}

	person := Person{Name: "Alice", Age: 25}
	actor.AttemptsTo(
		ensure.That(verity.ValueOf(person), expectations.Equals(Person{Name: "Alice", Age: 25})),
	)

	// Error values - check error message as string
	err := errors.New("connection failed")
	actor.AttemptsTo(
		ensure.That(verity.ValueOf(err.Error()), expectations.Contains("connection")),
	)
}

// TestAnswerableWithMixedQuestions demonstrates mixing static value questions
// with traditional dynamic questions from API interactions.
func TestAnswerableWithMixedQuestions(t *testing.T) {
	test := verity.NewVerityTest(t, verity.Scene{})
	apiBaseURL := localJSONPlaceholderURL(t)

	apiTester := test.ActorCalled("APITester").WhoCan(api.CallAnApiAt(apiBaseURL))

	// Mix of static value questions (using verity.ValueOf)
	// and dynamic questions (from API interactions)
	apiTester.AttemptsTo(
		// Dynamic: Get actual status from API
		api.SendGetRequest("/posts/1"),

		// Static: Compare against expected status code
		ensure.That(verity.ValueOf(200), expectations.Equals(200)),

		// Dynamic: Get actual response body
		ensure.That(api.LastResponseStatus{}, expectations.Equals(200)),

		// Static: Check response contains expected content
		ensure.That(api.LastResponseBody{}, expectations.Contains("title")),

		// Static: Verify expected response structure
		ensure.That(verity.ValueOf("userId"), expectations.Contains("user")),
	)
}

// TestAnswerableDescriptions shows how verity.ValueOf() generates
// clear descriptions that appear in test failure messages.
func TestAnswerableDescriptions(t *testing.T) {
	test := verity.NewVerityTest(t, verity.Scene{})

	actor := test.ActorCalled("DescriptionTester")

	// These descriptions will appear in test output:
	// "42 (int)"
	// "hello (string)"
	// "error something went wrong (error)"
	actor.AttemptsTo(
		ensure.That(verity.ValueOf(42), expectations.Equals(42)),
		ensure.That(verity.ValueOf("hello"), expectations.Contains("hell")),
		ensure.That(verity.ValueOf(errors.New("something went wrong").Error()), expectations.Contains("something")),
	)
}

// TestAnswerableEdgeCases demonstrates handling of edge cases
func TestAnswerableEdgeCases(t *testing.T) {
	test := verity.NewVerityTest(t, verity.Scene{})

	actor := test.ActorCalled("EdgeCaseTester")

	// Zero values
	actor.AttemptsTo(
		ensure.That(verity.ValueOf(0), expectations.Equals(0)),
		ensure.That(verity.ValueOf(false), expectations.Equals(false)),
		ensure.That(verity.ValueOf(""), expectations.Equals("")),
	)
}
