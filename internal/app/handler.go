package app

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		path := r.URL.Path
		value, err := GetUrl(path)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Location", value)
		w.WriteHeader(http.StatusTemporaryRedirect)
	case http.MethodPost:
		defer r.Body.Close()
		url, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		shortUrl, err := CreateShortUrl(string(url), 5)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		body := []byte(fmt.Sprintf("http://%s/%s", r.Host, shortUrl))
		_, err = w.Write(body)
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Length", strconv.Itoa(len(body)))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
