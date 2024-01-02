package collections

import (
	"encoding/json"
	"middleware/example/internal/models"
	songs "middleware/example/internal/services/songs"
	"net/http"

	"github.com/sirupsen/logrus"
)

// GetCollections
// @Tags         collections
// @Summary      Get collections.
// @Description  Get collections.
// @Success      200            {array}  models.Collection
// @Failure      500             "Something went wrong"
// @Router       /collections [get]
func AddSong(w http.ResponseWriter, r *http.Request) {
	//register a new song

	var newsong models.Song

	err := json.NewDecoder(r.Body).Decode(&newsong)
	song, err := songs.AddSong(&newsong)

	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	if err != nil {
		// logging error
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			// writing http code in header
			w.WriteHeader(customError.Code)
			// writing error message in body
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		body, _ := json.Marshal(song)
		_, _ = w.Write(body)

	}
}
