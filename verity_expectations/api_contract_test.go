package verity_expectations_test

import (
	"testing"

	answerable "github.com/nchursin/verity-bdd/verity_answerable"
	ve "github.com/nchursin/verity-bdd/verity_expectations"
	"github.com/nchursin/verity-bdd/verity_expectations/ensure"
)

func TestExpectationsAPIContractCompiles(t *testing.T) {
	q := answerable.ValueOf("hello")
	_ = ensure.That(q, ve.Contains("ell"))
}
