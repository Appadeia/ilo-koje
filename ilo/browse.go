package ilo

import (
	"encoding/json"

	"github.com/Necroforger/dgwidgets"

	"github.com/bwmarrin/discordgo"
)

func browse(s *discordgo.Session, m *discordgo.MessageCreate) {
	pu := PuDict()
	var words []puWord
	err := json.Unmarshal([]byte(pu), &words)

	if err != nil {
		embed := NewEmbed().
			SetTitle(_t("Error!", "pakala!", m)).
			SetDescription(err.Error()).
			SetColor(0xff0000)
		s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
		return
	}
	page := dgwidgets.NewPaginator(s, m.ChannelID)
	for _, val := range words {
		embed := NewEmbed()
		embed.SetTitle(val.Word)
		for _, meaning := range val.Meanings {
			embed.AddField(meaning[0], meaning[1])
		}
		page.Add(embed.MessageEmbed)
	}
	page.SetPageFooters()
	page.Spawn()
}
