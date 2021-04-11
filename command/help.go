package command

import "fmt"

const helpDescriprion = "you can use !rnd !roll"

type HelpCommandProcessor struct {
	registry *ProcessorRegistry
}

func (processor HelpCommandProcessor) Process(c *Command) (string, error) {
	return processor.processHelp(c), nil
}

func (processor HelpCommandProcessor) GetDescription() string {
	return helpDescriprion
}

func (processor HelpCommandProcessor) processHelp(c *Command) string {
	argCount := len(c.Arguments)
	if argCount == 0 {
		return helpDescriprion
	} else {
		result := ""
		for _, arg := range c.Arguments {
			result += processor.getDescription(&arg) + "\n"
		}
		return result
	}
}

func (processor HelpCommandProcessor) getDescription(arg *string) string {
	pLink := processor.registry.Get(arg)
	if pLink != nil {
		p := *pLink
		return p.GetDescription()
	} else {
		return fmt.Sprintf("[%v] is unsupported. Please, use !help", *arg)
	}
}
