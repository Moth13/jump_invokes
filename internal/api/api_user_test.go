package api

import (
	"encoding/json"
	"invokes/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
