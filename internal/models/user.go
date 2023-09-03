package models

import "gorm.io/gorm"

// Users alias
type Users = []*User

// User map the info about a user
type User struct {
	UserID       int     `gorm:"column:id" json:"user_id" binding:"required" example:"12"`
	FirstName    string  `json:"first_name" binding:"required" example:"Kevin"`
	LastName     string  `json:"last_name" binding:"required" example:"Findus"`
	Balance      int32   `json:"-" binding:"required" example:"49297"`
	BalanceFloat float64 `gorm:"-" json:"balance" binding:"required" example:"492.97"`
}

// AfterFind overload balance to fit to specs
func (u *User) AfterFind(tx *gorm.DB) (err error) {
	u.BalanceFloat = float64(u.Balance) / 100.0
	return
}
