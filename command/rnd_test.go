package command

import (
	"strconv"
	"testing"

	"gotest.tools/v3/assert"
)

func TestRndGetDescription(t *testing.T) {
	rnd := RndCommandProcessor{}
	actual := rnd.GetDescription()
	assert.Equal(t, RndDescription, actual)
}

func TestRndProcessWithoutArgs(t *testing.T) {
	rnd := RndCommandProcessor{}
	args := make([]string, 0)
	cmd := Command{"rnd", args}
	_, err := rnd.Process(&cmd)
	assert.ErrorContains(t, err, "incorrect")
	assert.ErrorContains(t, err, "0")
}

func TestRndProcessWithOneArg(t *testing.T) {
	rnd := RndCommandProcessor{}
	args := []string{"10"}
	cmd := Command{"rnd", args}
	_, err := rnd.Process(&cmd)
	assert.NilError(t, err)
}

func TestRndProcessWithTwoArg(t *testing.T) {
	rnd := RndCommandProcessor{}
	args := []string{"10", "20"}
	cmd := Command{"rnd", args}
	_, err := rnd.Process(&cmd)
	assert.NilError(t, err)
}

func TestRndProcessWithIncorrectArgsCount(t *testing.T) {
	rnd := RndCommandProcessor{}
	args := []string{"10", "20", "30", "40"}
	cmd := Command{"rnd", args}
	_, err := rnd.Process(&cmd)
	assert.ErrorContains(t, err, "incorrect")
	assert.ErrorContains(t, err, strconv.Itoa((len(args))))
}

func TestRndProcessWithOneIncorrectArg(t *testing.T) {
	rnd := RndCommandProcessor{}
	incorrectArg := "incorrectArg"
	args := []string{incorrectArg}
	cmd := Command{"rnd", args}
	_, err := rnd.Process(&cmd)
	assert.ErrorContains(t, err, incorrectArg)
}

func TestRndProcessWithFirstIncorrectArg(t *testing.T) {
	rnd := RndCommandProcessor{}
	incorrectArg := "incorrectArg"
	args := []string{incorrectArg, "10"}
	cmd := Command{"rnd", args}
	_, err := rnd.Process(&cmd)
	assert.ErrorContains(t, err, incorrectArg)
}

func TestRndProcessWithSecondIncorrectArg(t *testing.T) {
	rnd := RndCommandProcessor{}
	incorrectArg := "incorrectArg"
	args := []string{"10", incorrectArg}
	cmd := Command{"rnd", args}
	_, err := rnd.Process(&cmd)
	assert.ErrorContains(t, err, incorrectArg)
}
