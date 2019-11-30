package koje

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/dgraph-io/badger"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/ini.v1"
)

var db *badger.DB

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

type CommandStats struct {
	Cmds               []string
	CmdsRanAmountToday int
	LastTimeCmdRan     time.Time
}

var cStats CommandStats

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func saveCommands() {
	data, _ := json.Marshal(cStats)
	err := ioutil.WriteFile("./storage/commandsNow.json", data, 0644)
	check(err)
}

func loadCommands() {
	data, _ := ioutil.ReadFile("./storage/commandsNow.json")
	err := json.Unmarshal(data, &cStats)
	check(err)
}

func logCommand(m *discordgo.MessageCreate) {
	now := time.Now()
	if now.Month() != cStats.LastTimeCmdRan.Month() || now.Day() != cStats.LastTimeCmdRan.Day() {
		data, _ := ioutil.ReadFile("./storage/commandsNow.json")
		path := fmt.Sprintf("./storage/%d-%d-%d.json", cStats.LastTimeCmdRan.Year(), cStats.LastTimeCmdRan.Month(), cStats.LastTimeCmdRan.Day())
		err := ioutil.WriteFile(path, data, 0644)
		check(err)
		cStats.LastTimeCmdRan = now
		cStats.Cmds = []string{m.Content}
		cStats.CmdsRanAmountToday = 1
	} else {
		cStats.LastTimeCmdRan = now
		cStats.Cmds = append(cStats.Cmds, m.Content)
		cStats.CmdsRanAmountToday++
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.HasPrefix(m.Content, "k!") {
		cmds := map[string]cmd{
			"define":    pu,
			"d":         pu,
			"help":      help,
			"browse":    browse,
			"etymology": ety,
			"quiz":      quiz,
			"sitelen":   sitelenPona,
			"s":         sitelenPona,
			"etybrowse": browseety,
			"about":     about,
			"admin":     admin,
			"count":     count,
		}
		lex := strings.Split(strings.Split(m.Content, "!")[1], " ")
		if val, ok := cmds[lex[0]]; ok {
			if chanBlacklisted(m.ChannelID) && m.Author.ID != cfg.Section("Bot").Key("operator").String() {
				embed := NewEmbed().
					SetTitle("This channel is blacklisted for running commands.").
					SetColor(0xff0000)
				msg, _ := s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
				go func(msg *discordgo.Message) {
					time.Sleep(10 * time.Second)
					s.ChannelMessageDelete(msg.ChannelID, msg.ID)
				}(msg)
				return
			}
			go val(s, m)
			logCommand(m)
		}
	}
}

var cfg *ini.File

// Main function of ilo Koje
func Main() {
	defer saveCommands()

	loadCommands()
	var err error
	cfg, err = ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Failed to load config.ini")
		os.Exit(1)
	}
	db, err = badger.Open(badger.DefaultOptions("./storage/db"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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
