package helper

import (
	"context"
	"crypto/md5"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"unsafe"

	"github.com/dhowden/itl"
	_ "modernc.org/sqlite"
)

func HexMD5(text string) string {
	data := []byte(text)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func LoadTracks() map[string]itl.Track {

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

func InsertTracks(db *sql.DB, tracks map[string]itl.Track) error {
	var err error

	sql_i_t_itml := `INSERT INTO itunes_data (
        persistent_id,
        track_id,
        track_name,
        artist,
        album_artist,
        album, genre,
        disc_number,
        disc_count,
        track_number,
        track_count,
        album_year,
        date_modified,
        date_added,
        volume_adjustment,
        play_count,
        play_date_utc,
        artwork_count,
        md5_id
    ) VALUES (
        ?, ?, ?, ?, ?, ?, ?, ?, ?,
        ?, ?, ?, ?, ?, ?, ?, ?, ?
    ) ON CONFLICT (persistent_id) DO UPDATE SET
        track_id          = EXCLUDED.track_id,
        track_name        = EXCLUDED.track_name,
        artist            = EXCLUDED.artist,
        album_artist      = EXCLUDED.album_artist,
        album             = EXCLUDED.album,
        genre             = EXCLUDED.genre,
        disc_number       = EXCLUDED.disc_number,
        disc_count        = EXCLUDED.disc_count,
        track_number      = EXCLUDED.track_number,
        track_count       = EXCLUDED.track_count,
        album_year        = EXCLUDED.album_year,
        date_modified     = EXCLUDED.date_modified,
        date_added        = EXCLUDED.date_added,
        volume_adjustment = EXCLUDED.volume_adjustment,
        play_count        = EXCLUDED.play_count,
        play_date_utc     = EXCLUDED.play_date_utc,
        artwork_count     = EXCLUDED.artwork_count,
        md5_id            = EXCLUDED.md5_id;`

	stmt, err := db.Prepare(sql_i_t_itml)
	if err != nil {
		return err
	}

	for _, v := range tracks {
		v_size := int(unsafe.Sizeof(itl.Track{}))
		fmt.Println(v_size, ":", &v)
		md5_id := "abc"
		// track_id_b := []byte(strconv.Itoa(v.TrackID))
		// track_name := []byte(v.Name)

		// md5_in := []byte(v.TrackID)
		stmt.Exec(v.TrackID, v.Name, v.Artist, v.AlbumArtist, v.Album, v.Genre, v.DiscNumber, v.DiscCount,
			v.TrackNumber, v.TrackCount, v.Year, v.DateModified, v.DateAdded, v.VolumeAdjustment, v.PlayCount,
			v.PlayDateUTC, v.ArtworkCount, md5_id)
	}
	return nil
}

func ConnectDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	// Ping the database to ensure connectivity
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitDB(db *sql.DB, dbPath string) error {
	var err error
	sql_c_t_itml := `CREATE TABLE IF NOT EXISTS t_itml (
        persistent_id TEXT PRIMARY KEY,
		track_id INTEGER,
		track_name TEXT,
		artist TEXT,
		album_artist TEXT,
		album TEXT,
		genre TEXT,
		disc_number INTEGER,
        disc_count INTEGER,
        track_number INTEGER,
        track_count INTEGER,
        album_year TEXT,
        date_modified TEXT,
        date_added TEXT,
        volume_adjustment INTEGER,
        play_count INTEGER,
        play_date_utc TEXT,
        artwork_count INTEGER,
        md5_id TEXT);`

	_, err = db.ExecContext(context.Background(), sql_c_t_itml)
	if err != nil {
		return err
	}
	return nil
}
