package app

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/discordgo-ai-bot/config"
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

		log.Println(text)

		args := []string{
			"-X", "POST",
			"http://localhost:11434/api/generate",
			"-d", fmt.Sprintf(`{
				"model":"mistral",
				"prompt":"%s"
			}`, text),
		}

		response := exec.Command("curl", args...)

		log.Println(response)

		result, err := response.Output()
		if err != nil {
			log.Println("RunStart Error:", err)
			return
		}

		log.Println(string(result))

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
