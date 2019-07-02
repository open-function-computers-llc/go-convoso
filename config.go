package convoso

import "errors"

// Config a struct that we can fill from an external location to bootstrap this
// package
type Config struct {
	APIKey string
}

func (c Config) validate() error {
	if c.APIKey == "" {
		return errors.New("The Convoso API Key must be set")
	}

	return nil
}
