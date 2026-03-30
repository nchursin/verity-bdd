package verity

import (
	"context"

	internalanswerable "github.com/nchursin/verity-bdd/internal/answerable"
	internaltesting "github.com/nchursin/verity-bdd/internal/testing"
)

type ReporterProvider = internaltesting.ReporterProvider
type DefaultAbilityFactory = internaltesting.DefaultAbilityFactory
type Scene = internaltesting.Scene
type VerityTest = internaltesting.VerityTest
type TestContext = internaltesting.TestContext

var NewVerityTest = internaltesting.NewVerityTest
var NewVerityTestWithContext = internaltesting.NewVerityTestWithContext
var NewVerityTestWithReporter = internaltesting.NewVerityTestWithReporter

func ValueOf[T any](value T) Question[T] {
	return internalanswerable.ValueOf(value)
}

func ResultOf[T any](description string, fn func(Actor, context.Context) (T, error)) Question[T] {
	return internalanswerable.ResultOf(description, fn)
}
