package calenderutil

import (
	"fmt"
	"log"
	"time"
)

type Schedule struct {
	Title    string
	Location string
	Year     string
	Month    string
	Day      string
	Start    string
	End      string
}

func ShowEvents() {
	clt := SetClient()
	t := time.Now().Format(time.RFC3339)
	events, err := clt.Events.List("primary").ShowDeleted(false).SingleEvents(true).
		TimeMin(t).MaxResults(10).OrderBy("startTime").Do()

	if err != nil {
		log.Fatalf("Unable to retrieve next 10 user events")
	}
	fmt.Println("Upcoming events:")

	for _, item := range events.Items {
		date := item.Start.DateTime
		if date == "" {
			date = item.Start.Date
		}
		fmt.Printf("%v %v\n", item.Summary, date)
	}
}
