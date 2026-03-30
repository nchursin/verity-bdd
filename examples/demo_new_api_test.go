package examples

import (
	"context"
	"testing"

	"github.com/nchursin/verity-bdd/verity/abilities/api"
	"github.com/nchursin/verity-bdd/verity/expectations"
	"github.com/nchursin/verity-bdd/verity/expectations/ensure"
	verity "github.com/nchursin/verity-bdd/verity/testing"
)

// TestNewAPIDemonstration demonstrates the new TestContext API without require.NoError
func TestNewAPIDemonstration(t *testing.T) {
	// Create VerityTest context - no more manual error handling!
	ctx := context.Background()
	test := verity.NewVerityTestWithContext(ctx, t)

	// Create actor through test context
	apiTester := test.ActorCalled("APITester").WhoCan(api.CallAnApiAt("https://jsonplaceholder.typicode.com"))

	// Chain activities without require.NoError - errors are handled automatically!
	apiTester.AttemptsTo(
		api.SendGetRequest("/posts"),
		ensure.That(api.LastResponseStatus{}, expectations.Equals(200)),
		ensure.That(api.LastResponseBody{}, expectations.Contains("title")),
	)

	// Multiple actors are supported
	user := test.ActorCalled("RegularUser").WhoCan(api.CallAnApiAt("https://jsonplaceholder.typicode.com"))

	user.AttemptsTo(
		api.SendGetRequest("/users"),
		ensure.That(api.LastResponseStatus{}, expectations.Equals(200)),
		ensure.That(api.LastResponseBody{}, expectations.Contains("email")),
	)
}
