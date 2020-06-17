package datahub

import "fmt"

// Config provides a configuration struct to connect to Alameda-Datahub
type Config struct {
	Address string `mapstructure:"address"`
}

// NewDefaultConfig returns default configuration
func NewDefaultConfig(ns string) Config {
	return Config{
		Address: fmt.Sprintf("alameda-datahub.%s.svc.cluster.local:50050", ns),
	}
}
