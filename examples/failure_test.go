package examples

import (
	"testing"

	verity "github.com/nchursin/verity-bdd"
	"github.com/nchursin/verity-bdd/verity_abilities/api"
	expectations "github.com/nchursin/verity-bdd/verity_expectations"
	"github.com/nchursin/verity-bdd/verity_expectations/ensure"
)

// TestFailureHandling demonstrates how different failure modes work
func TestFailureHandling(t *testing.T) {
	test := verity.NewVerityTest(t, verity.Scene{})

	apiTester := test.ActorCalled("APITester").WhoCan(api.CallAnApiAt("https://jsonplaceholder.typicode.com"))

	// This should pass - all assertions are correct
	apiTester.AttemptsTo(
		api.SendGetRequest("/posts"),
		ensure.That(api.LastResponseStatus{}, expectations.Equals(200)),
	)

	// This should fail and stop the test - wrong status code
	// apiTester.AttemptsTo(
	// 	api.SendGetRequest("/posts"),
	// 	ensure.That(api.LastResponseStatus{}, expectations.Equals(202)),
	// 	ensure.That(api.LastResponseStatus{}, expectations.Equals(404)), // This will fail
	// )
}
