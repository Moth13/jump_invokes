package utils

import (
	models "invokes/internal/models"
	"io/ioutil"
	"net/http"

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

// Enum to define db error
const (
	InvalidContent = 100 // Error returned if data isn't correct
	AlreadyExist   = 101 // Error returned when data already exist

	InvoiceAlreadyPaid    = 200
	InvoiceAmountNotFound = 201
	InvoiceNotFound       = 203
)

// DBCodeToHTTPCode convert a db error code into the wanted http status code
func DBCodeToHTTPCode(dbcode int) int {
	httpcode := http.StatusNoContent
	switch dbcode {
	case InvoiceAlreadyPaid:
		httpcode = http.StatusUnprocessableEntity
	case InvoiceAmountNotFound:
		httpcode = http.StatusBadRequest
	case InvoiceNotFound:
		httpcode = http.StatusNotFound
	case InvalidContent:
		httpcode = http.StatusBadRequest
	case AlreadyExist:
		httpcode = http.StatusConflict
	default:
	}
	return httpcode
}
