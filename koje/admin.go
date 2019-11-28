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
				SetTitle(_t("Inducing panic!", "mi moli!", m)).
				SetDescription(_t("Goodbye, world!", "mi tawa!", m)).
				SetColor(0xff0000)
			s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
			panic(errors.New("induced panic"))
		case "blacklist":
			if chanBlacklisted(m.ChannelID) {
				embed := NewEmbed().
					SetTitle(_t("Channel no longer blacklisted for running commands", "tenpo ni la poki toki li jo e ken pi toki lawa", m)).
					SetColor(0xff0000)
				s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
				setChanBlacklisted(m.ChannelID, false)
			} else {
				embed := NewEmbed().
					SetTitle(_t("Channel blacklisted for running commands", "tenpo ni la poki toki li jo e ken pi toki lawa ala", m)).
					SetColor(0xff0000)
				s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
				setChanBlacklisted(m.ChannelID, true)
			}
		case "tokiponataso":
			if chanTped(m.ChannelID) {
				embed := NewEmbed().
					SetTitle("Channel is now no longer Toki Pona only.").
					SetColor(0xff0000)
				s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
				setChanTped(m.ChannelID, false)
			} else {
				embed := NewEmbed().
					SetTitle("tenpo ni la mi toki kepeken toki pona taso lon ni.").
					SetColor(0xff0000)
				s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
				setChanTped(m.ChannelID, true)
			}
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
