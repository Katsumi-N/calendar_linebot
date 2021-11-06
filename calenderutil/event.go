package calenderutil

import (
	"fmt"
	"strconv"
	"time"
)

type Schedule struct {
	Title    string
	Location string
	Start    string
	End      string
}

func (s *Schedule) SetSchedule(title string, location string, start string, end string) {
	s.Title = title
	s.Location = location
	s.Start = start
	s.End = end
}

func RetrieveEvents(eventNum int) []Schedule {
	clt := SetClient()
	t := time.Now().Format(time.RFC3339)
	events, err := clt.Events.List("primary").ShowDeleted(false).SingleEvents(true).
		TimeMin(t).MaxResults(int64(eventNum)).OrderBy("startTime").Do()

	if err != nil {
		fmt.Printf("Unable to retrieve next %v user events", eventNum)
	}
	fmt.Println("Upcoming events:")

	retEvents := make([]Schedule, len(events.Items))
	for i, item := range events.Items {
		date := item.Start.DateTime
		enddate := item.End.DateTime
		if date == "" {
			date = item.Start.Date
		}
		if enddate == "" {
			enddate = item.End.Date
		}
		retEvents[i].SetSchedule(item.Summary, item.Location, date, enddate)
		//fmt.Printf("%v %v\n", item.Summary, date)
	}

	return retEvents
}

func ParseDate(date string) (string, string, string, string, string) {
	layout := "2006-01-02T15:04:05+09:00"
	t, _ := time.Parse(layout, date)
	fmt.Println(t)
	hour, minutes := strconv.Itoa(t.Hour()), strconv.Itoa(t.Minute())
	if len(hour) == 1 {
		hour = "0" + hour
	}
	if len(minutes) == 1 {
		minutes = "0" + minutes
	}
	year, month, day := t.Date()
	return minutes, hour, strconv.Itoa(day), strconv.Itoa(int(month)), strconv.Itoa(year)

}
