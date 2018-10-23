package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/cors"
)

var nhac *ChiaSeNhac

// NewServer creates a new HTTP handler with default http client
func NewServer() http.Handler {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	return NewServerWithClient(client)
}

// NewServerWithClient creates a new HTTP handler with the specifed http client
func NewServerWithClient(c HttpClient) http.Handler {
	nhac = NewCsn(c)

	mux := http.NewServeMux()
	mux.HandleFunc("/media", require("link", handleGetLink))
	mux.HandleFunc("/detail", require("link", handleGetTrack))
	mux.HandleFunc("/tracks", require("link", handleGetList))
	mux.HandleFunc("/preset", require("link", handleGetPreset))

	return cors.Default().Handler(mux)
}

func handleGetList(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return nhac.GetList(r.URL.Query().Get("link"))
}

func handleGetPreset(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return nhac.GetPreset(r.URL.Query().Get("link"))
}

func handleGetTrack(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return nhac.GetTrack(r.URL.Query().Get("link"))
}

func handleGetLink(w http.ResponseWriter, r *http.Request) (link interface{}, e error) {
	media, e := nhac.GetMedia(r.URL.Query().Get("link"))
	if e == nil {
		media = fmt.Sprintf(`{"link": "%s"}`, link)
		link = map[string]string{"link": media}
	}
	return
}

func require(param string, handler func(w http.ResponseWriter, r *http.Request) (interface{}, error)) http.HandlerFunc {
	return withLog(func(w http.ResponseWriter, r *http.Request) {
		requestLink := r.URL.Query().Get(param)

		if requestLink == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("content-type", "application/json")
		data, e := handler(w, r)

		if e != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"error": "%s"}`, e.Error())
		} else {
			bytes, _ := json.Marshal(data)
			w.Write(bytes)
		}
	})
}

func withLog(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s: %s\n", r.Method, r.URL.String())
		handler(w, r)
	}
}
