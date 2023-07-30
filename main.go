package main

import (
	"flag"
	"log"

	tgClient "github.com/kenedyCO/tgBot/clients/telegram"
	"github.com/kenedyCO/tgBot/consumer/eventconsumer"
	"github.com/kenedyCO/tgBot/events/telegram"
	"github.com/kenedyCO/tgBot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

func main() {

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)
	log.Print("service strated")

	consumer := eventconsumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for tg bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
