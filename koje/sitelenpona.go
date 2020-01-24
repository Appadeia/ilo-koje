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

func validFont(name string) bool {
	switch name{
	case
		"linja pona",
		"linja pimeja":
		return true
    	}
    return false
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

	text := strings.Join(lex[1:], " ")
	fontname := "linja pimeja" // default font
	fontsize := "50"
	
	extra := strings.Split(text, ", ")"
	if len(extra) > 1 {
		lastextra := extra[len(extra)-1]
		extra = strings.SplitN(lastextra, " ", 2) // = ["first word after last comma", "rest of text after last comma"]
		if len(extra)==2 && (extra[0] == "kepeken" || extra[0] == "font") {
			if validFont(extra[1]) {
				fontname = extra[1]
				text = text[:len(text)-len(lastextra)-2] // dont render the options
			} else {
				embed := NewEmbed().
					SetTitle(_t("Error!", "pakala!", m)).
					SetDescription(_t("Unrecognised font: "+extra[1], "mi jo ala e sitelen pi "+extra[1], m)).
					SetColor(0xff0000)
				s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
				return
			}
		}
	}
	
	
	filename := "/tmp/" + randSeq(10) + ".png"
	cmd := exec.Command("pango-view", "-t", text, "--font", fontname+" "+fontsize, "-o", filename, "-q", "--align=center", "--hinting=slight", "--antialias=gray", "--margin=10px")
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
