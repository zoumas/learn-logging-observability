package main

import (
	"bufio"
	"errors"
	"fmt"
	"log/slog"
	"os"
)

type closeFunc func() error

func initializeLogger(filename string) (*slog.Logger, closeFunc, error) {
	debugHandler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	if filename == "" {
		return slog.New(debugHandler), func() error { return nil }, nil
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, func() error { return nil }, fmt.Errorf("creating log file: %w", err)
	}

	bufferedFile := bufio.NewWriterSize(file, 8192)

	infoHandler := slog.NewTextHandler(bufferedFile, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	multiHandler := slog.NewMultiHandler(debugHandler, infoHandler)

	logger := slog.New(multiHandler)
	return logger, func() error {
		return errors.Join(bufferedFile.Flush(), file.Close())
	}, nil
}
