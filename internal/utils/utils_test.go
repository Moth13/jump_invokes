package utils

import (
	"net/http"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestContains to test the function Contains
func TestContains(t *testing.T) {
	entryList := []string{"a", "b", "c"}

	assert.Equal(t, true, Contains(entryList, "b"))
	assert.Equal(t, false, Contains(entryList, "e"))
}

// TestConfig to test the function LoadConfiguration
func TestConfig(t *testing.T) {
	fileConfig, _ := filepath.Abs("../../configs/invokes.yml.sample")
	_, err := LoadConfiguration(&fileConfig)
	if err != nil {
		t.Errorf("Failed to load invoke sample conf %s", err.Error())
		return
	}
}

// TestDBCodeConverter to test the function DBCodeToHTTPCode
func TestDBCodeConverter(t *testing.T) {
	assert.Equal(t, http.StatusUnprocessableEntity, DBCodeToHTTPCode(InvoiceAlreadyPaid))
	assert.Equal(t, http.StatusBadRequest, DBCodeToHTTPCode(InvoiceAmountNotFound))
	assert.Equal(t, http.StatusNotFound, DBCodeToHTTPCode(InvoiceNotFound))
	assert.Equal(t, http.StatusBadRequest, DBCodeToHTTPCode(InvalidContent))
	assert.Equal(t, http.StatusConflict, DBCodeToHTTPCode(AlreadyExist))
}
