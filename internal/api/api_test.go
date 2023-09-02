package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"invokes/internal/models"
	"invokes/internal/utils"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func InitializeTestApp() *App {
	confFile := "../../configs/invokes.yml.test"
	config, _ := utils.LoadConfiguration(&confFile)
	utils.LoadLogger(&config)
	a := App{}
	a.Initialize(&config)
	return &a
}

// TestVersion to test the version endpoint
func TestVersion(t *testing.T) {
	a := InitializeTestApp()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/version", nil)
	a.Router.RouterGin.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var version map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &version)

	assert.Equal(t, nil, err)
	assert.Equal(t, map[string]string{"version": a.Env.Version}, version)
}

// TestUser to test the users endpoint
func TestUser(t *testing.T) {
	a := InitializeTestApp()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/users", nil)
	a.Router.RouterGin.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var users models.Users
	err := json.Unmarshal(w.Body.Bytes(), &users)

	assert.Equal(t, nil, err)
	assert.Equal(t, 17, len(users))
	// Check a random users
	assert.Equal(t, 6, users[5].UserID)
	assert.Equal(t, "Oscar", users[6].FirstName)
	assert.Equal(t, float64(5200.6), users[7].BalanceFloat)
}

// TestPostInvoice to test the post of an invoice
func TestPostInvoice(t *testing.T) {
	a := InitializeTestApp()

	rand.Seed(time.Now().UnixNano())

	// Test wrong float amount
	amount_random := float64(rand.Intn(200000-5137)+5137) / 1000

	invoice := models.Invoice{UserID: 2, AmountFloat: amount_random, Label: "Work for April"}
	json_data, err := json.Marshal(invoice)

	assert.Equal(t, nil, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/invoice", bytes.NewBuffer(json_data))
	a.Router.RouterGin.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Test correct amount
	amount_random = float64(rand.Intn(200000-5137)+5137) / 100

	invoice = models.Invoice{UserID: 2, AmountFloat: amount_random, Label: "Work for April"}
	json_data, err = json.Marshal(invoice)

	assert.Equal(t, nil, err)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/invoice", bytes.NewBuffer(json_data))
	a.Router.RouterGin.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)

	// repost the same
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/invoice", bytes.NewBuffer(json_data))
	a.Router.RouterGin.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)
}

// TestPostTransaction to test the post of a transaction
func TestPostTransaction(t *testing.T) {
	a := InitializeTestApp()

	rand.Seed(time.Now().UnixNano())

	amount_random := float64(rand.Intn(200000-5137)+5137) / 100

	// Post the invoice
	invoice := models.Invoice{UserID: 3, AmountFloat: amount_random, Label: "Work for September"}
	json_data, err := json.Marshal(invoice)

	assert.Equal(t, nil, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/invoice", bytes.NewBuffer(json_data))
	a.Router.RouterGin.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)

	// Get the next invoice id
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", fmt.Sprintf("/invoices?user_id=3&amount=%f&label=%s", amount_random, "Work%20for%20September"), nil)
	a.Router.RouterGin.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var invoices models.Invoices
	err = json.Unmarshal(w.Body.Bytes(), &invoices)

	assert.Equal(t, nil, err)

	invoice_id := invoices[len(invoices)-1].ID

	// Test Amount not the same
	transaction := models.Transaction{InvoiceID: invoice_id, AmountFloat: 150, Reference: "JMPINV200220117"}
	json_data, err = json.Marshal(transaction)

	assert.Equal(t, nil, err)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/transaction", bytes.NewBuffer(json_data))
	a.Router.RouterGin.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Test Paid it
	transaction = models.Transaction{InvoiceID: invoice_id, AmountFloat: amount_random, Reference: "JMPINV200220117"}
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
}
