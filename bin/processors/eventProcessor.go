package processors

import (
	"database/sql"

	"github.com/observly/bin/structs"
)

func ProcessEvent(incoming *structs.Event, db *sql.DB) {

	eventID, err := addEventToDatabase(incoming, db)

	if err != nil {
		panic(err)
	}

	incoming.EventID = eventID

	outgoing := structs.Action{

		Event:      incoming,
		ActionID:   0,
		ActionType: "",
	}

	go ProcessAction(&outgoing, db)

	db.Close()

}

func addEventToDatabase(incoming *structs.Event, db *sql.DB) (int64, error) {

	query := "INSERT INTO Event (TimeStamp, Node, Agent, Component, State, Message) VALUES (?, ?, ?, ?, ?, ?)"

	res, err := db.Exec(query, incoming.TimeStamp, incoming.Node, incoming.Agent, incoming.Component, incoming.State, incoming.Message)

	if err != nil {
		return -1, err
	}

	lastID, _ := res.LastInsertId()

	query = "INSERT INTO EventAttr (EventID, EventAttrKey, EventAttrValue) VALUES (?, ?, ?)"

	for key, value := range incoming.EventAttrs {
		_, err = db.Exec(query, lastID, key, value)
		if err != nil {
			return -1, err
		}
	}

	return lastID, nil

}
