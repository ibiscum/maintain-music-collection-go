package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/ibiscum/maintain-music-collection-go/internal/helper"
	"github.com/joho/godotenv"
)

var db *sql.DB

// func addAlbum(a *Album) (int64, error) {
// 	result, err := db.ExecContext(
// 		context.Background(),
// 		`INSERT INTO album (title, artist, price) VALUES (?,?,?);`, a.Title, a.Artist, a.Price,
// 	)
// 	id, err := result.LastInsertId()
// 	if err != nil {
// 		return 0, err
// 	}
// 	return id, nil
// }

// func albumsByArtist(artist string) ([]AlbumDbRow, error) {

// 	var albums []AlbumDbRow
// 	rows, err := db.QueryContext(
// 		context.Background(),
// 		`SELECT * FROM album WHERE artist=?;`, artist,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var album AlbumDbRow
// 		if err := rows.Scan(
// 			&album.ID, &album.Title, &album.Artist, &album.Price,
// 		); err != nil {
// 			return nil, err
// 		}
// 		albums = append(albums, album)
// 	}
// 	return albums, err
// }

// func albumByID(id int) (AlbumDbRow, error) {
// 	var album AlbumDbRow
// 	row := db.QueryRowContext(
// 		context.Background(),
// 		`SELECT * FROM album WHERE id=?`, id,
// 	)
// 	err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
// 	if err != nil {
// 		return album, err
// 	}
// 	return album, nil
// }

func main() {
	envFile := os.Args[1]
	logFileName := "preparedb.log"

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

	db, err = helper.ConnectDB(dbPath)
	if err != nil {
		log.Fatal("error connecting to DB: ", err)
	}

	err = helper.InitDB(db, dbPath)
	if err != nil {
		log.Fatal("error initializing DB: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("error initializing DB connection: ping error: ", err)
	}

	log.Println("database initialized.")

	// err = insertTestData()
	// if err != nil {
	// 	log.Fatal("error inserting test data: ", err)
	// }
	// fmt.Println("test data inserted..")

	// fmt.Println("querying test data by album ID..")
	// // query back each record with IDs 1 - 4
	// for i := 1; i <= 4; i++ {
	// 	album, err := albumByID(i)
	// 	if err != nil {
	// 		fmt.Printf("error querying album ID: %d, %s\n", i, err)
	// 	} else {
	// 		fmt.Printf("%v\n", album)
	// 	}
	// }
}
