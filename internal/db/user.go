package db

import (
	models "invokes/internal/models"
)

// GetUsers returns the list of users
func (db *Wrapper) GetUsers(filter *models.User) ([]*models.User, int, error) {

	var users []*models.User
	query := db.GormDB
	if filter != nil {
		query = query.Where(filter)
	}
	result := query.Order("id").Find(&users)

	return users, len(users), result.Error
}
