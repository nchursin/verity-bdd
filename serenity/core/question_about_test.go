package core_test

import (
	"context"
	stdtesting "testing"

	"github.com/nchursin/serenity-go/serenity/core"
	serenitytesting "github.com/nchursin/serenity-go/serenity/testing"
)

func TestQuestionAboutCreatesQuestion(t *stdtesting.T) {
	ctx := context.Background()
	test := serenitytesting.NewSerenityTestWithContext(ctx, t)

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
