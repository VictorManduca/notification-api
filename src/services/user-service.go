package services

import (
	config "notification-api/src/configuration"
	models "notification-api/src/models"
)

func CreateUser(user models.User) error {
	_, error := config.ConnectDb().Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	config.ConnectDb().Close()
	return error
}

func GetUsers() ([]*models.User, error) {
	users := make([]*models.User, 0)
	rows, error := config.ConnectDb().Query("SELECT name, email FROM users;")
	config.ConnectDb().Close()

	if error != nil {
		return nil, error
	}

	for rows.Next() {
		user := new(models.User)
		error := rows.Scan(&user.Email, &user.Name)
		if error != nil {
			return nil, error
		}

		users = append(users, user)
	}

	return users, nil
}
