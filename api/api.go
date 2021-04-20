package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// Api stateless type
type Api int

// Run Api server
// server is started within a go routine and wait against the context
func (p *Api) Run(ctx context.Context, port int) {
	server := http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 16,
		Handler:        http.HandlerFunc(helloWorld),
	}

	// shutdown the server if context returns
	defer func(delay time.Duration) {
		ctx, cancel := context.WithTimeout(context.Background(), delay)
		defer cancel()

		err := server.Shutdown(ctx)
		if err != nil {
			log.WithError(err).
				Fatal("Failed to shutdown http server")
			return
		}

		log.Info("Pong API exited")
	}(5 * time.Second)

	// start the server within a go routine
	go func() {
		log.WithField("Addr", server.Addr).
			Info("Pong API is starting")

		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.WithError(err).
				Fatal("Failed to start http server")
		}
	}()

	<-ctx.Done()
}
