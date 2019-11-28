package koje

import (
	"math/rand"
	"os"
	"os/exec"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func sitelenPona(s *discordgo.Session, m *discordgo.MessageCreate) {
	lex := strings.Split(strings.Split(m.Content, "!")[1], " ")

	if len(lex) < 2 {
		embed := NewEmbed().
			SetTitle(_t("Error!", "pakala!", m)).
			SetDescription(_t("No words given to render!", "sina pana e nimi tawa mi ala!", m)).
			SetColor(0xff0000)
		s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
		return
	}

	text := ""
	for _, word := range lex[1:] {
		text = text + " " + word
		text = strings.TrimSpace(text)
	}
	filename := "/tmp/" + randSeq(10) + ".png"
	cmd := exec.Command("pango-view", "-t", text, "--font", "linja pona 50", "-o", filename, "-q", "--align=center", "--hinting=slight", "--antialias=gray", "--margin=10px")
	if err := cmd.Run(); err != nil {
		embed := NewEmbed().
			SetTitle("Internal Error").
			SetDescription(err.Error()).
			SetColor(0xff0000)
		s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
		return
	}
	f, err := os.Open(filename)
	if err != nil {
		embed := NewEmbed().
			SetTitle("Internal Error").
			SetDescription(err.Error()).
			SetColor(0xff0000)
		s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
		return
	}
	ms := &discordgo.MessageSend{
		Files: []*discordgo.File{
			&discordgo.File{
				Name:   filename,
				Reader: f,
			},
		},
	}
	s.ChannelMessageSendComplex(m.ChannelID, ms)
}
