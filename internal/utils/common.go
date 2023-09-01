package utils

import (
	models "invokes/internal/models"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// LoadLogger loads logger from a config struct
func LoadLogger(config *models.Config) error {
	NewLogger(config)

	return nil
}

// LoadConfiguration loads a yaml config file into a Config struct.
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

// Contains checks if string is in a string list
func Contains(slist []string, value string) bool {
	for _, v := range slist {
		if v == value {
			return true
		}
	}
	return false
}
