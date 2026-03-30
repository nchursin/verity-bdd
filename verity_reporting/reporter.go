package verity_reporting

import internalreporting "github.com/nchursin/verity-bdd/internal/reporting"

type Reporter = internalreporting.Reporter
type TestResult = internalreporting.TestResult
type Attachment = internalreporting.Attachment
type Status = internalreporting.Status

const (
	StatusPassed  = internalreporting.StatusPassed
	StatusFailed  = internalreporting.StatusFailed
	StatusSkipped = internalreporting.StatusSkipped
)

type TestRunnerAdapter = internalreporting.TestRunnerAdapter
type ActivityTracker = internalreporting.ActivityTracker

var NewTestRunnerAdapter = internalreporting.NewTestRunnerAdapter
var NewActivityTracker = internalreporting.NewActivityTracker
var NewActivityTrackerWithActor = internalreporting.NewActivityTrackerWithActor
