// env.go

package env

import (
	"invokes/internal/db"
	models "invokes/internal/models"
)

// Env wrapped some elements to give them through packages
type Env struct {
	Config  *models.Config
	DB      db.Wrapper
	Version string
}
