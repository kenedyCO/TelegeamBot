package main

import (
	"context"
	"flag"
	"log"

	tgClient "github.com/kenedyCO/tgBot/clients/telegram"
	"github.com/kenedyCO/tgBot/consumer/eventconsumer"
	"github.com/kenedyCO/tgBot/events/telegram"
	"github.com/kenedyCO/tgBot/storage/files/sqlite"
)

const (
	tgBotHost         = "api.telegram.org"
	sqliteStoragePath = "data/sqlite/storage.db"
	batchSize         = 100
)

func main() {
	//s := files.New(storagePath)
	s, err := sqlite.New(sqliteStoragePath)
	if err != nil {
		log.Fatal("cant connect to storage", err)
	}

	if err := s.Init(context.TODO()); err != nil {
		log.Fatal("cant init storage", err)
	}

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		s,
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
