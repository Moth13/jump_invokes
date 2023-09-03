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
	result := db.GormDB.Where(&whereparams).FirstOrCreate(&newinvoice, invoice)

	if result.Error != nil {
		return &DBError{Msg: result.Error.Error(), Type: utils.InvalidParams}
	}

	// Means already exist before ask to add
	if result.RowsAffected == 0 {
		return &DBError{Msg: "invoice already exists", Type: utils.AlreadyExist}
	}
	return nil
}

// GetInvoices returns the list of Invoices
func (db *Wrapper) GetInvoices(filter *models.Invoice) ([]*models.Invoice, int, error) {
	var invoices []*models.Invoice
	query := db.GormDB
	if filter != nil {
		query = query.Where(filter)
	}
	result := query.Order("id").Find(&invoices)

	return invoices, len(invoices), result.Error
}
