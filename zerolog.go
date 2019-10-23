package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const url = "http://foo.com"

func main() {
	// UNIX Time is faster and smaller than most timestamps
	// If you set zerolog.TimeFieldFormat to an empty string,
	// logs will write with UNIX time
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	jsonFormat()    // performances are crucial
	consoleFormat() // dev environment, performances are important but not crucial
}

func jsonFormat() {
	log.Info().
		Str("url", url).
		Int("attempt", 3).
		Dur("backoff", time.Second).
		Msg("Failed to fetch URL")
}

func consoleFormat() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().
		Str("url", url).
		Int("attempt", 3).
		Dur("backoff", time.Second).
		Msg("Failed to fetch URL")
}
