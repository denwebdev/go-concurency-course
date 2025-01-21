package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"key-value-database/internal/database"
	"key-value-database/internal/database/compute"
	"key-value-database/internal/database/storage"
	"key-value-database/internal/initialization"
)

const (
	logLevel = "debug"
)

func main() {
	logger, err := initialization.CreateLogger(logLevel)
	if err != nil {
		log.Fatal("create logger error", err)
	}
	logger.Debug("debug mode on")

	waitExitSignal()

	db := database.NewDatabase(&compute.CommandParser{}, storage.NewInMemoryEngine())

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		result, err := db.HandleQuery(scanner.Text())
		if err != nil {
			fmt.Println("error: ", err)
			continue
		}
		fmt.Println(result)
	}
}

func waitExitSignal() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signals
		os.Exit(0)
	}()
}
