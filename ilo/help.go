package ilo

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
	embed.SetDescription(_t("Here are my commands:", "toki lawa mi li ni:", m))
	embed.SetColor(0x000099)

	s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)

	page := dgwidgets.NewPaginator(s, m.ChannelID)
	page.Add(
		cmdEmbed("help", _t("This command.", "ni.", m)),
	)
	page.Add(
		cmdEmbedWithArgs(
			"define", _t("Define a toki pona word.", "kama sona e nimi pi toki pona", m),
			[]arg{arg{argName: "word", argDesc: _t("The word to define", "nimi seme li wile kama sona", m)}},
		),
	)
	page.Add(
		cmdEmbed("browse", _t("Browse through Toki Pona words.", "kama sona e nimi mute pi toki pona", m)),
	)
	page.Add(
		cmdEmbedWithArgs(
			"etymology", _t("Get the etymology for a Toki Pona word.", "kama sona e tan nimi pi toki pona", m),
			[]arg{arg{argName: "word", argDesc: _t("The word to get etymology for", "nimi tan seme li wile kama sona", m)}},
		),
	)
	page.Add(
		cmdEmbed("etybrowse", _t("Browse through Toki Pona etymology.", "kama sona e nimi tan mute pi toki pona", m)),
	)
	page.Add(
		cmdEmbedWithArgs(
			"quiz", _t("Get quizzed on Toki Pona words.", "mi toki e ni tawa sina: 'nimi li seme?'", m),
			[]arg{arg{argName: "count", argDesc: _t("The number of words. Maximum 15.", "nanpa nimi. mi wile nanpa ≤ 15", m)}},
		),
	)
	page.Add(
		cmdEmbedWithArgs(
			"sitelen", _t("Draw some text in Sitelen Pona.", "sitelen e nimi kepeken sitelen pona", m),
			[]arg{arg{argName: "text", argDesc: _t("The text to draw.", "nimi tawa sitelen", m)}},
		),
	)
	page.Add(
		cmdEmbedWithArgs(
			"count", _t("Count in Toki Pona.", "toki e nanpa kepeken toki pona.", m),
			[]arg{arg{argName: "num", argDesc: _t("The number to count.", "nanpa toki.", m)}},
		),
	)
	page.Add(
		cmdEmbed("about", _t("About me, ilo Koje.", "sona mi.", m)),
	)
	page.SetPageFooters()
	page.Spawn()
}
