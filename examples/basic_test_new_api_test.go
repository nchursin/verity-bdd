package examples

import (
	"context"
	"testing"

	verity "github.com/nchursin/verity-bdd"
	"github.com/nchursin/verity-bdd/verity_abilities/api"
	expectations "github.com/nchursin/verity-bdd/verity_expectations"
	"github.com/nchursin/verity-bdd/verity_expectations/ensure"
)

// TestJSONPlaceholderBasicsNewAPI demonstrates basic API testing with JSONPlaceholder using new TestContext API
func TestJSONPlaceholderBasicsNewAPI(t *testing.T) {
	ctx := context.Background()
	test := verity.NewVerityTestWithContext(ctx, t)

	apiTester := test.ActorCalled("APITester").WhoCan(api.CallAnApiAt("https://jsonplaceholder.typicode.com"))

	// Test GET posts - should return existing posts
	apiTester.AttemptsTo(
		api.SendGetRequest("/posts"),
		ensure.That(api.LastResponseStatus{}, expectations.Equals(200)),
		ensure.That(api.LastResponseBody{}, expectations.Contains("title")),
	)

	// Test GET users - should return existing users
	apiTester.AttemptsTo(
		api.SendGetRequest("/users"),
		ensure.That(api.LastResponseStatus{}, expectations.Equals(200)),
		ensure.That(api.LastResponseBody{}, expectations.Contains("email")),
	)

	// The console output will now show detailed step-by-step execution
	// with emojis, timing, and activity tracking thanks to ConsoleReporter
}
