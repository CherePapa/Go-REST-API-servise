package main

import (
	"flag"
	"log"
	tgClient "telegram-bot/clients/telegram"
	eventconsumer "telegram-bot/consumer/event-consumer"
	"telegram-bot/events/telegram"
	"telegram-bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

// 8077488608:AAHboPmqvKIqVLBmfvNPcS96E-OOA2pKuXs
func main() {

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("servis started")

	consumer := eventconsumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to tg-bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not implaned")
	}

	return *token
}
