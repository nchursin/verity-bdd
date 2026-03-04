package notes

import (
	"context"
	"fmt"

	"github.com/nchursin/serenity-go/serenity/core"
)

// takeNoteBuilder constructs the TakeNote activity.
type takeNoteBuilder struct {
	value any
}

// TakeNoteOf starts a TakeNote activity definition.
func TakeNoteOf(value any) *takeNoteBuilder {
	return &takeNoteBuilder{value: value}
}

// As finalizes the TakeNote activity with the provided key.
func (b *takeNoteBuilder) As(key string) core.Activity {
	return &takeNoteActivity{key: key, value: b.value}
}

type takeNoteActivity struct {
	key   string
	value any
}

func (t *takeNoteActivity) Description() string {
	return fmt.Sprintf("#actor takes note %q", t.key)
}

func (t *takeNoteActivity) PerformAs(actor core.Actor, ctx context.Context) error {
	ability, err := actor.AbilityTo(&TakeNotesAbility{})
	if err != nil {
		return err
	}

	noteBookAbility, ok := ability.(*TakeNotesAbility)
	if !ok {
		return fmt.Errorf("notes ability must be *TakeNotesAbility, got %T", ability)
	}

	noteBookAbility.Set(t.key, t.value)
	return nil
}

func (t *takeNoteActivity) FailureMode() core.FailureMode {
	return core.FailFast
}
