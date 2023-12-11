package collections

import (
	"database/sql"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/songs"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// GET ALL
func GetAllSongs() ([]models.Song, error) {
	songs, err := repository.GetAllSongs()
	if err != nil {
		logrus.Errorf("error retrieving songs : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return songs, err
}

// GET (plus d'exemple à suivre D:)

func GetSongById(id uuid.UUID) (*models.Song, error) {
	song, err := repository.GetSongById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &models.CustomError{
				Message: "song not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving songs : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return song, err
}

// POST (tout seul comme un grand :D)
func AddSong(song *models.Song) (*models.Song, error) {
	song, err := repository.AddSong(song)
	if err != nil {
		logrus.Errorf("error adding song : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return song, err
}
