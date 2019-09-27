package convoso

import "github.com/sirupsen/logrus"

var log *logrus.Logger
var apiKEY string
var listMapper map[string]int

// Init initialize this package. This must be called once with a valid instance of a Config struct
func Init(c Config) error {
	err := c.validate()
	if err != nil {
		return err
	}

	apiKEY = c.APIKey
	log = c.Logger
	listMapper = c.ListMapper

	return nil
}
