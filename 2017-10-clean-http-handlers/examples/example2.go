package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/example1", GreeterWithError)
	http.ListenAndServe(":8080", nil)
}

func GreeterWithError(w http.ResponseWriter, r *http.Request) {
	err := fmt.Errorf("could not connect to database") // db.Query(...)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	fmt.Fprint(w, "Hola mundo!")
}
