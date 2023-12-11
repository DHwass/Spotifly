package collections

import (
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"

	"github.com/gofrs/uuid"
)

// GET ALL
func GetAllSongs() ([]models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM songs")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	songs := []models.Song{}
	for rows.Next() {
		var data models.Song
		err = rows.Scan(&data.Id, &data.Title, &data.Artist, &data.Duration)
		if err != nil {
			return nil, err
		}
		songs = append(songs, data)
	}

	_ = rows.Close()

	return songs, err
}

// GET
func GetSongById(id uuid.UUID) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM songs WHERE id = ?", id.String())
	helpers.CloseDB(db)

	var song models.Song
	err = row.Scan(&song.Id, &song.Title, &song.Artist, &song.Duration)
	if err != nil {
		return nil, err
	}

	return &song, err
}

// POST
func AddSong(song *models.Song) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	id, _ := uuid.NewV4()

	_, err = db.Exec("INSERT INTO songs (id, title, artist, duration) VALUES (?, ?, ?, ?)", id.String(), song.Title, song.Artist, song.Duration)
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	return song, err
}
