package command

import (
	"gotest.tools/v3/assert"
	"testing"
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
