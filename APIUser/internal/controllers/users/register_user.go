package collections

import (
	"encoding/json"
	"middleware/example/internal/models"
	Users "middleware/example/internal/services/users"
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
func RegUser(w http.ResponseWriter, r *http.Request) {
	//register a new user
	var newUser models.Users

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}
	user, err := Users.CreateUser(&newUser)

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
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	body, _ := json.Marshal(user)
	_, _ = w.Write(body)
	return

}
