package command

const helpDescriprion = "you can use !rnd !roll"

type HelpCommandProcessor struct{}

func (processor HelpCommandProcessor) Process(c *Command) (string, error) {
	return ProcessHelp(c), nil
}

func (processor HelpCommandProcessor) GetDescription() string {
	return helpDescriprion
}

func ProcessHelp(c *Command) string {
	argCount := len(c.Arguments)
	if argCount == 0 {
		return helpDescriprion
	} else {
		result := ""
		for _, arg := range c.Arguments {
			result += getDescription(&arg) + "\n"
		}
		return result
	}
}

func getDescription(arg *string) string {
	processor := ResolveProcessor(arg)
	return processor.GetDescription()
}
