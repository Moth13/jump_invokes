package db

import (
	models "invokes/internal/models"
	"invokes/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCheckAndUpdateInvoice to test the function TestCheckAndUpdateInvoice
func TestCheckAndUpdateInvoice(t *testing.T) {
	user := models.User{UserID: 1, FirstName: "John", LastName: "Doe", Balance: 19255, BalanceFloat: 192.55}
	invoice := models.Invoice{ID: 1, UserID: 1, Label: "test_invoice", Amount: 9745, AmountFloat: 97.45, User: user, Status: "pending"}
	transaction := models.Transaction{InvoiceID: 1, Reference: "JMPINV200220117", Amount: 9645, AmountFloat: 96.45}

	// check as amount aren't the same
	assert.Equal(t, &DBError{Msg: "Invoice amount isn't the same", Type: utils.InvoiceAmountNotFound}, checkAndUpdateInvoice(&invoice, &transaction))

	transaction.Amount = 9745
	transaction.AmountFloat = 97.45

	assert.Equal(t, nil, checkAndUpdateInvoice(&invoice, &transaction))
	assert.Equal(t, "paid", invoice.Status)
	assert.Equal(t, int32(29000), invoice.User.Balance)

	// invoice status has been passed to paid, check to set transaction again
	assert.Equal(t, &DBError{Msg: "Invoice has already been paid", Type: utils.InvoiceAlreadyPaid}, checkAndUpdateInvoice(&invoice, &transaction))
}
