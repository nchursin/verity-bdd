package verity_reporting

import internalreporting "github.com/nchursin/verity-bdd/internal/reporting"

// Reporter handles test execution reporting.
type Reporter = internalreporting.Reporter

// TestResult represents the result of a test or step execution.
type TestResult = internalreporting.TestResult

// Attachment represents additional data to include in a report entry.
// Content should be a serialized payload (for example, JSON).
type Attachment = internalreporting.Attachment

// Status represents the status of a test or step.
type Status = internalreporting.Status

const (
	// StatusPassed indicates a test or step passed.
	StatusPassed = internalreporting.StatusPassed
	// StatusFailed indicates a test or step failed.
	StatusFailed = internalreporting.StatusFailed
	// StatusSkipped indicates a test or step was skipped.
	StatusSkipped = internalreporting.StatusSkipped
)

// TestRunnerAdapter provides integration with test runners.
type TestRunnerAdapter = internalreporting.TestRunnerAdapter

// ActivityTracker tracks activity execution for reporting.
type ActivityTracker = internalreporting.ActivityTracker

// NewTestRunnerAdapter creates a new test runner adapter.
var NewTestRunnerAdapter = internalreporting.NewTestRunnerAdapter

// NewActivityTracker creates a new activity tracker.
var NewActivityTracker = internalreporting.NewActivityTracker

// NewActivityTrackerWithActor creates a new activity tracker with an actor name.
var NewActivityTrackerWithActor = internalreporting.NewActivityTrackerWithActor
