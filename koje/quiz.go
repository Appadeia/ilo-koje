package koje

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func waitForMessage(s *discordgo.Session) chan *discordgo.MessageCreate {
	channel := make(chan *discordgo.MessageCreate)
	s.AddHandlerOnce(func(_ *discordgo.Session, e *discordgo.MessageCreate) {
		channel <- e
	})
	return channel
}

func quiz(s *discordgo.Session, m *discordgo.MessageCreate) {
	pu := PuDict()
	var words []puWord
	err := json.Unmarshal([]byte(pu), &words)

	lex := strings.Split(strings.Split(m.Content, "!")[1], " ")

	if len(lex) < 2 {
		embed := NewEmbed().
			SetTitle(_t("Error!", "pakala!", m)).
			SetDescription(_t("No question count given.", "sina pana e nanpa tawa mi ala", m)).
			SetColor(0xff0000)
		s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
		return
	}

	if err != nil {
		embed := NewEmbed().
			SetTitle(_t("Error!", "pakala!", m)).
			SetDescription(err.Error()).
			SetColor(0xff0000)
		s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
		return
	}

	if i, err := strconv.Atoi(lex[1]); err != nil || i >= 15 || i <= 0 {
		embed := NewEmbed().
			SetTitle(_t("Invalid count!", "nanpa ike a!", m)).
			SetColor(0xff0000)
		s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
		return
	}

	reps, _ := strconv.Atoi(lex[1])

	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	correct := 0

	var quizMessages []string
Wait:
	for i := 1; i <= reps; i++ {
		word := words[rand.Intn(len(words))]
		embed := NewEmbed()
		embed.SetTitle(word.Word)
		embed.SetDescription(_t("What does this word mean?", "nimi ni li seme?", m))
		embed.SetFooter(_t("Question "+strconv.Itoa(i)+" out of "+strconv.Itoa(reps), "nimi "+strconv.Itoa(i)+"/"+strconv.Itoa(reps), m))

		mesg, _ := s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
		quizMessages = append(quizMessages, mesg.ID)

		timeoutChan := make(chan int)
		go func() {
			time.Sleep(7 * time.Second)
			timeoutChan <- 0
		}()
		concat := ""
		for _, val := range word.Meanings {
			concat = concat + " " + val[1] + ","
		}
		var arr []string
		for _, val := range strings.Split(concat, ",") {
			arr = append(arr, strings.ToLower(strings.TrimSpace(val)))
		}
		var arr2 []string
		for _, item := range arr {
			for _, val := range strings.Split(item, " ") {
				if val == "" {
					continue
				}
				arr2 = append(arr2, strings.Trim(val, ",! "))
			}
		}
		for {
			select {
			case usermsg := <-waitForMessage(s):
				quizMessages = append(quizMessages, usermsg.Message.ID)
				if usermsg.Author.ID != m.Author.ID {
					continue
				}
				if strings.Contains(usermsg.Content, "cancel") {
					embed := NewEmbed()
					embed.SetTitle(_t("Quiz Cancelled!", "tenpo ni la ni li lon ala", m))
					s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
					s.ChannelMessagesBulkDelete(m.ChannelID, quizMessages)
					return
				}
				for _, item := range strings.Split(usermsg.Content, " ") {
					for _, val := range arr2 {
						if strings.ToLower(strings.TrimSpace(item)) == val {
							embed := NewEmbed()
							embed.SetTitle(_t("Correct!", "pona a!", m))
							embed.SetColor(0x00ff00)
							correct = correct + 1
							m, _ := s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
							quizMessages = append(quizMessages, m.ID)
							s.ChannelTyping(m.ChannelID)
							time.Sleep(3 * time.Second)
							continue Wait
						}
					}
				}
			case <-timeoutChan:
				embed := NewEmbed()
				embed.SetTitle(_t("Time's up, here's what it means!", "tenpo ala a! ni li nimi:", m))
				for _, meaning := range word.Meanings {
					embed.AddField(meaning[0], meaning[1])
				}
				embed.SetColor(0xff0000)
				m, _ := s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
				quizMessages = append(quizMessages, m.ID)
				s.ChannelTyping(m.ChannelID)
				time.Sleep(3 * time.Second)
				continue Wait
			}
		}
	}
	embed := NewEmbed()
	embed.SetTitle(_t("Quiz Results ("+strconv.Itoa(reps)+" Questions)", "toki "+strconv.Itoa(reps), m))
	embed.SetColor(0x0000ff)
	embed.AddField(_t("Correct Answers", "toki pona", m), strconv.Itoa(correct))
	embed.AddField(_t("Correct Answers", "toki ike", m), strconv.Itoa(reps-correct))
	embed.SetFooter(_t(m.Author.Username+"'s Quiz", "ni li ijo pi "+m.Author.Username, m), m.Author.AvatarURL(""))
	s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
	s.ChannelMessagesBulkDelete(m.ChannelID, quizMessages)
}
