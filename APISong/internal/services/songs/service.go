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

// POST
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

// DELETE
func DeleteSong(id uuid.UUID) error {
	err := repository.DeleteSongById(id)
	if err != nil {
		logrus.Errorf("error deleting song : %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return err
}

// UPDATE
func ModifySong(id uuid.UUID, song *models.Song) (*models.Song, error) {
	// calling repository
	song, err := repository.UpdateSongById(id, song)
	// managing errors
	if err != nil {
		logrus.Errorf("error updating song : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return song, nil
}
