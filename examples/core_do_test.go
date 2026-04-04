package examples

import (
	"context"
	"fmt"
	"testing"

	verity "github.com/nchursin/verity-bdd"
)

// TestCoreDoFunction demonstrates the new verity.Do function for quick activity creation
func TestCoreDoFunction(t *testing.T) {
	ctx := context.Background()
	test := verity.NewVerityTestWithContext(ctx, t)

	actor := test.ActorCalled("TestActor")

	// Test the new verity.Do function with FailFast mode
	actor.AttemptsTo(
		verity.Do("#actor performs a simple action", func(ctx context.Context, actor verity.Actor) error {
			// Simple test action
			t.Logf("Actor %s is performing a custom action", actor.Name())
			return nil
		}),
	)

	// Test verity.Do with access to actor abilities
	actor.AttemptsTo(
		verity.Do("#actor accesses actor information", func(ctx context.Context, actor verity.Actor) error {
			// Verify we can access actor properties
			if actor.Name() != "TestActor" {
				return fmt.Errorf("expected actor name 'TestActor', got '%s'", actor.Name())
			}
			t.Logf("Successfully accessed actor: %s", actor.Name())
			return nil
		}),
	)
}
