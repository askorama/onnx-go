package main

import (
	"log"
	"net/http"
)

func wasmHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/wasm")
	http.ServeFile(w, r, "emotions.wasm")
}
func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))
	mux.HandleFunc("/emotions.wasm", wasmHandler)
	log.Fatal(http.ListenAndServe(":3000", mux))
}
