package db

import (
	models "invokes/internal/models"
	"invokes/internal/utils"

	"gorm.io/gorm"
)

// AddTransaction to add an Transaction
func (db *Wrapper) AddTransaction(transaction *models.Transaction) error {
	if transaction == nil {
		return &DBError{Msg: "transaction struct isn't valid", Type: InvalidContent}
	}

	invoice := models.Invoice{}
	whereparams := models.Invoice{
		ID: transaction.InvoiceID,
	}
	result := db.GormDB.Preload("User").Where(&whereparams).First(&invoice)

	// Means already exist before ask to add
	if result.RowsAffected == 0 {
		utils.Logger.Error("Invoice not found")
		return &DBError{Msg: "Invoice not found", Type: InvoiceNotFound}
	}

	if invoice.Status == "paid" {
		utils.Logger.Error("Invoice has already been paid")
		return &DBError{Msg: "Invoice has already been paid", Type: InvoiceAlreadyPaid}
	}

	if invoice.Amount != transaction.Amount {
		utils.Logger.Error("Invoice amount isn't the same")
		return &DBError{Msg: "Invoice amount isn't the same", Type: InvoiceAmountNotFound}
	}

	invoice.Status = "paid"
	invoice.User.Balance += transaction.Amount

	db.GormDB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&invoice)

	return nil
}
