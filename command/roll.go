package command

import (
	"fmt"
	"math/rand"
)

const RollDescription = `roll returns random element from provided list:
@gotbot !roll red green blue 
`
var ErrRollNoArgs = fmt.Errorf("there is no arguments for !roll")

type RollCommandProcessor struct{}

func (processor RollCommandProcessor) Process(c *Command) (string, error) {
	return ProcessRoll(c)
}

func (processor RollCommandProcessor) GetDescription() string {
	return RollDescription
}


func ProcessRoll(c *Command) (string, error) {
	argCount := len(c.Arguments)
	if argCount < 1 {
		return "", ErrRollNoArgs
	}
	return c.Arguments[rand.Intn(argCount)], nil
}