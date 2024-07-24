package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/yosa12978/mdpages/logging"
	"github.com/yosa12978/mdpages/session"
)

func init() {
	if os.Getenv("DEBUG") == "true" {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}
}

func Run() error {
	logger := logging.NewLogger(os.Stdin)

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()

	session.SetupStore()

	server := http.Server{
		Addr:    os.Getenv("ADDR"),
		Handler: NewRouter(ctx),
	}

	errCh := make(chan error, 1)
	go func() {
		logger.Info(fmt.Sprintf("server listening on %v", server.Addr))
		if err := server.ListenAndServe(); err != nil {
			errCh <- err
		}
		close(errCh)
	}()

	var err error
	select {
	case err = <-errCh:
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(
			context.Background(),
			10*time.Second,
		)
		defer cancel()
		err = server.Shutdown(timeout)
		logger.Info("server shutdown successfully")
	}
	return err
}
