package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime"
	"time"

	"path/filepath"

	"github.com/carlescere/scheduler"
	"github.com/dhowden/itl"
	"github.com/joho/godotenv"
)

func loadtracks() map[string]itl.Track {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	//DB_HOST := os.Getenv("DB_HOST")
	//DB_SCHEMA := os.Getenv("DB_SCHEMA")
	//DB_PASSWORD := os.Getenv("DB_PASSWORD")
	ITUNES_LIBRARY_FILE := os.Getenv("ITUNES_LIBRARY_FILE")
	ITUNES_MUSIC_DIR := os.Getenv("ITUNES_MUSIC_DIR")

	p := filepath.Join(ITUNES_MUSIC_DIR, ITUNES_LIBRARY_FILE)
	fmt.Println("p:", p)

	f, err := os.Open(p)
	if err != nil {
		log.Fatalf("Error opening plist file: %s", err)
	}

	l, err := itl.ReadFromXML(f)
	if err != nil {
		log.Fatalf("Error reading plist file: %s", err)
	}
	return l.Tracks
}

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
	job := func() {
		t := time.Now()
		tr := loadtracks()
		preparetracks(tr)
		fmt.Println("Time's up! @", t.UTC())
	}

	// Run now and every X.
	scheduler.Every(5).Minutes().Run(job)

	// Keep the program from not exiting.
	runtime.Goexit()
}
