package koje

import (
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type tpCounts struct {
	wan  int
	tu   int
	luka int
	mute int
	ale  int
}

func count(s *discordgo.Session, m *discordgo.MessageCreate) {
	lex := strings.Split(strings.Split(m.Content, "!")[1], " ")
	if len(lex) < 2 {
		embed := NewEmbed().
			SetTitle("No number to count given!").
			SetColor(0xff0000)
		s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
		return
	}
	if i, err := strconv.Atoi(lex[1]); err != nil || i < 0 {
		embed := NewEmbed().
			SetTitle("Invalid number!").
			SetDescription("Numbers must be non-negative integers").
			SetColor(0xff0000)
		s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
		return
	}
	i, _ := strconv.Atoi(lex[1])
	if i == 0 {
		embed := NewEmbed().
			SetTitle("ala")
		s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
		return
	}
	var count tpCounts
	for i > 0 {
		if i >= 100 {
			count.ale = (i / 100)
			i = i - ((i / 100) * 100)
		}
		if i >= 20 {
			count.mute = (i / 20)
			i = i - ((i / 20) * 20)
		}
		if i >= 5 {
			count.luka = (i / 5)
			i = i - ((i / 5) * 5)
		}
		if i >= 2 {
			count.tu = (i / 2)
			i = i - ((i / 2) * 2)
		}
		if i >= 1 {
			count.wan = (i / 1)
			i = i - ((i / 1) * 1)
		}
	}
	counts := ""
	counts = counts + strings.Repeat("ale ", count.ale)
	counts = counts + strings.Repeat("mute ", count.mute)
	counts = counts + strings.Repeat("luka ", count.luka)
	counts = counts + strings.Repeat("tu ", count.tu)
	counts = counts + strings.Repeat("wan ", count.wan)
	if len(counts) > 256 {
		embed := NewEmbed().
			SetDescription(counts).
			SetColor(0xfefe62)
		s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
		return
	} else if len(counts) <= 2048 {
		embed := NewEmbed().
			SetTitle(counts).
			SetColor(0xfefe62)
		s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
		return
	} else {
		embed := NewEmbed().
			SetTitle("Your number is too big to send in chat.").
			SetColor(0xff0000)
		s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
		return
	}
}
