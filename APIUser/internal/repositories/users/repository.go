package Users

import (
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
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
		err = rows.Scan(&data.Id, &data.Name, &data.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, data)
	}
	//closing rows
	_ = rows.Close()

	return users, err
}

func GetUserById(id int) (*models.Users, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM Users WHERE id=?", id)
	helpers.CloseDB(db)

	var data models.Users
	err = row.Scan(&data.Id, &data.Name, &data.Email)
	if err != nil {
		return nil, err
	}
	return &data, err
}
func CreateUser(user *models.Users) (*models.Users, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("INSERT INTO Users (name, email) VALUES (?, ?)", user.Name, user.Email)
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func DeleteUserByID(id int) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM Users WHERE id=?", id)
	helpers.CloseDB(db)
	if err != nil {
		return err
	}
	return err
}
