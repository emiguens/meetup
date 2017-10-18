package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// START 1 OMIT
func main() {
	port := os.Getenv("PORT")

	// Encapsulamos el handler custom en nuestro adapter
	http.Handle("/", Action(Index)) // HLxxx

	http.ListenAndServe(":"+port, nil)
}

func Index(r *http.Request) *Response {
	v, err := doSomething()
	if err != nil {
		return ErrorJSON(http.StatusServiceUnavailable, err, nil)
	}

	return DataJSON(http.StatusOK, v, nil)
}

func doSomething() ([]int, error) {
	return nil, fmt.Errorf("error fetching something")
}

// END 1 OMIT

// START 2 OMIT
func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		format := "[%s] User agent => %s Remote addr => %s"
		log.Printf(format, r.Method, r.UserAgent(), r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

// END 2 OMIT
