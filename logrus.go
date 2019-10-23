package main

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {

	// log.SetReportCaller(true)
	log.SetOutput(os.Stdout)

	consoleFormat() // dev environment
	jsonFormat()    // prod environment
}

func consoleFormat() {
	log.WithFields(log.Fields{
		"url":     "http://foo.com",
		"attempt": 3,
		"backoff": time.Second,
	}).Info("failed to fetch URL")
}

func jsonFormat() {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(log.Fields{
		"url":     "http://foo.com",
		"attempt": 3,
		"backoff": time.Second,
	}).Info("failed to fetch URL")
}
