package command

import (
	"fmt"
	"strings"
)

const (
	commandPrefix      = "!"
)

var (
	HelpProcessor HelpCommandProcessor = HelpCommandProcessor{}
	RndProcessor  RndCommandProcessor  = RndCommandProcessor{}
	RollProcessor RollCommandProcessor = RollCommandProcessor{}
)

type CommandProcessor interface {
	Process(c *Command) (string, error)
	GetDescription() string
}

type Command struct {
	Name      string
	Arguments []string
}

type DefaultCommandProcessor struct {
	commandName string
}

func (processor DefaultCommandProcessor) Process(c *Command) (string, error) {
	return "", fmt.Errorf("%v is unsupported", c.Name)
}

func (processor DefaultCommandProcessor) GetDescription() string {
	return processor.commandName + " is unsupported. Please, use !help"
}

// Provides command index in string.
// string looks like:
// "@!123123, @!321123 !commandName arg1 arg2"
func GetCommandIndex(message *string) int {
	return strings.LastIndex(*message, commandPrefix)
}

func ProcessCommand(c *Command) string {
	processor := ResolveProcessor(&c.Name)
	r, err := processor.Process(c)
	if err != nil {
		return err.Error()
	}
	return r
}

func Parse(message *string) Command {
	commandString := strings.Split(*message, " ")
	return Command{removeCommandPrefix(&commandString[0]), commandString[1:] /*arguments*/}
}

func removeCommandPrefix(message *string) string {
	return strings.TrimPrefix(*message, commandPrefix)
}

func ResolveProcessor(cNameLink *string) CommandProcessor {
	cName := *cNameLink
	switch cName {
	case "help":
		return HelpProcessor
	case "rnd":
		return RndProcessor
	case "roll":
		return RollProcessor
	default:
		return DefaultCommandProcessor{cName}
	}
}
