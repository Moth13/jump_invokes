package api

import (
	"bytes"
	"encoding/json"
	"invokes/internal/models"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestPostInvoice to test the post of an invoice
func TestPostInvoice(t *testing.T) {
	a := InitializeTestApp()

	rand.Seed(time.Now().UnixNano())

	// Get first user
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users?user_id=1", nil)
	a.Router.RouterGin.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var users models.Users
	err := json.Unmarshal(w.Body.Bytes(), &users)
	assert.Equal(t, nil, err)
	user := users[0]

	// Test wrong float amount
	amount_random := float64(rand.Intn(200000-5137)+5137) / 1000

	invoice := models.Invoice{UserID: user.UserID, AmountFloat: amount_random, Label: "Work for April"}
	json_data, err := json.Marshal(invoice)

	assert.Equal(t, nil, err)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/invoice", bytes.NewBuffer(json_data))
	a.Router.RouterGin.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Test correct amount
	amount_random = float64(rand.Intn(200000-5137)+5137) / 100

	invoice = models.Invoice{UserID: user.UserID, AmountFloat: amount_random, Label: "Work for April"}
	json_data, err = json.Marshal(invoice)

	assert.Equal(t, nil, err)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/invoice", bytes.NewBuffer(json_data))
	a.Router.RouterGin.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)

	// Repost the same
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/invoice", bytes.NewBuffer(json_data))
	a.Router.RouterGin.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)

	// Repost on a unvalid user
	w = httptest.NewRecorder()
	invoice = models.Invoice{UserID: 1984, AmountFloat: amount_random, Label: "Work for April"}
	json_data, err = json.Marshal(invoice)

	assert.Equal(t, nil, err)
	req, _ = http.NewRequest("POST", "/invoice", bytes.NewBuffer(json_data))
	a.Router.RouterGin.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
