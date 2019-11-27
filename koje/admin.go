package koje

import (
	"errors"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func admin(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == cfg.Section("Bot").Key("operator").String() {
		lex := strings.Split(strings.Split(m.Content, "!")[1], " ")
		switch lex[1] {
		case "panic":
			embed := NewEmbed().
				SetTitle("Inducing panic!").
				SetDescription("Goodbye, world!").
				SetColor(0xff0000)
			s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
			panic(errors.New("induced panic"))
		}
	} else {
		embed := NewEmbed().
			SetTitle("Permission denied.").
			SetDescription("You are not the operator of ilo Koje.").
			SetColor(0xff0000)
		s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
		return
	}
}
