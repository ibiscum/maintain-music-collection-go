package helper

import (
	"log"
	"os"
	"path/filepath"

	"github.com/dhowden/itl"
)

func loadTracks() map[string]itl.Track {

	ITML_FILE := os.Getenv("ITML_FILE")
	ITML_DIR := os.Getenv("ITML_DIR")

	p := filepath.Join(ITML_DIR, "/", ITML_FILE)

	f, err := os.Open(p)
	if err != nil {
		log.Fatalf("error opening plist file: %s", err)
	}

	l, err := itl.ReadFromXML(f)
	if err != nil {
		log.Fatalf("error reading plist file: %s", err)
	}

	return l.Tracks
}
