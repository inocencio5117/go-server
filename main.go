package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()
		next.ServeHTTP(w, r)
		fmt.Println(r.URL.String(), r.Method, time.Since(begin))
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	},
	)

	mux.HandleFunc("POST /api/echo/{text}", func(w http.ResponseWriter, r *http.Request) {
		text := r.PathValue("text")
		fmt.Fprintln(w, text)
	})

	srv := &http.Server{
		Addr:                         ":8080",
		Handler:                      Log(mux),
		DisableGeneralOptionsHandler: false,
		ReadTimeout:                  10 * time.Second,
		WriteTimeout:                 10 * time.Second,
		IdleTimeout:                  1 * time.Minute,
	}

	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}
}
