package main

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/dhowden/itl"
	"github.com/joho/godotenv"
)

func preparetracks(tracks map[string]itl.Track) {

	for _, v := range tracks {

		track_id := v.TrackID
		fmt.Println(reflect.TypeOf(track_id))

		track_name := v.Name
		fmt.Println(reflect.TypeOf(track_name))

		fmt.Println(v.PersistentID)
	}
}

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
	tr := helper.loadTracks()
	preparetracks(tr)
	log.Println("stop")
}
