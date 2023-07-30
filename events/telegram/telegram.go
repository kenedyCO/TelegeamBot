package telegram

import "github.com/kenedyCO/tgBot/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
	// storage
}

func New(client)
