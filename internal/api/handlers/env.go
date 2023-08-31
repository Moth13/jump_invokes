// env.go

package env

import (
	"invokes/internal/config"
)

type Env struct {
	Config  *config.Config
	Version string
}
