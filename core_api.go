package verity

import (
	"context"

	internalabilities "github.com/nchursin/verity-bdd/internal/abilities"
	internalcore "github.com/nchursin/verity-bdd/internal/core"
)

type Ability = internalabilities.Ability

type Actor = internalcore.Actor
type Activity = internalcore.Activity
type Interaction = internalcore.Interaction
type Task = internalcore.Task

type Question[T any] interface {
	Description() string
	AnsweredBy(actor Actor, ctx context.Context) (T, error)
}

type TestResult = internalcore.TestResult
type Status = internalcore.Status
type FailureMode = internalcore.FailureMode

const (
	StatusPending = internalcore.StatusPending
	StatusRunning = internalcore.StatusRunning
	StatusPassed  = internalcore.StatusPassed
	StatusFailed  = internalcore.StatusFailed
	StatusSkipped = internalcore.StatusSkipped
)

const (
	FailFast         = internalcore.FailFast
	ErrorButContinue = internalcore.ErrorButContinue
	Ignore           = internalcore.Ignore
)

var Do = internalcore.Do
var TaskWhere = internalcore.TaskWhere

var Critical = internalcore.Critical
var NonCritical = internalcore.NonCritical
var Optional = internalcore.Optional

var AbilityName = internalcore.AbilityName

func NewQuestion[T any](description string, ask func(actor Actor, ctx context.Context) (T, error)) Question[T] {
	return internalcore.NewQuestion(description, ask)
}

func QuestionAbout[T any](description string, ask func(actor Actor, ctx context.Context) (T, error)) Question[T] {
	return internalcore.QuestionAbout(description, ask)
}

func AbilityOf[T Ability](actor Actor) (T, error) {
	return internalcore.AbilityOf[T](actor)
}
