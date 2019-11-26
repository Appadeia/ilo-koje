package koje

import (
	"encoding/json"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/google/go-cmp/cmp"
)

func pu(s *discordgo.Session, m *discordgo.MessageCreate) {
	pu := PuDict()
	var words []puWord
	err := json.Unmarshal([]byte(pu), &words)

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
	var returnWord puWord
	for _, val := range words {
		if val.Word == lex[1] {
			returnWord = val
			break
		}
	}

	blank := puWord{}
	if cmp.Equal(returnWord, blank) {
		embed := NewEmbed().
			SetTitle("Error!").
			SetDescription("Word not found.").
			SetColor(0xff0000)
		s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
		return
	}

	embed := NewEmbed()
	embed.SetTitle(returnWord.Word)
	for _, meaning := range returnWord.Meanings {
		embed.AddField(meaning[0], meaning[1])
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
}
