package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func lineNotify() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	bot, err := linebot.New(os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	// append some message to messages

	if _, err = bot.BroadcastMessage(linebot.NewTextMessage("Hello world")).Do(); err != nil {
		log.Fatal(err)
	}
}
