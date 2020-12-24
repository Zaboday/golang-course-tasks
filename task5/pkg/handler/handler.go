package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpHandler struct {
	ctx    context.Context
	cancel context.CancelFunc
}

type RequestParameters struct {
	Text   string `json:"text"`
	Number int    `json:"number"`
}

func New(ctx context.Context, cancel context.CancelFunc) *HttpHandler {
	return &HttpHandler{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (h *HttpHandler) Text(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Bas Method Request.", http.StatusBadRequest)
		return
	}

	p := RequestParameters{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Parameters: %+v", p)
}

func (h *HttpHandler) Stat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Stat!")
}

func (h *HttpHandler) Stop(w http.ResponseWriter, r *http.Request) {
	h.cancel()

	fmt.Fprintf(w, "Server Stop!")
}
