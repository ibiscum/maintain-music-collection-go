package main

import (
	"log"
	"os"

	"github.com/ibiscum/maintain-music-collection-go/internal/helper"
	"github.com/joho/godotenv"
)

func main() {
	envFile := os.Args[1]
	logFileName := "cronitml.log"

	// open log file
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	// set log out put
	log.SetOutput(logFile)

	// optional: log date-time, filename, and line number
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	err = godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("error loading environment file: %s", err)
	}

	dbPath := os.Getenv("ITML_DB_PATH")
	if len(dbPath) == 0 {
		log.Fatal("specify the ITML_DB_PATH environment variable")
	}

	log.Println("start")
	tr := helper.LoadTracks()

	db, err := helper.ConnectDB(dbPath)
	if err != nil {
		log.Fatalf("error connecting DB: %s", err)
	}

	err = helper.InsertTracks(db, tr)
	if err != nil {
		log.Fatalf("error connecting DB: %s", err)
	}

	log.Println("stop")
}
