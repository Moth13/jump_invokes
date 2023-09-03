package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"invokes/internal/models"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestPostTransaction to test the post of a transaction
func TestPostTransaction(t *testing.T) {
	a := InitializeTestApp()

	rand.Seed(time.Now().UnixNano())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users?user_id=1", nil)
	a.Router.RouterGin.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var users models.Users
	err := json.Unmarshal(w.Body.Bytes(), &users)
	assert.Equal(t, nil, err)
	user := users[0]

	amount_random := float64(rand.Intn(200000-5137)+5137) / 100

	// Post the invoice
	invoice := models.Invoice{UserID: user.UserID, AmountFloat: amount_random, Label: "Work for September"}
	json_data, err := json.Marshal(invoice)

	assert.Equal(t, nil, err)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/invoice", bytes.NewBuffer(json_data))
	a.Router.RouterGin.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)

	// Get the next invoice id
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", fmt.Sprintf("/invoices?user_id=%d&amount=%f&label=%s",
		user.UserID,
		amount_random,
		"Work%20for%20September"), nil)
	a.Router.RouterGin.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var invoices models.Invoices
	err = json.Unmarshal(w.Body.Bytes(), &invoices)

	assert.Equal(t, nil, err)

	new_invoice := invoices[len(invoices)-1]
	assert.Equal(t, invoice.UserID, new_invoice.UserID)
	assert.Equal(t, invoice.Label, new_invoice.Label)

	// Test Amount not the same
	transaction := models.Transaction{InvoiceID: new_invoice.ID, AmountFloat: 150, Reference: "JMPINV200220117"}
	json_data, err = json.Marshal(transaction)

	assert.Equal(t, nil, err)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/transaction", bytes.NewBuffer(json_data))
	a.Router.RouterGin.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Test Paid it
	transaction = models.Transaction{InvoiceID: new_invoice.ID, AmountFloat: amount_random, Reference: "JMPINV200220117"}
	json_data, err = json.Marshal(transaction)

	assert.Equal(t, nil, err)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/transaction", bytes.NewBuffer(json_data))
	a.Router.RouterGin.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)

	// Test Paid it again
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/transaction", bytes.NewBuffer(json_data))
	a.Router.RouterGin.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)

	// Test Balance User after
	w = httptest.NewRecorder()

	req, _ = http.NewRequest("GET", "/users?user_id=1", nil)
	a.Router.RouterGin.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var new_users models.Users
	err = json.Unmarshal(w.Body.Bytes(), &new_users)

	assert.Equal(t, nil, err)
	assert.Equal(t, user.BalanceFloat+amount_random, new_users[0].BalanceFloat)

	// Test the invoice status
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", fmt.Sprintf("/invoices?invoiceID=%d", new_invoice.ID), nil)
	a.Router.RouterGin.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var new_invoices models.Invoices
	err = json.Unmarshal(w.Body.Bytes(), &new_invoices)

	assert.Equal(t, nil, err)
	assert.Equal(t, "paid", new_invoices[0].Status)
}
