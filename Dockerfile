# Build the go application into a binary
FROM golang:alpine as builder
WORKDIR /app
ADD . ./
RUN CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -a -installsuffix cgo -o bin/ollama-discord-bot .

FROM alpine:3.16
ENV DISCORD_PUBLIC_KEY=""
ENV AI_URL=""
ENV BOT_NAME=""
ENV MODEL_NAME=""
ENV APP_HOME=/app
WORKDIR ${APP_HOME}
RUN apt-get update && \
    apt-get install -y curl && \
    rm -rf /var/lib/apt/lists/*
COPY --from=builder /app/bin/ollama-discord-bot ./bin/ollama-discord-bot
ENTRYPOINT ["/app/bin/ollama-discord-bot"]
