package collections

import (
	"encoding/json"
	"middleware/example/internal/models"
	Users "middleware/example/internal/services/users"
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
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	UserId, _ := ctx.Value("UserId").(uuid.UUID)
	var User models.Users
	err := json.NewDecoder(r.Body).Decode(&User)

	user, err := Users.ModifyUser(UserId, &User)

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
		body, _ := json.Marshal(user)
		_, _ = w.Write(body)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
