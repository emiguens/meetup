package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/example1", GreeterWithErrorOK)
	http.ListenAndServe(":8080", nil)
}

// START 1 OMIT
func GreeterWithErrorOK(w http.ResponseWriter, r *http.Request) {
	err := fmt.Errorf("could not connect to database") // db.Query(...)
	if err != nil {
		// If we write a stream before setting status code, if defaults to 200 OK
		w.WriteHeader(http.StatusInternalServerError) // HLxxx
		fmt.Fprintf(w, err.Error())                   // HLxxx
		return                                        // HLxxx
	}

	fmt.Fprint(w, "Hola mundo!")
}

// END 1 OMIT
