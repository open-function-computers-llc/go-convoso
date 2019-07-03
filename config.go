package convoso

import (
	"errors"

	"github.com/Sirupsen/logrus"
)

// Config a struct that we can fill from an external location to bootstrap this
// package
type Config struct {
	APIKey string
	Logger *logrus.Logger
}

func (c Config) validate() error {
	if c.APIKey == "" {
		return errors.New("The Convoso API Key must be set")
	}

	return nil
}
