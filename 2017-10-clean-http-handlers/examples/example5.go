package main

import "net/http"

func main() {
	http.HandleFunc("/example1", ComplexHandler)
	http.ListenAndServe(":8080", nil)
}

// START 1 OMIT
func ComplexHandler(w http.ResponseWriter, r *http.Request) {
	// Do something
	_, err := r.GetBody()

	// check for err
	if err != nil {
		http.Error(w, err, http.StatusInternalServerError)
		return
	}

	// Do something // HLxxx
	// ... // HLxxx
	// Rinse, repeat // HLxxx
}

// END 1 OMIT
