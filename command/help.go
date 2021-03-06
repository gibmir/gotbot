package command

import "fmt"

const HelpDescriprion = "you can use !rnd !roll"

type HelpCommandProcessor struct {
	registry *ProcessorRegistry
}

func (processor *HelpCommandProcessor) Process(c *Command) (string, error) {
	return processor.processHelp(c), nil
}

func (processor *HelpCommandProcessor) GetDescription() string {
	return HelpDescriprion
}

func (processor *HelpCommandProcessor) processHelp(c *Command) string {
	argCount := len(c.Arguments)
	if argCount == 0 {
		return HelpDescriprion
	} else {
		result := ""
		for _, arg := range c.Arguments {
			result += processor.getDescription(&arg) + "\n"
		}
		return result
	}
}

func (processor *HelpCommandProcessor) getDescription(arg *string) string {
	p := processor.registry.Get(arg)
	if p != nil {
		return (*p).GetDescription()
	} else {
		return fmt.Sprintf("[%v] is unsupported. Please, use !help", *arg)
	}
}
