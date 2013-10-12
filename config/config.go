// Package config implements routines to read and structure
// the options for Moln.
package config

import (
	"encoding/json"
	"os"
	"time"
)

// TLS describes the cert/key options for a TLS connection.
type TLS struct {
	Cert string
	Key  string
}

// Config describes the options for Moln.
type Config struct {
	LogDir              string
	DBAddr              string
	DBNetwork           string
	DBMaxIdle           int
	DBMaxTimeoutStr     string
	DBMaxTimeout        time.Duration
	ServerAddr          string
	ServerNetwork       string
	ServerMaxTimeoutStr string
	ServerMaxTimeout    time.Duration
	TLS                 *TLS
}

// ReadFiles reads the given JSON config files and returns the combined config.
func ReadFiles(files ...string) (*Config, error) {
	var err error
	config := new(Config)

	for _, file := range files {
		fileF, err := os.Open(file)
		if err != nil {
			return nil, err
		}
		defer fileF.Close()

		decoder := json.NewDecoder(fileF)
		decoder.Decode(config)
	}

	config.DBMaxTimeout, err = time.ParseDuration(config.DBMaxTimeoutStr)
	config.ServerMaxTimeout, err = time.ParseDuration(config.ServerMaxTimeoutStr)
	return config, err
}
