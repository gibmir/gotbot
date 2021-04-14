package command

import (
	"fmt"
	"strings"
)

const (
	commandPrefix = "!"
)

type Command struct {
	Name      string
	Arguments []string
}
type CommandProcessor interface {
	Process(c *Command) (string, error)
	GetDescription() string
}

type ProcessorRegistry struct {
	cache map[string]*CommandProcessor
}

func NewRegistry() ProcessorRegistry {
	registry := ProcessorRegistry{map[string]*CommandProcessor{}}
	help := CommandProcessor(HelpCommandProcessor{&registry})
	rnd := CommandProcessor(RndCommandProcessor{})
	roll := CommandProcessor(RollCommandProcessor{})

	registry.Register("help", &help)
	registry.Register("rnd", &rnd)
	registry.Register("roll", &roll)
	return registry
}

func EmptyRegistry() ProcessorRegistry {
	registry := ProcessorRegistry{map[string]*CommandProcessor{}}
	return registry
}

func (registry ProcessorRegistry) Register(cName string, processor *CommandProcessor) {
	registry.cache[cName] = processor
}

func (registry ProcessorRegistry) Get(name *string) *CommandProcessor {
	return registry.cache[*name]
}

type DynamicProcessor struct {
	registry *ProcessorRegistry
}

func NewDynamicProcessor(r *ProcessorRegistry) DynamicProcessor {
	return DynamicProcessor{registry: r}
}

func (p DynamicProcessor) Process(c *Command) (string, error) {
	var processor *CommandProcessor = p.registry.Get(&c.Name)
	if processor != nil {
		return (*processor).Process(c)
	} else {
		return "", fmt.Errorf("%v is unsupported", c.Name)
	}
}

// Provides command index in string.
// string looks like:
// "@!123123, @!321123 !commandName arg1 arg2"
func GetCommandIndex(message *string) int {
	return strings.LastIndex(*message, commandPrefix)
}

func Parse(message *string) Command {
	commandString := strings.Split(*message, " ")
	return Command{removeCommandPrefix(&commandString[0]), commandString[1:] /*arguments*/}
}

func removeCommandPrefix(message *string) string {
	return strings.TrimPrefix(*message, commandPrefix)
}
