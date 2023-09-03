package api

import (
	handlers "invokes/internal/api/handlers"
	models "invokes/internal/models"
	"invokes/internal/utils"

	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// LoadConfiguration loads a json config file into a Config struct.
func LoadConfiguration(file *string) (models.Config, error) {
	var config models.Config
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

// App struct describes the application
type App struct {
	Env    handlers.Env
	Router Router

	DefaultTTL string
	ErrorTTL   string
}

// Initialize configure the app
func (a *App) Initialize(config *models.Config) error {
	err := a.Env.DB.Initialize(config)
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
