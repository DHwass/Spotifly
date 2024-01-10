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

	userr, err := Users.ModifyUser(UserId, &User)

	if err != nil {
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
		body, _ := json.Marshal(userr)
		_, _ = w.Write(body)

	}
}
