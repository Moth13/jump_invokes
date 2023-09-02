package db

import (
	models "invokes/internal/models"
	"invokes/internal/utils"
)

// AddInvoice to add an invoice
func (db *Wrapper) AddInvoice(invoice *models.Invoice) error {
	if invoice == nil {
		return &DBError{Msg: "invoice struct isn't valid", Type: utils.InvalidContent}
	}

	invoice.Status = "pending"

	newinvoice := models.Invoice{}
	whereparams := models.Invoice{
		UserID: invoice.UserID,
		Label:  invoice.Label,
		Amount: invoice.Amount,
	}
	result := db.GormDB.Where(&whereparams).Assign(invoice).FirstOrCreate(&newinvoice)

	// Means already exist before ask to add
	if result.RowsAffected == 0 {
		return &DBError{Msg: "invoice already exists", Type: utils.AlreadyExist}
	}
	return nil
}

// GetInvoices returns the list of Invoices
func (db *Wrapper) GetInvoices() ([]*models.Invoice, int, error) {

	var invoices []*models.Invoice
	result := db.GormDB.Order("id").Find(&invoices)

	return invoices, len(invoices), result.Error
}
