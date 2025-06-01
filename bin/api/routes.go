package api

import (
	"database/sql"
	"fmt"
	l "log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/observly/bin/processors"

	"github.com/observly/bin/structs"
)

var dbURL string
var dbPass string
var dbUser string
var dbName string

func ingestEvent(c *gin.Context) {

	var newEvent structs.Event

	if err := c.BindJSON(&newEvent); err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, newEvent)

	go processors.ProcessEvent(&newEvent, getDBConnection())

}

func getDBConnection() *sql.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?tls=false", dbUser, dbPass, dbURL, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	return db

}

func StartAPI() {

	err := godotenv.Load()

	if err != nil {
		l.Fatal("Error loading .env file")
	} else {
		dbURL = os.Getenv("DB_URL")
		dbUser = os.Getenv("DB_USER")
		dbPass = os.Getenv("DB_PASS")
		dbName = os.Getenv("DB_NAME")

		router := gin.Default()
		router.POST("/api/logs/ingest", ingestEvent)

		router.Run("localhost:8080")
	}

}
