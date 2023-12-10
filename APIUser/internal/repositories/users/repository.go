package Users

import (
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"

	"github.com/gofrs/uuid"
)

func GetAllUsers() ([]models.Users, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM Users")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	users := []models.Users{}
	for rows.Next() {
		var data models.Users
		err = rows.Scan(&data.Id)
		if err != nil {
			return nil, err
		}
		users = append(users, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return users, err
}

func GetUserById(id uuid.UUID) (*models.Users, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM Users WHERE id=?", id.String())
	helpers.CloseDB(db)

	var user models.Users
	err = row.Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, err
}
func CreateUser(user *models.Users) (*models.Users, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("INSERT INTO Users (name, email) VALUES (?, ?)", user.Id, user.Name, user.Email)
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}
	return user, nil
}
