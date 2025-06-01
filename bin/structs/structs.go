package structs

import (
	"time"
)

type Event struct {
	EventID    int64
	TimeStamp  time.Time         `json:"timestamp"`
	Node       string            `json:"node"`
	Agent      string            `json:"agent"`
	Component  string            `json:"component"`
	State      string            `json:"state"`
	Message    string            `json:"message"`
	EventAttrs map[string]string `json:"attributes"`
}

type Notification struct {
	NotifyID int64
}

type Action struct {
	Notification
	*Event
	ActionID   int64
	ActionType string
}

type Webhook struct {
	Action
	TargetURL string
	Body      string
	Headers   map[string]string
}
