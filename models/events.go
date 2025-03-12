package models

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserId      int
}

var events = []Event{}

func (e Event) Save() {
	// this is a GET method to add events to a database in the future
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
