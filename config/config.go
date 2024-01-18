package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	discordToken string
	botName      string
	aiUrl        string
}

var c *Config

func init() {
	env, err := godotenv.Read("./.env")

	if err != nil {
		fmt.Println("Not found .env File!")
		fmt.Println("It fetches the specified environment variables, not from the .env file.")
		specifiedEnv()
	} else {
		envFile(env)
	}
}

func envFile(env map[string]string) {

	c = &Config{}

	discrodToken := env["DISCORD_PUBLIC_KEY"]

	if discrodToken != "" {
		c.discordToken = discrodToken
	} else {
		panic("Not found in .env file the 'DISCORD_PUBLIC_KEY'!")
	}

	botName := env["BOT_NAME"]

	if botName != "" {
		c.botName = botName
	} else {
		fmt.Println("Not found in .env file the 'BOT_NAME'!")
		fmt.Println("Use the default for 'BOT_NAME'.")
		c.botName = "!chat"
	}

	aiUrl := env["AI_URL"]

	if botName != "" {
		c.aiUrl = aiUrl
	} else {
		panic("Not found the 'AI_URL'!")
	}

	fmt.Println("Config complete!")
}

func specifiedEnv() {

	c = &Config{}

	discrodToken := os.Getenv("DISCORD_PUBLIC_KEY")

	if discrodToken != "" {
		c.discordToken = discrodToken
	} else {
		panic("Not found the 'DISCORD_PUBLIC_KEY'!")
	}

	botName := os.Getenv("BOT_NAME")

	if botName != "" {
		c.botName = botName
	} else {
		fmt.Println("Not found the 'BOT_NAME'!")
		fmt.Println("Use the default for 'BOT_NAME'.")
		c.botName = "!chat"
	}

	aiUrl := os.Getenv("AI_URL")

	if botName != "" {
		c.aiUrl = aiUrl
	} else {
		panic("Not found the 'AI_URL'!")
	}

	fmt.Println("Config complete!")
}

func (c *Config) GetDiscordToken() string {
	return c.discordToken
}

func (c *Config) GetBotName() string {
	return c.botName
}

func (c *Config) GetAiUrl() string {
	return c.aiUrl
}

func GetConfig() *Config {
	return c
}
