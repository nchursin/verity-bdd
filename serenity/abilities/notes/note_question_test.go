package notes

import (
	"context"
	"testing"
)

func TestNoteQuestionReturnsTypedValue(t *testing.T) {
	ability := TakeNotes()
	ability.(*TakeNotesAbility).Set("name", "serenity")
	actor := newStubActor("reader", context.Background(), ability)

	answer, err := Note[string]("name").AnsweredBy(actor, context.Background())
	if err != nil {
		t.Fatalf("expected note to be answered, got error: %v", err)
	}
	if answer != "serenity" {
		t.Fatalf("expected 'serenity', got %v", answer)
	}
}

func TestNoteQuestionErrorsWhenMissing(t *testing.T) {
	actor := newStubActor("reader", context.Background(), TakeNotes())

	_, err := Note[string]("missing").AnsweredBy(actor, context.Background())
	if err == nil {
		t.Fatalf("expected error when note missing")
	}
}

func TestNoteQuestionErrorsWhenNoAbility(t *testing.T) {
	actor := newStubActor("reader", context.Background())

	_, err := Note[string]("missing").AnsweredBy(actor, context.Background())
	if err == nil {
		t.Fatalf("expected error when actor lacks notes ability")
	}
}

func TestNoteQuestionErrorsOnAbilityTypeMismatch(t *testing.T) {
	actor := newStubActor("reader", context.Background(), &dummyAbility{})

	_, err := Note[string]("missing").AnsweredBy(actor, context.Background())
	if err == nil {
		t.Fatalf("expected error when ability type mismatched")
	}
}

func TestNoteQuestionErrorsOnTypeMismatch(t *testing.T) {
	ability := TakeNotes()
	ability.(*TakeNotesAbility).Set("count", 123)
	actor := newStubActor("reader", context.Background(), ability)

	_, err := Note[string]("count").AnsweredBy(actor, context.Background())
	if err == nil {
		t.Fatalf("expected error on type mismatch")
	}
}

func TestNoteValueReturnsUntyped(t *testing.T) {
	ability := TakeNotes()
	ability.(*TakeNotesAbility).Set("count", 321)
	actor := newStubActor("reader", context.Background(), ability)

	value, err := NoteValue("count").AnsweredBy(actor, context.Background())
	if err != nil {
		t.Fatalf("expected note to be answered, got error: %v", err)
	}
	if value != 321 {
		t.Fatalf("expected 321, got %v", value)
	}
}
