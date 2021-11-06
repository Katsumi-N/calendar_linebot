package main

import (
	"fmt"
	"log"

	"github.com/Katsumi-N/calendar_linebot/calenderutil"
	"github.com/Katsumi-N/calendar_linebot/config"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	ev := calenderutil.RetrieveEvents(10)

	// mi, h, d, m, y := calenderutil.ParseDate(e[0].Start)
	// fmt.Println(y, m, d, h, mi)

	bot, err := linebot.New(
		config.Config.ChannelSecret, config.Config.ChannelToken,
	)

	if err != nil {
		log.Fatal(err)
	}

	messagestr := "直近の予定です\n"
	for _, m := range ev {
		title := m.Title
		smin, shour, sday, smonth, _ := calenderutil.ParseDate(m.Start)
		emin, ehour, eday, emonth, _ := calenderutil.ParseDate(m.End)
		messagestr += fmt.Sprintf("%s %s/%s %s:%s - %s/%s %s:%s\n ", title, smonth, sday, shour, smin,
			emonth, eday, ehour, emin)

	}
	message := linebot.NewTextMessage(messagestr)

	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}

}
