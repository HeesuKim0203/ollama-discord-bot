package app

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"

	"github.com/bwmarrin/discordgo"
	"github.com/discordgo-ai-bot/util"
)

func AiRequest(s *discordgo.Session, m *discordgo.MessageCreate, aiUrl string, modelName string, question string) {

	if question == "" {
		s.ChannelMessageSend(m.ChannelID, "Empty string. Please enter a question")
	}

	// Todo : "" -> err
	args := []string{
		"-X", "POST",
		aiUrl,
		"-d", fmt.Sprintf(`{
			"model":"mistral",
			"prompt":"%s"
		}`, question),
	}

	cmd := exec.Command("curl", args...)

	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating StdoutPipe for Cmd", err)
		aiErr(s, m)
		return
	}

	err = cmd.Start()
	if err != nil {
		fmt.Println("Error starting Cmd", err)
		aiErr(s, m)
		return
	}

	reader := bufio.NewReader(stdOut)
	line, err := reader.ReadString('\n')

	var response *util.Response

	result := ""

	for err == nil {
		err = json.Unmarshal([]byte(line), &response)
		if err != nil {
			log.Println(err)
			aiErr(s, m)
		} else {
			result += response.Data
		}
		line, err = reader.ReadString('\n')
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error waiting for Cmd", err)
		aiErr(s, m)
		return
	}

	s.ChannelMessageSend(m.ChannelID, result)
}

func aiErr(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "There was an error in the AI's response. Please ask your question again.")
}
