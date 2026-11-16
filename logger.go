package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type closeFunc func() error

func initializeLogger() (*log.Logger, closeFunc, error) {
	name, set := os.LookupEnv("LINKO_LOG_FILE")
	if !set {
		return log.New(os.Stderr, "", log.LstdFlags), func() error { return nil }, nil
	}

	file, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, func() error { return nil }, fmt.Errorf("creating log file: %w", err)
	}

	logger := log.New(io.MultiWriter(os.Stderr, file), "", log.LstdFlags)
	return logger, file.Close, nil
}
