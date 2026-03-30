package core_test

import (
	"context"
	stdtesting "testing"

	"github.com/nchursin/verity-bdd/internal/core"
	veritytesting "github.com/nchursin/verity-bdd/internal/testing"
)

func TestQuestionAboutCreatesQuestion(t *stdtesting.T) {
	ctx := context.Background()
	test := veritytesting.NewVerityTestWithContext(ctx, t)

	question := core.QuestionAbout[int]("number", func(actor core.Actor, ctx context.Context) (int, error) {
		return 7, nil
	})

	actor := test.ActorCalled("Questioner")
	value, err := question.AnsweredBy(actor, ctx)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if value != 7 {
		t.Fatalf("expected value 7, got %d", value)
	}

	if description := question.Description(); description != "asks number" {
		t.Fatalf("expected description 'asks number', got %q", description)
	}
}
