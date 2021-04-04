package command

import "strings"

const commandPrefix = "!"

type Command struct {
	Name      string
	Arguments []string
}

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
