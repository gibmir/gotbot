package handler

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gibmir/gotbot/command"
)

type DiscordHandler struct {
	processor *command.DynamicProcessor
}

func NewDiscordHandler(p *command.DynamicProcessor) DiscordHandler {
	return DiscordHandler{p}
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func (handler *DiscordHandler) OnMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by tFhe bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	if isBotWasMentioned(s, m) {
		log.Println("Bot was mentioned")
		contentLink := removeMentions(m)
		cmdIndex := command.GetCommandIndex(contentLink)
		if cmdIndex != -1 {
			content := *contentLink
			content = content[cmdIndex:]
			log.Printf("There is a command string [%v] in [%v]",
				content, m.Content)
			com := command.Parse(&content)
			var p = *handler.processor
			r, err := p.Process(&com)
			if err != nil {
				r = err.Error()
			}
			s.ChannelMessageSend(m.ChannelID,
				"@"+m.Author.Username+", result is: "+r)
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

func removeMentions(m *discordgo.MessageCreate) *string {
	result := m.Content
	for _, mention := range m.Mentions {
		result = strings.ReplaceAll(result, "<@!"+mention.ID+">", "")
	}
	r := strings.TrimSpace(result)
	return &r
}
