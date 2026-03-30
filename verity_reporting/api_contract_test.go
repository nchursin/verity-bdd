package verity_reporting_test

import (
	"testing"

	vr "github.com/nchursin/verity-bdd/verity_reporting"
	"github.com/nchursin/verity-bdd/verity_reporting/console_reporter"
)

func TestReportingAPIContractCompiles(t *testing.T) {
	var reporter vr.Reporter = console_reporter.NewConsoleReporter()
	_ = reporter
}
