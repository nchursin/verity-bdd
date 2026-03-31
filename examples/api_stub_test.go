package examples

import (
	"testing"

	"github.com/nchursin/verity-bdd/internal/testing/testserver"
)

func localJSONPlaceholderURL(t testing.TB) string {
	t.Helper()
	return testserver.StartJSONPlaceholderStub(t)
}
