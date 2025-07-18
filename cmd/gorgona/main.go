package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	gorgonaApp "github.com/zvfkjytytw/gorgona/internal/app"
)

func main() {
	var (
		// Maximum parallel flows
		maxGorutines int
		// Env variables of the maximum parallel flows
		envMaxGorutines string = "MAX_FLOWS"
	)

	flag.IntVar(&maxGorutines, "f", 3, "Maximum parallel flows")
	flag.Parse()

	envF, ok := os.LookupEnv(envMaxGorutines)
	if ok {
		value, err := strconv.Atoi(envF)
		if err == nil {
			maxGorutines = value
		}
	}

	app, err := gorgonaApp.New(maxGorutines)
	if err != nil {
		fmt.Printf("test failed")
		os.Exit(1)
	}

	app.Run()
}
