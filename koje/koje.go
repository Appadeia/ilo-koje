package koje

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/ini.v1"
)

type wordOrigin struct {
	Language   string `json:"language"`
	SourceWord string `json:"word"`
}

type puEty struct {
	Word   string     `json:"word"`
	Origin wordOrigin `json:"origins"`
}

type puWord struct {
	Word     string     `json:"word"`
	Meanings [][]string `json:"meanings"`
}

type cmd func(*discordgo.Session, *discordgo.MessageCreate)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.HasPrefix(m.Content, "k!") {
		cmds := map[string]cmd{
			"define":    pu,
			"help":      help,
			"browse":    browse,
			"etymology": ety,
			"quiz":      quiz,
			"sitelen":   sitelenPona,
		}
		lex := strings.Split(strings.Split(m.Content, "!")[1], " ")
		if val, ok := cmds[lex[0]]; ok {
			val(s, m)
		}
	}
}

// Main function of ilo Koje
func Main() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Failed to load config.ini")
		os.Exit(1)
	}

	discord, err := discordgo.New("Bot " + cfg.Section("Bot").Key("token").String())
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
	}

	discord.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return
	}

	fmt.Println("ilo Koje is now running.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()
}
