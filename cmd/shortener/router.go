package main

import (
	"github.com/DimKa163/go-shortener/internal/app"
	"net/http"
)

func run() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.Handler)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		return err
	}
	return nil
}
