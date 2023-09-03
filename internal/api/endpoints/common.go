//common.go

package endpoints

import (
	handlers "invokes/internal/api/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping godoc
// @Tags ping
// @ID ping
// @Summary Do a ping
// @Description If alive, answer a pong
// @Produce json
// @Success 200 {string} string "pong"
// @Router /ping [get]
func Ping(e *handlers.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	}
}

// VersionStruct to return application version
type VersionStruct struct {
	Version string `example:"9.17.84" json:"version"`
}

// GetVersion godoc
// @Tags version
// @ID version
// @Summary Get current RestApi version
// @Produce json
// @Success 200 {object} VersionStruct "Current RestApi version"
// @Router /version [get]
func GetVersion(e *handlers.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		var v VersionStruct
		v.Version = e.Version
		c.JSON(http.StatusOK, v)
	}
}
