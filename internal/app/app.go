package app

import (
	"context"
	"fmt"
	"os"

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

func(a *App) Run() {
	var testsCount int
	fmt.Print("how many tests should I run?: ")
    fmt.Fscan(os.Stdin, &testsCount)

	ctx, cancel := context.WithCancel(context.Background())
	execCh := make(chan func() error, maxChanBuffer)
	sucCh := make(chan struct{}, maxChanBuffer)
	errCh := make(chan error, maxChanBuffer)
	defer func() {
		close(execCh)
		close(sucCh)
		close(errCh)
	}()

	for i := 0; i < a.maxFlows; i++ {
		go flowExecuter(ctx, execCh, sucCh, errCh)
	}

	for i := 0; i < testsCount; i++ {
		execCh <- gorgonaTests.QuizTest
	}

	count := 0
	for count <= testsCount {
		select {
		case <- sucCh:
			count++
		case err := <- errCh:
			fmt.Printf("test error: %v", err)
			count++
		}
	}

	cancel()
}

func flowExecuter (ctx context.Context, execute <-chan func() error, success chan <- struct{}, errors chan <- error) {
	for {
		select {
		case <- ctx.Done():
			return
		case f := <- execute:
			err := f() 
			if err != nil {
				errors <- err
			} else {
				success <- struct{}{}
			}
		}
	}
}
