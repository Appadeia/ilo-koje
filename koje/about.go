package koje

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func contains(s []discordgo.User, e discordgo.User) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func about(s *discordgo.Session, m *discordgo.MessageCreate) {
	guildCount := len(s.State.Guilds)
	var users []discordgo.User
	for _, guild := range s.State.Guilds {
		for _, user := range guild.Members {
			if !contains(users, *user.User) {
				users = append(users, *user.User)
			}
		}
	}
	embed := NewEmbed()
	embed.SetTitle("About Me")
	embed.AddField("Guilds", strconv.Itoa(guildCount))
	embed.AddField("Users", strconv.Itoa(len(users)))
	embed.SetColor(0xfefe62)
	s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
}
