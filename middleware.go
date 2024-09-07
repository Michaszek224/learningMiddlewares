package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func jsonHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("Before json header", time.Now())
		next.ServeHTTP(w, r)
		fmt.Println("After json header", time.Now())
	})
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Path and so on %v %v %v", r.URL, r.Proto, r.Method)
		fmt.Println("Before logger", time.Now())
		next.ServeHTTP(w, r)
		fmt.Println("After logger", time.Now())
	})
}
