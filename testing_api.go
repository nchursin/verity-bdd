package verity

import (
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

