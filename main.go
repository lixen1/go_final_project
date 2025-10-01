package main

import (
	"go_final_project/pkg/db"
	"go_final_project/pkg/server"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

func main() {
	dbPath := os.Getenv("TODO_DBFILE")
	if dbPath == "" {
		dbPath = "scheduler.db"
	}

	err := db.Init("./" + dbPath)
	if err != nil {
		log.Fatalf("DB error: %v", err)
		return
	}
	defer db.DB.Close()

	server.Run()
}
