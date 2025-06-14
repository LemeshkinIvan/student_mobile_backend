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
		Addr:           "fck_all",
		Handler:        app.Router.Engine,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    30 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("We fucked up at the start: %s\n", err)
	}
}
