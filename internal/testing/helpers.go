// Package utilities for TestContext API implementation.
// These functions provide common functionality for test management and actor creation.
//
// The helper functions are designed to be used internally by the VerityTest
// implementation but can also be used by custom test managers that need
// TestContext integration.
package testing

import (
	"reflect"

	"github.com/nchursin/verity-bdd/internal/abilities"
)

// abilityMatchesType checks if the provided ability is assignable to or implements
// the target ability type (interface or concrete) based on runtime types.
func abilityMatchesType(ability abilities.Ability, abilityType abilities.Ability) bool {
	if ability == nil || abilityType == nil {
		return false
	}

	targetType := reflect.TypeOf(abilityType)
	abType := reflect.TypeOf(ability)

	if abType.AssignableTo(targetType) || (targetType.Kind() == reflect.Interface && abType.Implements(targetType)) {
		return true
	}

	if abType.Kind() == reflect.Pointer {
		abElem := abType.Elem()
		if abElem.AssignableTo(targetType) || (targetType.Kind() == reflect.Interface && abElem.Implements(targetType)) {
			return true
		}
	}

	if targetType.Kind() == reflect.Pointer {
		targetElem := targetType.Elem()
		if abType.AssignableTo(targetElem) || (targetElem.Kind() == reflect.Interface && abType.Implements(targetElem)) {
			return true
		}
		if abType.Kind() == reflect.Pointer {
			abElem := abType.Elem()
			if abElem.AssignableTo(targetElem) || (targetElem.Kind() == reflect.Interface && abElem.Implements(targetElem)) {
				return true
			}
		}
	}

	return abType.String() == targetType.String()
}
