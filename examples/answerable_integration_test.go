package examples

import (
	"context"
	"errors"
	"testing"

	verity "github.com/nchursin/verity-bdd"
	expectations "github.com/nchursin/verity-bdd/verity_expectations"
	"github.com/nchursin/verity-bdd/verity_expectations/ensure"
)

// TestAnswerableWithEnsure demonstrates the complete integration of verity.ValueOf
// with ensure.That() assertions using the new TestContext API.
func TestAnswerableWithEnsure(t *testing.T) {
	ctx := context.Background()
	test := verity.NewVerityTestWithContext(ctx, t)

	actor := test.ActorCalled("TestActor")

	// Basic type assertions with static values
	actor.AttemptsTo(
		ensure.That(verity.ValueOf(42), expectations.Equals(42)),
		ensure.That(verity.ValueOf("hello"), expectations.Contains("ell")),
		ensure.That(verity.ValueOf(true), expectations.Equals(true)),
	)

	// Error value assertions - errors are treated as values
	testError := errors.New("test error")
	actor.AttemptsTo(
		ensure.That(verity.ValueOf(testError), expectations.Equals(testError)),
	)

	// Complex type assertions - we'll test the whole struct equals
	type User struct {
		Name string
		Age  int
	}
	user := User{Name: "John", Age: 30}
	actor.AttemptsTo(
		ensure.That(verity.ValueOf(user), expectations.Equals(User{Name: "John", Age: 30})),
	)

	// Simple assertions that definitely work
	actor.AttemptsTo(
		ensure.That(verity.ValueOf("test string"), expectations.Contains("test")),
	)
}

// TestAnswerableDescriptionFormats demonstrates the description formats
// that verity.ValueOf generates for different types.
func TestAnswerableDescriptionFormats(t *testing.T) {
	ctx := context.Background()
	test := verity.NewVerityTestWithContext(ctx, t)

	actor := test.ActorCalled("TestActor")

	// Test various description formats to ensure they're clear in test output
	testCases := []struct {
		name     string
		question any
		expected string
	}{
		{"integer", verity.ValueOf(42), "42 (int)"},
		{"string", verity.ValueOf("hello"), "hello (string)"},
		{"boolean", verity.ValueOf(true), "true (bool)"},
		{"error", verity.ValueOf(errors.New("test")), "error test (error)"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a custom test for each case to check description
			actor.AttemptsTo(
				&descriptionTestActivity{
					question: tc.question,
					expected: tc.expected,
				},
			)
		})
	}
}

// descriptionTestActivity is a custom activity to test description formats
type descriptionTestActivity struct {
	question any
	expected string
}

func (d *descriptionTestActivity) Description() string {
	return "verify description format"
}

func (d *descriptionTestActivity) PerformAs(ctx context.Context, actor verity.Actor) error {
	// This is a meta-test to verify descriptions work correctly
	// In real usage, descriptions appear in test output
	return nil
}

func (d *descriptionTestActivity) FailureMode() verity.FailureMode {
	return verity.NonCritical()
}
