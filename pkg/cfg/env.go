package cfg

import (
	"fmt"
	"github.com/Netflix/go-env"
	"github.com/rs/zerolog/log"
)

// Vars environment variables
var Vars environment

type environment struct {
	LogLevel string `env:"GOTHY_LOG_LEVEL,default=info"`
	LogColor bool   `env:"GOTHY_LOG_COLOR,default=false"`

	Env string `env:"GOTHY_ENV,default=dev"`
}

// LoadEnv load env variables
func LoadEnv() error {
	_, err := env.UnmarshalFromEnviron(&Vars)
	if err != nil {
		return fmt.Errorf("unable to unmarshall env variables: %w", err)
	}
	return nil
}

// PrepareEnv global environment preparation
func PrepareEnv() (err error) {
	if err = LoadEnv(); err != nil {
		return
	}
	if err = ConfigureLogging(); err != nil {
		return
	}

	log.Trace().Msg("environment prepared")

	return
}
