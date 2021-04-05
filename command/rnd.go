package command

import (
	"fmt"
	"math/rand"
	"strconv"
)

const RndDescription = `rnd returns random number
with max value border:
@gotbot !rnd 10
with min and max value borders:
@gotbot !rnd 50 100
 `

var ErrRndZeroArgument = fmt.Errorf("you can't use 0 as argument")

type RndCommandProcessor struct{}

func (processor RndCommandProcessor) Process(c *Command) (string, error) {
	r, err := ProcessRnd(c)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(r), nil
}

func (processor RndCommandProcessor) GetDescription() string {
	return RndDescription
}

func ProcessRnd(c *Command) (int, error) {
	argCount := len(c.Arguments)
	if argCount == 1 {
		return rndMax(c)
	} else if argCount == 2 {
		return rndMinMax(c)
	} else {
		return 0, fmt.Errorf("incorrect arguments count [%v]",
			argCount)
	}
}

func rndMax(c *Command) (int, error) {
	arg, err := strconv.ParseInt(c.Arguments[0], 0, 32)
	if err != nil {
		return 0, fmt.Errorf("can't parse argument [%v], %v",
			c.Arguments[0], err)
	}
	if arg == 0 {
		return 0, ErrRndZeroArgument
	}
	return rand.Intn(int(arg)), nil
}

func rndMinMax(c *Command) (int, error) {
	min64, err := strconv.ParseInt(c.Arguments[0], 0, 32)
	if err != nil {
		return 0, fmt.Errorf("can't parse first argument [%v], %v",
			c.Arguments[0], err)
	}
	max64, err := strconv.ParseInt(c.Arguments[1], 0, 32)
	if err != nil {
		return 0, fmt.Errorf("can't parse second argument [%v], %v",
			c.Arguments[1], err)
	}

	max := int(max64)
	min := int(min64)
	n := max - min + 1
	if n < 0 {
		return 0, fmt.Errorf("incorrect borders [%v] and [%v]", min, max)
	}
	return int(rand.Intn(n) + min), nil
}
