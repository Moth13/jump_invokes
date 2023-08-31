package utils

import (
	"path/filepath"
	"testing"
)

// TestConfig to test the function LoadConfiguration
func TestConfig(t *testing.T) {
	fileConfig, _ := filepath.Abs("../../configs/invokes.yml.sample")
	_, err := LoadConfiguration(&fileConfig)
	if err != nil {
		t.Errorf("Failed to load invoke sample conf %s", err.Error())
		return
	}
}
