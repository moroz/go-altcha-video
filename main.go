package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/moroz/go-altcha-video/config"
	"github.com/moroz/go-altcha-video/db/queries"
	"github.com/moroz/go-altcha-video/handlers"
	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", config.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	queries.New(db).VacuumUsedChallenges(context.Background())

	e := handlers.Router(db)

	e.Logger.Fatal(e.Start(config.ListenOn))
}
