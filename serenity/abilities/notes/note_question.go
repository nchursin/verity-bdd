package notes

import (
	"context"
	"fmt"

	"github.com/nchursin/serenity-go/serenity/core"
)

type noteQuestion[T any] struct {
	key string
}

// Note returns a typed question for the given note key.
func Note[T any](key string) core.Question[T] {
	return &noteQuestion[T]{key: key}
}

// NoteValue returns an untyped note question.
func NoteValue(key string) core.Question[any] {
	return &noteQuestion[any]{key: key}
}

func (q *noteQuestion[T]) Description() string {
	return fmt.Sprintf("the note %q", q.key)
}

func (q *noteQuestion[T]) AnsweredBy(actor core.Actor, ctx context.Context) (T, error) {
	var zero T

	ability, err := core.AbilityOf[*TakeNotesAbility](actor)
	if err != nil {
		return zero, err
	}

	value, err := ability.Get(q.key)
	if err != nil {
		return zero, err
	}

	typedValue, ok := value.(T)
	if !ok {
		return zero, fmt.Errorf("note %q has type %T, not %T", q.key, value, zero)
	}

	return typedValue, nil
}
