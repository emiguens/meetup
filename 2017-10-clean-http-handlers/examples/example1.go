package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/example1", SimpleGreeter)
	http.ListenAndServe(":8080", nil)
}

func SimpleGreeter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hola mundo!")
}

