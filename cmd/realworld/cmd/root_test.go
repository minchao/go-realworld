package cmd

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestExecute(t *testing.T) {
	c := qt.New(t)
	got := Execute
	c.Assert(got, qt.Not(qt.IsNil))
}
