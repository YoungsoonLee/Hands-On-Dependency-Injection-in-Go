// +build bad

package config

import (
	"github.com/corsc/Hands-On-Dependency-Injection-in-Go/ch01/02_code_smells/circular_dependencies/payment"
	"errors"
)

// Config defines the JSON format of the config file
type Config struct {
	// Address is the host and port to bind to.
	// Default 0.0.0.0:8080
	Address string

	// DefaultCurrency is the default currency of the system
	DefaultCurrency payment.Currency
}

// Load will load the JSON config from the file supplied
func Load(filename string) (*Config, error) {
	// TODO: load currency from file
	return nil, errors.New("not implemented yet")
}
