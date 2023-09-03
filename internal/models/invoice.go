package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

// Invoices alias
type Invoices = []*Invoice

// Invoice map the info about an invoice
type Invoice struct {
	ID          int     `json:"invoice_id" example:"12"`
	UserID      int     `json:"user_id" binding:"required" example:"12"`
	Status      string  `json:"status" example:"pending"`
	Label       string  `json:"label" binding:"required" example:"Findus"`
	Amount      int32   `json:"-" binding:"required" example:"49297"`
	AmountFloat float64 `gorm:"-" json:"amount" example:"492.97"`
	User        User    `gorm:"foreignKey:user_id" json:"-" binding:"-"`
}

// InvoiceJSON map the info about an invoice
type InvoiceJSON struct {
	ID          int     `json:"invoice_id" example:"12"`
	UserID      int     `json:"user_id" binding:"required" example:"12"`
	Status      string  `json:"status" example:"pending"`
	Label       string  `json:"label" binding:"required" example:"Findus"`
	AmountFloat float64 `json:"amount" binding:"required" example:"492.97"`
}

// AfterFind overload balance to fit to specs
func (i *Invoice) AfterFind(tx *gorm.DB) (err error) {
	i.AmountFloat = float64(i.Amount) / 100.0
	return
}

// UnmarshalJSON overload to handle float to int conversion
func (i *Invoice) UnmarshalJSON(data []byte) error {
	var res InvoiceJSON

	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	i.ID = res.ID
	i.Status = res.Status
	i.UserID = res.UserID
	i.Label = res.Label
	i.Amount = int32(res.AmountFloat * 100)
	i.AmountFloat = res.AmountFloat

	return nil
}
