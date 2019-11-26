package koje

import (
	"encoding/json"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func ety(s *discordgo.Session, m *discordgo.MessageCreate) {
	ety := EtyDict()
	var etys []puEty
	err := json.Unmarshal([]byte(ety), &etys)

	lex := strings.Split(strings.Split(m.Content, "!")[1], " ")

	if len(lex) < 2 {
		embed := NewEmbed().
			SetTitle("Error!").
			SetDescription("No word given to look up in pu.").
			SetColor(0xff0000)
		s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
		return
	}
	if err != nil {
		embed := NewEmbed().
			SetTitle("Error!").
			SetDescription(err.Error()).
			SetColor(0xff0000)
		s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
		return
	}

	var returnEty puEty
	for _, val := range etys {
		if val.Word == lex[1] {
			returnEty = val
			break
		}
	}

	var blank puEty
	if returnEty == blank {
		embed := NewEmbed().
			SetTitle("Error!").
			SetDescription("Word not found.").
			SetColor(0xff0000)
		s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
		return
	}

	embed := NewEmbed()
	embed.SetTitle(returnEty.Word)
	if returnEty.Origin.SourceWord != "" {
		embed.AddField(returnEty.Origin.Language, returnEty.Origin.SourceWord)
	} else {
		embed.SetDescription(returnEty.Origin.Language)
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
}
