package cfg

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

// ConfigureLogging configure the logging - set global log level
func ConfigureLogging() (err error) {
	level, err := zerolog.ParseLevel(Vars.LogLevel)
	if err != nil {
		log.Warn().Err(err).Msg("unable to parse log level")
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)

	if Vars.LogColor {
		useColoredLogger()
	}

	return nil
}

func useColoredLogger() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}
