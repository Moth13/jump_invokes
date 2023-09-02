package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestInvoiceToJSON to test an unmarshalling of an incoming invoice
func TestInvoiceToJSON(t *testing.T) {
	invoice := Invoice{UserID: 1, Label: "test_invoice", Amount: 9745, AmountFloat: 97.45}

	// Check equality
	str := `{"user_id": 1, "label": "test_invoice", "amount": 97.45}`
	res := Invoice{}
	err := json.Unmarshal([]byte(str), &res)

	assert.Equal(t, nil, err)
	assert.Equal(t, invoice, res)

	// Check unequality
	str = `{"user_id": 2, "label": "test_invoice", "amount": 97.45}`
	res = Invoice{}
	err = json.Unmarshal([]byte(str), &res)

	assert.Equal(t, nil, err)
	assert.NotEqual(t, invoice, res)
}

// TestTransactionToJSON to test an unmarshalling of an incoming transaction

func TestTransactionToJSON(t *testing.T) {
	invoice := Transaction{InvoiceID: 1, Reference: "JMPINV200220117", Amount: 9745, AmountFloat: 97.45}

	// Check equality
	str := `{"invoice_id": 1, "reference": "JMPINV200220117", "amount": 97.45}`
	res := Transaction{}
	err := json.Unmarshal([]byte(str), &res)

	assert.Equal(t, nil, err)
	assert.Equal(t, invoice, res)

	// Check unequality
	str = `{"invoice_id": 2, "reference": "JMPINV200220117", "amount": 97.45}`
	res = Transaction{}
	err = json.Unmarshal([]byte(str), &res)

	assert.Equal(t, nil, err)
	assert.NotEqual(t, invoice, res)
}
