package core

import (
	"fmt"
	"reflect"

	"github.com/nchursin/serenity-go/serenity/abilities"
)

func abilityTypeName(t reflect.Type) string {
	if t == nil {
		return "<nil>"
	}
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return t.String()
}

func abilityTypeNameFor[T abilities.Ability]() string {
	return abilityTypeName(reflect.TypeOf((*T)(nil)).Elem())
}

// AbilityName returns the readable name of the provided ability instance.
func AbilityName(ability abilities.Ability) string {
	return abilityTypeName(reflect.TypeOf(ability))
}

// AbilityOf returns the first ability of type T held by the actor or an error
// with a human-friendly message when missing.
func AbilityOf[T abilities.Ability](actor Actor) (T, error) {
	var zero T
	abName := abilityTypeNameFor[T]()

	if actor == nil {
		return zero, fmt.Errorf("actor is nil; cannot get %s ability", abName)
	}

	ab, err := actor.AbilityTo(zero)
	if err != nil {
		return zero, fmt.Errorf("actor '%s' can't %s. Did you give them the ability?", actor.Name(), abName)
	}

	typed, ok := ab.(T)
	if !ok {
		return zero, fmt.Errorf("actor '%s' returned ability of wrong type (got %s, want %s)", actor.Name(), AbilityName(ab), abName)
	}

	return typed, nil
}
