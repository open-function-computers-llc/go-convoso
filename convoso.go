package convoso

import "github.com/Sirupsen/logrus"

var log *logrus.Logger

// Init initialize this package. This must be called once with a valid instance of a Config struct
func Init(c Config) error {
	err := c.validate()
	if err != nil {
		return err
	}
	return nil
}
