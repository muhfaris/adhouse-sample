package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/muhfaris/adhouse-sample/configs"
	"github.com/muhfaris/adhouse-sample/router"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func serveAPI(ctx context.Context) {
	// init config
	config := configs.NewConfig()
	config.InitializeConnectionPSQL()

	// cors
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "OPTIONS", "PATCH", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	router.HandlerV1(config, api)

	s := &http.Server{
		Addr:        fmt.Sprintf(":%d", config.Port),
		Handler:     cors(r),
		ReadTimeout: time.Duration(config.HTTP.ReadTimeout) * time.Minute,
	}

	done := make(chan struct{})
	go func() {
		<-ctx.Done()
		if err := s.Shutdown(context.Background()); err != nil {
			logrus.Error(err)
		}
		close(done)
	}()

	logrus.Infof("serving api at http://127.0.0.1:%d", config.Port)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		logrus.Error(err)
	}
	<-done
}

// this for gratefully shutdown
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serves the api",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, os.Interrupt)
			<-ch
			logrus.Info("signal caught. shutting down...")
			cancel()
		}()

		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			defer cancel()
			serveAPI(ctx)
		}()

		wg.Wait()
		return nil
	},
}
