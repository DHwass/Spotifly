package collections

import (
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"time"

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
		err = rows.Scan(&data.Id, &data.Title, &data.Artist, &data.Filename, &data.Published)
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
	err = row.Scan(&song.Id, &song.Title, &song.Artist, &song.Filename, &song.Published)
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
	song.Published = time.Now().Format("2006-01-02 15:04:05")
	_, err = db.Exec("INSERT INTO songs (id, title, artist, filename, published) VALUES (?, ?, ?, ?,?)", id.String(), song.Title, song.Artist, song.Filename, song.Published)
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	return song, nil
}

// DELETE
func DeleteSongById(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM songs WHERE id = ?", id.String())
	helpers.CloseDB(db)
	if err != nil {
		return err
	}

	return err
}

// PUT
func UpdateSongById(id uuid.UUID, song *models.Song) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("UPDATE songs SET title = ?, artist = ?, filename = ? WHERE id = ?", song.Title, song.Artist, song.Filename, id.String())
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	return song, err
}
