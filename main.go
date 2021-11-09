package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Katsumi-N/calendar_linebot/calenderutil"
	"github.com/Katsumi-N/calendar_linebot/config"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	timer(16, 56, 1)
}

func timer(shour int, sminute int, limitday int) {
	t := time.NewTicker(1 * time.Minute)
	for {
		select {
		case <-t.C:
			fmt.Println(time.Now())
			// 月曜日に一週間の予定を送る
			if time.Now().Weekday() == 1 && time.Now().Hour() == shour && time.Now().Minute() == sminute {
				sendRecentSchedule(100, 7, "今週一週間の予定です．張り切っていきましょー！\n")
			}
			// １日の予定を毎日送る
			if time.Now().Hour() == shour && time.Now().Minute() == sminute {
				sendRecentSchedule(100, limitday, "今日の予定です．忘れずにね！\n")
			}
		}
	}
}

func sendRecentSchedule(eventNum int, limitday int, header string) {
	ev := calenderutil.RetrieveEvents(eventNum, limitday)

	bot, err := linebot.New(
		config.Config.ChannelSecret, config.Config.ChannelToken,
	)

	if err != nil {
		log.Fatal(err)
	}

	messagestr := header + "\n"
	for _, m := range ev {
		title := m.Title
		smin, shour, sday, smonth, _ := calenderutil.ParseDate(m.Start)
		emin, ehour, eday, emonth, _ := calenderutil.ParseDate(m.End)
		messagestr += fmt.Sprintf("〇%s\n %s/%s %s:%s - %s/%s %s:%s\n", title, smonth, sday, shour, smin,
			emonth, eday, ehour, emin)

	}
	message := linebot.NewTextMessage(messagestr)

	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}
