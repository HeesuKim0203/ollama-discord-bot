package app

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/discordgo-ai-bot/util"
)

func AiRequest(s *discordgo.Session, m *discordgo.MessageCreate, aiUrl string, modelName string, question string) {

	if question == "" {
		s.ChannelMessageSend(m.ChannelID, "Empty string. Please enter a question")
	}

	requestData := util.NewRequst(question, modelName)

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		log.Fatalf("Error marshalling request data: %v", err)
	}

	response, err := http.Post(aiUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error making POST request: %v", err)
	}
	defer response.Body.Close()

	reader := bufio.NewReader(response.Body)
	var responseChunks []string

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break // 파일의 끝에 도달함
		}
		if err != nil {
			log.Fatalf("Error reading chunk: %v", err)
		}
		responseChunks = append(responseChunks, line)
	}

	log.Println(responseChunks)
}

func aiErr(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "There was an error in the AI's response. Please ask your question again.")
}
