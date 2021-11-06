package main

import (
	"fmt"
	"log"

	"github.com/Katsumi-N/calendar_linebot/calenderutil"
	"github.com/Katsumi-N/calendar_linebot/config"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	e := calenderutil.RetrieveEvents(10)
	fmt.Println(e)
	// mi, h, d, m, y := calenderutil.ParseDate(e[0].Start)
	// fmt.Println(y, m, d, h, mi)

	bot, err := linebot.New(
		config.Config.ChannelSecret, config.Config.ChannelToken,
	)

	if err != nil {
		log.Fatal(err)
	}

	message := linebot.NewTextMessage(e[0].Title)

	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}

}
