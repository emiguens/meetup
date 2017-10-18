package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/example1", GreeterWithJSON)
	http.ListenAndServe(":8080", nil)
}

// START 1 OMIT
func GreeterWithJSON(w http.ResponseWriter, r *http.Request) {
	res := struct {
		Data interface{} `json:"data"`
	}{
		Data: "Respuesta correcta",
	}

	b, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	// Send marshaled JSON to ResponseWriter
	w.Write(b) // HLxxx
}

// END 1 OMIT
