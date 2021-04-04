package handler

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gibmir/gotbot/command"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func OnMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by tFhe bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	if isBotWasMentioned(s, m) {
		log.Println("Bot was mentioned")
		commandIndex := command.GetCommandIndex(&m.Content)
		if commandIndex != -1 {
			commandString := m.Content[commandIndex:len(m.Content)]
			log.Printf("There is a command string [%v] in [%v]",
			 commandString, m.Content)
			com := command.Parse(&commandString)
			s.ChannelMessageSend(m.ChannelID,
				"@"+m.Author.Username+", result is: "+processCommand(&com))
			return
		} else {
			s.ChannelMessageSend(m.ChannelID,
				"@"+m.Author.Username+", please, use !help")
		}
	}
}

func isBotWasMentioned(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	for _, mention := range m.Mentions {
		if mention.ID == s.State.User.ID {
			return true
		}
	}
	return false
}

func processCommand(c *command.Command) string {
	switch commandName := c.Name; commandName {
	case "help":
		return "You can use !rnd and !roll"
	case "rnd":
		r, err := processRnd(c)
		if err != nil {
			return err.Error()
		}
		return strconv.Itoa(r)
	case "roll":
		return "roll is unsupported at this moment"
	default:
		return commandName + " is unsupported. Please, use help! command"
	}
}

func processRnd(c *command.Command) (int, error) {
	argCount := len(c.Arguments)
	if argCount == 1 {
		arg, err := strconv.ParseInt(c.Arguments[0], 0, 32)
		if err != nil {
			return 0, fmt.Errorf("can't parse argument %v, %v", c.Arguments[0], err)
		}
		if arg == 0 {
			return 0, fmt.Errorf("you can't use 0 as argument")
		}
		rand.Seed(time.Now().UnixNano())
		return rand.Intn(int(arg)), nil
	} else if argCount == 2 {
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
		rand.Seed(time.Now().UnixNano())

		max := int(max64)
		min := int(min64)
		n := max - min + 1
		if n < 0 {
			return 0, fmt.Errorf("incorrect borders [%v] and [%v]", min, max)
		}
		return int(rand.Intn(n) + min), nil
	} else {
		return 0, fmt.Errorf("incorrect arguments count [%v]",
			argCount)
	}
}
