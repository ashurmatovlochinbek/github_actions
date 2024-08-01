package main

import (
	"encoding/json"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		byteValue, _ := json.Marshal("Hello")
		w.Write(byteValue)
	})
	http.ListenAndServe(":8080", nil)
}
