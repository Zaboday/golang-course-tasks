package handler

import (
	"context"
	"fmt"
	"net/http"
)

type HttpHandler struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func New(ctx context.Context, cancel context.CancelFunc) *HttpHandler {
	return &HttpHandler{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (h *HttpHandler) Text(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Text!")
}

func (h *HttpHandler) Stat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Stat!")
}

func (h *HttpHandler) Stop(w http.ResponseWriter, r *http.Request) {
	h.cancel()

	fmt.Fprintf(w, "Server Stop!")
}
