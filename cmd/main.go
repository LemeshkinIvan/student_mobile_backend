package main

import (
	"app/internal/bootstrap"
	"log"
	"net/http"
	"time"
)

func main() {
	app := bootstrap.InitApp()

	srv := &http.Server{
		Addr:           app.Config.BASE_DEV_URL,
		Handler:        app.Engine,
		ReadTimeout:    app.Config.READ_TIMEOUT * time.Second,
		WriteTimeout:   app.Config.WRITE_TIMEOUT * time.Second,
		IdleTimeout:    app.Config.IDLE_TIMEOUT * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("We fucked up at the start: %s\n", err)
	}
}
