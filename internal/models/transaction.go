package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

// Transaction is a struct to describe a transaction
type Transaction struct {
	InvoiceID   int     `json:"invoice_id" binding:"required" example:"12"`
	Amount      int32   `json:"-" binding:"required" example:"49297"`
	AmountFloat float32 `json:"amount" example:"956.32"`
	Reference   string  `json:"reference" binding:"required" example:"JMPINV200220117"`
}

// TransactionJSON map the info about an transaction
type TransactionJSON struct {
	InvoiceID   int     `json:"invoice_id" binding:"required" example:"12"`
	AmountFloat float32 `json:"amount" binding:"required" example:"956.32"`
	Reference   string  `json:"reference" binding:"required" example:"JMPINV200220117"`
}

// AfterFind overload balance to fit to specs
func (i *Transaction) AfterFind(tx *gorm.DB) (err error) {
	i.AmountFloat = float32(i.Amount) / 100.0
	return
}

// UnmarshalJSON overload to handle float to int convertion
func (i *Transaction) UnmarshalJSON(data []byte) error {
	var res TransactionJSON

	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	i.InvoiceID = res.InvoiceID
	i.Reference = res.Reference
	i.Amount = int32(res.AmountFloat * 100)
	i.AmountFloat = res.AmountFloat

	return nil
}
