package command

import (
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestHelpGetDescription(t *testing.T) {
	registry := EmptyRegistry()
	help     := HelpCommandProcessor{&registry}
	actual := help.GetDescription()
	assert.Equal(t, HelpDescriprion, actual)
}

func TestHelpProcessWithoutArgs(t *testing.T) {
	registry := EmptyRegistry()
	help     := HelpCommandProcessor{&registry}
	args := make([]string, 0)
	cmd := Command{"help", args}
	actual, err := help.Process(&cmd)
	assert.NilError(t, err)
	assert.Equal(t, HelpDescriprion, actual)
}

func TestHelpProcessWithUnexpectedArg(t *testing.T) {
	registry := EmptyRegistry()
	help     := HelpCommandProcessor{&registry}
	args := []string{"unexpected command"}
	cmd := Command{"help", args}
	actual, err := help.Process(&cmd)
	assert.NilError(t, err)
	assert.Assert(t, strings.Contains(actual, "unsupported"))
}

func TestHelpProcessWithCorrectArgs(t *testing.T) {
	registry := NewRegistry()
	help     := HelpCommandProcessor{&registry}
	args := []string{"rnd","roll","help"}
	cmd := Command{"help", args}
	actual, err := help.Process(&cmd)
	assert.NilError(t, err)
	assert.Assert(t, strings.Contains(actual, RndDescription))
	assert.Assert(t, strings.Contains(actual, RollDescription))
	assert.Assert(t, strings.Contains(actual, HelpDescriprion))
}
