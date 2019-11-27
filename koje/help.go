package koje

import (
	"github.com/Necroforger/dgwidgets"
	"github.com/bwmarrin/discordgo"
)

type arg struct {
	argName string
	argDesc string
}

func cmdEmbedWithArgs(cmd string, desc string, args []arg) *discordgo.MessageEmbed {
	embed := NewEmbed()
	title := "k!" + cmd
	for _, val := range args {
		title = title + " [" + val.argName + "] "
	}
	embed.SetTitle(title)
	embed.SetDescription(desc)
	for _, val := range args {
		embed.AddField(val.argName, val.argDesc)
	}
	embed.SetColor(0xfefe62)
	return embed.MessageEmbed
}

func cmdEmbed(cmd string, desc string) *discordgo.MessageEmbed {
	embed := NewEmbed()
	embed.SetTitle("k!" + cmd)
	embed.SetDescription(desc)
	embed.SetColor(0xfefe62)
	return embed.MessageEmbed
}

func help(s *discordgo.Session, m *discordgo.MessageCreate) {
	embed := NewEmbed()
	embed.SetTitle("toki! mi ilo Koje!")
	embed.SetDescription("Here are my commands:")
	embed.SetColor(0x000099)

	s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)

	page := dgwidgets.NewPaginator(s, m.ChannelID)
	page.Add(
		cmdEmbed("help", "This command."),
	)
	page.Add(
		cmdEmbedWithArgs(
			"define", "Define a toki pona word.",
			[]arg{arg{argName: "word", argDesc: "The word to define"}},
		),
	)
	page.Add(
		cmdEmbed("browse", "Browse through Toki Pona words."),
	)
	page.Add(
		cmdEmbedWithArgs(
			"etymology", "Get the etymology for a Toki Pona word.",
			[]arg{arg{argName: "word", argDesc: "The word to get etymology for"}},
		),
	)
	page.Add(
		cmdEmbed("etybrowse", "Browse through Toki Pona etymology."),
	)
	page.Add(
		cmdEmbedWithArgs(
			"quiz", "Get quizzed on Toki Pona words.",
			[]arg{arg{argName: "count", argDesc: "The number of words. Maximum 15."}},
		),
	)
	page.Add(
		cmdEmbedWithArgs(
			"sitelen", "Draw some text in Sitelen Pona.",
			[]arg{arg{argName: "text", argDesc: "The text to draw."}},
		),
	)
	page.Add(
		cmdEmbedWithArgs(
			"count", "Count in Toki Pona.",
			[]arg{arg{argName: "num", argDesc: "The number to count."}},
		),
	)
	page.Add(
		cmdEmbed("about", "About me, ilo Koje."),
	)
	page.SetPageFooters()
	page.Spawn()
}
