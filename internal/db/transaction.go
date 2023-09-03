package db

import (
	models "invokes/internal/models"
	"invokes/internal/utils"

	"gorm.io/gorm"
)

// AddTransaction to add an Transaction
func (db *Wrapper) AddTransaction(transaction *models.Transaction) error {
	if transaction == nil {
		return &DBError{Msg: "transaction struct isn't valid", Type: utils.InvalidContent}
	}

	invoice := models.Invoice{}
	whereparams := models.Invoice{
		ID: transaction.InvoiceID,
	}
	result := db.GormDB.Preload("User").Where(&whereparams).First(&invoice)

	// Means already exist before ask to add
	if result.RowsAffected == 0 {
		utils.Logger.Error("Invoice not found")
		return &DBError{Msg: "Invoice not found", Type: utils.InvoiceNotFound}
	}

	if err := checkAndUpdateInvoice(&invoice, transaction); err != nil {
		utils.Logger.Error("An error happened", err)
		return err
	}

	db.GormDB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&invoice)

	return nil
}

// checkAndUpdateInvoice check invoice according to transaction, update it if can
func checkAndUpdateInvoice(invoice *models.Invoice, transaction *models.Transaction) error {
	if invoice.Amount != transaction.Amount {
		return &DBError{Msg: "Invoice amount isn't the same", Type: utils.InvoiceAmountNotFound}
	}

	if invoice.Status == "paid" {
		return &DBError{Msg: "Invoice has already been paid", Type: utils.InvoiceAlreadyPaid}
	}

	invoice.Status = "paid"
	invoice.User.Balance += transaction.Amount

	return nil
}
