package expectations_test

import (
	"testing"

	"github.com/nchursin/verity-bdd/internal/expectations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNot_PassesWhenInnerFails(t *testing.T) {
	err := expectations.Not(expectations.Equals("foo")).Evaluate("bar")
	assert.NoError(t, err)
}

func TestNot_FailsWhenInnerPasses(t *testing.T) {
	err := expectations.Not(expectations.Equals("foo")).Evaluate("foo")
	require.Error(t, err)
	assert.Equal(t, "not equals foo", err.Error())
}

func TestNot_Description(t *testing.T) {
	desc := expectations.Not(expectations.Equals("foo")).Description()
	assert.Equal(t, "not equals foo", desc)
}

func TestNot_IsEmpty_FailsOnEmptyString(t *testing.T) {
	err := expectations.Not(expectations.IsEmpty()).Evaluate("")
	require.Error(t, err)
	assert.Equal(t, "not is empty", err.Error())
}

func TestNot_IsEmpty_PassesOnNonEmptyString(t *testing.T) {
	err := expectations.Not(expectations.IsEmpty()).Evaluate("hello")
	assert.NoError(t, err)
}
