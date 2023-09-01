package db

import (
	models "invokes/internal/models"
)

// GetUsers returns the list of users
func (db *Wrapper) GetUsers() ([]*models.User, int, error) {

	var users []*models.User
	result := db.GormDB.Order("id").Find(&users)

	return users, len(users), result.Error
}
