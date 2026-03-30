package verity_expectations_test

import (
	"context"
	"testing"

	verity "github.com/nchursin/verity-bdd"
	ve "github.com/nchursin/verity-bdd/verity_expectations"
	"github.com/nchursin/verity-bdd/verity_expectations/ensure"
)

func TestExpectationsAPIContractCompiles(t *testing.T) {
	q := verity.ValueOf("hello")
	activity := ensure.That(q, ve.Contains("ell"))

	var _ verity.Activity = activity
	_ = context.Background()
}
