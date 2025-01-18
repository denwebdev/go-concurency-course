package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"key-value-database/internal/database/compute"
	"key-value-database/internal/database/storage"
)

const (
	logLevel = "debug"
)

func main() {
	logger, err := newLogger(logLevel)
	if err != nil {
		log.Fatal("create logger error", err)
	}
	logger.Debug("debug mode on")

	waitExitSignal(logger)

	engine := storage.NewInMemoryEngine(logger)
	computer := compute.NewCompute(&compute.CommandParser{}, engine)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		result, err := computer.Process(scanner.Text())
		if err != nil {
			fmt.Println("error: ", err)
			continue
		}
		fmt.Println(result)
	}
}

func waitExitSignal(logger *zap.Logger) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-signals
		logger.Info("service stop", zap.Stringer("signal", sig))
		os.Exit(0)
	}()
}

func newLogger(level string) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	atom := zap.NewAtomicLevel()
	err := atom.UnmarshalText([]byte(level))
	if err != nil {
		return nil, err
	}

	cfg.Level = atom

	return cfg.Build()
}
