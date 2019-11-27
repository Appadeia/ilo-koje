package koje

import (
	"encoding/json"

	"github.com/Necroforger/dgwidgets"
	"github.com/bwmarrin/discordgo"
)

func browseety(s *discordgo.Session, m *discordgo.MessageCreate) {
	ety := EtyDict()
	var etys []puEty
	err := json.Unmarshal([]byte(ety), &etys)

	if err != nil {
		embed := NewEmbed().
			SetTitle("Error!").
			SetDescription(err.Error()).
			SetColor(0xff0000)
		s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
		return
	}
	page := dgwidgets.NewPaginator(s, m.ChannelID)
	for _, val := range etys {
		embed := NewEmbed()
		embed.SetTitle(val.Word)
		if val.Origin.SourceWord != "" {
			embed.AddField(val.Origin.Language, val.Origin.SourceWord)
		} else {
			embed.SetDescription(val.Origin.Language)
		}
		page.Add(embed.MessageEmbed)
	}
	page.SetPageFooters()
	page.Spawn()
}
