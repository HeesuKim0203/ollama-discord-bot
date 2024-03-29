# Ai bot

This bot connects the output of [Ollama](https://ollama.ai/), enabling you to ask and answer questions on Discord using Ollama! Try asking and answering questions on Discord with this bot and see Ollama in action!

Before you start, please [download](https://ollama.ai/) and run the model from Ollama!

## ENV
The [godotenv](godotenv) library is in use.  
Do not include this value in a public repository.

### .env
|Name|Description|type|Default|
|:---|:---|:---:|:---:|
|DISCORD_PUBLIC_KEY|Please look at the [discord bot guide](https://discord.com/developers/docs/getting-started#step-1-creating-an-app) and create a token and put it in.|string|X|
|AI_URL|Accesses an AI server running locally and receives JSON formatted data through this address. In cases where it is downloaded from Docker Hub, the API address of the local AI should not be entered as localhost, but as host.docker.internal.|string|X|
|MODEL_NAME|Please specify the name of the model you want to run and receive output from in Ollama! Example: mistral|string|X|
|BOT_NAME|Indicates the name of the bot and the first command.(Case insensitive)|string|"!chat"|

## Command

You must use it with the commands set in the ```BOT_NAME``` mentioned above.

```
!chat [Message]
```

#### question

```
!chat What is the capital of the United States?
```

#### result

```
I believe you are asking for the capital city of the United States of America. The answer to that question is Washington, D.C. (District of Columbia). It's important to note that while "Washington" is often colloquially used to refer to the entire metropolitan area, the term "capital" specifically refers to the governmental seat or hub of political power within a country or region. In this case, Washington, D.C., fulfills that role for the United States.
```

## Docker Hub

The image file is provided on [Docker hub](https://hub.docker.com/r/heesu998/ollama-discord-bot)!  
After pulling it, you can use it by entering the commands in order.

#### Pull docker image
```
docker pull heesu998/ollama-discord-bot
```

#### Running a container
```
sudo docker run -e DISCORD_PUBLIC_KEY=[your discord public token] -e AI_URL=[your local ai api address] -e MODEL_NAME=[your ollama model name] --name [container name] heesu998/ollama-discord-bot
```
