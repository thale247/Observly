package structs

import "time"

type Event struct {
	TimeStamp time.Time `json:"timestamp"`
	Node      string    `json:"node"`
	Agent     string    `json:"agent"`
	Component string    `json:"component"`
	State     string    `json:"state"`
	Message   string    `json:"message"`
	Tags      []string  `json:"tags"`
}
