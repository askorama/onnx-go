// This file contains specific handler to run the demos

package main

import (
	"io"
	"net/http"
	"os"
)

func init() {
	http.HandleFunc("/image", imagePostHandler)
	http.HandleFunc("/model", imagePostHandler)
}

func imagePostHandler(w http.ResponseWriter, r *http.Request) {
	io.Copy(os.Stdout, r.Body)
}
func modelPostHandler(w http.ResponseWriter, r *http.Request) {
}
