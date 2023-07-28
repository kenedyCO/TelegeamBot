package main

import (
	"flag"
	"log"

	"github.com/kenedyCO/tgBot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(tgBotHost, mustToken())

	// tgClient = telegram.New(token)
	// fetcher = fetcher.New()
	// processor = processor.New()
	// Fconsumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String("token-bot-token", "", "token for tg bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
