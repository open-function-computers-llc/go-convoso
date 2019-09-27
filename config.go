package convoso

import (
	"errors"

	"github.com/sirupsen/logrus"
)

// Config a struct that we can fill from an external location to bootstrap this
// package
type Config struct {
	APIKey     string
	Logger     *logrus.Logger
	ListMapper map[string]int
}

func (c Config) validate() error {
	if c.APIKey == "" {
		return errors.New("The Convoso API Key must be set")
	}

	return nil
}
