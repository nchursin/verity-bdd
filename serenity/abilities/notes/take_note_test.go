package notes

import (
	"context"
	"fmt"
	"testing"

	"github.com/nchursin/serenity-go/serenity/abilities"
	"github.com/nchursin/serenity-go/serenity/core"
	"github.com/nchursin/serenity-go/serenity/reporting"
	"github.com/nchursin/serenity-go/serenity/reporting/mocks"
	serenitytesting "github.com/nchursin/serenity-go/serenity/testing"
	"go.uber.org/mock/gomock"
)

type stubActor struct {
	name      string
	ctx       context.Context
	abilities []abilities.Ability
}

func newStubActor(name string, ctx context.Context, abilities ...abilities.Ability) *stubActor {
	return &stubActor{name: name, ctx: ctx, abilities: abilities}
}

func (a *stubActor) Name() string { return a.name }

func (a *stubActor) Context() context.Context { return a.ctx }

func (a *stubActor) WhoCan(abilities ...abilities.Ability) core.Actor {
	a.abilities = append(a.abilities, abilities...)
	return a
}

func (a *stubActor) AbilityTo(target abilities.Ability) (abilities.Ability, error) {
	for _, ability := range a.abilities {
		if fmt.Sprintf("%T", ability) == fmt.Sprintf("%T", target) {
			return ability, nil
		}
	}
	return nil, fmt.Errorf("actor '%s' can't %s. Did you give them the ability?", a.name, core.AbilityName(target))
}

func (a *stubActor) AttemptsTo(activities ...core.Activity) {}

func (a *stubActor) AnswersTo(question core.Question[any]) (any, bool) {
	return nil, false
}

func TestTakeNoteStoresValue(t *testing.T) {
	noteBook := NewNoteBook()
	actor := newStubActor("alice", context.Background(), noteBook)

	activity := TakeNoteOf("secret-token").As("auth")

	if activity.Description() != "#actor takes note \"auth\"" {
		t.Fatalf("unexpected description: %s", activity.Description())
	}

	if activity.FailureMode() != core.FailFast {
		t.Fatalf("expected FailFast failure mode")
	}

	err := activity.PerformAs(actor, context.Background())
	if err != nil {
		t.Fatalf("expected no error performing take note, got %v", err)
	}

	value, err := noteBook.Get("auth")
	if err != nil {
		t.Fatalf("expected stored note, got error: %v", err)
	}
	if value != "secret-token" {
		t.Fatalf("expected value to be stored, got %v", value)
	}
}

func TestTakeNoteRequiresAbility(t *testing.T) {
	actor := newStubActor("bob", context.Background())

	activity := TakeNoteOf("value").As("missing")
	err := activity.PerformAs(actor, context.Background())
	if err == nil {
		t.Fatalf("expected error when actor lacks notebook ability")
	}
	if err.Error() == "" {
		t.Fatalf("expected error message, got empty string")
	}
}

func TestTakeNoteReportsStep(t *testing.T) {
	ctrl := gomock.NewController(t)
	reporter := mocks.NewMockReporter(ctrl)

	reporter.EXPECT().OnTestStart(t.Name())
	reporter.EXPECT().OnStepStart("Sam takes note \"remember\"")
	reporter.EXPECT().OnStepFinish(gomock.Any()).Do(func(result reporting.TestResult) {
		if result.Name() != "Sam takes note \"remember\"" {
			t.Fatalf("unexpected step name: %s", result.Name())
		}
		if result.Status() != reporting.StatusPassed {
			t.Fatalf("expected status passed, got %v", result.Status())
		}
	})
	reporter.EXPECT().OnTestFinish(gomock.Any())

	serenityTest := serenitytesting.NewSerenityTestWithReporter(context.Background(), t, reporter)
	actor := serenityTest.ActorCalled("Sam").WhoCan(TakeNotes())

	actor.AttemptsTo(TakeNoteOf("secret").As("remember"))
}
