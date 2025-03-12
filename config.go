package server

import (
	"flag"
	"os"

	"github.com/caarlos0/env"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

type envs struct {
	Host string `env:"CONTAINER_HOST,required"`
	Port string `env:"CONTAINER_PORT,required"`
}

var Envs envs

func init() {
	init_logger()
	init_env()
}

func init_logger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// for using stack trace in log
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	// for pretty print
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	// for log level, if run `./main -debug`, log level is set to DEBUG, otherwise INFO.
	debug := flag.Bool("debug", false, "sets log level to debug")
	flag.Parse()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}

func init_env() {
	if err := env.Parse(&Envs); err != nil {
		log.Fatal().Err(err).Msg("failed to parse config")
	}
}
