package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"github.com/cpanato/ditos_gauchos/ditos"
	"github.com/cpanato/ditos_gauchos/handler"
)

const (
	addr = ":8080"
)

func main() {
	ditos, err := ditos.New()
	if err != nil {
		log.Fatalf("failed to initialize ditos: %v\n", err)
	}

	handler := handler.New(ditos)
	router := mux.NewRouter()
	router.HandleFunc("/bah", handler.HandleBah).Methods("POST")

	srv := http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		log.Printf("listening on %s\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("failed to start server: %v\n", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig

	log.Println("interrupted, shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("failed to gracefully shutdown: %v\n", err)
	}
}
