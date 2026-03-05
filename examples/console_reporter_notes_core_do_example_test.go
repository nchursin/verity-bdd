package examples

import (
	"context"

	"github.com/nchursin/serenity-go/serenity/abilities/notes"
	"github.com/nchursin/serenity-go/serenity/core"
	serenity "github.com/nchursin/serenity-go/serenity/testing"
)

type exampleTestContext struct {
	name     string
	failed   bool
	cleanups []func()
}

func (e *exampleTestContext) Name() string { return e.name }
func (e *exampleTestContext) Logf(format string, args ...interface{}) {
	// no-op for example output
}
func (e *exampleTestContext) Errorf(format string, args ...interface{}) { e.failed = true }
func (e *exampleTestContext) FailNow()                                  { e.failed = true }
func (e *exampleTestContext) Failed() bool                              { return e.failed }
func (e *exampleTestContext) Cleanup(fn func()) {
	e.cleanups = append(e.cleanups, fn)
}
func (e *exampleTestContext) Helper() {}

func (e *exampleTestContext) RunCleanups() {
	for i := len(e.cleanups) - 1; i >= 0; i-- {
		e.cleanups[i]()
	}
}

// ExampleConsoleReporter_actorNotesWithCoreDo shows how an actor can record notes via core.Do
// and see them in the console reporter output.
func ExampleConsoleReporter_actorNotesWithCoreDo() {
	testCtx := &exampleTestContext{name: "ExampleNotesCoreDo"}
	test := serenity.NewSerenityTest(testCtx)
	defer testCtx.RunCleanups()

	actor := test.ActorCalled("Sam").WhoCan(notes.TakeNotes())
	actor.AttemptsTo(
		core.Do("#actor writes a secret", func(actor core.Actor, ctx context.Context) error {
			ability, err := actor.AbilityTo(&notes.TakeNotesAbility{})
			if err != nil {
				return err
			}

			notebook := ability.(*notes.TakeNotesAbility)
			notebook.Set("token", "secret")
			return nil
		}),
	)

	// Output:
	// 🚀 Starting: ExampleNotesCoreDo
	//   ✅ Sam writes a secret (0.00s)
	// ✅ ExampleNotesCoreDo: PASSED (0.00s)
	//    Attachments:
	//    - notes (application/json): {"Sam":{"token":"secret"}}
}
