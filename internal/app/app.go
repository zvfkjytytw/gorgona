package app

import (
	"context"
	"fmt"
	"os"
	"time"

	gorgonaTests "github.com/zvfkjytytw/gorgona/internal/tests"
)

const maxChanBuffer = 5

type App struct {
	maxFlows int
}

func New(maxFlows int) (*App, error) {
	return &App{
		maxFlows: maxFlows,
	}, nil
}

func (a *App) Run() {
	// Create control structure for tests
	seleniumTests, err := gorgonaTests.New()
	if err != nil {
		fmt.Printf("Test failed: %v", err)
		return
	}
	defer seleniumTests.Stop()

	// Asking for the number of tests to run
	var testsCount int
	fmt.Print("how many tests should I run?: ")
	fmt.Fscan(os.Stdin, &testsCount)

	// Creating a context and channels for running tests
	ctx, cancel := context.WithCancel(context.Background())
	execCh := make(chan func() error, maxChanBuffer)
	sucCh := make(chan struct{}, maxChanBuffer)
	errCh := make(chan error, maxChanBuffer)
	defer func() {
		close(execCh)
		close(sucCh)
		close(errCh)
	}()

	// Running gorutines
	for range a.maxFlows {
		go flowExecuter(ctx, execCh, sucCh, errCh)
	}

	// Running tsts
	for range testsCount {
		execCh <- seleniumTests.TestQuiz
	}

	// Checking exit statuses
	count := 0
	success := true
	for count < testsCount {
		select {
		case <-sucCh:
			count++
		case err := <-errCh:
			success = false
			fmt.Printf("test error: %v\n", err)
			count++
		}
	}

	cancel()
	time.Sleep(time.Second)
	if success {
		fmt.Println("All tests successful")
	}
}

// Run the functions from the execute channel and distribute the launch results to the appropriate channels
func flowExecuter(ctx context.Context, execute <-chan func() error, success chan<- struct{}, errors chan<- error) {
	for {
		select {
		case <-ctx.Done():
			return
		case f := <-execute:
			err := f()
			if err != nil {
				errors <- err
			} else {
				success <- struct{}{}
			}
		}
	}
}
