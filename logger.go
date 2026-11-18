package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
)

type closeFunc func() error

func initializeLogger(filename string) (*slog.Logger, closeFunc, error) {
	if filename == "" {
		return slog.New(slog.NewTextHandler(os.Stderr, nil)), func() error { return nil }, nil
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, func() error { return nil }, fmt.Errorf("creating log file: %w", err)
	}

	bufferedFile := bufio.NewWriterSize(file, 8192)

	logger := slog.New(slog.NewTextHandler(io.MultiWriter(os.Stderr, bufferedFile), nil))
	return logger, func() error {
		return errors.Join(bufferedFile.Flush(), file.Close())
	}, nil
}
