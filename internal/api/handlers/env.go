// env.go

package env

import (
	models "invokes/internal/models"
)

type Env struct {
	Config  *models.Config
	Version string
}
