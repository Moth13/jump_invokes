package api

import (
	handlers "invokes/internal/api/handlers"
	"invokes/internal/config"
	"invokes/internal/db"
	"invokes/internal/utils"

	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// LoadConfiguration loads a json config file into a Config struct.
func LoadConfiguration(file *string) (config.Config, error) {
	var config config.Config
	configFile, err := ioutil.ReadFile(*file)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}

type App struct {
	Env    handlers.Env
	Router Router
	DB     db.Wrapper

	DefaultTTL string
	ErrorTTL   string
}

func (a *App) Initialize(config *config.Config) error {
	err := a.DB.Initialize(config)
	if err != nil {
		return err
	}

	a.Env.Config = config

	a.Env.Version = Version

	a.Router.Initialize(&a.Env)

	a.DefaultTTL = fmt.Sprintf("max-age=%d", a.Env.Config.Caching.TTL)
	a.ErrorTTL = fmt.Sprintf("max-age=%d", a.Env.Config.Caching.ErrorTTL)

	return nil
}

// Run http server.
func (a *App) Run() {
	utils.Logger.Fatal(a.Router.Run())
}
