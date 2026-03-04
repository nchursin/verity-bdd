package notes

import (
	"context"
	"testing"
)

func TestNoteQuestionReturnsTypedValue(t *testing.T) {
	noteBook := NewNoteBook()
	noteBook.Set("name", "serenity")
	actor := newStubActor("reader", context.Background(), noteBook)

	answer, err := Note[string]("name").AnsweredBy(actor, context.Background())
	if err != nil {
		t.Fatalf("expected note to be answered, got error: %v", err)
	}
	if answer != "serenity" {
		t.Fatalf("expected 'serenity', got %v", answer)
	}
}

func TestNoteQuestionErrorsWhenMissing(t *testing.T) {
	actor := newStubActor("reader", context.Background(), NewNoteBook())

	_, err := Note[string]("missing").AnsweredBy(actor, context.Background())
	if err == nil {
		t.Fatalf("expected error when note missing")
	}
}

func TestNoteQuestionErrorsOnTypeMismatch(t *testing.T) {
	noteBook := NewNoteBook()
	noteBook.Set("count", 123)
	actor := newStubActor("reader", context.Background(), noteBook)

	_, err := Note[string]("count").AnsweredBy(actor, context.Background())
	if err == nil {
		t.Fatalf("expected error on type mismatch")
	}
}

func TestNoteValueReturnsUntyped(t *testing.T) {
	noteBook := NewNoteBook()
	noteBook.Set("count", 321)
	actor := newStubActor("reader", context.Background(), noteBook)

	value, err := NoteValue("count").AnsweredBy(actor, context.Background())
	if err != nil {
		t.Fatalf("expected note to be answered, got error: %v", err)
	}
	if value != 321 {
		t.Fatalf("expected 321, got %v", value)
	}
}
