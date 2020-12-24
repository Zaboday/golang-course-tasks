package main

import (
	"context"
	"log"
	"main/pkg/handler"
	"net/http"
)

func main() {
	StartHttpServer()
}

func StartHttpServer() {
	m := http.NewServeMux()
	s := http.Server{Addr: ":8080", Handler: m}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	h := handler.New(ctx, cancel)

	m.HandleFunc("/text", h.Text)
	m.HandleFunc("/stat", h.Stat)
	m.HandleFunc("/stop", h.Stop)

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	select {
	case <-ctx.Done():
		log.Println("Server Shutdown from context")
		s.Shutdown(ctx)
	}
}
