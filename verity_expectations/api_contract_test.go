package verity_expectations_test

import (
	"testing"

	verity "github.com/nchursin/verity-bdd"
	ve "github.com/nchursin/verity-bdd/verity_expectations"
	"github.com/nchursin/verity-bdd/verity_expectations/ensure"
)

func TestExpectationsAPIContractCompiles(t *testing.T) {
	q := verity.ValueOf("hello")
	_ = ensure.That(q, ve.Contains("ell"))
}
