package command

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestRollGetDescription(t *testing.T) {
	roll := RollCommandProcessor{}
	actual := roll.GetDescription()
	assert.Equal(t, RollDescription, actual)
}

func TestRollProcessWithoutArgs(t *testing.T) {
	roll := RollCommandProcessor{}
	args := make([]string, 0)
	cmd := Command{"roll", args}
	_, err := roll.Process(&cmd)
	assert.ErrorType(t, err, ErrRollNoArgs)
}

func TestRollProcessWithOneArgBug(t *testing.T) {
	roll := RollCommandProcessor{}
	expected := "10"
	args := []string{expected}
	cmd := Command{"roll", args}
	actual, err := roll.Process(&cmd)
	assert.NilError(t, err)
	assert.Equal(t, expected, actual)
}

func TestRollProcessWithFewArg(t *testing.T) {
	roll := RollCommandProcessor{}
	args := []string{"10", "20", "30"}
	cmd := Command{"roll", args}
	_, err := roll.Process(&cmd)
	assert.NilError(t, err)
}
