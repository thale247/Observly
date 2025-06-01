package processors

import (
	"database/sql"
	"fmt"

	"github.com/observly/bin/structs"
)

func ProcessEvent(incoming *structs.Event, db *sql.DB) {

	addEventToDatabase(incoming, db)
	db.Close()

}

func addEventToDatabase(incoming *structs.Event, db *sql.DB) {

	// Check the connection
	query := "INSERT INTO Event (TimeStamp, Node, Agent, Component, State, Message) VALUES (?, ?, ?, ?, ?, ?)"

	_, err := db.Exec(query, incoming.TimeStamp, incoming.Node, incoming.Agent, incoming.Component, incoming.State, incoming.Message)

	if err != nil {
		panic(err)
	}

	fmt.Println("Added to DB")

}
