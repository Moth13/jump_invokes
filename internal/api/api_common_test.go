package api

import (
	"encoding/json"
	"invokes/internal/utils"
	"net/http"
	"net/http/httptest"
	"testing"

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
