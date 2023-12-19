package collections

import (
	"encoding/json"
	"middleware/example/internal/models"
	songs "middleware/example/internal/services/songs"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// GetCollection
// @Tags         collections
// @Summary      Get a collection.
// @Description  Get a collection.
// @Param        id           	path      string  true  "Collection UUID formatted ID"
// @Success      200            {object}  models.Collection
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /collections/{id} [get]
func UpdateSong(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	SongId, _ := ctx.Value("songId").(uuid.UUID)
	var Song models.Song
	err := json.NewDecoder(r.Body).Decode(&Song)

	song, err := songs.ModifySong(SongId, &Song)

	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		body, _ := json.Marshal(song)
		_, _ = w.Write(body)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
