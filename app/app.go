package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	if os.Getenv("DEBUG") == "true" {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}
}

func Run() error {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()
	server := http.Server{
		Addr:    os.Getenv("ADDR"),
		Handler: NewRouter(ctx),
	}

	errCh := make(chan error, 1)
	go func() {
		log.Printf("server listening on %v\n", server.Addr)
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
	}
	return err
}
