package app

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/discordgo-ai-bot/config"
	"github.com/discordgo-ai-bot/util"
)

var (
	c       = config.GetConfig()
	botName = c.GetBotName()
)

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	content := strings.Split(m.Content, " ")

	if strings.EqualFold(content[0], botName) {
		text := ""
		for _, v := range content[1:] {
			text += v
		}

		args := []string{
			"-X", "POST",
			"http://localhost:11434/api/generate",
			"-d", fmt.Sprintf(`{
				"model":"mistral",
				"prompt":"%s"
			}`, text),
		}

		response := exec.Command("curl", args...)

		byteData, err := response.Output()
		if err != nil {
			log.Println("RunStart Error:", err)
			return
		}

		result := string(byteData)

		var prev string
		prevNum := 0
		jsonString := ""

		for index, data := range result {
			if data == '}' {

				prev = result[prevNum : index+1]
				prevNum = index + 1

				if index != len(result)-2 {
					jsonString += prev + ","
				} else {
					jsonString += prev
				}
			}
		}

		str := "[" + jsonString + "]"

		var test []interface{}
		var responses []*util.Response

		err = json.Unmarshal([]byte(str), &test)

		if err != nil {
			log.Println(err)
		}

		for index, item := range test {

			if index != len(test)-1 {
				response := item.(map[string]interface{})

				model := response["model"].(string)
				createdAt, _ := time.Parse(time.RFC3339, response["created_at"].(string))
				data := response["response"].(string)
				done := response["done"].(bool)

				responses = append(responses, util.NewResponse(model, createdAt, data, done))
			}
		}

		resultText := ""

		for _, item := range responses {

			resultText += item.GetData()
		}

		s.ChannelMessageSend(m.ChannelID, resultText)

		return
	}

}

func NewDiscord() *discordgo.Session {

	discord, err := discordgo.New("Bot " + c.GetDiscordToken())

	if err != nil {
		fmt.Println("discord Create Error : ")
		panic(err)
	}

	// Logging Server
	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	discord.AddHandler(messageHandler)

	return discord
}
