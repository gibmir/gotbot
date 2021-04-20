package command

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestDynamicProcessWithUnexpectedArgBug(t *testing.T) {
	registry := EmptyRegistry()
	dynamic := NewDynamicProcessor(&registry)
	unexpectedCommand := Command{"unexpected", make([]string, 0)}
	_, err := dynamic.Process(&unexpectedCommand)
	assert.ErrorContains(t, err, "unsupported")
}

func TestDynamicProcessWithHelp(t *testing.T) {
	registry := EmptyRegistry()
	var help CommandProcessor = &HelpCommandProcessor{&registry}
	registry.Register("help", &help)
	dynamic := NewDynamicProcessor(&registry)
	unexpectedCommand := Command{"help", make([]string, 0)}
	actual, err := dynamic.Process(&unexpectedCommand)
	assert.NilError(t, err)
	assert.Equal(t, HelpDescriprion, actual)
}

func TestGetCommandIndexFromCorrectString(t *testing.T) {
	correctString := "!command"
	cmdIndex := GetCommandIndex(&correctString)
	assert.Check(t, cmdIndex != -1)
}

func TestGetCommandIndexFromIncorrectString(t *testing.T) {
	correctString := "command"
	cmdIndex := GetCommandIndex(&correctString)
	assert.Check(t, cmdIndex == -1)
}

func TestParseCommandWithoutArgs(t *testing.T) {
	cmdString := "!command"
	cmd := Parse(&cmdString)
	assert.Equal(t, cmd.Name, "command")
	assert.Equal(t, len(cmd.Arguments), 0)
}

func TestParseCommandWithOneArg(t *testing.T) {
	cmdString := "!command arg"
	cmd := Parse(&cmdString)
	assert.Equal(t, cmd.Name, "command")
	assert.Equal(t, len(cmd.Arguments), 1)
	assert.Equal(t, cmd.Arguments[0], "arg")
}

func TestParseCommandWithFewArgs(t *testing.T) {
	cmdString := "!command arg1 arg2 arg3"
	cmd := Parse(&cmdString)
	assert.Equal(t, cmd.Name, "command")
	assert.Equal(t, len(cmd.Arguments), 3)
	assert.Equal(t, cmd.Arguments[0], "arg1")
	assert.Equal(t, cmd.Arguments[1], "arg2")
	assert.Equal(t, cmd.Arguments[2], "arg3")
}

func TestParseIncorrectCommand(t *testing.T) {
	incorrectString := "command arg1 arg2 arg3"
	cmd := Parse(&incorrectString)
	assert.Equal(t, cmd.Name, "command")
	assert.Equal(t, len(cmd.Arguments), 3)
	assert.Equal(t, cmd.Arguments[0], "arg1")
	assert.Equal(t, cmd.Arguments[1], "arg2")
	assert.Equal(t, cmd.Arguments[2], "arg3")
}