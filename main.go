package main

import (
	"log"
	"net/http"
)

type handler struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(200)
}

func main() {
	h := handler{}

	s := &http.Server{
		Addr:    ":8080",
		Handler: h,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
